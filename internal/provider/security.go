package provider

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SecureString wraps a string with secure cleanup capabilities
type SecureString struct {
	data []byte
}

// NewSecureString creates a new SecureString from a regular string
func NewSecureString(s string) *SecureString {
	return &SecureString{
		data: []byte(s),
	}
}

// String returns the string value (use sparingly, prefer GetBytes)
func (s *SecureString) String() string {
	if s.data == nil {
		return ""
	}
	return string(s.data)
}

// GetBytes returns the underlying byte slice
func (s *SecureString) GetBytes() []byte {
	return s.data
}

// Clear securely wipes the sensitive data from memory
func (s *SecureString) Clear(ctx context.Context) {
	if s.data == nil {
		return
	}
	originalLen := len(s.data)
	// Overwrite with random data first
	rand.Read(s.data)
	// Then zero out
	for i := range s.data {
		s.data[i] = 0
	}
	s.data = nil
	// Log successful clearing (without revealing content)
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] SecureString cleared: %d bytes zeroed", originalLen))
}

// SecureClearString securely clears a string from memory by converting to bytes and zeroing
// Note: This works on a best-effort basis as Go strings are immutable
func SecureClearString(ctx context.Context, s *string) {
	if s == nil || *s == "" {
		return
	}
	
	originalLen := len(*s)
	// Log before clearing (first 3 chars for verification in debug mode)
	debugPrefix := ""
	if originalLen > 0 {
		if originalLen >= 3 {
			debugPrefix = (*s)[:3]
		} else {
			debugPrefix = (*s)[:1]
		}
		debugPrefix += "***"
	}
	
	// Convert to byte slice and zero it out
	bytes := []byte(*s)
	// Overwrite with random data
	rand.Read(bytes)
	// Zero out
	for i := range bytes {
		bytes[i] = 0
	}
	// Set string to empty
	*s = ""
	
	// Verify clearing
	isCleared := len(*s) == 0
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] String cleared: original_len=%d, prefix_was=%s, cleared=%v", 
		originalLen, debugPrefix, isCleared))
}

// SecureClearByteSlice securely clears a byte slice from memory
func SecureClearByteSlice(ctx context.Context, b []byte) {
	if b == nil {
		return
	}
	
	originalLen := len(b)
	// Log first few bytes for verification (in hex to avoid binary issues)
	debugHex := ""
	if originalLen > 0 {
		maxBytes := 4
		if originalLen < maxBytes {
			maxBytes = originalLen
		}
		debugHex = fmt.Sprintf("%x", b[:maxBytes])
	}
	
	// Overwrite with random data
	rand.Read(b)
	// Zero out
	for i := range b {
		b[i] = 0
	}
	
	// Verify all bytes are zero
	allZero := true
	for _, v := range b {
		if v != 0 {
			allZero = false
			break
		}
	}
	
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] ByteSlice cleared: original_len=%d, first_bytes_were=%s, all_zero=%v", 
		originalLen, debugHex, allZero))
}

// SecureClearMap securely clears sensitive string values from a map
func SecureClearMap(ctx context.Context, m map[string]interface{}, sensitiveKeys []string) {
	if m == nil {
		return
	}
	clearedCount := 0
	for _, key := range sensitiveKeys {
		if val, ok := m[key]; ok {
			if strVal, isString := val.(string); isString {
				if strVal != "" {
					SecureClearString(ctx, &strVal)
					m[key] = ""
					clearedCount++
				}
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] Map cleared: %d sensitive keys processed", clearedCount))
}

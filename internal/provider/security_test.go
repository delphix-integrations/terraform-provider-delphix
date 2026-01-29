package provider

import (
	"context"
	"testing"
)

func TestSecureClearString(t *testing.T) {
	testPassword := "MySecretPassword123!"
	originalLen := len(testPassword)
	
	t.Logf("Original password length: %d", originalLen)
	t.Logf("Original password (first 3 chars): %s***", testPassword[:3])
	
	SecureClearString(context.Background(), &testPassword)
	
	if testPassword != "" {
		t.Errorf("Password was not cleared. Expected empty string, got: %s", testPassword)
	}
	
	if len(testPassword) != 0 {
		t.Errorf("Password length not zero after clearing. Got: %d", len(testPassword))
	}
	
	t.Logf("Password successfully cleared. Length after clear: %d", len(testPassword))
}

func TestSecureClearByteSlice(t *testing.T) {
	testData := []byte("SensitiveAPIKey12345")
	originalLen := len(testData)
	
	t.Logf("Original data length: %d", originalLen)
	t.Logf("Original data (hex): %x", testData[:4])
	
	SecureClearByteSlice(context.Background(), testData)
	
	// Verify all bytes are zero
	allZero := true
	for i, b := range testData {
		if b != 0 {
			allZero = false
			t.Errorf("Byte at index %d is not zero: %d", i, b)
		}
	}
	
	if !allZero {
		t.Errorf("Not all bytes were cleared to zero")
	} else {
		t.Logf("All %d bytes successfully cleared to zero", originalLen)
	}
}

func TestSecureClearMap(t *testing.T) {
	testMap := map[string]interface{}{
		"password":    "secret123",
		"api_key":     "key_12345",
		"safe_value":  "public_data",
		"access_key":  "aws_secret",
	}
	
	sensitiveKeys := []string{"password", "api_key", "access_key"}
	
	t.Logf("Before clear - password: %v", testMap["password"])
	
	SecureClearMap(context.Background(), testMap, sensitiveKeys)
	
	// Verify sensitive keys are cleared
	for _, key := range sensitiveKeys {
		if val, ok := testMap[key]; ok {
			if strVal, isString := val.(string); isString && strVal != "" {
				t.Errorf("Sensitive key '%s' was not cleared. Value: %s", key, strVal)
			}
		}
	}
	
	// Verify non-sensitive keys are untouched
	if testMap["safe_value"] != "public_data" {
		t.Errorf("Non-sensitive key was modified. Expected 'public_data', got: %v", testMap["safe_value"])
	}
	
	t.Logf("After clear - password: %v (should be empty)", testMap["password"])
	t.Log("Map clearing completed successfully")
}

func TestSecureString(t *testing.T) {
	original := "TestPassword456"
	secureStr := NewSecureString(original)
	
	t.Logf("SecureString created with length: %d", len(secureStr.GetBytes()))
	
	if secureStr.String() != original {
		t.Errorf("SecureString does not match original. Expected: %s, Got: %s", original, secureStr.String())
	}
	
	secureStr.Clear(context.Background())
	
	if secureStr.data != nil {
		t.Errorf("SecureString data was not set to nil after Clear()")
	}
	
	if secureStr.String() != "" {
		t.Errorf("SecureString should return empty string after Clear(). Got: %s", secureStr.String())
	}
	
	t.Log("SecureString cleared successfully")
}

func TestSecureClearNilValues(t *testing.T) {
	// Test nil string pointer
	var nilStr *string
	SecureClearString(context.Background(), nilStr) // Should not panic
	t.Log("Nil string pointer handled correctly")
	
	// Test empty string
	emptyStr := ""
	SecureClearString(context.Background(), &emptyStr)
	t.Log("Empty string handled correctly")
	
	// Test nil byte slice
	var nilBytes []byte
	SecureClearByteSlice(context.Background(), nilBytes) // Should not panic
	t.Log("Nil byte slice handled correctly")
	
	// Test nil map
	var nilMap map[string]interface{}
	SecureClearMap(context.Background(), nilMap, []string{"key"}) // Should not panic
	t.Log("Nil map handled correctly")
}

func BenchmarkSecureClearString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testStr := "BenchmarkPassword123456789"
		SecureClearString(context.Background(), &testStr)
	}
}

func BenchmarkSecureClearByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBytes := []byte("BenchmarkSecretKey123456789")
		SecureClearByteSlice(context.Background(), testBytes)
	}
}

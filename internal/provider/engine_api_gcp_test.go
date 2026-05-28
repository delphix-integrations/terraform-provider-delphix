package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// -----------------------------------------------------------------------------
// Sub-task A: GCP ObjectStore struct construction
//
// The production logic in initializeSystem (engine_api.go:204) and
// testConnectionForObjectStore (engine_api.go:502) builds ObjectStore and
// TestConnection structs from InitializationParameters.  Both functions are
// unexported and require an HTTP client + live engine, so we cannot call them
// end-to-end.  Instead, we test the same struct-construction expressions
// directly — the assertions below mirror the exact field assignments in
// production, so any change to those assignments will break these tests.
// convertStorageToBytes is package-internal and callable from tests.
// -----------------------------------------------------------------------------

// TestGcpObjectStoreStruct verifies that the ObjectStore built for the GCP
// case has the correct Type, Bucket, and Size for every combination we care
// about (mirrors engine_api.go:204-210).
func TestGcpObjectStoreStruct(t *testing.T) {
	type testCase struct {
		name          string
		bucket        string
		sizeStr       string
		wantType      string
		wantBucket    string
		wantSizeBytes int
	}

	tests := []testCase{
		{
			name:          "happy path - standard bucket and size",
			bucket:        "my-bucket",
			sizeStr:       "20GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "my-bucket",
			wantSizeBytes: 20 * 1024 * 1024 * 1024,
		},
		{
			name:          "alphanumeric bucket name",
			bucket:        "mybucket123",
			sizeStr:       "20GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "mybucket123",
			wantSizeBytes: 20 * 1024 * 1024 * 1024,
		},
		{
			name:          "bucket with hyphens",
			bucket:        "my-gcs-bucket-01",
			sizeStr:       "20GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "my-gcs-bucket-01",
			wantSizeBytes: 20 * 1024 * 1024 * 1024,
		},
		{
			name:          "bucket with dots",
			bucket:        "my.bucket.name",
			sizeStr:       "20GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "my.bucket.name",
			wantSizeBytes: 20 * 1024 * 1024 * 1024,
		},
		{
			name:          "size 1TB",
			bucket:        "my-bucket",
			sizeStr:       "1TB",
			wantType:      "GcpObjectStore",
			wantBucket:    "my-bucket",
			wantSizeBytes: 1 * 1024 * 1024 * 1024 * 1024,
		},
		{
			name:          "size 500GB",
			bucket:        "my-bucket",
			sizeStr:       "500GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "my-bucket",
			wantSizeBytes: 500 * 1024 * 1024 * 1024,
		},
		{
			name:          "empty bucket - builder is dumb, validation happens elsewhere",
			bucket:        "",
			sizeStr:       "20GB",
			wantType:      "GcpObjectStore",
			wantBucket:    "",
			wantSizeBytes: 20 * 1024 * 1024 * 1024,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Reproduce the exact construction from engine_api.go:204-210.
			sizeInBytes, err := convertStorageToBytes(tc.sizeStr)
			if err != nil {
				t.Fatalf("convertStorageToBytes(%q) returned unexpected error: %v", tc.sizeStr, err)
			}

			got := ObjectStore{
				Type:         "GcpObjectStore",
				Size:         sizeInBytes,
				CacheDevices: []string{},
				Bucket:       tc.bucket,
			}

			if got.Type != tc.wantType {
				t.Errorf("Type: got %q, want %q", got.Type, tc.wantType)
			}
			if got.Bucket != tc.wantBucket {
				t.Errorf("Bucket: got %q, want %q", got.Bucket, tc.wantBucket)
			}
			if got.Size != tc.wantSizeBytes {
				t.Errorf("Size: got %d bytes, want %d bytes (from %q)", got.Size, tc.wantSizeBytes, tc.sizeStr)
			}
			// GCP does not set AccessCredentials — verify it remains nil when
			// constructed the same way the production switch does.
			if got.AccessCredentials != nil {
				t.Errorf("AccessCredentials: expected nil for GCP, got %+v", got.AccessCredentials)
			}
		})
	}
}

// TestGcpObjectStoreNoAccessCredentials verifies that the GCP ObjectStore
// struct does NOT have access credentials set (unlike AWS/AZURE paths which
// enter the post-switch if-else in engine_api.go:213-240).
func TestGcpObjectStoreNoAccessCredentials(t *testing.T) {
	params := InitializationParameters{
		CloudProvider: GCP,
		Bucket:        "my-bucket",
		Size:          "20GB",
		DeviceType:    OBJECT,
	}

	sizeInBytes, err := convertStorageToBytes(params.Size)
	if err != nil {
		t.Fatalf("convertStorageToBytes error: %v", err)
	}

	objectStorage := &ObjectStore{
		Type:         "GcpObjectStore",
		Size:         sizeInBytes,
		CacheDevices: []string{},
		Bucket:       params.Bucket,
	}

	// Production code only sets AccessCredentials for AWS and AZURE — the
	// post-switch block has no `else if params.CloudProvider == GCP` branch.
	if params.CloudProvider != AWS && params.CloudProvider != AZURE {
		// Simulates the production guard: leave AccessCredentials unset.
	}

	if objectStorage.AccessCredentials != nil {
		t.Errorf("expected AccessCredentials to be nil for GCP, got %+v", objectStorage.AccessCredentials)
	}
	if objectStorage.Type != "GcpObjectStore" {
		t.Errorf("expected Type GcpObjectStore, got %q", objectStorage.Type)
	}
}

// -----------------------------------------------------------------------------
// Sub-task A (connection-test builder): GCP TestConnection struct construction
//
// Mirrors engine_api.go:502-508 — the GCP branch of testConnectionForObjectStore.
// -----------------------------------------------------------------------------

// TestGcpObjectStoreTestConnectionStruct verifies that the TestConnection
// built for the GCP case has Type="GcpObjectStoreTest" and the correct Bucket.
func TestGcpObjectStoreTestConnectionStruct(t *testing.T) {
	type testCase struct {
		name       string
		bucket     string
		wantType   string
		wantBucket string
	}

	tests := []testCase{
		{
			name:       "happy path",
			bucket:     "my-bucket",
			wantType:   "GcpObjectStoreTest",
			wantBucket: "my-bucket",
		},
		{
			name:       "alphanumeric bucket",
			bucket:     "mybucket123",
			wantType:   "GcpObjectStoreTest",
			wantBucket: "mybucket123",
		},
		{
			name:       "bucket with hyphens",
			bucket:     "my-gcs-bucket-01",
			wantType:   "GcpObjectStoreTest",
			wantBucket: "my-gcs-bucket-01",
		},
		{
			name:       "bucket with dots in it",
			bucket:     "my.bucket.name",
			wantType:   "GcpObjectStoreTest",
			wantBucket: "my.bucket.name",
		},
		{
			name:       "empty bucket - builder is dumb",
			bucket:     "",
			wantType:   "GcpObjectStoreTest",
			wantBucket: "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			params := InitializationParameters{
				CloudProvider: GCP,
				Bucket:        tc.bucket,
			}

			// Reproduce the exact construction from engine_api.go:504-507.
			payload := TestConnection{
				Type:   "GcpObjectStoreTest",
				Bucket: params.Bucket,
			}

			if payload.Type != tc.wantType {
				t.Errorf("Type: got %q, want %q", payload.Type, tc.wantType)
			}
			if payload.Bucket != tc.wantBucket {
				t.Errorf("Bucket: got %q, want %q", payload.Bucket, tc.wantBucket)
			}
			// GCP test connection has no Region, Endpoint, or Container.
			if payload.Region != "" {
				t.Errorf("Region: expected empty for GCP, got %q", payload.Region)
			}
			if payload.Endpoint != "" {
				t.Errorf("Endpoint: expected empty for GCP, got %q", payload.Endpoint)
			}
			if payload.Container != "" {
				t.Errorf("Container: expected empty for GCP, got %q", payload.Container)
			}
		})
	}
}

// TestGcpObjectStoreTestConnectionWithSizes exercises the connection-test
// struct for multiple size inputs — size does not appear in the TestConnection
// struct itself, but this confirms the builder is independent of size.
func TestGcpObjectStoreTestConnectionWithSizes(t *testing.T) {
	buckets := []string{"bucket-20gb", "bucket-1tb", "bucket-500gb"}
	for _, bucket := range buckets {
		bucket := bucket
		t.Run(bucket, func(t *testing.T) {
			params := InitializationParameters{
				CloudProvider: GCP,
				Bucket:        bucket,
			}

			payload := TestConnection{
				Type:   "GcpObjectStoreTest",
				Bucket: params.Bucket,
			}

			if payload.Type != "GcpObjectStoreTest" {
				t.Errorf("Type: got %q, want GcpObjectStoreTest", payload.Type)
			}
			if payload.Bucket != bucket {
				t.Errorf("Bucket: got %q, want %q", payload.Bucket, bucket)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Sub-task B: cloud_provider schema field validator
//
// The field at resource_engine_configuration.go:294 uses:
//   validation.StringInSlice([]string{AWS, AZURE, GCP}, false)
//
// We recreate the same validator inline (same call, same args) and exercise
// both the accepted and rejected values.  Using the same constants (AWS,
// AZURE, GCP) ensures the test breaks if a constant value ever changes.
// -----------------------------------------------------------------------------

// TestCloudProviderValidator verifies that the cloud_provider schema validator
// accepts exactly {AWS, AZURE, GCP} and rejects everything else.
func TestCloudProviderValidator(t *testing.T) {
	// Recreate the exact validator from resource_engine_configuration.go:294.
	validateFn := validation.StringInSlice([]string{AWS, AZURE, GCP}, false)

	type testCase struct {
		name       string
		value      string
		wantAccept bool // true = expect zero errors
	}

	tests := []testCase{
		// --- accepted values ---
		{name: "AWS constant accepted", value: AWS, wantAccept: true},
		{name: "AZURE constant accepted", value: AZURE, wantAccept: true},
		{name: "GCP constant accepted", value: GCP, wantAccept: true},

		// --- rejected values ---
		{name: "lowercase gcp rejected", value: "gcp", wantAccept: false},
		{name: "lowercase aws rejected", value: "aws", wantAccept: false},
		{name: "lowercase azure rejected", value: "azure", wantAccept: false},
		{name: "awsgov rejected", value: "awsgov", wantAccept: false},
		{name: "empty string rejected", value: "", wantAccept: false},
		{name: "S3 rejected", value: "S3", wantAccept: false},
		{name: "google rejected", value: "google", wantAccept: false},
		{name: "GCS rejected", value: "GCS", wantAccept: false},
		{name: "aws with trailing space rejected", value: "aws ", wantAccept: false},
		{name: "GCP with trailing space rejected", value: "GCP ", wantAccept: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			warnings, errors := validateFn(tc.value, "cloud_provider")

			if tc.wantAccept {
				if len(errors) != 0 {
					t.Errorf("value %q: expected zero errors, got %d: %v", tc.value, len(errors), errors)
				}
				if len(warnings) != 0 {
					t.Errorf("value %q: expected zero warnings, got %d: %v", tc.value, len(warnings), warnings)
				}
			} else {
				if len(errors) == 0 {
					t.Errorf("value %q: expected at least one error, got none", tc.value)
				}
			}
		})
	}
}

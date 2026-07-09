package provider

import (
	"testing"
)

// -----------------------------------------------------------------------------
// Factory
// -----------------------------------------------------------------------------

func TestCloudProviderForKnown(t *testing.T) {
	for _, name := range []string{AWS, AZURE, GCP} {
		p, err := cloudProviderFor(name)
		if err != nil {
			t.Errorf("cloudProviderFor(%q) returned error: %v", name, err)
		}
		if p == nil {
			t.Errorf("cloudProviderFor(%q) returned nil provider", name)
		}
	}
}

func TestCloudProviderForUnknown(t *testing.T) {
	p, err := cloudProviderFor("ORACLE")
	if err == nil {
		t.Error("cloudProviderFor(\"ORACLE\") expected error, got nil")
	}
	if p != nil {
		t.Errorf("cloudProviderFor(\"ORACLE\") expected nil provider, got %+v", p)
	}
}

// -----------------------------------------------------------------------------
// Validate
// -----------------------------------------------------------------------------

func TestAWSValidate(t *testing.T) {
	tests := []struct {
		name    string
		block   map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid with ROLE",
			block:   map[string]interface{}{"endpoint": "e", "region": "r", "bucket": "b", "auth_type": ROLE},
			wantErr: false,
		},
		{
			name:    "valid with ACCESS_KEY",
			block:   map[string]interface{}{"endpoint": "e", "region": "r", "bucket": "b", "auth_type": ACCESS_KEY, "access_id": "id", "access_key": "key"},
			wantErr: false,
		},
		{
			name:    "missing endpoint",
			block:   map[string]interface{}{"region": "r", "bucket": "b"},
			wantErr: true,
		},
		{
			name:    "missing region",
			block:   map[string]interface{}{"endpoint": "e", "bucket": "b"},
			wantErr: true,
		},
		{
			name:    "missing bucket",
			block:   map[string]interface{}{"endpoint": "e", "region": "r"},
			wantErr: true,
		},
		{
			name:    "invalid auth_type",
			block:   map[string]interface{}{"endpoint": "e", "region": "r", "bucket": "b", "auth_type": "BOGUS"},
			wantErr: true,
		},
		{
			name:    "ACCESS_KEY without credentials",
			block:   map[string]interface{}{"endpoint": "e", "region": "r", "bucket": "b", "auth_type": ACCESS_KEY, "access_id": "", "access_key": ""},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := awsProvider{}.Validate(tc.block)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestAzureValidate(t *testing.T) {
	tests := []struct {
		name    string
		block   map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid with MANAGED_IDENTITIES",
			block:   map[string]interface{}{"azure_container": "c", "azure_account": "a", "auth_type": MANAGED_IDENTITIES},
			wantErr: false,
		},
		{
			name:    "valid with ACCESS_KEY",
			block:   map[string]interface{}{"azure_container": "c", "azure_account": "a", "auth_type": ACCESS_KEY, "azure_key": "k"},
			wantErr: false,
		},
		{
			name:    "missing container",
			block:   map[string]interface{}{"azure_container": "", "azure_account": "a"},
			wantErr: true,
		},
		{
			name:    "missing account",
			block:   map[string]interface{}{"azure_container": "c", "azure_account": ""},
			wantErr: true,
		},
		{
			name:    "invalid auth_type",
			block:   map[string]interface{}{"azure_container": "c", "azure_account": "a", "auth_type": "BOGUS"},
			wantErr: true,
		},
		{
			name:    "ACCESS_KEY without azure_key",
			block:   map[string]interface{}{"azure_container": "c", "azure_account": "a", "auth_type": ACCESS_KEY, "azure_key": ""},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := azureProvider{}.Validate(tc.block)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestGCPValidate(t *testing.T) {
	tests := []struct {
		name    string
		block   map[string]interface{}
		wantErr bool
	}{
		{name: "valid bucket", block: map[string]interface{}{"bucket": "b"}, wantErr: false},
		{name: "empty bucket", block: map[string]interface{}{"bucket": ""}, wantErr: true},
		{name: "missing bucket", block: map[string]interface{}{}, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := gcpProvider{}.Validate(tc.block)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// ExtractParams
// -----------------------------------------------------------------------------

func TestAWSExtractParams(t *testing.T) {
	t.Run("ACCESS_KEY", func(t *testing.T) {
		block := map[string]interface{}{
			"endpoint": "e", "region": "r", "bucket": "b",
			"access_id": "id", "access_key": "key", "s3_instance_profile": "prof",
		}
		p := InitializationParameters{AuthType: ACCESS_KEY}
		awsProvider{}.ExtractParams(block, &p)
		if p.Endpoint != "e" || p.Region != "r" || p.Bucket != "b" {
			t.Errorf("base fields not extracted: %+v", p)
		}
		if p.ACCESS_ID != "id" || p.ACCESS_KEY != "key" {
			t.Errorf("access key fields not extracted: %+v", p)
		}
		if p.S3_INSTANCE_PROFILE != "" {
			t.Errorf("S3_INSTANCE_PROFILE should be empty for ACCESS_KEY, got %q", p.S3_INSTANCE_PROFILE)
		}
	})
	t.Run("ROLE", func(t *testing.T) {
		block := map[string]interface{}{
			"endpoint": "e", "region": "r", "bucket": "b",
			"access_id": "id", "access_key": "key", "s3_instance_profile": "prof",
		}
		p := InitializationParameters{AuthType: ROLE}
		awsProvider{}.ExtractParams(block, &p)
		if p.S3_INSTANCE_PROFILE != "prof" {
			t.Errorf("S3_INSTANCE_PROFILE not extracted: %+v", p)
		}
		if p.ACCESS_ID != "" || p.ACCESS_KEY != "" {
			t.Errorf("access key fields should be empty for ROLE: %+v", p)
		}
	})
}

func TestAzureExtractParams(t *testing.T) {
	t.Run("ACCESS_KEY", func(t *testing.T) {
		block := map[string]interface{}{
			"azure_container": "c", "azure_account": "a",
			"azure_key": "k", "azure_managed_identities": "mi",
		}
		p := InitializationParameters{AuthType: ACCESS_KEY}
		azureProvider{}.ExtractParams(block, &p)
		if p.Container != "c" || p.AZURE_ACCOUNT != "a" {
			t.Errorf("base fields not extracted: %+v", p)
		}
		if p.AZURE_KEY != "k" {
			t.Errorf("AZURE_KEY not extracted: %+v", p)
		}
		if p.AzureManagedIdentities != "" {
			t.Errorf("AzureManagedIdentities should be empty for ACCESS_KEY: %+v", p)
		}
	})
	t.Run("MANAGED_IDENTITIES", func(t *testing.T) {
		block := map[string]interface{}{
			"azure_container": "c", "azure_account": "a",
			"azure_key": "k", "azure_managed_identities": "mi",
		}
		p := InitializationParameters{AuthType: MANAGED_IDENTITIES}
		azureProvider{}.ExtractParams(block, &p)
		if p.AzureManagedIdentities != "mi" {
			t.Errorf("AzureManagedIdentities not extracted: %+v", p)
		}
		if p.AZURE_KEY != "" {
			t.Errorf("AZURE_KEY should be empty for MANAGED_IDENTITIES: %+v", p)
		}
	})
}

func TestGCPExtractParams(t *testing.T) {
	block := map[string]interface{}{"bucket": "b"}
	p := InitializationParameters{}
	gcpProvider{}.ExtractParams(block, &p)
	if p.Bucket != "b" {
		t.Errorf("Bucket not extracted: %+v", p)
	}
}

// -----------------------------------------------------------------------------
// BuildObjectStore
// -----------------------------------------------------------------------------

func TestAWSBuildObjectStore(t *testing.T) {
	devs := []string{"d1"}
	t.Run("ACCESS_KEY", func(t *testing.T) {
		p := InitializationParameters{CloudProvider: AWS, AuthType: ACCESS_KEY, Endpoint: "e", Region: "r", Bucket: "b", ACCESS_ID: "id", ACCESS_KEY: "key"}
		os := awsProvider{}.BuildObjectStore(p, 100, devs)
		if os.Type != "S3ObjectStore" || os.Endpoint != "e" || os.Region != "r" || os.Bucket != "b" || os.Size != 100 {
			t.Errorf("unexpected object store: %+v", os)
		}
		if os.AccessCredentials == nil || os.AccessCredentials.Type != "S3ObjectStoreAccessKey" || os.AccessCredentials.ACCESS_ID != "id" || os.AccessCredentials.ACCESS_KEY != "key" {
			t.Errorf("unexpected credentials: %+v", os.AccessCredentials)
		}
	})
	t.Run("ROLE", func(t *testing.T) {
		p := InitializationParameters{CloudProvider: AWS, AuthType: ROLE, Endpoint: "e", Region: "r", Bucket: "b", S3_INSTANCE_PROFILE: "prof"}
		os := awsProvider{}.BuildObjectStore(p, 100, devs)
		if os.AccessCredentials == nil || os.AccessCredentials.Type != "prof" {
			t.Errorf("unexpected credentials: %+v", os.AccessCredentials)
		}
	})
}

func TestAzureBuildObjectStore(t *testing.T) {
	devs := []string{"d1"}
	t.Run("ACCESS_KEY", func(t *testing.T) {
		p := InitializationParameters{CloudProvider: AZURE, AuthType: ACCESS_KEY, Container: "c", AZURE_ACCOUNT: "a", AZURE_KEY: "k"}
		os := azureProvider{}.BuildObjectStore(p, 100, devs)
		if os.Type != "BlobObjectStore" || os.Container != "c" || os.Size != 100 {
			t.Errorf("unexpected object store: %+v", os)
		}
		if os.AccessCredentials == nil || os.AccessCredentials.Type != "BlobObjectStoreAccessKey" || os.AccessCredentials.AZURE_ACCOUNT != "a" || os.AccessCredentials.AZURE_KEY != "k" {
			t.Errorf("unexpected credentials: %+v", os.AccessCredentials)
		}
	})
	t.Run("MANAGED_IDENTITIES", func(t *testing.T) {
		p := InitializationParameters{CloudProvider: AZURE, AuthType: MANAGED_IDENTITIES, Container: "c", AZURE_ACCOUNT: "a", AzureManagedIdentities: "mi"}
		os := azureProvider{}.BuildObjectStore(p, 100, devs)
		if os.AccessCredentials == nil || os.AccessCredentials.Type != "mi" || os.AccessCredentials.AZURE_ACCOUNT != "a" {
			t.Errorf("unexpected credentials: %+v", os.AccessCredentials)
		}
	})
}

func TestGCPBuildObjectStore(t *testing.T) {
	p := InitializationParameters{CloudProvider: GCP, Bucket: "b"}
	os := gcpProvider{}.BuildObjectStore(p, 100, []string{"d1"})
	if os.Type != "GcpObjectStore" || os.Bucket != "b" || os.Size != 100 {
		t.Errorf("unexpected object store: %+v", os)
	}
	if os.AccessCredentials != nil {
		t.Errorf("GCP should not set AccessCredentials, got %+v", os.AccessCredentials)
	}
}

// -----------------------------------------------------------------------------
// BuildTestConnection
// -----------------------------------------------------------------------------

func TestAWSBuildTestConnection(t *testing.T) {
	t.Run("ACCESS_KEY", func(t *testing.T) {
		p := InitializationParameters{AuthType: ACCESS_KEY, Endpoint: "e", Region: "r", Bucket: "b", ACCESS_ID: "id", ACCESS_KEY: "key"}
		tc := awsProvider{}.BuildTestConnection(p)
		if tc.Type != "S3ObjectStoreTest" || tc.Endpoint != "e" || tc.Region != "r" || tc.Bucket != "b" {
			t.Errorf("unexpected test connection: %+v", tc)
		}
		if tc.AccessCredentials.Type != "S3ObjectStoreAccessKey" || tc.AccessCredentials.ACCESS_ID != "id" {
			t.Errorf("unexpected credentials: %+v", tc.AccessCredentials)
		}
	})
	t.Run("ROLE", func(t *testing.T) {
		p := InitializationParameters{AuthType: ROLE, Endpoint: "e", Region: "r", Bucket: "b", S3_INSTANCE_PROFILE: "prof"}
		tc := awsProvider{}.BuildTestConnection(p)
		if tc.AccessCredentials.Type != "prof" {
			t.Errorf("unexpected credentials: %+v", tc.AccessCredentials)
		}
	})
}

func TestAzureBuildTestConnection(t *testing.T) {
	t.Run("MANAGED_IDENTITIES", func(t *testing.T) {
		p := InitializationParameters{AuthType: MANAGED_IDENTITIES, Container: "c", AZURE_ACCOUNT: "a", AzureManagedIdentities: "mi"}
		tc := azureProvider{}.BuildTestConnection(p)
		if tc.Type != "BlobObjectStoreTest" || tc.Container != "c" {
			t.Errorf("unexpected test connection: %+v", tc)
		}
		if tc.AccessCredentials.Type != "mi" || tc.AccessCredentials.AZURE_ACCOUNT != "a" {
			t.Errorf("unexpected credentials: %+v", tc.AccessCredentials)
		}
	})
	t.Run("ACCESS_KEY", func(t *testing.T) {
		p := InitializationParameters{AuthType: ACCESS_KEY, Container: "c", AZURE_ACCOUNT: "a", AZURE_KEY: "k"}
		tc := azureProvider{}.BuildTestConnection(p)
		if tc.AccessCredentials.Type != "BlobObjectStoreAccessKey" || tc.AccessCredentials.AZURE_KEY != "k" {
			t.Errorf("unexpected credentials: %+v", tc.AccessCredentials)
		}
	})
}

func TestGCPBuildTestConnection(t *testing.T) {
	p := InitializationParameters{Bucket: "b"}
	tc := gcpProvider{}.BuildTestConnection(p)
	if tc.Type != "GcpObjectStoreTest" || tc.Bucket != "b" {
		t.Errorf("unexpected test connection: %+v", tc)
	}
}

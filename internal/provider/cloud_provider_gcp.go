package provider

import "errors"

// gcpProvider implements CloudProvider for GCP object storage.
type gcpProvider struct{}

func (gcpProvider) Validate(block map[string]interface{}) error {
	if bucket, ok := block["bucket"].(string); !ok || bucket == "" {
		return errors.New("bucket must be a non-empty string in object_storage_params for GCP cloud_provider")
	}
	return nil
}

func (gcpProvider) ExtractParams(block map[string]interface{}, p *InitializationParameters) {
	p.Bucket = block["bucket"].(string)
}

func (gcpProvider) BuildObjectStore(p InitializationParameters, sizeBytes int, deviceRefs []string) *ObjectStore {
	return &ObjectStore{
		Type:         "GcpObjectStore",
		Size:         sizeBytes,
		CacheDevices: deviceRefs,
		Bucket:       p.Bucket,
	}
}

func (gcpProvider) BuildTestConnection(p InitializationParameters) TestConnection {
	return TestConnection{
		Type:   "GcpObjectStoreTest",
		Bucket: p.Bucket,
	}
}

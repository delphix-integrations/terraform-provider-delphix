package provider

import "errors"

// awsProvider implements CloudProvider for AWS S3 object storage.
type awsProvider struct{}

func (awsProvider) Validate(block map[string]interface{}) error {
	if _, ok := block["endpoint"]; !ok {
		return errors.New("endpoint must be provided in object_storage_params for AWS cloud_provider")
	}
	if _, ok := block["region"]; !ok {
		return errors.New("region must be provided in object_storage_params for AWS cloud_provider")
	}
	if _, ok := block["bucket"]; !ok {
		return errors.New("bucket must be provided in object_storage_params for AWS cloud_provider")
	}
	if authType, ok := block["auth_type"]; ok {
		authTypeStr := authType.(string)
		if authTypeStr != ROLE && authTypeStr != ACCESS_KEY {
			return errors.New("auth_type for AWS cloud_provider must be either ROLE or ACCESS_KEY")
		}
		if authType == ACCESS_KEY && (block["access_id"] == "" || block["access_key"] == "") {
			return errors.New("access_id and access_key must be provided when auth_type is ACCESS_KEY")
		}
	}
	return nil
}

func (awsProvider) ExtractParams(block map[string]interface{}, p *InitializationParameters) {
	p.Endpoint = block["endpoint"].(string)
	p.Region = block["region"].(string)
	p.Bucket = block["bucket"].(string)
	if p.AuthType == ACCESS_KEY {
		p.ACCESS_ID = block["access_id"].(string)
		p.ACCESS_KEY = block["access_key"].(string)
	} else {
		p.S3_INSTANCE_PROFILE = block["s3_instance_profile"].(string)
	}
}

func (awsProvider) BuildObjectStore(p InitializationParameters, sizeBytes int, deviceRefs []string) *ObjectStore {
	objectStorage := &ObjectStore{
		Type:         "S3ObjectStore",
		Size:         sizeBytes,
		CacheDevices: deviceRefs,
		Endpoint:     p.Endpoint,
		Region:       p.Region,
		Bucket:       p.Bucket,
	}
	switch p.AuthType {
	case ROLE:
		objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
			Type: p.S3_INSTANCE_PROFILE,
		}
	case ACCESS_KEY:
		objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
			Type:       "S3ObjectStoreAccessKey",
			ACCESS_ID:  p.ACCESS_ID,
			ACCESS_KEY: p.ACCESS_KEY,
		}
	}
	return objectStorage
}

func (awsProvider) BuildTestConnection(p InitializationParameters) TestConnection {
	if p.AuthType == ACCESS_KEY {
		return TestConnection{
			Type:     "S3ObjectStoreTest",
			Endpoint: p.Endpoint,
			Region:   p.Region,
			Bucket:   p.Bucket,
			AccessCredentials: ObjectStoreAccessCredentials{
				Type:       "S3ObjectStoreAccessKey",
				ACCESS_ID:  p.ACCESS_ID,
				ACCESS_KEY: p.ACCESS_KEY,
			},
		}
	}
	return TestConnection{
		Type:     "S3ObjectStoreTest",
		Endpoint: p.Endpoint,
		Region:   p.Region,
		Bucket:   p.Bucket,
		AccessCredentials: ObjectStoreAccessCredentials{
			Type: p.S3_INSTANCE_PROFILE,
		},
	}
}

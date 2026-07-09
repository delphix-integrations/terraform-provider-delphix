package provider

import "errors"

// azureProvider implements CloudProvider for Azure Blob object storage.
type azureProvider struct{}

func (azureProvider) Validate(block map[string]interface{}) error {
	if container, _ := block["azure_container"].(string); container == "" {
		return errors.New("azure_container must be provided in object_storage_params for AZURE cloud_provider")
	}
	if account, _ := block["azure_account"].(string); account == "" {
		return errors.New("azure_account must be provided in object_storage_params for AZURE cloud_provider")
	}
	if authType, ok := block["auth_type"]; ok {
		authTypeStr := authType.(string)
		if authTypeStr != MANAGED_IDENTITIES && authTypeStr != ACCESS_KEY {
			return errors.New("auth_type for AZURE cloud_provider must be either MANAGED_IDENTITIES or ACCESS_KEY")
		}
		if authTypeStr == ACCESS_KEY && block["azure_key"] == "" {
			return errors.New("azure_key must be provided when auth_type is ACCESS_KEY for AZURE cloud_provider")
		}
	}
	return nil
}

func (azureProvider) ExtractParams(block map[string]interface{}, p *InitializationParameters) {
	p.Container = block["azure_container"].(string)
	p.AZURE_ACCOUNT = block["azure_account"].(string)
	if p.AuthType == ACCESS_KEY {
		p.AZURE_KEY = block["azure_key"].(string)
	} else {
		p.AzureManagedIdentities = block["azure_managed_identities"].(string)
	}
}

func (azureProvider) BuildObjectStore(p InitializationParameters, sizeBytes int, deviceRefs []string) *ObjectStore {
	objectStorage := &ObjectStore{
		Type:         "BlobObjectStore",
		Size:         sizeBytes,
		CacheDevices: deviceRefs,
		Container:    p.Container,
	}
	switch p.AuthType {
	case ACCESS_KEY:
		objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
			Type:          "BlobObjectStoreAccessKey",
			AZURE_ACCOUNT: p.AZURE_ACCOUNT,
			AZURE_KEY:     p.AZURE_KEY,
		}
	case MANAGED_IDENTITIES:
		objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
			Type:          p.AzureManagedIdentities,
			AZURE_ACCOUNT: p.AZURE_ACCOUNT,
		}
	}
	return objectStorage
}

func (azureProvider) BuildTestConnection(p InitializationParameters) TestConnection {
	if p.AuthType == MANAGED_IDENTITIES {
		return TestConnection{
			Type:      "BlobObjectStoreTest",
			Container: p.Container,
			AccessCredentials: ObjectStoreAccessCredentials{
				Type:          p.AzureManagedIdentities,
				AZURE_ACCOUNT: p.AZURE_ACCOUNT,
			},
		}
	}
	return TestConnection{
		Type:      "BlobObjectStoreTest",
		Container: p.Container,
		AccessCredentials: ObjectStoreAccessCredentials{
			Type:          "BlobObjectStoreAccessKey",
			AZURE_ACCOUNT: p.AZURE_ACCOUNT,
			AZURE_KEY:     p.AZURE_KEY,
		},
	}
}

package provider

import "fmt"

// CloudProvider abstracts the per-cloud behavior for object-storage engine
// configuration. Each supported cloud (AWS, Azure, GCP) provides one
// implementation, so adding a new cloud means adding one file and one registry
// entry rather than editing every cloud_provider switch.
type CloudProvider interface {
	// Validate checks the object_storage_params block for this cloud during
	// CustomizeDiff. It returns an error describing the first problem found.
	Validate(block map[string]interface{}) error

	// ExtractParams copies the cloud-specific fields from the schema block into
	// p. Common fields (CloudProvider, Size, AuthType) are populated by the
	// caller before dispatch. ExtractParams runs after Validate, so it does not
	// return an error.
	ExtractParams(block map[string]interface{}, p *InitializationParameters)

	// BuildObjectStore constructs the ObjectStore payload (including access
	// credentials) used to initialize the engine.
	BuildObjectStore(p InitializationParameters, sizeBytes int, deviceRefs []string) *ObjectStore

	// BuildTestConnection constructs the TestConnection payload used to verify
	// object-store reachability before initialization.
	BuildTestConnection(p InitializationParameters) TestConnection
}

// cloudProviders is the registry of supported cloud providers keyed by the
// cloud_provider schema value.
var cloudProviders = map[string]CloudProvider{
	AWS:   awsProvider{},
	AZURE: azureProvider{},
	GCP:   gcpProvider{},
}

// cloudProviderFor returns the CloudProvider strategy for the given name, or an
// error if the name is unsupported. The schema's StringInSlice already gates
// input, so this error is defensive.
func cloudProviderFor(name string) (CloudProvider, error) {
	p, ok := cloudProviders[name]
	if !ok {
		return nil, fmt.Errorf("unsupported cloud_provider %q", name)
	}
	return p, nil
}

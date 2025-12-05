package provider

type APIVersion struct {
	Minor int    `json:"minor"`
	Major int    `json:"major"`
	Micro int    `json:"micro"`
	Type  string `json:"type"`
}

type APISession struct {
	Version APIVersion `json:"version"`
	Type    string     `json:"type"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Target   string `json:"target"`
	Type     string `json:"type"`
}

type StorageDevice struct {
	Type       string  `json:"type"`
	Reference  string  `json:"reference"`
	Configured bool    `json:"configured"`
	Size       float64 `json:"size"`
	// Include other relevant device properties here
}

type ListResult struct {
	Type     string          `json:"type"`
	Status   string          `json:"status"`
	Result   []StorageDevice `json:"result"`
	Job      interface{}     `json:"job"`
	Action   interface{}     `json:"action"`
	Total    int             `json:"total"`
	Overflow bool            `json:"overflow"`
}

type SystemInitializationParameters struct {
	Type            string   `json:"type"`
	DefaultUser     string   `json:"defaultUser"`
	DefaultPassword string   `json:"defaultPassword"`
	Devices         []string `json:"devices"`
	DefaultEmail    string   `json:"defaultEmail"`
}

// type ActionResult struct {
// 	Type      string `json:"type"`
// 	Reference string `json:"reference"`
// 	State     string `json:"state"`
// 	Action    string `json:"action"`
// }

type User struct {
	Name                        string     `json:"name"`
	UserType                    string     `json:"userType"`
	AuthenticationType          string     `json:"authenticationType"`
	Credential                  Credential `json:"credential"`
	AllowPasswordAuthentication bool       `json:"allowPasswordAuthentication"`
	Type                        string     `json:"type"`
	FirstName                   string     `json:"firstName"`
	LastName                    string     `json:"lastName"`
	Email                       string     `json:"email"`
}

type Credential struct {
	Password string `json:"password"`
	Type     string `json:"type"`
}

type ActionResult struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Action string `json:"action"`
	Result struct {
		Type      string `json:"type"`
		State     string `json:"state"`
		Reference string `json:"reference"`
	} `json:"result"`
}

type CredentialUpdateParameters struct {
	Type          string `json:"type"`
	NewCredential struct {
		Type     string `json:"type"`
		Password string `json:"password"`
	} `json:"newCredential"`
}

type SystemInfoResponse struct {
	Type   string                 `json:"type"`
	Status string                 `json:"status"`
	Result map[string]interface{} `json:"result"`
	Job    interface{}            `json:"job,omitempty"`
	Action interface{}            `json:"action,omitempty"`
}

type SystemInfo struct {
	Type       string `json:"type"`
	EngineType string `json:"engineType"`
}

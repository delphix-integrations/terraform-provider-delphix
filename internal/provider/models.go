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

type SystemInitializationBlockStorage struct {
	Type            string   `json:"type"`
	DefaultUser     string   `json:"defaultUser"`
	DefaultPassword string   `json:"defaultPassword"`
	DefaultEmail    string   `json:"defaultEmail,omitempty"`
	Devices         []string `json:"devices,omitempty"`
}
type SystemInitializationObjectStore struct {
	Type            string       `json:"type"`
	DefaultUser     string       `json:"defaultUser"`
	DefaultPassword string       `json:"defaultPassword"`
	DefaultEmail    string       `json:"defaultEmail,omitempty"`
	ObjectStore     *ObjectStore `json:"objectStore,omitempty"`
}

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

type InitializationParameters struct {
	User                   string
	Email                  string
	Password               string
	DeviceType             string
	Endpoint               string `json:"endpoint,omitempty"`
	Region                 string `json:"region,omitempty"`
	Bucket                 string `json:"bucket,omitempty"`
	Size                   string `json:"size,omitempty"`
	AuthType               string `json:"auth_type,omitempty"`
	ACCESS_ID              string `json:"access_id,omitempty"`
	ACCESS_KEY             string `json:"access_key,omitempty"`
	S3_INSTANCE_PROFILE    string `json:"s3_instance_profile,omitempty"`
	AzureManagedIdentities string `json:"azure_managed_identities,omitempty"`
	CloudProvider          string `json:"cloud_provider,omitempty"`
	Container              string `json:"container,omitempty"`
	AzureAccount           string `json:"azureAccount,omitempty"`
}

type TestConnection struct {
	Endpoint          string                       `json:"endpoint,omitempty"`
	Region            string                       `json:"region,omitempty"`
	Bucket            string                       `json:"bucket,omitempty"`
	Type              string                       `json:"type"`
	Container         string                       `json:"container,omitempty"`
	AccessCredentials ObjectStoreAccessCredentials `json:"accessCredentials"`
}

type TestConnectionResult struct {
	Type   string                `json:"type"`
	Status string                `json:"status"`
	Result ObjectStoreTestResult `json:"result"`
	Job    interface{}           `json:"job"`
	Action interface{}           `json:"action"`
}

type ObjectStoreAccessCredentials struct {
	Type         string `json:"type"`
	ACCESS_ID    string `json:"accessId,omitempty"`
	ACCESS_KEY   string `json:"accessKey,omitempty"`
	Azureaccount string `json:"azureAccount,omitempty"`
}

type ObjectStore struct {
	Type              string                        `json:"type"`
	Size              int                           `json:"size"`
	CacheDevices      []string                      `json:"cacheDevices"`
	Endpoint          string                        `json:"endpoint,omitempty"`
	Region            string                        `json:"region,omitempty"`
	Bucket            string                        `json:"bucket,omitempty"`
	Container         string                        `json:"container,omitempty"`
	AccessCredentials *ObjectStoreAccessCredentials `json:"accessCredentials"`
}

type ObjectStoreTestResult struct {
	Type         string `json:"type"`
	Result       bool   `json:"result"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type APIResponse struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Action string `json:"action"`
	Job    string `json:"job"`
	Result string `json:"result"`
}

type SetNTPConfig struct {
	Enabled bool     `json:"enabled"`
	Servers []string `json:"servers"`
	Type    string   `json:"type"`
}

type NTPServerParams struct {
	SystemTimeZone string       `json:"systemTimeZone"`
	NTPConfig      SetNTPConfig `json:"ntpConfig"`
	Type           string       `json:"type"`
}

type GetNTPResponse struct {
	Type   string      `json:"type"`
	Status string      `json:"status"`
	Result TimeConfig  `json:"result"`
	Job    interface{} `json:"job"`
	Action interface{} `json:"action"`
}

type TimeConfig struct {
	Type                       string    `json:"type"`
	CurrentTime                string    `json:"currentTime"`
	SystemTimeZone             string    `json:"systemTimeZone"`
	SystemTimeZoneOffset       int       `json:"systemTimeZoneOffset"`
	SystemTimeZoneOffsetString string    `json:"systemTimeZoneOffsetString"`
	NtpConfig                  NTPConfig `json:"ntpConfig"`
}

type NTPConfig struct {
	Type             string   `json:"type"`
	Enabled          bool     `json:"enabled"`
	Servers          []string `json:"servers"`
	UseMulticast     bool     `json:"useMulticast"`
	MulticastAddress string   `json:"multicastAddress"`
}

type SMTPConfig struct {
	Type                  string `json:"type"`
	Enabled               bool   `json:"enabled"`
	Server                string `json:"server"`
	Port                  int    `json:"port"`
	AuthenticationEnabled bool   `json:"authenticationEnabled"`
	Username              string `json:"username,omitempty"`
	Password              string `json:"password,omitempty"`
	TlsEnabled            bool   `json:"tlsEnabled"`
	FromAddress           string `json:"fromAddress"`
	SendTimeout           int    `json:"sendTimeout"`
}

type DNSConfig struct {
	Type    string   `json:"type"`
	Servers []string `json:"servers"`
	Domains []string `json:"domain,omitempty"`
}

type GetDNSResponse struct {
	Type   string      `json:"type"`
	Status string      `json:"status"`
	Result DNSConfig   `json:"result"`
	Job    interface{} `json:"job"`
	Action interface{} `json:"action"`
}

type PhoneHomeConfig struct {
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type UserAnalyticsConfig struct {
	Type             string `json:"type"`
	AnalyticsEnabled bool   `json:"analyticsEnabled"`
}

type WebProxyConfig struct {
	Type  string              `json:"type"`
	Https *ProxyConfiguration `json:"https,omitempty"`
}

type ProxyConfiguration struct {
	Type     string `json:"type"`
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SSOConfig struct {
	Type                 string `json:"type"`
	Enabled              bool   `json:"enabled"`
	EntityId             string `json:"entityId"`
	SamlMetadata         string `json:"samlMetadata"`
	ResponseSkewTime     int    `json:"responseSkewTime"`
	MaxAuthenticationAge int    `json:"maxAuthenticationAge"`
}

const (
	BLOCK                    = "BLOCK"
	OBJECT                   = "OBJECT"
	ROLE                     = "ROLE"
	ACCESS_KEY               = "ACCESS_KEY"
	MANAGED_IDENTITIES       = "MANAGED_IDENTITIES"
	DEFAULT_SSO_SKEW_TIME    = 120
	DEFAULT_SSO_MAX_AUTH_AGE = 86400
	DEFAULT_SEND_TIMEOUT     = 60
	CONTINUOUS_COMPLIANCE    = "CC"
	CONTINUOUS_DATA          = "CD"
	SYSTEM                   = "SYSTEM"
	AZURE                    = "AZURE"
	AWS                      = "AWS"
	ENGINE_API_VERSION       = "1.11.47"
)

type ConfigTask struct {
	Name      string
	Condition bool
	Task      func() error
}

var ENGINE_APIS = map[string]string{
	"SESSION":                      "/resources/json/delphix/session",
	"LOGIN":                        "/resources/json/delphix/login",
	"COMPLIANCE_LOGIN":             "/masking/api/login",
	"STORAGE_DEVICE":               "/resources/json/delphix/storage/device",
	"USER":                         "/resources/json/delphix/user/",
	"SYSTEM_INITIALIZATION":        "/resources/json/delphix/domain/initializeSystem",
	"NTP_CONFIG":                   "/resources/json/delphix/service/time",
	"SMTP_CONFIG":                  "/resources/json/delphix/service/smtp",
	"DNS_CONFIG":                   "/resources/json/delphix/service/dns",
	"PHONE_HOME_CONFIG":            "/resources/json/delphix/service/phonehome",
	"USER_ANALYTICS_CONFIG":        "/resources/json/delphix/service/userInterface",
	"WEB_PROXY_CONFIG":             "/resources/json/delphix/service/proxy",
	"SSO_CONFIG":                   "/resources/json/delphix/service/sso",
	"OBJECT_STORE_TEST_CONNECTION": "/resources/json/delphix/storage/objectStorage/testConnection",
	"ACTION":                       "/resources/json/delphix/action/",
	"SYSTEM_INFO":                  "/resources/json/delphix/system",
	"COMPLIANCE_USER":              "/masking/api/users",
	"START_MASKING":                "/resources/json/delphix/system/startMasking",
}

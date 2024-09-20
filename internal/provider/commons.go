package provider

const (
	Pending                string = "PENDING"
	Started                string = "STARTED"
	Timedout               string = "TIMEDOUT"
	Failed                 string = "FAILED"
	Completed              string = "COMPLETED"
	Canceled               string = "CANCELED"
	Abandoned              string = "ABANDONED"
	JOB_STATUS_SLEEP_TIME  int    = 5
	STATUS_POLL_SLEEP_TIME int    = 20
	DLPX                   string = "[DELPHIX] "
	INFO                   string = "[INFO] "
	WARN                   string = "[WARN] "
	ERROR                  string = "[ERROR] "
)

var updatableEnvKeys = map[string]bool{
	"name":                           true,
	"cluster_home":                   true,
	"connector_port":                 true,
	"username":                       true,
	"password":                       true,
	"java_home":                      true,
	"description":                    true,
	"hostname":                       true,
	"ssh_port":                       true,
	"toolkit_path":                   true,
	"nfs_address":                    true,
	"allow_provisioning":             true,
	"is_staging":                     true,
	"version":                        true,
	"oracle_base":                    true,
	"bits":                           true,
	"oracle_tde_keystores_root_path": true,
	"tags":                           true,
}

var isDestructiveUpdate = map[string]bool{
	"name":                           false,
	"cluster_home":                   true,
	"connector_port":                 true,
	"username":                       true,
	"password":                       true,
	"java_home":                      false,
	"description":                    false,
	"hostname":                       true,
	"ssh_port":                       true,
	"toolkit_path":                   true,
	"nfs_address":                    true,
	"allow_provisioning":             false,
	"is_staging":                     false,
	"version":                        false,
	"oracle_base":                    false,
	"bits":                           false,
	"oracle_tde_keystores_root_path": true,
	"tags":                           false,
}

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
	"name":               true,
	"cluster_home":       true,
	"connector_port":     true,
	"username":           true,
	"password":           true,
	"description":        true,
	"allow_provisioning": true,
	"is_staging":         true,
	"version":            true,
	"oracle_base":        true,
	"bits":               true,
	"tags":               true,
	"hosts":              true,
	// "java_home":      true,
	// "hostname":                       true,
	// "ssh_port":                       true,
	// "toolkit_path":                   true,
	// "nfs_addresses":                  true,
	// "oracle_tde_keystores_root_path": true,

}

var isDestructiveEnvUpdate = map[string]bool{
	"name":               false,
	"cluster_home":       true,
	"connector_port":     true,
	"username":           true,
	"password":           true,
	"description":        false,
	"allow_provisioning": false,
	"is_staging":         false,
	"version":            false,
	"oracle_base":        false,
	"bits":               false,
	"tags":               false,
	"hosts":              true,
	// "java_home":                      false,
	// "hostname":                       true,
	// "ssh_port":                       true,
	// "toolkit_path":                   true,
	// "nfs_addresses":                  true,
	// "oracle_tde_keystores_root_path": true,

}

var updatableVdbKeys = map[string]bool{
	"name":                  true,
	"db_username":           true,
	"db_password":           true,
	"pre_refresh":           true,
	"post_refresh":          true,
	"configure_clone":       true,
	"pre_snapshot":          true,
	"post_snapshot":         true,
	"pre_start":             true,
	"post_start":            true,
	"pre_stop":              true,
	"post_stop":             true,
	"template_id":           true,
	"pre_script":            true,
	"post_script":           true,
	"custom_env_vars":       true,
	"custom_env_files":      true,
	"appdata_source_params": true,
	"appdata_config_params": true,
	"config_params":         true,
}

var isDestructiveVdbUpdate = map[string]bool{
	"name":                  false,
	"db_username":           false,
	"db_password":           false,
	"pre_refresh":           false,
	"post_refresh":          false,
	"configure_clone":       false,
	"pre_snapshot":          false,
	"post_snapshot":         false,
	"pre_start":             false,
	"post_start":            false,
	"pre_stop":              false,
	"post_stop":             false,
	"template_id":           true,
	"pre_script":            false,
	"post_script":           false,
	"custom_env_vars":       false,
	"custom_env_files":      false,
	"appdata_source_params": true,
	"appdata_config_params": true,
	"config_params":         true,
}

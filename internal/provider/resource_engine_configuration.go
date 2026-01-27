package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEngineConfiguration() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Engine initialization.",

		CreateContext: engineConfigCreate,
		ReadContext:   engineConfigRead,
		UpdateContext: engineConfigUpdate,
		DeleteContext: engineConfigDelete,
		CustomizeDiff: func(ctx context.Context, rd *schema.ResourceDiff, i interface{}) error {

			device_type := rd.Get("device_type").(string)
			if device_type == OBJECT {
				ospList := rd.Get("object_storage_params").([]interface{})
				if len(ospList) == 0 {
					return errors.New("object_storage_params must be provided when device_type is OBJECT")
				}
				for _, item := range ospList {
					if item == nil {
						continue
					}
					block := item.(map[string]interface{})

					//Cloud Provider specific validations
					cloud_provider := block["cloud_provider"].(string)
					if cloud_provider == AWS {
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
						}

					} else if cloud_provider == AZURE {
						if _, ok := block["azure_container"]; !ok {
							return errors.New("azure_container must be provided in object_storage_params for AZURE cloud_provider")
						}

						if _, ok := block["azure_account"]; !ok {
							return errors.New("azure_account must be provided in object_storage_params for AZURE cloud_provider")
						}

						if authType, ok := block["auth_type"]; ok {
							authTypeStr := authType.(string)
							if authTypeStr != MANAGED_IDENTITIES && authTypeStr != ACCESS_KEY {
								return errors.New("auth_type for AZURE cloud_provider must be either MANAGED_IDENTITIES or ACCESS_KEY")
							}
						}
					}
					authType := block["auth_type"].(string)
					if authType == ACCESS_KEY && (block["access_id"] == "" || block["access_key"] == "") {
						return errors.New("access_id and access_key must be provided when auth_type is ACCESS_KEY")
					}

				}
				ntp_servers := rd.Get("ntp_servers").([]interface{})
				ntp_timezone := rd.Get("ntp_timezone").(string)
				if len(ntp_servers) == 0 || ntp_timezone == "" {
					return errors.New("ntp_servers and ntp_timezone must be provided when device_type is OBJECT")
				}
			}

			smtp_config := rd.Get("smtp_config").([]interface{})
			if len(smtp_config) > 0 {
				smtp_block := smtp_config[0].(map[string]interface{})
				if len(smtp_block["smtp_authentication"].([]interface{})) > 0 {
					if _, ok := smtp_block["smtp_authentication"].([]interface{})[0].(map[string]interface{})["user"]; !ok {
						return errors.New("username must be provided in smtp_authentication")
					}
					if _, ok := smtp_block["smtp_authentication"].([]interface{})[0].(map[string]interface{})["password"]; !ok {
						return errors.New("password must be provided in smtp_authentication")
					}
				}
			}

			engine_type := rd.Get("engine_type").(string)
			if engine_type == CONTINUOUS_COMPLIANCE {
				if _, ok := rd.GetOk("compliance_user"); !ok {
					return errors.New("compliance_user must be provided when engine_type is CONTINUOUS_COMPLIANCE")
				}
				if _, ok := rd.GetOk("compliance_password"); !ok {
					return errors.New("compliance_password must be provided when engine_type is CONTINUOUS_COMPLIANCE")
				}
				if _, ok := rd.GetOk("compliance_email"); !ok {
					return errors.New("compliance_email must be provided when engine_type is CONTINUOUS_COMPLIANCE")
				}
				if _, ok := rd.GetOk("compliance_new_password"); !ok {
					return errors.New("compliance_new_password must be provided when engine_type is CONTINUOUS_COMPLIANCE")
				}
			}
			return nil

		},

		Schema: map[string]*schema.Schema{
			"engine_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sys_user": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sys_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sys_new_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"user": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compliance_user": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"compliance_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"compliance_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_new_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"engine_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configured": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sso_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_qualifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled_features": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"toggleable_features": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"locales": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"banner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kernel_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{BLOCK, OBJECT}, false),
			},
			"object_storage_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_provider": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{AWS, AZURE}, false),
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"size": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateStorageSize,
						},
						"auth_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      ROLE,
							ValidateFunc: validation.StringInSlice([]string{ROLE, ACCESS_KEY, MANAGED_IDENTITIES}, false),
						},
						"s3_instance_profile": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "S3ObjectStoreAccessInstanceProfile",
						},
						"azure_managed_identities": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "BlobObjectStoreAccessManagedIdentities",
						},
						"access_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"access_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"azure_container": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"azure_account": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ntp_timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"smtp_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"from_email_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tls_authentication": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"send_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  DEFAULT_SEND_TIMEOUT,
						},
						"smtp_authentication": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user": {
										Type:     schema.TypeString,
										Required: true,
									},
									"password": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
								},
							},
						},
					},
				},
			},
			"dns_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"servers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.IsIPAddress, // Optional: validate IP addresses
							},
						},
						"domains": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringMatch(
									regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-\.]*[a-zA-Z0-9]$`),
									"must be a valid domain name",
								),
							},
						},
					},
				},
			},
			"phone_home_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_analytics_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"web_proxy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
			"sso_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"saml_metadata": {
							Type:     schema.TypeString,
							Required: true,
						},
						"response_skew_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  DEFAULT_SSO_SKEW_TIME,
						},
						"max_authentication_age": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  DEFAULT_SSO_MAX_AUTH_AGE,
						},
					},
				},
			},
		},
	}
}

func engineConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	var compl_email, compl_user, compl_password, compl_new_password, default_email string
	engine_host, _ := d.Get("engine_host").(string)

	//hardcoded for backward compatibility
	version := ENGINE_API_VERSION

	sys_user, _ := d.Get("sys_user").(string)
	sys_curr_pass, _ := d.Get("sys_password").(string)
	sys_new_pass, _ := d.Get("sys_new_password").(string)
	user, _ := d.Get("user").(string)
	email, has_email := d.GetOk("email")
	if has_email {
		default_email = email.(string)
	}
	password := d.Get("password").(string)
	engine_type := d.Get("engine_type").(string)
	device_type := d.Get("device_type").(string)
	ntp_timezone := d.Get("ntp_timezone").(string)
	ntp_servers_raw := d.Get("ntp_servers").([]interface{})
	ntp_servers := make([]string, len(ntp_servers_raw))

	for i, server := range ntp_servers_raw {
		ntp_servers[i] = server.(string)
	}

	smtp_config := d.Get("smtp_config").([]interface{})
	dns_config := d.Get("dns_config").([]interface{})
	phonehome := d.Get("phone_home_enabled").(bool)
	useranalytics := d.Get("user_analytics_enabled").(bool)
	web_proxy_config := d.Get("web_proxy_config").([]interface{})
	sso_config := d.Get("sso_config").([]interface{})

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] start Session for "+engine_host)
	err := startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] login as "+sys_user)
	err = login(ctx, client, engine_host, sys_user, sys_curr_pass, SYSTEM)
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	if engine_type == CONTINUOUS_COMPLIANCE {
		compl_user = d.Get("compliance_user").(string)
		compl_password = d.Get("compliance_password").(string)
		compl_email = d.Get("compliance_email").(string)
		compl_new_password = d.Get("compliance_new_password").(string)
		readDiag := startMasking(ctx, client, engine_host)
		if readDiag.HasError() {
			return readDiag
		}
		// Sleep for 10 seconds to allow masking service to start
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sleeping for 10 seconds to allow masking service to start")
		time.Sleep(time.Duration(10) * time.Second)

		token, err := loginComplianceUser(ctx, client, engine_host, compl_user, compl_password)
		if err != nil {
			return diag.Errorf("["+engine_host+"]Error logging in as compliance user: %s", err)
		}

		readDiags := updateComplianceUserDetails(ctx, client, engine_host, compl_password, compl_new_password, compl_email, compl_user, token)
		if readDiags.HasError() {
			return readDiags
		}
	}

	params := InitializationParameters{
		User:       user,
		Password:   password,
		Email:      default_email,
		DeviceType: device_type,
	}
	if device_type == OBJECT {
		object_storage_params := d.Get("object_storage_params").([]interface{})
		params.CloudProvider = object_storage_params[0].(map[string]interface{})["cloud_provider"].(string)
		params.Size = object_storage_params[0].(map[string]interface{})["size"].(string)
		params.AuthType = object_storage_params[0].(map[string]interface{})["auth_type"].(string)

		if params.CloudProvider == AWS {
			params.Endpoint = object_storage_params[0].(map[string]interface{})["endpoint"].(string)
			params.Region = object_storage_params[0].(map[string]interface{})["region"].(string)
			params.Bucket = object_storage_params[0].(map[string]interface{})["bucket"].(string)

			if params.AuthType == ACCESS_KEY {
				params.ACCESS_ID = object_storage_params[0].(map[string]interface{})["access_id"].(string)
				params.ACCESS_KEY = object_storage_params[0].(map[string]interface{})["access_key"].(string)
			} else {
				params.S3_INSTANCE_PROFILE = object_storage_params[0].(map[string]interface{})["s3_instance_profile"].(string)
			}
		} else if params.CloudProvider == AZURE {
			params.Container = object_storage_params[0].(map[string]interface{})["azure_container"].(string)
			params.AzureAccount = object_storage_params[0].(map[string]interface{})["azure_account"].(string)

			if params.AuthType == ACCESS_KEY {
				params.ACCESS_ID = object_storage_params[0].(map[string]interface{})["access_id"].(string)
				params.ACCESS_KEY = object_storage_params[0].(map[string]interface{})["access_key"].(string)
			} else {
				params.AzureManagedIdentities = object_storage_params[0].(map[string]interface{})["azure_managed_identities"].(string)
			}
		}

	}

	configTasks := []ConfigTask{
		{
			Name:      "SMTP",
			Condition: len(smtp_config) > 0,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring SMTP settings")
				smtp_block := smtp_config[0].(map[string]interface{})
				_, err := configureSMTP(ctx, client, engine_host, smtp_block)
				return err
			},
		},
		{
			Name:      "DNS",
			Condition: len(dns_config) > 0,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring DNS settings")
				dns_block := dns_config[0].(map[string]interface{})
				_, err := configureDNS(ctx, client, engine_host, dns_block)
				return err
			},
		},
		{
			Name:      "Phone Home",
			Condition: phonehome,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Enabling Phone Home")
				_, err := configurePhoneHome(ctx, client, engine_host, phonehome)
				return err
			},
		},
		{
			Name:      "User Analytics",
			Condition: useranalytics,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Enabling User Analytics")
				_, err := configureUserAnalytics(ctx, client, engine_host, useranalytics)
				return err
			},
		},
		{
			Name:      "Web Proxy",
			Condition: len(web_proxy_config) > 0,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring Web Proxy settings")
				web_proxy_block := web_proxy_config[0].(map[string]interface{})
				_, err := configureWebProxy(ctx, client, engine_host, web_proxy_block)
				return err
			},
		},
		{
			Name:      "NTP Servers",
			Condition: len(ntp_servers) > 0,
			Task: func() error {
				tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring NTP servers")
				_, err := setNtpServers(ctx, client, engine_host, ntp_servers, ntp_timezone)
				return err
			},
		},
	}

	for _, config := range configTasks {
		if config.Condition {
			tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Starting "+config.Name+" configuration")
			if err := config.Task(); err != nil {
				return diag.Errorf("["+engine_host+"] Error configuring %s: %s", config.Name, err)
			}
			tflog.Info(ctx, DLPX+INFO+config.Name+" ["+engine_host+"] configuration completed successfully")
		}
	}

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sleeping for 10 seconds before initialization for management services to restart properly")
	time.Sleep(time.Duration(10) * time.Second)

	// Initialize Engine
	result, readDiags := initializeSystemAndDevices(ctx, client, engine_host, params)
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Initialization action result: "+fmt.Sprintf("%+v", readDiags))

	if readDiags.HasError() {
		if engine_type == CONTINUOUS_COMPLIANCE {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Engine initialization failed, Rolling back compliance user password change!!")
			token, err := loginComplianceUser(ctx, client, engine_host, compl_user, compl_new_password)
			if err != nil {
				return diag.Errorf("["+engine_host+"]Error logging in as compliance user: %s", err)
			}
			readDiags := updateComplianceUserDetails(ctx, client, engine_host, compl_new_password, compl_password, compl_email, compl_user, token)
			if readDiags.HasError() {
				return readDiags
			}
		}
		return readDiags
	}

	// Poll for initialization to complete
	readDiags = pollActionStatus(ctx, client, engine_host, result.Action)
	if readDiags.HasError() {
		return readDiags
	}

	// Sleep a minute for the initialization to complete
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sleeping for 60 seconds to allow engine to restart after initialization")
	time.Sleep(time.Duration(60) * time.Second)

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"]start Session for "+engine_host)
	err = startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] login as "+sys_user)
	err = login(ctx, client, engine_host, sys_user, sys_curr_pass, SYSTEM)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error logging in: %s", err)
	}

	resp, err := setEngineType(ctx, client, engine_host, engine_type)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error setting engine type: %s", err)
	}

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] engine type resp "+string(resp))

	// Configure SSO
	if len(sso_config) > 0 {
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring SSO settings")
		sso_block := sso_config[0].(map[string]interface{})
		_, ssoErr := configureSSO(ctx, client, engine_host, sso_block)
		if ssoErr != nil {
			return diag.Errorf("["+engine_host+"] Error configuring SSO: %s", ssoErr)
		}
	}

	//Update defaultUser password
	readDiags = UpdateUserPassword(ctx, client, engine_host, version, user, password, password, "DOMAIN")
	if readDiags.HasError() {
		return readDiags
	}

	// Update sysadmin_user password
	readDiag := UpdateUserPassword(ctx, client, engine_host, version, sys_user, sys_curr_pass, sys_new_pass, SYSTEM)
	if readDiag.HasError() {
		return readDiag
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] System initialization successful!")

	d.SetId(engine_host)

	readDiags = engineConfigRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func engineConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}
	// revert and set the old value to the changed keys
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}

	return diag.Errorf("Action update not available for engine config")
}

func engineConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, DLPX+INFO+"Reading engine configuration")
	engineId := d.Id()
	version := ENGINE_API_VERSION

	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Start a session
	err := startSession(ctx, client, engineId, version)
	if err != nil {
		return diag.Errorf("Error starting session: %v", err)
	}

	// Authenticate/login
	err = login(ctx, client, engineId, d.Get("sys_user").(string), d.Get("sys_new_password").(string), SYSTEM)
	if err != nil {
		return diag.Errorf("Error logging in: %v", err)
	}

	body, err := getSystem(ctx, client, engineId)
	if err != nil {
		return diag.Errorf("Error getting system info: %v", err)
	}

	var response SystemInfoResponse
	sysErr := json.Unmarshal(body, &response)
	if sysErr != nil {
		return diag.Errorf("Error unmarshalling system info response: %v", sysErr)
	}

	d.Set("configured", response.Result["configured"])
	d.Set("hostname", response.Result["hostname"])
	d.Set("product_type", response.Result["productType"])
	//d.Set("engineType", response.Result["engineType"])
	d.Set("sso_enabled", response.Result["ssoEnabled"])
	d.Set("vendor_name", response.Result["vendorName"])
	d.Set("product_name", response.Result["productName"])
	d.Set("engine_qualifier", response.Result["engineQualifier"])
	d.Set("support_url", response.Result["supportURL"])
	d.Set("build_title", response.Result["buildTitle"])
	d.Set("build_timestamp", response.Result["buildTimestamp"])
	// d.Set("build_version", response.Result["buildVersion"])
	d.Set("enabled_features", response.Result["enabledFeatures"])
	d.Set("toggleable_features", response.Result["toggleableFeatures"])
	//d.Set("apiVersion", response.Result["apiVersion"])
	d.Set("banner", response.Result["banner"])
	d.Set("locales", response.Result["locales"])
	d.Set("current_locale", response.Result["currentLocale"])
	d.Set("kernel_name", response.Result["kernelName"])
	d.Set("platform", response.Result["platform"])

	return diags
}

func engineConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if os.Getenv("TF_ACC") == "1" {
		// Terraform acceptance test mode (destroy MUST succeed)
		d.SetId("")
		return nil
	}

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}
	// revert and set the old value to the changed keys
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}
	return diag.Errorf("Action delete not available for engine config")
}

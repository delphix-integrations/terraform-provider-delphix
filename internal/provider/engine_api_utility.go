package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func pollActionStatus(ctx context.Context, client *http.Client, engine_host string, action string) diag.Diagnostics {
	var diags diag.Diagnostics
	for {
		// Get Action
		tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] action_id "+action)
		actionData, err := getAction(ctx, client, engine_host, action)
		if err != nil {
			return diag.Errorf("["+engine_host+"] Error getting action: %s", err)
		}
		tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] action data "+string(actionData))
		var actionResult ActionResult
		err = json.Unmarshal(actionData, &actionResult)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+" ["+engine_host+"] Error unmarshalling "+err.Error())
			return diag.Errorf("["+engine_host+"] Error unmarshalling action result: %s", err)
		}
		tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] action state "+actionResult.Result.State)
		if actionResult.Result.State == "COMPLETED" {
			tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] Action completed!")
			break
		}
		time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
	}
	return diags
}

func UpdateUserPassword(ctx context.Context, client *http.Client, engine_host string, version string, user string, curr_pass string, new_pass string, target string) diag.Diagnostics {
	var diags diag.Diagnostics

	// Start a session
	tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] start Session")
	err := startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+" ["+engine_host+"] login as "+user)
	err = login(ctx, client, engine_host, user, curr_pass, target)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error logging in: %s", err)
	}

	// Get User Details
	userData, err := getCurrentUser(ctx, client, engine_host)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error getting curr user: %s", err)
	}

	var result ActionResult
	err = json.Unmarshal(userData, &result)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Unmarshal Error "+err.Error())
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Current User Reference "+result.Result.Reference)

	// Update sysadmin Password
	updateResp, updateErr := updatePassword(ctx, client, engine_host, result.Result.Reference, new_pass)
	if updateErr != nil {
		return diag.Errorf("["+engine_host+"] Error in update Password: %s", updateErr)
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] update password "+result.Result.Reference+" "+string(updateResp))

	return diags
}

func initializeSystemAndDevices(ctx context.Context, client *http.Client, engine_host string, params InitializationParameters) (APIResponse, diag.Diagnostics) {
	var diags diag.Diagnostics
	// Get Devices
	deviceData, err := getDevices(ctx, client, engine_host)
	if err != nil {
		return APIResponse{}, diag.Errorf("["+engine_host+"] Error getting devices: %s", err)
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] devices "+string(deviceData))

	// Parse device information
	var resultList ListResult
	err = json.Unmarshal(deviceData, &resultList)
	if err != nil {
		return APIResponse{}, diag.Errorf("["+engine_host+"] Error parsing device information: %s", err)
	}

	// Initialize System
	resp, err := initializeSystem(ctx, client, engine_host, resultList, params)
	if err != nil {
		return APIResponse{}, diag.Errorf("["+engine_host+"] Error initializing system: %s", err)
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Initializing system: "+string(resp))

	var result APIResponse
	unmarshalErr := json.Unmarshal(resp, &result)
	if unmarshalErr != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Error unmarshalling: "+unmarshalErr.Error())
		return APIResponse{}, diag.Errorf("["+engine_host+"] Error initializing system: %s", unmarshalErr)
	}
	return result, diags
}

func setNtpServers(ctx context.Context, client *http.Client, engine_host string, ntp_servers []string, ntp_timezone string) (APIResponse, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Setting NTP Servers")
	// Get default Timezone from Engine
	_, timezone, err := getNtpServersAndTimezones(ctx, client, engine_host)
	if err != nil {
		return APIResponse{}, nil
	}

	if ntp_timezone != "" {
		timezone = ntp_timezone
	}
	// Set NTP Servers on Engine
	ntpPayload := NTPServerParams{
		Type:           "TimeConfig",
		SystemTimeZone: timezone,
		NTPConfig: SetNTPConfig{
			Type:    "NTPConfig",
			Enabled: true,
			Servers: ntp_servers,
		},
	}

	ntpURL := engine_host + ENGINE_APIS["NTP_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, ntpPayload, ntpURL, "NTP", engine_host)
	return res, err
}

func configureSMTP(ctx context.Context, client *http.Client, engine_host string, smtp_config map[string]interface{}) (APIResponse, error) {
	var config SMTPConfig
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring SMTP Settings")
	var isSMTPAuthentication bool
	if len(smtp_config["smtp_authentication"].([]interface{})) > 0 {
		isSMTPAuthentication = true
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] SMTP Config Map: "+fmt.Sprintf("%+v", smtp_config))
	config = SMTPConfig{
		Type:                  "SMTPConfig",
		Enabled:               true,
		Server:                smtp_config["server"].(string),
		Port:                  smtp_config["port"].(int),
		AuthenticationEnabled: isSMTPAuthentication,
		TlsEnabled:            smtp_config["tls_authentication"].(bool),
		FromAddress:           smtp_config["from_email_address"].(string),
		SendTimeout:           smtp_config["send_timeout"].(int),
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] SMTP Config before adding auth details: "+fmt.Sprintf("%+v", config))
	if len(smtp_config["smtp_authentication"].([]interface{})) > 0 {
		config.Username = smtp_config["smtp_authentication"].([]interface{})[0].(map[string]interface{})["user"].(string)
		config.Password = smtp_config["smtp_authentication"].([]interface{})[0].(map[string]interface{})["password"].(string)
	}

	smtpURL := engine_host + ENGINE_APIS["SMTP_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, config, smtpURL, "SMTP", engine_host)
	return res, err
}

func configureDNS(ctx context.Context, client *http.Client, engine_host string, dns_config map[string]interface{}) (APIResponse, error) {
	existingDnsConfig, err := getDNSConfiguration(ctx, client, engine_host)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting existing DNS configuration: "+err.Error())
		return APIResponse{}, err
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Existing DNS Configuration: "+fmt.Sprintf("%+v", existingDnsConfig))
	var dnsServers []string
	var domains []string

	if len(existingDnsConfig.Result.Servers) > 0 {
		dnsServers = existingDnsConfig.Result.Servers
	}
	if len(existingDnsConfig.Result.Domains) > 0 {
		domains = existingDnsConfig.Result.Domains
	}

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Existing Domains"+fmt.Sprintf("%+v", domains))
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Existing DNS Servers"+fmt.Sprintf("%+v", dnsServers))
	dnsURL := engine_host + ENGINE_APIS["DNS_CONFIG"]
	var dnsPayload DNSConfig
	if len(toStringArray(dns_config["servers"])) > 0 {
		dnsServers = append(dnsServers, toStringArray(dns_config["servers"])...)
	}
	dnsPayload.Servers = dnsServers
	if len(toStringArray(dns_config["domains"])) > 0 {
		domains = append(domains, toStringArray(dns_config["domains"])...)
	}
	dnsPayload.Domains = domains
	dnsPayload.Type = "DNSConfig"
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Final DNS Config Struct: "+fmt.Sprintf("%+v", dnsPayload))
	res, err := processRequestAndResponse(ctx, client, dnsPayload, dnsURL, "DNS", engine_host)
	return res, err
}

func configurePhoneHome(ctx context.Context, client *http.Client, engine_host string, enable bool) (APIResponse, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring Phone Home Setting to "+fmt.Sprintf("%t", enable))
	phoneHomePayload := PhoneHomeConfig{
		Type:    "PhoneHomeService",
		Enabled: enable,
	}
	phoneHomeURL := engine_host + ENGINE_APIS["PHONE_HOME_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, phoneHomePayload, phoneHomeURL, "Phone Home", engine_host)
	return res, err
}

func configureUserAnalytics(ctx context.Context, client *http.Client, engine_host string, enable bool) (APIResponse, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring User Analytics Setting to "+fmt.Sprintf("%t", enable))
	userAnalyticsPayload := UserAnalyticsConfig{
		Type:             "UserInterfaceConfig",
		AnalyticsEnabled: enable,
	}
	userAnalyticsURL := engine_host + ENGINE_APIS["USER_ANALYTICS_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, userAnalyticsPayload, userAnalyticsURL, "User Analytics", engine_host)
	return res, err
}

func configureWebProxy(ctx context.Context, client *http.Client, engine_host string, web_proxy_config map[string]interface{}) (APIResponse, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring Web Proxy Settings")
	webProxyPayload := WebProxyConfig{
		Type: "ProxyService",
		Https: &ProxyConfiguration{
			Host:    web_proxy_config["host"].(string),
			Port:    web_proxy_config["port"].(int),
			Enabled: true,
			Type:    "ProxyConfiguration",
		},
	}

	if val, ok := web_proxy_config["username"]; ok && val.(string) != "" {
		webProxyPayload.Https.Username = web_proxy_config["username"].(string)
	}
	if val, ok := web_proxy_config["password"]; ok && val.(string) != "" {
		webProxyPayload.Https.Password = web_proxy_config["password"].(string)
	}

	webProxyURL := engine_host + ENGINE_APIS["WEB_PROXY_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, webProxyPayload, webProxyURL, "Web Proxy", engine_host)
	return res, err
}

func configureSSO(ctx context.Context, client *http.Client, engine_host string, sso_config map[string]interface{}) (APIResponse, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring SSO Settings")
	ssoPayload := SSOConfig{
		Type:         "SsoConfig",
		Enabled:      sso_config["enabled"].(bool),
		SamlMetadata: sso_config["saml_metadata"].(string),
	}

	if val, ok := sso_config["response_skew_time"]; ok && val != 0 {
		ssoPayload.ResponseSkewTime = sso_config["response_skew_time"].(int)
	} else {
		ssoPayload.ResponseSkewTime = DEFAULT_SSO_SKEW_TIME
	}

	if val, ok := sso_config["max_authentication_age"]; ok && val != 0 {
		ssoPayload.MaxAuthenticationAge = sso_config["max_authentication_age"].(int)
	} else {
		ssoPayload.MaxAuthenticationAge = DEFAULT_SSO_MAX_AUTH_AGE
	}

	entityId, err := getEntityIDForSSO(ctx, client, engine_host)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting existing SSO configuration: "+err.Error())
		return APIResponse{}, err
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Existing SSO Configuration: "+fmt.Sprintf("%+v", entityId))

	if entityId != "" {
		ssoPayload.EntityId = entityId
	}

	ssoURL := engine_host + ENGINE_APIS["SSO_CONFIG"]
	res, err := processRequestAndResponse(ctx, client, ssoPayload, ssoURL, "SSO", engine_host)
	return res, err
}

func validateStorageSize(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	// Regular expression to match number followed by GB, TB, or PB (case insensitive)
	pattern := `^\d+(?:\.\d+)?\s*(GB|TB|PB)$`
	matched, err := regexp.MatchString(pattern, value)

	if err != nil {
		errors = append(errors, fmt.Errorf("error validating %s: %s", k, err))
		return
	}

	if !matched {
		errors = append(errors, fmt.Errorf("%s must be a valid storage size with units (e.g., '20TB', '500GB', '1.5PB')", k))
		return
	}

	return
}

func processRequestAndResponse(ctx context.Context, client *http.Client, payload interface{}, apiURL string, config_name string, engine_host string) (APIResponse, error) {
	postData, err := json.Marshal(payload)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling "+config_name+" configuration: "+err.Error())
		return APIResponse{}, err
	}

	req, er := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(postData))
	if er != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"]error creating "+config_name+" request: "+er.Error())
		return APIResponse{}, er
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"]error configuring "+config_name+": "+err.Error())
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	bodyStr := string(body)
	if strings.Contains(bodyStr, `"status":"ERROR"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return APIResponse{}, fmt.Errorf("["+engine_host+"] %s configuration failed: %s", config_name, bodyStr)
	}

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] "+config_name+" Configuration Response Body: "+string(body))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading "+config_name+" response: "+err.Error())
		return APIResponse{}, err
	}
	var res APIResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshalling "+config_name+" response: "+err.Error())
		return APIResponse{}, err
	}
	return res, nil
}

func updateComplianceUserDetails(ctx context.Context, client *http.Client, engine_host string, comp_pass string,
	comp_new_pass string, email string, comp_user string, token string) diag.Diagnostics {

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Configuring Compliance User details for user "+comp_user)
	var diags diag.Diagnostics
	// Get User Details
	userId, err := getComplianceUserID(ctx, client, engine_host, comp_user, token)
	if err != nil {
		return diag.Errorf("["+engine_host+"] Error getting curr user: %s", err)
	}

	// Update compliance user details
	updateResp, updateErr := updateComplianceUser(ctx, client, engine_host, userId, comp_new_pass, email, comp_user, token)
	if updateErr != nil {
		return diag.Errorf("["+engine_host+"] Error in updating compliance user details: %s", updateErr)
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] update compliance user details "+userId+" "+string(updateResp))
	return diags
}

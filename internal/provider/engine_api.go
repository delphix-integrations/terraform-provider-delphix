package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// startSession starts a new session
func startSession(ctx context.Context, client *http.Client, engine_host string, version string) error {
	versionParts := strings.Split(version, ".")
	major, _ := strconv.Atoi(versionParts[0])
	minor, _ := strconv.Atoi(versionParts[1])
	micro, _ := strconv.Atoi(versionParts[2])
	tflog.Info(ctx, DLPX+INFO+"start session for "+engine_host+" version "+version)
	sessionURL := engine_host + ENGINE_APIS["SESSION"]
	apisessionData := APISession{
		Version: APIVersion{
			Minor: minor,
			Major: major,
			Micro: micro,
			Type:  "APIVersion",
		},
		Type: "APISession",
	}

	sessionData, err := json.Marshal(apisessionData)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling session data: "+err.Error())
		return err
	}

	req, err := http.NewRequest(http.MethodPost, sessionURL, bytes.NewReader(sessionData))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating session request: "+err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error starting session: "+err.Error())
		return err
	}
	defer resp.Body.Close()

	// Check for successful session start (can be modified to return response body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error starting session: "+strconv.Itoa(resp.StatusCode)+" "+string(body))
		return errors.New("[" + engine_host + "] error to start session: " + strconv.Itoa(resp.StatusCode) + " " + string(body))
	}
	return nil
}

func login(ctx context.Context, client *http.Client, engine_host string, user string, password string, target string) error {
	loginURL := engine_host + ENGINE_APIS["LOGIN"]
	loginData := LoginRequest{
		Password: password,
		Type:     "LoginRequest",
		Username: user,
		Target:   target,
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Logging in user "+user+" to target "+target)
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling login data: "+err.Error())
		return err
	}
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] Login JSON payload created: %d bytes", len(loginJSON)))
	// Clear sensitive data from loginJSON after use
	defer func() {
		tflog.Debug(ctx, "[SECURITY] Clearing login credentials from memory")
		SecureClearByteSlice(ctx, loginJSON)
		tflog.Debug(ctx, "[SECURITY] Login credentials cleared")
	}()
	
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Login ")
	req, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewReader(loginJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating login request: "+err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sending login request to "+loginURL)
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error authenticating: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Received login response with status code "+strconv.Itoa(resp.StatusCode))
	// Check for successful login (can be modified to return response body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Failed to login: "+strconv.Itoa(resp.StatusCode)+" "+string(body))
		return errors.New("[" + engine_host + "] Failed to login: " + strconv.Itoa(resp.StatusCode) + " " + string(body))
	}
	return nil
}

func getDevices(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	deviceURL := engine_host + ENGINE_APIS["STORAGE_DEVICE"]
	req, err := http.NewRequest(http.MethodGet, deviceURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating device request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting device: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading devices response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func getCurrentUser(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	userURL := engine_host + ENGINE_APIS["USER"] + "current"
	req, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating current user request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting current user: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading current user: "+err.Error())
		return nil, err
	}
	return body, nil
}

func initializeSystem(ctx context.Context, client *http.Client, engine_host string,
	resultList ListResult, params InitializationParameters) ([]byte, error) {

	initializeSystemURL := engine_host + ENGINE_APIS["SYSTEM_INITIALIZATION"]

	var deviceRefs []string
	for _, device := range resultList.Result {
		if !device.Configured {
			deviceRefs = append(deviceRefs, device.Reference)
		}
	}

	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"]Unconfigured Devices: "+fmt.Sprintf("%v", deviceRefs))
	var email string
	var initializationParams interface{}
	if params.Email != "" {
		email = params.Email
	}

	if params.DeviceType == OBJECT {
		var objectStorage *ObjectStore
		sizeInBytes, err := convertStorageToBytes(params.Size)
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Converted size in bytes: "+strconv.FormatInt(int64(sizeInBytes), 10))
		if err != nil {
			return nil, err
		}

		switch params.CloudProvider {
		case AWS:
			objectStorage = &ObjectStore{
				Type:         "S3ObjectStore",
				Size:         sizeInBytes,
				CacheDevices: deviceRefs,
				Endpoint:     params.Endpoint,
				Region:       params.Region,
				Bucket:       params.Bucket,
			}
		case AZURE:
			objectStorage = &ObjectStore{
				Type:         "BlobObjectStore",
				Size:         sizeInBytes,
				CacheDevices: deviceRefs,
				Container:    params.Container,
			}
		}

		if params.CloudProvider == AWS {
			switch params.AuthType {
			case ROLE:
				objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
					Type: params.S3_INSTANCE_PROFILE,
				}
			case ACCESS_KEY:
				objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
					Type:       "S3ObjectStoreAccessKey",
					ACCESS_ID:  params.ACCESS_ID,
					ACCESS_KEY: params.ACCESS_KEY,
				}
			}
		} else if params.CloudProvider == AZURE {
			switch params.AuthType {
			case ACCESS_KEY:
				objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
					Type:          "BlobObjectStoreAccessKey",
					AZURE_ACCOUNT: params.AZURE_ACCOUNT,
					AZURE_KEY:     params.AZURE_KEY,
				}
			case MANAGED_IDENTITIES:
				objectStorage.AccessCredentials = &ObjectStoreAccessCredentials{
					Type:          params.AzureManagedIdentities,
					AZURE_ACCOUNT: params.AZURE_ACCOUNT,
				}
			}
		}

		initializationParams = SystemInitializationObjectStore{
			Type:            "SystemInitializationParameters",
			DefaultUser:     params.User,
			DefaultPassword: params.Password,
			DefaultEmail:    email,
			ObjectStore:     objectStorage,
		}
		resBody, er := testConnectionForObjectStore(ctx, client, engine_host, params)
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] test connection response body: "+string(resBody))
		if er != nil {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error testing object store connection: "+er.Error())
			return nil, er
		}
		bodyStr := string(resBody)
		if strings.Contains(bodyStr, `"status":"ERROR"`) {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
			return nil, fmt.Errorf("initialization failed: %s", bodyStr)
		}
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] test conncection to object store successful.")

	} else {
		devices := deviceRefs
		initializationParams = SystemInitializationBlockStorage{
			Type:            "SystemInitializationParameters",
			DefaultUser:     params.User,
			DefaultPassword: params.Password,
			DefaultEmail:    email,
			Devices:         devices,
		}

	}

	postData, err := json.Marshal(initializationParams)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling initialization parameters: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, initializeSystemURL, bytes.NewReader(postData))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating initialize system request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error initializing system: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading initializing system response: "+err.Error())
		return nil, err
	}

	bodyStr := string(body)
	if strings.Contains(bodyStr, `"status":"ERROR"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return nil, fmt.Errorf("["+engine_host+"] initialization failed: %s", bodyStr)
	}
	return body, nil
}

func getAction(ctx context.Context, client *http.Client, engine_host string, action_id string) ([]byte, error) {
	actionURL := engine_host + ENGINE_APIS["ACTION"] + action_id
	req, err := http.NewRequest(http.MethodGet, actionURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating action request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting action: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading action response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func getSystem(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	actionURL := engine_host + ENGINE_APIS["SYSTEM_INFO"]
	req, err := http.NewRequest(http.MethodGet, actionURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating system request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting system: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading system response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func updatePassword(ctx context.Context, client *http.Client, engine_host string, user_id string, password string) ([]byte, error) {
	updateURL := engine_host + ENGINE_APIS["USER"] + user_id + "/updateCredential"
	// Create credential update parameters
	UpdateParameters := CredentialUpdateParameters{
		Type: "CredentialUpdateParameters",
		NewCredential: struct {
			Type     string `json:"type"`
			Password string `json:"password"`
		}{
			Type:     "PasswordCredential",
			Password: password,
		},
	}
	UpdateParametersJSON, err := json.Marshal(UpdateParameters)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling login data: "+err.Error())
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] Password update payload created: %d bytes for user %s", len(UpdateParametersJSON), user_id))
	// Clear sensitive data from UpdateParametersJSON after use
	defer func() {
		tflog.Debug(ctx, "[SECURITY] Clearing password update credentials from memory")
		SecureClearByteSlice(ctx, UpdateParametersJSON)
		tflog.Debug(ctx, "[SECURITY] Password update credentials cleared")
	}()

	req, err := http.NewRequest(http.MethodPost, updateURL, bytes.NewReader(UpdateParametersJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating login request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error authenticating: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func setEngineType(ctx context.Context, client *http.Client, engine_host string, engine_type string) ([]byte, error) {
	tflog.Info(ctx, DLPX+INFO+"Setting engine type to "+engine_type)
	updateURL := engine_host + ENGINE_APIS["SYSTEM_INFO"]
	var eg_type string

	switch engine_type {
	case CONTINUOUS_COMPLIANCE:
		eg_type = "MASKING"
	case CONTINUOUS_DATA:
		eg_type = "VIRTUALIZATION"
	default:
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Unknown engine type: "+engine_type)
	}
	// Prepare system information data
	data := SystemInfo{
		Type:       "SystemInfo",
		EngineType: eg_type,
	}
	systemInfoJSON, err := json.Marshal(data)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling login data: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, updateURL, bytes.NewReader(systemInfoJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error setting engine type: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func testConnectionForObjectStore(ctx context.Context, client *http.Client, engine_host string, params InitializationParameters) ([]byte, error) {
	testConnectionURL := engine_host + ENGINE_APIS["OBJECT_STORE_TEST_CONNECTION"]
	var payload TestConnection

	if params.CloudProvider == AWS {
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Testing connection for AWS S3 object store")
		if params.AuthType == ACCESS_KEY {
			payload = TestConnection{
				Type:     "S3ObjectStoreTest",
				Endpoint: params.Endpoint,
				Region:   params.Region,
				Bucket:   params.Bucket,
				AccessCredentials: ObjectStoreAccessCredentials{
					Type:       "S3ObjectStoreAccessKey",
					ACCESS_ID:  params.ACCESS_ID,
					ACCESS_KEY: params.ACCESS_KEY,
				},
			}
		} else {
			tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Using instance profile for S3")
			payload = TestConnection{
				Type:     "S3ObjectStoreTest",
				Endpoint: params.Endpoint,
				Region:   params.Region,
				Bucket:   params.Bucket,
				AccessCredentials: ObjectStoreAccessCredentials{
					Type: params.S3_INSTANCE_PROFILE,
				},
			}
		}
	} else if params.CloudProvider == AZURE {
		tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Testing connection for AZURE Blob object store")
		if params.AuthType == MANAGED_IDENTITIES {
			payload = TestConnection{
				Type:      "BlobObjectStoreTest",
				Container: params.Container,
				AccessCredentials: ObjectStoreAccessCredentials{
					Type:          params.AzureManagedIdentities,
					AZURE_ACCOUNT: params.AZURE_ACCOUNT,
				},
			}
		} else if params.AuthType == ACCESS_KEY {
			payload = TestConnection{
				Type:      "BlobObjectStoreTest",
				Container: params.Container,
				AccessCredentials: ObjectStoreAccessCredentials{
					Type:          "BlobObjectStoreAccessKey",
					AZURE_ACCOUNT: params.AZURE_ACCOUNT,
					AZURE_KEY:     params.AZURE_KEY,
				},
			}
		}
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling test connection data: "+err.Error())
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, testConnectionURL, bytes.NewReader(payloadJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating test connection request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] test connection response status: "+fmt.Sprintf("%v", resp))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error testing connection: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading test connection response: "+err.Error())
		return nil, err
	}

	var testResult TestConnectionResult
	if unmarshallErr := json.Unmarshal(body, &testResult); unmarshallErr == nil {
		if testResult.Status == "OK" && testResult.Result.Result {
			// Test connection successful
			tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Object store connection test successful")
			return body, nil
		} else if testResult.Status == "OK" && !testResult.Result.Result {
			// Test connection failed with specific error
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] Object store connection test failed: "+testResult.Result.ErrorMessage)
			return nil, fmt.Errorf("["+engine_host+"] object store connection test failed: %s", testResult.Result.ErrorMessage)
		} else if testResult.Status == "ERROR" {
			// API level error
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API error during connection test: "+string(body))
			return nil, fmt.Errorf("["+engine_host+"] API error during connection test: %s", string(body))
		}
	}
	return body, nil
}

// convertStorageToBytes converts storage size strings like "20TB", "500GB", "1.5PB" to bytes
func convertStorageToBytes(sizeStr string) (int, error) {
	const (
		BYTE = 1
		KB   = 1024 * BYTE
		MB   = 1024 * KB
		GB   = 1024 * MB
		TB   = 1024 * GB
		PB   = 1024 * TB
	)
	// Remove any spaces and convert to uppercase
	sizeStr = strings.TrimSpace(sizeStr)
	sizeStr = strings.ToUpper(sizeStr)

	// Regular expression to extract number and unit
	re := regexp.MustCompile(`^(\d+(?:\.\d+)?)\s*(GB|TB|PB)$`)
	matches := re.FindStringSubmatch(sizeStr)

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid size format: %q. Expected format like '20TB', '500GB', '1.5PB'", sizeStr)
	}

	// Parse the numeric value
	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse numeric value from %q: %v", sizeStr, err)
	}

	// Get the unit multiplier
	unit := matches[2]
	var multiplier int64

	switch unit {
	case "GB":
		multiplier = GB
	case "TB":
		multiplier = TB
	case "PB":
		multiplier = PB
	default:
		return 0, fmt.Errorf("unsupported storage unit: %q. Supported units: GB, TB, PB", unit)
	}

	// Calculate total bytes
	totalBytes := int64(value * float64(multiplier))

	return int(totalBytes), nil
}

func getNtpServersAndTimezones(ctx context.Context, client *http.Client, engine_host string) ([]string, string, error) {
	ntpURL := engine_host + ENGINE_APIS["NTP_CONFIG"]
	req, err := http.NewRequest(http.MethodGet, ntpURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating NTP request: "+err.Error())
		return nil, "", err
	}
	req.Header.Set("Content-Type", "application/json")
	tflog.Info(ctx, DLPX+INFO+"GET NTP Request URL: "+ntpURL)
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting NTP servers: "+err.Error())
		return nil, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading NTP response: "+err.Error())
		return nil, "", err
	}

	var ntpRes GetNTPResponse
	err = json.Unmarshal(body, &ntpRes)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshaling NTP response: "+err.Error())
		return nil, "", err
	}
	if ntpRes.Status != "OK" {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] NTP API returned non-OK status: "+ntpRes.Status)
		return nil, "", fmt.Errorf("["+engine_host+"] NTP API returned error status: %s", ntpRes.Status)
	}

	// Extract servers and timezone
	servers := ntpRes.Result.NtpConfig.Servers
	timezone := ntpRes.Result.SystemTimeZone
	return servers, timezone, nil
}

func getDNSConfiguration(ctx context.Context, client *http.Client, engine_host string) (GetDNSResponse, error) {
	{
		dnsURL := engine_host + ENGINE_APIS["DNS_CONFIG"]
		req, err := http.NewRequest(http.MethodGet, dnsURL, nil)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating DNS request: "+err.Error())
			return GetDNSResponse{}, err
		}
		req.Header.Set("Content-Type", "application/json")
		tflog.Info(ctx, DLPX+INFO+"GET DNS Request URL: "+dnsURL)
		resp, err := client.Do(req)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting DNS configuration: "+err.Error())
			return GetDNSResponse{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading DNS response: "+err.Error())
			return GetDNSResponse{}, err
		}

		var dnsRes GetDNSResponse
		err = json.Unmarshal(body, &dnsRes)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshaling DNS response: "+err.Error())
			return GetDNSResponse{}, err
		}
		if dnsRes.Status != "OK" {
			tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] DNS API returned non-OK status: "+dnsRes.Status)
			return GetDNSResponse{}, fmt.Errorf("["+engine_host+"] DNS API returned error status: %s", dnsRes.Status)
		}

		return dnsRes, nil
	}
}

func getEntityIDForSSO(ctx context.Context, client *http.Client, engine_host string) (string, error) {
	ssoURL := engine_host + ENGINE_APIS["SSO_CONFIG"]
	req, err := http.NewRequest(http.MethodGet, ssoURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating SSO request: "+err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	tflog.Info(ctx, DLPX+INFO+"GET SSO Request URL: "+ssoURL)
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting SSO details: "+err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading SSO response: "+err.Error())
		return "", err
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] SSO Response Body: "+string(body))
	var ssoRes map[string]interface{}
	err = json.Unmarshal(body, &ssoRes)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshaling SSO response: "+err.Error())
		return "", err
	}
	if bodyMap, ok := ssoRes["result"].(map[string]interface{}); ok {
		if entityId, exists := bodyMap["entityId"].(string); exists {
			return entityId, nil
		}
	}
	return "", fmt.Errorf("[%s] failed to retrieve EntityID for sso", engine_host)
}

func getComplianceUserID(ctx context.Context, client *http.Client, engine_host string, comp_user string, authToken string) (string, error) {
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Getting compliance user ID")
	userURL := engine_host + ENGINE_APIS["COMPLIANCE_USER"]
	req, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting compliance users request: "+err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error getting compliance users: "+err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading compliance users response: "+err.Error())
		return "", err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, `"status":"ERROR"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return "", fmt.Errorf("[%s] Failed to get compliance user ID: %s", engine_host, bodyStr)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshalling compliance users response: "+err.Error())
		return "", err
	}

	responseList, ok := response["responseList"].([]interface{})
	if !ok {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] responseList field not found or not an array in response")
		return "", fmt.Errorf("[%s] invalid response format: responseList field not found or not an array", engine_host)
	}
	for _, item := range responseList {
		if user, ok := item.(map[string]interface{}); ok {
			if userName, ok := user["userName"].(string); ok && userName == comp_user {
				if userID, ok := user["userId"].(float64); ok {
					userIDStr := strconv.Itoa(int(userID))
					tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Found compliance user - ID: "+userIDStr+", Name: "+userName+", Status: "+fmt.Sprintf("%v", user["userStatus"]))
					return userIDStr, nil
				} else {
					tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] userId field not found or not a number for user: "+comp_user)
					return "", fmt.Errorf("[%s] userId field not found or not a number for user: %s", engine_host, comp_user)
				}
			}
		}
	}
	tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] compliance user "+comp_user+" not found")
	return "", fmt.Errorf("[%s] compliance user %s not found", engine_host, comp_user)
}

func updateComplianceUser(ctx context.Context, client *http.Client, engine_host string, user_id string, comp_new_pass string, email string, comp_user string, authToken string) ([]byte, error) {
	updateURL := engine_host + ENGINE_APIS["COMPLIANCE_USER"] + "/" + user_id
	// Create compliance user update parameters
	UpdateParameters := map[string]interface{}{
		"email":      email,
		"password":   comp_new_pass,
		"userName":   comp_user,
		"userStatus": "ACTIVE",
		"isAdmin":    true,
		"firstName":  "",
		"lastName":   "",
	}
	UpdateParametersJSON, err := json.Marshal(UpdateParameters)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling compliance user update data: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, updateURL, bytes.NewReader(UpdateParametersJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating compliance user update request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error updating compliance user: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Compliance User Update Response Body: "+string(body))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading compliance user update response: "+err.Error())
		return nil, err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, `"errorMessage"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return nil, fmt.Errorf("["+engine_host+"] Failed to update compliance user: %s", bodyStr)
	}

	return body, nil
}

func loginComplianceUser(ctx context.Context, client *http.Client, engine_host string, comp_user string, comp_pass string) (string, error) {
	loginURL := engine_host + ENGINE_APIS["COMPLIANCE_LOGIN"]
	loginData := LoginRequest{
		Password: comp_pass,
		Username: comp_user,
	}
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Logging in compliance user "+comp_user)
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error marshalling compliance user login data: "+err.Error())
		return "", err
	}
	tflog.Debug(ctx, fmt.Sprintf("[SECURITY] Compliance login JSON payload created: %d bytes", len(loginJSON)))
	// Clear sensitive data from loginJSON after use
	defer func() {
		tflog.Debug(ctx, "[SECURITY] Clearing compliance login credentials from memory")
		SecureClearByteSlice(ctx, loginJSON)
		tflog.Debug(ctx, "[SECURITY] Compliance login credentials cleared")
	}()
	
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Compliance User Login ")
	req, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewReader(loginJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating compliance user login request: "+err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sending compliance user login request to "+loginURL)
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error authenticating compliance user: "+err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Compliance User Login Response Body: "+string(body))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error compliance user login response: "+err.Error())
		return "", err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, `"errorMessage"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return "", fmt.Errorf("["+engine_host+"] Failed to login as compliance user: %s", bodyStr)
	}

	var compResp map[string]interface{}
	if err := json.Unmarshal(body, &compResp); err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error unmarshalling compliance user login response: "+err.Error())
		return "", err
	}
	if auth, ok := compResp["Authorization"].(string); ok {
		return auth, nil
	}

	return "", fmt.Errorf("[%s] authorization token not found in compliance user login response", engine_host)
}

func startMasking(ctx context.Context, client *http.Client, engine_host string) diag.Diagnostics {
	var diags diag.Diagnostics
	startMaskingURL := engine_host + ENGINE_APIS["START_MASKING"]
	req, err := http.NewRequest(http.MethodPost, startMaskingURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error creating start masking request: "+err.Error())
		return diag.Errorf("[%s] error creating start masking request: %s", engine_host, err)
	}
	req.Header.Set("Content-Type", "application/json")
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Sending start masking request to "+startMaskingURL)
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error starting masking: "+err.Error())
		return diag.Errorf("["+engine_host+"] error starting masking: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	tflog.Info(ctx, DLPX+INFO+"["+engine_host+"] Start Masking Response Body: "+string(body))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] error reading start masking response: "+err.Error())
		return diag.Errorf("["+engine_host+"] error reading start masking response: %s", err)
	}

	bodyStr := string(body)
	if strings.Contains(bodyStr, `"status":"ERROR"`) {
		tflog.Error(ctx, DLPX+ERROR+"["+engine_host+"] API returned error response: "+bodyStr)
		return diag.Errorf("["+engine_host+"] Failed to start masking: %s", bodyStr)
	}

	return diags
}

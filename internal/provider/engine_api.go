package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// startSession starts a new session
func startSession(ctx context.Context, client *http.Client, engine_host string, version string) error {
	versionParts := strings.Split(version, ".")
	major, _ := strconv.Atoi(versionParts[0])
	minor, _ := strconv.Atoi(versionParts[1])
	micro, _ := strconv.Atoi(versionParts[2])
	tflog.Error(ctx, DLPX+INFO+"start session for "+engine_host+" version "+version)
	sessionURL := engine_host + "/resources/json/delphix/session"
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
		tflog.Error(ctx, DLPX+ERROR+"error marshalling session data: "+err.Error())
		return err
	}

	req, err := http.NewRequest(http.MethodPost, sessionURL, bytes.NewReader(sessionData))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating session request: "+err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error starting session: "+err.Error())
		return err
	}
	defer resp.Body.Close()

	// Check for successful session start (can be modified to return response body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		tflog.Error(ctx, DLPX+ERROR+"error starting session: "+strconv.Itoa(resp.StatusCode)+" "+string(body))
		return errors.New("error to start session: " + strconv.Itoa(resp.StatusCode) + " " + string(body))
	}
	return nil
}

func login(ctx context.Context, client *http.Client, engine_host string, user string, password string, target string) error {
	loginURL := engine_host + "/resources/json/delphix/login"
	loginData := LoginRequest{
		Password: password,
		Type:     "LoginRequest",
		Username: user,
		Target:   target,
	}
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error marshalling login data: "+err.Error())
		return err
	}

	req, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewReader(loginJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating login request: "+err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error authenticating: "+err.Error())
		return err
	}
	defer resp.Body.Close()

	// Check for successful login (can be modified to return response body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		tflog.Error(ctx, DLPX+ERROR+"Failed to login: "+strconv.Itoa(resp.StatusCode)+" "+string(body))
		return errors.New("Failed to login: " + strconv.Itoa(resp.StatusCode) + " " + string(body))
	}
	return nil
}

func getDevices(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	deviceURL := engine_host + "/resources/json/delphix/storage/device"
	req, err := http.NewRequest(http.MethodGet, deviceURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating device request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error getting device: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading devices response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func getCurrentUser(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	userURL := engine_host + "/resources/json/delphix/user/current"
	req, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating current user request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error getting current user: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading current user: "+err.Error())
		return nil, err
	}
	return body, nil
}

func initializeSystem(ctx context.Context, client *http.Client, engine_host string, resultList ListResult, user string, email string, password string) ([]byte, error) {
	initializeSystemURL := engine_host + "/resources/json/delphix/domain/initializeSystem"
	initializationParams := SystemInitializationParameters{
		Type:            "SystemInitializationParameters",
		DefaultUser:     user,
		DefaultPassword: password,
		Devices:         make([]string, 0),
	}
	if email != "" {
		initializationParams.DefaultEmail = email
	}

	for _, device := range resultList.Result {
		if !device.Configured {
			initializationParams.Devices = append(initializationParams.Devices, device.Reference)
		}
	}
	postData, err := json.Marshal(initializationParams)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error marshalling initialization parameters: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, initializeSystemURL, bytes.NewReader(postData))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating initialize system request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error initializing system: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading initializing system response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func getAction(ctx context.Context, client *http.Client, engine_host string, action_id string) ([]byte, error) {
	actionURL := engine_host + "/resources/json/delphix/action/" + action_id
	req, err := http.NewRequest(http.MethodGet, actionURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating action request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error getting action: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading action response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func getSystem(ctx context.Context, client *http.Client, engine_host string) ([]byte, error) {
	actionURL := engine_host + "/resources/json/delphix/system"
	req, err := http.NewRequest(http.MethodGet, actionURL, nil)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating system request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error getting system: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading system response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func updatePassword(ctx context.Context, client *http.Client, engine_host string, user_id string, password string) ([]byte, error) {
	updateURL := engine_host + "/resources/json/delphix/user/" + user_id + "/updateCredential"
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
		tflog.Error(ctx, DLPX+ERROR+"error marshalling login data: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, updateURL, bytes.NewReader(UpdateParametersJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating login request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error authenticating: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func createOrUpdateUser(ctx context.Context, client *http.Client, engine_host string, user_name string, password string, user_type string) ([]byte, error) {
	updateURL := engine_host + "/resources/json/delphix/user"
	// Create user parameters
	UserParameters := User{
		Name:                        user_name,
		UserType:                    user_type,
		AuthenticationType:          "NATIVE",
		Credential:                  Credential{Password: password, Type: "PasswordCredential"},
		AllowPasswordAuthentication: true,
		Type:                        "User",
	}

	UserParametersJSON, err := json.Marshal(UserParameters)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error marshalling user data: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, updateURL, bytes.NewReader(UserParametersJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating user request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error create/update user: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading response: "+err.Error())
		return nil, err
	}
	return body, nil
}

func setEngieType(ctx context.Context, client *http.Client, engine_host string, engine_type string) ([]byte, error) {
	updateURL := engine_host + "/resources/json/delphix/system"
	var eg_type string

	switch engine_type {
	case "CC":
		eg_type = "MASKING"
	case "CD":
		eg_type = "VIRTUALIZATION"
	default:
		tflog.Error(ctx, DLPX+ERROR+"Unknown engine type: "+engine_type)
	}
	// Prepare system information data
	data := SystemInfo{
		Type:       "SystemInfo",
		EngineType: eg_type,
	}
	systemInfoJSON, err := json.Marshal(data)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error marshalling login data: "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, updateURL, bytes.NewReader(systemInfoJSON))
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error creating request: "+err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error authenticating: "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"error reading response: "+err.Error())
		return nil, err
	}
	return body, nil
}

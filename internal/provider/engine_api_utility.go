package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func pollActionStatus(ctx context.Context, client *http.Client, engine_host string, action string) diag.Diagnostics {
	var diags diag.Diagnostics
	for {
		// Get Action
		tflog.Info(ctx, DLPX+INFO+" action_id "+action)
		actionData, err := getAction(ctx, client, engine_host, action)
		if err != nil {
			return diag.Errorf("Error getting action: %s", err)
		}
		var actionResult ActionResult
		err = json.Unmarshal(actionData, &actionResult)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+" Error unmarshalling "+err.Error())
		}
		tflog.Info(ctx, DLPX+INFO+" action state "+actionResult.Result.State)
		if actionResult.Result.State == "COMPLETED" {
			tflog.Info(ctx, DLPX+INFO+"Action completed!")
			break
		}
		time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
	}
	return diags
}

func UpdateUserPassword(ctx context.Context, client *http.Client, engine_host string, version string, user string, curr_pass string, new_pass string, target string) diag.Diagnostics {
	var diags diag.Diagnostics

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"start Session for "+engine_host)
	err := startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"login as "+user)
	err = login(ctx, client, engine_host, user, curr_pass, target)
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	// Get User Details
	userData, err := getCurrentUser(ctx, client, engine_host)
	if err != nil {
		return diag.Errorf("Error getting curr user: %s", err)
	}

	var result ActionResult
	err = json.Unmarshal(userData, &result)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Unmarshal Error "+err.Error())
	}
	tflog.Info(ctx, DLPX+INFO+"Current User Reference "+result.Result.Reference)

	// Update sysadmin Password
	updateResp, updateErr := updatePassword(ctx, client, engine_host, result.Result.Reference, new_pass)
	if updateErr != nil {
		return diag.Errorf("Error in update Password: %s", updateErr)
	}
	tflog.Info(ctx, DLPX+INFO+"update password "+result.Result.Reference+" "+string(updateResp))

	return diags
}

func initializeSystemAndDevices(ctx context.Context, client *http.Client, engine_host string, user string, default_email string, password string) (ActionResult, diag.Diagnostics) {
	var diags diag.Diagnostics
	// Get Devices
	deviceData, err := getDevices(ctx, client, engine_host)
	if err != nil {
		return ActionResult{}, diag.Errorf("Error getting devices: %s", err)
	}
	tflog.Info(ctx, DLPX+INFO+"devices "+string(deviceData))

	// Parse device information
	var resultList ListResult
	err = json.Unmarshal(deviceData, &resultList)
	if err != nil {
		return ActionResult{}, diag.Errorf("Error parsing device information: %s", err)
	}

	// Initialize System
	resp, err := initializeSystem(ctx, client, engine_host, resultList, user, default_email, password)
	if err != nil {
		return ActionResult{}, diag.Errorf("Error initializing system: %s", err)
	}
	tflog.Info(ctx, DLPX+INFO+"Initializing system: "+string(resp))

	var result ActionResult
	unmarshalErr := json.Unmarshal(resp, &result)
	if unmarshalErr != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error unmarshalling: "+unmarshalErr.Error())
	}
	return result, diags
}

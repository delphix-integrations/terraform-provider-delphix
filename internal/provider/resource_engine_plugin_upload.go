package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	dctapi "github.com/delphix/dct-sdk-go/v25"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabasePlugin() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Database plugin initialization.",

		CreateContext: databasePluginCreate,
		ReadContext:   databasePluginRead,
		UpdateContext: databasePluginUpdate,
		DeleteContext: databasePluginDelete,

		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					filePath := val.(string)
					// Check if file exists
					if _, err := os.Stat(filePath); os.IsNotExist(err) {
						errs = append(errs, fmt.Errorf("%q file does not exist: %s", key, filePath))
						return
					}

					// Check if file has .json extension
					if !strings.HasSuffix(strings.ToLower(filePath), ".json") {
						errs = append(errs, fmt.Errorf("%q must be a JSON file (with .json extension), got: %s", key, filePath))
						return
					}

					// Optional: Validate if it's a valid JSON file by trying to parse it
					file, err := os.Open(filePath)
					if err != nil {
						errs = append(errs, fmt.Errorf("%q cannot be opened: %v", key, err))
						return
					}
					defer file.Close()

					// Read and validate JSON content
					content, err := io.ReadAll(file)
					if err != nil {
						errs = append(errs, fmt.Errorf("%q cannot be read: %v", key, err))
						return
					}

					var jsonData interface{}
					if err := json.Unmarshal(content, &jsonData); err != nil {
						errs = append(errs, fmt.Errorf("%q is not a valid JSON file: %v", key, err))
						return
					}

					return
				},
			},
			"engine_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"toolkit_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func databasePluginCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Create a cookie jar to store session cookies
	client := meta.(*apiClient).client
	file_path := d.Get("file_path").(string)
	engine_name := d.Get("engine_host").(string)

	hostnameRegex := regexp.MustCompile(`^https?://(.+?)(?:[:/]|$)`)
	matches := hostnameRegex.FindStringSubmatch(engine_name)
	if len(matches) < 2 {
		return diag.Errorf("Invalid engine_name format. Expected format: http(s)://hostname, got: %s", engine_name)
	}

	hostname := matches[1]
	tflog.Info(ctx, DLPX+INFO+"Extracted hostname: "+hostname)

	filter_expression := fmt.Sprintf("hostname EQ '%s'", hostname)
	searchBody := *dctapi.NewSearchBody()
	searchBody.SetFilterExpression(filter_expression)
	_, httpRes, err := client.ManagementAPI.SearchEngines(ctx).SearchBody(searchBody).Execute()

	if err != nil {
		return diag.Errorf("Error fetching registered engines: %v", err)
	}
	tflog.Info(ctx, DLPX+INFO+"Search response: "+fmt.Sprintf("%v", httpRes.Body))
	engineResJson, err := io.ReadAll(httpRes.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error occurred in reading body of the response.")
		return diag.Errorf("Error reading response body: %v", err)
	}

	var engine_id string
	var engineData map[string]interface{}
	if err := json.Unmarshal(engineResJson, &engineData); err != nil {
		return diag.Errorf("Error parsing engine JSON: %v", err)
	}
	if _, ok := engineData["items"].([]interface{}); ok {
		items := engineData["items"].([]interface{})
		if len(items) == 0 {
			return diag.Errorf("No registered engine found with hostname: %s", hostname)
		}
		firstEngine := items[0].(map[string]interface{})
		engine_id = firstEngine["id"].(string)
		tflog.Info(ctx, DLPX+INFO+"Found engine ID: "+engine_id)
	}
	tflog.Info(ctx, DLPX+INFO+"Uploading database plugin from file: "+file_path)
	file, err := os.Open(file_path)
	if err != nil {
		return diag.Errorf("Error opening plugin file: %v", err)
	}
	defer file.Close()
	// Read the JSON content
	jsonContent, err := io.ReadAll(file)
	if err != nil {
		return diag.Errorf("Error reading plugin file: %v", err)
	}

	// Parse JSON to extract the name field
	var pluginData map[string]interface{}
	if err := json.Unmarshal(jsonContent, &pluginData); err != nil {
		return diag.Errorf("Error parsing plugin JSON: %v", err)
	}

	// Extract the name field
	var pluginName string
	if name, exists := pluginData["name"].(string); exists {
		pluginName = name
		tflog.Info(ctx, DLPX+INFO+"Plugin name from JSON: "+pluginName)
	} else {
		return diag.Errorf("Plugin JSON file does not contain a 'name' field")
	}

	file, err = os.Open(file_path)
	if err != nil {
		return diag.Errorf("Error reopening plugin file for upload: %v", err)
	}
	defer file.Close()
	apiReq := client.ToolkitsAPI.UploadToolkit(ctx).EngineId(engine_id).File(file)
	_, httpRes, apiErr := apiReq.Execute()
	if apiErr != nil {
		return diag.Errorf("Error uploading plugin to engine: %v", apiErr)
	}

	if httpRes != nil && httpRes.Body != nil {
		bodyBytes, err := io.ReadAll(httpRes.Body)
		if err != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error reading response body: "+err.Error())
			return diag.Errorf("Error reading response body: %v", err)
		}
		defer httpRes.Body.Close()

		tflog.Info(ctx, DLPX+INFO+"Upload response status: "+fmt.Sprintf("%s", bodyBytes))

		// Parse JSON response
		var jobResponse map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &jobResponse); err != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error parsing JSON response: "+err.Error())
			return diag.Errorf("Error parsing JSON response: %v", err)
		}

		// Extract job information
		if job, ok := jobResponse["job"].(map[string]interface{}); ok {
			if jobId, exists := job["id"].(string); exists {
				tflog.Info(ctx, DLPX+INFO+"Upload job ID: "+jobId)

				// Poll for job completion
				jobStatus, pollErr := PollJobStatus(jobId, ctx, client)
				if pollErr != "" {
					return diag.Errorf("Error polling job status: %v", pollErr)
				}

				tflog.Info(ctx, DLPX+INFO+"Job result is "+jobStatus)
			}

			if targetName, exists := job["target_name"].(string); exists {
				tflog.Info(ctx, DLPX+INFO+"Uploaded toolkit: "+targetName)
				d.Set("toolkit_name", pluginName)
			}

		}
	}

	readDiags := databasePluginRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

func databasePluginUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return diag.Errorf("Action update not available for engine plugin")
}

func databasePluginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, DLPX+INFO+"Reading engine configuration")
	toolkit_name := d.Get("toolkit_name").(string)
	tflog.Info(ctx, DLPX+INFO+"toolkit name: "+toolkit_name)
	client := meta.(*apiClient).client
	toolkitReq := client.ToolkitsAPI.GetToolkits(ctx)
	_, res, er := toolkitReq.Execute()
	if er != nil {
		return diag.Errorf("Error fetching existing toolkits %v", er)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error reading response body: "+err.Error())
		return diag.Errorf("Error reading response body: %v", err)
	}
	defer res.Body.Close()

	// Parse JSON response
	var toolkits map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &toolkits); err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error parsing JSON response: "+err.Error())
		return diag.Errorf("Error parsing JSON response for toolkits: %v", err)
	}

	tflog.Info(ctx, DLPX+INFO+"Toolkits fetched: "+fmt.Sprintf("%v", toolkits))
	if items, ok := toolkits["items"].([]interface{}); ok {
		tflog.Info(ctx, DLPX+INFO+"Number of toolkits found: "+fmt.Sprintf("%d", len(items)))
		for _, item := range items {
			if toolkit, ok := item.(map[string]interface{}); ok {
				if displayName, exists := toolkit["display_name"].(string); exists {
					tflog.Info(ctx, DLPX+INFO+"Found toolkit with display name: "+displayName)
					if displayName == toolkit_name {
						// Set the resource ID using the toolkit's id
						tflog.Info(ctx, DLPX+INFO+"Matching toolkit found: "+displayName)
						if id, exists := toolkit["id"].(string); exists {
							d.SetId(id)
							tflog.Info(ctx, DLPX+INFO+"Found toolkit with ID: "+id)
							return diags
						}
					}
				}
			}
		}
	}
	return diags
}

func databasePluginDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	toolkitReq := client.ToolkitsAPI.GetToolkits(ctx)
	_, _, er := toolkitReq.Execute()
	if er != nil {
		return diag.Errorf("Error fetching existing toolkits %v", er)
	}

	tflog.Info(ctx, DLPX+INFO+"Deleting toolkit with ID: "+d.Get("id").(string))
	delReq := client.ToolkitsAPI.DeleteToolkitById(ctx, d.Get("id").(string))
	apiRes, httpRes, err := delReq.Execute()
	tflog.Info(ctx, DLPX+INFO+"Delete response Body: "+fmt.Sprintf("%v", httpRes.Body))
	tflog.Info(ctx, DLPX+INFO+"Delete API response: "+fmt.Sprintf("%v", apiRes))
	if err != nil {
		return diag.Errorf("Error in deleting the toolkit %v, %v", err, httpRes.Body)
	}

	return diags
}

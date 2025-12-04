package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v25"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVdbGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing VDB Groups.",

		CreateContext: resourceVdbGroupCreate,
		ReadContext:   resourceVdbGroupRead,
		UpdateContext: resourceVdbGroupUpdate,
		DeleteContext: resourceVdbGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: CustomizeDiffTags,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vdb_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_tag_changes": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				DiffSuppressFunc: func(_, old, new string, d *schema.ResourceData) bool {
					if ignore, ok := d.GetOk("ignore_tag_changes"); ok && ignore.(bool) {
						return true
					}
					return false
				},
			},
		},
	}
}

func resourceVdbGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*apiClient).client

	vdbGroupCreateReq := *dctapi.NewCreateVDBGroupRequest(d.Get("name").(string))
	vdbGroupCreateReq.SetVdbIds(toStringArray(d.Get("vdb_ids")))

	if v, has_v := d.GetOk("tags"); has_v {
		vdbGroupCreateReq.SetTags(toTagArray(v))
	}

	apiRes, httpRes, err := client.VDBGroupsAPI.CreateVdbGroup(ctx).CreateVDBGroupRequest(vdbGroupCreateReq).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(apiRes.VdbGroup.GetId())

	readDiags := resourceVdbGroupRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

func resourceVdbGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbGroupId := d.Id()
	tflog.Info(ctx, DLPX+INFO+"VdbGroupId: "+vdbGroupId)
	apiRes, httpRes, err := client.VDBGroupsAPI.GetVdbGroup(ctx, vdbGroupId).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		d.SetId("")
		return diags
	}

	d.Set("name", apiRes.GetName())
	d.Set("vdb_ids", apiRes.GetVdbIds())
	tflog.Debug(ctx, "Getting Raw Config")
	errors := HandleRawConfigReadContext(ctx, d, apiRes)
	if errors != nil {
		return diag.FromErr(errors)
	}

	return diags
}

func resourceVdbGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Updating VDB Group")
	client := meta.(*apiClient).client
	var name_or_vdb_ids_changed bool = false
	vdbGroupId := d.Id()

	// Existing update logic
	updateVdbGroupReq := *dctapi.NewUpdateVDBGroupParameters()
	if d.HasChange("name") {
		updateVdbGroupReq.SetName(d.Get("name").(string))
		name_or_vdb_ids_changed = true
	}
	if d.HasChange("vdb_ids") {
		updateVdbGroupReq.SetVdbIds(toStringArray(d.Get("vdb_ids")))
		name_or_vdb_ids_changed = true
	}

	_, httpRes, err := client.VDBGroupsAPI.UpdateVdbGroupById(ctx, vdbGroupId).UpdateVDBGroupParameters(updateVdbGroupReq).Execute()
	if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
		return diags
	}

	// Polling logic for name/vdb_ids update
	if name_or_vdb_ids_changed {
		maxAttempts := 10
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			status, err := PollVdbGroupStatus(ctx, client, vdbGroupId)
			if err != nil {
				return diag.FromErr(err)
			}
			if status == "RUNNING" {
				break
			}
			if status == "FAILED" || status == "CANCELED" {
				return diag.Errorf("VDB group update failed with status: %s", status)
			}
			time.Sleep(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second)
		}
	}

	// Tag update logic
	tflog.Info(ctx, DLPX+INFO+"Updating tags for VDB Group: ")
	if !d.Get("ignore_tag_changes").(bool) {
		tflog.Info(ctx, DLPX+INFO+"Reading from Api to get existing tags")
		apiRes, httpRes, err := client.VDBGroupsAPI.GetVdbGroup(ctx, vdbGroupId).Execute()
		if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
			d.SetId("")
			return diags
		}
		tags := flattenTags(apiRes.GetTags())
		tflog.Debug(ctx, "Existing tags", map[string]interface{}{
			"tags": tags,
		})
		newRaw := d.GetRawConfig()
		if newRaw.IsKnown() || !newRaw.IsNull() {
			attr := newRaw.GetAttr("tags")
			tflog.Debug(ctx, "New tags raw config value", map[string]interface{}{
				"tags": attr.GoString(),
			})
			d.Set("tags", flattenTags(apiRes.GetTags()))
			if attr.IsNull() || !attr.IsKnown() || attr.LengthInt() == 0 {
				// This now correctly gives [] if the user set tags = []
				if len(tags) != 0 {
					tflog.Info(ctx, DLPX+INFO+"Tags field is not set, deleting all existing tags")
					httpRes, err := client.VDBGroupsAPI.DeleteVdbGroupTags(ctx, vdbGroupId).Execute()
					if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
						return diags
					}
				}
				return resourceVdbGroupRead(ctx, d, meta)
			}
		}

		if d.HasChange("tags") {
			oldTags, newTags := d.GetChange("tags")
			oldTagList := oldTags.([]interface{})
			newTagList := newTags.([]interface{})

			// Create a map of old tags for easy lookup
			oldTagMap := make(map[string]map[string]bool)
			for _, tag := range oldTagList {
				tagMap := tag.(map[string]interface{})
				key := tagMap["key"].(string)
				value := tagMap["value"].(string)
				if _, exists := oldTagMap[key]; !exists {
					oldTagMap[key] = make(map[string]bool)
				}
				oldTagMap[key][value] = true
			}

			// Create a map of new tags for easy lookup
			newTagMap := make(map[string]map[string]bool)
			for _, tag := range newTagList {
				tagMap := tag.(map[string]interface{})
				key := tagMap["key"].(string)
				value := tagMap["value"].(string)
				if _, exists := newTagMap[key]; !exists {
					newTagMap[key] = make(map[string]bool)
				}
				newTagMap[key][value] = true
			}

			// If newTagList is empty, delete all existing tags
			if len(newTagList) == 0 {
				for key := range oldTagMap {
					deleteTag := *dctapi.NewDeleteTag()
					deleteTag.SetKey(key)
					httpRes, err := client.VDBGroupsAPI.DeleteVdbGroupTags(ctx, vdbGroupId).DeleteTag(deleteTag).Execute()
					if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
						return diags
					}
				}
			} else {
				// Delete removed tags
				for key, oldValues := range oldTagMap {
					newValues, exists := newTagMap[key]
					if !exists {
						// Key doesn't exist in new tags, delete all values for this key
						deleteTag := *dctapi.NewDeleteTag()
						deleteTag.SetKey(key)
						httpRes, err := client.VDBGroupsAPI.DeleteVdbGroupTags(ctx, vdbGroupId).DeleteTag(deleteTag).Execute()
						if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
							return diags
						}
					} else {
						// Key exists, delete only values that are not in new tags
						for oldValue := range oldValues {
							if !newValues[oldValue] {
								deleteTag := *dctapi.NewDeleteTag()
								deleteTag.SetKey(key)
								deleteTag.SetValue(oldValue)
								httpRes, err := client.VDBGroupsAPI.DeleteVdbGroupTags(ctx, vdbGroupId).DeleteTag(deleteTag).Execute()
								if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
									return diags
								}
							}
						}
					}
				}

				// Create new tags
				var tags []dctapi.Tag
				for key, newValues := range newTagMap {
					oldValues, exists := oldTagMap[key]
					if !exists {
						// Key doesn't exist in old tags, create all values
						for value := range newValues {
							tag := *dctapi.NewTag(key, value)
							tags = append(tags, tag)
						}
					} else {
						// Key exists, create only new values
						for value := range newValues {
							if !oldValues[value] {
								tag := *dctapi.NewTag(key, value)
								tags = append(tags, tag)
							}
						}
					}
				}
				if len(tags) > 0 {
					tagsRequest := *dctapi.NewTagsRequest(tags)
					_, httpRes, err := client.VDBGroupsAPI.CreateVdbGroupsTags(ctx, vdbGroupId).TagsRequest(tagsRequest).Execute()
					if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
						return diags
					}
				}
			}
		}
	}

	return resourceVdbGroupRead(ctx, d, meta)
}

func PollVdbGroupStatus(ctx context.Context, client *dctapi.APIClient, vdbGroupId string) (string, error) {
	res, _, err := client.VDBGroupsAPI.GetVdbGroup(ctx, vdbGroupId).Execute()
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}

func resourceVdbGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbGroupId := d.Id()

	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	httpRes, err := client.VDBGroupsAPI.DeleteVdbGroup(ctx, vdbGroupId).Execute()

	if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
		return diags
	}
	if err != nil {
		resBody, err := ResponseBodyToString(ctx, httpRes.Body)
		if err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf(resBody)
	}

	return diags
}

package provider

import (
	"context"

	dctapi "github.com/delphix/dct-sdk-go"
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
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceVdbGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	apiRes, httpRes, err := client.VDBGroupsApi.CreateVdbGroup(ctx).CreateVDBGroupRequest(*dctapi.NewCreateVDBGroupRequest(
		d.Get("name").(string),
		toStringArray(d.Get("vdb_ids")),
	)).Execute()

	if diags := apiErrorResponseHelper(httpRes, err); diags != nil {
		return diags
	}

	d.SetId(apiRes.VdbGroup.GetId())

	return resourceVdbGroupRead(ctx, d, meta)
}

func resourceVdbGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	vdbGroupId := d.Id()
	InfoLog.Printf("VdbGroupId: %s", vdbGroupId)
	apiRes, httpRes, err := client.VDBGroupsApi.GetVdbGroup(ctx, vdbGroupId).Execute()

	if diags := apiErrorResponseHelper(httpRes, err); diags != nil {
		return diags
	}

	d.Set("name", apiRes.GetName())
	d.Set("vdb_ids", apiRes.GetVdbIds())
	return nil
}

func resourceVdbGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return diag.Errorf("not implemented")
}

func resourceVdbGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	vdbGroupId := d.Id()

	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	httpRes, err := client.VDBGroupsApi.DeleteVdbGroup(ctx, vdbGroupId).Execute()

	return apiErrorResponseHelper(httpRes, err)
}

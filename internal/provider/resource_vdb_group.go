package provider

import (
	"context"

	dctapi "github.com/delphix/dct-sdk-go/v14"
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

	var diags diag.Diagnostics

	client := meta.(*apiClient).client

	apiRes, httpRes, err := client.VDBGroupsApi.CreateVdbGroup(ctx).CreateVDBGroupRequest(*dctapi.NewCreateVDBGroupRequest(
		d.Get("name").(string),
		toStringArray(d.Get("vdb_ids")),
	)).Execute()

	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
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
	InfoLog.Printf("VdbGroupId: %s", vdbGroupId)
	apiRes, httpRes, err := client.VDBGroupsApi.GetVdbGroup(ctx, vdbGroupId).Execute()

	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.Set("name", apiRes.GetName())
	d.Set("vdb_ids", apiRes.GetVdbIds())
	return diags
}

func resourceVdbGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return diag.Errorf("not implemented")
}

func resourceVdbGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbGroupId := d.Id()

	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	httpRes, err := client.VDBGroupsApi.DeleteVdbGroup(ctx, vdbGroupId).Execute()

	if diags := apiErrorResponseHelper(nil, httpRes, err); diags != nil {
		return diags
	}
	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf(resBody)
	}

	return diags
}

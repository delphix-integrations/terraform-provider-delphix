package provider

import (
	"context"
	"log"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"

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

	res, httpRes, err := client.VDBGroupsApi.CreateVdbGroup(ctx).CreateVDBGroupRequest(*openapi.NewCreateVDBGroupRequest(
		d.Get("name").(string),
		toStringArray(d.Get("vdb_ids")),
	)).Execute()

	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			log.Fatal(err)
			return diag.FromErr(err)
		}
		return diag.Errorf(resBody)
	}

	d.SetId(res.VdbGroup.GetId())
	resourceVdbGroupRead(ctx, d, meta)
	return diags
}

func resourceVdbGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbGroupId := d.Id()
	log.Printf("VdbGroupId: %s", vdbGroupId)
	res, httpRes, err := client.VDBGroupsApi.GetVdbGroup(ctx, vdbGroupId).Execute()

	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf(resBody)
	}

	d.Set("name", res.GetName())
	d.Set("vdb_ids", res.GetVdbIds())
	return diags
}

func resourceVdbGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return diag.Errorf("not implemented")
}

func resourceVdbGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbGroupId := d.Id()

	deleteVdbParams := openapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	httpRes, err := client.VDBGroupsApi.DeleteVdbGroup(ctx, vdbGroupId).Execute()

	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf(resBody)
	}

	return diags
}

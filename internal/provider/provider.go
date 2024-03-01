package provider

import (
	"context"
	"crypto/tls"
	"net/http"

	dctapi "github.com/delphix/dct-sdk-go/v14"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_KEY", nil),
				},
				"tls_insecure_skip": {
					Type:        schema.TypeBool,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_TLS_INSECURE_SKIP", false),
				},
				"host": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_HOST", nil),
				},
				"host_scheme": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_HOST_SCHEME", "https"),
				},
				"debug": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"delphix_vdb":             resourceVdb(),
				"delphix_vdb_group":       resourceVdbGroup(),
				"delphix_environment":     resourceEnvironment(),
				"delphix_appdata_dsource": resourceAppdataDsource(),
				"delphix_oracle_dsource":  resourceOracleDsource(),
				"delphix_source":          resourceSource(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	client *dctapi.APIClient
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// configure client
		cfg := dctapi.NewConfiguration()
		cfg.Host = d.Get("host").(string)
		cfg.UserAgent = p.UserAgent("terraform-provider-delphix", version)
		cfg.Scheme = d.Get("host_scheme").(string)
		cfg.HTTPClient = &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: d.Get("tls_insecure_skip").(bool)},
		}}
		cfg.AddDefaultHeader("Authorization", "apk "+d.Get("key").(string))
		cfg.AddDefaultHeader("x-dct-client-name", "Terraform")

		client := dctapi.NewAPIClient(cfg)

		if d.Get("debug").(bool) {
			// print out raw api request body for debug purposes
			client.GetConfig().Debug = true
		}
		// make a test call

		req := client.ManagementApi.GetRegisteredEngines(ctx)
		_, _, err := client.ManagementApi.GetRegisteredEnginesExecute(req)

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return &apiClient{client}, nil
	}
}

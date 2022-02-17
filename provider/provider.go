package provider

import (
	"context"
	"crypto/tls"
	"net/http"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
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

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_KEY", nil),
				},
				"key_prefix": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("DCT_KEY_PREFIX", "apk"),
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
			},
			ResourcesMap: map[string]*schema.Resource{
				"delphix_resource": resourceScaffolding(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	client *openapi.APIClient
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// configure client
		cfg := openapi.NewConfiguration()
		cfg.Host = d.Get("host").(string)
		cfg.UserAgent = p.UserAgent("terraform-provider-delphix", version)
		cfg.Scheme = d.Get("host_scheme").(string)
		cfg.HTTPClient = &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}

		client := openapi.NewAPIClient(cfg)

		// make a test call
		apiKeyMap := make(map[string]openapi.APIKey)
		apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
			Key:    d.Get("key").(string),
			Prefix: d.Get("key_prefix").(string),
		}
		ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

		req := client.EnginesApi.GetEngines(ctx)
		_, _, err := client.EnginesApi.GetEnginesExecute(req)

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return &apiClient{client}, nil
	}
}

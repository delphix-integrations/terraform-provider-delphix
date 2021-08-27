module terraform-provider-delphix

go 1.16

replace github.com/ajaytho/delphix-go-sdk => ../delphix-go-sdk

require (
	github.com/ajaytho/delphix-go-sdk v0.0.0-20210825160814-bcc77c8ab45a // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.17.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0 // indirect
)

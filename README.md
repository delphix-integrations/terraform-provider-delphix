# Terraform Provider Delphix

![CodeQL](https://github.com/delphix-integrations/terraform-provider-delphix/actions/workflows/codeql.yml/badge.svg?branch=main)
![Release](https://github.com/delphix-integrations/terraform-provider-delphix/actions/workflows/release.yml/badge.svg?event=release)
![Version](https://img.shields.io/github/v/release/delphix-integrations/terraform-provider-delphix)

Terraform Provider for Delphix enables Terraform to create and manage Delphix Continuous Data &
Continuous Compliance infrastructure using the Delphix Control Tower (DCT) APIs.

Full implementation directions can be found on the [Delphix Ecosystem Documentation](https://help.delphix.com/eh/current/Content/Ecoystem/Terraform.htm) and [Terraform Provider Registry](https://registry.terraform.io/providers/delphix-integrations/delphix/latest/docs).

## Prerequisites

1. Install Delphix Control Tower (DCT). For more information, visit the [DCT documentation](https://docs.delphix.com/dct).
2. Delphix Continuous Data and Continuous Compliance engines must be connected to DCT.
3. An API key must be created for authenticating with DCT APIs. Refer to the [DCT API keys documentation](https://dct.delphix.com/docs/api-keys-2) for more info.
4. Additional infrastructure required for testing the provider operations. [e.g Hosts to be added as environments, dSources to create VDBs from]
5. Development setup for Golang.


## Getting Started (Development)

This guide covers the following

1. Install IDE [Visual Studio Code](https://code.visualstudio.com).

2. Install guide for [Golang](https://go.dev/dl/).

3. Install guide for [GoReleaser](https://goreleaser.com/install/).

4. Install Go Plugin for VS Code.

5. Install [Terraform](https://www.terraform.io/downloads).

6. Fork this repo and clone it locally. Switch to the `develop` branch which always heads to the latest development code.

7. Run following command to create binaries:

   ```goreleaser release --skip-publish --snapshot --rm-dist```

8. Execute example main.tf file under /examples/<resource> directory using the following commands:

    ``` 
        terraform init
        terraform plan
        terraform apply
    ```

## Contributing
This project is currently not accepting external contributions. 

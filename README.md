# Terraform Provider Delphix

![CodeQL](https://github.com/delphix-integrations/terraform-provider-delphix/actions/workflows/codeql.yml/badge.svg?branch=main)
![Release](https://github.com/delphix-integrations/terraform-provider-delphix/actions/workflows/release.yml/badge.svg?event=release)
![Version](https://img.shields.io/github/v/release/delphix-integrations/terraform-provider-delphix)

Terraform Provider for Delphix enables Terraform to create and manage Delphix Continuous Data &
Continuous Compliance infrastructure using the Delphix Control Tower (DCT) APIs.

Full documentation can he found [here](https://integrations.delphix.com/Terraform/)

## Prerequisites

1. Setup Delphix Control Tower(DCT) API Layer. For more information, visit [DCT Home.](https://docs.delphix.com/dct)
2. Delphix Engines must be registered with DCT APIs
3. API-Keys must be created for authenticating with DCT APIs. Refer to [DCT Home](https://docs.delphix.com/dct) for more info.
4. Additional infrastructure required for testing the provider operations [ e.g Hosts to be added as environments, dSources to create VDBs from]
5. Development setup for GoLang.


## Getting Started (Development)

This guide covers the following

1. Install IDE [Visual Studio Code](https://code.visualstudio.com)

2. Install guide for [golang](https://go.dev/dl/)

3. Install guide for [Goreleaser](https://goreleaser.com/install/)

4. Install Go Plugin for VS Code

5. Install [Terraform](https://www.terraform.io/downloads)

6. Fork this repo and clone it locally. Switch to develop branch which always heads to the latest development code.

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

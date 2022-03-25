# Terraform Provider Delphix

Terraform Provider for Delphix enables Terraform to create and manage Delphix Continuous Data & 
Continuous Compliance infrastructure using the Delphix Control Tower (DCT) APIs.

Full documentation can he found [here]()

## Getting Started (Development)
This guide will eventually cover the following


1. Setup DCT APi Gateway On Premise by following the [DCT API Gateway setup](https://github.com/delphix/orbital-api-gateway)

2. Install IDE [Visual Studio Code](https://code.visualstudio.com)

3. Install guide for [golang](https://go.dev/dl/)

4. Install guide for [Goreleaser](https://goreleaser.com/install/)

5. Install Go Plugin for VS Code

6. Install [Terraform](https://www.terraform.io/downloads)

7. Fork this repo and clone it locally. Switch to develop branch which always heads to the latest development code.

8. Run following command to create binaries:
    
    ```goreleaser release --skip-publish --snapshot --rm-dist```

9. Execute example main.tf file under /examples/<resource> directory using the following commands:

    ``` 
        terraform init
        terraform plan
        terraform apply
    ```

## Prerequisites

- Delphix Control Tower (DCT) API Gateway must be installed on-premise.
- Delphix Engines must be registered with DCT-OnPrem.
- API-Keys must be created for authenticating with DCT-OnPrem. Refer to DCT guide for more info.
- Additional infrastructure required for testing the provider operations [ e.g Hosts to be added as environments, dSources to create VDBs from]
- Development setup for GoLang.

## Contributing
This project is currently not accepting external contributions. 
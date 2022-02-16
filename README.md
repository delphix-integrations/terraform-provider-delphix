# POC repository for Delphix Terraform Provider

This is a project to build a Delphix Terraform Provider using the DCT-OnPrem APIs.

### Getting Started (Development)
This guide will eventually cover the following

    1. Setup DCT-OnPrem by following: https://github.com/delphix/orbital-api-gateway
    2. Install IDE: Visual Studio Code https://code.visualstudio.com
    3. Install "Go" language https://go.dev/dl/
    4. Install "Goreleaser" package https://goreleaser.com/install/
    5. Install "Go" Plugin for Visual Studio Code
    6. Install Terraform: https://www.terraform.io/downloads
    7. Fork this repo and clone it locally. Switch to develop branch which always heads to the latest development code.
    8. Run following command to create binaries:
        goreleaser release --skip-publish --snapshot --rm-dist
    9. Run example TF file in /examples/1/ by running following:
        terraform init
        terraform plan
        terraform apply

### Guides
    - For instructions on how to set up DCT-OnPrem, visit https://github.com/delphix/orbital-api-gateway

### Prerequisites

    - DCT-OnPrem must be installed and configured.
    - Delphix Engines must be registered with DCT-OnPrem
    - API-Keys for authenticatin with DCT-OnPrem
    - Additional infrastructure required for testing the provider operations [ e.g Environments, Source DBs]
    - IDE configured for GoLang Development.

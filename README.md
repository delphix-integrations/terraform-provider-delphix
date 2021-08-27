# terraform-provider-delphix
Prototyping for a Delphix provider for terraform

The following is a link to a video of the prototype in action: 
https://vimeo.com/263386228

This prototype is narrowly focused to support the following use cases:

## Use Case 1 - Adding/Removing a Unix Host Environment running Oracle
The prototype supports Create, Read (name only), and Delete

## Use Case 2 - Ingesting/Destroying an Oracle Single Instance dSource
The prototype supports Create, Read and Delete actions on that dSource.

## Use Case 3 - Provisioning/Destroying an Oracle Single Instance VDB
The prototype supports Create, Read, Update and Delete actions on that VDB.

## Use Case 4 - Provisioning/Destroying Delphix Dataset Groups
The prototype supports Create, Read, Update and Delete actions on that Group.

## Using the plugin
The following assumes you have a working knowledge of terraform and go

1. Clone this repo into your GOPATH
2. cd into the repo directory, and run ```go build```
3. Copy the terraform-provider-delphix file into your terraform directory
4. ```terraform init```
5. Now the plugin is loaded and ready for you to use. Create a simple blueprint to use with this plugin

## Example blueprint
```terraform apply``` will add a new Unix Host to the DE and then provision a VDB.
```terraform destroy``` will first delete the VDB, then delete the Host Environment.
Look to the main.tf.example for an example blueprint.

## Working locally with this repository
The following assumes you have a working knowledge of Golang

To get started, you will need to install Golang.

1. (a) If you use have a mac, and use homebrew, which I highly recommend, then it looks something like this:

```bash
	$ brew install golang
```

1. (b) If you do not have Golang, follow the instructions for installing
<a href="https://golang.org/dl">golang</a>.

2. If you do not have one already, create a directory to work with go.
3. Navigate into that directory in your terminal.
4. Create a directory called goprojects
5. Export the current working directory as GOPATH
6. Gather the required packages
7. Build
```bash
mkdir goprojects
export GOPATH=`pwd`
git clone https://github.com/delphix/terraform-provider-delphix goprojects/terraform-provider-delphix
cd goprojects/terraform-provider-delphix
go get
go build
```
## Known Issues
<a href="https://jira.delphix.com/browse/DLPX-57413" target="_blank">DLPX-57413</a>
Some sort of race condition with simultaneous API calls that can cause one of multiple same-type objects (i.e. environments) to fail to create.

## Contributing

All contributors are required to sign the Delphix Contributor Agreement prior to contributing code to an open source
repository. This process is handled automatically by [cla-assistant](https://cla-assistant.io/). Simply open a pull
request and a bot will automatically check to see if you have signed the latest agreement. If not, you will be prompted
to do so as part of the pull request process.

This project operates under the [Delphix Code of Conduct](https://delphix.github.io/code-of-conduct.html). By
participating in this project you agree to abide by its terms.

## Statement of Support

This software is provided as-is, without warranty of any kind or commercial support through Delphix. See the associated
license for additional details. Questions, issues, feature requests, and contributions should be directed to the
community as outlined in the [Delphix Community Guidelines](https://delphix.github.io/community-guidelines.html).

## License

This is code is licensed under the Apache License 2.0. Full license is available [here](./LICENSE).

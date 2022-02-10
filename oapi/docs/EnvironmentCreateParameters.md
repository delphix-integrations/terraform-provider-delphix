# EnvironmentCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the environment. | [optional] 
**EngineId** | **int64** | The ID of the Engine onto which to create the environment. | 
**OsName** | **string** | Operating system type of the environment. | 
**IsCluster** | Pointer to **bool** | Whether the environment to be created is a cluster. | [optional] [default to false]
**ClusterHome** | Pointer to **string** | Absolute path to cluster home drectory. This parameter is mandatory for UNIX cluster environments. | [optional] 
**Hostname** | **string** | host address of the machine. | 
**StagingEnvironment** | Pointer to **string** | Id of the connector environment which is used to connect to this source environment. This is mandatory parameter when creating Windows source environments. | [optional] 
**ConnectorPort** | Pointer to **int32** | Specify port on which Delphix connector will run. This is mandatory parameter when creating Windows target environments. | [optional] 
**SshPort** | Pointer to **int64** | ssh port of the host. | [optional] [default to 22]
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
**Username** | **string** | Username of the OS. | 
**Password** | **string** | Password of the OS. | 
**NfsAddresses** | Pointer to **[]string** | array of ip address or hostnames | [optional] 
**AseDbUsername** | Pointer to **string** | username of the SAP ASE database. | [optional] 
**AseDbPassword** | Pointer to **string** | password of the SAP ASE database. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**DspKeystorePath** | Pointer to **string** | DSP keystore path. | [optional] 
**DspKeystorePassword** | Pointer to **string** | DSP keystore password. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | DSP keystore alias. | [optional] 
**DspTruststorePath** | Pointer to **string** | DSP truststore path. | [optional] 
**DspTruststorePassword** | Pointer to **string** | DSP truststore password. | [optional] 
**Description** | Pointer to **string** | The environment description. | [optional] 

## Methods

### NewEnvironmentCreateParameters

`func NewEnvironmentCreateParameters(engineId int64, osName string, hostname string, username string, password string, ) *EnvironmentCreateParameters`

NewEnvironmentCreateParameters instantiates a new EnvironmentCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateParametersWithDefaults

`func NewEnvironmentCreateParametersWithDefaults() *EnvironmentCreateParameters`

NewEnvironmentCreateParametersWithDefaults instantiates a new EnvironmentCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCreateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentCreateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEngineId

`func (o *EnvironmentCreateParameters) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *EnvironmentCreateParameters) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *EnvironmentCreateParameters) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.


### GetOsName

`func (o *EnvironmentCreateParameters) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *EnvironmentCreateParameters) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *EnvironmentCreateParameters) SetOsName(v string)`

SetOsName sets OsName field to given value.


### GetIsCluster

`func (o *EnvironmentCreateParameters) GetIsCluster() bool`

GetIsCluster returns the IsCluster field if non-nil, zero value otherwise.

### GetIsClusterOk

`func (o *EnvironmentCreateParameters) GetIsClusterOk() (*bool, bool)`

GetIsClusterOk returns a tuple with the IsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCluster

`func (o *EnvironmentCreateParameters) SetIsCluster(v bool)`

SetIsCluster sets IsCluster field to given value.

### HasIsCluster

`func (o *EnvironmentCreateParameters) HasIsCluster() bool`

HasIsCluster returns a boolean if a field has been set.

### GetClusterHome

`func (o *EnvironmentCreateParameters) GetClusterHome() string`

GetClusterHome returns the ClusterHome field if non-nil, zero value otherwise.

### GetClusterHomeOk

`func (o *EnvironmentCreateParameters) GetClusterHomeOk() (*string, bool)`

GetClusterHomeOk returns a tuple with the ClusterHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHome

`func (o *EnvironmentCreateParameters) SetClusterHome(v string)`

SetClusterHome sets ClusterHome field to given value.

### HasClusterHome

`func (o *EnvironmentCreateParameters) HasClusterHome() bool`

HasClusterHome returns a boolean if a field has been set.

### GetHostname

`func (o *EnvironmentCreateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EnvironmentCreateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EnvironmentCreateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStagingEnvironment

`func (o *EnvironmentCreateParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *EnvironmentCreateParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *EnvironmentCreateParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *EnvironmentCreateParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetConnectorPort

`func (o *EnvironmentCreateParameters) GetConnectorPort() int32`

GetConnectorPort returns the ConnectorPort field if non-nil, zero value otherwise.

### GetConnectorPortOk

`func (o *EnvironmentCreateParameters) GetConnectorPortOk() (*int32, bool)`

GetConnectorPortOk returns a tuple with the ConnectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPort

`func (o *EnvironmentCreateParameters) SetConnectorPort(v int32)`

SetConnectorPort sets ConnectorPort field to given value.

### HasConnectorPort

`func (o *EnvironmentCreateParameters) HasConnectorPort() bool`

HasConnectorPort returns a boolean if a field has been set.

### GetSshPort

`func (o *EnvironmentCreateParameters) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *EnvironmentCreateParameters) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *EnvironmentCreateParameters) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *EnvironmentCreateParameters) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetToolkitPath

`func (o *EnvironmentCreateParameters) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *EnvironmentCreateParameters) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *EnvironmentCreateParameters) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *EnvironmentCreateParameters) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetUsername

`func (o *EnvironmentCreateParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentCreateParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentCreateParameters) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *EnvironmentCreateParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentCreateParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentCreateParameters) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetNfsAddresses

`func (o *EnvironmentCreateParameters) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *EnvironmentCreateParameters) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *EnvironmentCreateParameters) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *EnvironmentCreateParameters) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetAseDbUsername

`func (o *EnvironmentCreateParameters) GetAseDbUsername() string`

GetAseDbUsername returns the AseDbUsername field if non-nil, zero value otherwise.

### GetAseDbUsernameOk

`func (o *EnvironmentCreateParameters) GetAseDbUsernameOk() (*string, bool)`

GetAseDbUsernameOk returns a tuple with the AseDbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbUsername

`func (o *EnvironmentCreateParameters) SetAseDbUsername(v string)`

SetAseDbUsername sets AseDbUsername field to given value.

### HasAseDbUsername

`func (o *EnvironmentCreateParameters) HasAseDbUsername() bool`

HasAseDbUsername returns a boolean if a field has been set.

### GetAseDbPassword

`func (o *EnvironmentCreateParameters) GetAseDbPassword() string`

GetAseDbPassword returns the AseDbPassword field if non-nil, zero value otherwise.

### GetAseDbPasswordOk

`func (o *EnvironmentCreateParameters) GetAseDbPasswordOk() (*string, bool)`

GetAseDbPasswordOk returns a tuple with the AseDbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAseDbPassword

`func (o *EnvironmentCreateParameters) SetAseDbPassword(v string)`

SetAseDbPassword sets AseDbPassword field to given value.

### HasAseDbPassword

`func (o *EnvironmentCreateParameters) HasAseDbPassword() bool`

HasAseDbPassword returns a boolean if a field has been set.

### GetJavaHome

`func (o *EnvironmentCreateParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *EnvironmentCreateParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *EnvironmentCreateParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *EnvironmentCreateParameters) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *EnvironmentCreateParameters) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *EnvironmentCreateParameters) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *EnvironmentCreateParameters) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *EnvironmentCreateParameters) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspKeystorePassword

`func (o *EnvironmentCreateParameters) GetDspKeystorePassword() string`

GetDspKeystorePassword returns the DspKeystorePassword field if non-nil, zero value otherwise.

### GetDspKeystorePasswordOk

`func (o *EnvironmentCreateParameters) GetDspKeystorePasswordOk() (*string, bool)`

GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePassword

`func (o *EnvironmentCreateParameters) SetDspKeystorePassword(v string)`

SetDspKeystorePassword sets DspKeystorePassword field to given value.

### HasDspKeystorePassword

`func (o *EnvironmentCreateParameters) HasDspKeystorePassword() bool`

HasDspKeystorePassword returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *EnvironmentCreateParameters) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *EnvironmentCreateParameters) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *EnvironmentCreateParameters) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *EnvironmentCreateParameters) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *EnvironmentCreateParameters) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *EnvironmentCreateParameters) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *EnvironmentCreateParameters) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *EnvironmentCreateParameters) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetDspTruststorePassword

`func (o *EnvironmentCreateParameters) GetDspTruststorePassword() string`

GetDspTruststorePassword returns the DspTruststorePassword field if non-nil, zero value otherwise.

### GetDspTruststorePasswordOk

`func (o *EnvironmentCreateParameters) GetDspTruststorePasswordOk() (*string, bool)`

GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePassword

`func (o *EnvironmentCreateParameters) SetDspTruststorePassword(v string)`

SetDspTruststorePassword sets DspTruststorePassword field to given value.

### HasDspTruststorePassword

`func (o *EnvironmentCreateParameters) HasDspTruststorePassword() bool`

HasDspTruststorePassword returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentCreateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentCreateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentCreateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentCreateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



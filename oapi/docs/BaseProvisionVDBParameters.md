# BaseProvisionVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | Pointer to **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | [optional] 
**EngineId** | Pointer to **int64** | The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored. | [optional] 
**TargetGroupId** | Pointer to **string** | The ID of the group into which the VDB will be provisioned. If unset, a group is selected randomly on the Engine. | [optional] 
**VdbName** | Pointer to **string** | The unique name of the provisioned VDB within a group. If unset, a name is randomly generated. | [optional] 
**DatabaseName** | Pointer to **string** | The name of the database on the target environment. Defaults to vdb_name. | [optional] 
**TruncateLogOnCheckpoint** | Pointer to **bool** | Whether to truncate log on checkpoint (ASE only). | [optional] 
**Username** | Pointer to **string** | The name of the privileged user to run the provision operation (Oracle Only). | [optional] 
**Password** | Pointer to **string** | The password of the privileged user to run the provision operation (Oracle Only). | [optional] 
**EnvironmentId** | Pointer to **string** | The ID of the target environment where to provision the VDB. If repository_id unambigously identifies a repository, this is unnecessary and ignored. Otherwise, a compatible repository is randomly selected on the environment. | [optional] 
**EnvironmentUserId** | Pointer to **string** | The environment user ID to use to connect to the target environment. | [optional] 
**RepositoryId** | Pointer to **string** | The ID of the target repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, ...). Setting this attribute implicitly determines the environment where to provision the VDB. | [optional] 
**AutoSelectRepository** | Pointer to **bool** | Option to automatically select a compatible environment and repository. Mutually exclusive with repository_id. | [optional] 
**PreRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before refreshing the VDB. | [optional] 
**PostRefresh** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after refreshing the VDB. | [optional] 
**PreRollback** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before rewinding the VDB. | [optional] 
**PostRollback** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after rewinding the VDB. | [optional] 
**ConfigureClone** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment when the VDB is created or refreshed. | [optional] 
**PreSnapshot** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting. | [optional] 
**PostSnapshot** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after snapshotting a virtual source. | [optional] 
**PreStart** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before starting a virtual source. | [optional] 
**PostStart** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after starting a virtual source. | [optional] 
**PreStop** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment before stopping a virtual source. | [optional] 
**PostStop** | Pointer to [**[]Hook**](Hook.md) | The commands to execute on the target environment after stopping a virtual source. | [optional] 
**VdbRestart** | Pointer to **bool** | Indicates whether the Engine should automatically restart this virtual source when target host reboot is detected. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the target VDB Template (Oracle Only). | [optional] 
**FileMappingRules** | Pointer to **string** | Target VDB file mapping rules (Oracle Only). Rules must be line separated (\\n or \\r) and each line must have the format \&quot;pattern:replacement\&quot;. Lines are applied in order. | [optional] 
**OracleInstanceName** | Pointer to **string** | Target VDB SID name (Oracle Only). | [optional] 
**UniqueName** | Pointer to **string** | Target VDB db_unique_name (Oracle Only). | [optional] 
**MountPoint** | Pointer to **string** | Mount point for the VDB (Oracle, ASE Only). | [optional] 
**OpenResetLogs** | Pointer to **bool** | Whether to open the database after provision (Oracle Only). | [optional] 
**SnapshotPolicyId** | Pointer to **string** | The ID of the snapshot policy for the VDB. | [optional] 
**RetentionPolicyId** | Pointer to **string** | The ID of the retention policy for the VDB. | [optional] 
**RecoveryModel** | Pointer to **string** | Recovery model of the source database (MSSql Only). | [optional] 
**PreScript** | Pointer to **string** | PowerShell script or executable to run prior to provisioning (MSSql Only). | [optional] 
**PostScript** | Pointer to **string** | PowerShell script or executable to run after provisioning (MSSql Only). | [optional] 
**CdcOnProvision** | Pointer to **bool** | Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (MSSql Only). | [optional] 
**OnlineLogSize** | Pointer to **int32** | Online log size in MB (Oracle Only). | [optional] 
**OnlineLogGroups** | Pointer to **int32** | Number of online log groups (Oracle Only). | [optional] 
**ArchiveLog** | Pointer to **bool** | Option to create a VDB in archivelog mode (Oracle Only). | [optional] 
**NewDbid** | Pointer to **bool** | Option to generate a new DB ID for the created VDB (Oracle Only). | [optional] 
**ListenerIds** | Pointer to **[]string** | The listener IDs for this provision operation (Oracle Only). | [optional] 
**CustomEnvVars** | Pointer to **map[string]string** | Environment variable to be set when the engine creates a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. | [optional] 
**CustomEnvFiles** | Pointer to **[]string** | Environment files to be sourced when the Engine creates a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. | [optional] 

## Methods

### NewBaseProvisionVDBParameters

`func NewBaseProvisionVDBParameters() *BaseProvisionVDBParameters`

NewBaseProvisionVDBParameters instantiates a new BaseProvisionVDBParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseProvisionVDBParametersWithDefaults

`func NewBaseProvisionVDBParametersWithDefaults() *BaseProvisionVDBParameters`

NewBaseProvisionVDBParametersWithDefaults instantiates a new BaseProvisionVDBParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *BaseProvisionVDBParameters) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *BaseProvisionVDBParameters) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *BaseProvisionVDBParameters) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.

### HasSourceDataId

`func (o *BaseProvisionVDBParameters) HasSourceDataId() bool`

HasSourceDataId returns a boolean if a field has been set.

### GetEngineId

`func (o *BaseProvisionVDBParameters) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *BaseProvisionVDBParameters) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *BaseProvisionVDBParameters) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *BaseProvisionVDBParameters) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *BaseProvisionVDBParameters) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *BaseProvisionVDBParameters) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *BaseProvisionVDBParameters) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *BaseProvisionVDBParameters) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetVdbName

`func (o *BaseProvisionVDBParameters) GetVdbName() string`

GetVdbName returns the VdbName field if non-nil, zero value otherwise.

### GetVdbNameOk

`func (o *BaseProvisionVDBParameters) GetVdbNameOk() (*string, bool)`

GetVdbNameOk returns a tuple with the VdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbName

`func (o *BaseProvisionVDBParameters) SetVdbName(v string)`

SetVdbName sets VdbName field to given value.

### HasVdbName

`func (o *BaseProvisionVDBParameters) HasVdbName() bool`

HasVdbName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *BaseProvisionVDBParameters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *BaseProvisionVDBParameters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *BaseProvisionVDBParameters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *BaseProvisionVDBParameters) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParameters) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *BaseProvisionVDBParameters) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParameters) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *BaseProvisionVDBParameters) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetUsername

`func (o *BaseProvisionVDBParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseProvisionVDBParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseProvisionVDBParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BaseProvisionVDBParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *BaseProvisionVDBParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BaseProvisionVDBParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BaseProvisionVDBParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BaseProvisionVDBParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *BaseProvisionVDBParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BaseProvisionVDBParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BaseProvisionVDBParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *BaseProvisionVDBParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *BaseProvisionVDBParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *BaseProvisionVDBParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *BaseProvisionVDBParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *BaseProvisionVDBParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *BaseProvisionVDBParameters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *BaseProvisionVDBParameters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *BaseProvisionVDBParameters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *BaseProvisionVDBParameters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *BaseProvisionVDBParameters) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *BaseProvisionVDBParameters) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *BaseProvisionVDBParameters) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *BaseProvisionVDBParameters) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetPreRefresh

`func (o *BaseProvisionVDBParameters) GetPreRefresh() []Hook`

GetPreRefresh returns the PreRefresh field if non-nil, zero value otherwise.

### GetPreRefreshOk

`func (o *BaseProvisionVDBParameters) GetPreRefreshOk() (*[]Hook, bool)`

GetPreRefreshOk returns a tuple with the PreRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRefresh

`func (o *BaseProvisionVDBParameters) SetPreRefresh(v []Hook)`

SetPreRefresh sets PreRefresh field to given value.

### HasPreRefresh

`func (o *BaseProvisionVDBParameters) HasPreRefresh() bool`

HasPreRefresh returns a boolean if a field has been set.

### GetPostRefresh

`func (o *BaseProvisionVDBParameters) GetPostRefresh() []Hook`

GetPostRefresh returns the PostRefresh field if non-nil, zero value otherwise.

### GetPostRefreshOk

`func (o *BaseProvisionVDBParameters) GetPostRefreshOk() (*[]Hook, bool)`

GetPostRefreshOk returns a tuple with the PostRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRefresh

`func (o *BaseProvisionVDBParameters) SetPostRefresh(v []Hook)`

SetPostRefresh sets PostRefresh field to given value.

### HasPostRefresh

`func (o *BaseProvisionVDBParameters) HasPostRefresh() bool`

HasPostRefresh returns a boolean if a field has been set.

### GetPreRollback

`func (o *BaseProvisionVDBParameters) GetPreRollback() []Hook`

GetPreRollback returns the PreRollback field if non-nil, zero value otherwise.

### GetPreRollbackOk

`func (o *BaseProvisionVDBParameters) GetPreRollbackOk() (*[]Hook, bool)`

GetPreRollbackOk returns a tuple with the PreRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRollback

`func (o *BaseProvisionVDBParameters) SetPreRollback(v []Hook)`

SetPreRollback sets PreRollback field to given value.

### HasPreRollback

`func (o *BaseProvisionVDBParameters) HasPreRollback() bool`

HasPreRollback returns a boolean if a field has been set.

### GetPostRollback

`func (o *BaseProvisionVDBParameters) GetPostRollback() []Hook`

GetPostRollback returns the PostRollback field if non-nil, zero value otherwise.

### GetPostRollbackOk

`func (o *BaseProvisionVDBParameters) GetPostRollbackOk() (*[]Hook, bool)`

GetPostRollbackOk returns a tuple with the PostRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRollback

`func (o *BaseProvisionVDBParameters) SetPostRollback(v []Hook)`

SetPostRollback sets PostRollback field to given value.

### HasPostRollback

`func (o *BaseProvisionVDBParameters) HasPostRollback() bool`

HasPostRollback returns a boolean if a field has been set.

### GetConfigureClone

`func (o *BaseProvisionVDBParameters) GetConfigureClone() []Hook`

GetConfigureClone returns the ConfigureClone field if non-nil, zero value otherwise.

### GetConfigureCloneOk

`func (o *BaseProvisionVDBParameters) GetConfigureCloneOk() (*[]Hook, bool)`

GetConfigureCloneOk returns a tuple with the ConfigureClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureClone

`func (o *BaseProvisionVDBParameters) SetConfigureClone(v []Hook)`

SetConfigureClone sets ConfigureClone field to given value.

### HasConfigureClone

`func (o *BaseProvisionVDBParameters) HasConfigureClone() bool`

HasConfigureClone returns a boolean if a field has been set.

### GetPreSnapshot

`func (o *BaseProvisionVDBParameters) GetPreSnapshot() []Hook`

GetPreSnapshot returns the PreSnapshot field if non-nil, zero value otherwise.

### GetPreSnapshotOk

`func (o *BaseProvisionVDBParameters) GetPreSnapshotOk() (*[]Hook, bool)`

GetPreSnapshotOk returns a tuple with the PreSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshot

`func (o *BaseProvisionVDBParameters) SetPreSnapshot(v []Hook)`

SetPreSnapshot sets PreSnapshot field to given value.

### HasPreSnapshot

`func (o *BaseProvisionVDBParameters) HasPreSnapshot() bool`

HasPreSnapshot returns a boolean if a field has been set.

### GetPostSnapshot

`func (o *BaseProvisionVDBParameters) GetPostSnapshot() []Hook`

GetPostSnapshot returns the PostSnapshot field if non-nil, zero value otherwise.

### GetPostSnapshotOk

`func (o *BaseProvisionVDBParameters) GetPostSnapshotOk() (*[]Hook, bool)`

GetPostSnapshotOk returns a tuple with the PostSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshot

`func (o *BaseProvisionVDBParameters) SetPostSnapshot(v []Hook)`

SetPostSnapshot sets PostSnapshot field to given value.

### HasPostSnapshot

`func (o *BaseProvisionVDBParameters) HasPostSnapshot() bool`

HasPostSnapshot returns a boolean if a field has been set.

### GetPreStart

`func (o *BaseProvisionVDBParameters) GetPreStart() []Hook`

GetPreStart returns the PreStart field if non-nil, zero value otherwise.

### GetPreStartOk

`func (o *BaseProvisionVDBParameters) GetPreStartOk() (*[]Hook, bool)`

GetPreStartOk returns a tuple with the PreStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStart

`func (o *BaseProvisionVDBParameters) SetPreStart(v []Hook)`

SetPreStart sets PreStart field to given value.

### HasPreStart

`func (o *BaseProvisionVDBParameters) HasPreStart() bool`

HasPreStart returns a boolean if a field has been set.

### GetPostStart

`func (o *BaseProvisionVDBParameters) GetPostStart() []Hook`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *BaseProvisionVDBParameters) GetPostStartOk() (*[]Hook, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *BaseProvisionVDBParameters) SetPostStart(v []Hook)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *BaseProvisionVDBParameters) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *BaseProvisionVDBParameters) GetPreStop() []Hook`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *BaseProvisionVDBParameters) GetPreStopOk() (*[]Hook, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *BaseProvisionVDBParameters) SetPreStop(v []Hook)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *BaseProvisionVDBParameters) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.

### GetPostStop

`func (o *BaseProvisionVDBParameters) GetPostStop() []Hook`

GetPostStop returns the PostStop field if non-nil, zero value otherwise.

### GetPostStopOk

`func (o *BaseProvisionVDBParameters) GetPostStopOk() (*[]Hook, bool)`

GetPostStopOk returns a tuple with the PostStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStop

`func (o *BaseProvisionVDBParameters) SetPostStop(v []Hook)`

SetPostStop sets PostStop field to given value.

### HasPostStop

`func (o *BaseProvisionVDBParameters) HasPostStop() bool`

HasPostStop returns a boolean if a field has been set.

### GetVdbRestart

`func (o *BaseProvisionVDBParameters) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *BaseProvisionVDBParameters) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *BaseProvisionVDBParameters) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *BaseProvisionVDBParameters) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *BaseProvisionVDBParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *BaseProvisionVDBParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *BaseProvisionVDBParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *BaseProvisionVDBParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *BaseProvisionVDBParameters) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *BaseProvisionVDBParameters) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *BaseProvisionVDBParameters) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *BaseProvisionVDBParameters) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *BaseProvisionVDBParameters) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *BaseProvisionVDBParameters) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *BaseProvisionVDBParameters) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *BaseProvisionVDBParameters) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *BaseProvisionVDBParameters) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *BaseProvisionVDBParameters) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *BaseProvisionVDBParameters) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *BaseProvisionVDBParameters) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetMountPoint

`func (o *BaseProvisionVDBParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *BaseProvisionVDBParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *BaseProvisionVDBParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *BaseProvisionVDBParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *BaseProvisionVDBParameters) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *BaseProvisionVDBParameters) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *BaseProvisionVDBParameters) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *BaseProvisionVDBParameters) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *BaseProvisionVDBParameters) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *BaseProvisionVDBParameters) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *BaseProvisionVDBParameters) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *BaseProvisionVDBParameters) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *BaseProvisionVDBParameters) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *BaseProvisionVDBParameters) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *BaseProvisionVDBParameters) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *BaseProvisionVDBParameters) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *BaseProvisionVDBParameters) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *BaseProvisionVDBParameters) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *BaseProvisionVDBParameters) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *BaseProvisionVDBParameters) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *BaseProvisionVDBParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *BaseProvisionVDBParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *BaseProvisionVDBParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *BaseProvisionVDBParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *BaseProvisionVDBParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *BaseProvisionVDBParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *BaseProvisionVDBParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *BaseProvisionVDBParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *BaseProvisionVDBParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *BaseProvisionVDBParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *BaseProvisionVDBParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *BaseProvisionVDBParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *BaseProvisionVDBParameters) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *BaseProvisionVDBParameters) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *BaseProvisionVDBParameters) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *BaseProvisionVDBParameters) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *BaseProvisionVDBParameters) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *BaseProvisionVDBParameters) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *BaseProvisionVDBParameters) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *BaseProvisionVDBParameters) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *BaseProvisionVDBParameters) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *BaseProvisionVDBParameters) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *BaseProvisionVDBParameters) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *BaseProvisionVDBParameters) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *BaseProvisionVDBParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *BaseProvisionVDBParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *BaseProvisionVDBParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *BaseProvisionVDBParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *BaseProvisionVDBParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *BaseProvisionVDBParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *BaseProvisionVDBParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *BaseProvisionVDBParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *BaseProvisionVDBParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *BaseProvisionVDBParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *BaseProvisionVDBParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *BaseProvisionVDBParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *BaseProvisionVDBParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *BaseProvisionVDBParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *BaseProvisionVDBParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *BaseProvisionVDBParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



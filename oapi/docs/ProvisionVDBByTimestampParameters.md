# ProvisionVDBByTimestampParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDataId** | **string** | The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source. | 
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
**Timestamp** | Pointer to **time.Time** | The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point. | [optional] 
**TimestampInDatabaseTimezone** | Pointer to **string** | The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp. | [optional] 

## Methods

### NewProvisionVDBByTimestampParameters

`func NewProvisionVDBByTimestampParameters(sourceDataId string, ) *ProvisionVDBByTimestampParameters`

NewProvisionVDBByTimestampParameters instantiates a new ProvisionVDBByTimestampParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBByTimestampParametersWithDefaults

`func NewProvisionVDBByTimestampParametersWithDefaults() *ProvisionVDBByTimestampParameters`

NewProvisionVDBByTimestampParametersWithDefaults instantiates a new ProvisionVDBByTimestampParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDataId

`func (o *ProvisionVDBByTimestampParameters) GetSourceDataId() string`

GetSourceDataId returns the SourceDataId field if non-nil, zero value otherwise.

### GetSourceDataIdOk

`func (o *ProvisionVDBByTimestampParameters) GetSourceDataIdOk() (*string, bool)`

GetSourceDataIdOk returns a tuple with the SourceDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataId

`func (o *ProvisionVDBByTimestampParameters) SetSourceDataId(v string)`

SetSourceDataId sets SourceDataId field to given value.


### GetEngineId

`func (o *ProvisionVDBByTimestampParameters) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ProvisionVDBByTimestampParameters) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ProvisionVDBByTimestampParameters) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ProvisionVDBByTimestampParameters) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ProvisionVDBByTimestampParameters) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetVdbName

`func (o *ProvisionVDBByTimestampParameters) GetVdbName() string`

GetVdbName returns the VdbName field if non-nil, zero value otherwise.

### GetVdbNameOk

`func (o *ProvisionVDBByTimestampParameters) GetVdbNameOk() (*string, bool)`

GetVdbNameOk returns a tuple with the VdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbName

`func (o *ProvisionVDBByTimestampParameters) SetVdbName(v string)`

SetVdbName sets VdbName field to given value.

### HasVdbName

`func (o *ProvisionVDBByTimestampParameters) HasVdbName() bool`

HasVdbName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ProvisionVDBByTimestampParameters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ProvisionVDBByTimestampParameters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ProvisionVDBByTimestampParameters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ProvisionVDBByTimestampParameters) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) GetTruncateLogOnCheckpoint() bool`

GetTruncateLogOnCheckpoint returns the TruncateLogOnCheckpoint field if non-nil, zero value otherwise.

### GetTruncateLogOnCheckpointOk

`func (o *ProvisionVDBByTimestampParameters) GetTruncateLogOnCheckpointOk() (*bool, bool)`

GetTruncateLogOnCheckpointOk returns a tuple with the TruncateLogOnCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) SetTruncateLogOnCheckpoint(v bool)`

SetTruncateLogOnCheckpoint sets TruncateLogOnCheckpoint field to given value.

### HasTruncateLogOnCheckpoint

`func (o *ProvisionVDBByTimestampParameters) HasTruncateLogOnCheckpoint() bool`

HasTruncateLogOnCheckpoint returns a boolean if a field has been set.

### GetUsername

`func (o *ProvisionVDBByTimestampParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProvisionVDBByTimestampParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProvisionVDBByTimestampParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProvisionVDBByTimestampParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ProvisionVDBByTimestampParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProvisionVDBByTimestampParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProvisionVDBByTimestampParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProvisionVDBByTimestampParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ProvisionVDBByTimestampParameters) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *ProvisionVDBByTimestampParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *ProvisionVDBByTimestampParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ProvisionVDBByTimestampParameters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ProvisionVDBByTimestampParameters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ProvisionVDBByTimestampParameters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ProvisionVDBByTimestampParameters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) GetAutoSelectRepository() bool`

GetAutoSelectRepository returns the AutoSelectRepository field if non-nil, zero value otherwise.

### GetAutoSelectRepositoryOk

`func (o *ProvisionVDBByTimestampParameters) GetAutoSelectRepositoryOk() (*bool, bool)`

GetAutoSelectRepositoryOk returns a tuple with the AutoSelectRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) SetAutoSelectRepository(v bool)`

SetAutoSelectRepository sets AutoSelectRepository field to given value.

### HasAutoSelectRepository

`func (o *ProvisionVDBByTimestampParameters) HasAutoSelectRepository() bool`

HasAutoSelectRepository returns a boolean if a field has been set.

### GetPreRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPreRefresh() []Hook`

GetPreRefresh returns the PreRefresh field if non-nil, zero value otherwise.

### GetPreRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPreRefreshOk() (*[]Hook, bool)`

GetPreRefreshOk returns a tuple with the PreRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPreRefresh(v []Hook)`

SetPreRefresh sets PreRefresh field to given value.

### HasPreRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPreRefresh() bool`

HasPreRefresh returns a boolean if a field has been set.

### GetPostRefresh

`func (o *ProvisionVDBByTimestampParameters) GetPostRefresh() []Hook`

GetPostRefresh returns the PostRefresh field if non-nil, zero value otherwise.

### GetPostRefreshOk

`func (o *ProvisionVDBByTimestampParameters) GetPostRefreshOk() (*[]Hook, bool)`

GetPostRefreshOk returns a tuple with the PostRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRefresh

`func (o *ProvisionVDBByTimestampParameters) SetPostRefresh(v []Hook)`

SetPostRefresh sets PostRefresh field to given value.

### HasPostRefresh

`func (o *ProvisionVDBByTimestampParameters) HasPostRefresh() bool`

HasPostRefresh returns a boolean if a field has been set.

### GetPreRollback

`func (o *ProvisionVDBByTimestampParameters) GetPreRollback() []Hook`

GetPreRollback returns the PreRollback field if non-nil, zero value otherwise.

### GetPreRollbackOk

`func (o *ProvisionVDBByTimestampParameters) GetPreRollbackOk() (*[]Hook, bool)`

GetPreRollbackOk returns a tuple with the PreRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRollback

`func (o *ProvisionVDBByTimestampParameters) SetPreRollback(v []Hook)`

SetPreRollback sets PreRollback field to given value.

### HasPreRollback

`func (o *ProvisionVDBByTimestampParameters) HasPreRollback() bool`

HasPreRollback returns a boolean if a field has been set.

### GetPostRollback

`func (o *ProvisionVDBByTimestampParameters) GetPostRollback() []Hook`

GetPostRollback returns the PostRollback field if non-nil, zero value otherwise.

### GetPostRollbackOk

`func (o *ProvisionVDBByTimestampParameters) GetPostRollbackOk() (*[]Hook, bool)`

GetPostRollbackOk returns a tuple with the PostRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRollback

`func (o *ProvisionVDBByTimestampParameters) SetPostRollback(v []Hook)`

SetPostRollback sets PostRollback field to given value.

### HasPostRollback

`func (o *ProvisionVDBByTimestampParameters) HasPostRollback() bool`

HasPostRollback returns a boolean if a field has been set.

### GetConfigureClone

`func (o *ProvisionVDBByTimestampParameters) GetConfigureClone() []Hook`

GetConfigureClone returns the ConfigureClone field if non-nil, zero value otherwise.

### GetConfigureCloneOk

`func (o *ProvisionVDBByTimestampParameters) GetConfigureCloneOk() (*[]Hook, bool)`

GetConfigureCloneOk returns a tuple with the ConfigureClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureClone

`func (o *ProvisionVDBByTimestampParameters) SetConfigureClone(v []Hook)`

SetConfigureClone sets ConfigureClone field to given value.

### HasConfigureClone

`func (o *ProvisionVDBByTimestampParameters) HasConfigureClone() bool`

HasConfigureClone returns a boolean if a field has been set.

### GetPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) GetPreSnapshot() []Hook`

GetPreSnapshot returns the PreSnapshot field if non-nil, zero value otherwise.

### GetPreSnapshotOk

`func (o *ProvisionVDBByTimestampParameters) GetPreSnapshotOk() (*[]Hook, bool)`

GetPreSnapshotOk returns a tuple with the PreSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) SetPreSnapshot(v []Hook)`

SetPreSnapshot sets PreSnapshot field to given value.

### HasPreSnapshot

`func (o *ProvisionVDBByTimestampParameters) HasPreSnapshot() bool`

HasPreSnapshot returns a boolean if a field has been set.

### GetPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) GetPostSnapshot() []Hook`

GetPostSnapshot returns the PostSnapshot field if non-nil, zero value otherwise.

### GetPostSnapshotOk

`func (o *ProvisionVDBByTimestampParameters) GetPostSnapshotOk() (*[]Hook, bool)`

GetPostSnapshotOk returns a tuple with the PostSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) SetPostSnapshot(v []Hook)`

SetPostSnapshot sets PostSnapshot field to given value.

### HasPostSnapshot

`func (o *ProvisionVDBByTimestampParameters) HasPostSnapshot() bool`

HasPostSnapshot returns a boolean if a field has been set.

### GetPreStart

`func (o *ProvisionVDBByTimestampParameters) GetPreStart() []Hook`

GetPreStart returns the PreStart field if non-nil, zero value otherwise.

### GetPreStartOk

`func (o *ProvisionVDBByTimestampParameters) GetPreStartOk() (*[]Hook, bool)`

GetPreStartOk returns a tuple with the PreStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStart

`func (o *ProvisionVDBByTimestampParameters) SetPreStart(v []Hook)`

SetPreStart sets PreStart field to given value.

### HasPreStart

`func (o *ProvisionVDBByTimestampParameters) HasPreStart() bool`

HasPreStart returns a boolean if a field has been set.

### GetPostStart

`func (o *ProvisionVDBByTimestampParameters) GetPostStart() []Hook`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *ProvisionVDBByTimestampParameters) GetPostStartOk() (*[]Hook, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *ProvisionVDBByTimestampParameters) SetPostStart(v []Hook)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *ProvisionVDBByTimestampParameters) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *ProvisionVDBByTimestampParameters) GetPreStop() []Hook`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *ProvisionVDBByTimestampParameters) GetPreStopOk() (*[]Hook, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *ProvisionVDBByTimestampParameters) SetPreStop(v []Hook)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *ProvisionVDBByTimestampParameters) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.

### GetPostStop

`func (o *ProvisionVDBByTimestampParameters) GetPostStop() []Hook`

GetPostStop returns the PostStop field if non-nil, zero value otherwise.

### GetPostStopOk

`func (o *ProvisionVDBByTimestampParameters) GetPostStopOk() (*[]Hook, bool)`

GetPostStopOk returns a tuple with the PostStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStop

`func (o *ProvisionVDBByTimestampParameters) SetPostStop(v []Hook)`

SetPostStop sets PostStop field to given value.

### HasPostStop

`func (o *ProvisionVDBByTimestampParameters) HasPostStop() bool`

HasPostStop returns a boolean if a field has been set.

### GetVdbRestart

`func (o *ProvisionVDBByTimestampParameters) GetVdbRestart() bool`

GetVdbRestart returns the VdbRestart field if non-nil, zero value otherwise.

### GetVdbRestartOk

`func (o *ProvisionVDBByTimestampParameters) GetVdbRestartOk() (*bool, bool)`

GetVdbRestartOk returns a tuple with the VdbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRestart

`func (o *ProvisionVDBByTimestampParameters) SetVdbRestart(v bool)`

SetVdbRestart sets VdbRestart field to given value.

### HasVdbRestart

`func (o *ProvisionVDBByTimestampParameters) HasVdbRestart() bool`

HasVdbRestart returns a boolean if a field has been set.

### GetTemplateId

`func (o *ProvisionVDBByTimestampParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProvisionVDBByTimestampParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProvisionVDBByTimestampParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProvisionVDBByTimestampParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) GetFileMappingRules() string`

GetFileMappingRules returns the FileMappingRules field if non-nil, zero value otherwise.

### GetFileMappingRulesOk

`func (o *ProvisionVDBByTimestampParameters) GetFileMappingRulesOk() (*string, bool)`

GetFileMappingRulesOk returns a tuple with the FileMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) SetFileMappingRules(v string)`

SetFileMappingRules sets FileMappingRules field to given value.

### HasFileMappingRules

`func (o *ProvisionVDBByTimestampParameters) HasFileMappingRules() bool`

HasFileMappingRules returns a boolean if a field has been set.

### GetOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) GetOracleInstanceName() string`

GetOracleInstanceName returns the OracleInstanceName field if non-nil, zero value otherwise.

### GetOracleInstanceNameOk

`func (o *ProvisionVDBByTimestampParameters) GetOracleInstanceNameOk() (*string, bool)`

GetOracleInstanceNameOk returns a tuple with the OracleInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) SetOracleInstanceName(v string)`

SetOracleInstanceName sets OracleInstanceName field to given value.

### HasOracleInstanceName

`func (o *ProvisionVDBByTimestampParameters) HasOracleInstanceName() bool`

HasOracleInstanceName returns a boolean if a field has been set.

### GetUniqueName

`func (o *ProvisionVDBByTimestampParameters) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProvisionVDBByTimestampParameters) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProvisionVDBByTimestampParameters) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProvisionVDBByTimestampParameters) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetMountPoint

`func (o *ProvisionVDBByTimestampParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ProvisionVDBByTimestampParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ProvisionVDBByTimestampParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ProvisionVDBByTimestampParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) GetOpenResetLogs() bool`

GetOpenResetLogs returns the OpenResetLogs field if non-nil, zero value otherwise.

### GetOpenResetLogsOk

`func (o *ProvisionVDBByTimestampParameters) GetOpenResetLogsOk() (*bool, bool)`

GetOpenResetLogsOk returns a tuple with the OpenResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) SetOpenResetLogs(v bool)`

SetOpenResetLogs sets OpenResetLogs field to given value.

### HasOpenResetLogs

`func (o *ProvisionVDBByTimestampParameters) HasOpenResetLogs() bool`

HasOpenResetLogs returns a boolean if a field has been set.

### GetSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) GetSnapshotPolicyId() string`

GetSnapshotPolicyId returns the SnapshotPolicyId field if non-nil, zero value otherwise.

### GetSnapshotPolicyIdOk

`func (o *ProvisionVDBByTimestampParameters) GetSnapshotPolicyIdOk() (*string, bool)`

GetSnapshotPolicyIdOk returns a tuple with the SnapshotPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) SetSnapshotPolicyId(v string)`

SetSnapshotPolicyId sets SnapshotPolicyId field to given value.

### HasSnapshotPolicyId

`func (o *ProvisionVDBByTimestampParameters) HasSnapshotPolicyId() bool`

HasSnapshotPolicyId returns a boolean if a field has been set.

### GetRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) GetRetentionPolicyId() string`

GetRetentionPolicyId returns the RetentionPolicyId field if non-nil, zero value otherwise.

### GetRetentionPolicyIdOk

`func (o *ProvisionVDBByTimestampParameters) GetRetentionPolicyIdOk() (*string, bool)`

GetRetentionPolicyIdOk returns a tuple with the RetentionPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) SetRetentionPolicyId(v string)`

SetRetentionPolicyId sets RetentionPolicyId field to given value.

### HasRetentionPolicyId

`func (o *ProvisionVDBByTimestampParameters) HasRetentionPolicyId() bool`

HasRetentionPolicyId returns a boolean if a field has been set.

### GetRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *ProvisionVDBByTimestampParameters) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *ProvisionVDBByTimestampParameters) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### GetPreScript

`func (o *ProvisionVDBByTimestampParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *ProvisionVDBByTimestampParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *ProvisionVDBByTimestampParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *ProvisionVDBByTimestampParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *ProvisionVDBByTimestampParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *ProvisionVDBByTimestampParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *ProvisionVDBByTimestampParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *ProvisionVDBByTimestampParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *ProvisionVDBByTimestampParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *ProvisionVDBByTimestampParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogSize() int32`

GetOnlineLogSize returns the OnlineLogSize field if non-nil, zero value otherwise.

### GetOnlineLogSizeOk

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogSizeOk() (*int32, bool)`

GetOnlineLogSizeOk returns a tuple with the OnlineLogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) SetOnlineLogSize(v int32)`

SetOnlineLogSize sets OnlineLogSize field to given value.

### HasOnlineLogSize

`func (o *ProvisionVDBByTimestampParameters) HasOnlineLogSize() bool`

HasOnlineLogSize returns a boolean if a field has been set.

### GetOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogGroups() int32`

GetOnlineLogGroups returns the OnlineLogGroups field if non-nil, zero value otherwise.

### GetOnlineLogGroupsOk

`func (o *ProvisionVDBByTimestampParameters) GetOnlineLogGroupsOk() (*int32, bool)`

GetOnlineLogGroupsOk returns a tuple with the OnlineLogGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) SetOnlineLogGroups(v int32)`

SetOnlineLogGroups sets OnlineLogGroups field to given value.

### HasOnlineLogGroups

`func (o *ProvisionVDBByTimestampParameters) HasOnlineLogGroups() bool`

HasOnlineLogGroups returns a boolean if a field has been set.

### GetArchiveLog

`func (o *ProvisionVDBByTimestampParameters) GetArchiveLog() bool`

GetArchiveLog returns the ArchiveLog field if non-nil, zero value otherwise.

### GetArchiveLogOk

`func (o *ProvisionVDBByTimestampParameters) GetArchiveLogOk() (*bool, bool)`

GetArchiveLogOk returns a tuple with the ArchiveLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLog

`func (o *ProvisionVDBByTimestampParameters) SetArchiveLog(v bool)`

SetArchiveLog sets ArchiveLog field to given value.

### HasArchiveLog

`func (o *ProvisionVDBByTimestampParameters) HasArchiveLog() bool`

HasArchiveLog returns a boolean if a field has been set.

### GetNewDbid

`func (o *ProvisionVDBByTimestampParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *ProvisionVDBByTimestampParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *ProvisionVDBByTimestampParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *ProvisionVDBByTimestampParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetListenerIds

`func (o *ProvisionVDBByTimestampParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *ProvisionVDBByTimestampParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *ProvisionVDBByTimestampParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *ProvisionVDBByTimestampParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *ProvisionVDBByTimestampParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *ProvisionVDBByTimestampParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *ProvisionVDBByTimestampParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProvisionVDBByTimestampParameters) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProvisionVDBByTimestampParameters) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProvisionVDBByTimestampParameters) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProvisionVDBByTimestampParameters) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) GetTimestampInDatabaseTimezone() string`

GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field if non-nil, zero value otherwise.

### GetTimestampInDatabaseTimezoneOk

`func (o *ProvisionVDBByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool)`

GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) SetTimestampInDatabaseTimezone(v string)`

SetTimestampInDatabaseTimezone sets TimestampInDatabaseTimezone field to given value.

### HasTimestampInDatabaseTimezone

`func (o *ProvisionVDBByTimestampParameters) HasTimestampInDatabaseTimezone() bool`

HasTimestampInDatabaseTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



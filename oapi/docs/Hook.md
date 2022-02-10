# Hook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Command** | **string** |  | 
**Shell** | Pointer to **string** |  | [optional] [default to "bash"]

## Methods

### NewHook

`func NewHook(command string, ) *Hook`

NewHook instantiates a new Hook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookWithDefaults

`func NewHookWithDefaults() *Hook`

NewHookWithDefaults instantiates a new Hook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Hook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Hook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Hook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Hook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommand

`func (o *Hook) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Hook) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Hook) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetShell

`func (o *Hook) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *Hook) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *Hook) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *Hook) HasShell() bool`

HasShell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



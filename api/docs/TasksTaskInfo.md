# TasksTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CurrentProgressPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastExecutionResult** | Pointer to [**TasksTaskResult**](TasksTaskResult.md) |  | [optional] 
**Triggers** | Pointer to [**[]TasksTaskTriggerInfo**](TasksTaskTriggerInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewTasksTaskInfo

`func NewTasksTaskInfo() *TasksTaskInfo`

NewTasksTaskInfo instantiates a new TasksTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksTaskInfoWithDefaults

`func NewTasksTaskInfoWithDefaults() *TasksTaskInfo`

NewTasksTaskInfoWithDefaults instantiates a new TasksTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TasksTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TasksTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TasksTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TasksTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *TasksTaskInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TasksTaskInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TasksTaskInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TasksTaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentProgressPercentage

`func (o *TasksTaskInfo) GetCurrentProgressPercentage() float64`

GetCurrentProgressPercentage returns the CurrentProgressPercentage field if non-nil, zero value otherwise.

### GetCurrentProgressPercentageOk

`func (o *TasksTaskInfo) GetCurrentProgressPercentageOk() (*float64, bool)`

GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgressPercentage

`func (o *TasksTaskInfo) SetCurrentProgressPercentage(v float64)`

SetCurrentProgressPercentage sets CurrentProgressPercentage field to given value.

### HasCurrentProgressPercentage

`func (o *TasksTaskInfo) HasCurrentProgressPercentage() bool`

HasCurrentProgressPercentage returns a boolean if a field has been set.

### SetCurrentProgressPercentageNil

`func (o *TasksTaskInfo) SetCurrentProgressPercentageNil(b bool)`

 SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil

### UnsetCurrentProgressPercentage
`func (o *TasksTaskInfo) UnsetCurrentProgressPercentage()`

UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
### GetId

`func (o *TasksTaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TasksTaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TasksTaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TasksTaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastExecutionResult

`func (o *TasksTaskInfo) GetLastExecutionResult() TasksTaskResult`

GetLastExecutionResult returns the LastExecutionResult field if non-nil, zero value otherwise.

### GetLastExecutionResultOk

`func (o *TasksTaskInfo) GetLastExecutionResultOk() (*TasksTaskResult, bool)`

GetLastExecutionResultOk returns a tuple with the LastExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionResult

`func (o *TasksTaskInfo) SetLastExecutionResult(v TasksTaskResult)`

SetLastExecutionResult sets LastExecutionResult field to given value.

### HasLastExecutionResult

`func (o *TasksTaskInfo) HasLastExecutionResult() bool`

HasLastExecutionResult returns a boolean if a field has been set.

### GetTriggers

`func (o *TasksTaskInfo) GetTriggers() []TasksTaskTriggerInfo`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TasksTaskInfo) GetTriggersOk() (*[]TasksTaskTriggerInfo, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TasksTaskInfo) SetTriggers(v []TasksTaskTriggerInfo)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TasksTaskInfo) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetDescription

`func (o *TasksTaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TasksTaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TasksTaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TasksTaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *TasksTaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TasksTaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TasksTaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TasksTaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsHidden

`func (o *TasksTaskInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *TasksTaskInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *TasksTaskInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *TasksTaskInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetKey

`func (o *TasksTaskInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TasksTaskInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TasksTaskInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TasksTaskInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



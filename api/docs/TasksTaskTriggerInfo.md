# TasksTaskTriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**TimeOfDayTicks** | Pointer to **NullableInt64** |  | [optional] 
**IntervalTicks** | Pointer to **NullableInt64** |  | [optional] 
**SystemEvent** | Pointer to **string** |  | [optional] 
**DayOfWeek** | Pointer to **string** |  | [optional] 
**MaxRuntimeTicks** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewTasksTaskTriggerInfo

`func NewTasksTaskTriggerInfo() *TasksTaskTriggerInfo`

NewTasksTaskTriggerInfo instantiates a new TasksTaskTriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksTaskTriggerInfoWithDefaults

`func NewTasksTaskTriggerInfoWithDefaults() *TasksTaskTriggerInfo`

NewTasksTaskTriggerInfoWithDefaults instantiates a new TasksTaskTriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TasksTaskTriggerInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TasksTaskTriggerInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TasksTaskTriggerInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TasksTaskTriggerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeOfDayTicks

`func (o *TasksTaskTriggerInfo) GetTimeOfDayTicks() int64`

GetTimeOfDayTicks returns the TimeOfDayTicks field if non-nil, zero value otherwise.

### GetTimeOfDayTicksOk

`func (o *TasksTaskTriggerInfo) GetTimeOfDayTicksOk() (*int64, bool)`

GetTimeOfDayTicksOk returns a tuple with the TimeOfDayTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayTicks

`func (o *TasksTaskTriggerInfo) SetTimeOfDayTicks(v int64)`

SetTimeOfDayTicks sets TimeOfDayTicks field to given value.

### HasTimeOfDayTicks

`func (o *TasksTaskTriggerInfo) HasTimeOfDayTicks() bool`

HasTimeOfDayTicks returns a boolean if a field has been set.

### SetTimeOfDayTicksNil

`func (o *TasksTaskTriggerInfo) SetTimeOfDayTicksNil(b bool)`

 SetTimeOfDayTicksNil sets the value for TimeOfDayTicks to be an explicit nil

### UnsetTimeOfDayTicks
`func (o *TasksTaskTriggerInfo) UnsetTimeOfDayTicks()`

UnsetTimeOfDayTicks ensures that no value is present for TimeOfDayTicks, not even an explicit nil
### GetIntervalTicks

`func (o *TasksTaskTriggerInfo) GetIntervalTicks() int64`

GetIntervalTicks returns the IntervalTicks field if non-nil, zero value otherwise.

### GetIntervalTicksOk

`func (o *TasksTaskTriggerInfo) GetIntervalTicksOk() (*int64, bool)`

GetIntervalTicksOk returns a tuple with the IntervalTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalTicks

`func (o *TasksTaskTriggerInfo) SetIntervalTicks(v int64)`

SetIntervalTicks sets IntervalTicks field to given value.

### HasIntervalTicks

`func (o *TasksTaskTriggerInfo) HasIntervalTicks() bool`

HasIntervalTicks returns a boolean if a field has been set.

### SetIntervalTicksNil

`func (o *TasksTaskTriggerInfo) SetIntervalTicksNil(b bool)`

 SetIntervalTicksNil sets the value for IntervalTicks to be an explicit nil

### UnsetIntervalTicks
`func (o *TasksTaskTriggerInfo) UnsetIntervalTicks()`

UnsetIntervalTicks ensures that no value is present for IntervalTicks, not even an explicit nil
### GetSystemEvent

`func (o *TasksTaskTriggerInfo) GetSystemEvent() string`

GetSystemEvent returns the SystemEvent field if non-nil, zero value otherwise.

### GetSystemEventOk

`func (o *TasksTaskTriggerInfo) GetSystemEventOk() (*string, bool)`

GetSystemEventOk returns a tuple with the SystemEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemEvent

`func (o *TasksTaskTriggerInfo) SetSystemEvent(v string)`

SetSystemEvent sets SystemEvent field to given value.

### HasSystemEvent

`func (o *TasksTaskTriggerInfo) HasSystemEvent() bool`

HasSystemEvent returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *TasksTaskTriggerInfo) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TasksTaskTriggerInfo) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TasksTaskTriggerInfo) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TasksTaskTriggerInfo) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetMaxRuntimeTicks

`func (o *TasksTaskTriggerInfo) GetMaxRuntimeTicks() int64`

GetMaxRuntimeTicks returns the MaxRuntimeTicks field if non-nil, zero value otherwise.

### GetMaxRuntimeTicksOk

`func (o *TasksTaskTriggerInfo) GetMaxRuntimeTicksOk() (*int64, bool)`

GetMaxRuntimeTicksOk returns a tuple with the MaxRuntimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRuntimeTicks

`func (o *TasksTaskTriggerInfo) SetMaxRuntimeTicks(v int64)`

SetMaxRuntimeTicks sets MaxRuntimeTicks field to given value.

### HasMaxRuntimeTicks

`func (o *TasksTaskTriggerInfo) HasMaxRuntimeTicks() bool`

HasMaxRuntimeTicks returns a boolean if a field has been set.

### SetMaxRuntimeTicksNil

`func (o *TasksTaskTriggerInfo) SetMaxRuntimeTicksNil(b bool)`

 SetMaxRuntimeTicksNil sets the value for MaxRuntimeTicks to be an explicit nil

### UnsetMaxRuntimeTicks
`func (o *TasksTaskTriggerInfo) UnsetMaxRuntimeTicks()`

UnsetMaxRuntimeTicks ensures that no value is present for MaxRuntimeTicks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



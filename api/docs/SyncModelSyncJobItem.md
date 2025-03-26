# SyncModelSyncJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**JobId** | Pointer to **int64** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to [**MediaSourceInfo**](MediaSourceInfo.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**OutputPath** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**PrimaryImageItemId** | Pointer to **int64** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**TemporaryPath** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]SyncModelItemFileInfo**](SyncModelItemFileInfo.md) |  | [optional] 
**ItemDateModifiedTicks** | Pointer to **int64** |  | [optional] 

## Methods

### NewSyncModelSyncJobItem

`func NewSyncModelSyncJobItem() *SyncModelSyncJobItem`

NewSyncModelSyncJobItem instantiates a new SyncModelSyncJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncModelSyncJobItemWithDefaults

`func NewSyncModelSyncJobItemWithDefaults() *SyncModelSyncJobItem`

NewSyncModelSyncJobItemWithDefaults instantiates a new SyncModelSyncJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncModelSyncJobItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncModelSyncJobItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncModelSyncJobItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SyncModelSyncJobItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *SyncModelSyncJobItem) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SyncModelSyncJobItem) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SyncModelSyncJobItem) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SyncModelSyncJobItem) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetItemId

`func (o *SyncModelSyncJobItem) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SyncModelSyncJobItem) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SyncModelSyncJobItem) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SyncModelSyncJobItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemName

`func (o *SyncModelSyncJobItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *SyncModelSyncJobItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *SyncModelSyncJobItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *SyncModelSyncJobItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetMediaSourceId

`func (o *SyncModelSyncJobItem) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *SyncModelSyncJobItem) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *SyncModelSyncJobItem) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *SyncModelSyncJobItem) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetMediaSource

`func (o *SyncModelSyncJobItem) GetMediaSource() MediaSourceInfo`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *SyncModelSyncJobItem) GetMediaSourceOk() (*MediaSourceInfo, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *SyncModelSyncJobItem) SetMediaSource(v MediaSourceInfo)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *SyncModelSyncJobItem) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetTargetId

`func (o *SyncModelSyncJobItem) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SyncModelSyncJobItem) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SyncModelSyncJobItem) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SyncModelSyncJobItem) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetOutputPath

`func (o *SyncModelSyncJobItem) GetOutputPath() string`

GetOutputPath returns the OutputPath field if non-nil, zero value otherwise.

### GetOutputPathOk

`func (o *SyncModelSyncJobItem) GetOutputPathOk() (*string, bool)`

GetOutputPathOk returns a tuple with the OutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPath

`func (o *SyncModelSyncJobItem) SetOutputPath(v string)`

SetOutputPath sets OutputPath field to given value.

### HasOutputPath

`func (o *SyncModelSyncJobItem) HasOutputPath() bool`

HasOutputPath returns a boolean if a field has been set.

### GetStatus

`func (o *SyncModelSyncJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncModelSyncJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncModelSyncJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncModelSyncJobItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *SyncModelSyncJobItem) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SyncModelSyncJobItem) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SyncModelSyncJobItem) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SyncModelSyncJobItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *SyncModelSyncJobItem) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *SyncModelSyncJobItem) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetDateCreated

`func (o *SyncModelSyncJobItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncModelSyncJobItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncModelSyncJobItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncModelSyncJobItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *SyncModelSyncJobItem) GetPrimaryImageItemId() int64`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *SyncModelSyncJobItem) GetPrimaryImageItemIdOk() (*int64, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *SyncModelSyncJobItem) SetPrimaryImageItemId(v int64)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *SyncModelSyncJobItem) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *SyncModelSyncJobItem) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *SyncModelSyncJobItem) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *SyncModelSyncJobItem) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *SyncModelSyncJobItem) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetTemporaryPath

`func (o *SyncModelSyncJobItem) GetTemporaryPath() string`

GetTemporaryPath returns the TemporaryPath field if non-nil, zero value otherwise.

### GetTemporaryPathOk

`func (o *SyncModelSyncJobItem) GetTemporaryPathOk() (*string, bool)`

GetTemporaryPathOk returns a tuple with the TemporaryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryPath

`func (o *SyncModelSyncJobItem) SetTemporaryPath(v string)`

SetTemporaryPath sets TemporaryPath field to given value.

### HasTemporaryPath

`func (o *SyncModelSyncJobItem) HasTemporaryPath() bool`

HasTemporaryPath returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *SyncModelSyncJobItem) GetAdditionalFiles() []SyncModelItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *SyncModelSyncJobItem) GetAdditionalFilesOk() (*[]SyncModelItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *SyncModelSyncJobItem) SetAdditionalFiles(v []SyncModelItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *SyncModelSyncJobItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.

### GetItemDateModifiedTicks

`func (o *SyncModelSyncJobItem) GetItemDateModifiedTicks() int64`

GetItemDateModifiedTicks returns the ItemDateModifiedTicks field if non-nil, zero value otherwise.

### GetItemDateModifiedTicksOk

`func (o *SyncModelSyncJobItem) GetItemDateModifiedTicksOk() (*int64, bool)`

GetItemDateModifiedTicksOk returns a tuple with the ItemDateModifiedTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDateModifiedTicks

`func (o *SyncModelSyncJobItem) SetItemDateModifiedTicks(v int64)`

SetItemDateModifiedTicks sets ItemDateModifiedTicks field to given value.

### HasItemDateModifiedTicks

`func (o *SyncModelSyncJobItem) HasItemDateModifiedTicks() bool`

HasItemDateModifiedTicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



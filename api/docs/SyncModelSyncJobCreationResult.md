# SyncModelSyncJobCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**SyncSyncJob**](SyncSyncJob.md) |  | [optional] 
**JobItems** | Pointer to [**[]SyncModelSyncJobItem**](SyncModelSyncJobItem.md) |  | [optional] 

## Methods

### NewSyncModelSyncJobCreationResult

`func NewSyncModelSyncJobCreationResult() *SyncModelSyncJobCreationResult`

NewSyncModelSyncJobCreationResult instantiates a new SyncModelSyncJobCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncModelSyncJobCreationResultWithDefaults

`func NewSyncModelSyncJobCreationResultWithDefaults() *SyncModelSyncJobCreationResult`

NewSyncModelSyncJobCreationResultWithDefaults instantiates a new SyncModelSyncJobCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *SyncModelSyncJobCreationResult) GetJob() SyncSyncJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SyncModelSyncJobCreationResult) GetJobOk() (*SyncSyncJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SyncModelSyncJobCreationResult) SetJob(v SyncSyncJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *SyncModelSyncJobCreationResult) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobItems

`func (o *SyncModelSyncJobCreationResult) GetJobItems() []SyncModelSyncJobItem`

GetJobItems returns the JobItems field if non-nil, zero value otherwise.

### GetJobItemsOk

`func (o *SyncModelSyncJobCreationResult) GetJobItemsOk() (*[]SyncModelSyncJobItem, bool)`

GetJobItemsOk returns a tuple with the JobItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobItems

`func (o *SyncModelSyncJobCreationResult) SetJobItems(v []SyncModelSyncJobItem)`

SetJobItems sets JobItems field to given value.

### HasJobItems

`func (o *SyncModelSyncJobCreationResult) HasJobItems() bool`

HasJobItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



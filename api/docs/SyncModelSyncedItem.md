# SyncModelSyncedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**SyncJobId** | Pointer to **int64** |  | [optional] 
**SyncJobName** | Pointer to **string** |  | [optional] 
**SyncJobDateCreated** | Pointer to **time.Time** |  | [optional] 
**SyncJobItemId** | Pointer to **int64** |  | [optional] 
**OriginalFileName** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]SyncModelItemFileInfo**](SyncModelItemFileInfo.md) |  | [optional] 

## Methods

### NewSyncModelSyncedItem

`func NewSyncModelSyncedItem() *SyncModelSyncedItem`

NewSyncModelSyncedItem instantiates a new SyncModelSyncedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncModelSyncedItemWithDefaults

`func NewSyncModelSyncedItemWithDefaults() *SyncModelSyncedItem`

NewSyncModelSyncedItemWithDefaults instantiates a new SyncModelSyncedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *SyncModelSyncedItem) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SyncModelSyncedItem) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SyncModelSyncedItem) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SyncModelSyncedItem) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetSyncJobId

`func (o *SyncModelSyncedItem) GetSyncJobId() int64`

GetSyncJobId returns the SyncJobId field if non-nil, zero value otherwise.

### GetSyncJobIdOk

`func (o *SyncModelSyncedItem) GetSyncJobIdOk() (*int64, bool)`

GetSyncJobIdOk returns a tuple with the SyncJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobId

`func (o *SyncModelSyncedItem) SetSyncJobId(v int64)`

SetSyncJobId sets SyncJobId field to given value.

### HasSyncJobId

`func (o *SyncModelSyncedItem) HasSyncJobId() bool`

HasSyncJobId returns a boolean if a field has been set.

### GetSyncJobName

`func (o *SyncModelSyncedItem) GetSyncJobName() string`

GetSyncJobName returns the SyncJobName field if non-nil, zero value otherwise.

### GetSyncJobNameOk

`func (o *SyncModelSyncedItem) GetSyncJobNameOk() (*string, bool)`

GetSyncJobNameOk returns a tuple with the SyncJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobName

`func (o *SyncModelSyncedItem) SetSyncJobName(v string)`

SetSyncJobName sets SyncJobName field to given value.

### HasSyncJobName

`func (o *SyncModelSyncedItem) HasSyncJobName() bool`

HasSyncJobName returns a boolean if a field has been set.

### GetSyncJobDateCreated

`func (o *SyncModelSyncedItem) GetSyncJobDateCreated() time.Time`

GetSyncJobDateCreated returns the SyncJobDateCreated field if non-nil, zero value otherwise.

### GetSyncJobDateCreatedOk

`func (o *SyncModelSyncedItem) GetSyncJobDateCreatedOk() (*time.Time, bool)`

GetSyncJobDateCreatedOk returns a tuple with the SyncJobDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobDateCreated

`func (o *SyncModelSyncedItem) SetSyncJobDateCreated(v time.Time)`

SetSyncJobDateCreated sets SyncJobDateCreated field to given value.

### HasSyncJobDateCreated

`func (o *SyncModelSyncedItem) HasSyncJobDateCreated() bool`

HasSyncJobDateCreated returns a boolean if a field has been set.

### GetSyncJobItemId

`func (o *SyncModelSyncedItem) GetSyncJobItemId() int64`

GetSyncJobItemId returns the SyncJobItemId field if non-nil, zero value otherwise.

### GetSyncJobItemIdOk

`func (o *SyncModelSyncedItem) GetSyncJobItemIdOk() (*int64, bool)`

GetSyncJobItemIdOk returns a tuple with the SyncJobItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobItemId

`func (o *SyncModelSyncedItem) SetSyncJobItemId(v int64)`

SetSyncJobItemId sets SyncJobItemId field to given value.

### HasSyncJobItemId

`func (o *SyncModelSyncedItem) HasSyncJobItemId() bool`

HasSyncJobItemId returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *SyncModelSyncedItem) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *SyncModelSyncedItem) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *SyncModelSyncedItem) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *SyncModelSyncedItem) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetItem

`func (o *SyncModelSyncedItem) GetItem() BaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *SyncModelSyncedItem) GetItemOk() (*BaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *SyncModelSyncedItem) SetItem(v BaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *SyncModelSyncedItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetUserId

`func (o *SyncModelSyncedItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncModelSyncedItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncModelSyncedItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SyncModelSyncedItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *SyncModelSyncedItem) GetAdditionalFiles() []SyncModelItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *SyncModelSyncedItem) GetAdditionalFilesOk() (*[]SyncModelItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *SyncModelSyncedItem) SetAdditionalFiles(v []SyncModelItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *SyncModelSyncedItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



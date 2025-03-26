# SyncSyncJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**UnwatchedOnly** | Pointer to **bool** |  | [optional] 
**SyncNewContent** | Pointer to **bool** |  | [optional] 
**ItemLimit** | Pointer to **NullableInt32** |  | [optional] 
**RequestedItemIds** | Pointer to **[]int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateLastModified** | Pointer to **time.Time** |  | [optional] 
**ItemCount** | Pointer to **int32** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 

## Methods

### NewSyncSyncJob

`func NewSyncSyncJob() *SyncSyncJob`

NewSyncSyncJob instantiates a new SyncSyncJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncSyncJobWithDefaults

`func NewSyncSyncJobWithDefaults() *SyncSyncJob`

NewSyncSyncJobWithDefaults instantiates a new SyncSyncJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncSyncJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncSyncJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncSyncJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SyncSyncJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetId

`func (o *SyncSyncJob) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SyncSyncJob) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SyncSyncJob) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SyncSyncJob) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *SyncSyncJob) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *SyncSyncJob) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *SyncSyncJob) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *SyncSyncJob) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetQuality

`func (o *SyncSyncJob) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *SyncSyncJob) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *SyncSyncJob) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *SyncSyncJob) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetBitrate

`func (o *SyncSyncJob) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *SyncSyncJob) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *SyncSyncJob) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *SyncSyncJob) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *SyncSyncJob) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *SyncSyncJob) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetProfile

`func (o *SyncSyncJob) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SyncSyncJob) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SyncSyncJob) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SyncSyncJob) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCategory

`func (o *SyncSyncJob) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SyncSyncJob) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SyncSyncJob) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SyncSyncJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetParentId

`func (o *SyncSyncJob) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SyncSyncJob) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SyncSyncJob) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SyncSyncJob) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProgress

`func (o *SyncSyncJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SyncSyncJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SyncSyncJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SyncSyncJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetName

`func (o *SyncSyncJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncSyncJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncSyncJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyncSyncJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SyncSyncJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncSyncJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncSyncJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncSyncJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *SyncSyncJob) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncSyncJob) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncSyncJob) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SyncSyncJob) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUnwatchedOnly

`func (o *SyncSyncJob) GetUnwatchedOnly() bool`

GetUnwatchedOnly returns the UnwatchedOnly field if non-nil, zero value otherwise.

### GetUnwatchedOnlyOk

`func (o *SyncSyncJob) GetUnwatchedOnlyOk() (*bool, bool)`

GetUnwatchedOnlyOk returns a tuple with the UnwatchedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwatchedOnly

`func (o *SyncSyncJob) SetUnwatchedOnly(v bool)`

SetUnwatchedOnly sets UnwatchedOnly field to given value.

### HasUnwatchedOnly

`func (o *SyncSyncJob) HasUnwatchedOnly() bool`

HasUnwatchedOnly returns a boolean if a field has been set.

### GetSyncNewContent

`func (o *SyncSyncJob) GetSyncNewContent() bool`

GetSyncNewContent returns the SyncNewContent field if non-nil, zero value otherwise.

### GetSyncNewContentOk

`func (o *SyncSyncJob) GetSyncNewContentOk() (*bool, bool)`

GetSyncNewContentOk returns a tuple with the SyncNewContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncNewContent

`func (o *SyncSyncJob) SetSyncNewContent(v bool)`

SetSyncNewContent sets SyncNewContent field to given value.

### HasSyncNewContent

`func (o *SyncSyncJob) HasSyncNewContent() bool`

HasSyncNewContent returns a boolean if a field has been set.

### GetItemLimit

`func (o *SyncSyncJob) GetItemLimit() int32`

GetItemLimit returns the ItemLimit field if non-nil, zero value otherwise.

### GetItemLimitOk

`func (o *SyncSyncJob) GetItemLimitOk() (*int32, bool)`

GetItemLimitOk returns a tuple with the ItemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLimit

`func (o *SyncSyncJob) SetItemLimit(v int32)`

SetItemLimit sets ItemLimit field to given value.

### HasItemLimit

`func (o *SyncSyncJob) HasItemLimit() bool`

HasItemLimit returns a boolean if a field has been set.

### SetItemLimitNil

`func (o *SyncSyncJob) SetItemLimitNil(b bool)`

 SetItemLimitNil sets the value for ItemLimit to be an explicit nil

### UnsetItemLimit
`func (o *SyncSyncJob) UnsetItemLimit()`

UnsetItemLimit ensures that no value is present for ItemLimit, not even an explicit nil
### GetRequestedItemIds

`func (o *SyncSyncJob) GetRequestedItemIds() []int64`

GetRequestedItemIds returns the RequestedItemIds field if non-nil, zero value otherwise.

### GetRequestedItemIdsOk

`func (o *SyncSyncJob) GetRequestedItemIdsOk() (*[]int64, bool)`

GetRequestedItemIdsOk returns a tuple with the RequestedItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItemIds

`func (o *SyncSyncJob) SetRequestedItemIds(v []int64)`

SetRequestedItemIds sets RequestedItemIds field to given value.

### HasRequestedItemIds

`func (o *SyncSyncJob) HasRequestedItemIds() bool`

HasRequestedItemIds returns a boolean if a field has been set.

### GetDateCreated

`func (o *SyncSyncJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncSyncJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncSyncJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncSyncJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateLastModified

`func (o *SyncSyncJob) GetDateLastModified() time.Time`

GetDateLastModified returns the DateLastModified field if non-nil, zero value otherwise.

### GetDateLastModifiedOk

`func (o *SyncSyncJob) GetDateLastModifiedOk() (*time.Time, bool)`

GetDateLastModifiedOk returns a tuple with the DateLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastModified

`func (o *SyncSyncJob) SetDateLastModified(v time.Time)`

SetDateLastModified sets DateLastModified field to given value.

### HasDateLastModified

`func (o *SyncSyncJob) HasDateLastModified() bool`

HasDateLastModified returns a boolean if a field has been set.

### GetItemCount

`func (o *SyncSyncJob) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *SyncSyncJob) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *SyncSyncJob) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *SyncSyncJob) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetParentName

`func (o *SyncSyncJob) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *SyncSyncJob) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *SyncSyncJob) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *SyncSyncJob) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *SyncSyncJob) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *SyncSyncJob) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *SyncSyncJob) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *SyncSyncJob) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *SyncSyncJob) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *SyncSyncJob) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *SyncSyncJob) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *SyncSyncJob) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



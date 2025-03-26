# SyncModelSyncJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **string** |  | [optional] 
**ItemIds** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UnwatchedOnly** | Pointer to **bool** |  | [optional] 
**SyncNewContent** | Pointer to **bool** |  | [optional] 
**ItemLimit** | Pointer to **NullableInt32** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSyncModelSyncJobRequest

`func NewSyncModelSyncJobRequest() *SyncModelSyncJobRequest`

NewSyncModelSyncJobRequest instantiates a new SyncModelSyncJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncModelSyncJobRequestWithDefaults

`func NewSyncModelSyncJobRequestWithDefaults() *SyncModelSyncJobRequest`

NewSyncModelSyncJobRequestWithDefaults instantiates a new SyncModelSyncJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *SyncModelSyncJobRequest) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SyncModelSyncJobRequest) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SyncModelSyncJobRequest) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SyncModelSyncJobRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetItemIds

`func (o *SyncModelSyncJobRequest) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *SyncModelSyncJobRequest) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *SyncModelSyncJobRequest) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *SyncModelSyncJobRequest) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetCategory

`func (o *SyncModelSyncJobRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SyncModelSyncJobRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SyncModelSyncJobRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SyncModelSyncJobRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetParentId

`func (o *SyncModelSyncJobRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SyncModelSyncJobRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SyncModelSyncJobRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SyncModelSyncJobRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetQuality

`func (o *SyncModelSyncJobRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *SyncModelSyncJobRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *SyncModelSyncJobRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *SyncModelSyncJobRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetProfile

`func (o *SyncModelSyncJobRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SyncModelSyncJobRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SyncModelSyncJobRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SyncModelSyncJobRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetName

`func (o *SyncModelSyncJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncModelSyncJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncModelSyncJobRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyncModelSyncJobRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *SyncModelSyncJobRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncModelSyncJobRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncModelSyncJobRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SyncModelSyncJobRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUnwatchedOnly

`func (o *SyncModelSyncJobRequest) GetUnwatchedOnly() bool`

GetUnwatchedOnly returns the UnwatchedOnly field if non-nil, zero value otherwise.

### GetUnwatchedOnlyOk

`func (o *SyncModelSyncJobRequest) GetUnwatchedOnlyOk() (*bool, bool)`

GetUnwatchedOnlyOk returns a tuple with the UnwatchedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwatchedOnly

`func (o *SyncModelSyncJobRequest) SetUnwatchedOnly(v bool)`

SetUnwatchedOnly sets UnwatchedOnly field to given value.

### HasUnwatchedOnly

`func (o *SyncModelSyncJobRequest) HasUnwatchedOnly() bool`

HasUnwatchedOnly returns a boolean if a field has been set.

### GetSyncNewContent

`func (o *SyncModelSyncJobRequest) GetSyncNewContent() bool`

GetSyncNewContent returns the SyncNewContent field if non-nil, zero value otherwise.

### GetSyncNewContentOk

`func (o *SyncModelSyncJobRequest) GetSyncNewContentOk() (*bool, bool)`

GetSyncNewContentOk returns a tuple with the SyncNewContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncNewContent

`func (o *SyncModelSyncJobRequest) SetSyncNewContent(v bool)`

SetSyncNewContent sets SyncNewContent field to given value.

### HasSyncNewContent

`func (o *SyncModelSyncJobRequest) HasSyncNewContent() bool`

HasSyncNewContent returns a boolean if a field has been set.

### GetItemLimit

`func (o *SyncModelSyncJobRequest) GetItemLimit() int32`

GetItemLimit returns the ItemLimit field if non-nil, zero value otherwise.

### GetItemLimitOk

`func (o *SyncModelSyncJobRequest) GetItemLimitOk() (*int32, bool)`

GetItemLimitOk returns a tuple with the ItemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLimit

`func (o *SyncModelSyncJobRequest) SetItemLimit(v int32)`

SetItemLimit sets ItemLimit field to given value.

### HasItemLimit

`func (o *SyncModelSyncJobRequest) HasItemLimit() bool`

HasItemLimit returns a boolean if a field has been set.

### SetItemLimitNil

`func (o *SyncModelSyncJobRequest) SetItemLimitNil(b bool)`

 SetItemLimitNil sets the value for ItemLimit to be an explicit nil

### UnsetItemLimit
`func (o *SyncModelSyncJobRequest) UnsetItemLimit()`

UnsetItemLimit ensures that no value is present for ItemLimit, not even an explicit nil
### GetBitrate

`func (o *SyncModelSyncJobRequest) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *SyncModelSyncJobRequest) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *SyncModelSyncJobRequest) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *SyncModelSyncJobRequest) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *SyncModelSyncJobRequest) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *SyncModelSyncJobRequest) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



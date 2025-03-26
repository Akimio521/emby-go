# RemoteSubtitleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeLetterISOLanguageName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**CommunityRating** | Pointer to **NullableFloat32** |  | [optional] 
**DownloadCount** | Pointer to **NullableInt32** |  | [optional] 
**IsHashMatch** | Pointer to **NullableBool** |  | [optional] 
**IsForced** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewRemoteSubtitleInfo

`func NewRemoteSubtitleInfo() *RemoteSubtitleInfo`

NewRemoteSubtitleInfo instantiates a new RemoteSubtitleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSubtitleInfoWithDefaults

`func NewRemoteSubtitleInfoWithDefaults() *RemoteSubtitleInfo`

NewRemoteSubtitleInfoWithDefaults instantiates a new RemoteSubtitleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageName() string`

GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field if non-nil, zero value otherwise.

### GetThreeLetterISOLanguageNameOk

`func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageNameOk() (*string, bool)`

GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) SetThreeLetterISOLanguageName(v string)`

SetThreeLetterISOLanguageName sets ThreeLetterISOLanguageName field to given value.

### HasThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) HasThreeLetterISOLanguageName() bool`

HasThreeLetterISOLanguageName returns a boolean if a field has been set.

### GetId

`func (o *RemoteSubtitleInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteSubtitleInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteSubtitleInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteSubtitleInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *RemoteSubtitleInfo) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *RemoteSubtitleInfo) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *RemoteSubtitleInfo) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *RemoteSubtitleInfo) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetName

`func (o *RemoteSubtitleInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteSubtitleInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteSubtitleInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteSubtitleInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *RemoteSubtitleInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RemoteSubtitleInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RemoteSubtitleInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RemoteSubtitleInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetAuthor

`func (o *RemoteSubtitleInfo) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RemoteSubtitleInfo) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RemoteSubtitleInfo) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RemoteSubtitleInfo) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComment

`func (o *RemoteSubtitleInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemoteSubtitleInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemoteSubtitleInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RemoteSubtitleInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDateCreated

`func (o *RemoteSubtitleInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RemoteSubtitleInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RemoteSubtitleInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RemoteSubtitleInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *RemoteSubtitleInfo) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *RemoteSubtitleInfo) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCommunityRating

`func (o *RemoteSubtitleInfo) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *RemoteSubtitleInfo) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *RemoteSubtitleInfo) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *RemoteSubtitleInfo) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *RemoteSubtitleInfo) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *RemoteSubtitleInfo) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetDownloadCount

`func (o *RemoteSubtitleInfo) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *RemoteSubtitleInfo) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *RemoteSubtitleInfo) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *RemoteSubtitleInfo) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### SetDownloadCountNil

`func (o *RemoteSubtitleInfo) SetDownloadCountNil(b bool)`

 SetDownloadCountNil sets the value for DownloadCount to be an explicit nil

### UnsetDownloadCount
`func (o *RemoteSubtitleInfo) UnsetDownloadCount()`

UnsetDownloadCount ensures that no value is present for DownloadCount, not even an explicit nil
### GetIsHashMatch

`func (o *RemoteSubtitleInfo) GetIsHashMatch() bool`

GetIsHashMatch returns the IsHashMatch field if non-nil, zero value otherwise.

### GetIsHashMatchOk

`func (o *RemoteSubtitleInfo) GetIsHashMatchOk() (*bool, bool)`

GetIsHashMatchOk returns a tuple with the IsHashMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHashMatch

`func (o *RemoteSubtitleInfo) SetIsHashMatch(v bool)`

SetIsHashMatch sets IsHashMatch field to given value.

### HasIsHashMatch

`func (o *RemoteSubtitleInfo) HasIsHashMatch() bool`

HasIsHashMatch returns a boolean if a field has been set.

### SetIsHashMatchNil

`func (o *RemoteSubtitleInfo) SetIsHashMatchNil(b bool)`

 SetIsHashMatchNil sets the value for IsHashMatch to be an explicit nil

### UnsetIsHashMatch
`func (o *RemoteSubtitleInfo) UnsetIsHashMatch()`

UnsetIsHashMatch ensures that no value is present for IsHashMatch, not even an explicit nil
### GetIsForced

`func (o *RemoteSubtitleInfo) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *RemoteSubtitleInfo) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *RemoteSubtitleInfo) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.

### HasIsForced

`func (o *RemoteSubtitleInfo) HasIsForced() bool`

HasIsForced returns a boolean if a field has been set.

### SetIsForcedNil

`func (o *RemoteSubtitleInfo) SetIsForcedNil(b bool)`

 SetIsForcedNil sets the value for IsForced to be an explicit nil

### UnsetIsForced
`func (o *RemoteSubtitleInfo) UnsetIsForced()`

UnsetIsForced ensures that no value is present for IsForced, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



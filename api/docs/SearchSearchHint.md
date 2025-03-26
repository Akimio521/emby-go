# SearchSearchHint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MatchedTerm** | Pointer to **string** |  | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**ProductionYear** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**ThumbImageTag** | Pointer to **string** |  | [optional] 
**ThumbImageItemId** | Pointer to **string** |  | [optional] 
**BackdropImageTag** | Pointer to **string** |  | [optional] 
**BackdropImageItemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsFolder** | Pointer to **NullableBool** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Series** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Album** | Pointer to **string** |  | [optional] 
**AlbumId** | Pointer to **int64** |  | [optional] 
**AlbumArtist** | Pointer to **string** |  | [optional] 
**Artists** | Pointer to **[]string** |  | [optional] 
**SongCount** | Pointer to **NullableInt32** |  | [optional] 
**EpisodeCount** | Pointer to **NullableInt32** |  | [optional] 
**ChannelName** | Pointer to **string** |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewSearchSearchHint

`func NewSearchSearchHint() *SearchSearchHint`

NewSearchSearchHint instantiates a new SearchSearchHint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchHintWithDefaults

`func NewSearchSearchHintWithDefaults() *SearchSearchHint`

NewSearchSearchHintWithDefaults instantiates a new SearchSearchHint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *SearchSearchHint) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SearchSearchHint) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SearchSearchHint) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SearchSearchHint) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetId

`func (o *SearchSearchHint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchSearchHint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchSearchHint) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SearchSearchHint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchSearchHint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchSearchHint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchSearchHint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchSearchHint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatchedTerm

`func (o *SearchSearchHint) GetMatchedTerm() string`

GetMatchedTerm returns the MatchedTerm field if non-nil, zero value otherwise.

### GetMatchedTermOk

`func (o *SearchSearchHint) GetMatchedTermOk() (*string, bool)`

GetMatchedTermOk returns a tuple with the MatchedTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTerm

`func (o *SearchSearchHint) SetMatchedTerm(v string)`

SetMatchedTerm sets MatchedTerm field to given value.

### HasMatchedTerm

`func (o *SearchSearchHint) HasMatchedTerm() bool`

HasMatchedTerm returns a boolean if a field has been set.

### GetIndexNumber

`func (o *SearchSearchHint) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *SearchSearchHint) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *SearchSearchHint) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *SearchSearchHint) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *SearchSearchHint) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *SearchSearchHint) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetProductionYear

`func (o *SearchSearchHint) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *SearchSearchHint) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *SearchSearchHint) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *SearchSearchHint) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *SearchSearchHint) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *SearchSearchHint) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetParentIndexNumber

`func (o *SearchSearchHint) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *SearchSearchHint) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *SearchSearchHint) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *SearchSearchHint) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *SearchSearchHint) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *SearchSearchHint) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPrimaryImageTag

`func (o *SearchSearchHint) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *SearchSearchHint) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *SearchSearchHint) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *SearchSearchHint) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetThumbImageTag

`func (o *SearchSearchHint) GetThumbImageTag() string`

GetThumbImageTag returns the ThumbImageTag field if non-nil, zero value otherwise.

### GetThumbImageTagOk

`func (o *SearchSearchHint) GetThumbImageTagOk() (*string, bool)`

GetThumbImageTagOk returns a tuple with the ThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImageTag

`func (o *SearchSearchHint) SetThumbImageTag(v string)`

SetThumbImageTag sets ThumbImageTag field to given value.

### HasThumbImageTag

`func (o *SearchSearchHint) HasThumbImageTag() bool`

HasThumbImageTag returns a boolean if a field has been set.

### GetThumbImageItemId

`func (o *SearchSearchHint) GetThumbImageItemId() string`

GetThumbImageItemId returns the ThumbImageItemId field if non-nil, zero value otherwise.

### GetThumbImageItemIdOk

`func (o *SearchSearchHint) GetThumbImageItemIdOk() (*string, bool)`

GetThumbImageItemIdOk returns a tuple with the ThumbImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImageItemId

`func (o *SearchSearchHint) SetThumbImageItemId(v string)`

SetThumbImageItemId sets ThumbImageItemId field to given value.

### HasThumbImageItemId

`func (o *SearchSearchHint) HasThumbImageItemId() bool`

HasThumbImageItemId returns a boolean if a field has been set.

### GetBackdropImageTag

`func (o *SearchSearchHint) GetBackdropImageTag() string`

GetBackdropImageTag returns the BackdropImageTag field if non-nil, zero value otherwise.

### GetBackdropImageTagOk

`func (o *SearchSearchHint) GetBackdropImageTagOk() (*string, bool)`

GetBackdropImageTagOk returns a tuple with the BackdropImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageTag

`func (o *SearchSearchHint) SetBackdropImageTag(v string)`

SetBackdropImageTag sets BackdropImageTag field to given value.

### HasBackdropImageTag

`func (o *SearchSearchHint) HasBackdropImageTag() bool`

HasBackdropImageTag returns a boolean if a field has been set.

### GetBackdropImageItemId

`func (o *SearchSearchHint) GetBackdropImageItemId() string`

GetBackdropImageItemId returns the BackdropImageItemId field if non-nil, zero value otherwise.

### GetBackdropImageItemIdOk

`func (o *SearchSearchHint) GetBackdropImageItemIdOk() (*string, bool)`

GetBackdropImageItemIdOk returns a tuple with the BackdropImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageItemId

`func (o *SearchSearchHint) SetBackdropImageItemId(v string)`

SetBackdropImageItemId sets BackdropImageItemId field to given value.

### HasBackdropImageItemId

`func (o *SearchSearchHint) HasBackdropImageItemId() bool`

HasBackdropImageItemId returns a boolean if a field has been set.

### GetType

`func (o *SearchSearchHint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchSearchHint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchSearchHint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchSearchHint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsFolder

`func (o *SearchSearchHint) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SearchSearchHint) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SearchSearchHint) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SearchSearchHint) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *SearchSearchHint) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *SearchSearchHint) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetRunTimeTicks

`func (o *SearchSearchHint) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *SearchSearchHint) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *SearchSearchHint) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *SearchSearchHint) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *SearchSearchHint) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *SearchSearchHint) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetMediaType

`func (o *SearchSearchHint) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SearchSearchHint) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SearchSearchHint) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *SearchSearchHint) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetStartDate

`func (o *SearchSearchHint) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SearchSearchHint) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SearchSearchHint) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SearchSearchHint) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *SearchSearchHint) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *SearchSearchHint) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *SearchSearchHint) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SearchSearchHint) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SearchSearchHint) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SearchSearchHint) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *SearchSearchHint) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *SearchSearchHint) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetSeries

`func (o *SearchSearchHint) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SearchSearchHint) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SearchSearchHint) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *SearchSearchHint) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStatus

`func (o *SearchSearchHint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchSearchHint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchSearchHint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchSearchHint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlbum

`func (o *SearchSearchHint) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SearchSearchHint) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SearchSearchHint) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SearchSearchHint) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetAlbumId

`func (o *SearchSearchHint) GetAlbumId() int64`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *SearchSearchHint) GetAlbumIdOk() (*int64, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *SearchSearchHint) SetAlbumId(v int64)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *SearchSearchHint) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetAlbumArtist

`func (o *SearchSearchHint) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *SearchSearchHint) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *SearchSearchHint) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *SearchSearchHint) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### GetArtists

`func (o *SearchSearchHint) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SearchSearchHint) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SearchSearchHint) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *SearchSearchHint) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetSongCount

`func (o *SearchSearchHint) GetSongCount() int32`

GetSongCount returns the SongCount field if non-nil, zero value otherwise.

### GetSongCountOk

`func (o *SearchSearchHint) GetSongCountOk() (*int32, bool)`

GetSongCountOk returns a tuple with the SongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongCount

`func (o *SearchSearchHint) SetSongCount(v int32)`

SetSongCount sets SongCount field to given value.

### HasSongCount

`func (o *SearchSearchHint) HasSongCount() bool`

HasSongCount returns a boolean if a field has been set.

### SetSongCountNil

`func (o *SearchSearchHint) SetSongCountNil(b bool)`

 SetSongCountNil sets the value for SongCount to be an explicit nil

### UnsetSongCount
`func (o *SearchSearchHint) UnsetSongCount()`

UnsetSongCount ensures that no value is present for SongCount, not even an explicit nil
### GetEpisodeCount

`func (o *SearchSearchHint) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *SearchSearchHint) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *SearchSearchHint) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *SearchSearchHint) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### SetEpisodeCountNil

`func (o *SearchSearchHint) SetEpisodeCountNil(b bool)`

 SetEpisodeCountNil sets the value for EpisodeCount to be an explicit nil

### UnsetEpisodeCount
`func (o *SearchSearchHint) UnsetEpisodeCount()`

UnsetEpisodeCount ensures that no value is present for EpisodeCount, not even an explicit nil
### GetChannelName

`func (o *SearchSearchHint) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *SearchSearchHint) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *SearchSearchHint) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *SearchSearchHint) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetPrimaryImageAspectRatio

`func (o *SearchSearchHint) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *SearchSearchHint) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *SearchSearchHint) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *SearchSearchHint) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *SearchSearchHint) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *SearchSearchHint) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



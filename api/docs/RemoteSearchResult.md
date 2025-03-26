# RemoteSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ProviderIds** | Pointer to **map[string]string** |  | [optional] 
**ProductionYear** | Pointer to **NullableInt32** |  | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**IndexNumberEnd** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**PremiereDate** | Pointer to **NullableTime** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**SearchProviderName** | Pointer to **string** |  | [optional] 
**GameSystem** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**AlbumArtist** | Pointer to [**RemoteSearchResult**](RemoteSearchResult.md) |  | [optional] 
**Artists** | Pointer to [**[]RemoteSearchResult**](RemoteSearchResult.md) |  | [optional] 

## Methods

### NewRemoteSearchResult

`func NewRemoteSearchResult() *RemoteSearchResult`

NewRemoteSearchResult instantiates a new RemoteSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSearchResultWithDefaults

`func NewRemoteSearchResultWithDefaults() *RemoteSearchResult`

NewRemoteSearchResultWithDefaults instantiates a new RemoteSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RemoteSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteSearchResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteSearchResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderIds

`func (o *RemoteSearchResult) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *RemoteSearchResult) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *RemoteSearchResult) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *RemoteSearchResult) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetProductionYear

`func (o *RemoteSearchResult) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *RemoteSearchResult) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *RemoteSearchResult) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *RemoteSearchResult) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *RemoteSearchResult) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *RemoteSearchResult) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetIndexNumber

`func (o *RemoteSearchResult) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *RemoteSearchResult) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *RemoteSearchResult) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *RemoteSearchResult) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *RemoteSearchResult) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *RemoteSearchResult) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetIndexNumberEnd

`func (o *RemoteSearchResult) GetIndexNumberEnd() int32`

GetIndexNumberEnd returns the IndexNumberEnd field if non-nil, zero value otherwise.

### GetIndexNumberEndOk

`func (o *RemoteSearchResult) GetIndexNumberEndOk() (*int32, bool)`

GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumberEnd

`func (o *RemoteSearchResult) SetIndexNumberEnd(v int32)`

SetIndexNumberEnd sets IndexNumberEnd field to given value.

### HasIndexNumberEnd

`func (o *RemoteSearchResult) HasIndexNumberEnd() bool`

HasIndexNumberEnd returns a boolean if a field has been set.

### SetIndexNumberEndNil

`func (o *RemoteSearchResult) SetIndexNumberEndNil(b bool)`

 SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil

### UnsetIndexNumberEnd
`func (o *RemoteSearchResult) UnsetIndexNumberEnd()`

UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
### GetParentIndexNumber

`func (o *RemoteSearchResult) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *RemoteSearchResult) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *RemoteSearchResult) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *RemoteSearchResult) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *RemoteSearchResult) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *RemoteSearchResult) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *RemoteSearchResult) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *RemoteSearchResult) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *RemoteSearchResult) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *RemoteSearchResult) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *RemoteSearchResult) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *RemoteSearchResult) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetImageUrl

`func (o *RemoteSearchResult) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *RemoteSearchResult) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *RemoteSearchResult) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *RemoteSearchResult) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetSearchProviderName

`func (o *RemoteSearchResult) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *RemoteSearchResult) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *RemoteSearchResult) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *RemoteSearchResult) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### GetGameSystem

`func (o *RemoteSearchResult) GetGameSystem() string`

GetGameSystem returns the GameSystem field if non-nil, zero value otherwise.

### GetGameSystemOk

`func (o *RemoteSearchResult) GetGameSystemOk() (*string, bool)`

GetGameSystemOk returns a tuple with the GameSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameSystem

`func (o *RemoteSearchResult) SetGameSystem(v string)`

SetGameSystem sets GameSystem field to given value.

### HasGameSystem

`func (o *RemoteSearchResult) HasGameSystem() bool`

HasGameSystem returns a boolean if a field has been set.

### GetOverview

`func (o *RemoteSearchResult) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *RemoteSearchResult) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *RemoteSearchResult) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *RemoteSearchResult) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetAlbumArtist

`func (o *RemoteSearchResult) GetAlbumArtist() RemoteSearchResult`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *RemoteSearchResult) GetAlbumArtistOk() (*RemoteSearchResult, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *RemoteSearchResult) SetAlbumArtist(v RemoteSearchResult)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *RemoteSearchResult) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### GetArtists

`func (o *RemoteSearchResult) GetArtists() []RemoteSearchResult`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *RemoteSearchResult) GetArtistsOk() (*[]RemoteSearchResult, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *RemoteSearchResult) SetArtists(v []RemoteSearchResult)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *RemoteSearchResult) HasArtists() bool`

HasArtists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



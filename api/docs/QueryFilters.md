# QueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genres** | Pointer to [**[]NameLongIdPair**](NameLongIdPair.md) |  | [optional] 
**Studios** | Pointer to [**[]NameLongIdPair**](NameLongIdPair.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQueryFilters

`func NewQueryFilters() *QueryFilters`

NewQueryFilters instantiates a new QueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFiltersWithDefaults

`func NewQueryFiltersWithDefaults() *QueryFilters`

NewQueryFiltersWithDefaults instantiates a new QueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenres

`func (o *QueryFilters) GetGenres() []NameLongIdPair`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *QueryFilters) GetGenresOk() (*[]NameLongIdPair, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *QueryFilters) SetGenres(v []NameLongIdPair)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *QueryFilters) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetStudios

`func (o *QueryFilters) GetStudios() []NameLongIdPair`

GetStudios returns the Studios field if non-nil, zero value otherwise.

### GetStudiosOk

`func (o *QueryFilters) GetStudiosOk() (*[]NameLongIdPair, bool)`

GetStudiosOk returns a tuple with the Studios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudios

`func (o *QueryFilters) SetStudios(v []NameLongIdPair)`

SetStudios sets Studios field to given value.

### HasStudios

`func (o *QueryFilters) HasStudios() bool`

HasStudios returns a boolean if a field has been set.

### GetTags

`func (o *QueryFilters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryFilters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryFilters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



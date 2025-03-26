# VirtualFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to **[]string** |  | [optional] 
**CollectionType** | Pointer to **string** |  | [optional] 
**LibraryOptions** | Pointer to [**ConfigurationLibraryOptions**](ConfigurationLibraryOptions.md) |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**RefreshProgress** | Pointer to **NullableFloat64** |  | [optional] 
**RefreshStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualFolderInfo

`func NewVirtualFolderInfo() *VirtualFolderInfo`

NewVirtualFolderInfo instantiates a new VirtualFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualFolderInfoWithDefaults

`func NewVirtualFolderInfoWithDefaults() *VirtualFolderInfo`

NewVirtualFolderInfoWithDefaults instantiates a new VirtualFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualFolderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualFolderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualFolderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualFolderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocations

`func (o *VirtualFolderInfo) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *VirtualFolderInfo) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *VirtualFolderInfo) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *VirtualFolderInfo) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetCollectionType

`func (o *VirtualFolderInfo) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *VirtualFolderInfo) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *VirtualFolderInfo) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *VirtualFolderInfo) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetLibraryOptions

`func (o *VirtualFolderInfo) GetLibraryOptions() ConfigurationLibraryOptions`

GetLibraryOptions returns the LibraryOptions field if non-nil, zero value otherwise.

### GetLibraryOptionsOk

`func (o *VirtualFolderInfo) GetLibraryOptionsOk() (*ConfigurationLibraryOptions, bool)`

GetLibraryOptionsOk returns a tuple with the LibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryOptions

`func (o *VirtualFolderInfo) SetLibraryOptions(v ConfigurationLibraryOptions)`

SetLibraryOptions sets LibraryOptions field to given value.

### HasLibraryOptions

`func (o *VirtualFolderInfo) HasLibraryOptions() bool`

HasLibraryOptions returns a boolean if a field has been set.

### GetItemId

`func (o *VirtualFolderInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VirtualFolderInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VirtualFolderInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VirtualFolderInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *VirtualFolderInfo) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *VirtualFolderInfo) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *VirtualFolderInfo) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *VirtualFolderInfo) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetRefreshProgress

`func (o *VirtualFolderInfo) GetRefreshProgress() float64`

GetRefreshProgress returns the RefreshProgress field if non-nil, zero value otherwise.

### GetRefreshProgressOk

`func (o *VirtualFolderInfo) GetRefreshProgressOk() (*float64, bool)`

GetRefreshProgressOk returns a tuple with the RefreshProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshProgress

`func (o *VirtualFolderInfo) SetRefreshProgress(v float64)`

SetRefreshProgress sets RefreshProgress field to given value.

### HasRefreshProgress

`func (o *VirtualFolderInfo) HasRefreshProgress() bool`

HasRefreshProgress returns a boolean if a field has been set.

### SetRefreshProgressNil

`func (o *VirtualFolderInfo) SetRefreshProgressNil(b bool)`

 SetRefreshProgressNil sets the value for RefreshProgress to be an explicit nil

### UnsetRefreshProgress
`func (o *VirtualFolderInfo) UnsetRefreshProgress()`

UnsetRefreshProgress ensures that no value is present for RefreshProgress, not even an explicit nil
### GetRefreshStatus

`func (o *VirtualFolderInfo) GetRefreshStatus() string`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *VirtualFolderInfo) GetRefreshStatusOk() (*string, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *VirtualFolderInfo) SetRefreshStatus(v string)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *VirtualFolderInfo) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



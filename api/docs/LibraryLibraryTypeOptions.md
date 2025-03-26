# LibraryLibraryTypeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MetadataFetchers** | Pointer to [**[]LibraryLibraryOptionInfo**](LibraryLibraryOptionInfo.md) |  | [optional] 
**ImageFetchers** | Pointer to [**[]LibraryLibraryOptionInfo**](LibraryLibraryOptionInfo.md) |  | [optional] 
**SupportedImageTypes** | Pointer to **[]string** |  | [optional] 
**DefaultImageOptions** | Pointer to [**[]ConfigurationImageOption**](ConfigurationImageOption.md) |  | [optional] 

## Methods

### NewLibraryLibraryTypeOptions

`func NewLibraryLibraryTypeOptions() *LibraryLibraryTypeOptions`

NewLibraryLibraryTypeOptions instantiates a new LibraryLibraryTypeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryLibraryTypeOptionsWithDefaults

`func NewLibraryLibraryTypeOptionsWithDefaults() *LibraryLibraryTypeOptions`

NewLibraryLibraryTypeOptionsWithDefaults instantiates a new LibraryLibraryTypeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LibraryLibraryTypeOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LibraryLibraryTypeOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LibraryLibraryTypeOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LibraryLibraryTypeOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadataFetchers

`func (o *LibraryLibraryTypeOptions) GetMetadataFetchers() []LibraryLibraryOptionInfo`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *LibraryLibraryTypeOptions) GetMetadataFetchersOk() (*[]LibraryLibraryOptionInfo, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *LibraryLibraryTypeOptions) SetMetadataFetchers(v []LibraryLibraryOptionInfo)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *LibraryLibraryTypeOptions) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetImageFetchers

`func (o *LibraryLibraryTypeOptions) GetImageFetchers() []LibraryLibraryOptionInfo`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *LibraryLibraryTypeOptions) GetImageFetchersOk() (*[]LibraryLibraryOptionInfo, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *LibraryLibraryTypeOptions) SetImageFetchers(v []LibraryLibraryOptionInfo)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *LibraryLibraryTypeOptions) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetSupportedImageTypes

`func (o *LibraryLibraryTypeOptions) GetSupportedImageTypes() []string`

GetSupportedImageTypes returns the SupportedImageTypes field if non-nil, zero value otherwise.

### GetSupportedImageTypesOk

`func (o *LibraryLibraryTypeOptions) GetSupportedImageTypesOk() (*[]string, bool)`

GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImageTypes

`func (o *LibraryLibraryTypeOptions) SetSupportedImageTypes(v []string)`

SetSupportedImageTypes sets SupportedImageTypes field to given value.

### HasSupportedImageTypes

`func (o *LibraryLibraryTypeOptions) HasSupportedImageTypes() bool`

HasSupportedImageTypes returns a boolean if a field has been set.

### GetDefaultImageOptions

`func (o *LibraryLibraryTypeOptions) GetDefaultImageOptions() []ConfigurationImageOption`

GetDefaultImageOptions returns the DefaultImageOptions field if non-nil, zero value otherwise.

### GetDefaultImageOptionsOk

`func (o *LibraryLibraryTypeOptions) GetDefaultImageOptionsOk() (*[]ConfigurationImageOption, bool)`

GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageOptions

`func (o *LibraryLibraryTypeOptions) SetDefaultImageOptions(v []ConfigurationImageOption)`

SetDefaultImageOptions sets DefaultImageOptions field to given value.

### HasDefaultImageOptions

`func (o *LibraryLibraryTypeOptions) HasDefaultImageOptions() bool`

HasDefaultImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



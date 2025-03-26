# ConfigurationTypeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MetadataFetchers** | Pointer to **[]string** |  | [optional] 
**MetadataFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageFetchers** | Pointer to **[]string** |  | [optional] 
**ImageFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageOptions** | Pointer to [**[]ConfigurationImageOption**](ConfigurationImageOption.md) |  | [optional] 

## Methods

### NewConfigurationTypeOptions

`func NewConfigurationTypeOptions() *ConfigurationTypeOptions`

NewConfigurationTypeOptions instantiates a new ConfigurationTypeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeOptionsWithDefaults

`func NewConfigurationTypeOptionsWithDefaults() *ConfigurationTypeOptions`

NewConfigurationTypeOptionsWithDefaults instantiates a new ConfigurationTypeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConfigurationTypeOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationTypeOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationTypeOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigurationTypeOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadataFetchers

`func (o *ConfigurationTypeOptions) GetMetadataFetchers() []string`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *ConfigurationTypeOptions) GetMetadataFetchersOk() (*[]string, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *ConfigurationTypeOptions) SetMetadataFetchers(v []string)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *ConfigurationTypeOptions) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetMetadataFetcherOrder

`func (o *ConfigurationTypeOptions) GetMetadataFetcherOrder() []string`

GetMetadataFetcherOrder returns the MetadataFetcherOrder field if non-nil, zero value otherwise.

### GetMetadataFetcherOrderOk

`func (o *ConfigurationTypeOptions) GetMetadataFetcherOrderOk() (*[]string, bool)`

GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetcherOrder

`func (o *ConfigurationTypeOptions) SetMetadataFetcherOrder(v []string)`

SetMetadataFetcherOrder sets MetadataFetcherOrder field to given value.

### HasMetadataFetcherOrder

`func (o *ConfigurationTypeOptions) HasMetadataFetcherOrder() bool`

HasMetadataFetcherOrder returns a boolean if a field has been set.

### GetImageFetchers

`func (o *ConfigurationTypeOptions) GetImageFetchers() []string`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *ConfigurationTypeOptions) GetImageFetchersOk() (*[]string, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *ConfigurationTypeOptions) SetImageFetchers(v []string)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *ConfigurationTypeOptions) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetImageFetcherOrder

`func (o *ConfigurationTypeOptions) GetImageFetcherOrder() []string`

GetImageFetcherOrder returns the ImageFetcherOrder field if non-nil, zero value otherwise.

### GetImageFetcherOrderOk

`func (o *ConfigurationTypeOptions) GetImageFetcherOrderOk() (*[]string, bool)`

GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetcherOrder

`func (o *ConfigurationTypeOptions) SetImageFetcherOrder(v []string)`

SetImageFetcherOrder sets ImageFetcherOrder field to given value.

### HasImageFetcherOrder

`func (o *ConfigurationTypeOptions) HasImageFetcherOrder() bool`

HasImageFetcherOrder returns a boolean if a field has been set.

### GetImageOptions

`func (o *ConfigurationTypeOptions) GetImageOptions() []ConfigurationImageOption`

GetImageOptions returns the ImageOptions field if non-nil, zero value otherwise.

### GetImageOptionsOk

`func (o *ConfigurationTypeOptions) GetImageOptionsOk() (*[]ConfigurationImageOption, bool)`

GetImageOptionsOk returns a tuple with the ImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptions

`func (o *ConfigurationTypeOptions) SetImageOptions(v []ConfigurationImageOption)`

SetImageOptions sets ImageOptions field to given value.

### HasImageOptions

`func (o *ConfigurationTypeOptions) HasImageOptions() bool`

HasImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ImageProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SupportedImages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewImageProviderInfo

`func NewImageProviderInfo() *ImageProviderInfo`

NewImageProviderInfo instantiates a new ImageProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageProviderInfoWithDefaults

`func NewImageProviderInfoWithDefaults() *ImageProviderInfo`

NewImageProviderInfoWithDefaults instantiates a new ImageProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageProviderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageProviderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageProviderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageProviderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSupportedImages

`func (o *ImageProviderInfo) GetSupportedImages() []string`

GetSupportedImages returns the SupportedImages field if non-nil, zero value otherwise.

### GetSupportedImagesOk

`func (o *ImageProviderInfo) GetSupportedImagesOk() (*[]string, bool)`

GetSupportedImagesOk returns a tuple with the SupportedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImages

`func (o *ImageProviderInfo) SetSupportedImages(v []string)`

SetSupportedImages sets SupportedImages field to given value.

### HasSupportedImages

`func (o *ImageProviderInfo) HasSupportedImages() bool`

HasSupportedImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



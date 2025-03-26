# DlnaCodecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]DlnaProfileCondition**](DlnaProfileCondition.md) |  | [optional] 
**ApplyConditions** | Pointer to [**[]DlnaProfileCondition**](DlnaProfileCondition.md) |  | [optional] 
**Codec** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 

## Methods

### NewDlnaCodecProfile

`func NewDlnaCodecProfile() *DlnaCodecProfile`

NewDlnaCodecProfile instantiates a new DlnaCodecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaCodecProfileWithDefaults

`func NewDlnaCodecProfileWithDefaults() *DlnaCodecProfile`

NewDlnaCodecProfileWithDefaults instantiates a new DlnaCodecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DlnaCodecProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DlnaCodecProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DlnaCodecProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DlnaCodecProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *DlnaCodecProfile) GetConditions() []DlnaProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DlnaCodecProfile) GetConditionsOk() (*[]DlnaProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DlnaCodecProfile) SetConditions(v []DlnaProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DlnaCodecProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetApplyConditions

`func (o *DlnaCodecProfile) GetApplyConditions() []DlnaProfileCondition`

GetApplyConditions returns the ApplyConditions field if non-nil, zero value otherwise.

### GetApplyConditionsOk

`func (o *DlnaCodecProfile) GetApplyConditionsOk() (*[]DlnaProfileCondition, bool)`

GetApplyConditionsOk returns a tuple with the ApplyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyConditions

`func (o *DlnaCodecProfile) SetApplyConditions(v []DlnaProfileCondition)`

SetApplyConditions sets ApplyConditions field to given value.

### HasApplyConditions

`func (o *DlnaCodecProfile) HasApplyConditions() bool`

HasApplyConditions returns a boolean if a field has been set.

### GetCodec

`func (o *DlnaCodecProfile) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *DlnaCodecProfile) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *DlnaCodecProfile) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *DlnaCodecProfile) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetContainer

`func (o *DlnaCodecProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DlnaCodecProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DlnaCodecProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DlnaCodecProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



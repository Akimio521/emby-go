# DlnaContainerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]DlnaProfileCondition**](DlnaProfileCondition.md) |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 

## Methods

### NewDlnaContainerProfile

`func NewDlnaContainerProfile() *DlnaContainerProfile`

NewDlnaContainerProfile instantiates a new DlnaContainerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaContainerProfileWithDefaults

`func NewDlnaContainerProfileWithDefaults() *DlnaContainerProfile`

NewDlnaContainerProfileWithDefaults instantiates a new DlnaContainerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DlnaContainerProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DlnaContainerProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DlnaContainerProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DlnaContainerProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *DlnaContainerProfile) GetConditions() []DlnaProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DlnaContainerProfile) GetConditionsOk() (*[]DlnaProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DlnaContainerProfile) SetConditions(v []DlnaProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DlnaContainerProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainer

`func (o *DlnaContainerProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DlnaContainerProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DlnaContainerProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DlnaContainerProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



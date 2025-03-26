# UsersForgotPasswordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**PinFile** | Pointer to **string** |  | [optional] 
**PinExpirationDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewUsersForgotPasswordResult

`func NewUsersForgotPasswordResult() *UsersForgotPasswordResult`

NewUsersForgotPasswordResult instantiates a new UsersForgotPasswordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersForgotPasswordResultWithDefaults

`func NewUsersForgotPasswordResultWithDefaults() *UsersForgotPasswordResult`

NewUsersForgotPasswordResultWithDefaults instantiates a new UsersForgotPasswordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UsersForgotPasswordResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UsersForgotPasswordResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UsersForgotPasswordResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UsersForgotPasswordResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPinFile

`func (o *UsersForgotPasswordResult) GetPinFile() string`

GetPinFile returns the PinFile field if non-nil, zero value otherwise.

### GetPinFileOk

`func (o *UsersForgotPasswordResult) GetPinFileOk() (*string, bool)`

GetPinFileOk returns a tuple with the PinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFile

`func (o *UsersForgotPasswordResult) SetPinFile(v string)`

SetPinFile sets PinFile field to given value.

### HasPinFile

`func (o *UsersForgotPasswordResult) HasPinFile() bool`

HasPinFile returns a boolean if a field has been set.

### GetPinExpirationDate

`func (o *UsersForgotPasswordResult) GetPinExpirationDate() time.Time`

GetPinExpirationDate returns the PinExpirationDate field if non-nil, zero value otherwise.

### GetPinExpirationDateOk

`func (o *UsersForgotPasswordResult) GetPinExpirationDateOk() (*time.Time, bool)`

GetPinExpirationDateOk returns a tuple with the PinExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinExpirationDate

`func (o *UsersForgotPasswordResult) SetPinExpirationDate(v time.Time)`

SetPinExpirationDate sets PinExpirationDate field to given value.

### HasPinExpirationDate

`func (o *UsersForgotPasswordResult) HasPinExpirationDate() bool`

HasPinExpirationDate returns a boolean if a field has been set.

### SetPinExpirationDateNil

`func (o *UsersForgotPasswordResult) SetPinExpirationDateNil(b bool)`

 SetPinExpirationDateNil sets the value for PinExpirationDate to be an explicit nil

### UnsetPinExpirationDate
`func (o *UsersForgotPasswordResult) UnsetPinExpirationDate()`

UnsetPinExpirationDate ensures that no value is present for PinExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



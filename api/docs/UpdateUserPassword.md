# UpdateUserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CurrentPw** | Pointer to **string** |  | [optional] 
**NewPw** | Pointer to **string** |  | [optional] 
**ResetPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateUserPassword

`func NewUpdateUserPassword() *UpdateUserPassword`

NewUpdateUserPassword instantiates a new UpdateUserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPasswordWithDefaults

`func NewUpdateUserPasswordWithDefaults() *UpdateUserPassword`

NewUpdateUserPasswordWithDefaults instantiates a new UpdateUserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateUserPassword) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserPassword) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserPassword) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUserPassword) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCurrentPw

`func (o *UpdateUserPassword) GetCurrentPw() string`

GetCurrentPw returns the CurrentPw field if non-nil, zero value otherwise.

### GetCurrentPwOk

`func (o *UpdateUserPassword) GetCurrentPwOk() (*string, bool)`

GetCurrentPwOk returns a tuple with the CurrentPw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPw

`func (o *UpdateUserPassword) SetCurrentPw(v string)`

SetCurrentPw sets CurrentPw field to given value.

### HasCurrentPw

`func (o *UpdateUserPassword) HasCurrentPw() bool`

HasCurrentPw returns a boolean if a field has been set.

### GetNewPw

`func (o *UpdateUserPassword) GetNewPw() string`

GetNewPw returns the NewPw field if non-nil, zero value otherwise.

### GetNewPwOk

`func (o *UpdateUserPassword) GetNewPwOk() (*string, bool)`

GetNewPwOk returns a tuple with the NewPw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPw

`func (o *UpdateUserPassword) SetNewPw(v string)`

SetNewPw sets NewPw field to given value.

### HasNewPw

`func (o *UpdateUserPassword) HasNewPw() bool`

HasNewPw returns a boolean if a field has been set.

### GetResetPassword

`func (o *UpdateUserPassword) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *UpdateUserPassword) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *UpdateUserPassword) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *UpdateUserPassword) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



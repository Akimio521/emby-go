# AuthenticationAuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 
**SessionInfo** | Pointer to [**SessionSessionInfo**](SessionSessionInfo.md) |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationAuthenticationResult

`func NewAuthenticationAuthenticationResult() *AuthenticationAuthenticationResult`

NewAuthenticationAuthenticationResult instantiates a new AuthenticationAuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationAuthenticationResultWithDefaults

`func NewAuthenticationAuthenticationResultWithDefaults() *AuthenticationAuthenticationResult`

NewAuthenticationAuthenticationResultWithDefaults instantiates a new AuthenticationAuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AuthenticationAuthenticationResult) GetUser() UserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticationAuthenticationResult) GetUserOk() (*UserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticationAuthenticationResult) SetUser(v UserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticationAuthenticationResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSessionInfo

`func (o *AuthenticationAuthenticationResult) GetSessionInfo() SessionSessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *AuthenticationAuthenticationResult) GetSessionInfoOk() (*SessionSessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *AuthenticationAuthenticationResult) SetSessionInfo(v SessionSessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *AuthenticationAuthenticationResult) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthenticationAuthenticationResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthenticationAuthenticationResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthenticationAuthenticationResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthenticationAuthenticationResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetServerId

`func (o *AuthenticationAuthenticationResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *AuthenticationAuthenticationResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *AuthenticationAuthenticationResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *AuthenticationAuthenticationResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



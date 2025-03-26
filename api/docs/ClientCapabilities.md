# ClientCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**SupportedCommands** | Pointer to **[]string** |  | [optional] 
**SupportsMediaControl** | Pointer to **bool** |  | [optional] 
**PushToken** | Pointer to **string** |  | [optional] 
**PushTokenType** | Pointer to **string** |  | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** |  | [optional] 
**SupportsSync** | Pointer to **bool** |  | [optional] 
**DeviceProfile** | Pointer to [**DlnaDeviceProfile**](DlnaDeviceProfile.md) |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewClientCapabilities

`func NewClientCapabilities() *ClientCapabilities`

NewClientCapabilities instantiates a new ClientCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesWithDefaults

`func NewClientCapabilitiesWithDefaults() *ClientCapabilities`

NewClientCapabilitiesWithDefaults instantiates a new ClientCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *ClientCapabilities) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ClientCapabilities) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ClientCapabilities) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ClientCapabilities) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *ClientCapabilities) GetSupportedCommands() []string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ClientCapabilities) GetSupportedCommandsOk() (*[]string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ClientCapabilities) SetSupportedCommands(v []string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ClientCapabilities) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *ClientCapabilities) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ClientCapabilities) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ClientCapabilities) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ClientCapabilities) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetPushToken

`func (o *ClientCapabilities) GetPushToken() string`

GetPushToken returns the PushToken field if non-nil, zero value otherwise.

### GetPushTokenOk

`func (o *ClientCapabilities) GetPushTokenOk() (*string, bool)`

GetPushTokenOk returns a tuple with the PushToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToken

`func (o *ClientCapabilities) SetPushToken(v string)`

SetPushToken sets PushToken field to given value.

### HasPushToken

`func (o *ClientCapabilities) HasPushToken() bool`

HasPushToken returns a boolean if a field has been set.

### GetPushTokenType

`func (o *ClientCapabilities) GetPushTokenType() string`

GetPushTokenType returns the PushTokenType field if non-nil, zero value otherwise.

### GetPushTokenTypeOk

`func (o *ClientCapabilities) GetPushTokenTypeOk() (*string, bool)`

GetPushTokenTypeOk returns a tuple with the PushTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTokenType

`func (o *ClientCapabilities) SetPushTokenType(v string)`

SetPushTokenType sets PushTokenType field to given value.

### HasPushTokenType

`func (o *ClientCapabilities) HasPushTokenType() bool`

HasPushTokenType returns a boolean if a field has been set.

### GetSupportsPersistentIdentifier

`func (o *ClientCapabilities) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *ClientCapabilities) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *ClientCapabilities) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *ClientCapabilities) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetSupportsSync

`func (o *ClientCapabilities) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *ClientCapabilities) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *ClientCapabilities) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *ClientCapabilities) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ClientCapabilities) GetDeviceProfile() DlnaDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ClientCapabilities) GetDeviceProfileOk() (*DlnaDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ClientCapabilities) SetDeviceProfile(v DlnaDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ClientCapabilities) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetIconUrl

`func (o *ClientCapabilities) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ClientCapabilities) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ClientCapabilities) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ClientCapabilities) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAppId

`func (o *ClientCapabilities) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ClientCapabilities) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ClientCapabilities) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ClientCapabilities) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



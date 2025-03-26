# SessionSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**PlayerStateInfo**](PlayerStateInfo.md) |  | [optional] 
**AdditionalUsers** | Pointer to [**[]SessionUserInfo**](SessionUserInfo.md) |  | [optional] 
**Capabilities** | Pointer to [**ClientCapabilities**](ClientCapabilities.md) |  | [optional] 
**RemoteEndPoint** | Pointer to **string** |  | [optional] 
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**LastActivityDate** | Pointer to **time.Time** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**NowPlayingItem** | Pointer to [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**AppIconUrl** | Pointer to **string** |  | [optional] 
**SupportedCommands** | Pointer to **[]string** |  | [optional] 
**TranscodingInfo** | Pointer to [**TranscodingInfo**](TranscodingInfo.md) |  | [optional] 
**SupportsRemoteControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewSessionSessionInfo

`func NewSessionSessionInfo() *SessionSessionInfo`

NewSessionSessionInfo instantiates a new SessionSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSessionInfoWithDefaults

`func NewSessionSessionInfoWithDefaults() *SessionSessionInfo`

NewSessionSessionInfoWithDefaults instantiates a new SessionSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *SessionSessionInfo) GetPlayState() PlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *SessionSessionInfo) GetPlayStateOk() (*PlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *SessionSessionInfo) SetPlayState(v PlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *SessionSessionInfo) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### GetAdditionalUsers

`func (o *SessionSessionInfo) GetAdditionalUsers() []SessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *SessionSessionInfo) GetAdditionalUsersOk() (*[]SessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *SessionSessionInfo) SetAdditionalUsers(v []SessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *SessionSessionInfo) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *SessionSessionInfo) GetCapabilities() ClientCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SessionSessionInfo) GetCapabilitiesOk() (*ClientCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SessionSessionInfo) SetCapabilities(v ClientCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SessionSessionInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetRemoteEndPoint

`func (o *SessionSessionInfo) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *SessionSessionInfo) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *SessionSessionInfo) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *SessionSessionInfo) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### GetPlayableMediaTypes

`func (o *SessionSessionInfo) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *SessionSessionInfo) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *SessionSessionInfo) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *SessionSessionInfo) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SessionSessionInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SessionSessionInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SessionSessionInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SessionSessionInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetId

`func (o *SessionSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerId

`func (o *SessionSessionInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SessionSessionInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SessionSessionInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SessionSessionInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetUserId

`func (o *SessionSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SessionSessionInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionSessionInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionSessionInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionSessionInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *SessionSessionInfo) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *SessionSessionInfo) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *SessionSessionInfo) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *SessionSessionInfo) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### GetClient

`func (o *SessionSessionInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionSessionInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionSessionInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionSessionInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetLastActivityDate

`func (o *SessionSessionInfo) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *SessionSessionInfo) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *SessionSessionInfo) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *SessionSessionInfo) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetDeviceName

`func (o *SessionSessionInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionSessionInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionSessionInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SessionSessionInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *SessionSessionInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionSessionInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionSessionInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SessionSessionInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetNowPlayingItem

`func (o *SessionSessionInfo) GetNowPlayingItem() BaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *SessionSessionInfo) GetNowPlayingItemOk() (*BaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *SessionSessionInfo) SetNowPlayingItem(v BaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *SessionSessionInfo) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### GetDeviceId

`func (o *SessionSessionInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionSessionInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionSessionInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionSessionInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *SessionSessionInfo) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *SessionSessionInfo) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *SessionSessionInfo) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *SessionSessionInfo) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetAppIconUrl

`func (o *SessionSessionInfo) GetAppIconUrl() string`

GetAppIconUrl returns the AppIconUrl field if non-nil, zero value otherwise.

### GetAppIconUrlOk

`func (o *SessionSessionInfo) GetAppIconUrlOk() (*string, bool)`

GetAppIconUrlOk returns a tuple with the AppIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIconUrl

`func (o *SessionSessionInfo) SetAppIconUrl(v string)`

SetAppIconUrl sets AppIconUrl field to given value.

### HasAppIconUrl

`func (o *SessionSessionInfo) HasAppIconUrl() bool`

HasAppIconUrl returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *SessionSessionInfo) GetSupportedCommands() []string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *SessionSessionInfo) GetSupportedCommandsOk() (*[]string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *SessionSessionInfo) SetSupportedCommands(v []string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *SessionSessionInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetTranscodingInfo

`func (o *SessionSessionInfo) GetTranscodingInfo() TranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *SessionSessionInfo) GetTranscodingInfoOk() (*TranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *SessionSessionInfo) SetTranscodingInfo(v TranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *SessionSessionInfo) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *SessionSessionInfo) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *SessionSessionInfo) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *SessionSessionInfo) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *SessionSessionInfo) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



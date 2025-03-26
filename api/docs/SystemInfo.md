# SystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemUpdateLevel** | Pointer to **string** |  | [optional] 
**OperatingSystemDisplayName** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**HasPendingRestart** | Pointer to **bool** |  | [optional] 
**IsShuttingDown** | Pointer to **bool** |  | [optional] 
**SupportsLibraryMonitor** | Pointer to **bool** |  | [optional] 
**WebSocketPortNumber** | Pointer to **int32** |  | [optional] 
**CompletedInstallations** | Pointer to [**[]UpdatesInstallationInfo**](UpdatesInstallationInfo.md) |  | [optional] 
**CanSelfRestart** | Pointer to **bool** |  | [optional] 
**CanSelfUpdate** | Pointer to **bool** |  | [optional] 
**CanLaunchWebBrowser** | Pointer to **bool** |  | [optional] 
**ProgramDataPath** | Pointer to **string** |  | [optional] 
**ItemsByNamePath** | Pointer to **string** |  | [optional] 
**CachePath** | Pointer to **string** |  | [optional] 
**LogPath** | Pointer to **string** |  | [optional] 
**InternalMetadataPath** | Pointer to **string** |  | [optional] 
**TranscodingTempPath** | Pointer to **string** |  | [optional] 
**HttpServerPortNumber** | Pointer to **int32** |  | [optional] 
**SupportsHttps** | Pointer to **bool** |  | [optional] 
**HttpsPortNumber** | Pointer to **int32** |  | [optional] 
**HasUpdateAvailable** | Pointer to **bool** |  | [optional] 
**SupportsAutoRunAtStartup** | Pointer to **bool** |  | [optional] 
**HardwareAccelerationRequiresPremiere** | Pointer to **bool** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**WanAddress** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemInfo

`func NewSystemInfo() *SystemInfo`

NewSystemInfo instantiates a new SystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoWithDefaults

`func NewSystemInfoWithDefaults() *SystemInfo`

NewSystemInfoWithDefaults instantiates a new SystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemUpdateLevel

`func (o *SystemInfo) GetSystemUpdateLevel() string`

GetSystemUpdateLevel returns the SystemUpdateLevel field if non-nil, zero value otherwise.

### GetSystemUpdateLevelOk

`func (o *SystemInfo) GetSystemUpdateLevelOk() (*string, bool)`

GetSystemUpdateLevelOk returns a tuple with the SystemUpdateLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpdateLevel

`func (o *SystemInfo) SetSystemUpdateLevel(v string)`

SetSystemUpdateLevel sets SystemUpdateLevel field to given value.

### HasSystemUpdateLevel

`func (o *SystemInfo) HasSystemUpdateLevel() bool`

HasSystemUpdateLevel returns a boolean if a field has been set.

### GetOperatingSystemDisplayName

`func (o *SystemInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *SystemInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *SystemInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *SystemInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### GetPackageName

`func (o *SystemInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *SystemInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *SystemInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *SystemInfo) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetHasPendingRestart

`func (o *SystemInfo) GetHasPendingRestart() bool`

GetHasPendingRestart returns the HasPendingRestart field if non-nil, zero value otherwise.

### GetHasPendingRestartOk

`func (o *SystemInfo) GetHasPendingRestartOk() (*bool, bool)`

GetHasPendingRestartOk returns a tuple with the HasPendingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingRestart

`func (o *SystemInfo) SetHasPendingRestart(v bool)`

SetHasPendingRestart sets HasPendingRestart field to given value.

### HasHasPendingRestart

`func (o *SystemInfo) HasHasPendingRestart() bool`

HasHasPendingRestart returns a boolean if a field has been set.

### GetIsShuttingDown

`func (o *SystemInfo) GetIsShuttingDown() bool`

GetIsShuttingDown returns the IsShuttingDown field if non-nil, zero value otherwise.

### GetIsShuttingDownOk

`func (o *SystemInfo) GetIsShuttingDownOk() (*bool, bool)`

GetIsShuttingDownOk returns a tuple with the IsShuttingDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuttingDown

`func (o *SystemInfo) SetIsShuttingDown(v bool)`

SetIsShuttingDown sets IsShuttingDown field to given value.

### HasIsShuttingDown

`func (o *SystemInfo) HasIsShuttingDown() bool`

HasIsShuttingDown returns a boolean if a field has been set.

### GetSupportsLibraryMonitor

`func (o *SystemInfo) GetSupportsLibraryMonitor() bool`

GetSupportsLibraryMonitor returns the SupportsLibraryMonitor field if non-nil, zero value otherwise.

### GetSupportsLibraryMonitorOk

`func (o *SystemInfo) GetSupportsLibraryMonitorOk() (*bool, bool)`

GetSupportsLibraryMonitorOk returns a tuple with the SupportsLibraryMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLibraryMonitor

`func (o *SystemInfo) SetSupportsLibraryMonitor(v bool)`

SetSupportsLibraryMonitor sets SupportsLibraryMonitor field to given value.

### HasSupportsLibraryMonitor

`func (o *SystemInfo) HasSupportsLibraryMonitor() bool`

HasSupportsLibraryMonitor returns a boolean if a field has been set.

### GetWebSocketPortNumber

`func (o *SystemInfo) GetWebSocketPortNumber() int32`

GetWebSocketPortNumber returns the WebSocketPortNumber field if non-nil, zero value otherwise.

### GetWebSocketPortNumberOk

`func (o *SystemInfo) GetWebSocketPortNumberOk() (*int32, bool)`

GetWebSocketPortNumberOk returns a tuple with the WebSocketPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSocketPortNumber

`func (o *SystemInfo) SetWebSocketPortNumber(v int32)`

SetWebSocketPortNumber sets WebSocketPortNumber field to given value.

### HasWebSocketPortNumber

`func (o *SystemInfo) HasWebSocketPortNumber() bool`

HasWebSocketPortNumber returns a boolean if a field has been set.

### GetCompletedInstallations

`func (o *SystemInfo) GetCompletedInstallations() []UpdatesInstallationInfo`

GetCompletedInstallations returns the CompletedInstallations field if non-nil, zero value otherwise.

### GetCompletedInstallationsOk

`func (o *SystemInfo) GetCompletedInstallationsOk() (*[]UpdatesInstallationInfo, bool)`

GetCompletedInstallationsOk returns a tuple with the CompletedInstallations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInstallations

`func (o *SystemInfo) SetCompletedInstallations(v []UpdatesInstallationInfo)`

SetCompletedInstallations sets CompletedInstallations field to given value.

### HasCompletedInstallations

`func (o *SystemInfo) HasCompletedInstallations() bool`

HasCompletedInstallations returns a boolean if a field has been set.

### GetCanSelfRestart

`func (o *SystemInfo) GetCanSelfRestart() bool`

GetCanSelfRestart returns the CanSelfRestart field if non-nil, zero value otherwise.

### GetCanSelfRestartOk

`func (o *SystemInfo) GetCanSelfRestartOk() (*bool, bool)`

GetCanSelfRestartOk returns a tuple with the CanSelfRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfRestart

`func (o *SystemInfo) SetCanSelfRestart(v bool)`

SetCanSelfRestart sets CanSelfRestart field to given value.

### HasCanSelfRestart

`func (o *SystemInfo) HasCanSelfRestart() bool`

HasCanSelfRestart returns a boolean if a field has been set.

### GetCanSelfUpdate

`func (o *SystemInfo) GetCanSelfUpdate() bool`

GetCanSelfUpdate returns the CanSelfUpdate field if non-nil, zero value otherwise.

### GetCanSelfUpdateOk

`func (o *SystemInfo) GetCanSelfUpdateOk() (*bool, bool)`

GetCanSelfUpdateOk returns a tuple with the CanSelfUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfUpdate

`func (o *SystemInfo) SetCanSelfUpdate(v bool)`

SetCanSelfUpdate sets CanSelfUpdate field to given value.

### HasCanSelfUpdate

`func (o *SystemInfo) HasCanSelfUpdate() bool`

HasCanSelfUpdate returns a boolean if a field has been set.

### GetCanLaunchWebBrowser

`func (o *SystemInfo) GetCanLaunchWebBrowser() bool`

GetCanLaunchWebBrowser returns the CanLaunchWebBrowser field if non-nil, zero value otherwise.

### GetCanLaunchWebBrowserOk

`func (o *SystemInfo) GetCanLaunchWebBrowserOk() (*bool, bool)`

GetCanLaunchWebBrowserOk returns a tuple with the CanLaunchWebBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchWebBrowser

`func (o *SystemInfo) SetCanLaunchWebBrowser(v bool)`

SetCanLaunchWebBrowser sets CanLaunchWebBrowser field to given value.

### HasCanLaunchWebBrowser

`func (o *SystemInfo) HasCanLaunchWebBrowser() bool`

HasCanLaunchWebBrowser returns a boolean if a field has been set.

### GetProgramDataPath

`func (o *SystemInfo) GetProgramDataPath() string`

GetProgramDataPath returns the ProgramDataPath field if non-nil, zero value otherwise.

### GetProgramDataPathOk

`func (o *SystemInfo) GetProgramDataPathOk() (*string, bool)`

GetProgramDataPathOk returns a tuple with the ProgramDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDataPath

`func (o *SystemInfo) SetProgramDataPath(v string)`

SetProgramDataPath sets ProgramDataPath field to given value.

### HasProgramDataPath

`func (o *SystemInfo) HasProgramDataPath() bool`

HasProgramDataPath returns a boolean if a field has been set.

### GetItemsByNamePath

`func (o *SystemInfo) GetItemsByNamePath() string`

GetItemsByNamePath returns the ItemsByNamePath field if non-nil, zero value otherwise.

### GetItemsByNamePathOk

`func (o *SystemInfo) GetItemsByNamePathOk() (*string, bool)`

GetItemsByNamePathOk returns a tuple with the ItemsByNamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsByNamePath

`func (o *SystemInfo) SetItemsByNamePath(v string)`

SetItemsByNamePath sets ItemsByNamePath field to given value.

### HasItemsByNamePath

`func (o *SystemInfo) HasItemsByNamePath() bool`

HasItemsByNamePath returns a boolean if a field has been set.

### GetCachePath

`func (o *SystemInfo) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *SystemInfo) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *SystemInfo) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *SystemInfo) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### GetLogPath

`func (o *SystemInfo) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *SystemInfo) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *SystemInfo) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *SystemInfo) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetInternalMetadataPath

`func (o *SystemInfo) GetInternalMetadataPath() string`

GetInternalMetadataPath returns the InternalMetadataPath field if non-nil, zero value otherwise.

### GetInternalMetadataPathOk

`func (o *SystemInfo) GetInternalMetadataPathOk() (*string, bool)`

GetInternalMetadataPathOk returns a tuple with the InternalMetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadataPath

`func (o *SystemInfo) SetInternalMetadataPath(v string)`

SetInternalMetadataPath sets InternalMetadataPath field to given value.

### HasInternalMetadataPath

`func (o *SystemInfo) HasInternalMetadataPath() bool`

HasInternalMetadataPath returns a boolean if a field has been set.

### GetTranscodingTempPath

`func (o *SystemInfo) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *SystemInfo) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *SystemInfo) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *SystemInfo) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### GetHttpServerPortNumber

`func (o *SystemInfo) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *SystemInfo) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *SystemInfo) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *SystemInfo) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetSupportsHttps

`func (o *SystemInfo) GetSupportsHttps() bool`

GetSupportsHttps returns the SupportsHttps field if non-nil, zero value otherwise.

### GetSupportsHttpsOk

`func (o *SystemInfo) GetSupportsHttpsOk() (*bool, bool)`

GetSupportsHttpsOk returns a tuple with the SupportsHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHttps

`func (o *SystemInfo) SetSupportsHttps(v bool)`

SetSupportsHttps sets SupportsHttps field to given value.

### HasSupportsHttps

`func (o *SystemInfo) HasSupportsHttps() bool`

HasSupportsHttps returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *SystemInfo) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *SystemInfo) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *SystemInfo) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *SystemInfo) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

### GetHasUpdateAvailable

`func (o *SystemInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *SystemInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *SystemInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *SystemInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetSupportsAutoRunAtStartup

`func (o *SystemInfo) GetSupportsAutoRunAtStartup() bool`

GetSupportsAutoRunAtStartup returns the SupportsAutoRunAtStartup field if non-nil, zero value otherwise.

### GetSupportsAutoRunAtStartupOk

`func (o *SystemInfo) GetSupportsAutoRunAtStartupOk() (*bool, bool)`

GetSupportsAutoRunAtStartupOk returns a tuple with the SupportsAutoRunAtStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoRunAtStartup

`func (o *SystemInfo) SetSupportsAutoRunAtStartup(v bool)`

SetSupportsAutoRunAtStartup sets SupportsAutoRunAtStartup field to given value.

### HasSupportsAutoRunAtStartup

`func (o *SystemInfo) HasSupportsAutoRunAtStartup() bool`

HasSupportsAutoRunAtStartup returns a boolean if a field has been set.

### GetHardwareAccelerationRequiresPremiere

`func (o *SystemInfo) GetHardwareAccelerationRequiresPremiere() bool`

GetHardwareAccelerationRequiresPremiere returns the HardwareAccelerationRequiresPremiere field if non-nil, zero value otherwise.

### GetHardwareAccelerationRequiresPremiereOk

`func (o *SystemInfo) GetHardwareAccelerationRequiresPremiereOk() (*bool, bool)`

GetHardwareAccelerationRequiresPremiereOk returns a tuple with the HardwareAccelerationRequiresPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationRequiresPremiere

`func (o *SystemInfo) SetHardwareAccelerationRequiresPremiere(v bool)`

SetHardwareAccelerationRequiresPremiere sets HardwareAccelerationRequiresPremiere field to given value.

### HasHardwareAccelerationRequiresPremiere

`func (o *SystemInfo) HasHardwareAccelerationRequiresPremiere() bool`

HasHardwareAccelerationRequiresPremiere returns a boolean if a field has been set.

### GetLocalAddress

`func (o *SystemInfo) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *SystemInfo) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *SystemInfo) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *SystemInfo) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetWanAddress

`func (o *SystemInfo) GetWanAddress() string`

GetWanAddress returns the WanAddress field if non-nil, zero value otherwise.

### GetWanAddressOk

`func (o *SystemInfo) GetWanAddressOk() (*string, bool)`

GetWanAddressOk returns a tuple with the WanAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAddress

`func (o *SystemInfo) SetWanAddress(v string)`

SetWanAddress sets WanAddress field to given value.

### HasWanAddress

`func (o *SystemInfo) HasWanAddress() bool`

HasWanAddress returns a boolean if a field has been set.

### GetServerName

`func (o *SystemInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *SystemInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *SystemInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *SystemInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetVersion

`func (o *SystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetId

`func (o *SystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



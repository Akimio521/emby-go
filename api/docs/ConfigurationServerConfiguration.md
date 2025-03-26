# ConfigurationServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableUPnP** | Pointer to **bool** |  | [optional] 
**PublicPort** | Pointer to **int32** |  | [optional] 
**PublicHttpsPort** | Pointer to **int32** |  | [optional] 
**HttpServerPortNumber** | Pointer to **int32** |  | [optional] 
**HttpsPortNumber** | Pointer to **int32** |  | [optional] 
**EnableHttps** | Pointer to **bool** |  | [optional] 
**SubtitlePermissionsUpgraded** | Pointer to **bool** |  | [optional] 
**CertificatePath** | Pointer to **string** |  | [optional] 
**CertificatePassword** | Pointer to **string** |  | [optional] 
**IsPortAuthorized** | Pointer to **bool** |  | [optional] 
**AutoRunWebApp** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**LogAllQueryTimes** | Pointer to **bool** |  | [optional] 
**EnableCaseSensitiveItemIds** | Pointer to **bool** |  | [optional] 
**MetadataPath** | Pointer to **string** |  | [optional] 
**MetadataNetworkPath** | Pointer to **string** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**SortReplaceCharacters** | Pointer to **[]string** |  | [optional] 
**SortRemoveCharacters** | Pointer to **[]string** |  | [optional] 
**SortRemoveWords** | Pointer to **[]string** |  | [optional] 
**LibraryMonitorDelay** | Pointer to **int32** |  | [optional] 
**EnableDashboardResponseCaching** | Pointer to **bool** |  | [optional] 
**DashboardSourcePath** | Pointer to **string** |  | [optional] 
**ImageSavingConvention** | Pointer to **string** |  | [optional] 
**EnableAutomaticRestart** | Pointer to **bool** |  | [optional] 
**SkipDeserializationForBasicTypes** | Pointer to **bool** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**WanDdns** | Pointer to **string** |  | [optional] 
**UICulture** | Pointer to **string** |  | [optional] 
**SaveMetadataHidden** | Pointer to **bool** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**DisplaySpecialsWithinSeasons** | Pointer to **bool** |  | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** |  | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** |  | [optional] 
**EnableExternalContentInSuggestions** | Pointer to **bool** |  | [optional] 
**RequireHttps** | Pointer to **bool** |  | [optional] 
**IsBehindProxy** | Pointer to **bool** |  | [optional] 
**RemoteIPFilter** | Pointer to **[]string** |  | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** |  | [optional] 
**ImageExtractionTimeoutMs** | Pointer to **int32** |  | [optional] 
**PathSubstitutions** | Pointer to [**[]ConfigurationPathSubstitution**](ConfigurationPathSubstitution.md) |  | [optional] 
**UninstalledPlugins** | Pointer to **[]string** |  | [optional] 
**CollapseVideoFolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalTrackTitles** | Pointer to **bool** |  | [optional] 
**EnableDebugLevelLogging** | Pointer to **bool** |  | [optional] 
**EnableAutoUpdate** | Pointer to **bool** |  | [optional] 
**LogFileRetentionDays** | Pointer to **int32** |  | [optional] 
**RunAtStartup** | Pointer to **bool** |  | [optional] 
**IsStartupWizardCompleted** | Pointer to **bool** |  | [optional] 
**CachePath** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigurationServerConfiguration

`func NewConfigurationServerConfiguration() *ConfigurationServerConfiguration`

NewConfigurationServerConfiguration instantiates a new ConfigurationServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationServerConfigurationWithDefaults

`func NewConfigurationServerConfigurationWithDefaults() *ConfigurationServerConfiguration`

NewConfigurationServerConfigurationWithDefaults instantiates a new ConfigurationServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableUPnP

`func (o *ConfigurationServerConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *ConfigurationServerConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *ConfigurationServerConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *ConfigurationServerConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetPublicPort

`func (o *ConfigurationServerConfiguration) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *ConfigurationServerConfiguration) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *ConfigurationServerConfiguration) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *ConfigurationServerConfiguration) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *ConfigurationServerConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *ConfigurationServerConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *ConfigurationServerConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *ConfigurationServerConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetHttpServerPortNumber

`func (o *ConfigurationServerConfiguration) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *ConfigurationServerConfiguration) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *ConfigurationServerConfiguration) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *ConfigurationServerConfiguration) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *ConfigurationServerConfiguration) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *ConfigurationServerConfiguration) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *ConfigurationServerConfiguration) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *ConfigurationServerConfiguration) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

### GetEnableHttps

`func (o *ConfigurationServerConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *ConfigurationServerConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *ConfigurationServerConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *ConfigurationServerConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetSubtitlePermissionsUpgraded

`func (o *ConfigurationServerConfiguration) GetSubtitlePermissionsUpgraded() bool`

GetSubtitlePermissionsUpgraded returns the SubtitlePermissionsUpgraded field if non-nil, zero value otherwise.

### GetSubtitlePermissionsUpgradedOk

`func (o *ConfigurationServerConfiguration) GetSubtitlePermissionsUpgradedOk() (*bool, bool)`

GetSubtitlePermissionsUpgradedOk returns a tuple with the SubtitlePermissionsUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitlePermissionsUpgraded

`func (o *ConfigurationServerConfiguration) SetSubtitlePermissionsUpgraded(v bool)`

SetSubtitlePermissionsUpgraded sets SubtitlePermissionsUpgraded field to given value.

### HasSubtitlePermissionsUpgraded

`func (o *ConfigurationServerConfiguration) HasSubtitlePermissionsUpgraded() bool`

HasSubtitlePermissionsUpgraded returns a boolean if a field has been set.

### GetCertificatePath

`func (o *ConfigurationServerConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *ConfigurationServerConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *ConfigurationServerConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *ConfigurationServerConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *ConfigurationServerConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *ConfigurationServerConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *ConfigurationServerConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *ConfigurationServerConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetIsPortAuthorized

`func (o *ConfigurationServerConfiguration) GetIsPortAuthorized() bool`

GetIsPortAuthorized returns the IsPortAuthorized field if non-nil, zero value otherwise.

### GetIsPortAuthorizedOk

`func (o *ConfigurationServerConfiguration) GetIsPortAuthorizedOk() (*bool, bool)`

GetIsPortAuthorizedOk returns a tuple with the IsPortAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPortAuthorized

`func (o *ConfigurationServerConfiguration) SetIsPortAuthorized(v bool)`

SetIsPortAuthorized sets IsPortAuthorized field to given value.

### HasIsPortAuthorized

`func (o *ConfigurationServerConfiguration) HasIsPortAuthorized() bool`

HasIsPortAuthorized returns a boolean if a field has been set.

### GetAutoRunWebApp

`func (o *ConfigurationServerConfiguration) GetAutoRunWebApp() bool`

GetAutoRunWebApp returns the AutoRunWebApp field if non-nil, zero value otherwise.

### GetAutoRunWebAppOk

`func (o *ConfigurationServerConfiguration) GetAutoRunWebAppOk() (*bool, bool)`

GetAutoRunWebAppOk returns a tuple with the AutoRunWebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunWebApp

`func (o *ConfigurationServerConfiguration) SetAutoRunWebApp(v bool)`

SetAutoRunWebApp sets AutoRunWebApp field to given value.

### HasAutoRunWebApp

`func (o *ConfigurationServerConfiguration) HasAutoRunWebApp() bool`

HasAutoRunWebApp returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *ConfigurationServerConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *ConfigurationServerConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *ConfigurationServerConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *ConfigurationServerConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLogAllQueryTimes

`func (o *ConfigurationServerConfiguration) GetLogAllQueryTimes() bool`

GetLogAllQueryTimes returns the LogAllQueryTimes field if non-nil, zero value otherwise.

### GetLogAllQueryTimesOk

`func (o *ConfigurationServerConfiguration) GetLogAllQueryTimesOk() (*bool, bool)`

GetLogAllQueryTimesOk returns a tuple with the LogAllQueryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAllQueryTimes

`func (o *ConfigurationServerConfiguration) SetLogAllQueryTimes(v bool)`

SetLogAllQueryTimes sets LogAllQueryTimes field to given value.

### HasLogAllQueryTimes

`func (o *ConfigurationServerConfiguration) HasLogAllQueryTimes() bool`

HasLogAllQueryTimes returns a boolean if a field has been set.

### GetEnableCaseSensitiveItemIds

`func (o *ConfigurationServerConfiguration) GetEnableCaseSensitiveItemIds() bool`

GetEnableCaseSensitiveItemIds returns the EnableCaseSensitiveItemIds field if non-nil, zero value otherwise.

### GetEnableCaseSensitiveItemIdsOk

`func (o *ConfigurationServerConfiguration) GetEnableCaseSensitiveItemIdsOk() (*bool, bool)`

GetEnableCaseSensitiveItemIdsOk returns a tuple with the EnableCaseSensitiveItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaseSensitiveItemIds

`func (o *ConfigurationServerConfiguration) SetEnableCaseSensitiveItemIds(v bool)`

SetEnableCaseSensitiveItemIds sets EnableCaseSensitiveItemIds field to given value.

### HasEnableCaseSensitiveItemIds

`func (o *ConfigurationServerConfiguration) HasEnableCaseSensitiveItemIds() bool`

HasEnableCaseSensitiveItemIds returns a boolean if a field has been set.

### GetMetadataPath

`func (o *ConfigurationServerConfiguration) GetMetadataPath() string`

GetMetadataPath returns the MetadataPath field if non-nil, zero value otherwise.

### GetMetadataPathOk

`func (o *ConfigurationServerConfiguration) GetMetadataPathOk() (*string, bool)`

GetMetadataPathOk returns a tuple with the MetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPath

`func (o *ConfigurationServerConfiguration) SetMetadataPath(v string)`

SetMetadataPath sets MetadataPath field to given value.

### HasMetadataPath

`func (o *ConfigurationServerConfiguration) HasMetadataPath() bool`

HasMetadataPath returns a boolean if a field has been set.

### GetMetadataNetworkPath

`func (o *ConfigurationServerConfiguration) GetMetadataNetworkPath() string`

GetMetadataNetworkPath returns the MetadataNetworkPath field if non-nil, zero value otherwise.

### GetMetadataNetworkPathOk

`func (o *ConfigurationServerConfiguration) GetMetadataNetworkPathOk() (*string, bool)`

GetMetadataNetworkPathOk returns a tuple with the MetadataNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataNetworkPath

`func (o *ConfigurationServerConfiguration) SetMetadataNetworkPath(v string)`

SetMetadataNetworkPath sets MetadataNetworkPath field to given value.

### HasMetadataNetworkPath

`func (o *ConfigurationServerConfiguration) HasMetadataNetworkPath() bool`

HasMetadataNetworkPath returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ConfigurationServerConfiguration) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ConfigurationServerConfiguration) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ConfigurationServerConfiguration) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ConfigurationServerConfiguration) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ConfigurationServerConfiguration) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ConfigurationServerConfiguration) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ConfigurationServerConfiguration) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ConfigurationServerConfiguration) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSortReplaceCharacters

`func (o *ConfigurationServerConfiguration) GetSortReplaceCharacters() []string`

GetSortReplaceCharacters returns the SortReplaceCharacters field if non-nil, zero value otherwise.

### GetSortReplaceCharactersOk

`func (o *ConfigurationServerConfiguration) GetSortReplaceCharactersOk() (*[]string, bool)`

GetSortReplaceCharactersOk returns a tuple with the SortReplaceCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortReplaceCharacters

`func (o *ConfigurationServerConfiguration) SetSortReplaceCharacters(v []string)`

SetSortReplaceCharacters sets SortReplaceCharacters field to given value.

### HasSortReplaceCharacters

`func (o *ConfigurationServerConfiguration) HasSortReplaceCharacters() bool`

HasSortReplaceCharacters returns a boolean if a field has been set.

### GetSortRemoveCharacters

`func (o *ConfigurationServerConfiguration) GetSortRemoveCharacters() []string`

GetSortRemoveCharacters returns the SortRemoveCharacters field if non-nil, zero value otherwise.

### GetSortRemoveCharactersOk

`func (o *ConfigurationServerConfiguration) GetSortRemoveCharactersOk() (*[]string, bool)`

GetSortRemoveCharactersOk returns a tuple with the SortRemoveCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveCharacters

`func (o *ConfigurationServerConfiguration) SetSortRemoveCharacters(v []string)`

SetSortRemoveCharacters sets SortRemoveCharacters field to given value.

### HasSortRemoveCharacters

`func (o *ConfigurationServerConfiguration) HasSortRemoveCharacters() bool`

HasSortRemoveCharacters returns a boolean if a field has been set.

### GetSortRemoveWords

`func (o *ConfigurationServerConfiguration) GetSortRemoveWords() []string`

GetSortRemoveWords returns the SortRemoveWords field if non-nil, zero value otherwise.

### GetSortRemoveWordsOk

`func (o *ConfigurationServerConfiguration) GetSortRemoveWordsOk() (*[]string, bool)`

GetSortRemoveWordsOk returns a tuple with the SortRemoveWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveWords

`func (o *ConfigurationServerConfiguration) SetSortRemoveWords(v []string)`

SetSortRemoveWords sets SortRemoveWords field to given value.

### HasSortRemoveWords

`func (o *ConfigurationServerConfiguration) HasSortRemoveWords() bool`

HasSortRemoveWords returns a boolean if a field has been set.

### GetLibraryMonitorDelay

`func (o *ConfigurationServerConfiguration) GetLibraryMonitorDelay() int32`

GetLibraryMonitorDelay returns the LibraryMonitorDelay field if non-nil, zero value otherwise.

### GetLibraryMonitorDelayOk

`func (o *ConfigurationServerConfiguration) GetLibraryMonitorDelayOk() (*int32, bool)`

GetLibraryMonitorDelayOk returns a tuple with the LibraryMonitorDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMonitorDelay

`func (o *ConfigurationServerConfiguration) SetLibraryMonitorDelay(v int32)`

SetLibraryMonitorDelay sets LibraryMonitorDelay field to given value.

### HasLibraryMonitorDelay

`func (o *ConfigurationServerConfiguration) HasLibraryMonitorDelay() bool`

HasLibraryMonitorDelay returns a boolean if a field has been set.

### GetEnableDashboardResponseCaching

`func (o *ConfigurationServerConfiguration) GetEnableDashboardResponseCaching() bool`

GetEnableDashboardResponseCaching returns the EnableDashboardResponseCaching field if non-nil, zero value otherwise.

### GetEnableDashboardResponseCachingOk

`func (o *ConfigurationServerConfiguration) GetEnableDashboardResponseCachingOk() (*bool, bool)`

GetEnableDashboardResponseCachingOk returns a tuple with the EnableDashboardResponseCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDashboardResponseCaching

`func (o *ConfigurationServerConfiguration) SetEnableDashboardResponseCaching(v bool)`

SetEnableDashboardResponseCaching sets EnableDashboardResponseCaching field to given value.

### HasEnableDashboardResponseCaching

`func (o *ConfigurationServerConfiguration) HasEnableDashboardResponseCaching() bool`

HasEnableDashboardResponseCaching returns a boolean if a field has been set.

### GetDashboardSourcePath

`func (o *ConfigurationServerConfiguration) GetDashboardSourcePath() string`

GetDashboardSourcePath returns the DashboardSourcePath field if non-nil, zero value otherwise.

### GetDashboardSourcePathOk

`func (o *ConfigurationServerConfiguration) GetDashboardSourcePathOk() (*string, bool)`

GetDashboardSourcePathOk returns a tuple with the DashboardSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSourcePath

`func (o *ConfigurationServerConfiguration) SetDashboardSourcePath(v string)`

SetDashboardSourcePath sets DashboardSourcePath field to given value.

### HasDashboardSourcePath

`func (o *ConfigurationServerConfiguration) HasDashboardSourcePath() bool`

HasDashboardSourcePath returns a boolean if a field has been set.

### GetImageSavingConvention

`func (o *ConfigurationServerConfiguration) GetImageSavingConvention() string`

GetImageSavingConvention returns the ImageSavingConvention field if non-nil, zero value otherwise.

### GetImageSavingConventionOk

`func (o *ConfigurationServerConfiguration) GetImageSavingConventionOk() (*string, bool)`

GetImageSavingConventionOk returns a tuple with the ImageSavingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSavingConvention

`func (o *ConfigurationServerConfiguration) SetImageSavingConvention(v string)`

SetImageSavingConvention sets ImageSavingConvention field to given value.

### HasImageSavingConvention

`func (o *ConfigurationServerConfiguration) HasImageSavingConvention() bool`

HasImageSavingConvention returns a boolean if a field has been set.

### GetEnableAutomaticRestart

`func (o *ConfigurationServerConfiguration) GetEnableAutomaticRestart() bool`

GetEnableAutomaticRestart returns the EnableAutomaticRestart field if non-nil, zero value otherwise.

### GetEnableAutomaticRestartOk

`func (o *ConfigurationServerConfiguration) GetEnableAutomaticRestartOk() (*bool, bool)`

GetEnableAutomaticRestartOk returns a tuple with the EnableAutomaticRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRestart

`func (o *ConfigurationServerConfiguration) SetEnableAutomaticRestart(v bool)`

SetEnableAutomaticRestart sets EnableAutomaticRestart field to given value.

### HasEnableAutomaticRestart

`func (o *ConfigurationServerConfiguration) HasEnableAutomaticRestart() bool`

HasEnableAutomaticRestart returns a boolean if a field has been set.

### GetSkipDeserializationForBasicTypes

`func (o *ConfigurationServerConfiguration) GetSkipDeserializationForBasicTypes() bool`

GetSkipDeserializationForBasicTypes returns the SkipDeserializationForBasicTypes field if non-nil, zero value otherwise.

### GetSkipDeserializationForBasicTypesOk

`func (o *ConfigurationServerConfiguration) GetSkipDeserializationForBasicTypesOk() (*bool, bool)`

GetSkipDeserializationForBasicTypesOk returns a tuple with the SkipDeserializationForBasicTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDeserializationForBasicTypes

`func (o *ConfigurationServerConfiguration) SetSkipDeserializationForBasicTypes(v bool)`

SetSkipDeserializationForBasicTypes sets SkipDeserializationForBasicTypes field to given value.

### HasSkipDeserializationForBasicTypes

`func (o *ConfigurationServerConfiguration) HasSkipDeserializationForBasicTypes() bool`

HasSkipDeserializationForBasicTypes returns a boolean if a field has been set.

### GetServerName

`func (o *ConfigurationServerConfiguration) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ConfigurationServerConfiguration) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ConfigurationServerConfiguration) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ConfigurationServerConfiguration) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetWanDdns

`func (o *ConfigurationServerConfiguration) GetWanDdns() string`

GetWanDdns returns the WanDdns field if non-nil, zero value otherwise.

### GetWanDdnsOk

`func (o *ConfigurationServerConfiguration) GetWanDdnsOk() (*string, bool)`

GetWanDdnsOk returns a tuple with the WanDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanDdns

`func (o *ConfigurationServerConfiguration) SetWanDdns(v string)`

SetWanDdns sets WanDdns field to given value.

### HasWanDdns

`func (o *ConfigurationServerConfiguration) HasWanDdns() bool`

HasWanDdns returns a boolean if a field has been set.

### GetUICulture

`func (o *ConfigurationServerConfiguration) GetUICulture() string`

GetUICulture returns the UICulture field if non-nil, zero value otherwise.

### GetUICultureOk

`func (o *ConfigurationServerConfiguration) GetUICultureOk() (*string, bool)`

GetUICultureOk returns a tuple with the UICulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUICulture

`func (o *ConfigurationServerConfiguration) SetUICulture(v string)`

SetUICulture sets UICulture field to given value.

### HasUICulture

`func (o *ConfigurationServerConfiguration) HasUICulture() bool`

HasUICulture returns a boolean if a field has been set.

### GetSaveMetadataHidden

`func (o *ConfigurationServerConfiguration) GetSaveMetadataHidden() bool`

GetSaveMetadataHidden returns the SaveMetadataHidden field if non-nil, zero value otherwise.

### GetSaveMetadataHiddenOk

`func (o *ConfigurationServerConfiguration) GetSaveMetadataHiddenOk() (*bool, bool)`

GetSaveMetadataHiddenOk returns a tuple with the SaveMetadataHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMetadataHidden

`func (o *ConfigurationServerConfiguration) SetSaveMetadataHidden(v bool)`

SetSaveMetadataHidden sets SaveMetadataHidden field to given value.

### HasSaveMetadataHidden

`func (o *ConfigurationServerConfiguration) HasSaveMetadataHidden() bool`

HasSaveMetadataHidden returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *ConfigurationServerConfiguration) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *ConfigurationServerConfiguration) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *ConfigurationServerConfiguration) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *ConfigurationServerConfiguration) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ConfigurationServerConfiguration) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ConfigurationServerConfiguration) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ConfigurationServerConfiguration) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ConfigurationServerConfiguration) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetDisplaySpecialsWithinSeasons

`func (o *ConfigurationServerConfiguration) GetDisplaySpecialsWithinSeasons() bool`

GetDisplaySpecialsWithinSeasons returns the DisplaySpecialsWithinSeasons field if non-nil, zero value otherwise.

### GetDisplaySpecialsWithinSeasonsOk

`func (o *ConfigurationServerConfiguration) GetDisplaySpecialsWithinSeasonsOk() (*bool, bool)`

GetDisplaySpecialsWithinSeasonsOk returns a tuple with the DisplaySpecialsWithinSeasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySpecialsWithinSeasons

`func (o *ConfigurationServerConfiguration) SetDisplaySpecialsWithinSeasons(v bool)`

SetDisplaySpecialsWithinSeasons sets DisplaySpecialsWithinSeasons field to given value.

### HasDisplaySpecialsWithinSeasons

`func (o *ConfigurationServerConfiguration) HasDisplaySpecialsWithinSeasons() bool`

HasDisplaySpecialsWithinSeasons returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *ConfigurationServerConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *ConfigurationServerConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *ConfigurationServerConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *ConfigurationServerConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *ConfigurationServerConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *ConfigurationServerConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *ConfigurationServerConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *ConfigurationServerConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetEnableExternalContentInSuggestions

`func (o *ConfigurationServerConfiguration) GetEnableExternalContentInSuggestions() bool`

GetEnableExternalContentInSuggestions returns the EnableExternalContentInSuggestions field if non-nil, zero value otherwise.

### GetEnableExternalContentInSuggestionsOk

`func (o *ConfigurationServerConfiguration) GetEnableExternalContentInSuggestionsOk() (*bool, bool)`

GetEnableExternalContentInSuggestionsOk returns a tuple with the EnableExternalContentInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalContentInSuggestions

`func (o *ConfigurationServerConfiguration) SetEnableExternalContentInSuggestions(v bool)`

SetEnableExternalContentInSuggestions sets EnableExternalContentInSuggestions field to given value.

### HasEnableExternalContentInSuggestions

`func (o *ConfigurationServerConfiguration) HasEnableExternalContentInSuggestions() bool`

HasEnableExternalContentInSuggestions returns a boolean if a field has been set.

### GetRequireHttps

`func (o *ConfigurationServerConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *ConfigurationServerConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *ConfigurationServerConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *ConfigurationServerConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetIsBehindProxy

`func (o *ConfigurationServerConfiguration) GetIsBehindProxy() bool`

GetIsBehindProxy returns the IsBehindProxy field if non-nil, zero value otherwise.

### GetIsBehindProxyOk

`func (o *ConfigurationServerConfiguration) GetIsBehindProxyOk() (*bool, bool)`

GetIsBehindProxyOk returns a tuple with the IsBehindProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBehindProxy

`func (o *ConfigurationServerConfiguration) SetIsBehindProxy(v bool)`

SetIsBehindProxy sets IsBehindProxy field to given value.

### HasIsBehindProxy

`func (o *ConfigurationServerConfiguration) HasIsBehindProxy() bool`

HasIsBehindProxy returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *ConfigurationServerConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *ConfigurationServerConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *ConfigurationServerConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *ConfigurationServerConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *ConfigurationServerConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *ConfigurationServerConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *ConfigurationServerConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *ConfigurationServerConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.

### GetImageExtractionTimeoutMs

`func (o *ConfigurationServerConfiguration) GetImageExtractionTimeoutMs() int32`

GetImageExtractionTimeoutMs returns the ImageExtractionTimeoutMs field if non-nil, zero value otherwise.

### GetImageExtractionTimeoutMsOk

`func (o *ConfigurationServerConfiguration) GetImageExtractionTimeoutMsOk() (*int32, bool)`

GetImageExtractionTimeoutMsOk returns a tuple with the ImageExtractionTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractionTimeoutMs

`func (o *ConfigurationServerConfiguration) SetImageExtractionTimeoutMs(v int32)`

SetImageExtractionTimeoutMs sets ImageExtractionTimeoutMs field to given value.

### HasImageExtractionTimeoutMs

`func (o *ConfigurationServerConfiguration) HasImageExtractionTimeoutMs() bool`

HasImageExtractionTimeoutMs returns a boolean if a field has been set.

### GetPathSubstitutions

`func (o *ConfigurationServerConfiguration) GetPathSubstitutions() []ConfigurationPathSubstitution`

GetPathSubstitutions returns the PathSubstitutions field if non-nil, zero value otherwise.

### GetPathSubstitutionsOk

`func (o *ConfigurationServerConfiguration) GetPathSubstitutionsOk() (*[]ConfigurationPathSubstitution, bool)`

GetPathSubstitutionsOk returns a tuple with the PathSubstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSubstitutions

`func (o *ConfigurationServerConfiguration) SetPathSubstitutions(v []ConfigurationPathSubstitution)`

SetPathSubstitutions sets PathSubstitutions field to given value.

### HasPathSubstitutions

`func (o *ConfigurationServerConfiguration) HasPathSubstitutions() bool`

HasPathSubstitutions returns a boolean if a field has been set.

### GetUninstalledPlugins

`func (o *ConfigurationServerConfiguration) GetUninstalledPlugins() []string`

GetUninstalledPlugins returns the UninstalledPlugins field if non-nil, zero value otherwise.

### GetUninstalledPluginsOk

`func (o *ConfigurationServerConfiguration) GetUninstalledPluginsOk() (*[]string, bool)`

GetUninstalledPluginsOk returns a tuple with the UninstalledPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledPlugins

`func (o *ConfigurationServerConfiguration) SetUninstalledPlugins(v []string)`

SetUninstalledPlugins sets UninstalledPlugins field to given value.

### HasUninstalledPlugins

`func (o *ConfigurationServerConfiguration) HasUninstalledPlugins() bool`

HasUninstalledPlugins returns a boolean if a field has been set.

### GetCollapseVideoFolders

`func (o *ConfigurationServerConfiguration) GetCollapseVideoFolders() bool`

GetCollapseVideoFolders returns the CollapseVideoFolders field if non-nil, zero value otherwise.

### GetCollapseVideoFoldersOk

`func (o *ConfigurationServerConfiguration) GetCollapseVideoFoldersOk() (*bool, bool)`

GetCollapseVideoFoldersOk returns a tuple with the CollapseVideoFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseVideoFolders

`func (o *ConfigurationServerConfiguration) SetCollapseVideoFolders(v bool)`

SetCollapseVideoFolders sets CollapseVideoFolders field to given value.

### HasCollapseVideoFolders

`func (o *ConfigurationServerConfiguration) HasCollapseVideoFolders() bool`

HasCollapseVideoFolders returns a boolean if a field has been set.

### GetEnableOriginalTrackTitles

`func (o *ConfigurationServerConfiguration) GetEnableOriginalTrackTitles() bool`

GetEnableOriginalTrackTitles returns the EnableOriginalTrackTitles field if non-nil, zero value otherwise.

### GetEnableOriginalTrackTitlesOk

`func (o *ConfigurationServerConfiguration) GetEnableOriginalTrackTitlesOk() (*bool, bool)`

GetEnableOriginalTrackTitlesOk returns a tuple with the EnableOriginalTrackTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalTrackTitles

`func (o *ConfigurationServerConfiguration) SetEnableOriginalTrackTitles(v bool)`

SetEnableOriginalTrackTitles sets EnableOriginalTrackTitles field to given value.

### HasEnableOriginalTrackTitles

`func (o *ConfigurationServerConfiguration) HasEnableOriginalTrackTitles() bool`

HasEnableOriginalTrackTitles returns a boolean if a field has been set.

### GetEnableDebugLevelLogging

`func (o *ConfigurationServerConfiguration) GetEnableDebugLevelLogging() bool`

GetEnableDebugLevelLogging returns the EnableDebugLevelLogging field if non-nil, zero value otherwise.

### GetEnableDebugLevelLoggingOk

`func (o *ConfigurationServerConfiguration) GetEnableDebugLevelLoggingOk() (*bool, bool)`

GetEnableDebugLevelLoggingOk returns a tuple with the EnableDebugLevelLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLevelLogging

`func (o *ConfigurationServerConfiguration) SetEnableDebugLevelLogging(v bool)`

SetEnableDebugLevelLogging sets EnableDebugLevelLogging field to given value.

### HasEnableDebugLevelLogging

`func (o *ConfigurationServerConfiguration) HasEnableDebugLevelLogging() bool`

HasEnableDebugLevelLogging returns a boolean if a field has been set.

### GetEnableAutoUpdate

`func (o *ConfigurationServerConfiguration) GetEnableAutoUpdate() bool`

GetEnableAutoUpdate returns the EnableAutoUpdate field if non-nil, zero value otherwise.

### GetEnableAutoUpdateOk

`func (o *ConfigurationServerConfiguration) GetEnableAutoUpdateOk() (*bool, bool)`

GetEnableAutoUpdateOk returns a tuple with the EnableAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdate

`func (o *ConfigurationServerConfiguration) SetEnableAutoUpdate(v bool)`

SetEnableAutoUpdate sets EnableAutoUpdate field to given value.

### HasEnableAutoUpdate

`func (o *ConfigurationServerConfiguration) HasEnableAutoUpdate() bool`

HasEnableAutoUpdate returns a boolean if a field has been set.

### GetLogFileRetentionDays

`func (o *ConfigurationServerConfiguration) GetLogFileRetentionDays() int32`

GetLogFileRetentionDays returns the LogFileRetentionDays field if non-nil, zero value otherwise.

### GetLogFileRetentionDaysOk

`func (o *ConfigurationServerConfiguration) GetLogFileRetentionDaysOk() (*int32, bool)`

GetLogFileRetentionDaysOk returns a tuple with the LogFileRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileRetentionDays

`func (o *ConfigurationServerConfiguration) SetLogFileRetentionDays(v int32)`

SetLogFileRetentionDays sets LogFileRetentionDays field to given value.

### HasLogFileRetentionDays

`func (o *ConfigurationServerConfiguration) HasLogFileRetentionDays() bool`

HasLogFileRetentionDays returns a boolean if a field has been set.

### GetRunAtStartup

`func (o *ConfigurationServerConfiguration) GetRunAtStartup() bool`

GetRunAtStartup returns the RunAtStartup field if non-nil, zero value otherwise.

### GetRunAtStartupOk

`func (o *ConfigurationServerConfiguration) GetRunAtStartupOk() (*bool, bool)`

GetRunAtStartupOk returns a tuple with the RunAtStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAtStartup

`func (o *ConfigurationServerConfiguration) SetRunAtStartup(v bool)`

SetRunAtStartup sets RunAtStartup field to given value.

### HasRunAtStartup

`func (o *ConfigurationServerConfiguration) HasRunAtStartup() bool`

HasRunAtStartup returns a boolean if a field has been set.

### GetIsStartupWizardCompleted

`func (o *ConfigurationServerConfiguration) GetIsStartupWizardCompleted() bool`

GetIsStartupWizardCompleted returns the IsStartupWizardCompleted field if non-nil, zero value otherwise.

### GetIsStartupWizardCompletedOk

`func (o *ConfigurationServerConfiguration) GetIsStartupWizardCompletedOk() (*bool, bool)`

GetIsStartupWizardCompletedOk returns a tuple with the IsStartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartupWizardCompleted

`func (o *ConfigurationServerConfiguration) SetIsStartupWizardCompleted(v bool)`

SetIsStartupWizardCompleted sets IsStartupWizardCompleted field to given value.

### HasIsStartupWizardCompleted

`func (o *ConfigurationServerConfiguration) HasIsStartupWizardCompleted() bool`

HasIsStartupWizardCompleted returns a boolean if a field has been set.

### GetCachePath

`func (o *ConfigurationServerConfiguration) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ConfigurationServerConfiguration) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ConfigurationServerConfiguration) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ConfigurationServerConfiguration) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DlnaDeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identification** | Pointer to [**DlnaDeviceIdentification**](DlnaDeviceIdentification.md) |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerUrl** | Pointer to **string** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**ModelDescription** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**ModelUrl** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**EnableAlbumArtInDidl** | Pointer to **bool** |  | [optional] 
**EnableSingleAlbumArtLimit** | Pointer to **bool** |  | [optional] 
**EnableSingleSubtitleLimit** | Pointer to **bool** |  | [optional] 
**SupportedMediaTypes** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AlbumArtPn** | Pointer to **string** |  | [optional] 
**MaxAlbumArtWidth** | Pointer to **int32** |  | [optional] 
**MaxAlbumArtHeight** | Pointer to **int32** |  | [optional] 
**MaxIconWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxIconHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**MaxStaticBitrate** | Pointer to **NullableInt64** |  | [optional] 
**MusicStreamingTranscodingBitrate** | Pointer to **NullableInt32** |  | [optional] 
**MaxStaticMusicBitrate** | Pointer to **NullableInt32** |  | [optional] 
**SonyAggregationFlags** | Pointer to **string** |  | [optional] 
**ProtocolInfo** | Pointer to **string** |  | [optional] 
**TimelineOffsetSeconds** | Pointer to **int32** |  | [optional] 
**RequiresPlainVideoItems** | Pointer to **bool** |  | [optional] 
**RequiresPlainFolders** | Pointer to **bool** |  | [optional] 
**EnableMSMediaReceiverRegistrar** | Pointer to **bool** |  | [optional] 
**IgnoreTranscodeByteRangeRequests** | Pointer to **bool** |  | [optional] 
**XmlRootAttributes** | Pointer to [**[]DlnaXmlAttribute**](DlnaXmlAttribute.md) |  | [optional] 
**DirectPlayProfiles** | Pointer to [**[]DlnaDirectPlayProfile**](DlnaDirectPlayProfile.md) |  | [optional] 
**TranscodingProfiles** | Pointer to [**[]DlnaTranscodingProfile**](DlnaTranscodingProfile.md) |  | [optional] 
**ContainerProfiles** | Pointer to [**[]DlnaContainerProfile**](DlnaContainerProfile.md) |  | [optional] 
**CodecProfiles** | Pointer to [**[]DlnaCodecProfile**](DlnaCodecProfile.md) |  | [optional] 
**ResponseProfiles** | Pointer to [**[]DlnaResponseProfile**](DlnaResponseProfile.md) |  | [optional] 
**SubtitleProfiles** | Pointer to [**[]DlnaSubtitleProfile**](DlnaSubtitleProfile.md) |  | [optional] 

## Methods

### NewDlnaDeviceProfile

`func NewDlnaDeviceProfile() *DlnaDeviceProfile`

NewDlnaDeviceProfile instantiates a new DlnaDeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaDeviceProfileWithDefaults

`func NewDlnaDeviceProfileWithDefaults() *DlnaDeviceProfile`

NewDlnaDeviceProfileWithDefaults instantiates a new DlnaDeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DlnaDeviceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DlnaDeviceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DlnaDeviceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DlnaDeviceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *DlnaDeviceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DlnaDeviceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DlnaDeviceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DlnaDeviceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentification

`func (o *DlnaDeviceProfile) GetIdentification() DlnaDeviceIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *DlnaDeviceProfile) GetIdentificationOk() (*DlnaDeviceIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *DlnaDeviceProfile) SetIdentification(v DlnaDeviceIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *DlnaDeviceProfile) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetFriendlyName

`func (o *DlnaDeviceProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DlnaDeviceProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DlnaDeviceProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DlnaDeviceProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetManufacturer

`func (o *DlnaDeviceProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DlnaDeviceProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DlnaDeviceProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DlnaDeviceProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerUrl

`func (o *DlnaDeviceProfile) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *DlnaDeviceProfile) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *DlnaDeviceProfile) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *DlnaDeviceProfile) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### GetModelName

`func (o *DlnaDeviceProfile) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DlnaDeviceProfile) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DlnaDeviceProfile) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DlnaDeviceProfile) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelDescription

`func (o *DlnaDeviceProfile) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *DlnaDeviceProfile) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *DlnaDeviceProfile) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *DlnaDeviceProfile) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### GetModelNumber

`func (o *DlnaDeviceProfile) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DlnaDeviceProfile) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DlnaDeviceProfile) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DlnaDeviceProfile) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetModelUrl

`func (o *DlnaDeviceProfile) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *DlnaDeviceProfile) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *DlnaDeviceProfile) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *DlnaDeviceProfile) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DlnaDeviceProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DlnaDeviceProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DlnaDeviceProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DlnaDeviceProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetEnableAlbumArtInDidl

`func (o *DlnaDeviceProfile) GetEnableAlbumArtInDidl() bool`

GetEnableAlbumArtInDidl returns the EnableAlbumArtInDidl field if non-nil, zero value otherwise.

### GetEnableAlbumArtInDidlOk

`func (o *DlnaDeviceProfile) GetEnableAlbumArtInDidlOk() (*bool, bool)`

GetEnableAlbumArtInDidlOk returns a tuple with the EnableAlbumArtInDidl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlbumArtInDidl

`func (o *DlnaDeviceProfile) SetEnableAlbumArtInDidl(v bool)`

SetEnableAlbumArtInDidl sets EnableAlbumArtInDidl field to given value.

### HasEnableAlbumArtInDidl

`func (o *DlnaDeviceProfile) HasEnableAlbumArtInDidl() bool`

HasEnableAlbumArtInDidl returns a boolean if a field has been set.

### GetEnableSingleAlbumArtLimit

`func (o *DlnaDeviceProfile) GetEnableSingleAlbumArtLimit() bool`

GetEnableSingleAlbumArtLimit returns the EnableSingleAlbumArtLimit field if non-nil, zero value otherwise.

### GetEnableSingleAlbumArtLimitOk

`func (o *DlnaDeviceProfile) GetEnableSingleAlbumArtLimitOk() (*bool, bool)`

GetEnableSingleAlbumArtLimitOk returns a tuple with the EnableSingleAlbumArtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleAlbumArtLimit

`func (o *DlnaDeviceProfile) SetEnableSingleAlbumArtLimit(v bool)`

SetEnableSingleAlbumArtLimit sets EnableSingleAlbumArtLimit field to given value.

### HasEnableSingleAlbumArtLimit

`func (o *DlnaDeviceProfile) HasEnableSingleAlbumArtLimit() bool`

HasEnableSingleAlbumArtLimit returns a boolean if a field has been set.

### GetEnableSingleSubtitleLimit

`func (o *DlnaDeviceProfile) GetEnableSingleSubtitleLimit() bool`

GetEnableSingleSubtitleLimit returns the EnableSingleSubtitleLimit field if non-nil, zero value otherwise.

### GetEnableSingleSubtitleLimitOk

`func (o *DlnaDeviceProfile) GetEnableSingleSubtitleLimitOk() (*bool, bool)`

GetEnableSingleSubtitleLimitOk returns a tuple with the EnableSingleSubtitleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSubtitleLimit

`func (o *DlnaDeviceProfile) SetEnableSingleSubtitleLimit(v bool)`

SetEnableSingleSubtitleLimit sets EnableSingleSubtitleLimit field to given value.

### HasEnableSingleSubtitleLimit

`func (o *DlnaDeviceProfile) HasEnableSingleSubtitleLimit() bool`

HasEnableSingleSubtitleLimit returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *DlnaDeviceProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *DlnaDeviceProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *DlnaDeviceProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *DlnaDeviceProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetUserId

`func (o *DlnaDeviceProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DlnaDeviceProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DlnaDeviceProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DlnaDeviceProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAlbumArtPn

`func (o *DlnaDeviceProfile) GetAlbumArtPn() string`

GetAlbumArtPn returns the AlbumArtPn field if non-nil, zero value otherwise.

### GetAlbumArtPnOk

`func (o *DlnaDeviceProfile) GetAlbumArtPnOk() (*string, bool)`

GetAlbumArtPnOk returns a tuple with the AlbumArtPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtPn

`func (o *DlnaDeviceProfile) SetAlbumArtPn(v string)`

SetAlbumArtPn sets AlbumArtPn field to given value.

### HasAlbumArtPn

`func (o *DlnaDeviceProfile) HasAlbumArtPn() bool`

HasAlbumArtPn returns a boolean if a field has been set.

### GetMaxAlbumArtWidth

`func (o *DlnaDeviceProfile) GetMaxAlbumArtWidth() int32`

GetMaxAlbumArtWidth returns the MaxAlbumArtWidth field if non-nil, zero value otherwise.

### GetMaxAlbumArtWidthOk

`func (o *DlnaDeviceProfile) GetMaxAlbumArtWidthOk() (*int32, bool)`

GetMaxAlbumArtWidthOk returns a tuple with the MaxAlbumArtWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtWidth

`func (o *DlnaDeviceProfile) SetMaxAlbumArtWidth(v int32)`

SetMaxAlbumArtWidth sets MaxAlbumArtWidth field to given value.

### HasMaxAlbumArtWidth

`func (o *DlnaDeviceProfile) HasMaxAlbumArtWidth() bool`

HasMaxAlbumArtWidth returns a boolean if a field has been set.

### GetMaxAlbumArtHeight

`func (o *DlnaDeviceProfile) GetMaxAlbumArtHeight() int32`

GetMaxAlbumArtHeight returns the MaxAlbumArtHeight field if non-nil, zero value otherwise.

### GetMaxAlbumArtHeightOk

`func (o *DlnaDeviceProfile) GetMaxAlbumArtHeightOk() (*int32, bool)`

GetMaxAlbumArtHeightOk returns a tuple with the MaxAlbumArtHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtHeight

`func (o *DlnaDeviceProfile) SetMaxAlbumArtHeight(v int32)`

SetMaxAlbumArtHeight sets MaxAlbumArtHeight field to given value.

### HasMaxAlbumArtHeight

`func (o *DlnaDeviceProfile) HasMaxAlbumArtHeight() bool`

HasMaxAlbumArtHeight returns a boolean if a field has been set.

### GetMaxIconWidth

`func (o *DlnaDeviceProfile) GetMaxIconWidth() int32`

GetMaxIconWidth returns the MaxIconWidth field if non-nil, zero value otherwise.

### GetMaxIconWidthOk

`func (o *DlnaDeviceProfile) GetMaxIconWidthOk() (*int32, bool)`

GetMaxIconWidthOk returns a tuple with the MaxIconWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconWidth

`func (o *DlnaDeviceProfile) SetMaxIconWidth(v int32)`

SetMaxIconWidth sets MaxIconWidth field to given value.

### HasMaxIconWidth

`func (o *DlnaDeviceProfile) HasMaxIconWidth() bool`

HasMaxIconWidth returns a boolean if a field has been set.

### SetMaxIconWidthNil

`func (o *DlnaDeviceProfile) SetMaxIconWidthNil(b bool)`

 SetMaxIconWidthNil sets the value for MaxIconWidth to be an explicit nil

### UnsetMaxIconWidth
`func (o *DlnaDeviceProfile) UnsetMaxIconWidth()`

UnsetMaxIconWidth ensures that no value is present for MaxIconWidth, not even an explicit nil
### GetMaxIconHeight

`func (o *DlnaDeviceProfile) GetMaxIconHeight() int32`

GetMaxIconHeight returns the MaxIconHeight field if non-nil, zero value otherwise.

### GetMaxIconHeightOk

`func (o *DlnaDeviceProfile) GetMaxIconHeightOk() (*int32, bool)`

GetMaxIconHeightOk returns a tuple with the MaxIconHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconHeight

`func (o *DlnaDeviceProfile) SetMaxIconHeight(v int32)`

SetMaxIconHeight sets MaxIconHeight field to given value.

### HasMaxIconHeight

`func (o *DlnaDeviceProfile) HasMaxIconHeight() bool`

HasMaxIconHeight returns a boolean if a field has been set.

### SetMaxIconHeightNil

`func (o *DlnaDeviceProfile) SetMaxIconHeightNil(b bool)`

 SetMaxIconHeightNil sets the value for MaxIconHeight to be an explicit nil

### UnsetMaxIconHeight
`func (o *DlnaDeviceProfile) UnsetMaxIconHeight()`

UnsetMaxIconHeight ensures that no value is present for MaxIconHeight, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *DlnaDeviceProfile) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *DlnaDeviceProfile) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *DlnaDeviceProfile) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *DlnaDeviceProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *DlnaDeviceProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *DlnaDeviceProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMaxStaticBitrate

`func (o *DlnaDeviceProfile) GetMaxStaticBitrate() int64`

GetMaxStaticBitrate returns the MaxStaticBitrate field if non-nil, zero value otherwise.

### GetMaxStaticBitrateOk

`func (o *DlnaDeviceProfile) GetMaxStaticBitrateOk() (*int64, bool)`

GetMaxStaticBitrateOk returns a tuple with the MaxStaticBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticBitrate

`func (o *DlnaDeviceProfile) SetMaxStaticBitrate(v int64)`

SetMaxStaticBitrate sets MaxStaticBitrate field to given value.

### HasMaxStaticBitrate

`func (o *DlnaDeviceProfile) HasMaxStaticBitrate() bool`

HasMaxStaticBitrate returns a boolean if a field has been set.

### SetMaxStaticBitrateNil

`func (o *DlnaDeviceProfile) SetMaxStaticBitrateNil(b bool)`

 SetMaxStaticBitrateNil sets the value for MaxStaticBitrate to be an explicit nil

### UnsetMaxStaticBitrate
`func (o *DlnaDeviceProfile) UnsetMaxStaticBitrate()`

UnsetMaxStaticBitrate ensures that no value is present for MaxStaticBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *DlnaDeviceProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *DlnaDeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *DlnaDeviceProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *DlnaDeviceProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *DlnaDeviceProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *DlnaDeviceProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *DlnaDeviceProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *DlnaDeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *DlnaDeviceProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *DlnaDeviceProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *DlnaDeviceProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *DlnaDeviceProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetSonyAggregationFlags

`func (o *DlnaDeviceProfile) GetSonyAggregationFlags() string`

GetSonyAggregationFlags returns the SonyAggregationFlags field if non-nil, zero value otherwise.

### GetSonyAggregationFlagsOk

`func (o *DlnaDeviceProfile) GetSonyAggregationFlagsOk() (*string, bool)`

GetSonyAggregationFlagsOk returns a tuple with the SonyAggregationFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonyAggregationFlags

`func (o *DlnaDeviceProfile) SetSonyAggregationFlags(v string)`

SetSonyAggregationFlags sets SonyAggregationFlags field to given value.

### HasSonyAggregationFlags

`func (o *DlnaDeviceProfile) HasSonyAggregationFlags() bool`

HasSonyAggregationFlags returns a boolean if a field has been set.

### GetProtocolInfo

`func (o *DlnaDeviceProfile) GetProtocolInfo() string`

GetProtocolInfo returns the ProtocolInfo field if non-nil, zero value otherwise.

### GetProtocolInfoOk

`func (o *DlnaDeviceProfile) GetProtocolInfoOk() (*string, bool)`

GetProtocolInfoOk returns a tuple with the ProtocolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfo

`func (o *DlnaDeviceProfile) SetProtocolInfo(v string)`

SetProtocolInfo sets ProtocolInfo field to given value.

### HasProtocolInfo

`func (o *DlnaDeviceProfile) HasProtocolInfo() bool`

HasProtocolInfo returns a boolean if a field has been set.

### GetTimelineOffsetSeconds

`func (o *DlnaDeviceProfile) GetTimelineOffsetSeconds() int32`

GetTimelineOffsetSeconds returns the TimelineOffsetSeconds field if non-nil, zero value otherwise.

### GetTimelineOffsetSecondsOk

`func (o *DlnaDeviceProfile) GetTimelineOffsetSecondsOk() (*int32, bool)`

GetTimelineOffsetSecondsOk returns a tuple with the TimelineOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineOffsetSeconds

`func (o *DlnaDeviceProfile) SetTimelineOffsetSeconds(v int32)`

SetTimelineOffsetSeconds sets TimelineOffsetSeconds field to given value.

### HasTimelineOffsetSeconds

`func (o *DlnaDeviceProfile) HasTimelineOffsetSeconds() bool`

HasTimelineOffsetSeconds returns a boolean if a field has been set.

### GetRequiresPlainVideoItems

`func (o *DlnaDeviceProfile) GetRequiresPlainVideoItems() bool`

GetRequiresPlainVideoItems returns the RequiresPlainVideoItems field if non-nil, zero value otherwise.

### GetRequiresPlainVideoItemsOk

`func (o *DlnaDeviceProfile) GetRequiresPlainVideoItemsOk() (*bool, bool)`

GetRequiresPlainVideoItemsOk returns a tuple with the RequiresPlainVideoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainVideoItems

`func (o *DlnaDeviceProfile) SetRequiresPlainVideoItems(v bool)`

SetRequiresPlainVideoItems sets RequiresPlainVideoItems field to given value.

### HasRequiresPlainVideoItems

`func (o *DlnaDeviceProfile) HasRequiresPlainVideoItems() bool`

HasRequiresPlainVideoItems returns a boolean if a field has been set.

### GetRequiresPlainFolders

`func (o *DlnaDeviceProfile) GetRequiresPlainFolders() bool`

GetRequiresPlainFolders returns the RequiresPlainFolders field if non-nil, zero value otherwise.

### GetRequiresPlainFoldersOk

`func (o *DlnaDeviceProfile) GetRequiresPlainFoldersOk() (*bool, bool)`

GetRequiresPlainFoldersOk returns a tuple with the RequiresPlainFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainFolders

`func (o *DlnaDeviceProfile) SetRequiresPlainFolders(v bool)`

SetRequiresPlainFolders sets RequiresPlainFolders field to given value.

### HasRequiresPlainFolders

`func (o *DlnaDeviceProfile) HasRequiresPlainFolders() bool`

HasRequiresPlainFolders returns a boolean if a field has been set.

### GetEnableMSMediaReceiverRegistrar

`func (o *DlnaDeviceProfile) GetEnableMSMediaReceiverRegistrar() bool`

GetEnableMSMediaReceiverRegistrar returns the EnableMSMediaReceiverRegistrar field if non-nil, zero value otherwise.

### GetEnableMSMediaReceiverRegistrarOk

`func (o *DlnaDeviceProfile) GetEnableMSMediaReceiverRegistrarOk() (*bool, bool)`

GetEnableMSMediaReceiverRegistrarOk returns a tuple with the EnableMSMediaReceiverRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMSMediaReceiverRegistrar

`func (o *DlnaDeviceProfile) SetEnableMSMediaReceiverRegistrar(v bool)`

SetEnableMSMediaReceiverRegistrar sets EnableMSMediaReceiverRegistrar field to given value.

### HasEnableMSMediaReceiverRegistrar

`func (o *DlnaDeviceProfile) HasEnableMSMediaReceiverRegistrar() bool`

HasEnableMSMediaReceiverRegistrar returns a boolean if a field has been set.

### GetIgnoreTranscodeByteRangeRequests

`func (o *DlnaDeviceProfile) GetIgnoreTranscodeByteRangeRequests() bool`

GetIgnoreTranscodeByteRangeRequests returns the IgnoreTranscodeByteRangeRequests field if non-nil, zero value otherwise.

### GetIgnoreTranscodeByteRangeRequestsOk

`func (o *DlnaDeviceProfile) GetIgnoreTranscodeByteRangeRequestsOk() (*bool, bool)`

GetIgnoreTranscodeByteRangeRequestsOk returns a tuple with the IgnoreTranscodeByteRangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTranscodeByteRangeRequests

`func (o *DlnaDeviceProfile) SetIgnoreTranscodeByteRangeRequests(v bool)`

SetIgnoreTranscodeByteRangeRequests sets IgnoreTranscodeByteRangeRequests field to given value.

### HasIgnoreTranscodeByteRangeRequests

`func (o *DlnaDeviceProfile) HasIgnoreTranscodeByteRangeRequests() bool`

HasIgnoreTranscodeByteRangeRequests returns a boolean if a field has been set.

### GetXmlRootAttributes

`func (o *DlnaDeviceProfile) GetXmlRootAttributes() []DlnaXmlAttribute`

GetXmlRootAttributes returns the XmlRootAttributes field if non-nil, zero value otherwise.

### GetXmlRootAttributesOk

`func (o *DlnaDeviceProfile) GetXmlRootAttributesOk() (*[]DlnaXmlAttribute, bool)`

GetXmlRootAttributesOk returns a tuple with the XmlRootAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlRootAttributes

`func (o *DlnaDeviceProfile) SetXmlRootAttributes(v []DlnaXmlAttribute)`

SetXmlRootAttributes sets XmlRootAttributes field to given value.

### HasXmlRootAttributes

`func (o *DlnaDeviceProfile) HasXmlRootAttributes() bool`

HasXmlRootAttributes returns a boolean if a field has been set.

### GetDirectPlayProfiles

`func (o *DlnaDeviceProfile) GetDirectPlayProfiles() []DlnaDirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *DlnaDeviceProfile) GetDirectPlayProfilesOk() (*[]DlnaDirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *DlnaDeviceProfile) SetDirectPlayProfiles(v []DlnaDirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *DlnaDeviceProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *DlnaDeviceProfile) GetTranscodingProfiles() []DlnaTranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *DlnaDeviceProfile) GetTranscodingProfilesOk() (*[]DlnaTranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *DlnaDeviceProfile) SetTranscodingProfiles(v []DlnaTranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *DlnaDeviceProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *DlnaDeviceProfile) GetContainerProfiles() []DlnaContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *DlnaDeviceProfile) GetContainerProfilesOk() (*[]DlnaContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *DlnaDeviceProfile) SetContainerProfiles(v []DlnaContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *DlnaDeviceProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *DlnaDeviceProfile) GetCodecProfiles() []DlnaCodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *DlnaDeviceProfile) GetCodecProfilesOk() (*[]DlnaCodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *DlnaDeviceProfile) SetCodecProfiles(v []DlnaCodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *DlnaDeviceProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *DlnaDeviceProfile) GetResponseProfiles() []DlnaResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *DlnaDeviceProfile) GetResponseProfilesOk() (*[]DlnaResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *DlnaDeviceProfile) SetResponseProfiles(v []DlnaResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *DlnaDeviceProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *DlnaDeviceProfile) GetSubtitleProfiles() []DlnaSubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *DlnaDeviceProfile) GetSubtitleProfilesOk() (*[]DlnaSubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *DlnaDeviceProfile) SetSubtitleProfiles(v []DlnaSubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *DlnaDeviceProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MediaStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to **string** |  | [optional] 
**CodecTag** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**ColorTransfer** | Pointer to **string** |  | [optional] 
**ColorPrimaries** | Pointer to **string** |  | [optional] 
**ColorSpace** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**TimeBase** | Pointer to **string** |  | [optional] 
**CodecTimeBase** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Extradata** | Pointer to **string** |  | [optional] 
**VideoRange** | Pointer to **string** |  | [optional] 
**DisplayTitle** | Pointer to **string** |  | [optional] 
**DisplayLanguage** | Pointer to **string** |  | [optional] 
**NalLengthSize** | Pointer to **string** |  | [optional] 
**IsInterlaced** | Pointer to **bool** |  | [optional] 
**IsAVC** | Pointer to **NullableBool** |  | [optional] 
**ChannelLayout** | Pointer to **string** |  | [optional] 
**BitRate** | Pointer to **NullableInt32** |  | [optional] 
**BitDepth** | Pointer to **NullableInt32** |  | [optional] 
**RefFrames** | Pointer to **NullableInt32** |  | [optional] 
**PacketLength** | Pointer to **NullableInt32** |  | [optional] 
**Channels** | Pointer to **NullableInt32** |  | [optional] 
**SampleRate** | Pointer to **NullableInt32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsForced** | Pointer to **bool** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**AverageFrameRate** | Pointer to **NullableFloat32** |  | [optional] 
**RealFrameRate** | Pointer to **NullableFloat32** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AspectRatio** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **NullableInt32** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**DeliveryMethod** | Pointer to **string** |  | [optional] 
**DeliveryUrl** | Pointer to **string** |  | [optional] 
**IsExternalUrl** | Pointer to **NullableBool** |  | [optional] 
**IsTextSubtitleStream** | Pointer to **bool** |  | [optional] 
**SupportsExternalStream** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PixelFormat** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **NullableFloat64** |  | [optional] 
**IsAnamorphic** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMediaStream

`func NewMediaStream() *MediaStream`

NewMediaStream instantiates a new MediaStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStreamWithDefaults

`func NewMediaStreamWithDefaults() *MediaStream`

NewMediaStreamWithDefaults instantiates a new MediaStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *MediaStream) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *MediaStream) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *MediaStream) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *MediaStream) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetCodecTag

`func (o *MediaStream) GetCodecTag() string`

GetCodecTag returns the CodecTag field if non-nil, zero value otherwise.

### GetCodecTagOk

`func (o *MediaStream) GetCodecTagOk() (*string, bool)`

GetCodecTagOk returns a tuple with the CodecTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTag

`func (o *MediaStream) SetCodecTag(v string)`

SetCodecTag sets CodecTag field to given value.

### HasCodecTag

`func (o *MediaStream) HasCodecTag() bool`

HasCodecTag returns a boolean if a field has been set.

### GetLanguage

`func (o *MediaStream) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MediaStream) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MediaStream) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MediaStream) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetColorTransfer

`func (o *MediaStream) GetColorTransfer() string`

GetColorTransfer returns the ColorTransfer field if non-nil, zero value otherwise.

### GetColorTransferOk

`func (o *MediaStream) GetColorTransferOk() (*string, bool)`

GetColorTransferOk returns a tuple with the ColorTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorTransfer

`func (o *MediaStream) SetColorTransfer(v string)`

SetColorTransfer sets ColorTransfer field to given value.

### HasColorTransfer

`func (o *MediaStream) HasColorTransfer() bool`

HasColorTransfer returns a boolean if a field has been set.

### GetColorPrimaries

`func (o *MediaStream) GetColorPrimaries() string`

GetColorPrimaries returns the ColorPrimaries field if non-nil, zero value otherwise.

### GetColorPrimariesOk

`func (o *MediaStream) GetColorPrimariesOk() (*string, bool)`

GetColorPrimariesOk returns a tuple with the ColorPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPrimaries

`func (o *MediaStream) SetColorPrimaries(v string)`

SetColorPrimaries sets ColorPrimaries field to given value.

### HasColorPrimaries

`func (o *MediaStream) HasColorPrimaries() bool`

HasColorPrimaries returns a boolean if a field has been set.

### GetColorSpace

`func (o *MediaStream) GetColorSpace() string`

GetColorSpace returns the ColorSpace field if non-nil, zero value otherwise.

### GetColorSpaceOk

`func (o *MediaStream) GetColorSpaceOk() (*string, bool)`

GetColorSpaceOk returns a tuple with the ColorSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorSpace

`func (o *MediaStream) SetColorSpace(v string)`

SetColorSpace sets ColorSpace field to given value.

### HasColorSpace

`func (o *MediaStream) HasColorSpace() bool`

HasColorSpace returns a boolean if a field has been set.

### GetComment

`func (o *MediaStream) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MediaStream) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MediaStream) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MediaStream) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTimeBase

`func (o *MediaStream) GetTimeBase() string`

GetTimeBase returns the TimeBase field if non-nil, zero value otherwise.

### GetTimeBaseOk

`func (o *MediaStream) GetTimeBaseOk() (*string, bool)`

GetTimeBaseOk returns a tuple with the TimeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBase

`func (o *MediaStream) SetTimeBase(v string)`

SetTimeBase sets TimeBase field to given value.

### HasTimeBase

`func (o *MediaStream) HasTimeBase() bool`

HasTimeBase returns a boolean if a field has been set.

### GetCodecTimeBase

`func (o *MediaStream) GetCodecTimeBase() string`

GetCodecTimeBase returns the CodecTimeBase field if non-nil, zero value otherwise.

### GetCodecTimeBaseOk

`func (o *MediaStream) GetCodecTimeBaseOk() (*string, bool)`

GetCodecTimeBaseOk returns a tuple with the CodecTimeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTimeBase

`func (o *MediaStream) SetCodecTimeBase(v string)`

SetCodecTimeBase sets CodecTimeBase field to given value.

### HasCodecTimeBase

`func (o *MediaStream) HasCodecTimeBase() bool`

HasCodecTimeBase returns a boolean if a field has been set.

### GetTitle

`func (o *MediaStream) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MediaStream) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MediaStream) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MediaStream) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetExtradata

`func (o *MediaStream) GetExtradata() string`

GetExtradata returns the Extradata field if non-nil, zero value otherwise.

### GetExtradataOk

`func (o *MediaStream) GetExtradataOk() (*string, bool)`

GetExtradataOk returns a tuple with the Extradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtradata

`func (o *MediaStream) SetExtradata(v string)`

SetExtradata sets Extradata field to given value.

### HasExtradata

`func (o *MediaStream) HasExtradata() bool`

HasExtradata returns a boolean if a field has been set.

### GetVideoRange

`func (o *MediaStream) GetVideoRange() string`

GetVideoRange returns the VideoRange field if non-nil, zero value otherwise.

### GetVideoRangeOk

`func (o *MediaStream) GetVideoRangeOk() (*string, bool)`

GetVideoRangeOk returns a tuple with the VideoRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRange

`func (o *MediaStream) SetVideoRange(v string)`

SetVideoRange sets VideoRange field to given value.

### HasVideoRange

`func (o *MediaStream) HasVideoRange() bool`

HasVideoRange returns a boolean if a field has been set.

### GetDisplayTitle

`func (o *MediaStream) GetDisplayTitle() string`

GetDisplayTitle returns the DisplayTitle field if non-nil, zero value otherwise.

### GetDisplayTitleOk

`func (o *MediaStream) GetDisplayTitleOk() (*string, bool)`

GetDisplayTitleOk returns a tuple with the DisplayTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTitle

`func (o *MediaStream) SetDisplayTitle(v string)`

SetDisplayTitle sets DisplayTitle field to given value.

### HasDisplayTitle

`func (o *MediaStream) HasDisplayTitle() bool`

HasDisplayTitle returns a boolean if a field has been set.

### GetDisplayLanguage

`func (o *MediaStream) GetDisplayLanguage() string`

GetDisplayLanguage returns the DisplayLanguage field if non-nil, zero value otherwise.

### GetDisplayLanguageOk

`func (o *MediaStream) GetDisplayLanguageOk() (*string, bool)`

GetDisplayLanguageOk returns a tuple with the DisplayLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLanguage

`func (o *MediaStream) SetDisplayLanguage(v string)`

SetDisplayLanguage sets DisplayLanguage field to given value.

### HasDisplayLanguage

`func (o *MediaStream) HasDisplayLanguage() bool`

HasDisplayLanguage returns a boolean if a field has been set.

### GetNalLengthSize

`func (o *MediaStream) GetNalLengthSize() string`

GetNalLengthSize returns the NalLengthSize field if non-nil, zero value otherwise.

### GetNalLengthSizeOk

`func (o *MediaStream) GetNalLengthSizeOk() (*string, bool)`

GetNalLengthSizeOk returns a tuple with the NalLengthSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNalLengthSize

`func (o *MediaStream) SetNalLengthSize(v string)`

SetNalLengthSize sets NalLengthSize field to given value.

### HasNalLengthSize

`func (o *MediaStream) HasNalLengthSize() bool`

HasNalLengthSize returns a boolean if a field has been set.

### GetIsInterlaced

`func (o *MediaStream) GetIsInterlaced() bool`

GetIsInterlaced returns the IsInterlaced field if non-nil, zero value otherwise.

### GetIsInterlacedOk

`func (o *MediaStream) GetIsInterlacedOk() (*bool, bool)`

GetIsInterlacedOk returns a tuple with the IsInterlaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterlaced

`func (o *MediaStream) SetIsInterlaced(v bool)`

SetIsInterlaced sets IsInterlaced field to given value.

### HasIsInterlaced

`func (o *MediaStream) HasIsInterlaced() bool`

HasIsInterlaced returns a boolean if a field has been set.

### GetIsAVC

`func (o *MediaStream) GetIsAVC() bool`

GetIsAVC returns the IsAVC field if non-nil, zero value otherwise.

### GetIsAVCOk

`func (o *MediaStream) GetIsAVCOk() (*bool, bool)`

GetIsAVCOk returns a tuple with the IsAVC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAVC

`func (o *MediaStream) SetIsAVC(v bool)`

SetIsAVC sets IsAVC field to given value.

### HasIsAVC

`func (o *MediaStream) HasIsAVC() bool`

HasIsAVC returns a boolean if a field has been set.

### SetIsAVCNil

`func (o *MediaStream) SetIsAVCNil(b bool)`

 SetIsAVCNil sets the value for IsAVC to be an explicit nil

### UnsetIsAVC
`func (o *MediaStream) UnsetIsAVC()`

UnsetIsAVC ensures that no value is present for IsAVC, not even an explicit nil
### GetChannelLayout

`func (o *MediaStream) GetChannelLayout() string`

GetChannelLayout returns the ChannelLayout field if non-nil, zero value otherwise.

### GetChannelLayoutOk

`func (o *MediaStream) GetChannelLayoutOk() (*string, bool)`

GetChannelLayoutOk returns a tuple with the ChannelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLayout

`func (o *MediaStream) SetChannelLayout(v string)`

SetChannelLayout sets ChannelLayout field to given value.

### HasChannelLayout

`func (o *MediaStream) HasChannelLayout() bool`

HasChannelLayout returns a boolean if a field has been set.

### GetBitRate

`func (o *MediaStream) GetBitRate() int32`

GetBitRate returns the BitRate field if non-nil, zero value otherwise.

### GetBitRateOk

`func (o *MediaStream) GetBitRateOk() (*int32, bool)`

GetBitRateOk returns a tuple with the BitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRate

`func (o *MediaStream) SetBitRate(v int32)`

SetBitRate sets BitRate field to given value.

### HasBitRate

`func (o *MediaStream) HasBitRate() bool`

HasBitRate returns a boolean if a field has been set.

### SetBitRateNil

`func (o *MediaStream) SetBitRateNil(b bool)`

 SetBitRateNil sets the value for BitRate to be an explicit nil

### UnsetBitRate
`func (o *MediaStream) UnsetBitRate()`

UnsetBitRate ensures that no value is present for BitRate, not even an explicit nil
### GetBitDepth

`func (o *MediaStream) GetBitDepth() int32`

GetBitDepth returns the BitDepth field if non-nil, zero value otherwise.

### GetBitDepthOk

`func (o *MediaStream) GetBitDepthOk() (*int32, bool)`

GetBitDepthOk returns a tuple with the BitDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitDepth

`func (o *MediaStream) SetBitDepth(v int32)`

SetBitDepth sets BitDepth field to given value.

### HasBitDepth

`func (o *MediaStream) HasBitDepth() bool`

HasBitDepth returns a boolean if a field has been set.

### SetBitDepthNil

`func (o *MediaStream) SetBitDepthNil(b bool)`

 SetBitDepthNil sets the value for BitDepth to be an explicit nil

### UnsetBitDepth
`func (o *MediaStream) UnsetBitDepth()`

UnsetBitDepth ensures that no value is present for BitDepth, not even an explicit nil
### GetRefFrames

`func (o *MediaStream) GetRefFrames() int32`

GetRefFrames returns the RefFrames field if non-nil, zero value otherwise.

### GetRefFramesOk

`func (o *MediaStream) GetRefFramesOk() (*int32, bool)`

GetRefFramesOk returns a tuple with the RefFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFrames

`func (o *MediaStream) SetRefFrames(v int32)`

SetRefFrames sets RefFrames field to given value.

### HasRefFrames

`func (o *MediaStream) HasRefFrames() bool`

HasRefFrames returns a boolean if a field has been set.

### SetRefFramesNil

`func (o *MediaStream) SetRefFramesNil(b bool)`

 SetRefFramesNil sets the value for RefFrames to be an explicit nil

### UnsetRefFrames
`func (o *MediaStream) UnsetRefFrames()`

UnsetRefFrames ensures that no value is present for RefFrames, not even an explicit nil
### GetPacketLength

`func (o *MediaStream) GetPacketLength() int32`

GetPacketLength returns the PacketLength field if non-nil, zero value otherwise.

### GetPacketLengthOk

`func (o *MediaStream) GetPacketLengthOk() (*int32, bool)`

GetPacketLengthOk returns a tuple with the PacketLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLength

`func (o *MediaStream) SetPacketLength(v int32)`

SetPacketLength sets PacketLength field to given value.

### HasPacketLength

`func (o *MediaStream) HasPacketLength() bool`

HasPacketLength returns a boolean if a field has been set.

### SetPacketLengthNil

`func (o *MediaStream) SetPacketLengthNil(b bool)`

 SetPacketLengthNil sets the value for PacketLength to be an explicit nil

### UnsetPacketLength
`func (o *MediaStream) UnsetPacketLength()`

UnsetPacketLength ensures that no value is present for PacketLength, not even an explicit nil
### GetChannels

`func (o *MediaStream) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MediaStream) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *MediaStream) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *MediaStream) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *MediaStream) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *MediaStream) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetSampleRate

`func (o *MediaStream) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *MediaStream) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *MediaStream) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *MediaStream) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRateNil

`func (o *MediaStream) SetSampleRateNil(b bool)`

 SetSampleRateNil sets the value for SampleRate to be an explicit nil

### UnsetSampleRate
`func (o *MediaStream) UnsetSampleRate()`

UnsetSampleRate ensures that no value is present for SampleRate, not even an explicit nil
### GetIsDefault

`func (o *MediaStream) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MediaStream) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MediaStream) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MediaStream) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsForced

`func (o *MediaStream) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *MediaStream) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *MediaStream) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.

### HasIsForced

`func (o *MediaStream) HasIsForced() bool`

HasIsForced returns a boolean if a field has been set.

### GetHeight

`func (o *MediaStream) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MediaStream) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MediaStream) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MediaStream) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *MediaStream) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *MediaStream) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *MediaStream) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MediaStream) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MediaStream) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MediaStream) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *MediaStream) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *MediaStream) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetAverageFrameRate

`func (o *MediaStream) GetAverageFrameRate() float32`

GetAverageFrameRate returns the AverageFrameRate field if non-nil, zero value otherwise.

### GetAverageFrameRateOk

`func (o *MediaStream) GetAverageFrameRateOk() (*float32, bool)`

GetAverageFrameRateOk returns a tuple with the AverageFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFrameRate

`func (o *MediaStream) SetAverageFrameRate(v float32)`

SetAverageFrameRate sets AverageFrameRate field to given value.

### HasAverageFrameRate

`func (o *MediaStream) HasAverageFrameRate() bool`

HasAverageFrameRate returns a boolean if a field has been set.

### SetAverageFrameRateNil

`func (o *MediaStream) SetAverageFrameRateNil(b bool)`

 SetAverageFrameRateNil sets the value for AverageFrameRate to be an explicit nil

### UnsetAverageFrameRate
`func (o *MediaStream) UnsetAverageFrameRate()`

UnsetAverageFrameRate ensures that no value is present for AverageFrameRate, not even an explicit nil
### GetRealFrameRate

`func (o *MediaStream) GetRealFrameRate() float32`

GetRealFrameRate returns the RealFrameRate field if non-nil, zero value otherwise.

### GetRealFrameRateOk

`func (o *MediaStream) GetRealFrameRateOk() (*float32, bool)`

GetRealFrameRateOk returns a tuple with the RealFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealFrameRate

`func (o *MediaStream) SetRealFrameRate(v float32)`

SetRealFrameRate sets RealFrameRate field to given value.

### HasRealFrameRate

`func (o *MediaStream) HasRealFrameRate() bool`

HasRealFrameRate returns a boolean if a field has been set.

### SetRealFrameRateNil

`func (o *MediaStream) SetRealFrameRateNil(b bool)`

 SetRealFrameRateNil sets the value for RealFrameRate to be an explicit nil

### UnsetRealFrameRate
`func (o *MediaStream) UnsetRealFrameRate()`

UnsetRealFrameRate ensures that no value is present for RealFrameRate, not even an explicit nil
### GetProfile

`func (o *MediaStream) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MediaStream) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MediaStream) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MediaStream) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetType

`func (o *MediaStream) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaStream) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaStream) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MediaStream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAspectRatio

`func (o *MediaStream) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *MediaStream) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *MediaStream) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *MediaStream) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetIndex

`func (o *MediaStream) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MediaStream) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MediaStream) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MediaStream) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetScore

`func (o *MediaStream) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MediaStream) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MediaStream) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *MediaStream) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *MediaStream) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *MediaStream) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetIsExternal

`func (o *MediaStream) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *MediaStream) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *MediaStream) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *MediaStream) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *MediaStream) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *MediaStream) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *MediaStream) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *MediaStream) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetDeliveryUrl

`func (o *MediaStream) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *MediaStream) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *MediaStream) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *MediaStream) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### GetIsExternalUrl

`func (o *MediaStream) GetIsExternalUrl() bool`

GetIsExternalUrl returns the IsExternalUrl field if non-nil, zero value otherwise.

### GetIsExternalUrlOk

`func (o *MediaStream) GetIsExternalUrlOk() (*bool, bool)`

GetIsExternalUrlOk returns a tuple with the IsExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalUrl

`func (o *MediaStream) SetIsExternalUrl(v bool)`

SetIsExternalUrl sets IsExternalUrl field to given value.

### HasIsExternalUrl

`func (o *MediaStream) HasIsExternalUrl() bool`

HasIsExternalUrl returns a boolean if a field has been set.

### SetIsExternalUrlNil

`func (o *MediaStream) SetIsExternalUrlNil(b bool)`

 SetIsExternalUrlNil sets the value for IsExternalUrl to be an explicit nil

### UnsetIsExternalUrl
`func (o *MediaStream) UnsetIsExternalUrl()`

UnsetIsExternalUrl ensures that no value is present for IsExternalUrl, not even an explicit nil
### GetIsTextSubtitleStream

`func (o *MediaStream) GetIsTextSubtitleStream() bool`

GetIsTextSubtitleStream returns the IsTextSubtitleStream field if non-nil, zero value otherwise.

### GetIsTextSubtitleStreamOk

`func (o *MediaStream) GetIsTextSubtitleStreamOk() (*bool, bool)`

GetIsTextSubtitleStreamOk returns a tuple with the IsTextSubtitleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTextSubtitleStream

`func (o *MediaStream) SetIsTextSubtitleStream(v bool)`

SetIsTextSubtitleStream sets IsTextSubtitleStream field to given value.

### HasIsTextSubtitleStream

`func (o *MediaStream) HasIsTextSubtitleStream() bool`

HasIsTextSubtitleStream returns a boolean if a field has been set.

### GetSupportsExternalStream

`func (o *MediaStream) GetSupportsExternalStream() bool`

GetSupportsExternalStream returns the SupportsExternalStream field if non-nil, zero value otherwise.

### GetSupportsExternalStreamOk

`func (o *MediaStream) GetSupportsExternalStreamOk() (*bool, bool)`

GetSupportsExternalStreamOk returns a tuple with the SupportsExternalStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExternalStream

`func (o *MediaStream) SetSupportsExternalStream(v bool)`

SetSupportsExternalStream sets SupportsExternalStream field to given value.

### HasSupportsExternalStream

`func (o *MediaStream) HasSupportsExternalStream() bool`

HasSupportsExternalStream returns a boolean if a field has been set.

### GetPath

`func (o *MediaStream) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MediaStream) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MediaStream) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MediaStream) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPixelFormat

`func (o *MediaStream) GetPixelFormat() string`

GetPixelFormat returns the PixelFormat field if non-nil, zero value otherwise.

### GetPixelFormatOk

`func (o *MediaStream) GetPixelFormatOk() (*string, bool)`

GetPixelFormatOk returns a tuple with the PixelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelFormat

`func (o *MediaStream) SetPixelFormat(v string)`

SetPixelFormat sets PixelFormat field to given value.

### HasPixelFormat

`func (o *MediaStream) HasPixelFormat() bool`

HasPixelFormat returns a boolean if a field has been set.

### GetLevel

`func (o *MediaStream) GetLevel() float64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MediaStream) GetLevelOk() (*float64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MediaStream) SetLevel(v float64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MediaStream) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *MediaStream) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *MediaStream) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetIsAnamorphic

`func (o *MediaStream) GetIsAnamorphic() bool`

GetIsAnamorphic returns the IsAnamorphic field if non-nil, zero value otherwise.

### GetIsAnamorphicOk

`func (o *MediaStream) GetIsAnamorphicOk() (*bool, bool)`

GetIsAnamorphicOk returns a tuple with the IsAnamorphic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnamorphic

`func (o *MediaStream) SetIsAnamorphic(v bool)`

SetIsAnamorphic sets IsAnamorphic field to given value.

### HasIsAnamorphic

`func (o *MediaStream) HasIsAnamorphic() bool`

HasIsAnamorphic returns a boolean if a field has been set.

### SetIsAnamorphicNil

`func (o *MediaStream) SetIsAnamorphicNil(b bool)`

 SetIsAnamorphicNil sets the value for IsAnamorphic to be an explicit nil

### UnsetIsAnamorphic
`func (o *MediaStream) UnsetIsAnamorphic()`

UnsetIsAnamorphic ensures that no value is present for IsAnamorphic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



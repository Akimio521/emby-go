# MediaEncodingCodecsVideoCodecsVideoCodecBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodecKind** | Pointer to **string** |  | [optional] 
**MediaTypeName** | Pointer to **string** |  | [optional] 
**VideoMediaType** | Pointer to **string** |  | [optional] 
**MinWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MinHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**WidthAlignment** | Pointer to **NullableInt32** |  | [optional] 
**HeightAlignment** | Pointer to **NullableInt32** |  | [optional] 
**MinFrameRate** | Pointer to **NullableInt32** |  | [optional] 
**MaxFrameRate** | Pointer to **NullableInt32** |  | [optional] 
**SupportedColorFormats** | Pointer to **[]string** |  | [optional] 
**SupportedColorFormatStrings** | Pointer to **[]string** |  | [optional] 
**ProfileAndLevelInformation** | Pointer to [**[]MediaEncodingCodecsCommonTypesProfileLevelInformation**](MediaEncodingCodecsCommonTypesProfileLevelInformation.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FrameworkCodec** | Pointer to **string** |  | [optional] 
**IsHardwareCodec** | Pointer to **bool** |  | [optional] 
**SecondaryFramework** | Pointer to **string** |  | [optional] 
**SecondaryFrameworkCodec** | Pointer to **string** |  | [optional] 
**MaxInstanceCount** | Pointer to **NullableInt32** |  | [optional] 
**MinBitRate** | Pointer to [**MediaEncodingCodecsCommonTypesBitRate**](MediaEncodingCodecsCommonTypesBitRate.md) |  | [optional] 
**MaxBitRate** | Pointer to [**MediaEncodingCodecsCommonTypesBitRate**](MediaEncodingCodecsCommonTypesBitRate.md) |  | [optional] 
**IsEnabledByDefault** | Pointer to **bool** |  | [optional] 
**DefaultPriority** | Pointer to **int32** |  | [optional] 

## Methods

### NewMediaEncodingCodecsVideoCodecsVideoCodecBase

`func NewMediaEncodingCodecsVideoCodecsVideoCodecBase() *MediaEncodingCodecsVideoCodecsVideoCodecBase`

NewMediaEncodingCodecsVideoCodecsVideoCodecBase instantiates a new MediaEncodingCodecsVideoCodecsVideoCodecBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaEncodingCodecsVideoCodecsVideoCodecBaseWithDefaults

`func NewMediaEncodingCodecsVideoCodecsVideoCodecBaseWithDefaults() *MediaEncodingCodecsVideoCodecsVideoCodecBase`

NewMediaEncodingCodecsVideoCodecsVideoCodecBaseWithDefaults instantiates a new MediaEncodingCodecsVideoCodecsVideoCodecBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodecKind

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetCodecKind() string`

GetCodecKind returns the CodecKind field if non-nil, zero value otherwise.

### GetCodecKindOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetCodecKindOk() (*string, bool)`

GetCodecKindOk returns a tuple with the CodecKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecKind

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetCodecKind(v string)`

SetCodecKind sets CodecKind field to given value.

### HasCodecKind

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasCodecKind() bool`

HasCodecKind returns a boolean if a field has been set.

### GetMediaTypeName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMediaTypeName() string`

GetMediaTypeName returns the MediaTypeName field if non-nil, zero value otherwise.

### GetMediaTypeNameOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMediaTypeNameOk() (*string, bool)`

GetMediaTypeNameOk returns a tuple with the MediaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMediaTypeName(v string)`

SetMediaTypeName sets MediaTypeName field to given value.

### HasMediaTypeName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMediaTypeName() bool`

HasMediaTypeName returns a boolean if a field has been set.

### GetVideoMediaType

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetVideoMediaType() string`

GetVideoMediaType returns the VideoMediaType field if non-nil, zero value otherwise.

### GetVideoMediaTypeOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetVideoMediaTypeOk() (*string, bool)`

GetVideoMediaTypeOk returns a tuple with the VideoMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMediaType

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetVideoMediaType(v string)`

SetVideoMediaType sets VideoMediaType field to given value.

### HasVideoMediaType

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasVideoMediaType() bool`

HasVideoMediaType returns a boolean if a field has been set.

### GetMinWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMaxWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMinHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetWidthAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetWidthAlignment() int32`

GetWidthAlignment returns the WidthAlignment field if non-nil, zero value otherwise.

### GetWidthAlignmentOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetWidthAlignmentOk() (*int32, bool)`

GetWidthAlignmentOk returns a tuple with the WidthAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetWidthAlignment(v int32)`

SetWidthAlignment sets WidthAlignment field to given value.

### HasWidthAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasWidthAlignment() bool`

HasWidthAlignment returns a boolean if a field has been set.

### SetWidthAlignmentNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetWidthAlignmentNil(b bool)`

 SetWidthAlignmentNil sets the value for WidthAlignment to be an explicit nil

### UnsetWidthAlignment
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetWidthAlignment()`

UnsetWidthAlignment ensures that no value is present for WidthAlignment, not even an explicit nil
### GetHeightAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetHeightAlignment() int32`

GetHeightAlignment returns the HeightAlignment field if non-nil, zero value otherwise.

### GetHeightAlignmentOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetHeightAlignmentOk() (*int32, bool)`

GetHeightAlignmentOk returns a tuple with the HeightAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetHeightAlignment(v int32)`

SetHeightAlignment sets HeightAlignment field to given value.

### HasHeightAlignment

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasHeightAlignment() bool`

HasHeightAlignment returns a boolean if a field has been set.

### SetHeightAlignmentNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetHeightAlignmentNil(b bool)`

 SetHeightAlignmentNil sets the value for HeightAlignment to be an explicit nil

### UnsetHeightAlignment
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetHeightAlignment()`

UnsetHeightAlignment ensures that no value is present for HeightAlignment, not even an explicit nil
### GetMinFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinFrameRate() int32`

GetMinFrameRate returns the MinFrameRate field if non-nil, zero value otherwise.

### GetMinFrameRateOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinFrameRateOk() (*int32, bool)`

GetMinFrameRateOk returns a tuple with the MinFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinFrameRate(v int32)`

SetMinFrameRate sets MinFrameRate field to given value.

### HasMinFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMinFrameRate() bool`

HasMinFrameRate returns a boolean if a field has been set.

### SetMinFrameRateNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinFrameRateNil(b bool)`

 SetMinFrameRateNil sets the value for MinFrameRate to be an explicit nil

### UnsetMinFrameRate
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMinFrameRate()`

UnsetMinFrameRate ensures that no value is present for MinFrameRate, not even an explicit nil
### GetMaxFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxFrameRate() int32`

GetMaxFrameRate returns the MaxFrameRate field if non-nil, zero value otherwise.

### GetMaxFrameRateOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxFrameRateOk() (*int32, bool)`

GetMaxFrameRateOk returns a tuple with the MaxFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxFrameRate(v int32)`

SetMaxFrameRate sets MaxFrameRate field to given value.

### HasMaxFrameRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMaxFrameRate() bool`

HasMaxFrameRate returns a boolean if a field has been set.

### SetMaxFrameRateNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxFrameRateNil(b bool)`

 SetMaxFrameRateNil sets the value for MaxFrameRate to be an explicit nil

### UnsetMaxFrameRate
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMaxFrameRate()`

UnsetMaxFrameRate ensures that no value is present for MaxFrameRate, not even an explicit nil
### GetSupportedColorFormats

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSupportedColorFormats() []string`

GetSupportedColorFormats returns the SupportedColorFormats field if non-nil, zero value otherwise.

### GetSupportedColorFormatsOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSupportedColorFormatsOk() (*[]string, bool)`

GetSupportedColorFormatsOk returns a tuple with the SupportedColorFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormats

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetSupportedColorFormats(v []string)`

SetSupportedColorFormats sets SupportedColorFormats field to given value.

### HasSupportedColorFormats

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasSupportedColorFormats() bool`

HasSupportedColorFormats returns a boolean if a field has been set.

### GetSupportedColorFormatStrings

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSupportedColorFormatStrings() []string`

GetSupportedColorFormatStrings returns the SupportedColorFormatStrings field if non-nil, zero value otherwise.

### GetSupportedColorFormatStringsOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSupportedColorFormatStringsOk() (*[]string, bool)`

GetSupportedColorFormatStringsOk returns a tuple with the SupportedColorFormatStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormatStrings

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetSupportedColorFormatStrings(v []string)`

SetSupportedColorFormatStrings sets SupportedColorFormatStrings field to given value.

### HasSupportedColorFormatStrings

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasSupportedColorFormatStrings() bool`

HasSupportedColorFormatStrings returns a boolean if a field has been set.

### GetProfileAndLevelInformation

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetProfileAndLevelInformation() []MediaEncodingCodecsCommonTypesProfileLevelInformation`

GetProfileAndLevelInformation returns the ProfileAndLevelInformation field if non-nil, zero value otherwise.

### GetProfileAndLevelInformationOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetProfileAndLevelInformationOk() (*[]MediaEncodingCodecsCommonTypesProfileLevelInformation, bool)`

GetProfileAndLevelInformationOk returns a tuple with the ProfileAndLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAndLevelInformation

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetProfileAndLevelInformation(v []MediaEncodingCodecsCommonTypesProfileLevelInformation)`

SetProfileAndLevelInformation sets ProfileAndLevelInformation field to given value.

### HasProfileAndLevelInformation

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasProfileAndLevelInformation() bool`

HasProfileAndLevelInformation returns a boolean if a field has been set.

### GetId

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetFrameworkCodec() string`

GetFrameworkCodec returns the FrameworkCodec field if non-nil, zero value otherwise.

### GetFrameworkCodecOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetFrameworkCodecOk() (*string, bool)`

GetFrameworkCodecOk returns a tuple with the FrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetFrameworkCodec(v string)`

SetFrameworkCodec sets FrameworkCodec field to given value.

### HasFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasFrameworkCodec() bool`

HasFrameworkCodec returns a boolean if a field has been set.

### GetIsHardwareCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetIsHardwareCodec() bool`

GetIsHardwareCodec returns the IsHardwareCodec field if non-nil, zero value otherwise.

### GetIsHardwareCodecOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetIsHardwareCodecOk() (*bool, bool)`

GetIsHardwareCodecOk returns a tuple with the IsHardwareCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetIsHardwareCodec(v bool)`

SetIsHardwareCodec sets IsHardwareCodec field to given value.

### HasIsHardwareCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasIsHardwareCodec() bool`

HasIsHardwareCodec returns a boolean if a field has been set.

### GetSecondaryFramework

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSecondaryFramework() string`

GetSecondaryFramework returns the SecondaryFramework field if non-nil, zero value otherwise.

### GetSecondaryFrameworkOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSecondaryFrameworkOk() (*string, bool)`

GetSecondaryFrameworkOk returns a tuple with the SecondaryFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFramework

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetSecondaryFramework(v string)`

SetSecondaryFramework sets SecondaryFramework field to given value.

### HasSecondaryFramework

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasSecondaryFramework() bool`

HasSecondaryFramework returns a boolean if a field has been set.

### GetSecondaryFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSecondaryFrameworkCodec() string`

GetSecondaryFrameworkCodec returns the SecondaryFrameworkCodec field if non-nil, zero value otherwise.

### GetSecondaryFrameworkCodecOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetSecondaryFrameworkCodecOk() (*string, bool)`

GetSecondaryFrameworkCodecOk returns a tuple with the SecondaryFrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetSecondaryFrameworkCodec(v string)`

SetSecondaryFrameworkCodec sets SecondaryFrameworkCodec field to given value.

### HasSecondaryFrameworkCodec

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasSecondaryFrameworkCodec() bool`

HasSecondaryFrameworkCodec returns a boolean if a field has been set.

### GetMaxInstanceCount

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxInstanceCount() int32`

GetMaxInstanceCount returns the MaxInstanceCount field if non-nil, zero value otherwise.

### GetMaxInstanceCountOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxInstanceCountOk() (*int32, bool)`

GetMaxInstanceCountOk returns a tuple with the MaxInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceCount

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxInstanceCount(v int32)`

SetMaxInstanceCount sets MaxInstanceCount field to given value.

### HasMaxInstanceCount

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMaxInstanceCount() bool`

HasMaxInstanceCount returns a boolean if a field has been set.

### SetMaxInstanceCountNil

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxInstanceCountNil(b bool)`

 SetMaxInstanceCountNil sets the value for MaxInstanceCount to be an explicit nil

### UnsetMaxInstanceCount
`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) UnsetMaxInstanceCount()`

UnsetMaxInstanceCount ensures that no value is present for MaxInstanceCount, not even an explicit nil
### GetMinBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinBitRate() MediaEncodingCodecsCommonTypesBitRate`

GetMinBitRate returns the MinBitRate field if non-nil, zero value otherwise.

### GetMinBitRateOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMinBitRateOk() (*MediaEncodingCodecsCommonTypesBitRate, bool)`

GetMinBitRateOk returns a tuple with the MinBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMinBitRate(v MediaEncodingCodecsCommonTypesBitRate)`

SetMinBitRate sets MinBitRate field to given value.

### HasMinBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMinBitRate() bool`

HasMinBitRate returns a boolean if a field has been set.

### GetMaxBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxBitRate() MediaEncodingCodecsCommonTypesBitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetMaxBitRateOk() (*MediaEncodingCodecsCommonTypesBitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetMaxBitRate(v MediaEncodingCodecsCommonTypesBitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetIsEnabledByDefault

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetIsEnabledByDefault() bool`

GetIsEnabledByDefault returns the IsEnabledByDefault field if non-nil, zero value otherwise.

### GetIsEnabledByDefaultOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetIsEnabledByDefaultOk() (*bool, bool)`

GetIsEnabledByDefaultOk returns a tuple with the IsEnabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledByDefault

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetIsEnabledByDefault(v bool)`

SetIsEnabledByDefault sets IsEnabledByDefault field to given value.

### HasIsEnabledByDefault

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasIsEnabledByDefault() bool`

HasIsEnabledByDefault returns a boolean if a field has been set.

### GetDefaultPriority

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDefaultPriority() int32`

GetDefaultPriority returns the DefaultPriority field if non-nil, zero value otherwise.

### GetDefaultPriorityOk

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) GetDefaultPriorityOk() (*int32, bool)`

GetDefaultPriorityOk returns a tuple with the DefaultPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriority

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) SetDefaultPriority(v int32)`

SetDefaultPriority sets DefaultPriority field to given value.

### HasDefaultPriority

`func (o *MediaEncodingCodecsVideoCodecsVideoCodecBase) HasDefaultPriority() bool`

HasDefaultPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



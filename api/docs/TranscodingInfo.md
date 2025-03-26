# TranscodingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**IsVideoDirect** | Pointer to **bool** |  | [optional] 
**IsAudioDirect** | Pointer to **bool** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Framerate** | Pointer to **NullableFloat32** |  | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**TranscodingPositionTicks** | Pointer to **NullableFloat64** |  | [optional] 
**TranscodingStartPositionTicks** | Pointer to **NullableFloat64** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**AudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**TranscodeReasons** | Pointer to **[]string** |  | [optional] 
**CurrentCpuUsage** | Pointer to **NullableFloat64** |  | [optional] 
**AverageCpuUsage** | Pointer to **NullableFloat64** |  | [optional] 
**CpuHistory** | Pointer to [**[]TupleDoubleDouble**](TupleDoubleDouble.md) |  | [optional] 
**CurrentThrottle** | Pointer to **NullableInt32** |  | [optional] 
**VideoDecoder** | Pointer to **string** |  | [optional] 
**VideoDecoderIsHardware** | Pointer to **bool** |  | [optional] 
**VideoDecoderMediaType** | Pointer to **string** |  | [optional] 
**VideoDecoderHwAccel** | Pointer to **string** |  | [optional] 
**VideoEncoder** | Pointer to **string** |  | [optional] 
**VideoEncoderIsHardware** | Pointer to **bool** |  | [optional] 
**VideoEncoderMediaType** | Pointer to **string** |  | [optional] 
**VideoEncoderHwAccel** | Pointer to **string** |  | [optional] 

## Methods

### NewTranscodingInfo

`func NewTranscodingInfo() *TranscodingInfo`

NewTranscodingInfo instantiates a new TranscodingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscodingInfoWithDefaults

`func NewTranscodingInfoWithDefaults() *TranscodingInfo`

NewTranscodingInfoWithDefaults instantiates a new TranscodingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioCodec

`func (o *TranscodingInfo) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *TranscodingInfo) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *TranscodingInfo) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *TranscodingInfo) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *TranscodingInfo) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *TranscodingInfo) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *TranscodingInfo) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *TranscodingInfo) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetContainer

`func (o *TranscodingInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *TranscodingInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *TranscodingInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *TranscodingInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetIsVideoDirect

`func (o *TranscodingInfo) GetIsVideoDirect() bool`

GetIsVideoDirect returns the IsVideoDirect field if non-nil, zero value otherwise.

### GetIsVideoDirectOk

`func (o *TranscodingInfo) GetIsVideoDirectOk() (*bool, bool)`

GetIsVideoDirectOk returns a tuple with the IsVideoDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideoDirect

`func (o *TranscodingInfo) SetIsVideoDirect(v bool)`

SetIsVideoDirect sets IsVideoDirect field to given value.

### HasIsVideoDirect

`func (o *TranscodingInfo) HasIsVideoDirect() bool`

HasIsVideoDirect returns a boolean if a field has been set.

### GetIsAudioDirect

`func (o *TranscodingInfo) GetIsAudioDirect() bool`

GetIsAudioDirect returns the IsAudioDirect field if non-nil, zero value otherwise.

### GetIsAudioDirectOk

`func (o *TranscodingInfo) GetIsAudioDirectOk() (*bool, bool)`

GetIsAudioDirectOk returns a tuple with the IsAudioDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAudioDirect

`func (o *TranscodingInfo) SetIsAudioDirect(v bool)`

SetIsAudioDirect sets IsAudioDirect field to given value.

### HasIsAudioDirect

`func (o *TranscodingInfo) HasIsAudioDirect() bool`

HasIsAudioDirect returns a boolean if a field has been set.

### GetBitrate

`func (o *TranscodingInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *TranscodingInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *TranscodingInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *TranscodingInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *TranscodingInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *TranscodingInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetFramerate

`func (o *TranscodingInfo) GetFramerate() float32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *TranscodingInfo) GetFramerateOk() (*float32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *TranscodingInfo) SetFramerate(v float32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *TranscodingInfo) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### SetFramerateNil

`func (o *TranscodingInfo) SetFramerateNil(b bool)`

 SetFramerateNil sets the value for Framerate to be an explicit nil

### UnsetFramerate
`func (o *TranscodingInfo) UnsetFramerate()`

UnsetFramerate ensures that no value is present for Framerate, not even an explicit nil
### GetCompletionPercentage

`func (o *TranscodingInfo) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *TranscodingInfo) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *TranscodingInfo) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *TranscodingInfo) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *TranscodingInfo) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *TranscodingInfo) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetTranscodingPositionTicks

`func (o *TranscodingInfo) GetTranscodingPositionTicks() float64`

GetTranscodingPositionTicks returns the TranscodingPositionTicks field if non-nil, zero value otherwise.

### GetTranscodingPositionTicksOk

`func (o *TranscodingInfo) GetTranscodingPositionTicksOk() (*float64, bool)`

GetTranscodingPositionTicksOk returns a tuple with the TranscodingPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingPositionTicks

`func (o *TranscodingInfo) SetTranscodingPositionTicks(v float64)`

SetTranscodingPositionTicks sets TranscodingPositionTicks field to given value.

### HasTranscodingPositionTicks

`func (o *TranscodingInfo) HasTranscodingPositionTicks() bool`

HasTranscodingPositionTicks returns a boolean if a field has been set.

### SetTranscodingPositionTicksNil

`func (o *TranscodingInfo) SetTranscodingPositionTicksNil(b bool)`

 SetTranscodingPositionTicksNil sets the value for TranscodingPositionTicks to be an explicit nil

### UnsetTranscodingPositionTicks
`func (o *TranscodingInfo) UnsetTranscodingPositionTicks()`

UnsetTranscodingPositionTicks ensures that no value is present for TranscodingPositionTicks, not even an explicit nil
### GetTranscodingStartPositionTicks

`func (o *TranscodingInfo) GetTranscodingStartPositionTicks() float64`

GetTranscodingStartPositionTicks returns the TranscodingStartPositionTicks field if non-nil, zero value otherwise.

### GetTranscodingStartPositionTicksOk

`func (o *TranscodingInfo) GetTranscodingStartPositionTicksOk() (*float64, bool)`

GetTranscodingStartPositionTicksOk returns a tuple with the TranscodingStartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingStartPositionTicks

`func (o *TranscodingInfo) SetTranscodingStartPositionTicks(v float64)`

SetTranscodingStartPositionTicks sets TranscodingStartPositionTicks field to given value.

### HasTranscodingStartPositionTicks

`func (o *TranscodingInfo) HasTranscodingStartPositionTicks() bool`

HasTranscodingStartPositionTicks returns a boolean if a field has been set.

### SetTranscodingStartPositionTicksNil

`func (o *TranscodingInfo) SetTranscodingStartPositionTicksNil(b bool)`

 SetTranscodingStartPositionTicksNil sets the value for TranscodingStartPositionTicks to be an explicit nil

### UnsetTranscodingStartPositionTicks
`func (o *TranscodingInfo) UnsetTranscodingStartPositionTicks()`

UnsetTranscodingStartPositionTicks ensures that no value is present for TranscodingStartPositionTicks, not even an explicit nil
### GetWidth

`func (o *TranscodingInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TranscodingInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TranscodingInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TranscodingInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *TranscodingInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *TranscodingInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *TranscodingInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TranscodingInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TranscodingInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TranscodingInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *TranscodingInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *TranscodingInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetAudioChannels

`func (o *TranscodingInfo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *TranscodingInfo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *TranscodingInfo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *TranscodingInfo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannelsNil

`func (o *TranscodingInfo) SetAudioChannelsNil(b bool)`

 SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil

### UnsetAudioChannels
`func (o *TranscodingInfo) UnsetAudioChannels()`

UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
### GetTranscodeReasons

`func (o *TranscodingInfo) GetTranscodeReasons() []string`

GetTranscodeReasons returns the TranscodeReasons field if non-nil, zero value otherwise.

### GetTranscodeReasonsOk

`func (o *TranscodingInfo) GetTranscodeReasonsOk() (*[]string, bool)`

GetTranscodeReasonsOk returns a tuple with the TranscodeReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeReasons

`func (o *TranscodingInfo) SetTranscodeReasons(v []string)`

SetTranscodeReasons sets TranscodeReasons field to given value.

### HasTranscodeReasons

`func (o *TranscodingInfo) HasTranscodeReasons() bool`

HasTranscodeReasons returns a boolean if a field has been set.

### GetCurrentCpuUsage

`func (o *TranscodingInfo) GetCurrentCpuUsage() float64`

GetCurrentCpuUsage returns the CurrentCpuUsage field if non-nil, zero value otherwise.

### GetCurrentCpuUsageOk

`func (o *TranscodingInfo) GetCurrentCpuUsageOk() (*float64, bool)`

GetCurrentCpuUsageOk returns a tuple with the CurrentCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuUsage

`func (o *TranscodingInfo) SetCurrentCpuUsage(v float64)`

SetCurrentCpuUsage sets CurrentCpuUsage field to given value.

### HasCurrentCpuUsage

`func (o *TranscodingInfo) HasCurrentCpuUsage() bool`

HasCurrentCpuUsage returns a boolean if a field has been set.

### SetCurrentCpuUsageNil

`func (o *TranscodingInfo) SetCurrentCpuUsageNil(b bool)`

 SetCurrentCpuUsageNil sets the value for CurrentCpuUsage to be an explicit nil

### UnsetCurrentCpuUsage
`func (o *TranscodingInfo) UnsetCurrentCpuUsage()`

UnsetCurrentCpuUsage ensures that no value is present for CurrentCpuUsage, not even an explicit nil
### GetAverageCpuUsage

`func (o *TranscodingInfo) GetAverageCpuUsage() float64`

GetAverageCpuUsage returns the AverageCpuUsage field if non-nil, zero value otherwise.

### GetAverageCpuUsageOk

`func (o *TranscodingInfo) GetAverageCpuUsageOk() (*float64, bool)`

GetAverageCpuUsageOk returns a tuple with the AverageCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCpuUsage

`func (o *TranscodingInfo) SetAverageCpuUsage(v float64)`

SetAverageCpuUsage sets AverageCpuUsage field to given value.

### HasAverageCpuUsage

`func (o *TranscodingInfo) HasAverageCpuUsage() bool`

HasAverageCpuUsage returns a boolean if a field has been set.

### SetAverageCpuUsageNil

`func (o *TranscodingInfo) SetAverageCpuUsageNil(b bool)`

 SetAverageCpuUsageNil sets the value for AverageCpuUsage to be an explicit nil

### UnsetAverageCpuUsage
`func (o *TranscodingInfo) UnsetAverageCpuUsage()`

UnsetAverageCpuUsage ensures that no value is present for AverageCpuUsage, not even an explicit nil
### GetCpuHistory

`func (o *TranscodingInfo) GetCpuHistory() []TupleDoubleDouble`

GetCpuHistory returns the CpuHistory field if non-nil, zero value otherwise.

### GetCpuHistoryOk

`func (o *TranscodingInfo) GetCpuHistoryOk() (*[]TupleDoubleDouble, bool)`

GetCpuHistoryOk returns a tuple with the CpuHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHistory

`func (o *TranscodingInfo) SetCpuHistory(v []TupleDoubleDouble)`

SetCpuHistory sets CpuHistory field to given value.

### HasCpuHistory

`func (o *TranscodingInfo) HasCpuHistory() bool`

HasCpuHistory returns a boolean if a field has been set.

### GetCurrentThrottle

`func (o *TranscodingInfo) GetCurrentThrottle() int32`

GetCurrentThrottle returns the CurrentThrottle field if non-nil, zero value otherwise.

### GetCurrentThrottleOk

`func (o *TranscodingInfo) GetCurrentThrottleOk() (*int32, bool)`

GetCurrentThrottleOk returns a tuple with the CurrentThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentThrottle

`func (o *TranscodingInfo) SetCurrentThrottle(v int32)`

SetCurrentThrottle sets CurrentThrottle field to given value.

### HasCurrentThrottle

`func (o *TranscodingInfo) HasCurrentThrottle() bool`

HasCurrentThrottle returns a boolean if a field has been set.

### SetCurrentThrottleNil

`func (o *TranscodingInfo) SetCurrentThrottleNil(b bool)`

 SetCurrentThrottleNil sets the value for CurrentThrottle to be an explicit nil

### UnsetCurrentThrottle
`func (o *TranscodingInfo) UnsetCurrentThrottle()`

UnsetCurrentThrottle ensures that no value is present for CurrentThrottle, not even an explicit nil
### GetVideoDecoder

`func (o *TranscodingInfo) GetVideoDecoder() string`

GetVideoDecoder returns the VideoDecoder field if non-nil, zero value otherwise.

### GetVideoDecoderOk

`func (o *TranscodingInfo) GetVideoDecoderOk() (*string, bool)`

GetVideoDecoderOk returns a tuple with the VideoDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoder

`func (o *TranscodingInfo) SetVideoDecoder(v string)`

SetVideoDecoder sets VideoDecoder field to given value.

### HasVideoDecoder

`func (o *TranscodingInfo) HasVideoDecoder() bool`

HasVideoDecoder returns a boolean if a field has been set.

### GetVideoDecoderIsHardware

`func (o *TranscodingInfo) GetVideoDecoderIsHardware() bool`

GetVideoDecoderIsHardware returns the VideoDecoderIsHardware field if non-nil, zero value otherwise.

### GetVideoDecoderIsHardwareOk

`func (o *TranscodingInfo) GetVideoDecoderIsHardwareOk() (*bool, bool)`

GetVideoDecoderIsHardwareOk returns a tuple with the VideoDecoderIsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderIsHardware

`func (o *TranscodingInfo) SetVideoDecoderIsHardware(v bool)`

SetVideoDecoderIsHardware sets VideoDecoderIsHardware field to given value.

### HasVideoDecoderIsHardware

`func (o *TranscodingInfo) HasVideoDecoderIsHardware() bool`

HasVideoDecoderIsHardware returns a boolean if a field has been set.

### GetVideoDecoderMediaType

`func (o *TranscodingInfo) GetVideoDecoderMediaType() string`

GetVideoDecoderMediaType returns the VideoDecoderMediaType field if non-nil, zero value otherwise.

### GetVideoDecoderMediaTypeOk

`func (o *TranscodingInfo) GetVideoDecoderMediaTypeOk() (*string, bool)`

GetVideoDecoderMediaTypeOk returns a tuple with the VideoDecoderMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderMediaType

`func (o *TranscodingInfo) SetVideoDecoderMediaType(v string)`

SetVideoDecoderMediaType sets VideoDecoderMediaType field to given value.

### HasVideoDecoderMediaType

`func (o *TranscodingInfo) HasVideoDecoderMediaType() bool`

HasVideoDecoderMediaType returns a boolean if a field has been set.

### GetVideoDecoderHwAccel

`func (o *TranscodingInfo) GetVideoDecoderHwAccel() string`

GetVideoDecoderHwAccel returns the VideoDecoderHwAccel field if non-nil, zero value otherwise.

### GetVideoDecoderHwAccelOk

`func (o *TranscodingInfo) GetVideoDecoderHwAccelOk() (*string, bool)`

GetVideoDecoderHwAccelOk returns a tuple with the VideoDecoderHwAccel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderHwAccel

`func (o *TranscodingInfo) SetVideoDecoderHwAccel(v string)`

SetVideoDecoderHwAccel sets VideoDecoderHwAccel field to given value.

### HasVideoDecoderHwAccel

`func (o *TranscodingInfo) HasVideoDecoderHwAccel() bool`

HasVideoDecoderHwAccel returns a boolean if a field has been set.

### GetVideoEncoder

`func (o *TranscodingInfo) GetVideoEncoder() string`

GetVideoEncoder returns the VideoEncoder field if non-nil, zero value otherwise.

### GetVideoEncoderOk

`func (o *TranscodingInfo) GetVideoEncoderOk() (*string, bool)`

GetVideoEncoderOk returns a tuple with the VideoEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoder

`func (o *TranscodingInfo) SetVideoEncoder(v string)`

SetVideoEncoder sets VideoEncoder field to given value.

### HasVideoEncoder

`func (o *TranscodingInfo) HasVideoEncoder() bool`

HasVideoEncoder returns a boolean if a field has been set.

### GetVideoEncoderIsHardware

`func (o *TranscodingInfo) GetVideoEncoderIsHardware() bool`

GetVideoEncoderIsHardware returns the VideoEncoderIsHardware field if non-nil, zero value otherwise.

### GetVideoEncoderIsHardwareOk

`func (o *TranscodingInfo) GetVideoEncoderIsHardwareOk() (*bool, bool)`

GetVideoEncoderIsHardwareOk returns a tuple with the VideoEncoderIsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderIsHardware

`func (o *TranscodingInfo) SetVideoEncoderIsHardware(v bool)`

SetVideoEncoderIsHardware sets VideoEncoderIsHardware field to given value.

### HasVideoEncoderIsHardware

`func (o *TranscodingInfo) HasVideoEncoderIsHardware() bool`

HasVideoEncoderIsHardware returns a boolean if a field has been set.

### GetVideoEncoderMediaType

`func (o *TranscodingInfo) GetVideoEncoderMediaType() string`

GetVideoEncoderMediaType returns the VideoEncoderMediaType field if non-nil, zero value otherwise.

### GetVideoEncoderMediaTypeOk

`func (o *TranscodingInfo) GetVideoEncoderMediaTypeOk() (*string, bool)`

GetVideoEncoderMediaTypeOk returns a tuple with the VideoEncoderMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderMediaType

`func (o *TranscodingInfo) SetVideoEncoderMediaType(v string)`

SetVideoEncoderMediaType sets VideoEncoderMediaType field to given value.

### HasVideoEncoderMediaType

`func (o *TranscodingInfo) HasVideoEncoderMediaType() bool`

HasVideoEncoderMediaType returns a boolean if a field has been set.

### GetVideoEncoderHwAccel

`func (o *TranscodingInfo) GetVideoEncoderHwAccel() string`

GetVideoEncoderHwAccel returns the VideoEncoderHwAccel field if non-nil, zero value otherwise.

### GetVideoEncoderHwAccelOk

`func (o *TranscodingInfo) GetVideoEncoderHwAccelOk() (*string, bool)`

GetVideoEncoderHwAccelOk returns a tuple with the VideoEncoderHwAccel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderHwAccel

`func (o *TranscodingInfo) SetVideoEncoderHwAccel(v string)`

SetVideoEncoderHwAccel sets VideoEncoderHwAccel field to given value.

### HasVideoEncoderHwAccel

`func (o *TranscodingInfo) HasVideoEncoderHwAccel() bool`

HasVideoEncoderHwAccel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



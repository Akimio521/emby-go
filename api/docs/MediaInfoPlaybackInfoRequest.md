# MediaInfoPlaybackInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**DeviceProfile** | Pointer to [**DlnaDeviceProfile**](DlnaDeviceProfile.md) |  | [optional] 
**EnableDirectPlay** | Pointer to **bool** |  | [optional] 
**EnableDirectStream** | Pointer to **bool** |  | [optional] 
**EnableTranscoding** | Pointer to **bool** |  | [optional] 
**AllowVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**AllowAudioStreamCopy** | Pointer to **bool** |  | [optional] 
**IsPlayback** | Pointer to **bool** |  | [optional] 
**AutoOpenLiveStream** | Pointer to **bool** |  | [optional] 
**DirectPlayProtocols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMediaInfoPlaybackInfoRequest

`func NewMediaInfoPlaybackInfoRequest() *MediaInfoPlaybackInfoRequest`

NewMediaInfoPlaybackInfoRequest instantiates a new MediaInfoPlaybackInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaInfoPlaybackInfoRequestWithDefaults

`func NewMediaInfoPlaybackInfoRequestWithDefaults() *MediaInfoPlaybackInfoRequest`

NewMediaInfoPlaybackInfoRequestWithDefaults instantiates a new MediaInfoPlaybackInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaInfoPlaybackInfoRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaInfoPlaybackInfoRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaInfoPlaybackInfoRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaInfoPlaybackInfoRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *MediaInfoPlaybackInfoRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MediaInfoPlaybackInfoRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MediaInfoPlaybackInfoRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MediaInfoPlaybackInfoRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *MediaInfoPlaybackInfoRequest) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *MediaInfoPlaybackInfoRequest) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *MediaInfoPlaybackInfoRequest) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *MediaInfoPlaybackInfoRequest) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *MediaInfoPlaybackInfoRequest) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *MediaInfoPlaybackInfoRequest) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *MediaInfoPlaybackInfoRequest) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *MediaInfoPlaybackInfoRequest) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *MediaInfoPlaybackInfoRequest) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *MediaInfoPlaybackInfoRequest) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *MediaInfoPlaybackInfoRequest) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *MediaInfoPlaybackInfoRequest) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *MediaInfoPlaybackInfoRequest) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *MediaInfoPlaybackInfoRequest) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *MediaInfoPlaybackInfoRequest) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *MediaInfoPlaybackInfoRequest) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *MediaInfoPlaybackInfoRequest) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *MediaInfoPlaybackInfoRequest) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *MediaInfoPlaybackInfoRequest) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *MediaInfoPlaybackInfoRequest) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *MediaInfoPlaybackInfoRequest) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *MediaInfoPlaybackInfoRequest) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *MediaInfoPlaybackInfoRequest) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *MediaInfoPlaybackInfoRequest) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *MediaInfoPlaybackInfoRequest) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetMediaSourceId

`func (o *MediaInfoPlaybackInfoRequest) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *MediaInfoPlaybackInfoRequest) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *MediaInfoPlaybackInfoRequest) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *MediaInfoPlaybackInfoRequest) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *MediaInfoPlaybackInfoRequest) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *MediaInfoPlaybackInfoRequest) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *MediaInfoPlaybackInfoRequest) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *MediaInfoPlaybackInfoRequest) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *MediaInfoPlaybackInfoRequest) GetDeviceProfile() DlnaDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *MediaInfoPlaybackInfoRequest) GetDeviceProfileOk() (*DlnaDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *MediaInfoPlaybackInfoRequest) SetDeviceProfile(v DlnaDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *MediaInfoPlaybackInfoRequest) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetEnableDirectPlay

`func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *MediaInfoPlaybackInfoRequest) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *MediaInfoPlaybackInfoRequest) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### GetEnableDirectStream

`func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *MediaInfoPlaybackInfoRequest) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *MediaInfoPlaybackInfoRequest) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### GetEnableTranscoding

`func (o *MediaInfoPlaybackInfoRequest) GetEnableTranscoding() bool`

GetEnableTranscoding returns the EnableTranscoding field if non-nil, zero value otherwise.

### GetEnableTranscodingOk

`func (o *MediaInfoPlaybackInfoRequest) GetEnableTranscodingOk() (*bool, bool)`

GetEnableTranscodingOk returns a tuple with the EnableTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTranscoding

`func (o *MediaInfoPlaybackInfoRequest) SetEnableTranscoding(v bool)`

SetEnableTranscoding sets EnableTranscoding field to given value.

### HasEnableTranscoding

`func (o *MediaInfoPlaybackInfoRequest) HasEnableTranscoding() bool`

HasEnableTranscoding returns a boolean if a field has been set.

### GetAllowVideoStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) GetAllowVideoStreamCopy() bool`

GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowVideoStreamCopyOk

`func (o *MediaInfoPlaybackInfoRequest) GetAllowVideoStreamCopyOk() (*bool, bool)`

GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideoStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) SetAllowVideoStreamCopy(v bool)`

SetAllowVideoStreamCopy sets AllowVideoStreamCopy field to given value.

### HasAllowVideoStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) HasAllowVideoStreamCopy() bool`

HasAllowVideoStreamCopy returns a boolean if a field has been set.

### GetAllowAudioStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) GetAllowAudioStreamCopy() bool`

GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field if non-nil, zero value otherwise.

### GetAllowAudioStreamCopyOk

`func (o *MediaInfoPlaybackInfoRequest) GetAllowAudioStreamCopyOk() (*bool, bool)`

GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAudioStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) SetAllowAudioStreamCopy(v bool)`

SetAllowAudioStreamCopy sets AllowAudioStreamCopy field to given value.

### HasAllowAudioStreamCopy

`func (o *MediaInfoPlaybackInfoRequest) HasAllowAudioStreamCopy() bool`

HasAllowAudioStreamCopy returns a boolean if a field has been set.

### GetIsPlayback

`func (o *MediaInfoPlaybackInfoRequest) GetIsPlayback() bool`

GetIsPlayback returns the IsPlayback field if non-nil, zero value otherwise.

### GetIsPlaybackOk

`func (o *MediaInfoPlaybackInfoRequest) GetIsPlaybackOk() (*bool, bool)`

GetIsPlaybackOk returns a tuple with the IsPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayback

`func (o *MediaInfoPlaybackInfoRequest) SetIsPlayback(v bool)`

SetIsPlayback sets IsPlayback field to given value.

### HasIsPlayback

`func (o *MediaInfoPlaybackInfoRequest) HasIsPlayback() bool`

HasIsPlayback returns a boolean if a field has been set.

### GetAutoOpenLiveStream

`func (o *MediaInfoPlaybackInfoRequest) GetAutoOpenLiveStream() bool`

GetAutoOpenLiveStream returns the AutoOpenLiveStream field if non-nil, zero value otherwise.

### GetAutoOpenLiveStreamOk

`func (o *MediaInfoPlaybackInfoRequest) GetAutoOpenLiveStreamOk() (*bool, bool)`

GetAutoOpenLiveStreamOk returns a tuple with the AutoOpenLiveStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOpenLiveStream

`func (o *MediaInfoPlaybackInfoRequest) SetAutoOpenLiveStream(v bool)`

SetAutoOpenLiveStream sets AutoOpenLiveStream field to given value.

### HasAutoOpenLiveStream

`func (o *MediaInfoPlaybackInfoRequest) HasAutoOpenLiveStream() bool`

HasAutoOpenLiveStream returns a boolean if a field has been set.

### GetDirectPlayProtocols

`func (o *MediaInfoPlaybackInfoRequest) GetDirectPlayProtocols() []string`

GetDirectPlayProtocols returns the DirectPlayProtocols field if non-nil, zero value otherwise.

### GetDirectPlayProtocolsOk

`func (o *MediaInfoPlaybackInfoRequest) GetDirectPlayProtocolsOk() (*[]string, bool)`

GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProtocols

`func (o *MediaInfoPlaybackInfoRequest) SetDirectPlayProtocols(v []string)`

SetDirectPlayProtocols sets DirectPlayProtocols field to given value.

### HasDirectPlayProtocols

`func (o *MediaInfoPlaybackInfoRequest) HasDirectPlayProtocols() bool`

HasDirectPlayProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



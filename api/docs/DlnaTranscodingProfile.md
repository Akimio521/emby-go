# DlnaTranscodingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**EstimateContentLength** | Pointer to **bool** |  | [optional] 
**EnableMpegtsM2TsMode** | Pointer to **bool** |  | [optional] 
**TranscodeSeekInfo** | Pointer to **string** |  | [optional] 
**CopyTimestamps** | Pointer to **bool** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**MaxAudioChannels** | Pointer to **string** |  | [optional] 
**MinSegments** | Pointer to **int32** |  | [optional] 
**SegmentLength** | Pointer to **int32** |  | [optional] 
**BreakOnNonKeyFrames** | Pointer to **bool** |  | [optional] 
**ManifestSubtitles** | Pointer to **string** |  | [optional] 

## Methods

### NewDlnaTranscodingProfile

`func NewDlnaTranscodingProfile() *DlnaTranscodingProfile`

NewDlnaTranscodingProfile instantiates a new DlnaTranscodingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaTranscodingProfileWithDefaults

`func NewDlnaTranscodingProfileWithDefaults() *DlnaTranscodingProfile`

NewDlnaTranscodingProfileWithDefaults instantiates a new DlnaTranscodingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *DlnaTranscodingProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DlnaTranscodingProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DlnaTranscodingProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DlnaTranscodingProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetType

`func (o *DlnaTranscodingProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DlnaTranscodingProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DlnaTranscodingProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DlnaTranscodingProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVideoCodec

`func (o *DlnaTranscodingProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *DlnaTranscodingProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *DlnaTranscodingProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *DlnaTranscodingProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *DlnaTranscodingProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *DlnaTranscodingProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *DlnaTranscodingProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *DlnaTranscodingProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetProtocol

`func (o *DlnaTranscodingProfile) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DlnaTranscodingProfile) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DlnaTranscodingProfile) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DlnaTranscodingProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEstimateContentLength

`func (o *DlnaTranscodingProfile) GetEstimateContentLength() bool`

GetEstimateContentLength returns the EstimateContentLength field if non-nil, zero value otherwise.

### GetEstimateContentLengthOk

`func (o *DlnaTranscodingProfile) GetEstimateContentLengthOk() (*bool, bool)`

GetEstimateContentLengthOk returns a tuple with the EstimateContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateContentLength

`func (o *DlnaTranscodingProfile) SetEstimateContentLength(v bool)`

SetEstimateContentLength sets EstimateContentLength field to given value.

### HasEstimateContentLength

`func (o *DlnaTranscodingProfile) HasEstimateContentLength() bool`

HasEstimateContentLength returns a boolean if a field has been set.

### GetEnableMpegtsM2TsMode

`func (o *DlnaTranscodingProfile) GetEnableMpegtsM2TsMode() bool`

GetEnableMpegtsM2TsMode returns the EnableMpegtsM2TsMode field if non-nil, zero value otherwise.

### GetEnableMpegtsM2TsModeOk

`func (o *DlnaTranscodingProfile) GetEnableMpegtsM2TsModeOk() (*bool, bool)`

GetEnableMpegtsM2TsModeOk returns a tuple with the EnableMpegtsM2TsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMpegtsM2TsMode

`func (o *DlnaTranscodingProfile) SetEnableMpegtsM2TsMode(v bool)`

SetEnableMpegtsM2TsMode sets EnableMpegtsM2TsMode field to given value.

### HasEnableMpegtsM2TsMode

`func (o *DlnaTranscodingProfile) HasEnableMpegtsM2TsMode() bool`

HasEnableMpegtsM2TsMode returns a boolean if a field has been set.

### GetTranscodeSeekInfo

`func (o *DlnaTranscodingProfile) GetTranscodeSeekInfo() string`

GetTranscodeSeekInfo returns the TranscodeSeekInfo field if non-nil, zero value otherwise.

### GetTranscodeSeekInfoOk

`func (o *DlnaTranscodingProfile) GetTranscodeSeekInfoOk() (*string, bool)`

GetTranscodeSeekInfoOk returns a tuple with the TranscodeSeekInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeSeekInfo

`func (o *DlnaTranscodingProfile) SetTranscodeSeekInfo(v string)`

SetTranscodeSeekInfo sets TranscodeSeekInfo field to given value.

### HasTranscodeSeekInfo

`func (o *DlnaTranscodingProfile) HasTranscodeSeekInfo() bool`

HasTranscodeSeekInfo returns a boolean if a field has been set.

### GetCopyTimestamps

`func (o *DlnaTranscodingProfile) GetCopyTimestamps() bool`

GetCopyTimestamps returns the CopyTimestamps field if non-nil, zero value otherwise.

### GetCopyTimestampsOk

`func (o *DlnaTranscodingProfile) GetCopyTimestampsOk() (*bool, bool)`

GetCopyTimestampsOk returns a tuple with the CopyTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTimestamps

`func (o *DlnaTranscodingProfile) SetCopyTimestamps(v bool)`

SetCopyTimestamps sets CopyTimestamps field to given value.

### HasCopyTimestamps

`func (o *DlnaTranscodingProfile) HasCopyTimestamps() bool`

HasCopyTimestamps returns a boolean if a field has been set.

### GetContext

`func (o *DlnaTranscodingProfile) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DlnaTranscodingProfile) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DlnaTranscodingProfile) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DlnaTranscodingProfile) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMaxAudioChannels

`func (o *DlnaTranscodingProfile) GetMaxAudioChannels() string`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *DlnaTranscodingProfile) GetMaxAudioChannelsOk() (*string, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *DlnaTranscodingProfile) SetMaxAudioChannels(v string)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *DlnaTranscodingProfile) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### GetMinSegments

`func (o *DlnaTranscodingProfile) GetMinSegments() int32`

GetMinSegments returns the MinSegments field if non-nil, zero value otherwise.

### GetMinSegmentsOk

`func (o *DlnaTranscodingProfile) GetMinSegmentsOk() (*int32, bool)`

GetMinSegmentsOk returns a tuple with the MinSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSegments

`func (o *DlnaTranscodingProfile) SetMinSegments(v int32)`

SetMinSegments sets MinSegments field to given value.

### HasMinSegments

`func (o *DlnaTranscodingProfile) HasMinSegments() bool`

HasMinSegments returns a boolean if a field has been set.

### GetSegmentLength

`func (o *DlnaTranscodingProfile) GetSegmentLength() int32`

GetSegmentLength returns the SegmentLength field if non-nil, zero value otherwise.

### GetSegmentLengthOk

`func (o *DlnaTranscodingProfile) GetSegmentLengthOk() (*int32, bool)`

GetSegmentLengthOk returns a tuple with the SegmentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentLength

`func (o *DlnaTranscodingProfile) SetSegmentLength(v int32)`

SetSegmentLength sets SegmentLength field to given value.

### HasSegmentLength

`func (o *DlnaTranscodingProfile) HasSegmentLength() bool`

HasSegmentLength returns a boolean if a field has been set.

### GetBreakOnNonKeyFrames

`func (o *DlnaTranscodingProfile) GetBreakOnNonKeyFrames() bool`

GetBreakOnNonKeyFrames returns the BreakOnNonKeyFrames field if non-nil, zero value otherwise.

### GetBreakOnNonKeyFramesOk

`func (o *DlnaTranscodingProfile) GetBreakOnNonKeyFramesOk() (*bool, bool)`

GetBreakOnNonKeyFramesOk returns a tuple with the BreakOnNonKeyFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakOnNonKeyFrames

`func (o *DlnaTranscodingProfile) SetBreakOnNonKeyFrames(v bool)`

SetBreakOnNonKeyFrames sets BreakOnNonKeyFrames field to given value.

### HasBreakOnNonKeyFrames

`func (o *DlnaTranscodingProfile) HasBreakOnNonKeyFrames() bool`

HasBreakOnNonKeyFrames returns a boolean if a field has been set.

### GetManifestSubtitles

`func (o *DlnaTranscodingProfile) GetManifestSubtitles() string`

GetManifestSubtitles returns the ManifestSubtitles field if non-nil, zero value otherwise.

### GetManifestSubtitlesOk

`func (o *DlnaTranscodingProfile) GetManifestSubtitlesOk() (*string, bool)`

GetManifestSubtitlesOk returns a tuple with the ManifestSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSubtitles

`func (o *DlnaTranscodingProfile) SetManifestSubtitles(v string)`

SetManifestSubtitles sets ManifestSubtitles field to given value.

### HasManifestSubtitles

`func (o *DlnaTranscodingProfile) HasManifestSubtitles() bool`

HasManifestSubtitles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



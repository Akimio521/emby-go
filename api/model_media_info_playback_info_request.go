/*
Emby Server API

Explore the Emby Server API

API version: 4.1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaInfoPlaybackInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaInfoPlaybackInfoRequest{}

// MediaInfoPlaybackInfoRequest struct for MediaInfoPlaybackInfoRequest
type MediaInfoPlaybackInfoRequest struct {
	Id *string `json:"Id,omitempty"`
	UserId *string `json:"UserId,omitempty"`
	MaxStreamingBitrate NullableInt64 `json:"MaxStreamingBitrate,omitempty"`
	StartTimeTicks NullableInt64 `json:"StartTimeTicks,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	MaxAudioChannels NullableInt32 `json:"MaxAudioChannels,omitempty"`
	MediaSourceId *string `json:"MediaSourceId,omitempty"`
	LiveStreamId *string `json:"LiveStreamId,omitempty"`
	DeviceProfile *DlnaDeviceProfile `json:"DeviceProfile,omitempty"`
	EnableDirectPlay *bool `json:"EnableDirectPlay,omitempty"`
	EnableDirectStream *bool `json:"EnableDirectStream,omitempty"`
	EnableTranscoding *bool `json:"EnableTranscoding,omitempty"`
	AllowVideoStreamCopy *bool `json:"AllowVideoStreamCopy,omitempty"`
	AllowAudioStreamCopy *bool `json:"AllowAudioStreamCopy,omitempty"`
	IsPlayback *bool `json:"IsPlayback,omitempty"`
	AutoOpenLiveStream *bool `json:"AutoOpenLiveStream,omitempty"`
	DirectPlayProtocols []string `json:"DirectPlayProtocols,omitempty"`
}

// NewMediaInfoPlaybackInfoRequest instantiates a new MediaInfoPlaybackInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoPlaybackInfoRequest() *MediaInfoPlaybackInfoRequest {
	this := MediaInfoPlaybackInfoRequest{}
	return &this
}

// NewMediaInfoPlaybackInfoRequestWithDefaults instantiates a new MediaInfoPlaybackInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoPlaybackInfoRequestWithDefaults() *MediaInfoPlaybackInfoRequest {
	this := MediaInfoPlaybackInfoRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MediaInfoPlaybackInfoRequest) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MediaInfoPlaybackInfoRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoPlaybackInfoRequest) GetMaxStreamingBitrate() int64 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoPlaybackInfoRequest) GetMaxStreamingBitrateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt64 and assigns it to the MaxStreamingBitrate field.
func (o *MediaInfoPlaybackInfoRequest) SetMaxStreamingBitrate(v int64) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *MediaInfoPlaybackInfoRequest) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *MediaInfoPlaybackInfoRequest) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetStartTimeTicks returns the StartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoPlaybackInfoRequest) GetStartTimeTicks() int64 {
	if o == nil || IsNil(o.StartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTimeTicks.Get()
}

// GetStartTimeTicksOk returns a tuple with the StartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoPlaybackInfoRequest) GetStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeTicks.Get(), o.StartTimeTicks.IsSet()
}

// HasStartTimeTicks returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasStartTimeTicks() bool {
	if o != nil && o.StartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the StartTimeTicks field.
func (o *MediaInfoPlaybackInfoRequest) SetStartTimeTicks(v int64) {
	o.StartTimeTicks.Set(&v)
}
// SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil
func (o *MediaInfoPlaybackInfoRequest) SetStartTimeTicksNil() {
	o.StartTimeTicks.Set(nil)
}

// UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
func (o *MediaInfoPlaybackInfoRequest) UnsetStartTimeTicks() {
	o.StartTimeTicks.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoPlaybackInfoRequest) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoPlaybackInfoRequest) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *MediaInfoPlaybackInfoRequest) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *MediaInfoPlaybackInfoRequest) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *MediaInfoPlaybackInfoRequest) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoPlaybackInfoRequest) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoPlaybackInfoRequest) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *MediaInfoPlaybackInfoRequest) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *MediaInfoPlaybackInfoRequest) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *MediaInfoPlaybackInfoRequest) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMaxAudioChannels returns the MaxAudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoPlaybackInfoRequest) GetMaxAudioChannels() int32 {
	if o == nil || IsNil(o.MaxAudioChannels.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAudioChannels.Get()
}

// GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoPlaybackInfoRequest) GetMaxAudioChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAudioChannels.Get(), o.MaxAudioChannels.IsSet()
}

// HasMaxAudioChannels returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasMaxAudioChannels() bool {
	if o != nil && o.MaxAudioChannels.IsSet() {
		return true
	}

	return false
}

// SetMaxAudioChannels gets a reference to the given NullableInt32 and assigns it to the MaxAudioChannels field.
func (o *MediaInfoPlaybackInfoRequest) SetMaxAudioChannels(v int32) {
	o.MaxAudioChannels.Set(&v)
}
// SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil
func (o *MediaInfoPlaybackInfoRequest) SetMaxAudioChannelsNil() {
	o.MaxAudioChannels.Set(nil)
}

// UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
func (o *MediaInfoPlaybackInfoRequest) UnsetMaxAudioChannels() {
	o.MaxAudioChannels.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId) {
		var ret string
		return ret
	}
	return *o.MediaSourceId
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetMediaSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSourceId) {
		return nil, false
	}
	return o.MediaSourceId, true
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasMediaSourceId() bool {
	if o != nil && !IsNil(o.MediaSourceId) {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given string and assigns it to the MediaSourceId field.
func (o *MediaInfoPlaybackInfoRequest) SetMediaSourceId(v string) {
	o.MediaSourceId = &v
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId) {
		var ret string
		return ret
	}
	return *o.LiveStreamId
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetLiveStreamIdOk() (*string, bool) {
	if o == nil || IsNil(o.LiveStreamId) {
		return nil, false
	}
	return o.LiveStreamId, true
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasLiveStreamId() bool {
	if o != nil && !IsNil(o.LiveStreamId) {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given string and assigns it to the LiveStreamId field.
func (o *MediaInfoPlaybackInfoRequest) SetLiveStreamId(v string) {
	o.LiveStreamId = &v
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetDeviceProfile() DlnaDeviceProfile {
	if o == nil || IsNil(o.DeviceProfile) {
		var ret DlnaDeviceProfile
		return ret
	}
	return *o.DeviceProfile
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetDeviceProfileOk() (*DlnaDeviceProfile, bool) {
	if o == nil || IsNil(o.DeviceProfile) {
		return nil, false
	}
	return o.DeviceProfile, true
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasDeviceProfile() bool {
	if o != nil && !IsNil(o.DeviceProfile) {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given DlnaDeviceProfile and assigns it to the DeviceProfile field.
func (o *MediaInfoPlaybackInfoRequest) SetDeviceProfile(v DlnaDeviceProfile) {
	o.DeviceProfile = &v
}

// GetEnableDirectPlay returns the EnableDirectPlay field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectPlay() bool {
	if o == nil || IsNil(o.EnableDirectPlay) {
		var ret bool
		return ret
	}
	return *o.EnableDirectPlay
}

// GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectPlayOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDirectPlay) {
		return nil, false
	}
	return o.EnableDirectPlay, true
}

// HasEnableDirectPlay returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasEnableDirectPlay() bool {
	if o != nil && !IsNil(o.EnableDirectPlay) {
		return true
	}

	return false
}

// SetEnableDirectPlay gets a reference to the given bool and assigns it to the EnableDirectPlay field.
func (o *MediaInfoPlaybackInfoRequest) SetEnableDirectPlay(v bool) {
	o.EnableDirectPlay = &v
}

// GetEnableDirectStream returns the EnableDirectStream field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectStream() bool {
	if o == nil || IsNil(o.EnableDirectStream) {
		var ret bool
		return ret
	}
	return *o.EnableDirectStream
}

// GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetEnableDirectStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDirectStream) {
		return nil, false
	}
	return o.EnableDirectStream, true
}

// HasEnableDirectStream returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasEnableDirectStream() bool {
	if o != nil && !IsNil(o.EnableDirectStream) {
		return true
	}

	return false
}

// SetEnableDirectStream gets a reference to the given bool and assigns it to the EnableDirectStream field.
func (o *MediaInfoPlaybackInfoRequest) SetEnableDirectStream(v bool) {
	o.EnableDirectStream = &v
}

// GetEnableTranscoding returns the EnableTranscoding field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetEnableTranscoding() bool {
	if o == nil || IsNil(o.EnableTranscoding) {
		var ret bool
		return ret
	}
	return *o.EnableTranscoding
}

// GetEnableTranscodingOk returns a tuple with the EnableTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetEnableTranscodingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTranscoding) {
		return nil, false
	}
	return o.EnableTranscoding, true
}

// HasEnableTranscoding returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasEnableTranscoding() bool {
	if o != nil && !IsNil(o.EnableTranscoding) {
		return true
	}

	return false
}

// SetEnableTranscoding gets a reference to the given bool and assigns it to the EnableTranscoding field.
func (o *MediaInfoPlaybackInfoRequest) SetEnableTranscoding(v bool) {
	o.EnableTranscoding = &v
}

// GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetAllowVideoStreamCopy() bool {
	if o == nil || IsNil(o.AllowVideoStreamCopy) {
		var ret bool
		return ret
	}
	return *o.AllowVideoStreamCopy
}

// GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetAllowVideoStreamCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowVideoStreamCopy) {
		return nil, false
	}
	return o.AllowVideoStreamCopy, true
}

// HasAllowVideoStreamCopy returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasAllowVideoStreamCopy() bool {
	if o != nil && !IsNil(o.AllowVideoStreamCopy) {
		return true
	}

	return false
}

// SetAllowVideoStreamCopy gets a reference to the given bool and assigns it to the AllowVideoStreamCopy field.
func (o *MediaInfoPlaybackInfoRequest) SetAllowVideoStreamCopy(v bool) {
	o.AllowVideoStreamCopy = &v
}

// GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetAllowAudioStreamCopy() bool {
	if o == nil || IsNil(o.AllowAudioStreamCopy) {
		var ret bool
		return ret
	}
	return *o.AllowAudioStreamCopy
}

// GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetAllowAudioStreamCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAudioStreamCopy) {
		return nil, false
	}
	return o.AllowAudioStreamCopy, true
}

// HasAllowAudioStreamCopy returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasAllowAudioStreamCopy() bool {
	if o != nil && !IsNil(o.AllowAudioStreamCopy) {
		return true
	}

	return false
}

// SetAllowAudioStreamCopy gets a reference to the given bool and assigns it to the AllowAudioStreamCopy field.
func (o *MediaInfoPlaybackInfoRequest) SetAllowAudioStreamCopy(v bool) {
	o.AllowAudioStreamCopy = &v
}

// GetIsPlayback returns the IsPlayback field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetIsPlayback() bool {
	if o == nil || IsNil(o.IsPlayback) {
		var ret bool
		return ret
	}
	return *o.IsPlayback
}

// GetIsPlaybackOk returns a tuple with the IsPlayback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetIsPlaybackOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlayback) {
		return nil, false
	}
	return o.IsPlayback, true
}

// HasIsPlayback returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasIsPlayback() bool {
	if o != nil && !IsNil(o.IsPlayback) {
		return true
	}

	return false
}

// SetIsPlayback gets a reference to the given bool and assigns it to the IsPlayback field.
func (o *MediaInfoPlaybackInfoRequest) SetIsPlayback(v bool) {
	o.IsPlayback = &v
}

// GetAutoOpenLiveStream returns the AutoOpenLiveStream field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetAutoOpenLiveStream() bool {
	if o == nil || IsNil(o.AutoOpenLiveStream) {
		var ret bool
		return ret
	}
	return *o.AutoOpenLiveStream
}

// GetAutoOpenLiveStreamOk returns a tuple with the AutoOpenLiveStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetAutoOpenLiveStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoOpenLiveStream) {
		return nil, false
	}
	return o.AutoOpenLiveStream, true
}

// HasAutoOpenLiveStream returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasAutoOpenLiveStream() bool {
	if o != nil && !IsNil(o.AutoOpenLiveStream) {
		return true
	}

	return false
}

// SetAutoOpenLiveStream gets a reference to the given bool and assigns it to the AutoOpenLiveStream field.
func (o *MediaInfoPlaybackInfoRequest) SetAutoOpenLiveStream(v bool) {
	o.AutoOpenLiveStream = &v
}

// GetDirectPlayProtocols returns the DirectPlayProtocols field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoRequest) GetDirectPlayProtocols() []string {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		var ret []string
		return ret
	}
	return o.DirectPlayProtocols
}

// GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoRequest) GetDirectPlayProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		return nil, false
	}
	return o.DirectPlayProtocols, true
}

// HasDirectPlayProtocols returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoRequest) HasDirectPlayProtocols() bool {
	if o != nil && !IsNil(o.DirectPlayProtocols) {
		return true
	}

	return false
}

// SetDirectPlayProtocols gets a reference to the given []string and assigns it to the DirectPlayProtocols field.
func (o *MediaInfoPlaybackInfoRequest) SetDirectPlayProtocols(v []string) {
	o.DirectPlayProtocols = v
}

func (o MediaInfoPlaybackInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaInfoPlaybackInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if o.MaxStreamingBitrate.IsSet() {
		toSerialize["MaxStreamingBitrate"] = o.MaxStreamingBitrate.Get()
	}
	if o.StartTimeTicks.IsSet() {
		toSerialize["StartTimeTicks"] = o.StartTimeTicks.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if o.MaxAudioChannels.IsSet() {
		toSerialize["MaxAudioChannels"] = o.MaxAudioChannels.Get()
	}
	if !IsNil(o.MediaSourceId) {
		toSerialize["MediaSourceId"] = o.MediaSourceId
	}
	if !IsNil(o.LiveStreamId) {
		toSerialize["LiveStreamId"] = o.LiveStreamId
	}
	if !IsNil(o.DeviceProfile) {
		toSerialize["DeviceProfile"] = o.DeviceProfile
	}
	if !IsNil(o.EnableDirectPlay) {
		toSerialize["EnableDirectPlay"] = o.EnableDirectPlay
	}
	if !IsNil(o.EnableDirectStream) {
		toSerialize["EnableDirectStream"] = o.EnableDirectStream
	}
	if !IsNil(o.EnableTranscoding) {
		toSerialize["EnableTranscoding"] = o.EnableTranscoding
	}
	if !IsNil(o.AllowVideoStreamCopy) {
		toSerialize["AllowVideoStreamCopy"] = o.AllowVideoStreamCopy
	}
	if !IsNil(o.AllowAudioStreamCopy) {
		toSerialize["AllowAudioStreamCopy"] = o.AllowAudioStreamCopy
	}
	if !IsNil(o.IsPlayback) {
		toSerialize["IsPlayback"] = o.IsPlayback
	}
	if !IsNil(o.AutoOpenLiveStream) {
		toSerialize["AutoOpenLiveStream"] = o.AutoOpenLiveStream
	}
	if !IsNil(o.DirectPlayProtocols) {
		toSerialize["DirectPlayProtocols"] = o.DirectPlayProtocols
	}
	return toSerialize, nil
}

type NullableMediaInfoPlaybackInfoRequest struct {
	value *MediaInfoPlaybackInfoRequest
	isSet bool
}

func (v NullableMediaInfoPlaybackInfoRequest) Get() *MediaInfoPlaybackInfoRequest {
	return v.value
}

func (v *NullableMediaInfoPlaybackInfoRequest) Set(val *MediaInfoPlaybackInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoPlaybackInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoPlaybackInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoPlaybackInfoRequest(val *MediaInfoPlaybackInfoRequest) *NullableMediaInfoPlaybackInfoRequest {
	return &NullableMediaInfoPlaybackInfoRequest{value: val, isSet: true}
}

func (v NullableMediaInfoPlaybackInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoPlaybackInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



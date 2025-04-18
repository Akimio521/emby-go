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

// checks if the MediaInfoLiveStreamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaInfoLiveStreamRequest{}

// MediaInfoLiveStreamRequest struct for MediaInfoLiveStreamRequest
type MediaInfoLiveStreamRequest struct {
	OpenToken *string `json:"OpenToken,omitempty"`
	UserId *string `json:"UserId,omitempty"`
	PlaySessionId *string `json:"PlaySessionId,omitempty"`
	MaxStreamingBitrate NullableInt64 `json:"MaxStreamingBitrate,omitempty"`
	StartTimeTicks NullableInt64 `json:"StartTimeTicks,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	MaxAudioChannels NullableInt32 `json:"MaxAudioChannels,omitempty"`
	ItemId *int64 `json:"ItemId,omitempty"`
	DeviceProfile *DlnaDeviceProfile `json:"DeviceProfile,omitempty"`
	EnableDirectPlay *bool `json:"EnableDirectPlay,omitempty"`
	EnableDirectStream *bool `json:"EnableDirectStream,omitempty"`
	EnableTranscoding *bool `json:"EnableTranscoding,omitempty"`
	AllowVideoStreamCopy *bool `json:"AllowVideoStreamCopy,omitempty"`
	AllowAudioStreamCopy *bool `json:"AllowAudioStreamCopy,omitempty"`
	DirectPlayProtocols []string `json:"DirectPlayProtocols,omitempty"`
}

// NewMediaInfoLiveStreamRequest instantiates a new MediaInfoLiveStreamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoLiveStreamRequest() *MediaInfoLiveStreamRequest {
	this := MediaInfoLiveStreamRequest{}
	return &this
}

// NewMediaInfoLiveStreamRequestWithDefaults instantiates a new MediaInfoLiveStreamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoLiveStreamRequestWithDefaults() *MediaInfoLiveStreamRequest {
	this := MediaInfoLiveStreamRequest{}
	return &this
}

// GetOpenToken returns the OpenToken field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetOpenToken() string {
	if o == nil || IsNil(o.OpenToken) {
		var ret string
		return ret
	}
	return *o.OpenToken
}

// GetOpenTokenOk returns a tuple with the OpenToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetOpenTokenOk() (*string, bool) {
	if o == nil || IsNil(o.OpenToken) {
		return nil, false
	}
	return o.OpenToken, true
}

// HasOpenToken returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasOpenToken() bool {
	if o != nil && !IsNil(o.OpenToken) {
		return true
	}

	return false
}

// SetOpenToken gets a reference to the given string and assigns it to the OpenToken field.
func (o *MediaInfoLiveStreamRequest) SetOpenToken(v string) {
	o.OpenToken = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MediaInfoLiveStreamRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId) {
		var ret string
		return ret
	}
	return *o.PlaySessionId
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetPlaySessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaySessionId) {
		return nil, false
	}
	return o.PlaySessionId, true
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasPlaySessionId() bool {
	if o != nil && !IsNil(o.PlaySessionId) {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given string and assigns it to the PlaySessionId field.
func (o *MediaInfoLiveStreamRequest) SetPlaySessionId(v string) {
	o.PlaySessionId = &v
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoLiveStreamRequest) GetMaxStreamingBitrate() int64 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoLiveStreamRequest) GetMaxStreamingBitrateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt64 and assigns it to the MaxStreamingBitrate field.
func (o *MediaInfoLiveStreamRequest) SetMaxStreamingBitrate(v int64) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *MediaInfoLiveStreamRequest) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *MediaInfoLiveStreamRequest) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetStartTimeTicks returns the StartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoLiveStreamRequest) GetStartTimeTicks() int64 {
	if o == nil || IsNil(o.StartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTimeTicks.Get()
}

// GetStartTimeTicksOk returns a tuple with the StartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoLiveStreamRequest) GetStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeTicks.Get(), o.StartTimeTicks.IsSet()
}

// HasStartTimeTicks returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasStartTimeTicks() bool {
	if o != nil && o.StartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the StartTimeTicks field.
func (o *MediaInfoLiveStreamRequest) SetStartTimeTicks(v int64) {
	o.StartTimeTicks.Set(&v)
}
// SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil
func (o *MediaInfoLiveStreamRequest) SetStartTimeTicksNil() {
	o.StartTimeTicks.Set(nil)
}

// UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
func (o *MediaInfoLiveStreamRequest) UnsetStartTimeTicks() {
	o.StartTimeTicks.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoLiveStreamRequest) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoLiveStreamRequest) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *MediaInfoLiveStreamRequest) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *MediaInfoLiveStreamRequest) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *MediaInfoLiveStreamRequest) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoLiveStreamRequest) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoLiveStreamRequest) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *MediaInfoLiveStreamRequest) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *MediaInfoLiveStreamRequest) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *MediaInfoLiveStreamRequest) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMaxAudioChannels returns the MaxAudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoLiveStreamRequest) GetMaxAudioChannels() int32 {
	if o == nil || IsNil(o.MaxAudioChannels.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAudioChannels.Get()
}

// GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoLiveStreamRequest) GetMaxAudioChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAudioChannels.Get(), o.MaxAudioChannels.IsSet()
}

// HasMaxAudioChannels returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasMaxAudioChannels() bool {
	if o != nil && o.MaxAudioChannels.IsSet() {
		return true
	}

	return false
}

// SetMaxAudioChannels gets a reference to the given NullableInt32 and assigns it to the MaxAudioChannels field.
func (o *MediaInfoLiveStreamRequest) SetMaxAudioChannels(v int32) {
	o.MaxAudioChannels.Set(&v)
}
// SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil
func (o *MediaInfoLiveStreamRequest) SetMaxAudioChannelsNil() {
	o.MaxAudioChannels.Set(nil)
}

// UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
func (o *MediaInfoLiveStreamRequest) UnsetMaxAudioChannels() {
	o.MaxAudioChannels.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *MediaInfoLiveStreamRequest) SetItemId(v int64) {
	o.ItemId = &v
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetDeviceProfile() DlnaDeviceProfile {
	if o == nil || IsNil(o.DeviceProfile) {
		var ret DlnaDeviceProfile
		return ret
	}
	return *o.DeviceProfile
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetDeviceProfileOk() (*DlnaDeviceProfile, bool) {
	if o == nil || IsNil(o.DeviceProfile) {
		return nil, false
	}
	return o.DeviceProfile, true
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasDeviceProfile() bool {
	if o != nil && !IsNil(o.DeviceProfile) {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given DlnaDeviceProfile and assigns it to the DeviceProfile field.
func (o *MediaInfoLiveStreamRequest) SetDeviceProfile(v DlnaDeviceProfile) {
	o.DeviceProfile = &v
}

// GetEnableDirectPlay returns the EnableDirectPlay field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetEnableDirectPlay() bool {
	if o == nil || IsNil(o.EnableDirectPlay) {
		var ret bool
		return ret
	}
	return *o.EnableDirectPlay
}

// GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetEnableDirectPlayOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDirectPlay) {
		return nil, false
	}
	return o.EnableDirectPlay, true
}

// HasEnableDirectPlay returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasEnableDirectPlay() bool {
	if o != nil && !IsNil(o.EnableDirectPlay) {
		return true
	}

	return false
}

// SetEnableDirectPlay gets a reference to the given bool and assigns it to the EnableDirectPlay field.
func (o *MediaInfoLiveStreamRequest) SetEnableDirectPlay(v bool) {
	o.EnableDirectPlay = &v
}

// GetEnableDirectStream returns the EnableDirectStream field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetEnableDirectStream() bool {
	if o == nil || IsNil(o.EnableDirectStream) {
		var ret bool
		return ret
	}
	return *o.EnableDirectStream
}

// GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetEnableDirectStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDirectStream) {
		return nil, false
	}
	return o.EnableDirectStream, true
}

// HasEnableDirectStream returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasEnableDirectStream() bool {
	if o != nil && !IsNil(o.EnableDirectStream) {
		return true
	}

	return false
}

// SetEnableDirectStream gets a reference to the given bool and assigns it to the EnableDirectStream field.
func (o *MediaInfoLiveStreamRequest) SetEnableDirectStream(v bool) {
	o.EnableDirectStream = &v
}

// GetEnableTranscoding returns the EnableTranscoding field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetEnableTranscoding() bool {
	if o == nil || IsNil(o.EnableTranscoding) {
		var ret bool
		return ret
	}
	return *o.EnableTranscoding
}

// GetEnableTranscodingOk returns a tuple with the EnableTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetEnableTranscodingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTranscoding) {
		return nil, false
	}
	return o.EnableTranscoding, true
}

// HasEnableTranscoding returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasEnableTranscoding() bool {
	if o != nil && !IsNil(o.EnableTranscoding) {
		return true
	}

	return false
}

// SetEnableTranscoding gets a reference to the given bool and assigns it to the EnableTranscoding field.
func (o *MediaInfoLiveStreamRequest) SetEnableTranscoding(v bool) {
	o.EnableTranscoding = &v
}

// GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetAllowVideoStreamCopy() bool {
	if o == nil || IsNil(o.AllowVideoStreamCopy) {
		var ret bool
		return ret
	}
	return *o.AllowVideoStreamCopy
}

// GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetAllowVideoStreamCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowVideoStreamCopy) {
		return nil, false
	}
	return o.AllowVideoStreamCopy, true
}

// HasAllowVideoStreamCopy returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasAllowVideoStreamCopy() bool {
	if o != nil && !IsNil(o.AllowVideoStreamCopy) {
		return true
	}

	return false
}

// SetAllowVideoStreamCopy gets a reference to the given bool and assigns it to the AllowVideoStreamCopy field.
func (o *MediaInfoLiveStreamRequest) SetAllowVideoStreamCopy(v bool) {
	o.AllowVideoStreamCopy = &v
}

// GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetAllowAudioStreamCopy() bool {
	if o == nil || IsNil(o.AllowAudioStreamCopy) {
		var ret bool
		return ret
	}
	return *o.AllowAudioStreamCopy
}

// GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetAllowAudioStreamCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAudioStreamCopy) {
		return nil, false
	}
	return o.AllowAudioStreamCopy, true
}

// HasAllowAudioStreamCopy returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasAllowAudioStreamCopy() bool {
	if o != nil && !IsNil(o.AllowAudioStreamCopy) {
		return true
	}

	return false
}

// SetAllowAudioStreamCopy gets a reference to the given bool and assigns it to the AllowAudioStreamCopy field.
func (o *MediaInfoLiveStreamRequest) SetAllowAudioStreamCopy(v bool) {
	o.AllowAudioStreamCopy = &v
}

// GetDirectPlayProtocols returns the DirectPlayProtocols field value if set, zero value otherwise.
func (o *MediaInfoLiveStreamRequest) GetDirectPlayProtocols() []string {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		var ret []string
		return ret
	}
	return o.DirectPlayProtocols
}

// GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoLiveStreamRequest) GetDirectPlayProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		return nil, false
	}
	return o.DirectPlayProtocols, true
}

// HasDirectPlayProtocols returns a boolean if a field has been set.
func (o *MediaInfoLiveStreamRequest) HasDirectPlayProtocols() bool {
	if o != nil && !IsNil(o.DirectPlayProtocols) {
		return true
	}

	return false
}

// SetDirectPlayProtocols gets a reference to the given []string and assigns it to the DirectPlayProtocols field.
func (o *MediaInfoLiveStreamRequest) SetDirectPlayProtocols(v []string) {
	o.DirectPlayProtocols = v
}

func (o MediaInfoLiveStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaInfoLiveStreamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenToken) {
		toSerialize["OpenToken"] = o.OpenToken
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.PlaySessionId) {
		toSerialize["PlaySessionId"] = o.PlaySessionId
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
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
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
	if !IsNil(o.DirectPlayProtocols) {
		toSerialize["DirectPlayProtocols"] = o.DirectPlayProtocols
	}
	return toSerialize, nil
}

type NullableMediaInfoLiveStreamRequest struct {
	value *MediaInfoLiveStreamRequest
	isSet bool
}

func (v NullableMediaInfoLiveStreamRequest) Get() *MediaInfoLiveStreamRequest {
	return v.value
}

func (v *NullableMediaInfoLiveStreamRequest) Set(val *MediaInfoLiveStreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoLiveStreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoLiveStreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoLiveStreamRequest(val *MediaInfoLiveStreamRequest) *NullableMediaInfoLiveStreamRequest {
	return &NullableMediaInfoLiveStreamRequest{value: val, isSet: true}
}

func (v NullableMediaInfoLiveStreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoLiveStreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the PlaybackProgressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybackProgressInfo{}

// PlaybackProgressInfo struct for PlaybackProgressInfo
type PlaybackProgressInfo struct {
	CanSeek *bool `json:"CanSeek,omitempty"`
	Item *BaseItemDto `json:"Item,omitempty"`
	NowPlayingQueue []QueueItem `json:"NowPlayingQueue,omitempty"`
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	SessionId *string `json:"SessionId,omitempty"`
	MediaSourceId *string `json:"MediaSourceId,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	IsPaused *bool `json:"IsPaused,omitempty"`
	IsMuted *bool `json:"IsMuted,omitempty"`
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	RunTimeTicks NullableInt64 `json:"RunTimeTicks,omitempty"`
	PlaybackStartTimeTicks NullableInt64 `json:"PlaybackStartTimeTicks,omitempty"`
	VolumeLevel NullableInt32 `json:"VolumeLevel,omitempty"`
	Brightness NullableInt32 `json:"Brightness,omitempty"`
	AspectRatio *string `json:"AspectRatio,omitempty"`
	PlayMethod *string `json:"PlayMethod,omitempty"`
	LiveStreamId *string `json:"LiveStreamId,omitempty"`
	PlaySessionId *string `json:"PlaySessionId,omitempty"`
	RepeatMode *string `json:"RepeatMode,omitempty"`
}

// NewPlaybackProgressInfo instantiates a new PlaybackProgressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybackProgressInfo() *PlaybackProgressInfo {
	this := PlaybackProgressInfo{}
	return &this
}

// NewPlaybackProgressInfoWithDefaults instantiates a new PlaybackProgressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybackProgressInfoWithDefaults() *PlaybackProgressInfo {
	this := PlaybackProgressInfo{}
	return &this
}

// GetCanSeek returns the CanSeek field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetCanSeek() bool {
	if o == nil || IsNil(o.CanSeek) {
		var ret bool
		return ret
	}
	return *o.CanSeek
}

// GetCanSeekOk returns a tuple with the CanSeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetCanSeekOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSeek) {
		return nil, false
	}
	return o.CanSeek, true
}

// HasCanSeek returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasCanSeek() bool {
	if o != nil && !IsNil(o.CanSeek) {
		return true
	}

	return false
}

// SetCanSeek gets a reference to the given bool and assigns it to the CanSeek field.
func (o *PlaybackProgressInfo) SetCanSeek(v bool) {
	o.CanSeek = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetItem() BaseItemDto {
	if o == nil || IsNil(o.Item) {
		var ret BaseItemDto
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetItemOk() (*BaseItemDto, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given BaseItemDto and assigns it to the Item field.
func (o *PlaybackProgressInfo) SetItem(v BaseItemDto) {
	o.Item = &v
}

// GetNowPlayingQueue returns the NowPlayingQueue field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetNowPlayingQueue() []QueueItem {
	if o == nil || IsNil(o.NowPlayingQueue) {
		var ret []QueueItem
		return ret
	}
	return o.NowPlayingQueue
}

// GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetNowPlayingQueueOk() ([]QueueItem, bool) {
	if o == nil || IsNil(o.NowPlayingQueue) {
		return nil, false
	}
	return o.NowPlayingQueue, true
}

// HasNowPlayingQueue returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasNowPlayingQueue() bool {
	if o != nil && !IsNil(o.NowPlayingQueue) {
		return true
	}

	return false
}

// SetNowPlayingQueue gets a reference to the given []QueueItem and assigns it to the NowPlayingQueue field.
func (o *PlaybackProgressInfo) SetNowPlayingQueue(v []QueueItem) {
	o.NowPlayingQueue = v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *PlaybackProgressInfo) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *PlaybackProgressInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *PlaybackProgressInfo) SetSessionId(v string) {
	o.SessionId = &v
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId) {
		var ret string
		return ret
	}
	return *o.MediaSourceId
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSourceId) {
		return nil, false
	}
	return o.MediaSourceId, true
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasMediaSourceId() bool {
	if o != nil && !IsNil(o.MediaSourceId) {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given string and assigns it to the MediaSourceId field.
func (o *PlaybackProgressInfo) SetMediaSourceId(v string) {
	o.MediaSourceId = &v
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *PlaybackProgressInfo) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *PlaybackProgressInfo) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *PlaybackProgressInfo) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *PlaybackProgressInfo) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused) {
		var ret bool
		return ret
	}
	return *o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetIsPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaused) {
		return nil, false
	}
	return o.IsPaused, true
}

// HasIsPaused returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasIsPaused() bool {
	if o != nil && !IsNil(o.IsPaused) {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given bool and assigns it to the IsPaused field.
func (o *PlaybackProgressInfo) SetIsPaused(v bool) {
	o.IsPaused = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *PlaybackProgressInfo) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *PlaybackProgressInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *PlaybackProgressInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetRunTimeTicks returns the RunTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetRunTimeTicks() int64 {
	if o == nil || IsNil(o.RunTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.RunTimeTicks.Get()
}

// GetRunTimeTicksOk returns a tuple with the RunTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetRunTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunTimeTicks.Get(), o.RunTimeTicks.IsSet()
}

// HasRunTimeTicks returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasRunTimeTicks() bool {
	if o != nil && o.RunTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetRunTimeTicks gets a reference to the given NullableInt64 and assigns it to the RunTimeTicks field.
func (o *PlaybackProgressInfo) SetRunTimeTicks(v int64) {
	o.RunTimeTicks.Set(&v)
}
// SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil
func (o *PlaybackProgressInfo) SetRunTimeTicksNil() {
	o.RunTimeTicks.Set(nil)
}

// UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetRunTimeTicks() {
	o.RunTimeTicks.Unset()
}

// GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetPlaybackStartTimeTicks() int64 {
	if o == nil || IsNil(o.PlaybackStartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PlaybackStartTimeTicks.Get()
}

// GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetPlaybackStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaybackStartTimeTicks.Get(), o.PlaybackStartTimeTicks.IsSet()
}

// HasPlaybackStartTimeTicks returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasPlaybackStartTimeTicks() bool {
	if o != nil && o.PlaybackStartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetPlaybackStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the PlaybackStartTimeTicks field.
func (o *PlaybackProgressInfo) SetPlaybackStartTimeTicks(v int64) {
	o.PlaybackStartTimeTicks.Set(&v)
}
// SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil
func (o *PlaybackProgressInfo) SetPlaybackStartTimeTicksNil() {
	o.PlaybackStartTimeTicks.Set(nil)
}

// UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetPlaybackStartTimeTicks() {
	o.PlaybackStartTimeTicks.Unset()
}

// GetVolumeLevel returns the VolumeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetVolumeLevel() int32 {
	if o == nil || IsNil(o.VolumeLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumeLevel.Get()
}

// GetVolumeLevelOk returns a tuple with the VolumeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetVolumeLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeLevel.Get(), o.VolumeLevel.IsSet()
}

// HasVolumeLevel returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasVolumeLevel() bool {
	if o != nil && o.VolumeLevel.IsSet() {
		return true
	}

	return false
}

// SetVolumeLevel gets a reference to the given NullableInt32 and assigns it to the VolumeLevel field.
func (o *PlaybackProgressInfo) SetVolumeLevel(v int32) {
	o.VolumeLevel.Set(&v)
}
// SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil
func (o *PlaybackProgressInfo) SetVolumeLevelNil() {
	o.VolumeLevel.Set(nil)
}

// UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetVolumeLevel() {
	o.VolumeLevel.Unset()
}

// GetBrightness returns the Brightness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackProgressInfo) GetBrightness() int32 {
	if o == nil || IsNil(o.Brightness.Get()) {
		var ret int32
		return ret
	}
	return *o.Brightness.Get()
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackProgressInfo) GetBrightnessOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brightness.Get(), o.Brightness.IsSet()
}

// HasBrightness returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasBrightness() bool {
	if o != nil && o.Brightness.IsSet() {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given NullableInt32 and assigns it to the Brightness field.
func (o *PlaybackProgressInfo) SetBrightness(v int32) {
	o.Brightness.Set(&v)
}
// SetBrightnessNil sets the value for Brightness to be an explicit nil
func (o *PlaybackProgressInfo) SetBrightnessNil() {
	o.Brightness.Set(nil)
}

// UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
func (o *PlaybackProgressInfo) UnsetBrightness() {
	o.Brightness.Unset()
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetAspectRatio() string {
	if o == nil || IsNil(o.AspectRatio) {
		var ret string
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetAspectRatioOk() (*string, bool) {
	if o == nil || IsNil(o.AspectRatio) {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasAspectRatio() bool {
	if o != nil && !IsNil(o.AspectRatio) {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given string and assigns it to the AspectRatio field.
func (o *PlaybackProgressInfo) SetAspectRatio(v string) {
	o.AspectRatio = &v
}

// GetPlayMethod returns the PlayMethod field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetPlayMethod() string {
	if o == nil || IsNil(o.PlayMethod) {
		var ret string
		return ret
	}
	return *o.PlayMethod
}

// GetPlayMethodOk returns a tuple with the PlayMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetPlayMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PlayMethod) {
		return nil, false
	}
	return o.PlayMethod, true
}

// HasPlayMethod returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasPlayMethod() bool {
	if o != nil && !IsNil(o.PlayMethod) {
		return true
	}

	return false
}

// SetPlayMethod gets a reference to the given string and assigns it to the PlayMethod field.
func (o *PlaybackProgressInfo) SetPlayMethod(v string) {
	o.PlayMethod = &v
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId) {
		var ret string
		return ret
	}
	return *o.LiveStreamId
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetLiveStreamIdOk() (*string, bool) {
	if o == nil || IsNil(o.LiveStreamId) {
		return nil, false
	}
	return o.LiveStreamId, true
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasLiveStreamId() bool {
	if o != nil && !IsNil(o.LiveStreamId) {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given string and assigns it to the LiveStreamId field.
func (o *PlaybackProgressInfo) SetLiveStreamId(v string) {
	o.LiveStreamId = &v
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId) {
		var ret string
		return ret
	}
	return *o.PlaySessionId
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetPlaySessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaySessionId) {
		return nil, false
	}
	return o.PlaySessionId, true
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasPlaySessionId() bool {
	if o != nil && !IsNil(o.PlaySessionId) {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given string and assigns it to the PlaySessionId field.
func (o *PlaybackProgressInfo) SetPlaySessionId(v string) {
	o.PlaySessionId = &v
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *PlaybackProgressInfo) GetRepeatMode() string {
	if o == nil || IsNil(o.RepeatMode) {
		var ret string
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackProgressInfo) GetRepeatModeOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *PlaybackProgressInfo) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given string and assigns it to the RepeatMode field.
func (o *PlaybackProgressInfo) SetRepeatMode(v string) {
	o.RepeatMode = &v
}

func (o PlaybackProgressInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybackProgressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanSeek) {
		toSerialize["CanSeek"] = o.CanSeek
	}
	if !IsNil(o.Item) {
		toSerialize["Item"] = o.Item
	}
	if !IsNil(o.NowPlayingQueue) {
		toSerialize["NowPlayingQueue"] = o.NowPlayingQueue
	}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.SessionId) {
		toSerialize["SessionId"] = o.SessionId
	}
	if !IsNil(o.MediaSourceId) {
		toSerialize["MediaSourceId"] = o.MediaSourceId
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if !IsNil(o.IsPaused) {
		toSerialize["IsPaused"] = o.IsPaused
	}
	if !IsNil(o.IsMuted) {
		toSerialize["IsMuted"] = o.IsMuted
	}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if o.RunTimeTicks.IsSet() {
		toSerialize["RunTimeTicks"] = o.RunTimeTicks.Get()
	}
	if o.PlaybackStartTimeTicks.IsSet() {
		toSerialize["PlaybackStartTimeTicks"] = o.PlaybackStartTimeTicks.Get()
	}
	if o.VolumeLevel.IsSet() {
		toSerialize["VolumeLevel"] = o.VolumeLevel.Get()
	}
	if o.Brightness.IsSet() {
		toSerialize["Brightness"] = o.Brightness.Get()
	}
	if !IsNil(o.AspectRatio) {
		toSerialize["AspectRatio"] = o.AspectRatio
	}
	if !IsNil(o.PlayMethod) {
		toSerialize["PlayMethod"] = o.PlayMethod
	}
	if !IsNil(o.LiveStreamId) {
		toSerialize["LiveStreamId"] = o.LiveStreamId
	}
	if !IsNil(o.PlaySessionId) {
		toSerialize["PlaySessionId"] = o.PlaySessionId
	}
	if !IsNil(o.RepeatMode) {
		toSerialize["RepeatMode"] = o.RepeatMode
	}
	return toSerialize, nil
}

type NullablePlaybackProgressInfo struct {
	value *PlaybackProgressInfo
	isSet bool
}

func (v NullablePlaybackProgressInfo) Get() *PlaybackProgressInfo {
	return v.value
}

func (v *NullablePlaybackProgressInfo) Set(val *PlaybackProgressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackProgressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackProgressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackProgressInfo(val *PlaybackProgressInfo) *NullablePlaybackProgressInfo {
	return &NullablePlaybackProgressInfo{value: val, isSet: true}
}

func (v NullablePlaybackProgressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackProgressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



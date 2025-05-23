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

// checks if the MediaInfoPlaybackInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaInfoPlaybackInfoResponse{}

// MediaInfoPlaybackInfoResponse struct for MediaInfoPlaybackInfoResponse
type MediaInfoPlaybackInfoResponse struct {
	MediaSources []MediaSourceInfo `json:"MediaSources,omitempty"`
	PlaySessionId *string `json:"PlaySessionId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty"`
}

// NewMediaInfoPlaybackInfoResponse instantiates a new MediaInfoPlaybackInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoPlaybackInfoResponse() *MediaInfoPlaybackInfoResponse {
	this := MediaInfoPlaybackInfoResponse{}
	return &this
}

// NewMediaInfoPlaybackInfoResponseWithDefaults instantiates a new MediaInfoPlaybackInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoPlaybackInfoResponseWithDefaults() *MediaInfoPlaybackInfoResponse {
	this := MediaInfoPlaybackInfoResponse{}
	return &this
}

// GetMediaSources returns the MediaSources field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoResponse) GetMediaSources() []MediaSourceInfo {
	if o == nil || IsNil(o.MediaSources) {
		var ret []MediaSourceInfo
		return ret
	}
	return o.MediaSources
}

// GetMediaSourcesOk returns a tuple with the MediaSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoResponse) GetMediaSourcesOk() ([]MediaSourceInfo, bool) {
	if o == nil || IsNil(o.MediaSources) {
		return nil, false
	}
	return o.MediaSources, true
}

// HasMediaSources returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoResponse) HasMediaSources() bool {
	if o != nil && !IsNil(o.MediaSources) {
		return true
	}

	return false
}

// SetMediaSources gets a reference to the given []MediaSourceInfo and assigns it to the MediaSources field.
func (o *MediaInfoPlaybackInfoResponse) SetMediaSources(v []MediaSourceInfo) {
	o.MediaSources = v
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoResponse) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId) {
		var ret string
		return ret
	}
	return *o.PlaySessionId
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoResponse) GetPlaySessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaySessionId) {
		return nil, false
	}
	return o.PlaySessionId, true
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoResponse) HasPlaySessionId() bool {
	if o != nil && !IsNil(o.PlaySessionId) {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given string and assigns it to the PlaySessionId field.
func (o *MediaInfoPlaybackInfoResponse) SetPlaySessionId(v string) {
	o.PlaySessionId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *MediaInfoPlaybackInfoResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoPlaybackInfoResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *MediaInfoPlaybackInfoResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *MediaInfoPlaybackInfoResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o MediaInfoPlaybackInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaInfoPlaybackInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaSources) {
		toSerialize["MediaSources"] = o.MediaSources
	}
	if !IsNil(o.PlaySessionId) {
		toSerialize["PlaySessionId"] = o.PlaySessionId
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableMediaInfoPlaybackInfoResponse struct {
	value *MediaInfoPlaybackInfoResponse
	isSet bool
}

func (v NullableMediaInfoPlaybackInfoResponse) Get() *MediaInfoPlaybackInfoResponse {
	return v.value
}

func (v *NullableMediaInfoPlaybackInfoResponse) Set(val *MediaInfoPlaybackInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoPlaybackInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoPlaybackInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoPlaybackInfoResponse(val *MediaInfoPlaybackInfoResponse) *NullableMediaInfoPlaybackInfoResponse {
	return &NullableMediaInfoPlaybackInfoResponse{value: val, isSet: true}
}

func (v NullableMediaInfoPlaybackInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoPlaybackInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the LiveTvSetChannelMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvSetChannelMapping{}

// LiveTvSetChannelMapping struct for LiveTvSetChannelMapping
type LiveTvSetChannelMapping struct {
	TunerChannelId *string `json:"TunerChannelId,omitempty"`
	ProviderChannelId *string `json:"ProviderChannelId,omitempty"`
}

// NewLiveTvSetChannelMapping instantiates a new LiveTvSetChannelMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvSetChannelMapping() *LiveTvSetChannelMapping {
	this := LiveTvSetChannelMapping{}
	return &this
}

// NewLiveTvSetChannelMappingWithDefaults instantiates a new LiveTvSetChannelMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvSetChannelMappingWithDefaults() *LiveTvSetChannelMapping {
	this := LiveTvSetChannelMapping{}
	return &this
}

// GetTunerChannelId returns the TunerChannelId field value if set, zero value otherwise.
func (o *LiveTvSetChannelMapping) GetTunerChannelId() string {
	if o == nil || IsNil(o.TunerChannelId) {
		var ret string
		return ret
	}
	return *o.TunerChannelId
}

// GetTunerChannelIdOk returns a tuple with the TunerChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvSetChannelMapping) GetTunerChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.TunerChannelId) {
		return nil, false
	}
	return o.TunerChannelId, true
}

// HasTunerChannelId returns a boolean if a field has been set.
func (o *LiveTvSetChannelMapping) HasTunerChannelId() bool {
	if o != nil && !IsNil(o.TunerChannelId) {
		return true
	}

	return false
}

// SetTunerChannelId gets a reference to the given string and assigns it to the TunerChannelId field.
func (o *LiveTvSetChannelMapping) SetTunerChannelId(v string) {
	o.TunerChannelId = &v
}

// GetProviderChannelId returns the ProviderChannelId field value if set, zero value otherwise.
func (o *LiveTvSetChannelMapping) GetProviderChannelId() string {
	if o == nil || IsNil(o.ProviderChannelId) {
		var ret string
		return ret
	}
	return *o.ProviderChannelId
}

// GetProviderChannelIdOk returns a tuple with the ProviderChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvSetChannelMapping) GetProviderChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderChannelId) {
		return nil, false
	}
	return o.ProviderChannelId, true
}

// HasProviderChannelId returns a boolean if a field has been set.
func (o *LiveTvSetChannelMapping) HasProviderChannelId() bool {
	if o != nil && !IsNil(o.ProviderChannelId) {
		return true
	}

	return false
}

// SetProviderChannelId gets a reference to the given string and assigns it to the ProviderChannelId field.
func (o *LiveTvSetChannelMapping) SetProviderChannelId(v string) {
	o.ProviderChannelId = &v
}

func (o LiveTvSetChannelMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvSetChannelMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TunerChannelId) {
		toSerialize["TunerChannelId"] = o.TunerChannelId
	}
	if !IsNil(o.ProviderChannelId) {
		toSerialize["ProviderChannelId"] = o.ProviderChannelId
	}
	return toSerialize, nil
}

type NullableLiveTvSetChannelMapping struct {
	value *LiveTvSetChannelMapping
	isSet bool
}

func (v NullableLiveTvSetChannelMapping) Get() *LiveTvSetChannelMapping {
	return v.value
}

func (v *NullableLiveTvSetChannelMapping) Set(val *LiveTvSetChannelMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvSetChannelMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvSetChannelMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvSetChannelMapping(val *LiveTvSetChannelMapping) *NullableLiveTvSetChannelMapping {
	return &NullableLiveTvSetChannelMapping{value: val, isSet: true}
}

func (v NullableLiveTvSetChannelMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvSetChannelMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



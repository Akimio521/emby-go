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

// checks if the PlaylistsPlaylistCreationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaylistsPlaylistCreationResult{}

// PlaylistsPlaylistCreationResult struct for PlaylistsPlaylistCreationResult
type PlaylistsPlaylistCreationResult struct {
	Id *string `json:"Id,omitempty"`
}

// NewPlaylistsPlaylistCreationResult instantiates a new PlaylistsPlaylistCreationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistsPlaylistCreationResult() *PlaylistsPlaylistCreationResult {
	this := PlaylistsPlaylistCreationResult{}
	return &this
}

// NewPlaylistsPlaylistCreationResultWithDefaults instantiates a new PlaylistsPlaylistCreationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistsPlaylistCreationResultWithDefaults() *PlaylistsPlaylistCreationResult {
	this := PlaylistsPlaylistCreationResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlaylistsPlaylistCreationResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistsPlaylistCreationResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlaylistsPlaylistCreationResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlaylistsPlaylistCreationResult) SetId(v string) {
	o.Id = &v
}

func (o PlaylistsPlaylistCreationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaylistsPlaylistCreationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePlaylistsPlaylistCreationResult struct {
	value *PlaylistsPlaylistCreationResult
	isSet bool
}

func (v NullablePlaylistsPlaylistCreationResult) Get() *PlaylistsPlaylistCreationResult {
	return v.value
}

func (v *NullablePlaylistsPlaylistCreationResult) Set(val *PlaylistsPlaylistCreationResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistsPlaylistCreationResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistsPlaylistCreationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistsPlaylistCreationResult(val *PlaylistsPlaylistCreationResult) *NullablePlaylistsPlaylistCreationResult {
	return &NullablePlaylistsPlaylistCreationResult{value: val, isSet: true}
}

func (v NullablePlaylistsPlaylistCreationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistsPlaylistCreationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



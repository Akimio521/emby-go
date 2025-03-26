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

// checks if the ChapterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChapterInfo{}

// ChapterInfo struct for ChapterInfo
type ChapterInfo struct {
	StartPositionTicks *int64 `json:"StartPositionTicks,omitempty"`
	Name *string `json:"Name,omitempty"`
	ImageTag *string `json:"ImageTag,omitempty"`
}

// NewChapterInfo instantiates a new ChapterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChapterInfo() *ChapterInfo {
	this := ChapterInfo{}
	return &this
}

// NewChapterInfoWithDefaults instantiates a new ChapterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChapterInfoWithDefaults() *ChapterInfo {
	this := ChapterInfo{}
	return &this
}

// GetStartPositionTicks returns the StartPositionTicks field value if set, zero value otherwise.
func (o *ChapterInfo) GetStartPositionTicks() int64 {
	if o == nil || IsNil(o.StartPositionTicks) {
		var ret int64
		return ret
	}
	return *o.StartPositionTicks
}

// GetStartPositionTicksOk returns a tuple with the StartPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterInfo) GetStartPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPositionTicks) {
		return nil, false
	}
	return o.StartPositionTicks, true
}

// HasStartPositionTicks returns a boolean if a field has been set.
func (o *ChapterInfo) HasStartPositionTicks() bool {
	if o != nil && !IsNil(o.StartPositionTicks) {
		return true
	}

	return false
}

// SetStartPositionTicks gets a reference to the given int64 and assigns it to the StartPositionTicks field.
func (o *ChapterInfo) SetStartPositionTicks(v int64) {
	o.StartPositionTicks = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChapterInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChapterInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChapterInfo) SetName(v string) {
	o.Name = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *ChapterInfo) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag) {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterInfo) GetImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTag) {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *ChapterInfo) HasImageTag() bool {
	if o != nil && !IsNil(o.ImageTag) {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *ChapterInfo) SetImageTag(v string) {
	o.ImageTag = &v
}

func (o ChapterInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChapterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartPositionTicks) {
		toSerialize["StartPositionTicks"] = o.StartPositionTicks
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ImageTag) {
		toSerialize["ImageTag"] = o.ImageTag
	}
	return toSerialize, nil
}

type NullableChapterInfo struct {
	value *ChapterInfo
	isSet bool
}

func (v NullableChapterInfo) Get() *ChapterInfo {
	return v.value
}

func (v *NullableChapterInfo) Set(val *ChapterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChapterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChapterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChapterInfo(val *ChapterInfo) *NullableChapterInfo {
	return &NullableChapterInfo{value: val, isSet: true}
}

func (v NullableChapterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChapterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



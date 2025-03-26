/*
Emby Server API

Explore the Emby Server API

API version: 4.1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the LiveTvGuideInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvGuideInfo{}

// LiveTvGuideInfo struct for LiveTvGuideInfo
type LiveTvGuideInfo struct {
	StartDate *time.Time `json:"StartDate,omitempty"`
	EndDate *time.Time `json:"EndDate,omitempty"`
}

// NewLiveTvGuideInfo instantiates a new LiveTvGuideInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvGuideInfo() *LiveTvGuideInfo {
	this := LiveTvGuideInfo{}
	return &this
}

// NewLiveTvGuideInfoWithDefaults instantiates a new LiveTvGuideInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvGuideInfoWithDefaults() *LiveTvGuideInfo {
	this := LiveTvGuideInfo{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LiveTvGuideInfo) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvGuideInfo) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LiveTvGuideInfo) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *LiveTvGuideInfo) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *LiveTvGuideInfo) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvGuideInfo) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *LiveTvGuideInfo) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *LiveTvGuideInfo) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o LiveTvGuideInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvGuideInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableLiveTvGuideInfo struct {
	value *LiveTvGuideInfo
	isSet bool
}

func (v NullableLiveTvGuideInfo) Get() *LiveTvGuideInfo {
	return v.value
}

func (v *NullableLiveTvGuideInfo) Set(val *LiveTvGuideInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvGuideInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvGuideInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvGuideInfo(val *LiveTvGuideInfo) *NullableLiveTvGuideInfo {
	return &NullableLiveTvGuideInfo{value: val, isSet: true}
}

func (v NullableLiveTvGuideInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvGuideInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



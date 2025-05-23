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

// checks if the DlnaHttpHeaderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DlnaHttpHeaderInfo{}

// DlnaHttpHeaderInfo struct for DlnaHttpHeaderInfo
type DlnaHttpHeaderInfo struct {
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
	Match *string `json:"Match,omitempty"`
}

// NewDlnaHttpHeaderInfo instantiates a new DlnaHttpHeaderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDlnaHttpHeaderInfo() *DlnaHttpHeaderInfo {
	this := DlnaHttpHeaderInfo{}
	return &this
}

// NewDlnaHttpHeaderInfoWithDefaults instantiates a new DlnaHttpHeaderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDlnaHttpHeaderInfoWithDefaults() *DlnaHttpHeaderInfo {
	this := DlnaHttpHeaderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DlnaHttpHeaderInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaHttpHeaderInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DlnaHttpHeaderInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DlnaHttpHeaderInfo) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DlnaHttpHeaderInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaHttpHeaderInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DlnaHttpHeaderInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DlnaHttpHeaderInfo) SetValue(v string) {
	o.Value = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *DlnaHttpHeaderInfo) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaHttpHeaderInfo) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *DlnaHttpHeaderInfo) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *DlnaHttpHeaderInfo) SetMatch(v string) {
	o.Match = &v
}

func (o DlnaHttpHeaderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DlnaHttpHeaderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if !IsNil(o.Match) {
		toSerialize["Match"] = o.Match
	}
	return toSerialize, nil
}

type NullableDlnaHttpHeaderInfo struct {
	value *DlnaHttpHeaderInfo
	isSet bool
}

func (v NullableDlnaHttpHeaderInfo) Get() *DlnaHttpHeaderInfo {
	return v.value
}

func (v *NullableDlnaHttpHeaderInfo) Set(val *DlnaHttpHeaderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDlnaHttpHeaderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDlnaHttpHeaderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlnaHttpHeaderInfo(val *DlnaHttpHeaderInfo) *NullableDlnaHttpHeaderInfo {
	return &NullableDlnaHttpHeaderInfo{value: val, isSet: true}
}

func (v NullableDlnaHttpHeaderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlnaHttpHeaderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



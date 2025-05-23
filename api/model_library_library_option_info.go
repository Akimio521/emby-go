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

// checks if the LibraryLibraryOptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryLibraryOptionInfo{}

// LibraryLibraryOptionInfo struct for LibraryLibraryOptionInfo
type LibraryLibraryOptionInfo struct {
	Name *string `json:"Name,omitempty"`
	DefaultEnabled *bool `json:"DefaultEnabled,omitempty"`
}

// NewLibraryLibraryOptionInfo instantiates a new LibraryLibraryOptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryLibraryOptionInfo() *LibraryLibraryOptionInfo {
	this := LibraryLibraryOptionInfo{}
	return &this
}

// NewLibraryLibraryOptionInfoWithDefaults instantiates a new LibraryLibraryOptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryLibraryOptionInfoWithDefaults() *LibraryLibraryOptionInfo {
	this := LibraryLibraryOptionInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LibraryLibraryOptionInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryLibraryOptionInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LibraryLibraryOptionInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LibraryLibraryOptionInfo) SetName(v string) {
	o.Name = &v
}

// GetDefaultEnabled returns the DefaultEnabled field value if set, zero value otherwise.
func (o *LibraryLibraryOptionInfo) GetDefaultEnabled() bool {
	if o == nil || IsNil(o.DefaultEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultEnabled
}

// GetDefaultEnabledOk returns a tuple with the DefaultEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryLibraryOptionInfo) GetDefaultEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultEnabled) {
		return nil, false
	}
	return o.DefaultEnabled, true
}

// HasDefaultEnabled returns a boolean if a field has been set.
func (o *LibraryLibraryOptionInfo) HasDefaultEnabled() bool {
	if o != nil && !IsNil(o.DefaultEnabled) {
		return true
	}

	return false
}

// SetDefaultEnabled gets a reference to the given bool and assigns it to the DefaultEnabled field.
func (o *LibraryLibraryOptionInfo) SetDefaultEnabled(v bool) {
	o.DefaultEnabled = &v
}

func (o LibraryLibraryOptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryLibraryOptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.DefaultEnabled) {
		toSerialize["DefaultEnabled"] = o.DefaultEnabled
	}
	return toSerialize, nil
}

type NullableLibraryLibraryOptionInfo struct {
	value *LibraryLibraryOptionInfo
	isSet bool
}

func (v NullableLibraryLibraryOptionInfo) Get() *LibraryLibraryOptionInfo {
	return v.value
}

func (v *NullableLibraryLibraryOptionInfo) Set(val *LibraryLibraryOptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryLibraryOptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryLibraryOptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryLibraryOptionInfo(val *LibraryLibraryOptionInfo) *NullableLibraryLibraryOptionInfo {
	return &NullableLibraryLibraryOptionInfo{value: val, isSet: true}
}

func (v NullableLibraryLibraryOptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryLibraryOptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



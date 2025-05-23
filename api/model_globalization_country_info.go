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

// checks if the GlobalizationCountryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalizationCountryInfo{}

// GlobalizationCountryInfo struct for GlobalizationCountryInfo
type GlobalizationCountryInfo struct {
	Name *string `json:"Name,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty"`
	TwoLetterISORegionName *string `json:"TwoLetterISORegionName,omitempty"`
	ThreeLetterISORegionName *string `json:"ThreeLetterISORegionName,omitempty"`
}

// NewGlobalizationCountryInfo instantiates a new GlobalizationCountryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalizationCountryInfo() *GlobalizationCountryInfo {
	this := GlobalizationCountryInfo{}
	return &this
}

// NewGlobalizationCountryInfoWithDefaults instantiates a new GlobalizationCountryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalizationCountryInfoWithDefaults() *GlobalizationCountryInfo {
	this := GlobalizationCountryInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalizationCountryInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalizationCountryInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalizationCountryInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalizationCountryInfo) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GlobalizationCountryInfo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalizationCountryInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GlobalizationCountryInfo) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GlobalizationCountryInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTwoLetterISORegionName returns the TwoLetterISORegionName field value if set, zero value otherwise.
func (o *GlobalizationCountryInfo) GetTwoLetterISORegionName() string {
	if o == nil || IsNil(o.TwoLetterISORegionName) {
		var ret string
		return ret
	}
	return *o.TwoLetterISORegionName
}

// GetTwoLetterISORegionNameOk returns a tuple with the TwoLetterISORegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalizationCountryInfo) GetTwoLetterISORegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.TwoLetterISORegionName) {
		return nil, false
	}
	return o.TwoLetterISORegionName, true
}

// HasTwoLetterISORegionName returns a boolean if a field has been set.
func (o *GlobalizationCountryInfo) HasTwoLetterISORegionName() bool {
	if o != nil && !IsNil(o.TwoLetterISORegionName) {
		return true
	}

	return false
}

// SetTwoLetterISORegionName gets a reference to the given string and assigns it to the TwoLetterISORegionName field.
func (o *GlobalizationCountryInfo) SetTwoLetterISORegionName(v string) {
	o.TwoLetterISORegionName = &v
}

// GetThreeLetterISORegionName returns the ThreeLetterISORegionName field value if set, zero value otherwise.
func (o *GlobalizationCountryInfo) GetThreeLetterISORegionName() string {
	if o == nil || IsNil(o.ThreeLetterISORegionName) {
		var ret string
		return ret
	}
	return *o.ThreeLetterISORegionName
}

// GetThreeLetterISORegionNameOk returns a tuple with the ThreeLetterISORegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalizationCountryInfo) GetThreeLetterISORegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ThreeLetterISORegionName) {
		return nil, false
	}
	return o.ThreeLetterISORegionName, true
}

// HasThreeLetterISORegionName returns a boolean if a field has been set.
func (o *GlobalizationCountryInfo) HasThreeLetterISORegionName() bool {
	if o != nil && !IsNil(o.ThreeLetterISORegionName) {
		return true
	}

	return false
}

// SetThreeLetterISORegionName gets a reference to the given string and assigns it to the ThreeLetterISORegionName field.
func (o *GlobalizationCountryInfo) SetThreeLetterISORegionName(v string) {
	o.ThreeLetterISORegionName = &v
}

func (o GlobalizationCountryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalizationCountryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !IsNil(o.TwoLetterISORegionName) {
		toSerialize["TwoLetterISORegionName"] = o.TwoLetterISORegionName
	}
	if !IsNil(o.ThreeLetterISORegionName) {
		toSerialize["ThreeLetterISORegionName"] = o.ThreeLetterISORegionName
	}
	return toSerialize, nil
}

type NullableGlobalizationCountryInfo struct {
	value *GlobalizationCountryInfo
	isSet bool
}

func (v NullableGlobalizationCountryInfo) Get() *GlobalizationCountryInfo {
	return v.value
}

func (v *NullableGlobalizationCountryInfo) Set(val *GlobalizationCountryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalizationCountryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalizationCountryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalizationCountryInfo(val *GlobalizationCountryInfo) *NullableGlobalizationCountryInfo {
	return &NullableGlobalizationCountryInfo{value: val, isSet: true}
}

func (v NullableGlobalizationCountryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalizationCountryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the WakeOnLanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WakeOnLanInfo{}

// WakeOnLanInfo struct for WakeOnLanInfo
type WakeOnLanInfo struct {
	MacAddress *string `json:"MacAddress,omitempty"`
	BroadcastAddress *string `json:"BroadcastAddress,omitempty"`
	Port *int32 `json:"Port,omitempty"`
}

// NewWakeOnLanInfo instantiates a new WakeOnLanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWakeOnLanInfo() *WakeOnLanInfo {
	this := WakeOnLanInfo{}
	return &this
}

// NewWakeOnLanInfoWithDefaults instantiates a new WakeOnLanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWakeOnLanInfoWithDefaults() *WakeOnLanInfo {
	this := WakeOnLanInfo{}
	return &this
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *WakeOnLanInfo) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WakeOnLanInfo) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WakeOnLanInfo) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *WakeOnLanInfo) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetBroadcastAddress returns the BroadcastAddress field value if set, zero value otherwise.
func (o *WakeOnLanInfo) GetBroadcastAddress() string {
	if o == nil || IsNil(o.BroadcastAddress) {
		var ret string
		return ret
	}
	return *o.BroadcastAddress
}

// GetBroadcastAddressOk returns a tuple with the BroadcastAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WakeOnLanInfo) GetBroadcastAddressOk() (*string, bool) {
	if o == nil || IsNil(o.BroadcastAddress) {
		return nil, false
	}
	return o.BroadcastAddress, true
}

// HasBroadcastAddress returns a boolean if a field has been set.
func (o *WakeOnLanInfo) HasBroadcastAddress() bool {
	if o != nil && !IsNil(o.BroadcastAddress) {
		return true
	}

	return false
}

// SetBroadcastAddress gets a reference to the given string and assigns it to the BroadcastAddress field.
func (o *WakeOnLanInfo) SetBroadcastAddress(v string) {
	o.BroadcastAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *WakeOnLanInfo) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WakeOnLanInfo) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *WakeOnLanInfo) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *WakeOnLanInfo) SetPort(v int32) {
	o.Port = &v
}

func (o WakeOnLanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WakeOnLanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MacAddress) {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if !IsNil(o.BroadcastAddress) {
		toSerialize["BroadcastAddress"] = o.BroadcastAddress
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	return toSerialize, nil
}

type NullableWakeOnLanInfo struct {
	value *WakeOnLanInfo
	isSet bool
}

func (v NullableWakeOnLanInfo) Get() *WakeOnLanInfo {
	return v.value
}

func (v *NullableWakeOnLanInfo) Set(val *WakeOnLanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWakeOnLanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWakeOnLanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWakeOnLanInfo(val *WakeOnLanInfo) *NullableWakeOnLanInfo {
	return &NullableWakeOnLanInfo{value: val, isSet: true}
}

func (v NullableWakeOnLanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWakeOnLanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



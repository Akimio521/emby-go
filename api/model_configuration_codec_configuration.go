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

// checks if the ConfigurationCodecConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationCodecConfiguration{}

// ConfigurationCodecConfiguration struct for ConfigurationCodecConfiguration
type ConfigurationCodecConfiguration struct {
	IsEnabled *bool `json:"IsEnabled,omitempty"`
	Priority *int32 `json:"Priority,omitempty"`
	CodecId *string `json:"CodecId,omitempty"`
}

// NewConfigurationCodecConfiguration instantiates a new ConfigurationCodecConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationCodecConfiguration() *ConfigurationCodecConfiguration {
	this := ConfigurationCodecConfiguration{}
	return &this
}

// NewConfigurationCodecConfigurationWithDefaults instantiates a new ConfigurationCodecConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationCodecConfigurationWithDefaults() *ConfigurationCodecConfiguration {
	this := ConfigurationCodecConfiguration{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ConfigurationCodecConfiguration) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationCodecConfiguration) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ConfigurationCodecConfiguration) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ConfigurationCodecConfiguration) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ConfigurationCodecConfiguration) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationCodecConfiguration) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ConfigurationCodecConfiguration) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ConfigurationCodecConfiguration) SetPriority(v int32) {
	o.Priority = &v
}

// GetCodecId returns the CodecId field value if set, zero value otherwise.
func (o *ConfigurationCodecConfiguration) GetCodecId() string {
	if o == nil || IsNil(o.CodecId) {
		var ret string
		return ret
	}
	return *o.CodecId
}

// GetCodecIdOk returns a tuple with the CodecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationCodecConfiguration) GetCodecIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodecId) {
		return nil, false
	}
	return o.CodecId, true
}

// HasCodecId returns a boolean if a field has been set.
func (o *ConfigurationCodecConfiguration) HasCodecId() bool {
	if o != nil && !IsNil(o.CodecId) {
		return true
	}

	return false
}

// SetCodecId gets a reference to the given string and assigns it to the CodecId field.
func (o *ConfigurationCodecConfiguration) SetCodecId(v string) {
	o.CodecId = &v
}

func (o ConfigurationCodecConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationCodecConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEnabled) {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.CodecId) {
		toSerialize["CodecId"] = o.CodecId
	}
	return toSerialize, nil
}

type NullableConfigurationCodecConfiguration struct {
	value *ConfigurationCodecConfiguration
	isSet bool
}

func (v NullableConfigurationCodecConfiguration) Get() *ConfigurationCodecConfiguration {
	return v.value
}

func (v *NullableConfigurationCodecConfiguration) Set(val *ConfigurationCodecConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationCodecConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationCodecConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationCodecConfiguration(val *ConfigurationCodecConfiguration) *NullableConfigurationCodecConfiguration {
	return &NullableConfigurationCodecConfiguration{value: val, isSet: true}
}

func (v NullableConfigurationCodecConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationCodecConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



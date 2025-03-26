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

// checks if the NotificationsNotificationTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsNotificationTypeInfo{}

// NotificationsNotificationTypeInfo struct for NotificationsNotificationTypeInfo
type NotificationsNotificationTypeInfo struct {
	Type *string `json:"Type,omitempty"`
	Name *string `json:"Name,omitempty"`
	Enabled *bool `json:"Enabled,omitempty"`
	Category *string `json:"Category,omitempty"`
	IsBasedOnUserEvent *bool `json:"IsBasedOnUserEvent,omitempty"`
}

// NewNotificationsNotificationTypeInfo instantiates a new NotificationsNotificationTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsNotificationTypeInfo() *NotificationsNotificationTypeInfo {
	this := NotificationsNotificationTypeInfo{}
	return &this
}

// NewNotificationsNotificationTypeInfoWithDefaults instantiates a new NotificationsNotificationTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsNotificationTypeInfoWithDefaults() *NotificationsNotificationTypeInfo {
	this := NotificationsNotificationTypeInfo{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationsNotificationTypeInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsNotificationTypeInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationsNotificationTypeInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationsNotificationTypeInfo) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationsNotificationTypeInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsNotificationTypeInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationsNotificationTypeInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationsNotificationTypeInfo) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationsNotificationTypeInfo) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsNotificationTypeInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationsNotificationTypeInfo) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationsNotificationTypeInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *NotificationsNotificationTypeInfo) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsNotificationTypeInfo) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *NotificationsNotificationTypeInfo) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *NotificationsNotificationTypeInfo) SetCategory(v string) {
	o.Category = &v
}

// GetIsBasedOnUserEvent returns the IsBasedOnUserEvent field value if set, zero value otherwise.
func (o *NotificationsNotificationTypeInfo) GetIsBasedOnUserEvent() bool {
	if o == nil || IsNil(o.IsBasedOnUserEvent) {
		var ret bool
		return ret
	}
	return *o.IsBasedOnUserEvent
}

// GetIsBasedOnUserEventOk returns a tuple with the IsBasedOnUserEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsNotificationTypeInfo) GetIsBasedOnUserEventOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBasedOnUserEvent) {
		return nil, false
	}
	return o.IsBasedOnUserEvent, true
}

// HasIsBasedOnUserEvent returns a boolean if a field has been set.
func (o *NotificationsNotificationTypeInfo) HasIsBasedOnUserEvent() bool {
	if o != nil && !IsNil(o.IsBasedOnUserEvent) {
		return true
	}

	return false
}

// SetIsBasedOnUserEvent gets a reference to the given bool and assigns it to the IsBasedOnUserEvent field.
func (o *NotificationsNotificationTypeInfo) SetIsBasedOnUserEvent(v bool) {
	o.IsBasedOnUserEvent = &v
}

func (o NotificationsNotificationTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsNotificationTypeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.IsBasedOnUserEvent) {
		toSerialize["IsBasedOnUserEvent"] = o.IsBasedOnUserEvent
	}
	return toSerialize, nil
}

type NullableNotificationsNotificationTypeInfo struct {
	value *NotificationsNotificationTypeInfo
	isSet bool
}

func (v NullableNotificationsNotificationTypeInfo) Get() *NotificationsNotificationTypeInfo {
	return v.value
}

func (v *NullableNotificationsNotificationTypeInfo) Set(val *NotificationsNotificationTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsNotificationTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsNotificationTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsNotificationTypeInfo(val *NotificationsNotificationTypeInfo) *NullableNotificationsNotificationTypeInfo {
	return &NullableNotificationsNotificationTypeInfo{value: val, isSet: true}
}

func (v NullableNotificationsNotificationTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsNotificationTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



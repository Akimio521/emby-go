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

// checks if the SyncModelSyncedItemProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncModelSyncedItemProgress{}

// SyncModelSyncedItemProgress struct for SyncModelSyncedItemProgress
type SyncModelSyncedItemProgress struct {
	Progress NullableFloat64 `json:"Progress,omitempty"`
	Status *string `json:"Status,omitempty"`
}

// NewSyncModelSyncedItemProgress instantiates a new SyncModelSyncedItemProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncModelSyncedItemProgress() *SyncModelSyncedItemProgress {
	this := SyncModelSyncedItemProgress{}
	return &this
}

// NewSyncModelSyncedItemProgressWithDefaults instantiates a new SyncModelSyncedItemProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncModelSyncedItemProgressWithDefaults() *SyncModelSyncedItemProgress {
	this := SyncModelSyncedItemProgress{}
	return &this
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncModelSyncedItemProgress) GetProgress() float64 {
	if o == nil || IsNil(o.Progress.Get()) {
		var ret float64
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncModelSyncedItemProgress) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *SyncModelSyncedItemProgress) HasProgress() bool {
	if o != nil && o.Progress.IsSet() {
		return true
	}

	return false
}

// SetProgress gets a reference to the given NullableFloat64 and assigns it to the Progress field.
func (o *SyncModelSyncedItemProgress) SetProgress(v float64) {
	o.Progress.Set(&v)
}
// SetProgressNil sets the value for Progress to be an explicit nil
func (o *SyncModelSyncedItemProgress) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil
func (o *SyncModelSyncedItemProgress) UnsetProgress() {
	o.Progress.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyncModelSyncedItemProgress) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncModelSyncedItemProgress) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyncModelSyncedItemProgress) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyncModelSyncedItemProgress) SetStatus(v string) {
	o.Status = &v
}

func (o SyncModelSyncedItemProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncModelSyncedItemProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Progress.IsSet() {
		toSerialize["Progress"] = o.Progress.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSyncModelSyncedItemProgress struct {
	value *SyncModelSyncedItemProgress
	isSet bool
}

func (v NullableSyncModelSyncedItemProgress) Get() *SyncModelSyncedItemProgress {
	return v.value
}

func (v *NullableSyncModelSyncedItemProgress) Set(val *SyncModelSyncedItemProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncModelSyncedItemProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncModelSyncedItemProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncModelSyncedItemProgress(val *SyncModelSyncedItemProgress) *NullableSyncModelSyncedItemProgress {
	return &NullableSyncModelSyncedItemProgress{value: val, isSet: true}
}

func (v NullableSyncModelSyncedItemProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncModelSyncedItemProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



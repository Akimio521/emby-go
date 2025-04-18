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

// checks if the UpdatesInstallationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatesInstallationInfo{}

// UpdatesInstallationInfo struct for UpdatesInstallationInfo
type UpdatesInstallationInfo struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	AssemblyGuid *string `json:"AssemblyGuid,omitempty"`
	Version *string `json:"Version,omitempty"`
	UpdateClass *string `json:"UpdateClass,omitempty"`
	PercentComplete NullableFloat64 `json:"PercentComplete,omitempty"`
}

// NewUpdatesInstallationInfo instantiates a new UpdatesInstallationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesInstallationInfo() *UpdatesInstallationInfo {
	this := UpdatesInstallationInfo{}
	return &this
}

// NewUpdatesInstallationInfoWithDefaults instantiates a new UpdatesInstallationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesInstallationInfoWithDefaults() *UpdatesInstallationInfo {
	this := UpdatesInstallationInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdatesInstallationInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesInstallationInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdatesInstallationInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatesInstallationInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesInstallationInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatesInstallationInfo) SetName(v string) {
	o.Name = &v
}

// GetAssemblyGuid returns the AssemblyGuid field value if set, zero value otherwise.
func (o *UpdatesInstallationInfo) GetAssemblyGuid() string {
	if o == nil || IsNil(o.AssemblyGuid) {
		var ret string
		return ret
	}
	return *o.AssemblyGuid
}

// GetAssemblyGuidOk returns a tuple with the AssemblyGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesInstallationInfo) GetAssemblyGuidOk() (*string, bool) {
	if o == nil || IsNil(o.AssemblyGuid) {
		return nil, false
	}
	return o.AssemblyGuid, true
}

// HasAssemblyGuid returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasAssemblyGuid() bool {
	if o != nil && !IsNil(o.AssemblyGuid) {
		return true
	}

	return false
}

// SetAssemblyGuid gets a reference to the given string and assigns it to the AssemblyGuid field.
func (o *UpdatesInstallationInfo) SetAssemblyGuid(v string) {
	o.AssemblyGuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpdatesInstallationInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesInstallationInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpdatesInstallationInfo) SetVersion(v string) {
	o.Version = &v
}

// GetUpdateClass returns the UpdateClass field value if set, zero value otherwise.
func (o *UpdatesInstallationInfo) GetUpdateClass() string {
	if o == nil || IsNil(o.UpdateClass) {
		var ret string
		return ret
	}
	return *o.UpdateClass
}

// GetUpdateClassOk returns a tuple with the UpdateClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesInstallationInfo) GetUpdateClassOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateClass) {
		return nil, false
	}
	return o.UpdateClass, true
}

// HasUpdateClass returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasUpdateClass() bool {
	if o != nil && !IsNil(o.UpdateClass) {
		return true
	}

	return false
}

// SetUpdateClass gets a reference to the given string and assigns it to the UpdateClass field.
func (o *UpdatesInstallationInfo) SetUpdateClass(v string) {
	o.UpdateClass = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatesInstallationInfo) GetPercentComplete() float64 {
	if o == nil || IsNil(o.PercentComplete.Get()) {
		var ret float64
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatesInstallationInfo) GetPercentCompleteOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *UpdatesInstallationInfo) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableFloat64 and assigns it to the PercentComplete field.
func (o *UpdatesInstallationInfo) SetPercentComplete(v float64) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *UpdatesInstallationInfo) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *UpdatesInstallationInfo) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

func (o UpdatesInstallationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatesInstallationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.AssemblyGuid) {
		toSerialize["AssemblyGuid"] = o.AssemblyGuid
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.UpdateClass) {
		toSerialize["UpdateClass"] = o.UpdateClass
	}
	if o.PercentComplete.IsSet() {
		toSerialize["PercentComplete"] = o.PercentComplete.Get()
	}
	return toSerialize, nil
}

type NullableUpdatesInstallationInfo struct {
	value *UpdatesInstallationInfo
	isSet bool
}

func (v NullableUpdatesInstallationInfo) Get() *UpdatesInstallationInfo {
	return v.value
}

func (v *NullableUpdatesInstallationInfo) Set(val *UpdatesInstallationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesInstallationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesInstallationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesInstallationInfo(val *UpdatesInstallationInfo) *NullableUpdatesInstallationInfo {
	return &NullableUpdatesInstallationInfo{value: val, isSet: true}
}

func (v NullableUpdatesInstallationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesInstallationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



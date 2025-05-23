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

// checks if the IOFileSystemEntryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IOFileSystemEntryInfo{}

// IOFileSystemEntryInfo struct for IOFileSystemEntryInfo
type IOFileSystemEntryInfo struct {
	Name *string `json:"Name,omitempty"`
	Path *string `json:"Path,omitempty"`
	Type *string `json:"Type,omitempty"`
}

// NewIOFileSystemEntryInfo instantiates a new IOFileSystemEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIOFileSystemEntryInfo() *IOFileSystemEntryInfo {
	this := IOFileSystemEntryInfo{}
	return &this
}

// NewIOFileSystemEntryInfoWithDefaults instantiates a new IOFileSystemEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIOFileSystemEntryInfoWithDefaults() *IOFileSystemEntryInfo {
	this := IOFileSystemEntryInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IOFileSystemEntryInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IOFileSystemEntryInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IOFileSystemEntryInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IOFileSystemEntryInfo) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *IOFileSystemEntryInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IOFileSystemEntryInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *IOFileSystemEntryInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *IOFileSystemEntryInfo) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IOFileSystemEntryInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IOFileSystemEntryInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IOFileSystemEntryInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IOFileSystemEntryInfo) SetType(v string) {
	o.Type = &v
}

func (o IOFileSystemEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IOFileSystemEntryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIOFileSystemEntryInfo struct {
	value *IOFileSystemEntryInfo
	isSet bool
}

func (v NullableIOFileSystemEntryInfo) Get() *IOFileSystemEntryInfo {
	return v.value
}

func (v *NullableIOFileSystemEntryInfo) Set(val *IOFileSystemEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIOFileSystemEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIOFileSystemEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIOFileSystemEntryInfo(val *IOFileSystemEntryInfo) *NullableIOFileSystemEntryInfo {
	return &NullableIOFileSystemEntryInfo{value: val, isSet: true}
}

func (v NullableIOFileSystemEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIOFileSystemEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



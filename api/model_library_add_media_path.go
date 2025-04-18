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

// checks if the LibraryAddMediaPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryAddMediaPath{}

// LibraryAddMediaPath struct for LibraryAddMediaPath
type LibraryAddMediaPath struct {
	Name *string `json:"Name,omitempty"`
	Path *string `json:"Path,omitempty"`
	PathInfo *ConfigurationMediaPathInfo `json:"PathInfo,omitempty"`
	RefreshLibrary *bool `json:"RefreshLibrary,omitempty"`
}

// NewLibraryAddMediaPath instantiates a new LibraryAddMediaPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryAddMediaPath() *LibraryAddMediaPath {
	this := LibraryAddMediaPath{}
	return &this
}

// NewLibraryAddMediaPathWithDefaults instantiates a new LibraryAddMediaPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryAddMediaPathWithDefaults() *LibraryAddMediaPath {
	this := LibraryAddMediaPath{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LibraryAddMediaPath) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryAddMediaPath) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LibraryAddMediaPath) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LibraryAddMediaPath) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LibraryAddMediaPath) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryAddMediaPath) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LibraryAddMediaPath) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LibraryAddMediaPath) SetPath(v string) {
	o.Path = &v
}

// GetPathInfo returns the PathInfo field value if set, zero value otherwise.
func (o *LibraryAddMediaPath) GetPathInfo() ConfigurationMediaPathInfo {
	if o == nil || IsNil(o.PathInfo) {
		var ret ConfigurationMediaPathInfo
		return ret
	}
	return *o.PathInfo
}

// GetPathInfoOk returns a tuple with the PathInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryAddMediaPath) GetPathInfoOk() (*ConfigurationMediaPathInfo, bool) {
	if o == nil || IsNil(o.PathInfo) {
		return nil, false
	}
	return o.PathInfo, true
}

// HasPathInfo returns a boolean if a field has been set.
func (o *LibraryAddMediaPath) HasPathInfo() bool {
	if o != nil && !IsNil(o.PathInfo) {
		return true
	}

	return false
}

// SetPathInfo gets a reference to the given ConfigurationMediaPathInfo and assigns it to the PathInfo field.
func (o *LibraryAddMediaPath) SetPathInfo(v ConfigurationMediaPathInfo) {
	o.PathInfo = &v
}

// GetRefreshLibrary returns the RefreshLibrary field value if set, zero value otherwise.
func (o *LibraryAddMediaPath) GetRefreshLibrary() bool {
	if o == nil || IsNil(o.RefreshLibrary) {
		var ret bool
		return ret
	}
	return *o.RefreshLibrary
}

// GetRefreshLibraryOk returns a tuple with the RefreshLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryAddMediaPath) GetRefreshLibraryOk() (*bool, bool) {
	if o == nil || IsNil(o.RefreshLibrary) {
		return nil, false
	}
	return o.RefreshLibrary, true
}

// HasRefreshLibrary returns a boolean if a field has been set.
func (o *LibraryAddMediaPath) HasRefreshLibrary() bool {
	if o != nil && !IsNil(o.RefreshLibrary) {
		return true
	}

	return false
}

// SetRefreshLibrary gets a reference to the given bool and assigns it to the RefreshLibrary field.
func (o *LibraryAddMediaPath) SetRefreshLibrary(v bool) {
	o.RefreshLibrary = &v
}

func (o LibraryAddMediaPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryAddMediaPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.PathInfo) {
		toSerialize["PathInfo"] = o.PathInfo
	}
	if !IsNil(o.RefreshLibrary) {
		toSerialize["RefreshLibrary"] = o.RefreshLibrary
	}
	return toSerialize, nil
}

type NullableLibraryAddMediaPath struct {
	value *LibraryAddMediaPath
	isSet bool
}

func (v NullableLibraryAddMediaPath) Get() *LibraryAddMediaPath {
	return v.value
}

func (v *NullableLibraryAddMediaPath) Set(val *LibraryAddMediaPath) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryAddMediaPath) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryAddMediaPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryAddMediaPath(val *LibraryAddMediaPath) *NullableLibraryAddMediaPath {
	return &NullableLibraryAddMediaPath{value: val, isSet: true}
}

func (v NullableLibraryAddMediaPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryAddMediaPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the LibraryMediaFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryMediaFolder{}

// LibraryMediaFolder struct for LibraryMediaFolder
type LibraryMediaFolder struct {
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	SubFolders []LibrarySubFolder `json:"SubFolders,omitempty"`
}

// NewLibraryMediaFolder instantiates a new LibraryMediaFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryMediaFolder() *LibraryMediaFolder {
	this := LibraryMediaFolder{}
	return &this
}

// NewLibraryMediaFolderWithDefaults instantiates a new LibraryMediaFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryMediaFolderWithDefaults() *LibraryMediaFolder {
	this := LibraryMediaFolder{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LibraryMediaFolder) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryMediaFolder) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LibraryMediaFolder) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LibraryMediaFolder) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LibraryMediaFolder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryMediaFolder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LibraryMediaFolder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LibraryMediaFolder) SetId(v string) {
	o.Id = &v
}

// GetSubFolders returns the SubFolders field value if set, zero value otherwise.
func (o *LibraryMediaFolder) GetSubFolders() []LibrarySubFolder {
	if o == nil || IsNil(o.SubFolders) {
		var ret []LibrarySubFolder
		return ret
	}
	return o.SubFolders
}

// GetSubFoldersOk returns a tuple with the SubFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryMediaFolder) GetSubFoldersOk() ([]LibrarySubFolder, bool) {
	if o == nil || IsNil(o.SubFolders) {
		return nil, false
	}
	return o.SubFolders, true
}

// HasSubFolders returns a boolean if a field has been set.
func (o *LibraryMediaFolder) HasSubFolders() bool {
	if o != nil && !IsNil(o.SubFolders) {
		return true
	}

	return false
}

// SetSubFolders gets a reference to the given []LibrarySubFolder and assigns it to the SubFolders field.
func (o *LibraryMediaFolder) SetSubFolders(v []LibrarySubFolder) {
	o.SubFolders = v
}

func (o LibraryMediaFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryMediaFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.SubFolders) {
		toSerialize["SubFolders"] = o.SubFolders
	}
	return toSerialize, nil
}

type NullableLibraryMediaFolder struct {
	value *LibraryMediaFolder
	isSet bool
}

func (v NullableLibraryMediaFolder) Get() *LibraryMediaFolder {
	return v.value
}

func (v *NullableLibraryMediaFolder) Set(val *LibraryMediaFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryMediaFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryMediaFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryMediaFolder(val *LibraryMediaFolder) *NullableLibraryMediaFolder {
	return &NullableLibraryMediaFolder{value: val, isSet: true}
}

func (v NullableLibraryMediaFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryMediaFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



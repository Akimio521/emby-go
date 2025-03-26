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

// checks if the UserLibraryOfficialRatingItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLibraryOfficialRatingItem{}

// UserLibraryOfficialRatingItem struct for UserLibraryOfficialRatingItem
type UserLibraryOfficialRatingItem struct {
	Name *string `json:"Name,omitempty"`
}

// NewUserLibraryOfficialRatingItem instantiates a new UserLibraryOfficialRatingItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLibraryOfficialRatingItem() *UserLibraryOfficialRatingItem {
	this := UserLibraryOfficialRatingItem{}
	return &this
}

// NewUserLibraryOfficialRatingItemWithDefaults instantiates a new UserLibraryOfficialRatingItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLibraryOfficialRatingItemWithDefaults() *UserLibraryOfficialRatingItem {
	this := UserLibraryOfficialRatingItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserLibraryOfficialRatingItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLibraryOfficialRatingItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserLibraryOfficialRatingItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserLibraryOfficialRatingItem) SetName(v string) {
	o.Name = &v
}

func (o UserLibraryOfficialRatingItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLibraryOfficialRatingItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUserLibraryOfficialRatingItem struct {
	value *UserLibraryOfficialRatingItem
	isSet bool
}

func (v NullableUserLibraryOfficialRatingItem) Get() *UserLibraryOfficialRatingItem {
	return v.value
}

func (v *NullableUserLibraryOfficialRatingItem) Set(val *UserLibraryOfficialRatingItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLibraryOfficialRatingItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLibraryOfficialRatingItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLibraryOfficialRatingItem(val *UserLibraryOfficialRatingItem) *NullableUserLibraryOfficialRatingItem {
	return &NullableUserLibraryOfficialRatingItem{value: val, isSet: true}
}

func (v NullableUserLibraryOfficialRatingItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLibraryOfficialRatingItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



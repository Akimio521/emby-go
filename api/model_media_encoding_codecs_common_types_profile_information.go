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

// checks if the MediaEncodingCodecsCommonTypesProfileInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaEncodingCodecsCommonTypesProfileInformation{}

// MediaEncodingCodecsCommonTypesProfileInformation struct for MediaEncodingCodecsCommonTypesProfileInformation
type MediaEncodingCodecsCommonTypesProfileInformation struct {
	ShortName *string `json:"ShortName,omitempty"`
	Description *string `json:"Description,omitempty"`
	Details *string `json:"Details,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewMediaEncodingCodecsCommonTypesProfileInformation instantiates a new MediaEncodingCodecsCommonTypesProfileInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaEncodingCodecsCommonTypesProfileInformation() *MediaEncodingCodecsCommonTypesProfileInformation {
	this := MediaEncodingCodecsCommonTypesProfileInformation{}
	return &this
}

// NewMediaEncodingCodecsCommonTypesProfileInformationWithDefaults instantiates a new MediaEncodingCodecsCommonTypesProfileInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaEncodingCodecsCommonTypesProfileInformationWithDefaults() *MediaEncodingCodecsCommonTypesProfileInformation {
	this := MediaEncodingCodecsCommonTypesProfileInformation{}
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) SetShortName(v string) {
	o.ShortName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) SetDetails(v string) {
	o.Details = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MediaEncodingCodecsCommonTypesProfileInformation) SetId(v string) {
	o.Id = &v
}

func (o MediaEncodingCodecsCommonTypesProfileInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaEncodingCodecsCommonTypesProfileInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShortName) {
		toSerialize["ShortName"] = o.ShortName
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["Details"] = o.Details
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullableMediaEncodingCodecsCommonTypesProfileInformation struct {
	value *MediaEncodingCodecsCommonTypesProfileInformation
	isSet bool
}

func (v NullableMediaEncodingCodecsCommonTypesProfileInformation) Get() *MediaEncodingCodecsCommonTypesProfileInformation {
	return v.value
}

func (v *NullableMediaEncodingCodecsCommonTypesProfileInformation) Set(val *MediaEncodingCodecsCommonTypesProfileInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaEncodingCodecsCommonTypesProfileInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaEncodingCodecsCommonTypesProfileInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaEncodingCodecsCommonTypesProfileInformation(val *MediaEncodingCodecsCommonTypesProfileInformation) *NullableMediaEncodingCodecsCommonTypesProfileInformation {
	return &NullableMediaEncodingCodecsCommonTypesProfileInformation{value: val, isSet: true}
}

func (v NullableMediaEncodingCodecsCommonTypesProfileInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaEncodingCodecsCommonTypesProfileInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



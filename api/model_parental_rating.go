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

// checks if the ParentalRating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentalRating{}

// ParentalRating struct for ParentalRating
type ParentalRating struct {
	Name *string `json:"Name,omitempty"`
	Value *int32 `json:"Value,omitempty"`
}

// NewParentalRating instantiates a new ParentalRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentalRating() *ParentalRating {
	this := ParentalRating{}
	return &this
}

// NewParentalRatingWithDefaults instantiates a new ParentalRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentalRatingWithDefaults() *ParentalRating {
	this := ParentalRating{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParentalRating) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalRating) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParentalRating) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParentalRating) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ParentalRating) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalRating) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ParentalRating) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ParentalRating) SetValue(v int32) {
	o.Value = &v
}

func (o ParentalRating) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentalRating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableParentalRating struct {
	value *ParentalRating
	isSet bool
}

func (v NullableParentalRating) Get() *ParentalRating {
	return v.value
}

func (v *NullableParentalRating) Set(val *ParentalRating) {
	v.value = val
	v.isSet = true
}

func (v NullableParentalRating) IsSet() bool {
	return v.isSet
}

func (v *NullableParentalRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentalRating(val *ParentalRating) *NullableParentalRating {
	return &NullableParentalRating{value: val, isSet: true}
}

func (v NullableParentalRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentalRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



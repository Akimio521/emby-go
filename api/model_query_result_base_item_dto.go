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

// checks if the QueryResultBaseItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryResultBaseItemDto{}

// QueryResultBaseItemDto struct for QueryResultBaseItemDto
type QueryResultBaseItemDto struct {
	Items []BaseItemDto `json:"Items,omitempty"`
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
}

// NewQueryResultBaseItemDto instantiates a new QueryResultBaseItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryResultBaseItemDto() *QueryResultBaseItemDto {
	this := QueryResultBaseItemDto{}
	return &this
}

// NewQueryResultBaseItemDtoWithDefaults instantiates a new QueryResultBaseItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryResultBaseItemDtoWithDefaults() *QueryResultBaseItemDto {
	this := QueryResultBaseItemDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QueryResultBaseItemDto) GetItems() []BaseItemDto {
	if o == nil || IsNil(o.Items) {
		var ret []BaseItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResultBaseItemDto) GetItemsOk() ([]BaseItemDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QueryResultBaseItemDto) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BaseItemDto and assigns it to the Items field.
func (o *QueryResultBaseItemDto) SetItems(v []BaseItemDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *QueryResultBaseItemDto) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResultBaseItemDto) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *QueryResultBaseItemDto) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *QueryResultBaseItemDto) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

func (o QueryResultBaseItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryResultBaseItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	return toSerialize, nil
}

type NullableQueryResultBaseItemDto struct {
	value *QueryResultBaseItemDto
	isSet bool
}

func (v NullableQueryResultBaseItemDto) Get() *QueryResultBaseItemDto {
	return v.value
}

func (v *NullableQueryResultBaseItemDto) Set(val *QueryResultBaseItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryResultBaseItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryResultBaseItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryResultBaseItemDto(val *QueryResultBaseItemDto) *NullableQueryResultBaseItemDto {
	return &NullableQueryResultBaseItemDto{value: val, isSet: true}
}

func (v NullableQueryResultBaseItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryResultBaseItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



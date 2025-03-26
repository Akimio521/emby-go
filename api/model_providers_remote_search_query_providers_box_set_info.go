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

// checks if the ProvidersRemoteSearchQueryProvidersBoxSetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvidersRemoteSearchQueryProvidersBoxSetInfo{}

// ProvidersRemoteSearchQueryProvidersBoxSetInfo struct for ProvidersRemoteSearchQueryProvidersBoxSetInfo
type ProvidersRemoteSearchQueryProvidersBoxSetInfo struct {
	SearchInfo *ProvidersBoxSetInfo `json:"SearchInfo,omitempty"`
	ItemId *int64 `json:"ItemId,omitempty"`
	SearchProviderName *string `json:"SearchProviderName,omitempty"`
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewProvidersRemoteSearchQueryProvidersBoxSetInfo instantiates a new ProvidersRemoteSearchQueryProvidersBoxSetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvidersRemoteSearchQueryProvidersBoxSetInfo() *ProvidersRemoteSearchQueryProvidersBoxSetInfo {
	this := ProvidersRemoteSearchQueryProvidersBoxSetInfo{}
	return &this
}

// NewProvidersRemoteSearchQueryProvidersBoxSetInfoWithDefaults instantiates a new ProvidersRemoteSearchQueryProvidersBoxSetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvidersRemoteSearchQueryProvidersBoxSetInfoWithDefaults() *ProvidersRemoteSearchQueryProvidersBoxSetInfo {
	this := ProvidersRemoteSearchQueryProvidersBoxSetInfo{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetSearchInfo() ProvidersBoxSetInfo {
	if o == nil || IsNil(o.SearchInfo) {
		var ret ProvidersBoxSetInfo
		return ret
	}
	return *o.SearchInfo
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetSearchInfoOk() (*ProvidersBoxSetInfo, bool) {
	if o == nil || IsNil(o.SearchInfo) {
		return nil, false
	}
	return o.SearchInfo, true
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) HasSearchInfo() bool {
	if o != nil && !IsNil(o.SearchInfo) {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given ProvidersBoxSetInfo and assigns it to the SearchInfo field.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) SetSearchInfo(v ProvidersBoxSetInfo) {
	o.SearchInfo = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) SetItemId(v int64) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName) {
		var ret string
		return ret
	}
	return *o.SearchProviderName
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetSearchProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SearchProviderName) {
		return nil, false
	}
	return o.SearchProviderName, true
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) HasSearchProviderName() bool {
	if o != nil && !IsNil(o.SearchProviderName) {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given string and assigns it to the SearchProviderName field.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) SetSearchProviderName(v string) {
	o.SearchProviderName = &v
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *ProvidersRemoteSearchQueryProvidersBoxSetInfo) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o ProvidersRemoteSearchQueryProvidersBoxSetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvidersRemoteSearchQueryProvidersBoxSetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchInfo) {
		toSerialize["SearchInfo"] = o.SearchInfo
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.SearchProviderName) {
		toSerialize["SearchProviderName"] = o.SearchProviderName
	}
	if !IsNil(o.IncludeDisabledProviders) {
		toSerialize["IncludeDisabledProviders"] = o.IncludeDisabledProviders
	}
	return toSerialize, nil
}

type NullableProvidersRemoteSearchQueryProvidersBoxSetInfo struct {
	value *ProvidersRemoteSearchQueryProvidersBoxSetInfo
	isSet bool
}

func (v NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) Get() *ProvidersRemoteSearchQueryProvidersBoxSetInfo {
	return v.value
}

func (v *NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) Set(val *ProvidersRemoteSearchQueryProvidersBoxSetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvidersRemoteSearchQueryProvidersBoxSetInfo(val *ProvidersRemoteSearchQueryProvidersBoxSetInfo) *NullableProvidersRemoteSearchQueryProvidersBoxSetInfo {
	return &NullableProvidersRemoteSearchQueryProvidersBoxSetInfo{value: val, isSet: true}
}

func (v NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvidersRemoteSearchQueryProvidersBoxSetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



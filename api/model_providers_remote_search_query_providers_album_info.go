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

// checks if the ProvidersRemoteSearchQueryProvidersAlbumInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvidersRemoteSearchQueryProvidersAlbumInfo{}

// ProvidersRemoteSearchQueryProvidersAlbumInfo struct for ProvidersRemoteSearchQueryProvidersAlbumInfo
type ProvidersRemoteSearchQueryProvidersAlbumInfo struct {
	SearchInfo *ProvidersAlbumInfo `json:"SearchInfo,omitempty"`
	ItemId *int64 `json:"ItemId,omitempty"`
	SearchProviderName *string `json:"SearchProviderName,omitempty"`
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewProvidersRemoteSearchQueryProvidersAlbumInfo instantiates a new ProvidersRemoteSearchQueryProvidersAlbumInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvidersRemoteSearchQueryProvidersAlbumInfo() *ProvidersRemoteSearchQueryProvidersAlbumInfo {
	this := ProvidersRemoteSearchQueryProvidersAlbumInfo{}
	return &this
}

// NewProvidersRemoteSearchQueryProvidersAlbumInfoWithDefaults instantiates a new ProvidersRemoteSearchQueryProvidersAlbumInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvidersRemoteSearchQueryProvidersAlbumInfoWithDefaults() *ProvidersRemoteSearchQueryProvidersAlbumInfo {
	this := ProvidersRemoteSearchQueryProvidersAlbumInfo{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetSearchInfo() ProvidersAlbumInfo {
	if o == nil || IsNil(o.SearchInfo) {
		var ret ProvidersAlbumInfo
		return ret
	}
	return *o.SearchInfo
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetSearchInfoOk() (*ProvidersAlbumInfo, bool) {
	if o == nil || IsNil(o.SearchInfo) {
		return nil, false
	}
	return o.SearchInfo, true
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) HasSearchInfo() bool {
	if o != nil && !IsNil(o.SearchInfo) {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given ProvidersAlbumInfo and assigns it to the SearchInfo field.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) SetSearchInfo(v ProvidersAlbumInfo) {
	o.SearchInfo = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) SetItemId(v int64) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName) {
		var ret string
		return ret
	}
	return *o.SearchProviderName
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetSearchProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SearchProviderName) {
		return nil, false
	}
	return o.SearchProviderName, true
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) HasSearchProviderName() bool {
	if o != nil && !IsNil(o.SearchProviderName) {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given string and assigns it to the SearchProviderName field.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) SetSearchProviderName(v string) {
	o.SearchProviderName = &v
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *ProvidersRemoteSearchQueryProvidersAlbumInfo) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o ProvidersRemoteSearchQueryProvidersAlbumInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvidersRemoteSearchQueryProvidersAlbumInfo) ToMap() (map[string]interface{}, error) {
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

type NullableProvidersRemoteSearchQueryProvidersAlbumInfo struct {
	value *ProvidersRemoteSearchQueryProvidersAlbumInfo
	isSet bool
}

func (v NullableProvidersRemoteSearchQueryProvidersAlbumInfo) Get() *ProvidersRemoteSearchQueryProvidersAlbumInfo {
	return v.value
}

func (v *NullableProvidersRemoteSearchQueryProvidersAlbumInfo) Set(val *ProvidersRemoteSearchQueryProvidersAlbumInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProvidersRemoteSearchQueryProvidersAlbumInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProvidersRemoteSearchQueryProvidersAlbumInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvidersRemoteSearchQueryProvidersAlbumInfo(val *ProvidersRemoteSearchQueryProvidersAlbumInfo) *NullableProvidersRemoteSearchQueryProvidersAlbumInfo {
	return &NullableProvidersRemoteSearchQueryProvidersAlbumInfo{value: val, isSet: true}
}

func (v NullableProvidersRemoteSearchQueryProvidersAlbumInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvidersRemoteSearchQueryProvidersAlbumInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the VirtualFolderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualFolderInfo{}

// VirtualFolderInfo struct for VirtualFolderInfo
type VirtualFolderInfo struct {
	Name *string `json:"Name,omitempty"`
	Locations []string `json:"Locations,omitempty"`
	CollectionType *string `json:"CollectionType,omitempty"`
	LibraryOptions *ConfigurationLibraryOptions `json:"LibraryOptions,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	PrimaryImageItemId *string `json:"PrimaryImageItemId,omitempty"`
	RefreshProgress NullableFloat64 `json:"RefreshProgress,omitempty"`
	RefreshStatus *string `json:"RefreshStatus,omitempty"`
}

// NewVirtualFolderInfo instantiates a new VirtualFolderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualFolderInfo() *VirtualFolderInfo {
	this := VirtualFolderInfo{}
	return &this
}

// NewVirtualFolderInfoWithDefaults instantiates a new VirtualFolderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualFolderInfoWithDefaults() *VirtualFolderInfo {
	this := VirtualFolderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualFolderInfo) SetName(v string) {
	o.Name = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetLocations() []string {
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetLocationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *VirtualFolderInfo) SetLocations(v []string) {
	o.Locations = v
}

// GetCollectionType returns the CollectionType field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetCollectionType() string {
	if o == nil || IsNil(o.CollectionType) {
		var ret string
		return ret
	}
	return *o.CollectionType
}

// GetCollectionTypeOk returns a tuple with the CollectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetCollectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionType) {
		return nil, false
	}
	return o.CollectionType, true
}

// HasCollectionType returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasCollectionType() bool {
	if o != nil && !IsNil(o.CollectionType) {
		return true
	}

	return false
}

// SetCollectionType gets a reference to the given string and assigns it to the CollectionType field.
func (o *VirtualFolderInfo) SetCollectionType(v string) {
	o.CollectionType = &v
}

// GetLibraryOptions returns the LibraryOptions field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetLibraryOptions() ConfigurationLibraryOptions {
	if o == nil || IsNil(o.LibraryOptions) {
		var ret ConfigurationLibraryOptions
		return ret
	}
	return *o.LibraryOptions
}

// GetLibraryOptionsOk returns a tuple with the LibraryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetLibraryOptionsOk() (*ConfigurationLibraryOptions, bool) {
	if o == nil || IsNil(o.LibraryOptions) {
		return nil, false
	}
	return o.LibraryOptions, true
}

// HasLibraryOptions returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasLibraryOptions() bool {
	if o != nil && !IsNil(o.LibraryOptions) {
		return true
	}

	return false
}

// SetLibraryOptions gets a reference to the given ConfigurationLibraryOptions and assigns it to the LibraryOptions field.
func (o *VirtualFolderInfo) SetLibraryOptions(v ConfigurationLibraryOptions) {
	o.LibraryOptions = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *VirtualFolderInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetPrimaryImageItemId returns the PrimaryImageItemId field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetPrimaryImageItemId() string {
	if o == nil || IsNil(o.PrimaryImageItemId) {
		var ret string
		return ret
	}
	return *o.PrimaryImageItemId
}

// GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetPrimaryImageItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryImageItemId) {
		return nil, false
	}
	return o.PrimaryImageItemId, true
}

// HasPrimaryImageItemId returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasPrimaryImageItemId() bool {
	if o != nil && !IsNil(o.PrimaryImageItemId) {
		return true
	}

	return false
}

// SetPrimaryImageItemId gets a reference to the given string and assigns it to the PrimaryImageItemId field.
func (o *VirtualFolderInfo) SetPrimaryImageItemId(v string) {
	o.PrimaryImageItemId = &v
}

// GetRefreshProgress returns the RefreshProgress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualFolderInfo) GetRefreshProgress() float64 {
	if o == nil || IsNil(o.RefreshProgress.Get()) {
		var ret float64
		return ret
	}
	return *o.RefreshProgress.Get()
}

// GetRefreshProgressOk returns a tuple with the RefreshProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualFolderInfo) GetRefreshProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshProgress.Get(), o.RefreshProgress.IsSet()
}

// HasRefreshProgress returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasRefreshProgress() bool {
	if o != nil && o.RefreshProgress.IsSet() {
		return true
	}

	return false
}

// SetRefreshProgress gets a reference to the given NullableFloat64 and assigns it to the RefreshProgress field.
func (o *VirtualFolderInfo) SetRefreshProgress(v float64) {
	o.RefreshProgress.Set(&v)
}
// SetRefreshProgressNil sets the value for RefreshProgress to be an explicit nil
func (o *VirtualFolderInfo) SetRefreshProgressNil() {
	o.RefreshProgress.Set(nil)
}

// UnsetRefreshProgress ensures that no value is present for RefreshProgress, not even an explicit nil
func (o *VirtualFolderInfo) UnsetRefreshProgress() {
	o.RefreshProgress.Unset()
}

// GetRefreshStatus returns the RefreshStatus field value if set, zero value otherwise.
func (o *VirtualFolderInfo) GetRefreshStatus() string {
	if o == nil || IsNil(o.RefreshStatus) {
		var ret string
		return ret
	}
	return *o.RefreshStatus
}

// GetRefreshStatusOk returns a tuple with the RefreshStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolderInfo) GetRefreshStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshStatus) {
		return nil, false
	}
	return o.RefreshStatus, true
}

// HasRefreshStatus returns a boolean if a field has been set.
func (o *VirtualFolderInfo) HasRefreshStatus() bool {
	if o != nil && !IsNil(o.RefreshStatus) {
		return true
	}

	return false
}

// SetRefreshStatus gets a reference to the given string and assigns it to the RefreshStatus field.
func (o *VirtualFolderInfo) SetRefreshStatus(v string) {
	o.RefreshStatus = &v
}

func (o VirtualFolderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualFolderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Locations) {
		toSerialize["Locations"] = o.Locations
	}
	if !IsNil(o.CollectionType) {
		toSerialize["CollectionType"] = o.CollectionType
	}
	if !IsNil(o.LibraryOptions) {
		toSerialize["LibraryOptions"] = o.LibraryOptions
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.PrimaryImageItemId) {
		toSerialize["PrimaryImageItemId"] = o.PrimaryImageItemId
	}
	if o.RefreshProgress.IsSet() {
		toSerialize["RefreshProgress"] = o.RefreshProgress.Get()
	}
	if !IsNil(o.RefreshStatus) {
		toSerialize["RefreshStatus"] = o.RefreshStatus
	}
	return toSerialize, nil
}

type NullableVirtualFolderInfo struct {
	value *VirtualFolderInfo
	isSet bool
}

func (v NullableVirtualFolderInfo) Get() *VirtualFolderInfo {
	return v.value
}

func (v *NullableVirtualFolderInfo) Set(val *VirtualFolderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualFolderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualFolderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualFolderInfo(val *VirtualFolderInfo) *NullableVirtualFolderInfo {
	return &NullableVirtualFolderInfo{value: val, isSet: true}
}

func (v NullableVirtualFolderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualFolderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



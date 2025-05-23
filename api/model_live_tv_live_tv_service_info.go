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

// checks if the LiveTvLiveTvServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvLiveTvServiceInfo{}

// LiveTvLiveTvServiceInfo struct for LiveTvLiveTvServiceInfo
type LiveTvLiveTvServiceInfo struct {
	Name *string `json:"Name,omitempty"`
	HomePageUrl *string `json:"HomePageUrl,omitempty"`
	Status *string `json:"Status,omitempty"`
	StatusMessage *string `json:"StatusMessage,omitempty"`
	Version *string `json:"Version,omitempty"`
	HasUpdateAvailable *bool `json:"HasUpdateAvailable,omitempty"`
	IsVisible *bool `json:"IsVisible,omitempty"`
	Tuners []string `json:"Tuners,omitempty"`
}

// NewLiveTvLiveTvServiceInfo instantiates a new LiveTvLiveTvServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvLiveTvServiceInfo() *LiveTvLiveTvServiceInfo {
	this := LiveTvLiveTvServiceInfo{}
	return &this
}

// NewLiveTvLiveTvServiceInfoWithDefaults instantiates a new LiveTvLiveTvServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvLiveTvServiceInfoWithDefaults() *LiveTvLiveTvServiceInfo {
	this := LiveTvLiveTvServiceInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LiveTvLiveTvServiceInfo) SetName(v string) {
	o.Name = &v
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetHomePageUrl() string {
	if o == nil || IsNil(o.HomePageUrl) {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetHomePageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomePageUrl) {
		return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasHomePageUrl() bool {
	if o != nil && !IsNil(o.HomePageUrl) {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *LiveTvLiveTvServiceInfo) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiveTvLiveTvServiceInfo) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *LiveTvLiveTvServiceInfo) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LiveTvLiveTvServiceInfo) SetVersion(v string) {
	o.Version = &v
}

// GetHasUpdateAvailable returns the HasUpdateAvailable field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetHasUpdateAvailable() bool {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.HasUpdateAvailable
}

// GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetHasUpdateAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		return nil, false
	}
	return o.HasUpdateAvailable, true
}

// HasHasUpdateAvailable returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasHasUpdateAvailable() bool {
	if o != nil && !IsNil(o.HasUpdateAvailable) {
		return true
	}

	return false
}

// SetHasUpdateAvailable gets a reference to the given bool and assigns it to the HasUpdateAvailable field.
func (o *LiveTvLiveTvServiceInfo) SetHasUpdateAvailable(v bool) {
	o.HasUpdateAvailable = &v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetIsVisible() bool {
	if o == nil || IsNil(o.IsVisible) {
		var ret bool
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetIsVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVisible) {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasIsVisible() bool {
	if o != nil && !IsNil(o.IsVisible) {
		return true
	}

	return false
}

// SetIsVisible gets a reference to the given bool and assigns it to the IsVisible field.
func (o *LiveTvLiveTvServiceInfo) SetIsVisible(v bool) {
	o.IsVisible = &v
}

// GetTuners returns the Tuners field value if set, zero value otherwise.
func (o *LiveTvLiveTvServiceInfo) GetTuners() []string {
	if o == nil || IsNil(o.Tuners) {
		var ret []string
		return ret
	}
	return o.Tuners
}

// GetTunersOk returns a tuple with the Tuners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvLiveTvServiceInfo) GetTunersOk() ([]string, bool) {
	if o == nil || IsNil(o.Tuners) {
		return nil, false
	}
	return o.Tuners, true
}

// HasTuners returns a boolean if a field has been set.
func (o *LiveTvLiveTvServiceInfo) HasTuners() bool {
	if o != nil && !IsNil(o.Tuners) {
		return true
	}

	return false
}

// SetTuners gets a reference to the given []string and assigns it to the Tuners field.
func (o *LiveTvLiveTvServiceInfo) SetTuners(v []string) {
	o.Tuners = v
}

func (o LiveTvLiveTvServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvLiveTvServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.HomePageUrl) {
		toSerialize["HomePageUrl"] = o.HomePageUrl
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.HasUpdateAvailable) {
		toSerialize["HasUpdateAvailable"] = o.HasUpdateAvailable
	}
	if !IsNil(o.IsVisible) {
		toSerialize["IsVisible"] = o.IsVisible
	}
	if !IsNil(o.Tuners) {
		toSerialize["Tuners"] = o.Tuners
	}
	return toSerialize, nil
}

type NullableLiveTvLiveTvServiceInfo struct {
	value *LiveTvLiveTvServiceInfo
	isSet bool
}

func (v NullableLiveTvLiveTvServiceInfo) Get() *LiveTvLiveTvServiceInfo {
	return v.value
}

func (v *NullableLiveTvLiveTvServiceInfo) Set(val *LiveTvLiveTvServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvLiveTvServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvLiveTvServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvLiveTvServiceInfo(val *LiveTvLiveTvServiceInfo) *NullableLiveTvLiveTvServiceInfo {
	return &NullableLiveTvLiveTvServiceInfo{value: val, isSet: true}
}

func (v NullableLiveTvLiveTvServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvLiveTvServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



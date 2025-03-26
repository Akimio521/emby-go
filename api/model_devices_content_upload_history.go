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

// checks if the DevicesContentUploadHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesContentUploadHistory{}

// DevicesContentUploadHistory struct for DevicesContentUploadHistory
type DevicesContentUploadHistory struct {
	DeviceId *string `json:"DeviceId,omitempty"`
	FilesUploaded []DevicesLocalFileInfo `json:"FilesUploaded,omitempty"`
}

// NewDevicesContentUploadHistory instantiates a new DevicesContentUploadHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesContentUploadHistory() *DevicesContentUploadHistory {
	this := DevicesContentUploadHistory{}
	return &this
}

// NewDevicesContentUploadHistoryWithDefaults instantiates a new DevicesContentUploadHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesContentUploadHistoryWithDefaults() *DevicesContentUploadHistory {
	this := DevicesContentUploadHistory{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DevicesContentUploadHistory) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesContentUploadHistory) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DevicesContentUploadHistory) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DevicesContentUploadHistory) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetFilesUploaded returns the FilesUploaded field value if set, zero value otherwise.
func (o *DevicesContentUploadHistory) GetFilesUploaded() []DevicesLocalFileInfo {
	if o == nil || IsNil(o.FilesUploaded) {
		var ret []DevicesLocalFileInfo
		return ret
	}
	return o.FilesUploaded
}

// GetFilesUploadedOk returns a tuple with the FilesUploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesContentUploadHistory) GetFilesUploadedOk() ([]DevicesLocalFileInfo, bool) {
	if o == nil || IsNil(o.FilesUploaded) {
		return nil, false
	}
	return o.FilesUploaded, true
}

// HasFilesUploaded returns a boolean if a field has been set.
func (o *DevicesContentUploadHistory) HasFilesUploaded() bool {
	if o != nil && !IsNil(o.FilesUploaded) {
		return true
	}

	return false
}

// SetFilesUploaded gets a reference to the given []DevicesLocalFileInfo and assigns it to the FilesUploaded field.
func (o *DevicesContentUploadHistory) SetFilesUploaded(v []DevicesLocalFileInfo) {
	o.FilesUploaded = v
}

func (o DevicesContentUploadHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesContentUploadHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.FilesUploaded) {
		toSerialize["FilesUploaded"] = o.FilesUploaded
	}
	return toSerialize, nil
}

type NullableDevicesContentUploadHistory struct {
	value *DevicesContentUploadHistory
	isSet bool
}

func (v NullableDevicesContentUploadHistory) Get() *DevicesContentUploadHistory {
	return v.value
}

func (v *NullableDevicesContentUploadHistory) Set(val *DevicesContentUploadHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesContentUploadHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesContentUploadHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesContentUploadHistory(val *DevicesContentUploadHistory) *NullableDevicesContentUploadHistory {
	return &NullableDevicesContentUploadHistory{value: val, isSet: true}
}

func (v NullableDevicesContentUploadHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesContentUploadHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



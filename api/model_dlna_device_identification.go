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

// checks if the DlnaDeviceIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DlnaDeviceIdentification{}

// DlnaDeviceIdentification struct for DlnaDeviceIdentification
type DlnaDeviceIdentification struct {
	FriendlyName *string `json:"FriendlyName,omitempty"`
	ModelNumber *string `json:"ModelNumber,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty"`
	ModelName *string `json:"ModelName,omitempty"`
	ModelDescription *string `json:"ModelDescription,omitempty"`
	DeviceDescription *string `json:"DeviceDescription,omitempty"`
	ModelUrl *string `json:"ModelUrl,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty"`
	ManufacturerUrl *string `json:"ManufacturerUrl,omitempty"`
	Headers []DlnaHttpHeaderInfo `json:"Headers,omitempty"`
}

// NewDlnaDeviceIdentification instantiates a new DlnaDeviceIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDlnaDeviceIdentification() *DlnaDeviceIdentification {
	this := DlnaDeviceIdentification{}
	return &this
}

// NewDlnaDeviceIdentificationWithDefaults instantiates a new DlnaDeviceIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDlnaDeviceIdentificationWithDefaults() *DlnaDeviceIdentification {
	this := DlnaDeviceIdentification{}
	return &this
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *DlnaDeviceIdentification) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetModelNumber() string {
	if o == nil || IsNil(o.ModelNumber) {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetModelNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ModelNumber) {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasModelNumber() bool {
	if o != nil && !IsNil(o.ModelNumber) {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *DlnaDeviceIdentification) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DlnaDeviceIdentification) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *DlnaDeviceIdentification) SetModelName(v string) {
	o.ModelName = &v
}

// GetModelDescription returns the ModelDescription field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetModelDescription() string {
	if o == nil || IsNil(o.ModelDescription) {
		var ret string
		return ret
	}
	return *o.ModelDescription
}

// GetModelDescriptionOk returns a tuple with the ModelDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetModelDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ModelDescription) {
		return nil, false
	}
	return o.ModelDescription, true
}

// HasModelDescription returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasModelDescription() bool {
	if o != nil && !IsNil(o.ModelDescription) {
		return true
	}

	return false
}

// SetModelDescription gets a reference to the given string and assigns it to the ModelDescription field.
func (o *DlnaDeviceIdentification) SetModelDescription(v string) {
	o.ModelDescription = &v
}

// GetDeviceDescription returns the DeviceDescription field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetDeviceDescription() string {
	if o == nil || IsNil(o.DeviceDescription) {
		var ret string
		return ret
	}
	return *o.DeviceDescription
}

// GetDeviceDescriptionOk returns a tuple with the DeviceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetDeviceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceDescription) {
		return nil, false
	}
	return o.DeviceDescription, true
}

// HasDeviceDescription returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasDeviceDescription() bool {
	if o != nil && !IsNil(o.DeviceDescription) {
		return true
	}

	return false
}

// SetDeviceDescription gets a reference to the given string and assigns it to the DeviceDescription field.
func (o *DlnaDeviceIdentification) SetDeviceDescription(v string) {
	o.DeviceDescription = &v
}

// GetModelUrl returns the ModelUrl field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetModelUrl() string {
	if o == nil || IsNil(o.ModelUrl) {
		var ret string
		return ret
	}
	return *o.ModelUrl
}

// GetModelUrlOk returns a tuple with the ModelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetModelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ModelUrl) {
		return nil, false
	}
	return o.ModelUrl, true
}

// HasModelUrl returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasModelUrl() bool {
	if o != nil && !IsNil(o.ModelUrl) {
		return true
	}

	return false
}

// SetModelUrl gets a reference to the given string and assigns it to the ModelUrl field.
func (o *DlnaDeviceIdentification) SetModelUrl(v string) {
	o.ModelUrl = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *DlnaDeviceIdentification) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetManufacturerUrl returns the ManufacturerUrl field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetManufacturerUrl() string {
	if o == nil || IsNil(o.ManufacturerUrl) {
		var ret string
		return ret
	}
	return *o.ManufacturerUrl
}

// GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetManufacturerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ManufacturerUrl) {
		return nil, false
	}
	return o.ManufacturerUrl, true
}

// HasManufacturerUrl returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasManufacturerUrl() bool {
	if o != nil && !IsNil(o.ManufacturerUrl) {
		return true
	}

	return false
}

// SetManufacturerUrl gets a reference to the given string and assigns it to the ManufacturerUrl field.
func (o *DlnaDeviceIdentification) SetManufacturerUrl(v string) {
	o.ManufacturerUrl = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *DlnaDeviceIdentification) GetHeaders() []DlnaHttpHeaderInfo {
	if o == nil || IsNil(o.Headers) {
		var ret []DlnaHttpHeaderInfo
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlnaDeviceIdentification) GetHeadersOk() ([]DlnaHttpHeaderInfo, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *DlnaDeviceIdentification) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []DlnaHttpHeaderInfo and assigns it to the Headers field.
func (o *DlnaDeviceIdentification) SetHeaders(v []DlnaHttpHeaderInfo) {
	o.Headers = v
}

func (o DlnaDeviceIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DlnaDeviceIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FriendlyName) {
		toSerialize["FriendlyName"] = o.FriendlyName
	}
	if !IsNil(o.ModelNumber) {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if !IsNil(o.ModelName) {
		toSerialize["ModelName"] = o.ModelName
	}
	if !IsNil(o.ModelDescription) {
		toSerialize["ModelDescription"] = o.ModelDescription
	}
	if !IsNil(o.DeviceDescription) {
		toSerialize["DeviceDescription"] = o.DeviceDescription
	}
	if !IsNil(o.ModelUrl) {
		toSerialize["ModelUrl"] = o.ModelUrl
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["Manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.ManufacturerUrl) {
		toSerialize["ManufacturerUrl"] = o.ManufacturerUrl
	}
	if !IsNil(o.Headers) {
		toSerialize["Headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableDlnaDeviceIdentification struct {
	value *DlnaDeviceIdentification
	isSet bool
}

func (v NullableDlnaDeviceIdentification) Get() *DlnaDeviceIdentification {
	return v.value
}

func (v *NullableDlnaDeviceIdentification) Set(val *DlnaDeviceIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableDlnaDeviceIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableDlnaDeviceIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlnaDeviceIdentification(val *DlnaDeviceIdentification) *NullableDlnaDeviceIdentification {
	return &NullableDlnaDeviceIdentification{value: val, isSet: true}
}

func (v NullableDlnaDeviceIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlnaDeviceIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



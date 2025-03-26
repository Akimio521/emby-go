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

// checks if the MediaEncodingCodecsCommonTypesResolutionWithRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaEncodingCodecsCommonTypesResolutionWithRate{}

// MediaEncodingCodecsCommonTypesResolutionWithRate struct for MediaEncodingCodecsCommonTypesResolutionWithRate
type MediaEncodingCodecsCommonTypesResolutionWithRate struct {
	Width *int32 `json:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty"`
	FrameRate *float64 `json:"FrameRate,omitempty"`
	Resolution *MediaEncodingCodecsCommonTypesResolution `json:"Resolution,omitempty"`
}

// NewMediaEncodingCodecsCommonTypesResolutionWithRate instantiates a new MediaEncodingCodecsCommonTypesResolutionWithRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaEncodingCodecsCommonTypesResolutionWithRate() *MediaEncodingCodecsCommonTypesResolutionWithRate {
	this := MediaEncodingCodecsCommonTypesResolutionWithRate{}
	return &this
}

// NewMediaEncodingCodecsCommonTypesResolutionWithRateWithDefaults instantiates a new MediaEncodingCodecsCommonTypesResolutionWithRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaEncodingCodecsCommonTypesResolutionWithRateWithDefaults() *MediaEncodingCodecsCommonTypesResolutionWithRate {
	this := MediaEncodingCodecsCommonTypesResolutionWithRate{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) SetHeight(v int32) {
	o.Height = &v
}

// GetFrameRate returns the FrameRate field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetFrameRate() float64 {
	if o == nil || IsNil(o.FrameRate) {
		var ret float64
		return ret
	}
	return *o.FrameRate
}

// GetFrameRateOk returns a tuple with the FrameRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetFrameRateOk() (*float64, bool) {
	if o == nil || IsNil(o.FrameRate) {
		return nil, false
	}
	return o.FrameRate, true
}

// HasFrameRate returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) HasFrameRate() bool {
	if o != nil && !IsNil(o.FrameRate) {
		return true
	}

	return false
}

// SetFrameRate gets a reference to the given float64 and assigns it to the FrameRate field.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) SetFrameRate(v float64) {
	o.FrameRate = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetResolution() MediaEncodingCodecsCommonTypesResolution {
	if o == nil || IsNil(o.Resolution) {
		var ret MediaEncodingCodecsCommonTypesResolution
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) GetResolutionOk() (*MediaEncodingCodecsCommonTypesResolution, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given MediaEncodingCodecsCommonTypesResolution and assigns it to the Resolution field.
func (o *MediaEncodingCodecsCommonTypesResolutionWithRate) SetResolution(v MediaEncodingCodecsCommonTypesResolution) {
	o.Resolution = &v
}

func (o MediaEncodingCodecsCommonTypesResolutionWithRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaEncodingCodecsCommonTypesResolutionWithRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["Width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["Height"] = o.Height
	}
	if !IsNil(o.FrameRate) {
		toSerialize["FrameRate"] = o.FrameRate
	}
	if !IsNil(o.Resolution) {
		toSerialize["Resolution"] = o.Resolution
	}
	return toSerialize, nil
}

type NullableMediaEncodingCodecsCommonTypesResolutionWithRate struct {
	value *MediaEncodingCodecsCommonTypesResolutionWithRate
	isSet bool
}

func (v NullableMediaEncodingCodecsCommonTypesResolutionWithRate) Get() *MediaEncodingCodecsCommonTypesResolutionWithRate {
	return v.value
}

func (v *NullableMediaEncodingCodecsCommonTypesResolutionWithRate) Set(val *MediaEncodingCodecsCommonTypesResolutionWithRate) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaEncodingCodecsCommonTypesResolutionWithRate) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaEncodingCodecsCommonTypesResolutionWithRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaEncodingCodecsCommonTypesResolutionWithRate(val *MediaEncodingCodecsCommonTypesResolutionWithRate) *NullableMediaEncodingCodecsCommonTypesResolutionWithRate {
	return &NullableMediaEncodingCodecsCommonTypesResolutionWithRate{value: val, isSet: true}
}

func (v NullableMediaEncodingCodecsCommonTypesResolutionWithRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaEncodingCodecsCommonTypesResolutionWithRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the PluginsPluginInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginsPluginInfo{}

// PluginsPluginInfo struct for PluginsPluginInfo
type PluginsPluginInfo struct {
	Name *string `json:"Name,omitempty"`
	Version *string `json:"Version,omitempty"`
	ConfigurationFileName *string `json:"ConfigurationFileName,omitempty"`
	Description *string `json:"Description,omitempty"`
	Id *string `json:"Id,omitempty"`
	ImageTag *string `json:"ImageTag,omitempty"`
}

// NewPluginsPluginInfo instantiates a new PluginsPluginInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginsPluginInfo() *PluginsPluginInfo {
	this := PluginsPluginInfo{}
	return &this
}

// NewPluginsPluginInfoWithDefaults instantiates a new PluginsPluginInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsPluginInfoWithDefaults() *PluginsPluginInfo {
	this := PluginsPluginInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginsPluginInfo) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PluginsPluginInfo) SetVersion(v string) {
	o.Version = &v
}

// GetConfigurationFileName returns the ConfigurationFileName field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetConfigurationFileName() string {
	if o == nil || IsNil(o.ConfigurationFileName) {
		var ret string
		return ret
	}
	return *o.ConfigurationFileName
}

// GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetConfigurationFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationFileName) {
		return nil, false
	}
	return o.ConfigurationFileName, true
}

// HasConfigurationFileName returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasConfigurationFileName() bool {
	if o != nil && !IsNil(o.ConfigurationFileName) {
		return true
	}

	return false
}

// SetConfigurationFileName gets a reference to the given string and assigns it to the ConfigurationFileName field.
func (o *PluginsPluginInfo) SetConfigurationFileName(v string) {
	o.ConfigurationFileName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluginsPluginInfo) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PluginsPluginInfo) SetId(v string) {
	o.Id = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *PluginsPluginInfo) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag) {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginsPluginInfo) GetImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTag) {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *PluginsPluginInfo) HasImageTag() bool {
	if o != nil && !IsNil(o.ImageTag) {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *PluginsPluginInfo) SetImageTag(v string) {
	o.ImageTag = &v
}

func (o PluginsPluginInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginsPluginInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.ConfigurationFileName) {
		toSerialize["ConfigurationFileName"] = o.ConfigurationFileName
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.ImageTag) {
		toSerialize["ImageTag"] = o.ImageTag
	}
	return toSerialize, nil
}

type NullablePluginsPluginInfo struct {
	value *PluginsPluginInfo
	isSet bool
}

func (v NullablePluginsPluginInfo) Get() *PluginsPluginInfo {
	return v.value
}

func (v *NullablePluginsPluginInfo) Set(val *PluginsPluginInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginsPluginInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginsPluginInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginsPluginInfo(val *PluginsPluginInfo) *NullablePluginsPluginInfo {
	return &NullablePluginsPluginInfo{value: val, isSet: true}
}

func (v NullablePluginsPluginInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginsPluginInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



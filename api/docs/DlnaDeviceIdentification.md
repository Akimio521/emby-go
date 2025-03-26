# DlnaDeviceIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**ModelDescription** | Pointer to **string** |  | [optional] 
**DeviceDescription** | Pointer to **string** |  | [optional] 
**ModelUrl** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerUrl** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]DlnaHttpHeaderInfo**](DlnaHttpHeaderInfo.md) |  | [optional] 

## Methods

### NewDlnaDeviceIdentification

`func NewDlnaDeviceIdentification() *DlnaDeviceIdentification`

NewDlnaDeviceIdentification instantiates a new DlnaDeviceIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaDeviceIdentificationWithDefaults

`func NewDlnaDeviceIdentificationWithDefaults() *DlnaDeviceIdentification`

NewDlnaDeviceIdentificationWithDefaults instantiates a new DlnaDeviceIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *DlnaDeviceIdentification) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DlnaDeviceIdentification) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DlnaDeviceIdentification) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DlnaDeviceIdentification) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetModelNumber

`func (o *DlnaDeviceIdentification) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DlnaDeviceIdentification) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DlnaDeviceIdentification) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DlnaDeviceIdentification) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DlnaDeviceIdentification) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DlnaDeviceIdentification) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DlnaDeviceIdentification) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DlnaDeviceIdentification) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetModelName

`func (o *DlnaDeviceIdentification) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DlnaDeviceIdentification) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DlnaDeviceIdentification) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DlnaDeviceIdentification) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelDescription

`func (o *DlnaDeviceIdentification) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *DlnaDeviceIdentification) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *DlnaDeviceIdentification) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *DlnaDeviceIdentification) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### GetDeviceDescription

`func (o *DlnaDeviceIdentification) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *DlnaDeviceIdentification) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *DlnaDeviceIdentification) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *DlnaDeviceIdentification) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetModelUrl

`func (o *DlnaDeviceIdentification) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *DlnaDeviceIdentification) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *DlnaDeviceIdentification) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *DlnaDeviceIdentification) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### GetManufacturer

`func (o *DlnaDeviceIdentification) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DlnaDeviceIdentification) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DlnaDeviceIdentification) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DlnaDeviceIdentification) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerUrl

`func (o *DlnaDeviceIdentification) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *DlnaDeviceIdentification) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *DlnaDeviceIdentification) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *DlnaDeviceIdentification) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *DlnaDeviceIdentification) GetHeaders() []DlnaHttpHeaderInfo`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DlnaDeviceIdentification) GetHeadersOk() (*[]DlnaHttpHeaderInfo, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DlnaDeviceIdentification) SetHeaders(v []DlnaHttpHeaderInfo)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DlnaDeviceIdentification) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DevicesContentUploadHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**FilesUploaded** | Pointer to [**[]DevicesLocalFileInfo**](DevicesLocalFileInfo.md) |  | [optional] 

## Methods

### NewDevicesContentUploadHistory

`func NewDevicesContentUploadHistory() *DevicesContentUploadHistory`

NewDevicesContentUploadHistory instantiates a new DevicesContentUploadHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesContentUploadHistoryWithDefaults

`func NewDevicesContentUploadHistoryWithDefaults() *DevicesContentUploadHistory`

NewDevicesContentUploadHistoryWithDefaults instantiates a new DevicesContentUploadHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DevicesContentUploadHistory) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DevicesContentUploadHistory) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DevicesContentUploadHistory) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DevicesContentUploadHistory) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFilesUploaded

`func (o *DevicesContentUploadHistory) GetFilesUploaded() []DevicesLocalFileInfo`

GetFilesUploaded returns the FilesUploaded field if non-nil, zero value otherwise.

### GetFilesUploadedOk

`func (o *DevicesContentUploadHistory) GetFilesUploadedOk() (*[]DevicesLocalFileInfo, bool)`

GetFilesUploadedOk returns a tuple with the FilesUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesUploaded

`func (o *DevicesContentUploadHistory) SetFilesUploaded(v []DevicesLocalFileInfo)`

SetFilesUploaded sets FilesUploaded field to given value.

### HasFilesUploaded

`func (o *DevicesContentUploadHistory) HasFilesUploaded() bool`

HasFilesUploaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



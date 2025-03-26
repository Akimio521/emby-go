# UpdatesPackageVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**VersionStr** | Pointer to **string** |  | [optional] 
**Classification** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiredVersionStr** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**TargetFilename** | Pointer to **string** |  | [optional] 
**InfoUrl** | Pointer to **string** |  | [optional] 
**Runtimes** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatesPackageVersionInfo

`func NewUpdatesPackageVersionInfo() *UpdatesPackageVersionInfo`

NewUpdatesPackageVersionInfo instantiates a new UpdatesPackageVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesPackageVersionInfoWithDefaults

`func NewUpdatesPackageVersionInfoWithDefaults() *UpdatesPackageVersionInfo`

NewUpdatesPackageVersionInfoWithDefaults instantiates a new UpdatesPackageVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatesPackageVersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatesPackageVersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatesPackageVersionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatesPackageVersionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGuid

`func (o *UpdatesPackageVersionInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *UpdatesPackageVersionInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *UpdatesPackageVersionInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *UpdatesPackageVersionInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetVersionStr

`func (o *UpdatesPackageVersionInfo) GetVersionStr() string`

GetVersionStr returns the VersionStr field if non-nil, zero value otherwise.

### GetVersionStrOk

`func (o *UpdatesPackageVersionInfo) GetVersionStrOk() (*string, bool)`

GetVersionStrOk returns a tuple with the VersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStr

`func (o *UpdatesPackageVersionInfo) SetVersionStr(v string)`

SetVersionStr sets VersionStr field to given value.

### HasVersionStr

`func (o *UpdatesPackageVersionInfo) HasVersionStr() bool`

HasVersionStr returns a boolean if a field has been set.

### GetClassification

`func (o *UpdatesPackageVersionInfo) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *UpdatesPackageVersionInfo) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *UpdatesPackageVersionInfo) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *UpdatesPackageVersionInfo) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatesPackageVersionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatesPackageVersionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatesPackageVersionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatesPackageVersionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredVersionStr

`func (o *UpdatesPackageVersionInfo) GetRequiredVersionStr() string`

GetRequiredVersionStr returns the RequiredVersionStr field if non-nil, zero value otherwise.

### GetRequiredVersionStrOk

`func (o *UpdatesPackageVersionInfo) GetRequiredVersionStrOk() (*string, bool)`

GetRequiredVersionStrOk returns a tuple with the RequiredVersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredVersionStr

`func (o *UpdatesPackageVersionInfo) SetRequiredVersionStr(v string)`

SetRequiredVersionStr sets RequiredVersionStr field to given value.

### HasRequiredVersionStr

`func (o *UpdatesPackageVersionInfo) HasRequiredVersionStr() bool`

HasRequiredVersionStr returns a boolean if a field has been set.

### GetSourceUrl

`func (o *UpdatesPackageVersionInfo) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *UpdatesPackageVersionInfo) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *UpdatesPackageVersionInfo) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *UpdatesPackageVersionInfo) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetChecksum

`func (o *UpdatesPackageVersionInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *UpdatesPackageVersionInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *UpdatesPackageVersionInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *UpdatesPackageVersionInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetTargetFilename

`func (o *UpdatesPackageVersionInfo) GetTargetFilename() string`

GetTargetFilename returns the TargetFilename field if non-nil, zero value otherwise.

### GetTargetFilenameOk

`func (o *UpdatesPackageVersionInfo) GetTargetFilenameOk() (*string, bool)`

GetTargetFilenameOk returns a tuple with the TargetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilename

`func (o *UpdatesPackageVersionInfo) SetTargetFilename(v string)`

SetTargetFilename sets TargetFilename field to given value.

### HasTargetFilename

`func (o *UpdatesPackageVersionInfo) HasTargetFilename() bool`

HasTargetFilename returns a boolean if a field has been set.

### GetInfoUrl

`func (o *UpdatesPackageVersionInfo) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *UpdatesPackageVersionInfo) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *UpdatesPackageVersionInfo) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *UpdatesPackageVersionInfo) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetRuntimes

`func (o *UpdatesPackageVersionInfo) GetRuntimes() string`

GetRuntimes returns the Runtimes field if non-nil, zero value otherwise.

### GetRuntimesOk

`func (o *UpdatesPackageVersionInfo) GetRuntimesOk() (*string, bool)`

GetRuntimesOk returns a tuple with the Runtimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimes

`func (o *UpdatesPackageVersionInfo) SetRuntimes(v string)`

SetRuntimes sets Runtimes field to given value.

### HasRuntimes

`func (o *UpdatesPackageVersionInfo) HasRuntimes() bool`

HasRuntimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



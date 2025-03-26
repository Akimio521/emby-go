# LibraryMediaFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SubFolders** | Pointer to [**[]LibrarySubFolder**](LibrarySubFolder.md) |  | [optional] 

## Methods

### NewLibraryMediaFolder

`func NewLibraryMediaFolder() *LibraryMediaFolder`

NewLibraryMediaFolder instantiates a new LibraryMediaFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryMediaFolderWithDefaults

`func NewLibraryMediaFolderWithDefaults() *LibraryMediaFolder`

NewLibraryMediaFolderWithDefaults instantiates a new LibraryMediaFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LibraryMediaFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryMediaFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryMediaFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryMediaFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *LibraryMediaFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryMediaFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryMediaFolder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LibraryMediaFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubFolders

`func (o *LibraryMediaFolder) GetSubFolders() []LibrarySubFolder`

GetSubFolders returns the SubFolders field if non-nil, zero value otherwise.

### GetSubFoldersOk

`func (o *LibraryMediaFolder) GetSubFoldersOk() (*[]LibrarySubFolder, bool)`

GetSubFoldersOk returns a tuple with the SubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFolders

`func (o *LibraryMediaFolder) SetSubFolders(v []LibrarySubFolder)`

SetSubFolders sets SubFolders field to given value.

### HasSubFolders

`func (o *LibraryMediaFolder) HasSubFolders() bool`

HasSubFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



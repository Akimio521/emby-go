# DisplayPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**IndexBy** | Pointer to **string** |  | [optional] 
**RememberIndexing** | Pointer to **bool** |  | [optional] 
**PrimaryImageHeight** | Pointer to **int32** |  | [optional] 
**PrimaryImageWidth** | Pointer to **int32** |  | [optional] 
**CustomPrefs** | Pointer to **map[string]string** |  | [optional] 
**ScrollDirection** | Pointer to **string** |  | [optional] 
**ShowBackdrop** | Pointer to **bool** |  | [optional] 
**RememberSorting** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **string** |  | [optional] 
**ShowSidebar** | Pointer to **bool** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 

## Methods

### NewDisplayPreferences

`func NewDisplayPreferences() *DisplayPreferences`

NewDisplayPreferences instantiates a new DisplayPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayPreferencesWithDefaults

`func NewDisplayPreferencesWithDefaults() *DisplayPreferences`

NewDisplayPreferencesWithDefaults instantiates a new DisplayPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisplayPreferences) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisplayPreferences) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisplayPreferences) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DisplayPreferences) HasId() bool`

HasId returns a boolean if a field has been set.

### GetViewType

`func (o *DisplayPreferences) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *DisplayPreferences) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *DisplayPreferences) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *DisplayPreferences) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetSortBy

`func (o *DisplayPreferences) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *DisplayPreferences) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *DisplayPreferences) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *DisplayPreferences) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetIndexBy

`func (o *DisplayPreferences) GetIndexBy() string`

GetIndexBy returns the IndexBy field if non-nil, zero value otherwise.

### GetIndexByOk

`func (o *DisplayPreferences) GetIndexByOk() (*string, bool)`

GetIndexByOk returns a tuple with the IndexBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBy

`func (o *DisplayPreferences) SetIndexBy(v string)`

SetIndexBy sets IndexBy field to given value.

### HasIndexBy

`func (o *DisplayPreferences) HasIndexBy() bool`

HasIndexBy returns a boolean if a field has been set.

### GetRememberIndexing

`func (o *DisplayPreferences) GetRememberIndexing() bool`

GetRememberIndexing returns the RememberIndexing field if non-nil, zero value otherwise.

### GetRememberIndexingOk

`func (o *DisplayPreferences) GetRememberIndexingOk() (*bool, bool)`

GetRememberIndexingOk returns a tuple with the RememberIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberIndexing

`func (o *DisplayPreferences) SetRememberIndexing(v bool)`

SetRememberIndexing sets RememberIndexing field to given value.

### HasRememberIndexing

`func (o *DisplayPreferences) HasRememberIndexing() bool`

HasRememberIndexing returns a boolean if a field has been set.

### GetPrimaryImageHeight

`func (o *DisplayPreferences) GetPrimaryImageHeight() int32`

GetPrimaryImageHeight returns the PrimaryImageHeight field if non-nil, zero value otherwise.

### GetPrimaryImageHeightOk

`func (o *DisplayPreferences) GetPrimaryImageHeightOk() (*int32, bool)`

GetPrimaryImageHeightOk returns a tuple with the PrimaryImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageHeight

`func (o *DisplayPreferences) SetPrimaryImageHeight(v int32)`

SetPrimaryImageHeight sets PrimaryImageHeight field to given value.

### HasPrimaryImageHeight

`func (o *DisplayPreferences) HasPrimaryImageHeight() bool`

HasPrimaryImageHeight returns a boolean if a field has been set.

### GetPrimaryImageWidth

`func (o *DisplayPreferences) GetPrimaryImageWidth() int32`

GetPrimaryImageWidth returns the PrimaryImageWidth field if non-nil, zero value otherwise.

### GetPrimaryImageWidthOk

`func (o *DisplayPreferences) GetPrimaryImageWidthOk() (*int32, bool)`

GetPrimaryImageWidthOk returns a tuple with the PrimaryImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageWidth

`func (o *DisplayPreferences) SetPrimaryImageWidth(v int32)`

SetPrimaryImageWidth sets PrimaryImageWidth field to given value.

### HasPrimaryImageWidth

`func (o *DisplayPreferences) HasPrimaryImageWidth() bool`

HasPrimaryImageWidth returns a boolean if a field has been set.

### GetCustomPrefs

`func (o *DisplayPreferences) GetCustomPrefs() map[string]string`

GetCustomPrefs returns the CustomPrefs field if non-nil, zero value otherwise.

### GetCustomPrefsOk

`func (o *DisplayPreferences) GetCustomPrefsOk() (*map[string]string, bool)`

GetCustomPrefsOk returns a tuple with the CustomPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrefs

`func (o *DisplayPreferences) SetCustomPrefs(v map[string]string)`

SetCustomPrefs sets CustomPrefs field to given value.

### HasCustomPrefs

`func (o *DisplayPreferences) HasCustomPrefs() bool`

HasCustomPrefs returns a boolean if a field has been set.

### GetScrollDirection

`func (o *DisplayPreferences) GetScrollDirection() string`

GetScrollDirection returns the ScrollDirection field if non-nil, zero value otherwise.

### GetScrollDirectionOk

`func (o *DisplayPreferences) GetScrollDirectionOk() (*string, bool)`

GetScrollDirectionOk returns a tuple with the ScrollDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollDirection

`func (o *DisplayPreferences) SetScrollDirection(v string)`

SetScrollDirection sets ScrollDirection field to given value.

### HasScrollDirection

`func (o *DisplayPreferences) HasScrollDirection() bool`

HasScrollDirection returns a boolean if a field has been set.

### GetShowBackdrop

`func (o *DisplayPreferences) GetShowBackdrop() bool`

GetShowBackdrop returns the ShowBackdrop field if non-nil, zero value otherwise.

### GetShowBackdropOk

`func (o *DisplayPreferences) GetShowBackdropOk() (*bool, bool)`

GetShowBackdropOk returns a tuple with the ShowBackdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBackdrop

`func (o *DisplayPreferences) SetShowBackdrop(v bool)`

SetShowBackdrop sets ShowBackdrop field to given value.

### HasShowBackdrop

`func (o *DisplayPreferences) HasShowBackdrop() bool`

HasShowBackdrop returns a boolean if a field has been set.

### GetRememberSorting

`func (o *DisplayPreferences) GetRememberSorting() bool`

GetRememberSorting returns the RememberSorting field if non-nil, zero value otherwise.

### GetRememberSortingOk

`func (o *DisplayPreferences) GetRememberSortingOk() (*bool, bool)`

GetRememberSortingOk returns a tuple with the RememberSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSorting

`func (o *DisplayPreferences) SetRememberSorting(v bool)`

SetRememberSorting sets RememberSorting field to given value.

### HasRememberSorting

`func (o *DisplayPreferences) HasRememberSorting() bool`

HasRememberSorting returns a boolean if a field has been set.

### GetSortOrder

`func (o *DisplayPreferences) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *DisplayPreferences) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *DisplayPreferences) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *DisplayPreferences) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetShowSidebar

`func (o *DisplayPreferences) GetShowSidebar() bool`

GetShowSidebar returns the ShowSidebar field if non-nil, zero value otherwise.

### GetShowSidebarOk

`func (o *DisplayPreferences) GetShowSidebarOk() (*bool, bool)`

GetShowSidebarOk returns a tuple with the ShowSidebar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSidebar

`func (o *DisplayPreferences) SetShowSidebar(v bool)`

SetShowSidebar sets ShowSidebar field to given value.

### HasShowSidebar

`func (o *DisplayPreferences) HasShowSidebar() bool`

HasShowSidebar returns a boolean if a field has been set.

### GetClient

`func (o *DisplayPreferences) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *DisplayPreferences) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *DisplayPreferences) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *DisplayPreferences) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



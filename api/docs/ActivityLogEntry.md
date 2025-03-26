# ActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**ShortOverview** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewActivityLogEntry

`func NewActivityLogEntry() *ActivityLogEntry`

NewActivityLogEntry instantiates a new ActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntryWithDefaults

`func NewActivityLogEntryWithDefaults() *ActivityLogEntry`

NewActivityLogEntryWithDefaults instantiates a new ActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityLogEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogEntry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivityLogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *ActivityLogEntry) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ActivityLogEntry) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ActivityLogEntry) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ActivityLogEntry) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetShortOverview

`func (o *ActivityLogEntry) GetShortOverview() string`

GetShortOverview returns the ShortOverview field if non-nil, zero value otherwise.

### GetShortOverviewOk

`func (o *ActivityLogEntry) GetShortOverviewOk() (*string, bool)`

GetShortOverviewOk returns a tuple with the ShortOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOverview

`func (o *ActivityLogEntry) SetShortOverview(v string)`

SetShortOverview sets ShortOverview field to given value.

### HasShortOverview

`func (o *ActivityLogEntry) HasShortOverview() bool`

HasShortOverview returns a boolean if a field has been set.

### GetType

`func (o *ActivityLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemId

`func (o *ActivityLogEntry) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ActivityLogEntry) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ActivityLogEntry) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ActivityLogEntry) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDate

`func (o *ActivityLogEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityLogEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityLogEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ActivityLogEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUserId

`func (o *ActivityLogEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActivityLogEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActivityLogEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActivityLogEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *ActivityLogEntry) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *ActivityLogEntry) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *ActivityLogEntry) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *ActivityLogEntry) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### GetSeverity

`func (o *ActivityLogEntry) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ActivityLogEntry) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ActivityLogEntry) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ActivityLogEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



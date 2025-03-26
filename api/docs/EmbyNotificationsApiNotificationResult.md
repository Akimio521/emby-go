# EmbyNotificationsApiNotificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]EmbyNotificationsApiNotification**](EmbyNotificationsApiNotification.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewEmbyNotificationsApiNotificationResult

`func NewEmbyNotificationsApiNotificationResult() *EmbyNotificationsApiNotificationResult`

NewEmbyNotificationsApiNotificationResult instantiates a new EmbyNotificationsApiNotificationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbyNotificationsApiNotificationResultWithDefaults

`func NewEmbyNotificationsApiNotificationResultWithDefaults() *EmbyNotificationsApiNotificationResult`

NewEmbyNotificationsApiNotificationResultWithDefaults instantiates a new EmbyNotificationsApiNotificationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *EmbyNotificationsApiNotificationResult) GetNotifications() []EmbyNotificationsApiNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *EmbyNotificationsApiNotificationResult) GetNotificationsOk() (*[]EmbyNotificationsApiNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *EmbyNotificationsApiNotificationResult) SetNotifications(v []EmbyNotificationsApiNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *EmbyNotificationsApiNotificationResult) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *EmbyNotificationsApiNotificationResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *EmbyNotificationsApiNotificationResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *EmbyNotificationsApiNotificationResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *EmbyNotificationsApiNotificationResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



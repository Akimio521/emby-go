# MediaInfoPlaybackInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaSources** | Pointer to [**[]MediaSourceInfo**](MediaSourceInfo.md) |  | [optional] 
**PlaySessionId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaInfoPlaybackInfoResponse

`func NewMediaInfoPlaybackInfoResponse() *MediaInfoPlaybackInfoResponse`

NewMediaInfoPlaybackInfoResponse instantiates a new MediaInfoPlaybackInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaInfoPlaybackInfoResponseWithDefaults

`func NewMediaInfoPlaybackInfoResponseWithDefaults() *MediaInfoPlaybackInfoResponse`

NewMediaInfoPlaybackInfoResponseWithDefaults instantiates a new MediaInfoPlaybackInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaSources

`func (o *MediaInfoPlaybackInfoResponse) GetMediaSources() []MediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *MediaInfoPlaybackInfoResponse) GetMediaSourcesOk() (*[]MediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *MediaInfoPlaybackInfoResponse) SetMediaSources(v []MediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *MediaInfoPlaybackInfoResponse) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *MediaInfoPlaybackInfoResponse) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *MediaInfoPlaybackInfoResponse) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *MediaInfoPlaybackInfoResponse) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *MediaInfoPlaybackInfoResponse) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### GetErrorCode

`func (o *MediaInfoPlaybackInfoResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MediaInfoPlaybackInfoResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MediaInfoPlaybackInfoResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MediaInfoPlaybackInfoResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



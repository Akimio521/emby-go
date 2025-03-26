# EmbyReportsApiModelReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]EmbyReportsApiModelReportRow**](EmbyReportsApiModelReportRow.md) |  | [optional] 
**Headers** | Pointer to [**[]EmbyReportsApiModelReportHeader**](EmbyReportsApiModelReportHeader.md) |  | [optional] 
**Groups** | Pointer to [**[]EmbyReportsApiModelReportGroup**](EmbyReportsApiModelReportGroup.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 
**IsGrouped** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmbyReportsApiModelReportResult

`func NewEmbyReportsApiModelReportResult() *EmbyReportsApiModelReportResult`

NewEmbyReportsApiModelReportResult instantiates a new EmbyReportsApiModelReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbyReportsApiModelReportResultWithDefaults

`func NewEmbyReportsApiModelReportResultWithDefaults() *EmbyReportsApiModelReportResult`

NewEmbyReportsApiModelReportResultWithDefaults instantiates a new EmbyReportsApiModelReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *EmbyReportsApiModelReportResult) GetRows() []EmbyReportsApiModelReportRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EmbyReportsApiModelReportResult) GetRowsOk() (*[]EmbyReportsApiModelReportRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EmbyReportsApiModelReportResult) SetRows(v []EmbyReportsApiModelReportRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EmbyReportsApiModelReportResult) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetHeaders

`func (o *EmbyReportsApiModelReportResult) GetHeaders() []EmbyReportsApiModelReportHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmbyReportsApiModelReportResult) GetHeadersOk() (*[]EmbyReportsApiModelReportHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmbyReportsApiModelReportResult) SetHeaders(v []EmbyReportsApiModelReportHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmbyReportsApiModelReportResult) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetGroups

`func (o *EmbyReportsApiModelReportResult) GetGroups() []EmbyReportsApiModelReportGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmbyReportsApiModelReportResult) GetGroupsOk() (*[]EmbyReportsApiModelReportGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmbyReportsApiModelReportResult) SetGroups(v []EmbyReportsApiModelReportGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmbyReportsApiModelReportResult) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *EmbyReportsApiModelReportResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *EmbyReportsApiModelReportResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *EmbyReportsApiModelReportResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *EmbyReportsApiModelReportResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetIsGrouped

`func (o *EmbyReportsApiModelReportResult) GetIsGrouped() bool`

GetIsGrouped returns the IsGrouped field if non-nil, zero value otherwise.

### GetIsGroupedOk

`func (o *EmbyReportsApiModelReportResult) GetIsGroupedOk() (*bool, bool)`

GetIsGroupedOk returns a tuple with the IsGrouped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGrouped

`func (o *EmbyReportsApiModelReportResult) SetIsGrouped(v bool)`

SetIsGrouped sets IsGrouped field to given value.

### HasIsGrouped

`func (o *EmbyReportsApiModelReportResult) HasIsGrouped() bool`

HasIsGrouped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



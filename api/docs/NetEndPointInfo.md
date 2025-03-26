# NetEndPointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocal** | Pointer to **bool** |  | [optional] 
**IsInNetwork** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetEndPointInfo

`func NewNetEndPointInfo() *NetEndPointInfo`

NewNetEndPointInfo instantiates a new NetEndPointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetEndPointInfoWithDefaults

`func NewNetEndPointInfoWithDefaults() *NetEndPointInfo`

NewNetEndPointInfoWithDefaults instantiates a new NetEndPointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocal

`func (o *NetEndPointInfo) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *NetEndPointInfo) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *NetEndPointInfo) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *NetEndPointInfo) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetIsInNetwork

`func (o *NetEndPointInfo) GetIsInNetwork() bool`

GetIsInNetwork returns the IsInNetwork field if non-nil, zero value otherwise.

### GetIsInNetworkOk

`func (o *NetEndPointInfo) GetIsInNetworkOk() (*bool, bool)`

GetIsInNetworkOk returns a tuple with the IsInNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInNetwork

`func (o *NetEndPointInfo) SetIsInNetwork(v bool)`

SetIsInNetwork sets IsInNetwork field to given value.

### HasIsInNetwork

`func (o *NetEndPointInfo) HasIsInNetwork() bool`

HasIsInNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



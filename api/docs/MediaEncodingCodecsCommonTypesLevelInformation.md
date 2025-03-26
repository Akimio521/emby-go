# MediaEncodingCodecsCommonTypesLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ordinal** | Pointer to **NullableInt32** |  | [optional] 
**MaxBitRate** | Pointer to [**MediaEncodingCodecsCommonTypesBitRate**](MediaEncodingCodecsCommonTypesBitRate.md) |  | [optional] 
**MaxBitRateDisplay** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ResolutionRates** | Pointer to [**[]MediaEncodingCodecsCommonTypesResolutionWithRate**](MediaEncodingCodecsCommonTypesResolutionWithRate.md) |  | [optional] 
**ResolutionRateStrings** | Pointer to **[]string** |  | [optional] 
**ResolutionRatesDisplay** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaEncodingCodecsCommonTypesLevelInformation

`func NewMediaEncodingCodecsCommonTypesLevelInformation() *MediaEncodingCodecsCommonTypesLevelInformation`

NewMediaEncodingCodecsCommonTypesLevelInformation instantiates a new MediaEncodingCodecsCommonTypesLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaEncodingCodecsCommonTypesLevelInformationWithDefaults

`func NewMediaEncodingCodecsCommonTypesLevelInformationWithDefaults() *MediaEncodingCodecsCommonTypesLevelInformation`

NewMediaEncodingCodecsCommonTypesLevelInformationWithDefaults instantiates a new MediaEncodingCodecsCommonTypesLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetDescription

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrdinal

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetOrdinal() int32`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetOrdinalOk() (*int32, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetOrdinal(v int32)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### SetOrdinalNil

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetOrdinalNil(b bool)`

 SetOrdinalNil sets the value for Ordinal to be an explicit nil

### UnsetOrdinal
`func (o *MediaEncodingCodecsCommonTypesLevelInformation) UnsetOrdinal()`

UnsetOrdinal ensures that no value is present for Ordinal, not even an explicit nil
### GetMaxBitRate

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetMaxBitRate() MediaEncodingCodecsCommonTypesBitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetMaxBitRateOk() (*MediaEncodingCodecsCommonTypesBitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetMaxBitRate(v MediaEncodingCodecsCommonTypesBitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetMaxBitRateDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetMaxBitRateDisplay() string`

GetMaxBitRateDisplay returns the MaxBitRateDisplay field if non-nil, zero value otherwise.

### GetMaxBitRateDisplayOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetMaxBitRateDisplayOk() (*string, bool)`

GetMaxBitRateDisplayOk returns a tuple with the MaxBitRateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRateDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetMaxBitRateDisplay(v string)`

SetMaxBitRateDisplay sets MaxBitRateDisplay field to given value.

### HasMaxBitRateDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasMaxBitRateDisplay() bool`

HasMaxBitRateDisplay returns a boolean if a field has been set.

### GetId

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResolutionRates

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRates() []MediaEncodingCodecsCommonTypesResolutionWithRate`

GetResolutionRates returns the ResolutionRates field if non-nil, zero value otherwise.

### GetResolutionRatesOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRatesOk() (*[]MediaEncodingCodecsCommonTypesResolutionWithRate, bool)`

GetResolutionRatesOk returns a tuple with the ResolutionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRates

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetResolutionRates(v []MediaEncodingCodecsCommonTypesResolutionWithRate)`

SetResolutionRates sets ResolutionRates field to given value.

### HasResolutionRates

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasResolutionRates() bool`

HasResolutionRates returns a boolean if a field has been set.

### GetResolutionRateStrings

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRateStrings() []string`

GetResolutionRateStrings returns the ResolutionRateStrings field if non-nil, zero value otherwise.

### GetResolutionRateStringsOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRateStringsOk() (*[]string, bool)`

GetResolutionRateStringsOk returns a tuple with the ResolutionRateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRateStrings

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetResolutionRateStrings(v []string)`

SetResolutionRateStrings sets ResolutionRateStrings field to given value.

### HasResolutionRateStrings

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasResolutionRateStrings() bool`

HasResolutionRateStrings returns a boolean if a field has been set.

### GetResolutionRatesDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRatesDisplay() string`

GetResolutionRatesDisplay returns the ResolutionRatesDisplay field if non-nil, zero value otherwise.

### GetResolutionRatesDisplayOk

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) GetResolutionRatesDisplayOk() (*string, bool)`

GetResolutionRatesDisplayOk returns a tuple with the ResolutionRatesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRatesDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) SetResolutionRatesDisplay(v string)`

SetResolutionRatesDisplay sets ResolutionRatesDisplay field to given value.

### HasResolutionRatesDisplay

`func (o *MediaEncodingCodecsCommonTypesLevelInformation) HasResolutionRatesDisplay() bool`

HasResolutionRatesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



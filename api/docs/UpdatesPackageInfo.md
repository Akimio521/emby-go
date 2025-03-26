# UpdatesPackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**IsPremium** | Pointer to **bool** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**RichDescUrl** | Pointer to **string** |  | [optional] 
**ThumbImage** | Pointer to **string** |  | [optional] 
**PreviewImage** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TargetFilename** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**TileColor** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**RegInfo** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**TargetSystem** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**TotalRatings** | Pointer to **NullableInt32** |  | [optional] 
**AvgRating** | Pointer to **float32** |  | [optional] 
**IsRegistered** | Pointer to **bool** |  | [optional] 
**ExpDate** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to [**[]UpdatesPackageVersionInfo**](UpdatesPackageVersionInfo.md) |  | [optional] 
**EnableInAppStore** | Pointer to **bool** |  | [optional] 
**Installs** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdatesPackageInfo

`func NewUpdatesPackageInfo() *UpdatesPackageInfo`

NewUpdatesPackageInfo instantiates a new UpdatesPackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesPackageInfoWithDefaults

`func NewUpdatesPackageInfoWithDefaults() *UpdatesPackageInfo`

NewUpdatesPackageInfoWithDefaults instantiates a new UpdatesPackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatesPackageInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatesPackageInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatesPackageInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatesPackageInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdatesPackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatesPackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatesPackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatesPackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *UpdatesPackageInfo) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *UpdatesPackageInfo) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *UpdatesPackageInfo) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *UpdatesPackageInfo) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetOverview

`func (o *UpdatesPackageInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *UpdatesPackageInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *UpdatesPackageInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *UpdatesPackageInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetIsPremium

`func (o *UpdatesPackageInfo) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *UpdatesPackageInfo) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *UpdatesPackageInfo) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *UpdatesPackageInfo) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetAdult

`func (o *UpdatesPackageInfo) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *UpdatesPackageInfo) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *UpdatesPackageInfo) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *UpdatesPackageInfo) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetRichDescUrl

`func (o *UpdatesPackageInfo) GetRichDescUrl() string`

GetRichDescUrl returns the RichDescUrl field if non-nil, zero value otherwise.

### GetRichDescUrlOk

`func (o *UpdatesPackageInfo) GetRichDescUrlOk() (*string, bool)`

GetRichDescUrlOk returns a tuple with the RichDescUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichDescUrl

`func (o *UpdatesPackageInfo) SetRichDescUrl(v string)`

SetRichDescUrl sets RichDescUrl field to given value.

### HasRichDescUrl

`func (o *UpdatesPackageInfo) HasRichDescUrl() bool`

HasRichDescUrl returns a boolean if a field has been set.

### GetThumbImage

`func (o *UpdatesPackageInfo) GetThumbImage() string`

GetThumbImage returns the ThumbImage field if non-nil, zero value otherwise.

### GetThumbImageOk

`func (o *UpdatesPackageInfo) GetThumbImageOk() (*string, bool)`

GetThumbImageOk returns a tuple with the ThumbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImage

`func (o *UpdatesPackageInfo) SetThumbImage(v string)`

SetThumbImage sets ThumbImage field to given value.

### HasThumbImage

`func (o *UpdatesPackageInfo) HasThumbImage() bool`

HasThumbImage returns a boolean if a field has been set.

### GetPreviewImage

`func (o *UpdatesPackageInfo) GetPreviewImage() string`

GetPreviewImage returns the PreviewImage field if non-nil, zero value otherwise.

### GetPreviewImageOk

`func (o *UpdatesPackageInfo) GetPreviewImageOk() (*string, bool)`

GetPreviewImageOk returns a tuple with the PreviewImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImage

`func (o *UpdatesPackageInfo) SetPreviewImage(v string)`

SetPreviewImage sets PreviewImage field to given value.

### HasPreviewImage

`func (o *UpdatesPackageInfo) HasPreviewImage() bool`

HasPreviewImage returns a boolean if a field has been set.

### GetType

`func (o *UpdatesPackageInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatesPackageInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatesPackageInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatesPackageInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTargetFilename

`func (o *UpdatesPackageInfo) GetTargetFilename() string`

GetTargetFilename returns the TargetFilename field if non-nil, zero value otherwise.

### GetTargetFilenameOk

`func (o *UpdatesPackageInfo) GetTargetFilenameOk() (*string, bool)`

GetTargetFilenameOk returns a tuple with the TargetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilename

`func (o *UpdatesPackageInfo) SetTargetFilename(v string)`

SetTargetFilename sets TargetFilename field to given value.

### HasTargetFilename

`func (o *UpdatesPackageInfo) HasTargetFilename() bool`

HasTargetFilename returns a boolean if a field has been set.

### GetOwner

`func (o *UpdatesPackageInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdatesPackageInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdatesPackageInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdatesPackageInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *UpdatesPackageInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdatesPackageInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdatesPackageInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdatesPackageInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTileColor

`func (o *UpdatesPackageInfo) GetTileColor() string`

GetTileColor returns the TileColor field if non-nil, zero value otherwise.

### GetTileColorOk

`func (o *UpdatesPackageInfo) GetTileColorOk() (*string, bool)`

GetTileColorOk returns a tuple with the TileColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileColor

`func (o *UpdatesPackageInfo) SetTileColor(v string)`

SetTileColor sets TileColor field to given value.

### HasTileColor

`func (o *UpdatesPackageInfo) HasTileColor() bool`

HasTileColor returns a boolean if a field has been set.

### GetFeatureId

`func (o *UpdatesPackageInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *UpdatesPackageInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *UpdatesPackageInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *UpdatesPackageInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetRegInfo

`func (o *UpdatesPackageInfo) GetRegInfo() string`

GetRegInfo returns the RegInfo field if non-nil, zero value otherwise.

### GetRegInfoOk

`func (o *UpdatesPackageInfo) GetRegInfoOk() (*string, bool)`

GetRegInfoOk returns a tuple with the RegInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegInfo

`func (o *UpdatesPackageInfo) SetRegInfo(v string)`

SetRegInfo sets RegInfo field to given value.

### HasRegInfo

`func (o *UpdatesPackageInfo) HasRegInfo() bool`

HasRegInfo returns a boolean if a field has been set.

### GetPrice

`func (o *UpdatesPackageInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatesPackageInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatesPackageInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdatesPackageInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTargetSystem

`func (o *UpdatesPackageInfo) GetTargetSystem() string`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *UpdatesPackageInfo) GetTargetSystemOk() (*string, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *UpdatesPackageInfo) SetTargetSystem(v string)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *UpdatesPackageInfo) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.

### GetGuid

`func (o *UpdatesPackageInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *UpdatesPackageInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *UpdatesPackageInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *UpdatesPackageInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetTotalRatings

`func (o *UpdatesPackageInfo) GetTotalRatings() int32`

GetTotalRatings returns the TotalRatings field if non-nil, zero value otherwise.

### GetTotalRatingsOk

`func (o *UpdatesPackageInfo) GetTotalRatingsOk() (*int32, bool)`

GetTotalRatingsOk returns a tuple with the TotalRatings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRatings

`func (o *UpdatesPackageInfo) SetTotalRatings(v int32)`

SetTotalRatings sets TotalRatings field to given value.

### HasTotalRatings

`func (o *UpdatesPackageInfo) HasTotalRatings() bool`

HasTotalRatings returns a boolean if a field has been set.

### SetTotalRatingsNil

`func (o *UpdatesPackageInfo) SetTotalRatingsNil(b bool)`

 SetTotalRatingsNil sets the value for TotalRatings to be an explicit nil

### UnsetTotalRatings
`func (o *UpdatesPackageInfo) UnsetTotalRatings()`

UnsetTotalRatings ensures that no value is present for TotalRatings, not even an explicit nil
### GetAvgRating

`func (o *UpdatesPackageInfo) GetAvgRating() float32`

GetAvgRating returns the AvgRating field if non-nil, zero value otherwise.

### GetAvgRatingOk

`func (o *UpdatesPackageInfo) GetAvgRatingOk() (*float32, bool)`

GetAvgRatingOk returns a tuple with the AvgRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRating

`func (o *UpdatesPackageInfo) SetAvgRating(v float32)`

SetAvgRating sets AvgRating field to given value.

### HasAvgRating

`func (o *UpdatesPackageInfo) HasAvgRating() bool`

HasAvgRating returns a boolean if a field has been set.

### GetIsRegistered

`func (o *UpdatesPackageInfo) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *UpdatesPackageInfo) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *UpdatesPackageInfo) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *UpdatesPackageInfo) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### GetExpDate

`func (o *UpdatesPackageInfo) GetExpDate() time.Time`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *UpdatesPackageInfo) GetExpDateOk() (*time.Time, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *UpdatesPackageInfo) SetExpDate(v time.Time)`

SetExpDate sets ExpDate field to given value.

### HasExpDate

`func (o *UpdatesPackageInfo) HasExpDate() bool`

HasExpDate returns a boolean if a field has been set.

### GetVersions

`func (o *UpdatesPackageInfo) GetVersions() []UpdatesPackageVersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *UpdatesPackageInfo) GetVersionsOk() (*[]UpdatesPackageVersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *UpdatesPackageInfo) SetVersions(v []UpdatesPackageVersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *UpdatesPackageInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetEnableInAppStore

`func (o *UpdatesPackageInfo) GetEnableInAppStore() bool`

GetEnableInAppStore returns the EnableInAppStore field if non-nil, zero value otherwise.

### GetEnableInAppStoreOk

`func (o *UpdatesPackageInfo) GetEnableInAppStoreOk() (*bool, bool)`

GetEnableInAppStoreOk returns a tuple with the EnableInAppStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInAppStore

`func (o *UpdatesPackageInfo) SetEnableInAppStore(v bool)`

SetEnableInAppStore sets EnableInAppStore field to given value.

### HasEnableInAppStore

`func (o *UpdatesPackageInfo) HasEnableInAppStore() bool`

HasEnableInAppStore returns a boolean if a field has been set.

### GetInstalls

`func (o *UpdatesPackageInfo) GetInstalls() int32`

GetInstalls returns the Installs field if non-nil, zero value otherwise.

### GetInstallsOk

`func (o *UpdatesPackageInfo) GetInstallsOk() (*int32, bool)`

GetInstallsOk returns a tuple with the Installs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalls

`func (o *UpdatesPackageInfo) SetInstalls(v int32)`

SetInstalls sets Installs field to given value.

### HasInstalls

`func (o *UpdatesPackageInfo) HasInstalls() bool`

HasInstalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



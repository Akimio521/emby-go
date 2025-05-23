/*
Emby Server API

Explore the Emby Server API

API version: 4.1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the SyncSyncJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncSyncJob{}

// SyncSyncJob struct for SyncSyncJob
type SyncSyncJob struct {
	Id *int64 `json:"Id,omitempty"`
	TargetId *string `json:"TargetId,omitempty"`
	TargetName *string `json:"TargetName,omitempty"`
	Quality *string `json:"Quality,omitempty"`
	Bitrate NullableInt32 `json:"Bitrate,omitempty"`
	Profile *string `json:"Profile,omitempty"`
	Category *string `json:"Category,omitempty"`
	ParentId *int64 `json:"ParentId,omitempty"`
	Progress *float64 `json:"Progress,omitempty"`
	Name *string `json:"Name,omitempty"`
	Status *string `json:"Status,omitempty"`
	UserId *int64 `json:"UserId,omitempty"`
	UnwatchedOnly *bool `json:"UnwatchedOnly,omitempty"`
	SyncNewContent *bool `json:"SyncNewContent,omitempty"`
	ItemLimit NullableInt32 `json:"ItemLimit,omitempty"`
	RequestedItemIds []int64 `json:"RequestedItemIds,omitempty"`
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	DateLastModified *time.Time `json:"DateLastModified,omitempty"`
	ItemCount *int32 `json:"ItemCount,omitempty"`
	ParentName *string `json:"ParentName,omitempty"`
	PrimaryImageItemId *string `json:"PrimaryImageItemId,omitempty"`
	PrimaryImageTag *string `json:"PrimaryImageTag,omitempty"`
}

// NewSyncSyncJob instantiates a new SyncSyncJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncSyncJob() *SyncSyncJob {
	this := SyncSyncJob{}
	return &this
}

// NewSyncSyncJobWithDefaults instantiates a new SyncSyncJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncSyncJobWithDefaults() *SyncSyncJob {
	this := SyncSyncJob{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyncSyncJob) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyncSyncJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SyncSyncJob) SetId(v int64) {
	o.Id = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SyncSyncJob) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SyncSyncJob) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SyncSyncJob) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *SyncSyncJob) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *SyncSyncJob) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *SyncSyncJob) SetTargetName(v string) {
	o.TargetName = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *SyncSyncJob) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *SyncSyncJob) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *SyncSyncJob) SetQuality(v string) {
	o.Quality = &v
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncSyncJob) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncSyncJob) GetBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *SyncSyncJob) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *SyncSyncJob) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}
// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *SyncSyncJob) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *SyncSyncJob) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SyncSyncJob) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SyncSyncJob) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *SyncSyncJob) SetProfile(v string) {
	o.Profile = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SyncSyncJob) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SyncSyncJob) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SyncSyncJob) SetCategory(v string) {
	o.Category = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SyncSyncJob) GetParentId() int64 {
	if o == nil || IsNil(o.ParentId) {
		var ret int64
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetParentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SyncSyncJob) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int64 and assigns it to the ParentId field.
func (o *SyncSyncJob) SetParentId(v int64) {
	o.ParentId = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *SyncSyncJob) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *SyncSyncJob) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *SyncSyncJob) SetProgress(v float64) {
	o.Progress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyncSyncJob) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyncSyncJob) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyncSyncJob) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyncSyncJob) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyncSyncJob) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyncSyncJob) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SyncSyncJob) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SyncSyncJob) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *SyncSyncJob) SetUserId(v int64) {
	o.UserId = &v
}

// GetUnwatchedOnly returns the UnwatchedOnly field value if set, zero value otherwise.
func (o *SyncSyncJob) GetUnwatchedOnly() bool {
	if o == nil || IsNil(o.UnwatchedOnly) {
		var ret bool
		return ret
	}
	return *o.UnwatchedOnly
}

// GetUnwatchedOnlyOk returns a tuple with the UnwatchedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetUnwatchedOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.UnwatchedOnly) {
		return nil, false
	}
	return o.UnwatchedOnly, true
}

// HasUnwatchedOnly returns a boolean if a field has been set.
func (o *SyncSyncJob) HasUnwatchedOnly() bool {
	if o != nil && !IsNil(o.UnwatchedOnly) {
		return true
	}

	return false
}

// SetUnwatchedOnly gets a reference to the given bool and assigns it to the UnwatchedOnly field.
func (o *SyncSyncJob) SetUnwatchedOnly(v bool) {
	o.UnwatchedOnly = &v
}

// GetSyncNewContent returns the SyncNewContent field value if set, zero value otherwise.
func (o *SyncSyncJob) GetSyncNewContent() bool {
	if o == nil || IsNil(o.SyncNewContent) {
		var ret bool
		return ret
	}
	return *o.SyncNewContent
}

// GetSyncNewContentOk returns a tuple with the SyncNewContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetSyncNewContentOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncNewContent) {
		return nil, false
	}
	return o.SyncNewContent, true
}

// HasSyncNewContent returns a boolean if a field has been set.
func (o *SyncSyncJob) HasSyncNewContent() bool {
	if o != nil && !IsNil(o.SyncNewContent) {
		return true
	}

	return false
}

// SetSyncNewContent gets a reference to the given bool and assigns it to the SyncNewContent field.
func (o *SyncSyncJob) SetSyncNewContent(v bool) {
	o.SyncNewContent = &v
}

// GetItemLimit returns the ItemLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyncSyncJob) GetItemLimit() int32 {
	if o == nil || IsNil(o.ItemLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.ItemLimit.Get()
}

// GetItemLimitOk returns a tuple with the ItemLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncSyncJob) GetItemLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemLimit.Get(), o.ItemLimit.IsSet()
}

// HasItemLimit returns a boolean if a field has been set.
func (o *SyncSyncJob) HasItemLimit() bool {
	if o != nil && o.ItemLimit.IsSet() {
		return true
	}

	return false
}

// SetItemLimit gets a reference to the given NullableInt32 and assigns it to the ItemLimit field.
func (o *SyncSyncJob) SetItemLimit(v int32) {
	o.ItemLimit.Set(&v)
}
// SetItemLimitNil sets the value for ItemLimit to be an explicit nil
func (o *SyncSyncJob) SetItemLimitNil() {
	o.ItemLimit.Set(nil)
}

// UnsetItemLimit ensures that no value is present for ItemLimit, not even an explicit nil
func (o *SyncSyncJob) UnsetItemLimit() {
	o.ItemLimit.Unset()
}

// GetRequestedItemIds returns the RequestedItemIds field value if set, zero value otherwise.
func (o *SyncSyncJob) GetRequestedItemIds() []int64 {
	if o == nil || IsNil(o.RequestedItemIds) {
		var ret []int64
		return ret
	}
	return o.RequestedItemIds
}

// GetRequestedItemIdsOk returns a tuple with the RequestedItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetRequestedItemIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.RequestedItemIds) {
		return nil, false
	}
	return o.RequestedItemIds, true
}

// HasRequestedItemIds returns a boolean if a field has been set.
func (o *SyncSyncJob) HasRequestedItemIds() bool {
	if o != nil && !IsNil(o.RequestedItemIds) {
		return true
	}

	return false
}

// SetRequestedItemIds gets a reference to the given []int64 and assigns it to the RequestedItemIds field.
func (o *SyncSyncJob) SetRequestedItemIds(v []int64) {
	o.RequestedItemIds = v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SyncSyncJob) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SyncSyncJob) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SyncSyncJob) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateLastModified returns the DateLastModified field value if set, zero value otherwise.
func (o *SyncSyncJob) GetDateLastModified() time.Time {
	if o == nil || IsNil(o.DateLastModified) {
		var ret time.Time
		return ret
	}
	return *o.DateLastModified
}

// GetDateLastModifiedOk returns a tuple with the DateLastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetDateLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateLastModified) {
		return nil, false
	}
	return o.DateLastModified, true
}

// HasDateLastModified returns a boolean if a field has been set.
func (o *SyncSyncJob) HasDateLastModified() bool {
	if o != nil && !IsNil(o.DateLastModified) {
		return true
	}

	return false
}

// SetDateLastModified gets a reference to the given time.Time and assigns it to the DateLastModified field.
func (o *SyncSyncJob) SetDateLastModified(v time.Time) {
	o.DateLastModified = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *SyncSyncJob) GetItemCount() int32 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetItemCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *SyncSyncJob) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *SyncSyncJob) SetItemCount(v int32) {
	o.ItemCount = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *SyncSyncJob) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *SyncSyncJob) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *SyncSyncJob) SetParentName(v string) {
	o.ParentName = &v
}

// GetPrimaryImageItemId returns the PrimaryImageItemId field value if set, zero value otherwise.
func (o *SyncSyncJob) GetPrimaryImageItemId() string {
	if o == nil || IsNil(o.PrimaryImageItemId) {
		var ret string
		return ret
	}
	return *o.PrimaryImageItemId
}

// GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetPrimaryImageItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryImageItemId) {
		return nil, false
	}
	return o.PrimaryImageItemId, true
}

// HasPrimaryImageItemId returns a boolean if a field has been set.
func (o *SyncSyncJob) HasPrimaryImageItemId() bool {
	if o != nil && !IsNil(o.PrimaryImageItemId) {
		return true
	}

	return false
}

// SetPrimaryImageItemId gets a reference to the given string and assigns it to the PrimaryImageItemId field.
func (o *SyncSyncJob) SetPrimaryImageItemId(v string) {
	o.PrimaryImageItemId = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise.
func (o *SyncSyncJob) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncSyncJob) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryImageTag) {
		return nil, false
	}
	return o.PrimaryImageTag, true
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *SyncSyncJob) HasPrimaryImageTag() bool {
	if o != nil && !IsNil(o.PrimaryImageTag) {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given string and assigns it to the PrimaryImageTag field.
func (o *SyncSyncJob) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag = &v
}

func (o SyncSyncJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncSyncJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.TargetId) {
		toSerialize["TargetId"] = o.TargetId
	}
	if !IsNil(o.TargetName) {
		toSerialize["TargetName"] = o.TargetName
	}
	if !IsNil(o.Quality) {
		toSerialize["Quality"] = o.Quality
	}
	if o.Bitrate.IsSet() {
		toSerialize["Bitrate"] = o.Bitrate.Get()
	}
	if !IsNil(o.Profile) {
		toSerialize["Profile"] = o.Profile
	}
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.ParentId) {
		toSerialize["ParentId"] = o.ParentId
	}
	if !IsNil(o.Progress) {
		toSerialize["Progress"] = o.Progress
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.UnwatchedOnly) {
		toSerialize["UnwatchedOnly"] = o.UnwatchedOnly
	}
	if !IsNil(o.SyncNewContent) {
		toSerialize["SyncNewContent"] = o.SyncNewContent
	}
	if o.ItemLimit.IsSet() {
		toSerialize["ItemLimit"] = o.ItemLimit.Get()
	}
	if !IsNil(o.RequestedItemIds) {
		toSerialize["RequestedItemIds"] = o.RequestedItemIds
	}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if !IsNil(o.DateLastModified) {
		toSerialize["DateLastModified"] = o.DateLastModified
	}
	if !IsNil(o.ItemCount) {
		toSerialize["ItemCount"] = o.ItemCount
	}
	if !IsNil(o.ParentName) {
		toSerialize["ParentName"] = o.ParentName
	}
	if !IsNil(o.PrimaryImageItemId) {
		toSerialize["PrimaryImageItemId"] = o.PrimaryImageItemId
	}
	if !IsNil(o.PrimaryImageTag) {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag
	}
	return toSerialize, nil
}

type NullableSyncSyncJob struct {
	value *SyncSyncJob
	isSet bool
}

func (v NullableSyncSyncJob) Get() *SyncSyncJob {
	return v.value
}

func (v *NullableSyncSyncJob) Set(val *SyncSyncJob) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncSyncJob) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncSyncJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncSyncJob(val *SyncSyncJob) *NullableSyncSyncJob {
	return &NullableSyncSyncJob{value: val, isSet: true}
}

func (v NullableSyncSyncJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncSyncJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



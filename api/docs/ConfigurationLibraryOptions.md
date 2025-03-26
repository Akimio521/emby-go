# ConfigurationLibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableArchiveMediaFiles** | Pointer to **bool** |  | [optional] 
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**DownloadImagesInAdvance** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]ConfigurationMediaPathInfo**](ConfigurationMediaPathInfo.md) |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**SaveLocalThumbnailSets** | Pointer to **bool** |  | [optional] 
**ImportMissingEpisodes** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableAudioResume** | Pointer to **bool** |  | [optional] 
**AutomaticRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**SeasonZeroDisplayName** | Pointer to **string** |  | [optional] 
**MetadataSavers** | Pointer to **[]string** |  | [optional] 
**DisabledLocalMetadataReaders** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledSubtitleFetchers** | Pointer to **[]string** |  | [optional] 
**SubtitleFetcherOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**ForcedSubtitlesOnly** | Pointer to **bool** |  | [optional] 
**TypeOptions** | Pointer to [**[]ConfigurationTypeOptions**](ConfigurationTypeOptions.md) |  | [optional] 
**CollapseSingleItemFolders** | Pointer to **bool** |  | [optional] 
**MinResumePct** | Pointer to **int32** |  | [optional] 
**MaxResumePct** | Pointer to **int32** |  | [optional] 
**MinResumeDurationSeconds** | Pointer to **int32** |  | [optional] 
**ThumbnailImagesIntervalSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewConfigurationLibraryOptions

`func NewConfigurationLibraryOptions() *ConfigurationLibraryOptions`

NewConfigurationLibraryOptions instantiates a new ConfigurationLibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationLibraryOptionsWithDefaults

`func NewConfigurationLibraryOptionsWithDefaults() *ConfigurationLibraryOptions`

NewConfigurationLibraryOptionsWithDefaults instantiates a new ConfigurationLibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableArchiveMediaFiles

`func (o *ConfigurationLibraryOptions) GetEnableArchiveMediaFiles() bool`

GetEnableArchiveMediaFiles returns the EnableArchiveMediaFiles field if non-nil, zero value otherwise.

### GetEnableArchiveMediaFilesOk

`func (o *ConfigurationLibraryOptions) GetEnableArchiveMediaFilesOk() (*bool, bool)`

GetEnableArchiveMediaFilesOk returns a tuple with the EnableArchiveMediaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveMediaFiles

`func (o *ConfigurationLibraryOptions) SetEnableArchiveMediaFiles(v bool)`

SetEnableArchiveMediaFiles sets EnableArchiveMediaFiles field to given value.

### HasEnableArchiveMediaFiles

`func (o *ConfigurationLibraryOptions) HasEnableArchiveMediaFiles() bool`

HasEnableArchiveMediaFiles returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *ConfigurationLibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *ConfigurationLibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *ConfigurationLibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *ConfigurationLibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *ConfigurationLibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *ConfigurationLibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *ConfigurationLibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *ConfigurationLibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *ConfigurationLibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *ConfigurationLibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *ConfigurationLibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *ConfigurationLibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *ConfigurationLibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *ConfigurationLibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *ConfigurationLibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *ConfigurationLibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetDownloadImagesInAdvance

`func (o *ConfigurationLibraryOptions) GetDownloadImagesInAdvance() bool`

GetDownloadImagesInAdvance returns the DownloadImagesInAdvance field if non-nil, zero value otherwise.

### GetDownloadImagesInAdvanceOk

`func (o *ConfigurationLibraryOptions) GetDownloadImagesInAdvanceOk() (*bool, bool)`

GetDownloadImagesInAdvanceOk returns a tuple with the DownloadImagesInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadImagesInAdvance

`func (o *ConfigurationLibraryOptions) SetDownloadImagesInAdvance(v bool)`

SetDownloadImagesInAdvance sets DownloadImagesInAdvance field to given value.

### HasDownloadImagesInAdvance

`func (o *ConfigurationLibraryOptions) HasDownloadImagesInAdvance() bool`

HasDownloadImagesInAdvance returns a boolean if a field has been set.

### GetPathInfos

`func (o *ConfigurationLibraryOptions) GetPathInfos() []ConfigurationMediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *ConfigurationLibraryOptions) GetPathInfosOk() (*[]ConfigurationMediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *ConfigurationLibraryOptions) SetPathInfos(v []ConfigurationMediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *ConfigurationLibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *ConfigurationLibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *ConfigurationLibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *ConfigurationLibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *ConfigurationLibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetSaveLocalThumbnailSets

`func (o *ConfigurationLibraryOptions) GetSaveLocalThumbnailSets() bool`

GetSaveLocalThumbnailSets returns the SaveLocalThumbnailSets field if non-nil, zero value otherwise.

### GetSaveLocalThumbnailSetsOk

`func (o *ConfigurationLibraryOptions) GetSaveLocalThumbnailSetsOk() (*bool, bool)`

GetSaveLocalThumbnailSetsOk returns a tuple with the SaveLocalThumbnailSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalThumbnailSets

`func (o *ConfigurationLibraryOptions) SetSaveLocalThumbnailSets(v bool)`

SetSaveLocalThumbnailSets sets SaveLocalThumbnailSets field to given value.

### HasSaveLocalThumbnailSets

`func (o *ConfigurationLibraryOptions) HasSaveLocalThumbnailSets() bool`

HasSaveLocalThumbnailSets returns a boolean if a field has been set.

### GetImportMissingEpisodes

`func (o *ConfigurationLibraryOptions) GetImportMissingEpisodes() bool`

GetImportMissingEpisodes returns the ImportMissingEpisodes field if non-nil, zero value otherwise.

### GetImportMissingEpisodesOk

`func (o *ConfigurationLibraryOptions) GetImportMissingEpisodesOk() (*bool, bool)`

GetImportMissingEpisodesOk returns a tuple with the ImportMissingEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMissingEpisodes

`func (o *ConfigurationLibraryOptions) SetImportMissingEpisodes(v bool)`

SetImportMissingEpisodes sets ImportMissingEpisodes field to given value.

### HasImportMissingEpisodes

`func (o *ConfigurationLibraryOptions) HasImportMissingEpisodes() bool`

HasImportMissingEpisodes returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *ConfigurationLibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *ConfigurationLibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *ConfigurationLibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *ConfigurationLibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *ConfigurationLibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *ConfigurationLibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *ConfigurationLibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *ConfigurationLibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableAudioResume

`func (o *ConfigurationLibraryOptions) GetEnableAudioResume() bool`

GetEnableAudioResume returns the EnableAudioResume field if non-nil, zero value otherwise.

### GetEnableAudioResumeOk

`func (o *ConfigurationLibraryOptions) GetEnableAudioResumeOk() (*bool, bool)`

GetEnableAudioResumeOk returns a tuple with the EnableAudioResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioResume

`func (o *ConfigurationLibraryOptions) SetEnableAudioResume(v bool)`

SetEnableAudioResume sets EnableAudioResume field to given value.

### HasEnableAudioResume

`func (o *ConfigurationLibraryOptions) HasEnableAudioResume() bool`

HasEnableAudioResume returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *ConfigurationLibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *ConfigurationLibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *ConfigurationLibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *ConfigurationLibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ConfigurationLibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ConfigurationLibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ConfigurationLibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ConfigurationLibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetContentType

`func (o *ConfigurationLibraryOptions) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ConfigurationLibraryOptions) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ConfigurationLibraryOptions) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ConfigurationLibraryOptions) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ConfigurationLibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ConfigurationLibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ConfigurationLibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ConfigurationLibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSeasonZeroDisplayName

`func (o *ConfigurationLibraryOptions) GetSeasonZeroDisplayName() string`

GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field if non-nil, zero value otherwise.

### GetSeasonZeroDisplayNameOk

`func (o *ConfigurationLibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool)`

GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonZeroDisplayName

`func (o *ConfigurationLibraryOptions) SetSeasonZeroDisplayName(v string)`

SetSeasonZeroDisplayName sets SeasonZeroDisplayName field to given value.

### HasSeasonZeroDisplayName

`func (o *ConfigurationLibraryOptions) HasSeasonZeroDisplayName() bool`

HasSeasonZeroDisplayName returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *ConfigurationLibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *ConfigurationLibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *ConfigurationLibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *ConfigurationLibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### GetDisabledLocalMetadataReaders

`func (o *ConfigurationLibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *ConfigurationLibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *ConfigurationLibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *ConfigurationLibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *ConfigurationLibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *ConfigurationLibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *ConfigurationLibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *ConfigurationLibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### GetDisabledSubtitleFetchers

`func (o *ConfigurationLibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *ConfigurationLibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *ConfigurationLibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *ConfigurationLibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *ConfigurationLibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *ConfigurationLibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *ConfigurationLibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *ConfigurationLibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ConfigurationLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *ConfigurationLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ConfigurationLibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ConfigurationLibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *ConfigurationLibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *ConfigurationLibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *ConfigurationLibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *ConfigurationLibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *ConfigurationLibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *ConfigurationLibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *ConfigurationLibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *ConfigurationLibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### GetRequirePerfectSubtitleMatch

`func (o *ConfigurationLibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *ConfigurationLibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *ConfigurationLibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *ConfigurationLibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *ConfigurationLibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *ConfigurationLibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *ConfigurationLibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *ConfigurationLibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetForcedSubtitlesOnly

`func (o *ConfigurationLibraryOptions) GetForcedSubtitlesOnly() bool`

GetForcedSubtitlesOnly returns the ForcedSubtitlesOnly field if non-nil, zero value otherwise.

### GetForcedSubtitlesOnlyOk

`func (o *ConfigurationLibraryOptions) GetForcedSubtitlesOnlyOk() (*bool, bool)`

GetForcedSubtitlesOnlyOk returns a tuple with the ForcedSubtitlesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSubtitlesOnly

`func (o *ConfigurationLibraryOptions) SetForcedSubtitlesOnly(v bool)`

SetForcedSubtitlesOnly sets ForcedSubtitlesOnly field to given value.

### HasForcedSubtitlesOnly

`func (o *ConfigurationLibraryOptions) HasForcedSubtitlesOnly() bool`

HasForcedSubtitlesOnly returns a boolean if a field has been set.

### GetTypeOptions

`func (o *ConfigurationLibraryOptions) GetTypeOptions() []ConfigurationTypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *ConfigurationLibraryOptions) GetTypeOptionsOk() (*[]ConfigurationTypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *ConfigurationLibraryOptions) SetTypeOptions(v []ConfigurationTypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *ConfigurationLibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.

### GetCollapseSingleItemFolders

`func (o *ConfigurationLibraryOptions) GetCollapseSingleItemFolders() bool`

GetCollapseSingleItemFolders returns the CollapseSingleItemFolders field if non-nil, zero value otherwise.

### GetCollapseSingleItemFoldersOk

`func (o *ConfigurationLibraryOptions) GetCollapseSingleItemFoldersOk() (*bool, bool)`

GetCollapseSingleItemFoldersOk returns a tuple with the CollapseSingleItemFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseSingleItemFolders

`func (o *ConfigurationLibraryOptions) SetCollapseSingleItemFolders(v bool)`

SetCollapseSingleItemFolders sets CollapseSingleItemFolders field to given value.

### HasCollapseSingleItemFolders

`func (o *ConfigurationLibraryOptions) HasCollapseSingleItemFolders() bool`

HasCollapseSingleItemFolders returns a boolean if a field has been set.

### GetMinResumePct

`func (o *ConfigurationLibraryOptions) GetMinResumePct() int32`

GetMinResumePct returns the MinResumePct field if non-nil, zero value otherwise.

### GetMinResumePctOk

`func (o *ConfigurationLibraryOptions) GetMinResumePctOk() (*int32, bool)`

GetMinResumePctOk returns a tuple with the MinResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumePct

`func (o *ConfigurationLibraryOptions) SetMinResumePct(v int32)`

SetMinResumePct sets MinResumePct field to given value.

### HasMinResumePct

`func (o *ConfigurationLibraryOptions) HasMinResumePct() bool`

HasMinResumePct returns a boolean if a field has been set.

### GetMaxResumePct

`func (o *ConfigurationLibraryOptions) GetMaxResumePct() int32`

GetMaxResumePct returns the MaxResumePct field if non-nil, zero value otherwise.

### GetMaxResumePctOk

`func (o *ConfigurationLibraryOptions) GetMaxResumePctOk() (*int32, bool)`

GetMaxResumePctOk returns a tuple with the MaxResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResumePct

`func (o *ConfigurationLibraryOptions) SetMaxResumePct(v int32)`

SetMaxResumePct sets MaxResumePct field to given value.

### HasMaxResumePct

`func (o *ConfigurationLibraryOptions) HasMaxResumePct() bool`

HasMaxResumePct returns a boolean if a field has been set.

### GetMinResumeDurationSeconds

`func (o *ConfigurationLibraryOptions) GetMinResumeDurationSeconds() int32`

GetMinResumeDurationSeconds returns the MinResumeDurationSeconds field if non-nil, zero value otherwise.

### GetMinResumeDurationSecondsOk

`func (o *ConfigurationLibraryOptions) GetMinResumeDurationSecondsOk() (*int32, bool)`

GetMinResumeDurationSecondsOk returns a tuple with the MinResumeDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumeDurationSeconds

`func (o *ConfigurationLibraryOptions) SetMinResumeDurationSeconds(v int32)`

SetMinResumeDurationSeconds sets MinResumeDurationSeconds field to given value.

### HasMinResumeDurationSeconds

`func (o *ConfigurationLibraryOptions) HasMinResumeDurationSeconds() bool`

HasMinResumeDurationSeconds returns a boolean if a field has been set.

### GetThumbnailImagesIntervalSeconds

`func (o *ConfigurationLibraryOptions) GetThumbnailImagesIntervalSeconds() int32`

GetThumbnailImagesIntervalSeconds returns the ThumbnailImagesIntervalSeconds field if non-nil, zero value otherwise.

### GetThumbnailImagesIntervalSecondsOk

`func (o *ConfigurationLibraryOptions) GetThumbnailImagesIntervalSecondsOk() (*int32, bool)`

GetThumbnailImagesIntervalSecondsOk returns a tuple with the ThumbnailImagesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImagesIntervalSeconds

`func (o *ConfigurationLibraryOptions) SetThumbnailImagesIntervalSeconds(v int32)`

SetThumbnailImagesIntervalSeconds sets ThumbnailImagesIntervalSeconds field to given value.

### HasThumbnailImagesIntervalSeconds

`func (o *ConfigurationLibraryOptions) HasThumbnailImagesIntervalSeconds() bool`

HasThumbnailImagesIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



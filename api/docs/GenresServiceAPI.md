# \GenresServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenres**](GenresServiceAPI.md#GetGenres) | **Get** /Genres | Gets all genres from a given item, folder, or the entire library
[**GetGenresByName**](GenresServiceAPI.md#GetGenresByName) | **Get** /Genres/{Name} | Gets a genre, by name



## GetGenres

> QueryResultBaseItemDto GetGenres(ctx).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).VideoCodecs(videoCodecs).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets all genres from a given item, folder, or the entire library



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Akimio521/emby-go/api"
)

func main() {
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	locationTypes := "locationTypes_example" // string | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted. (optional)
	excludeLocationTypes := "excludeLocationTypes_example" // string | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted. (optional)
	isMissing := true // bool | Optional filter by items that are missing episodes or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for movies. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresServiceAPI.GetGenres(context.Background()).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).VideoCodecs(videoCodecs).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresServiceAPI.GetGenres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenres`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `GenresServiceAPI.GetGenres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGenresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **locationTypes** | **string** | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted. | 
 **excludeLocationTypes** | **string** | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted. | 
 **isMissing** | **bool** | Optional filter by items that are missing episodes or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for movies. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenresByName

> BaseItemDto GetGenresByName(ctx, name).UserId(userId).Execute()

Gets a genre, by name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Akimio521/emby-go/api"
)

func main() {
	name := "name_example" // string | The genre name
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresServiceAPI.GetGenresByName(context.Background(), name).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresServiceAPI.GetGenresByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenresByName`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `GenresServiceAPI.GetGenresByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The genre name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenresByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


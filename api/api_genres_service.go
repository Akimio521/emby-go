/*
Emby Server API

Explore the Emby Server API

API version: 4.1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type GenresServiceAPI interface {

	/*
	GetGenres Gets all genres from a given item, folder, or the entire library

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetGenresRequest
	*/
	GetGenres(ctx context.Context) ApiGetGenresRequest

	// GetGenresExecute executes the request
	//  @return QueryResultBaseItemDto
	GetGenresExecute(r ApiGetGenresRequest) (*QueryResultBaseItemDto, *http.Response, error)

	/*
	GetGenresByName Gets a genre, by name

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The genre name
	@return ApiGetGenresByNameRequest
	*/
	GetGenresByName(ctx context.Context, name string) ApiGetGenresByNameRequest

	// GetGenresByNameExecute executes the request
	//  @return BaseItemDto
	GetGenresByNameExecute(r ApiGetGenresByNameRequest) (*BaseItemDto, *http.Response, error)
}

// GenresServiceAPIService GenresServiceAPI service
type GenresServiceAPIService service

type ApiGetGenresRequest struct {
	ctx context.Context
	ApiService GenresServiceAPI
	artistType *string
	maxOfficialRating *string
	hasThemeSong *bool
	hasThemeVideo *bool
	hasSubtitles *bool
	hasSpecialFeature *bool
	hasTrailer *bool
	adjacentTo *string
	minIndexNumber *int32
	minPlayers *int32
	maxPlayers *int32
	parentIndexNumber *int32
	hasParentalRating *bool
	isHD *bool
	locationTypes *string
	excludeLocationTypes *string
	isMissing *bool
	isUnaired *bool
	minCommunityRating *float64
	minCriticRating *float64
	airedDuringSeason *int32
	minPremiereDate *string
	minDateLastSaved *string
	minDateLastSavedForUser *string
	maxPremiereDate *string
	hasOverview *bool
	hasImdbId *bool
	hasTmdbId *bool
	hasTvdbId *bool
	excludeItemIds *string
	startIndex *int32
	limit *int32
	recursive *bool
	sortOrder *string
	parentId *string
	fields *string
	excludeItemTypes *string
	includeItemTypes *string
	anyProviderIdEquals *string
	filters *string
	isFavorite *bool
	isMovie *bool
	isSeries *bool
	isNews *bool
	isKids *bool
	isSports *bool
	mediaTypes *string
	imageTypes *string
	sortBy *string
	isPlayed *bool
	genres *string
	officialRatings *string
	tags *string
	years *string
	enableImages *bool
	enableUserData *bool
	imageTypeLimit *int32
	enableImageTypes *string
	person *string
	personIds *string
	personTypes *string
	studios *string
	studioIds *string
	artists *string
	artistIds *string
	albums *string
	ids *string
	videoTypes *string
	containers *string
	audioCodecs *string
	videoCodecs *string
	subtitleCodecs *string
	path *string
	userId *string
	minOfficialRating *string
	isLocked *bool
	isPlaceHolder *bool
	hasOfficialRating *bool
	groupItemsIntoCollections *bool
	is3D *bool
	seriesStatus *string
	nameStartsWithOrGreater *string
	nameStartsWith *string
	nameLessThan *string
}

// Artist or AlbumArtist
func (r ApiGetGenresRequest) ArtistType(artistType string) ApiGetGenresRequest {
	r.artistType = &artistType
	return r
}

// Optional filter by maximum official rating (PG, PG-13, TV-MA, etc).
func (r ApiGetGenresRequest) MaxOfficialRating(maxOfficialRating string) ApiGetGenresRequest {
	r.maxOfficialRating = &maxOfficialRating
	return r
}

// Optional filter by items with theme songs.
func (r ApiGetGenresRequest) HasThemeSong(hasThemeSong bool) ApiGetGenresRequest {
	r.hasThemeSong = &hasThemeSong
	return r
}

// Optional filter by items with theme videos.
func (r ApiGetGenresRequest) HasThemeVideo(hasThemeVideo bool) ApiGetGenresRequest {
	r.hasThemeVideo = &hasThemeVideo
	return r
}

// Optional filter by items with subtitles.
func (r ApiGetGenresRequest) HasSubtitles(hasSubtitles bool) ApiGetGenresRequest {
	r.hasSubtitles = &hasSubtitles
	return r
}

// Optional filter by items with special features.
func (r ApiGetGenresRequest) HasSpecialFeature(hasSpecialFeature bool) ApiGetGenresRequest {
	r.hasSpecialFeature = &hasSpecialFeature
	return r
}

// Optional filter by items with trailers.
func (r ApiGetGenresRequest) HasTrailer(hasTrailer bool) ApiGetGenresRequest {
	r.hasTrailer = &hasTrailer
	return r
}

// Optional. Return items that are siblings of a supplied item.
func (r ApiGetGenresRequest) AdjacentTo(adjacentTo string) ApiGetGenresRequest {
	r.adjacentTo = &adjacentTo
	return r
}

// Optional filter by minimum index number.
func (r ApiGetGenresRequest) MinIndexNumber(minIndexNumber int32) ApiGetGenresRequest {
	r.minIndexNumber = &minIndexNumber
	return r
}

// Optional filter by minimum number of game players.
func (r ApiGetGenresRequest) MinPlayers(minPlayers int32) ApiGetGenresRequest {
	r.minPlayers = &minPlayers
	return r
}

// Optional filter by maximum number of game players.
func (r ApiGetGenresRequest) MaxPlayers(maxPlayers int32) ApiGetGenresRequest {
	r.maxPlayers = &maxPlayers
	return r
}

// Optional filter by parent index number.
func (r ApiGetGenresRequest) ParentIndexNumber(parentIndexNumber int32) ApiGetGenresRequest {
	r.parentIndexNumber = &parentIndexNumber
	return r
}

// Optional filter by items that have or do not have a parental rating
func (r ApiGetGenresRequest) HasParentalRating(hasParentalRating bool) ApiGetGenresRequest {
	r.hasParentalRating = &hasParentalRating
	return r
}

// Optional filter by items that are HD or not.
func (r ApiGetGenresRequest) IsHD(isHD bool) ApiGetGenresRequest {
	r.isHD = &isHD
	return r
}

// Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) LocationTypes(locationTypes string) ApiGetGenresRequest {
	r.locationTypes = &locationTypes
	return r
}

// Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) ExcludeLocationTypes(excludeLocationTypes string) ApiGetGenresRequest {
	r.excludeLocationTypes = &excludeLocationTypes
	return r
}

// Optional filter by items that are missing episodes or not.
func (r ApiGetGenresRequest) IsMissing(isMissing bool) ApiGetGenresRequest {
	r.isMissing = &isMissing
	return r
}

// Optional filter by items that are unaired episodes or not.
func (r ApiGetGenresRequest) IsUnaired(isUnaired bool) ApiGetGenresRequest {
	r.isUnaired = &isUnaired
	return r
}

// Optional filter by minimum community rating.
func (r ApiGetGenresRequest) MinCommunityRating(minCommunityRating float64) ApiGetGenresRequest {
	r.minCommunityRating = &minCommunityRating
	return r
}

// Optional filter by minimum critic rating.
func (r ApiGetGenresRequest) MinCriticRating(minCriticRating float64) ApiGetGenresRequest {
	r.minCriticRating = &minCriticRating
	return r
}

// Gets all episodes that aired during a season, including specials.
func (r ApiGetGenresRequest) AiredDuringSeason(airedDuringSeason int32) ApiGetGenresRequest {
	r.airedDuringSeason = &airedDuringSeason
	return r
}

// Optional. The minimum premiere date. Format &#x3D; ISO
func (r ApiGetGenresRequest) MinPremiereDate(minPremiereDate string) ApiGetGenresRequest {
	r.minPremiereDate = &minPremiereDate
	return r
}

// Optional. The minimum premiere date. Format &#x3D; ISO
func (r ApiGetGenresRequest) MinDateLastSaved(minDateLastSaved string) ApiGetGenresRequest {
	r.minDateLastSaved = &minDateLastSaved
	return r
}

// Optional. The minimum premiere date. Format &#x3D; ISO
func (r ApiGetGenresRequest) MinDateLastSavedForUser(minDateLastSavedForUser string) ApiGetGenresRequest {
	r.minDateLastSavedForUser = &minDateLastSavedForUser
	return r
}

// Optional. The maximum premiere date. Format &#x3D; ISO
func (r ApiGetGenresRequest) MaxPremiereDate(maxPremiereDate string) ApiGetGenresRequest {
	r.maxPremiereDate = &maxPremiereDate
	return r
}

// Optional filter by items that have an overview or not.
func (r ApiGetGenresRequest) HasOverview(hasOverview bool) ApiGetGenresRequest {
	r.hasOverview = &hasOverview
	return r
}

// Optional filter by items that have an imdb id or not.
func (r ApiGetGenresRequest) HasImdbId(hasImdbId bool) ApiGetGenresRequest {
	r.hasImdbId = &hasImdbId
	return r
}

// Optional filter by items that have a tmdb id or not.
func (r ApiGetGenresRequest) HasTmdbId(hasTmdbId bool) ApiGetGenresRequest {
	r.hasTmdbId = &hasTmdbId
	return r
}

// Optional filter by items that have a tvdb id or not.
func (r ApiGetGenresRequest) HasTvdbId(hasTvdbId bool) ApiGetGenresRequest {
	r.hasTvdbId = &hasTvdbId
	return r
}

// Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) ExcludeItemIds(excludeItemIds string) ApiGetGenresRequest {
	r.excludeItemIds = &excludeItemIds
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetGenresRequest) StartIndex(startIndex int32) ApiGetGenresRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return
func (r ApiGetGenresRequest) Limit(limit int32) ApiGetGenresRequest {
	r.limit = &limit
	return r
}

// When searching within folders, this determines whether or not the search will be recursive. true/false
func (r ApiGetGenresRequest) Recursive(recursive bool) ApiGetGenresRequest {
	r.recursive = &recursive
	return r
}

// Sort Order - Ascending,Descending
func (r ApiGetGenresRequest) SortOrder(sortOrder string) ApiGetGenresRequest {
	r.sortOrder = &sortOrder
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root
func (r ApiGetGenresRequest) ParentId(parentId string) ApiGetGenresRequest {
	r.parentId = &parentId
	return r
}

// Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines
func (r ApiGetGenresRequest) Fields(fields string) ApiGetGenresRequest {
	r.fields = &fields
	return r
}

// Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) ExcludeItemTypes(excludeItemTypes string) ApiGetGenresRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) IncludeItemTypes(includeItemTypes string) ApiGetGenresRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs.
func (r ApiGetGenresRequest) AnyProviderIdEquals(anyProviderIdEquals string) ApiGetGenresRequest {
	r.anyProviderIdEquals = &anyProviderIdEquals
	return r
}

// Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes
func (r ApiGetGenresRequest) Filters(filters string) ApiGetGenresRequest {
	r.filters = &filters
	return r
}

// Optional filter by items that are marked as favorite, or not.
func (r ApiGetGenresRequest) IsFavorite(isFavorite bool) ApiGetGenresRequest {
	r.isFavorite = &isFavorite
	return r
}

// Optional filter for movies.
func (r ApiGetGenresRequest) IsMovie(isMovie bool) ApiGetGenresRequest {
	r.isMovie = &isMovie
	return r
}

// Optional filter for movies.
func (r ApiGetGenresRequest) IsSeries(isSeries bool) ApiGetGenresRequest {
	r.isSeries = &isSeries
	return r
}

// Optional filter for news.
func (r ApiGetGenresRequest) IsNews(isNews bool) ApiGetGenresRequest {
	r.isNews = &isNews
	return r
}

// Optional filter for kids.
func (r ApiGetGenresRequest) IsKids(isKids bool) ApiGetGenresRequest {
	r.isKids = &isKids
	return r
}

// Optional filter for sports.
func (r ApiGetGenresRequest) IsSports(isSports bool) ApiGetGenresRequest {
	r.isSports = &isSports
	return r
}

// Optional filter by MediaType. Allows multiple, comma delimited.
func (r ApiGetGenresRequest) MediaTypes(mediaTypes string) ApiGetGenresRequest {
	r.mediaTypes = &mediaTypes
	return r
}

// Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited.
func (r ApiGetGenresRequest) ImageTypes(imageTypes string) ApiGetGenresRequest {
	r.imageTypes = &imageTypes
	return r
}

// Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime
func (r ApiGetGenresRequest) SortBy(sortBy string) ApiGetGenresRequest {
	r.sortBy = &sortBy
	return r
}

// Optional filter by items that are played, or not.
func (r ApiGetGenresRequest) IsPlayed(isPlayed bool) ApiGetGenresRequest {
	r.isPlayed = &isPlayed
	return r
}

// Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) Genres(genres string) ApiGetGenresRequest {
	r.genres = &genres
	return r
}

// Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) OfficialRatings(officialRatings string) ApiGetGenresRequest {
	r.officialRatings = &officialRatings
	return r
}

// Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) Tags(tags string) ApiGetGenresRequest {
	r.tags = &tags
	return r
}

// Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted.
func (r ApiGetGenresRequest) Years(years string) ApiGetGenresRequest {
	r.years = &years
	return r
}

// Optional, include image information in output
func (r ApiGetGenresRequest) EnableImages(enableImages bool) ApiGetGenresRequest {
	r.enableImages = &enableImages
	return r
}

// Optional, include user data
func (r ApiGetGenresRequest) EnableUserData(enableUserData bool) ApiGetGenresRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional, the max number of images to return, per image type
func (r ApiGetGenresRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetGenresRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetGenresRequest) EnableImageTypes(enableImageTypes string) ApiGetGenresRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. If specified, results will be filtered to include only those containing the specified person.
func (r ApiGetGenresRequest) Person(person string) ApiGetGenresRequest {
	r.person = &person
	return r
}

// Optional. If specified, results will be filtered to include only those containing the specified person.
func (r ApiGetGenresRequest) PersonIds(personIds string) ApiGetGenresRequest {
	r.personIds = &personIds
	return r
}

// Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited
func (r ApiGetGenresRequest) PersonTypes(personTypes string) ApiGetGenresRequest {
	r.personTypes = &personTypes
	return r
}

// Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) Studios(studios string) ApiGetGenresRequest {
	r.studios = &studios
	return r
}

// Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) StudioIds(studioIds string) ApiGetGenresRequest {
	r.studioIds = &studioIds
	return r
}

// Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) Artists(artists string) ApiGetGenresRequest {
	r.artists = &artists
	return r
}

// Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) ArtistIds(artistIds string) ApiGetGenresRequest {
	r.artistIds = &artistIds
	return r
}

// Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted.
func (r ApiGetGenresRequest) Albums(albums string) ApiGetGenresRequest {
	r.albums = &albums
	return r
}

// Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited.
func (r ApiGetGenresRequest) Ids(ids string) ApiGetGenresRequest {
	r.ids = &ids
	return r
}

// Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) VideoTypes(videoTypes string) ApiGetGenresRequest {
	r.videoTypes = &videoTypes
	return r
}

// Optional filter by Container. Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) Containers(containers string) ApiGetGenresRequest {
	r.containers = &containers
	return r
}

// Optional filter by AudioCodec. Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) AudioCodecs(audioCodecs string) ApiGetGenresRequest {
	r.audioCodecs = &audioCodecs
	return r
}

// Optional filter by VideoCodec. Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) VideoCodecs(videoCodecs string) ApiGetGenresRequest {
	r.videoCodecs = &videoCodecs
	return r
}

// Optional filter by SubtitleCodec. Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) SubtitleCodecs(subtitleCodecs string) ApiGetGenresRequest {
	r.subtitleCodecs = &subtitleCodecs
	return r
}

// Optional filter by Path.
func (r ApiGetGenresRequest) Path(path string) ApiGetGenresRequest {
	r.path = &path
	return r
}

// User Id
func (r ApiGetGenresRequest) UserId(userId string) ApiGetGenresRequest {
	r.userId = &userId
	return r
}

// Optional filter by minimum official rating (PG, PG-13, TV-MA, etc).
func (r ApiGetGenresRequest) MinOfficialRating(minOfficialRating string) ApiGetGenresRequest {
	r.minOfficialRating = &minOfficialRating
	return r
}

// Optional filter by items that are locked.
func (r ApiGetGenresRequest) IsLocked(isLocked bool) ApiGetGenresRequest {
	r.isLocked = &isLocked
	return r
}

// Optional filter by items that are placeholders
func (r ApiGetGenresRequest) IsPlaceHolder(isPlaceHolder bool) ApiGetGenresRequest {
	r.isPlaceHolder = &isPlaceHolder
	return r
}

// Optional filter by items that have official ratings
func (r ApiGetGenresRequest) HasOfficialRating(hasOfficialRating bool) ApiGetGenresRequest {
	r.hasOfficialRating = &hasOfficialRating
	return r
}

// Whether or not to hide items behind their boxsets.
func (r ApiGetGenresRequest) GroupItemsIntoCollections(groupItemsIntoCollections bool) ApiGetGenresRequest {
	r.groupItemsIntoCollections = &groupItemsIntoCollections
	return r
}

// Optional filter by items that are 3D, or not.
func (r ApiGetGenresRequest) Is3D(is3D bool) ApiGetGenresRequest {
	r.is3D = &is3D
	return r
}

// Optional filter by Series Status. Allows multiple, comma delimeted.
func (r ApiGetGenresRequest) SeriesStatus(seriesStatus string) ApiGetGenresRequest {
	r.seriesStatus = &seriesStatus
	return r
}

// Optional filter by items whose name is sorted equally or greater than a given input string.
func (r ApiGetGenresRequest) NameStartsWithOrGreater(nameStartsWithOrGreater string) ApiGetGenresRequest {
	r.nameStartsWithOrGreater = &nameStartsWithOrGreater
	return r
}

// Optional filter by items whose name is sorted equally than a given input string.
func (r ApiGetGenresRequest) NameStartsWith(nameStartsWith string) ApiGetGenresRequest {
	r.nameStartsWith = &nameStartsWith
	return r
}

// Optional filter by items whose name is equally or lesser than a given input string.
func (r ApiGetGenresRequest) NameLessThan(nameLessThan string) ApiGetGenresRequest {
	r.nameLessThan = &nameLessThan
	return r
}

func (r ApiGetGenresRequest) Execute() (*QueryResultBaseItemDto, *http.Response, error) {
	return r.ApiService.GetGenresExecute(r)
}

/*
GetGenres Gets all genres from a given item, folder, or the entire library

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGenresRequest
*/
func (a *GenresServiceAPIService) GetGenres(ctx context.Context) ApiGetGenresRequest {
	return ApiGetGenresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryResultBaseItemDto
func (a *GenresServiceAPIService) GetGenresExecute(r ApiGetGenresRequest) (*QueryResultBaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryResultBaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenresServiceAPIService.GetGenres")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Genres"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.artistType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ArtistType", r.artistType, "form", "")
	}
	if r.maxOfficialRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxOfficialRating", r.maxOfficialRating, "form", "")
	}
	if r.hasThemeSong != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasThemeSong", r.hasThemeSong, "form", "")
	}
	if r.hasThemeVideo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasThemeVideo", r.hasThemeVideo, "form", "")
	}
	if r.hasSubtitles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasSubtitles", r.hasSubtitles, "form", "")
	}
	if r.hasSpecialFeature != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasSpecialFeature", r.hasSpecialFeature, "form", "")
	}
	if r.hasTrailer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasTrailer", r.hasTrailer, "form", "")
	}
	if r.adjacentTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AdjacentTo", r.adjacentTo, "form", "")
	}
	if r.minIndexNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinIndexNumber", r.minIndexNumber, "form", "")
	}
	if r.minPlayers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinPlayers", r.minPlayers, "form", "")
	}
	if r.maxPlayers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxPlayers", r.maxPlayers, "form", "")
	}
	if r.parentIndexNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ParentIndexNumber", r.parentIndexNumber, "form", "")
	}
	if r.hasParentalRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasParentalRating", r.hasParentalRating, "form", "")
	}
	if r.isHD != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsHD", r.isHD, "form", "")
	}
	if r.locationTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "LocationTypes", r.locationTypes, "form", "")
	}
	if r.excludeLocationTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ExcludeLocationTypes", r.excludeLocationTypes, "form", "")
	}
	if r.isMissing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsMissing", r.isMissing, "form", "")
	}
	if r.isUnaired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsUnaired", r.isUnaired, "form", "")
	}
	if r.minCommunityRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinCommunityRating", r.minCommunityRating, "form", "")
	}
	if r.minCriticRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinCriticRating", r.minCriticRating, "form", "")
	}
	if r.airedDuringSeason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AiredDuringSeason", r.airedDuringSeason, "form", "")
	}
	if r.minPremiereDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinPremiereDate", r.minPremiereDate, "form", "")
	}
	if r.minDateLastSaved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinDateLastSaved", r.minDateLastSaved, "form", "")
	}
	if r.minDateLastSavedForUser != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinDateLastSavedForUser", r.minDateLastSavedForUser, "form", "")
	}
	if r.maxPremiereDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxPremiereDate", r.maxPremiereDate, "form", "")
	}
	if r.hasOverview != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasOverview", r.hasOverview, "form", "")
	}
	if r.hasImdbId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasImdbId", r.hasImdbId, "form", "")
	}
	if r.hasTmdbId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasTmdbId", r.hasTmdbId, "form", "")
	}
	if r.hasTvdbId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasTvdbId", r.hasTvdbId, "form", "")
	}
	if r.excludeItemIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ExcludeItemIds", r.excludeItemIds, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Limit", r.limit, "form", "")
	}
	if r.recursive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Recursive", r.recursive, "form", "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SortOrder", r.sortOrder, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ParentId", r.parentId, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Fields", r.fields, "form", "")
	}
	if r.excludeItemTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ExcludeItemTypes", r.excludeItemTypes, "form", "")
	}
	if r.includeItemTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IncludeItemTypes", r.includeItemTypes, "form", "")
	}
	if r.anyProviderIdEquals != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AnyProviderIdEquals", r.anyProviderIdEquals, "form", "")
	}
	if r.filters != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Filters", r.filters, "form", "")
	}
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsFavorite", r.isFavorite, "form", "")
	}
	if r.isMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsMovie", r.isMovie, "form", "")
	}
	if r.isSeries != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsSeries", r.isSeries, "form", "")
	}
	if r.isNews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsNews", r.isNews, "form", "")
	}
	if r.isKids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsKids", r.isKids, "form", "")
	}
	if r.isSports != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsSports", r.isSports, "form", "")
	}
	if r.mediaTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MediaTypes", r.mediaTypes, "form", "")
	}
	if r.imageTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ImageTypes", r.imageTypes, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SortBy", r.sortBy, "form", "")
	}
	if r.isPlayed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsPlayed", r.isPlayed, "form", "")
	}
	if r.genres != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Genres", r.genres, "form", "")
	}
	if r.officialRatings != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OfficialRatings", r.officialRatings, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Tags", r.tags, "form", "")
	}
	if r.years != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Years", r.years, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EnableImages", r.enableImages, "form", "")
	}
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EnableUserData", r.enableUserData, "form", "")
	}
	if r.imageTypeLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ImageTypeLimit", r.imageTypeLimit, "form", "")
	}
	if r.enableImageTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EnableImageTypes", r.enableImageTypes, "form", "")
	}
	if r.person != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Person", r.person, "form", "")
	}
	if r.personIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PersonIds", r.personIds, "form", "")
	}
	if r.personTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PersonTypes", r.personTypes, "form", "")
	}
	if r.studios != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Studios", r.studios, "form", "")
	}
	if r.studioIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StudioIds", r.studioIds, "form", "")
	}
	if r.artists != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Artists", r.artists, "form", "")
	}
	if r.artistIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ArtistIds", r.artistIds, "form", "")
	}
	if r.albums != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Albums", r.albums, "form", "")
	}
	if r.ids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Ids", r.ids, "form", "")
	}
	if r.videoTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "VideoTypes", r.videoTypes, "form", "")
	}
	if r.containers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Containers", r.containers, "form", "")
	}
	if r.audioCodecs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AudioCodecs", r.audioCodecs, "form", "")
	}
	if r.videoCodecs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "VideoCodecs", r.videoCodecs, "form", "")
	}
	if r.subtitleCodecs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SubtitleCodecs", r.subtitleCodecs, "form", "")
	}
	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Path", r.path, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "UserId", r.userId, "form", "")
	}
	if r.minOfficialRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MinOfficialRating", r.minOfficialRating, "form", "")
	}
	if r.isLocked != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsLocked", r.isLocked, "form", "")
	}
	if r.isPlaceHolder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsPlaceHolder", r.isPlaceHolder, "form", "")
	}
	if r.hasOfficialRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "HasOfficialRating", r.hasOfficialRating, "form", "")
	}
	if r.groupItemsIntoCollections != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "GroupItemsIntoCollections", r.groupItemsIntoCollections, "form", "")
	}
	if r.is3D != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Is3D", r.is3D, "form", "")
	}
	if r.seriesStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SeriesStatus", r.seriesStatus, "form", "")
	}
	if r.nameStartsWithOrGreater != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NameStartsWithOrGreater", r.nameStartsWithOrGreater, "form", "")
	}
	if r.nameStartsWith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NameStartsWith", r.nameStartsWith, "form", "")
	}
	if r.nameLessThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NameLessThan", r.nameLessThan, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGenresByNameRequest struct {
	ctx context.Context
	ApiService GenresServiceAPI
	name string
	userId *string
}

// Optional. Filter by user id, and attach user data
func (r ApiGetGenresByNameRequest) UserId(userId string) ApiGetGenresByNameRequest {
	r.userId = &userId
	return r
}

func (r ApiGetGenresByNameRequest) Execute() (*BaseItemDto, *http.Response, error) {
	return r.ApiService.GetGenresByNameExecute(r)
}

/*
GetGenresByName Gets a genre, by name

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name The genre name
 @return ApiGetGenresByNameRequest
*/
func (a *GenresServiceAPIService) GetGenresByName(ctx context.Context, name string) ApiGetGenresByNameRequest {
	return ApiGetGenresByNameRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return BaseItemDto
func (a *GenresServiceAPIService) GetGenresByNameExecute(r ApiGetGenresByNameRequest) (*BaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenresServiceAPIService.GetGenresByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Genres/{Name}"
	localVarPath = strings.Replace(localVarPath, "{"+"Name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "UserId", r.userId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

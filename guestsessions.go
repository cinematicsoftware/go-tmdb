package tmdb

import (
	"fmt"
)

// GetGuestSessionRatedMovies gets the list of rated movies for a specific guest session id
// https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-movies
func (tmdb *TMDb) GetGuestSessionRatedMovies(sessionID string, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":       {},
		"sort_by":    {},
		"sort_order": {},
		"language":   {}}
	var favorites MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/guest_session/%v/rated/movies?api_key=%s%s", baseURL, sessionID, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*MoviePagedResults), err
}

// GetGuestSessionRatedTvShows gets the list of rated tv shows for a specific guest session id
// https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-tv-shows
func (tmdb *TMDb) GetGuestSessionRatedTvShows(sessionID string, options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":       {},
		"sort_by":    {},
		"sort_order": {},
		"language":   {}}
	var favorites TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/guest_session/%v/rated/tv?api_key=%s%s", baseURL, sessionID, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*TvPagedResults), err
}

// GetGuestSessionRatedTvEpisodes gets the list of rated tv episodes for a specific guest session id
// https://developers.themoviedb.org/3/guest-sessions/get-gest-session-rated-tv-episodes
func (tmdb *TMDb) GetGuestSessionRatedTvEpisodes(sessionID string, options map[string]string) (*TvEpisodePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":       {},
		"sort_by":    {},
		"sort_order": {},
		"language":   {}}
	var favorites TvEpisodePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/guest_session/%v/rated/tv/episodes?api_key=%s%s", baseURL, sessionID, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*TvEpisodePagedResults), err
}

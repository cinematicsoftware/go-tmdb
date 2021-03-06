package tmdb

import (
	"fmt"
)

// Movie struct
type Movie struct {
	Adult               bool
	BackdropPath        string           `json:"backdrop_path"`
	BelongsToCollection *CollectionShort `json:"belongs_to_collection,omitempty"`
	Budget              uint32
	Genres              []struct {
		ID   int
		Name string
	}
	Homepage            string
	ID                  int
	ImdbID              string `json:"imdb_id"`
	OriginalLanguage    string `json:"original_language"`
	OriginalTitle       string `json:"original_title"`
	Overview            string
	Popularity          float32
	PosterPath          string `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int
		Name string
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         uint32
	Runtime         uint32
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string
	} `json:"spoken_languages"`
	Status            string
	Tagline           string
	Title             string
	Video             bool
	VoteAverage       float32                 `json:"vote_average"`
	VoteCount         uint32                  `json:"vote_count"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieShort struct
type MovieShort struct {
	Adult         bool
	BackdropPath  string `json:"backdrop_path"`
	ID            int
	OriginalTitle string `json:"original_title"`
	Popularity    float32
	PosterPath    string `json:"poster_path"`
	ReleaseDate   string `json:"release_date"`
	Title         string
	Video         bool
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
}

// MovieDatedResults struct
type MovieDatedResults struct {
	Dates struct {
		Minimum string
		Maximum string
	}
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MoviePagedResults struct
type MoviePagedResults struct {
	ID                int
	Page              int
	Results           []MovieShort
	TotalPages        int                     `json:"total_pages"`
	TotalResults      int                     `json:"total_results"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieAccountState struct
type MovieAccountState struct {
	ID        int
	Favorite  bool
	Watchlist bool
	Rated     struct {
		Value float32
	}
}

// MovieAlternativeTitles struct
type MovieAlternativeTitles struct {
	ID     int
	Titles []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string
	}
	AccountStates *MovieAccountState `json:"account_states,omitempty"`
	Credits       *MovieCredits      `json:",omitempty"`
	Images        *MovieImages       `json:",omitempty"`
	Keywords      *MovieKeywords     `json:",omitempty"`
	Releases      *MovieReleases     `json:"release_dates,omitempty"`
	Videos        *MovieVideos       `json:",omitempty"`
	Translations  *MovieTranslations `json:",omitempty"`
	Similar       *MoviePagedResults `json:",omitempty"`
	Reviews       *MovieReviews      `json:",omitempty"`
	Lists         *MovieLists        `json:",omitempty"`
	Changes       *MovieChanges      `json:",omitempty"`
	// Rating        *MovieRating       `json:",omitempty"`
}

// MovieChanges struct
type MovieChanges struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// MovieCredits struct
type MovieCredits struct {
	ID   int
	Cast []struct {
		CastID      int `json:"cast_id"`
		Character   string
		CreditID    string `json:"credit_id"`
		ID          int
		Name        string
		Order       int
		ProfilePath string `json:"profile_path"`
	}
	Crew []struct {
		CreditID    string `json:"credit_id"`
		Department  string
		ID          int
		Job         string
		Name        string
		ProfilePath string `json:"profile_path"`
	}
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieImage struct
type MovieImage struct {
	FilePath    string `json:"file_path"`
	Width       int
	Height      int
	Iso639_1    string  `json:"iso_639_1"`
	AspectRatio float32 `json:"aspect_ratio"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   uint32  `json:"vote_count"`
}

// MovieImages struct
type MovieImages struct {
	ID                int
	Backdrops         []MovieImage
	Posters           []MovieImage
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieKeywords struct
type MovieKeywords struct {
	ID       int
	Keywords []struct {
		ID   int
		Name string
	}
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieLists struct
type MovieLists struct {
	ID      int
	Page    int
	Results []struct {
		Description   string
		FavoriteCount int `json:"favorite_count"`
		ID            string
		ItemCount     int    `json:"item_count"`
		Iso639_1      string `json:"iso_639_1"`
		Name          string
		PosterPath    string `json:"poster_path"`
	}
	TotalPages        int                     `json:"total_pages"`
	TotalResults      int                     `json:"total_results"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieRelease struct
type MovieRelease struct {
	Iso3166_1    string `json:"iso_3166_1"`
	ReleaseDates []struct {
		Certification string
		Iso639_1      string `json:"iso_639_1"`
		ReleaseDate   string `json:"release_date"`
		Type          int
	} `json:"release_dates"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieReleases struct
type MovieReleases struct {
	Results []MovieRelease
}

// MovieRating struct
type MovieRating struct {
	*MovieShort
	Rating int
}

// MovieRatingResults struct
type MovieRatingResults struct {
	*PaginatedResults
	*MovieRating
}

// MovieReviews struct
type MovieReviews struct {
	ID      int
	Page    int
	Results []struct {
		ID      string
		Author  string
		Content string
		URL     string
	}
	TotalPages        int                     `json:"total_pages"`
	TotalResults      int                     `json:"total_results"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieTranslations struct
type MovieTranslations struct {
	ID           int
	Translations []struct {
		Iso639_1    string `json:"iso_639_1"`
		Name        string
		EnglishName string `json:"english_name"`
	}
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Videos            *MovieVideos            `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// MovieVideos struct
type MovieVideos struct {
	ID      int
	Results []struct {
		ID        string
		Iso639_1  string `json:"iso_639_1"`
		Iso3166_1 string `json:"iso_3166_1"`
		Key       string
		Name      string
		Site      string
		Size      int
		Type      string
	}
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
	Credits           *MovieCredits           `json:",omitempty"`
	Images            *MovieImages            `json:",omitempty"`
	Keywords          *MovieKeywords          `json:",omitempty"`
	Releases          *MovieReleases          `json:"release_dates,omitempty"`
	Translations      *MovieTranslations      `json:",omitempty"`
	Similar           *MoviePagedResults      `json:",omitempty"`
	Reviews           *MovieReviews           `json:",omitempty"`
	Lists             *MovieLists             `json:",omitempty"`
	Changes           *MovieChanges           `json:",omitempty"`
	// Rating            *MovieRating            `json:",omitempty"`
}

// GetMovieInfo for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-details
func (tmdb *TMDb) GetMovieInfo(id int, options map[string]string) (*Movie, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var movie Movie
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movie)
	return result.(*Movie), err
}

// GetMovieAccountStates gets the status of whether or not the movie has been rated or added to their favourite or movie watch list
// https://developers.themoviedb.org/3/movies/get-movie-account-states
func (tmdb *TMDb) GetMovieAccountStates(id int, sessionID string) (*MovieAccountState, error) {
	var state MovieAccountState
	uri := fmt.Sprintf("%s/movie/%v/account_states?api_key=%s&session_id=%s", baseURL, id, tmdb.apiKey, sessionID)
	result, err := getTmdb(uri, &state)
	return result.(*MovieAccountState), err
}

// GetMovieAlternativeTitles for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (tmdb *TMDb) GetMovieAlternativeTitles(id int, options map[string]string) (*MovieAlternativeTitles, error) {
	var availableOptions = map[string]struct{}{
		"country":            {},
		"append_to_response": {}}
	var titles MovieAlternativeTitles
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/alternative_titles?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &titles)
	return result.(*MovieAlternativeTitles), err
}

// GetMovieChanges for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (tmdb *TMDb) GetMovieChanges(id int, options map[string]string) (*MovieChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes MovieChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*MovieChanges), err
}

// GetMovieCredits for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (tmdb *TMDb) GetMovieCredits(id int, options map[string]string) (*MovieCredits, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var credits MovieCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*MovieCredits), err
}

// GetMovieImages for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-images
func (tmdb *TMDb) GetMovieImages(id int, options map[string]string) (*MovieImages, error) {
	var availableOptions = map[string]struct{}{
		"language":               {},
		"append_to_response":     {},
		"include_image_language": {}}
	var images MovieImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/images?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*MovieImages), err
}

// GetMovieKeywords for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-keywords
func (tmdb *TMDb) GetMovieKeywords(id int, options map[string]string) (*MovieKeywords, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var keywords MovieKeywords
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/keywords?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &keywords)
	return result.(*MovieKeywords), err
}

// GetMovieLatest gets the latest movie
// https://developers.themoviedb.org/3/movies/get-latest-movie
func (tmdb *TMDb) GetMovieLatest() (*Movie, error) {
	var movie Movie
	uri := fmt.Sprintf("%s/movie/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &movie)
	return result.(*Movie), err
}

// GetMovieLists that the movie belongs to
// https://developers.themoviedb.org/3/movies/get-movie-lists
func (tmdb *TMDb) GetMovieLists(id int, options map[string]string) (*MovieLists, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var lists MovieLists
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/lists?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &lists)
	return result.(*MovieLists), err
}

// GetMovieNowPlaying that have been, or are being released this week
// https://developers.themoviedb.org/3/movies/get-now-playing
func (tmdb *TMDb) GetMovieNowPlaying(options map[string]string) (*MovieDatedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var nowPlaying MovieDatedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/now_playing?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &nowPlaying)
	return result.(*MovieDatedResults), err
}

// GetMoviePopular gets the list of popular movies on The Movie Database
// https://developers.themoviedb.org/3/movies/get-popular-movies
func (tmdb *TMDb) GetMoviePopular(options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var popular MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/popular?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &popular)
	return result.(*MoviePagedResults), err
}

// GetMovieReleases for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (tmdb *TMDb) GetMovieReleaseDates(id int, options map[string]string) (*MovieReleases, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var releases MovieReleases
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/releases?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &releases)
	return result.(*MovieReleases), err
}

// GetMovieReviews for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-reviews
func (tmdb *TMDb) GetMovieReviews(id int, options map[string]string) (*MovieReviews, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var reviews MovieReviews
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/reviews?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &reviews)
	return result.(*MovieReviews), err
}

// GetMovieSimilar for a specific movie id
// https://developers.themoviedb.org/3/movies/get-similar-movies
func (tmdb *TMDb) GetMovieSimilar(id int, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var similar MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/similar?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &similar)
	return result.(*MoviePagedResults), err
}

// GetMovieTopRated gets the list of top rated movies
// https://developers.themoviedb.org/3/movies/get-top-rated-movies
func (tmdb *TMDb) GetMovieTopRated(options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var topRated MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/top_rated?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &topRated)
	return result.(*MoviePagedResults), err
}

// GetMovieTranslations for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (tmdb *TMDb) GetMovieTranslations(id int, options map[string]string) (*MovieTranslations, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var translations MovieTranslations
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/translations?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &translations)
	return result.(*MovieTranslations), err
}

// GetMovieVideos for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (tmdb *TMDb) GetMovieVideos(id int, options map[string]string) (*MovieVideos, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var videos MovieVideos
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/videos?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &videos)
	return result.(*MovieVideos), err
}

// GetMovieUpcoming by release date
// https://developers.themoviedb.org/3/movies/get-upcoming
func (tmdb *TMDb) GetMovieUpcoming(options map[string]string) (*MovieDatedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var upcoming MovieDatedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/upcoming?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &upcoming)
	return result.(*MovieDatedResults), err
}

// // SetMovieRating lets users rate a movie
// // http://docs.themoviedb.apiary.io/#reference/movies/movieidrating/post
// func (tmdb *TMDb) SetMovieRating(id int) (*MovieRating, error) {
// 	// TODO
// 	var rating MovieRating
// 	var err error
// 	return &rating, err
// }

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETMovies(t *testing.T) {
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "12", Isbn: "438228", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	t.Run("returns movies", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/movies", nil)
		response := httptest.NewRecorder()

		getMovies(response, request)

		bytes, _ := json.Marshal(&movies)

		assertResponseBody(t, response.Body.String(), string(bytes)+"\n")
	})
	t.Run("returns movie 1", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/movies/1", nil)
		response := httptest.NewRecorder()

		getMovie(response, request)

		bytes, _ := json.Marshal(&movies[0])

		assertResponseBody(t, response.Body.String(), string(bytes)+"\n")
	})
	movies = []Movie{}
}
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
func assertNotError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Should not be an error")
	}
}
func assert(t testing.TB, boolean bool, err string) {
	t.Helper()
	if !boolean {
		t.Errorf(err)
	}
}
func assertSame(t testing.TB, expected *Movie, got *Movie) {
	t.Helper()
	same := expected.Isbn == got.Isbn && expected.Title == got.Title && expected.Director.FirstName == got.Director.FirstName && expected.Director.LastName == got.Director.LastName
	if !same {
		t.Errorf("%v is not equals to %v", expected, got)
	}
}
func TestCreateMovies(t *testing.T) {
	movie := Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}}

	bytess, err := json.Marshal(movie)
	assertNotError(t, err)

	request, err := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(bytess))
	assertNotError(t, err)
	response := httptest.NewRecorder()

	var newMovie Movie
	createMovie(response, request)

	err = json.Unmarshal(response.Body.Bytes(), &newMovie)
	assertNotError(t, err)

	assertSame(t, &movie, &newMovie)
	assertSame(t, &movie, &movies[0])
	movies = []Movie{}
}

func TestUpdateMovies(t *testing.T) {
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "12", Isbn: "438228", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	movie := Movie{ID: "1", Isbn: "438227", Title: "Movie Three", Director: &Director{FirstName: "John", LastName: "Doe"}}

	bytess, err := json.Marshal(movie)
	assertNotError(t, err)

	request, err := http.NewRequest(http.MethodPut, "/movies/1", bytes.NewBuffer(bytess))
	assertNotError(t, err)
	response := httptest.NewRecorder()

	var newMovie Movie
	updateMovie(response, request)

	err = json.Unmarshal(response.Body.Bytes(), &newMovie)
	assertNotError(t, err)

	assertSame(t, &movie, &newMovie)
	movies = []Movie{}
}

func TestDeleteMovies(t *testing.T) {
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "12", Isbn: "438228", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	request, err := http.NewRequest(http.MethodDelete, "/movies/1", nil)
	assertNotError(t, err)
	response := httptest.NewRecorder()

	deleteMovie(response, request)

	bytes, _ := json.Marshal(&movies)

	assertResponseBody(t, response.Body.String(), string(bytes)+"\n")
	assert(t, len(movies) == 1, "Should not have two elements")
	movies = []Movie{}
}

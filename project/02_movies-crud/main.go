package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438202", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Don"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438203", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	setApplication(writer)
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	setApplication(writer)
	params := mux.Vars(request)
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			moviePtr := &movies[index]
			if "" != movie.Isbn {
				moviePtr.Isbn = movie.Isbn
			}
			if "" != movie.Title {
				moviePtr.Title = movie.Title
			}
			if nil != movie.Director {
				moviePtr.Director = movie.Director
			}
			json.NewEncoder(writer).Encode(moviePtr)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	setApplication(writer)
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	setApplication(writer)
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	setApplication(writer)
	json.NewEncoder(writer).Encode(movies)
}

func setApplication(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// Movie 电影
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director 导演
type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

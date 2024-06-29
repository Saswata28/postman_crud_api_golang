package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"math/rand"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func mainPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello")
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)	
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies{
		if item.ID == param["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){ //We are deleting the movie then creating a new movie with the new updated elements with the same id number. Not recommended.
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	
	for index, item := range movies {
		if item.ID == param["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			// creating new movie with same id.
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)
			break
		}
	}
	

	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"56565", Title:"hello", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn:"56566", Title:"something", Director: &Director{Firstname: "sam", Lastname: "Tarly"}})
	
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting on port 8080:\nGo to http://127.0.0.1:8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}

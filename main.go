package main

import (
"fmt"
"net/http"//Para poder usar todos los metodos http
"log"
"encoding/json"
"github.com/gorilla/mux"
)


func homepage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"Entonces el http.ResposeWriter es lo que nos permite enviar el contenido a la pagina web")
	fmt.Println("Endpoint logrado: homepage")
}

func handleRequests(){

	myRouter:=mux.NewRouter().StrictSlash(true)
myRouter.HandleFunc("/",homepage)
myRouter.HandleFunc("/articles",returnAllArticles)
myRouter.HandleFunc("/articles/{id}",returnSingleArticle)
log.Fatal(http.ListenAndServe(":10000",myRouter))
}

type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}


//This function will return all articles in a json string, all 
//elements stored on our Articles Array
func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}


//If we do not need all articles then we can create a function that returns
//only a single article thanks to our Gorilla/mux router
//These are Path Varianles (variables handled on our path or URL)
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
vars:=mux.Vars(r)
key:= vars["id"]
fmt.Fprintf(w, "Key: "+key)


//We loop over all our articles and once we find it we return 
//the alement encoded as JSON object
for _, article:= range Articles{
	if article.Id==key{
		json.NewEncoder(w).Encode(article)
	}
}

}

var Articles []Article

func main(){

	Articles=[]Article{
		Article{Id: "1", Title: "Hello", Desc: "Descripcion",Content: "Article Content"},
		Article{Id: "2", Title: "Hellwwwo", Desc: "Descripcion",Content: "Article Content"},
	}

	handleRequests()
}
package main

import (
"fmt"
"net/http"//Para poder usar todos los metodos http
"log"
"encoding/json"
"github.com/gorilla/mux"
"io/ioutil"
)


func homepage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"Entonces el http.ResposeWriter es lo que nos permite enviar el contenido a la pagina web")
	fmt.Println("Endpoint logrado: homepage")
}

func handleRequests(){

	myRouter:=mux.NewRouter().StrictSlash(true)
	//Order on our requests is very important to take into consideration when building an REST API
myRouter.HandleFunc("/",homepage)
myRouter.HandleFunc("/articles",returnAllArticles)
myRouter.HandleFunc("/article",createNewArticle).Methods("POST")
//add our new DELEEETE endpoint here
myRouter.HandleFunc("/article/{id}",deleteArticle).Methods("DELETE")
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

//This can be taken as a GET petition :)

}

func createNewArticle(w http.ResponseWriter, r *http.Request){
//In order to create a new Article, it will be necesary to implement a POST petition

//get the body of our POST request
//return the string response containing the request body

reqBody,_:=ioutil.ReadAll(r.Body)
fmt.Fprintf(w,"%+v", string(reqBody))

/*Now we can add the body object into our array. In order to do so we need to unmarshal
the reqBody that was sent through the POST request
*/

  // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    

	var article Article
	json.Unmarshal(reqBody, &article)//aqui la info obtenida del reqbody la metemos a la variable article

	//Update our global Articles array to include our new article sent by the POST request

	Articles=append(Articles,article)
	//json.NewEncoder(w).Encode(article)

}

/*
Remember that ALL methods implemented to handle information on our REST API
need to have the w http.ResposeWriter and r *http.Request parameters so we know
that those methos are the ones being part of the actual API\
*/
func deleteArticle(w http.ResponseWriter, r *http.Request){

	//once again we will need to parse the path parameters (the parameter we will sent through the URL)
	vars:=mux.Vars(r) //Path parameter taken


	//we will need to extract the `id` of the article we wish to delete
	id:=vars["id"]

	//we then need to loop trough all our articles

	for index, article:=range Articles{
		//if our id path parameter matches one of our articles

		if article.Id==id{
			//uptade our Articles array to remove the artice we chose
			Articles=append(Articles[:index],Articles[index+1:]...)
			json.NewEncoder(w).Encode("Article: "+article.Id+" deleted\n"+"Hehe")
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
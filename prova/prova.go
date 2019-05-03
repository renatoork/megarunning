package prova

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"projetos/megarunning/atleta"
	"projetos/megarunning/conexao"
)

type ValidaToken struct {
	Error string `json:"error"`
	Email string `json:"email"`
}

var collection *mgo.Collection = conexao.GetDb().C("Prova")

func Handler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	fmt.Println("Entrou no handler")
	fmt.Println(len(request.Cookies()))
	token, err := request.Cookie("token")
	fmt.Println(token.Value, err)
	data := url.Values{}
	data.Set("access_token", token.Value)
	//req = http.NewRequest("POST", "https://www.googleapis.com/oauth2/v1/tokeninfo", nil)
	resp, err := http.Post("https://www.googleapis.com/oauth2/v1/tokeninfo", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param="+token)
	//http.
	var retorno ValidaToken
	json.NewDecoder(resp.Body).Decode(&retorno)
	fmt.Println(retorno.Error, retorno.Email)

	switch request.Method {
	case "GET":
		provas := Provas()
		fmt.Println(provas)
		js, _ := json.MarshalIndent(provas, " ", "   ")
		response.Write(js)

	case "POST":
		var prova Prova
		json.NewDecoder(request.Body).Decode(&prova)
		err := Insere(&prova)
		fmt.Println(err)
	}
}

func HandlerId(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	fmt.Println("Entrou no handlerId")

	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		id := bson.ObjectIdHex(vars["id"])

		prova := GetProva(id)
		fmt.Println(prova)
		js, _ := json.MarshalIndent(prova, " ", "   ")
		response.Write(js)

	case "POST":
		var prova Prova
		json.NewDecoder(request.Body).Decode(&prova)
		fmt.Println("prova.Id")
		if prova.Id == "" {
			err := Insere(&prova)
			fmt.Println(err)
		} else {
			err := Update(&prova)
			fmt.Println(err)
		}

	case "DELETE":
		vars := mux.Vars(request)
		id := bson.ObjectIdHex(vars["id"])
		err := Delete(id)
		fmt.Println(err)
	}
}

func HandlerInscricao(response http.ResponseWriter, request *http.Request) {
	atleta := atleta.Atleta{}
	id := bson.ObjectIdHex(mux.Vars(request)["id"])
	atleta.Nome = "João Fernando Barros"
	atleta.Email = "joao.exemplo@mega.com.br"
	prova := GetProva(id)
	prova.Participantes = append(prova.Participantes, Participacao{ValorPatrocinio: 0, Atleta: atleta})
	err := Update(&prova)
	fmt.Println(err)
}

func Provas() []Prova {
	var provas []Prova
	collection.Find(nil).All(&provas)
	return provas
}

func GetProva(id bson.ObjectId) Prova {
	var prova Prova
	collection.FindId(id).One(&prova)
	return prova
}

func Insere(prova *Prova) error {

	prova.Id = bson.NewObjectId()
	fmt.Println(prova)
	return collection.Insert(prova)

}

func Update(prova *Prova) error {

	fmt.Println(prova.Id)
	return collection.Update(bson.M{"_id": prova.Id}, &prova)

}

func Delete(id bson.ObjectId) error {

	fmt.Println(id)
	return collection.RemoveId(id)

}

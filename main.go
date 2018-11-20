package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-password/password"
	"log"
	"net/http"
	"strings"
)
type Client struct {
	Name, URL, Database, DBPassword	string

}
func (c Client) clientMysqlPassword() string {
	clientMysqlPassword, err := password.Generate(16, 4, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}
	return clientMysqlPassword
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	var client = Client{}
	clients := strings.Split(string(client.lsHelm()), "\n")
	clientsJson, _ := json.Marshal(clients[:len(clients)-1])
	w.Header().Set("Content-Type", "application/json")
	w.Write(clientsJson)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	var client = Client{
		Name: r.FormValue("client"),
		URL: r.FormValue("url"),
	}
	if client.Name != "" && client.URL != "" {
		client.Database = "db_" + client.Name
		client.createDatabase()
		output := strings.Split(string(client.installHelm()), "\n")
		clientsJson, _ := json.Marshal(output)
		w.Header().Set("Content-Type", "application/json")
		w.Write(clientsJson)
	} else {
		log.Println("I need client's name and URL")
	}

}

func UpgradeClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PATCH")
	var client = Client{
		Name: r.FormValue("client"),
		URL: r.FormValue("url"),
	}
	clientsJson, _ := json.Marshal(client)
	w.Header().Set("Content-Type", "application/json")
	w.Write(clientsJson)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	var client = Client{
		Name: r.FormValue("client"),
	}
	if client.Name != "" {
		client.Database = "db_" + client.Name
		res := client.delHelm()
		client.dropDatabase()
		clientsJson, _ := json.Marshal(string(res))
		w.Header().Set("Content-Type", "application/json")
		w.Write(clientsJson)
	} else {
		log.Println("I need client's name")
	}
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/user/", GetClients).Methods("GET")
	router.HandleFunc("/user/", CreateClient).Methods("POST")
	router.HandleFunc("/user/", UpgradeClient).Methods("PATCH")
	router.HandleFunc("/user/", DeleteClient).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

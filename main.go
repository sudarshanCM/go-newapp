package main

import (
	"fmt"
	"log"
	"net/http"
	
	
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error


type Details struct {

	
	Name  string `json:"name"`
	AadharNumber  string  `json:"aadharNumber"`
	Phone   string  `json:"phone"`

}


func InitialMigration() {

	db, err := gorm.Open("mysql", "root:root@/AngularDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	} else {
		
		fmt.Println("SERVER1");
		fmt.Println("Connected successfully")
	
		
	}
	defer db.Close()
	db.AutoMigrate(&Details{})
	fmt.Println("SERVER2");
	fmt.Println("HELLO");
	
	}

// func helloworld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Helloworld")
// }
func handleRequests() {
	fmt.Println("SERVER");
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user", Newdetails).Methods("POST")
	fmt.Println("SERVER3");
	log.Fatal(http.ListenAndServe(":3007", cors.Default().Handler(myRouter)))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	InitialMigration()
	handleRequests()
}
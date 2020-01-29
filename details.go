package main

import (
	
	"fmt"
	"encoding/json"
	
	"net/http"
	


	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Newdetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var userDetails Details
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userDetails) 
	defer r.Body.Close()
	
	db, err := gorm.Open("mysql", "root:root@/AngularDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println(db,"Connection Successfull")
	}
	
	name := userDetails.Name
	aadharNumber := userDetails.AadharNumber
	phone:=userDetails.Phone
	
	result:=db.Exec("INSERT INTO details(name,aadharNumber,phone) VALUES(?,?,?)",name,aadharNumber,phone)
	json.NewEncoder(w).Encode(result)
	fmt.Println("result",result);
	

	fmt.Println(w, "New user created successfully")
	

}
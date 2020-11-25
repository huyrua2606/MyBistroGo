package main

import (
	"github.com/gorilla/mux"
)

func API(route *mux.Router) {
	route.HandleFunc("/regis", createAccount).Methods("POST")

	route.HandleFunc("/getaccount", getAccount).Methods("GET")

	route.HandleFunc("/createmenu", createMenu).Methods("POST")

	route.HandleFunc("/getmenu", getMenu).Methods("GET")

	route.HandleFunc("/updatemenu", updateMenu1).Methods("PUT")

	route.HandleFunc("/createorder", createOrder).Methods("POST")

	route.HandleFunc("/getorderhistory", getOrderHistory).Methods("GET")

	route.HandleFunc("/addingredient", addIngredient).Methods("POST")

	route.HandleFunc("/deleteingredient", deleteIngredient).Methods("DELETE")
}

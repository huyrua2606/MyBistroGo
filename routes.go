package main

import (
	"net/http"

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

	route.HandleFunc("/adddish", addDish).Methods("POST")

	route.HandleFunc("/createbistro", createBistro).Methods("POST")

	route.HandleFunc("/regisbistroman", regisBistroManager).Methods("POST")

	route.HandleFunc("/getbistroman", getBistroMan).Methods("GET")

	route.HandleFunc("/regiscustomer", regisCustomer).Methods("POST")

	route.PathPrefix("/customerimage/").Handler(http.StripPrefix("/customerimage/", http.FileServer(http.Dir("./CustomerImage"))))

	route.PathPrefix("/bistromanagerimage/").Handler(http.StripPrefix("/bistromanagerimage/", http.FileServer(http.Dir("./BistroManagerImage"))))

	route.PathPrefix("/bistroimage/").Handler(http.StripPrefix("/bistroimage/", http.FileServer(http.Dir("./BistroImage"))))

	route.PathPrefix("/dishimage/").Handler(http.StripPrefix("/dishimage/", http.FileServer(http.Dir("./DishImage"))))

	route.HandleFunc("/getaccounttype", getAccountType).Methods("GET")

	route.HandleFunc("/updatedoneorder", updateOrderStatusDone).Methods("PUT")

	route.HandleFunc("/getbistro", getBistro).Methods("GET")

	route.HandleFunc("/getcustomer", getCustomer).Methods("GET")

	route.HandleFunc("/getbistrolist", getBistroList).Methods("GET")

	route.HandleFunc("/getdishfrommenu", getDishFromMenu).Methods("GET")

	route.HandleFunc("/getmenulist", getMenuList).Methods("GET")

	route.HandleFunc("/acceptorder", acceptOrder).Methods("PUT")

	route.HandleFunc("/declineorder", declineOrder).Methods("PUT")

	route.HandleFunc("/finishorder", finishOrder).Methods("PUT")

	route.HandleFunc("/getdishesbybid", getDishesByBID).Methods("GET")

	route.HandleFunc("/addorderdetail", addOrderDetail).Methods("POST")

	route.HandleFunc("/createcart", createCart).Methods("POST")

	route.HandleFunc("/adddishincart", addDishInCart).Methods("POST")

	route.HandleFunc("/getdishbycartid", getDishByCartID).Methods("GET")

	route.HandleFunc("/getcartid", getCartID).Methods("GET")
}

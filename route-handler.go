package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func createAccount(response http.ResponseWriter, request *http.Request) {
	var cus User
	var ReturnData string
	cus.username = request.URL.Query().Get("username")
	cus.password = request.URL.Query().Get("password")

	if checkAccountCreated(cus) == true {
		ReturnData = "Tai khoan da duoc tao tu truoc"
	} else {
		db.Query("INSERT INTO `mybistro`.`user` (`username`,`password`) VALUES ('" + cus.username + "','" + cus.password + "')")
		ReturnData = "Tao tai khoan thanh cong"
	}
	jsonResponse, jsonError := json.Marshal(ReturnData)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}
	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}

}

func checkAccountCreated(user User) bool {
	var UserCheck User
	createAccept, err := db.Query("SELECT * FROM `mybistro`.`user` where `username` = '" + user.username + "'")
	if err != nil {
		panic(err.Error())
		return true
	}
	for createAccept.Next() {
		err2 := createAccept.Scan(&UserCheck.UID, &UserCheck.username, &UserCheck.password)
		if err2 != nil {
			panic(err2.Error())
			return true
		}

	}
	if createAccept != nil {
		if UserCheck.username == user.username {
			return true
		} else {
			return false
		}
	}
	return true
}

func getAccount(response http.ResponseWriter, request *http.Request) {
	var (
		userlogin  User
		userlogin2 User
	)
	userlogin.username = request.URL.Query().Get("username")
	userlogin.password = request.URL.Query().Get("password")

	loginaccept, err := db.Query("Select * FROM `user` WHERE `username` = '" + userlogin.username + "' AND `password` = '" + userlogin.password + "'")
	for loginaccept.Next() {
		err2 := loginaccept.Scan(&userlogin2.UID, &userlogin2.username, &userlogin2.password)

		if err2 != nil {
			fmt.Println(err2)
		}
	}
	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}

	if loginaccept != nil {
		var ReturnData string
		if userlogin2.username == userlogin.username && userlogin2.password == userlogin.password {
			ReturnData = "Dang Nhap Thanh Cong"
		} else {
			ReturnData = "Dang Nhap That Bai"
		}
		jsonResponse, jsonError := json.Marshal(ReturnData)
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}

	fmt.Println(userlogin.username)
	fmt.Println(userlogin.username)

}

func createMenu(response http.ResponseWriter, request *http.Request) {
	var menu Menu
	menu.BID = request.URL.Query().Get("bid")
	menu.MenuName = request.URL.Query().Get("name")

	_, err := db.Query("INSERT INTO `mybistro`.`menu` (`BID`, `MenuName`) VALUES (" + menu.BID + ",'" + menu.MenuName + "')")
	if err != nil {
		fmt.Println(err)
	}
}

func getMenu(response http.ResponseWriter, request *http.Request) {
	var menu Menu

	menu.MID = request.URL.Query().Get("menuid")

	menu1, err := db.Query("SELECT * FROM `menu` WHERE `MID` = '" + menu.MID + "'")
	if err != nil {
		fmt.Println(err)

	}
	for menu1.Next() {
		err2 := menu1.Scan(&menu.MID, &menu.BID, &menu.MenuName)

		if err2 != nil {
			fmt.Println(err2)
		}
	}
	if menu1 != nil {

		jsonResponse, jsonError := json.Marshal(menu)
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}

	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("It's not you it's me.")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}

//Update MenuName
func updateMenu1(response http.ResponseWriter, request *http.Request) {
	var (
		user User
		menu Menu
	)
	user.username = request.URL.Query().Get("username")
	user.password = request.URL.Query().Get("password")
	menu.MenuName = request.URL.Query().Get("menuname")

	_, err := db.Query("UPDATE `mybistro`.`menu` SET `MenuName` = '" + menu.MenuName + "' WHERE `menu`.`BID` = `bistro`.`BID` AND `bistro`.`UID` = `user`.`iduser` and `user`.`username` = `" + user.username + "` AND `user`.`password` = `" + user.password + "` ")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprint(response, "Sua thong tin thanh cong")
	}
}

func createOrder(response http.ResponseWriter, request *http.Request) {
	var (
		o Order
	)
	o.UID = request.URL.Query().Get("userid")
	o.BID = request.URL.Query().Get("bistroid")
	o.Price = request.URL.Query().Get("price")
	o.BeginDate = strconv.FormatInt(time.Now().Unix(), 10)
	o.EndDate = strconv.FormatInt(time.Now().Unix(), 10)

	_, err := db.Query("INSERT INTO `mybistro`.`order_history` (`UID`,`BID`,`Price`,`BeginDate`,`EndDate`) VALUES ('" + o.UID + "','" + o.BID + "','" + o.Price + "'," + o.BeginDate + "," + o.EndDate + ");")

	if err != nil {
		panic(err.Error())
	}

	jsonResponse, jsonError := json.Marshal("Order thanh cong")
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}
	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}

}

func getOrderHistory(response http.ResponseWriter, request *http.Request) {
	var (
		order  Order
		orders []Order
		BID    string
	)

	UID := request.URL.Query().Get("userid")

	rows, err := db.Query("SELECT * FROM `mybistro`.`order_history` WHERE `UID` = " + UID + ";")

	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&order.OrderID, &order.UID, &BID, &order.Price, &order.BeginDate, &order.EndDate)

			rows1, err1 := db.Query("Select Name FROM `mybistro`.`bistro` WHERE `BID` = '" + BID + "';")
			if err1 != nil {
				panic(err.Error())
			} else {
				for rows1.Next() {
					rows1.Scan(&order.BID)

				}
				defer rows1.Close()

			}

			orders = append(orders, order)
			defer rows1.Close()
		}
		defer rows.Close()
		jsonResponse, jsonError := json.Marshal(orders)
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

func addIngredient(response http.ResponseWriter, request *http.Request) {
	var ingredient Ingredient

	ingredient.IngreName = request.URL.Query().Get("name")
	ingredient.MinQuantity = request.URL.Query().Get("minquantity")
	ingredient.Measure = request.URL.Query().Get("measure")

	_, err := db.Query("INSERT INTO `mybistro`.`ingredient` (`Name`,`MinQuantity`,`Measure`) VALUES ('" + ingredient.IngreName + "'," + ingredient.MinQuantity + ",'" + ingredient.Measure + " ')")
	if err != nil {
		panic(err.Error())
	} else {
		jsonResponse, jsonError := json.Marshal("Tao ingredient thanh cong")
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

func deleteIngredient(response http.ResponseWriter, request *http.Request) {
	var ingredient Ingredient

	ingredient.IngreID = request.URL.Query().Get("ingreid")

	_, err := db.Query("DELETE FROM `mybistro`.`ingredient` WHERE `IngreID` = " + ingredient.IngreID + ";")
	if err != nil {
		panic(err.Error())
	} else {
		jsonResponse, jsonError := json.Marshal("Xoa ingredient thanh cong")
		if jsonError != nil {
			fmt.Println(jsonError)
			returnErrorResponse(response, request)
		}
		if jsonResponse == nil {
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

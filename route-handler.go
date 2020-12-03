package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const maxUploadSize = 2 * 1024 * 1024

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
			ReturnData = userlogin2.UID
		} else {
			ReturnData = "-1"
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

	fmt.Println("Account login:" + userlogin.username)
	fmt.Println(userlogin.password)

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
		o          Order
		returnData string
	)
	o.UID = request.URL.Query().Get("userid")
	o.BID = request.URL.Query().Get("bistroid")
	o.Price = request.URL.Query().Get("price")
	o.BeginDate = strconv.FormatInt(time.Now().Unix(), 10)
	o.EndDate = strconv.FormatInt(time.Now().Unix(), 10)
	o.Status = newOrder

	_, err := db.Query("INSERT INTO `mybistro`.`order_history` (`UID`,`BID`,`Price`,`BeginDate`,`EndDate`,`Status`) VALUES (" + o.UID + "," + o.BID + ",'" + o.Price + "'," + o.BeginDate + "," + o.EndDate + "," + o.Status + ");")

	if err != nil {
		panic(err.Error())
	}
	time.Sleep(50 * time.Millisecond)
	rows, err1 := db.Query("SELECT OrderID FROM `mybistro`.`order_history` ORDER By OrderID desc Limit 1")
	if err1 != nil {
		panic(err.Error())
	}
	for rows.Next() {
		rows.Scan(&returnData)
	}
	jsonResponse, jsonError := json.Marshal(returnData)
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
			rows.Scan(&order.OrderID, &order.UID, &BID, &order.Price, &order.BeginDate, &order.EndDate, &order.Status)

			rows1, err1 := db.Query("Select Name FROM `mybistro`.`bistro` WHERE `BID` = " + BID + ";")
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

func addDish(response http.ResponseWriter, request *http.Request) {
	var dish Dish
	var ReturnData string
	dish.MID = request.URL.Query().Get("menuid")
	dish.Name = request.URL.Query().Get("name")
	dish.Price = request.URL.Query().Get("price")

	_, err := db.Query("INSERT INTO `mybistro`.`dish` (`MID`,`Name`,`Price`) VALUES (" + dish.MID + ",'" + dish.Name + "'," + dish.Price + ")")
	if err != nil {
		panic(err.Error())
		ReturnData = "Tao mon an khong thanh cong"
	} else {
		ReturnData = "Tao mon an thanh cong"
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

func createBistro(response http.ResponseWriter, request *http.Request) {
	var bistro Bistro
	var ReturnData string
	bistro.Name = request.URL.Query().Get("name")
	bistro.Address = request.URL.Query().Get("address")
	bistro.PhoneNum = request.URL.Query().Get("phonenum")
	bistro.Email = request.URL.Query().Get("email")

	exist := checkBistroCreated(bistro.Name, bistro.Address)
	if exist == true {
		ReturnData = "Bistro da duoc tao tu truoc"
	} else {
		_, err := db.Query("INSERT INTO `mybistro`.`bistro` (`Name`,`Address`,`PhoneNum`,`Email`) VALUES ('" + bistro.Name + "','" + bistro.Address + "','" + bistro.PhoneNum + "','" + bistro.Email + "')")

		if err != nil {
			panic(err.Error())
			ReturnData = err.Error()
		} else {
			ReturnData = "Dang ky Bistro Thanh cong"
		}

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

func regisBistroManager(response http.ResponseWriter, request *http.Request) {
	var bm Bistro_Manager
	var cus User
	var ReturnData string

	bm.Name = request.URL.Query().Get("name")
	bm.BID = request.URL.Query().Get("bid")

	cus.username = request.URL.Query().Get("username")
	cus.password = request.URL.Query().Get("password")

	if checkAccountCreated(cus) == true {
		ReturnData = "Tai khoan da duoc tao tu truoc"
	} else {
		go func() {
			_, err := db.Query("INSERT INTO `mybistro`.`user` (`username`,`password`) VALUES ('" + cus.username + "','" + cus.password + "');")

			if err != nil {
				panic(err.Error())

			}

		}()
		time.Sleep(100 * time.Millisecond)
		rows, err := db.Query("SELECT iduser FROM `mybistro`.`user` WHERE `username`='" + cus.username + "' AND `password`='" + cus.password + "';")
		if err != nil {
			panic(err.Error())
		}

		for rows.Next() {
			rows.Scan(&bm.UID)
			fmt.Println(bm.UID)
		}
		db.Query("INSERT INTO `mybistro`.`bistro_manager` (`UID`,`Name`,`BID`) VALUES (" + bm.UID + ",'" + bm.Name + "'," + bm.BID + ");")
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

func checkBistroCreated(Name, Address string) bool {
	var BID string
	rows, err := db.Query("SELECT BID from `mybistro`.`bistro` WHERE `Name` = '" + Name + "' AND `Address` = '" + Address + "'")
	for rows.Next() {
		rows.Scan(&BID)
	}

	fmt.Println(BID)
	if err != nil {
		panic(err.Error())
		return true
	} else if BID == "" {
		return false
	} else {
		return true
	}
}

func getBistroMan(response http.ResponseWriter, request *http.Request) {
	UID := request.URL.Query().Get("uid")
	var bm Bistro_Manager
	rows, err := db.Query("SELECT * from `mybistro`.`bistro_manager` WHERE `UID` = '" + UID + "' ")
	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&bm.UID, &bm.Name, &bm.BID, &bm.Image)
		}
	}
	jsonResponse, jsonError := json.Marshal(bm)
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

func regisCustomer(response http.ResponseWriter, request *http.Request) {
	var cus Customer
	var usr User
	var ReturnData string

	cus.FullName = request.URL.Query().Get("name")
	cus.Email = request.URL.Query().Get("email")
	cus.DoB = request.URL.Query().Get("dob")
	cus.PhoneNum = request.URL.Query().Get("phonenum")

	usr.username = request.URL.Query().Get("username")
	usr.password = request.URL.Query().Get("password")

	if checkAccountCreated(usr) == true {
		ReturnData = "Tai khoan da duoc tao tu truoc"
	} else {
		go func() {
			_, err := db.Query("INSERT INTO `mybistro`.`user` (`username`,`password`) VALUES ('" + usr.username + "','" + usr.password + "');")

			if err != nil {
				panic(err.Error())

			}

		}()
		time.Sleep(100 * time.Millisecond)
		rows, err := db.Query("SELECT iduser FROM `mybistro`.`user` WHERE `username`='" + usr.username + "' AND `password`='" + usr.password + "';")
		if err != nil {
			panic(err.Error())
		}

		for rows.Next() {
			rows.Scan(&cus.UID)
			fmt.Println(cus.UID)
		}

		_, err2 := db.Query("INSERT INTO `mybistro`.`customer` (`UID`,`FullName`,`Email`,`DoB`,`PhoneNum`) VALUES (" + cus.UID + ",'" + cus.FullName + "','" + cus.Email + "'," + cus.DoB + ",'" + cus.FullName + "');")
		if err2 != nil {
			panic(err.Error())
			ReturnData = "Tao tai khoan khong thanh cong"
		} else {
			ReturnData = "Tao tai khoan thanh cong"
		}

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

func getAccountType(response http.ResponseWriter, request *http.Request) {
	var cus User
	var UID string
	var ReturnData string
	cus.username = request.URL.Query().Get("username")
	cus.password = request.URL.Query().Get("password")

	rows, err := db.Query("SELECT UID FROM `mybistro`.`user`,`mybistro`.`bistro_manager` WHERE `user`.`username` = '" + cus.username + "' AND `user`.`password` = '" + cus.password + "' AND `mybistro`.`user`.`iduser` = `mybistro`.`bistro_manager`.`UID`;")
	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&UID)
		}
		if UID == "" {
			ReturnData = "Customer"
		} else {
			ReturnData = "BM"
		}
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

func updateOrderStatusDone(response http.ResponseWriter, request *http.Request) {

	var order Order
	var ReturnData string
	order.UID = request.URL.Query().Get("userid")
	order.BID = request.URL.Query().Get("bistroid")
	order.BeginDate = request.URL.Query().Get("begindate")
	order.EndDate = request.URL.Query().Get("enddate")

	_, err := db.Query("UPDATE `mybistro`.`order_history` SET `Status`='" + orderFinish + "', `EndDate`='" + strconv.FormatInt(time.Now().Unix(), 10) + "' WHERE `UID`='" + order.UID + "' AND `BID`='" + order.BID + "' AND `BeginDate`='" + order.BeginDate + "' AND `EndDate`='" + order.EndDate + "'")

	if err != nil {
		panic(err.Error())
		ReturnData = "Update status that bai"
	} else {
		ReturnData = "Update status thanh cong"
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

func getBistro(response http.ResponseWriter, request *http.Request) {
	var bistro Bistro

	bistro.BID = request.URL.Query().Get("bid")

	rows, err := db.Query("SELECT * from `mybistro`.`bistro` WHERE `BID` = '" + bistro.BID + "'")

	if err != nil {
		panic(err.Error())
		jsonResponse, jsonError := json.Marshal("Truy xuat thong tin that bai")
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
	} else {
		for rows.Next() {
			rows.Scan(&bistro.BID, &bistro.Name, &bistro.Address, &bistro.PhoneNum, &bistro.Email, &bistro.Status, &bistro.Rating, &bistro.Reputation, &bistro.Image)
		}
		jsonResponse, jsonError := json.Marshal(bistro)
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

func getCustomer(response http.ResponseWriter, request *http.Request) {
	UID := request.URL.Query().Get("uid")
	var cus Customer
	rows, err := db.Query("SELECT * from `mybistro`.`customer` WHERE `UID` = '" + UID + "';")
	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&cus.UID, &cus.FullName, &cus.Email, &cus.DoB, &cus.PhoneNum, &cus.Image)
		}
		jsonResponse, jsonError := json.Marshal(cus)
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

func getBistroList(response http.ResponseWriter, request *http.Request) {
	var (
		bistro  Bistro
		bistros []Bistro
	)

	rows, err := db.Query("SELECT * FROM `mybistro`.`bistro`;")

	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&bistro.BID, &bistro.Name, &bistro.Address, &bistro.PhoneNum, &bistro.Email, &bistro.Status, &bistro.Rating, &bistro.Reputation, &bistro.Image)

			bistros = append(bistros, bistro)

		}
		defer rows.Close()
		jsonResponse, jsonError := json.Marshal(bistros)
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

func getDishFromMenu(response http.ResponseWriter, request *http.Request) {
	var (
		dish   Dish
		dishes []Dish
		menuid string
	)

	menuid = request.URL.Query().Get("menuid")

	rows, err := db.Query("SELECT * FROM `mybistro`.`dish` WHERE `MID`='" + menuid + "';")

	if err != nil {
		panic(err.Error())
		jsonResponse, jsonError := json.Marshal("Lay data that bai")
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
	} else {
		for rows.Next() {
			rows.Scan(&dish.DID, &dish.MID, &dish.Name, &dish.Price, &dish.Image)

			dishes = append(dishes, dish)

		}
		defer rows.Close()
		jsonResponse, jsonError := json.Marshal(dishes)
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

func getMenuList(response http.ResponseWriter, request *http.Request) {
	var (
		menu  Menu
		menus []Menu
		BID   string
	)
	BID = request.URL.Query().Get("bid")

	rows, err := db.Query("SELECT * FROM `mybistro`.`menu` WHERE `BID`='" + BID + "';")

	if err != nil {
		panic(err.Error())
		jsonResponse, jsonError := json.Marshal("Lay data that bai")
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
	} else {
		for rows.Next() {
			rows.Scan(&menu.MID, &menu.BID, &menu.MenuName)

			menus = append(menus, menu)
		}
		jsonResponse, jsonError := json.Marshal(menus)
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

func acceptOrder(response http.ResponseWriter, request *http.Request) {
	var order Order
	var ReturnData string
	order.OrderID = request.URL.Query().Get("oid")
	Status := checkOrderStatus(order.OrderID)

	if Status == "Khong the lay du lieu" {
		ReturnData = Status
	} else if Status == "-1" || Status == "1" || Status == "2" || Status == "" {
		ReturnData = "Du lieu sai"
	} else {
		_, err := db.Query("UPDATE `mybistro`.`order_history` SET `Status`='1' WHERE `OrderID` ='" + order.OrderID + "'")
		if err != nil {
			panic(err.Error())
			ReturnData = "Khong the update du lieu"
		} else {
			ReturnData = "Chap nhan don hang thanh cong"
		}
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

func checkOrderStatus(OrderID string) string {
	var Status string
	rows, err := db.Query("SELECT Status FROM `mybistro`.`order_history` WHERE `OrderID` = '" + OrderID + "'")
	if err != nil {
		panic(err.Error())
		return "Khong the lay du lieu"
	} else {
		for rows.Next() {
			rows.Scan(&Status)
		}
		return Status
	}
}

func declineOrder(response http.ResponseWriter, request *http.Request) {
	var order Order
	var ReturnData string
	order.OrderID = request.URL.Query().Get("oid")
	Status := checkOrderStatus(order.OrderID)

	if Status == "Khong the lay du lieu" {
		ReturnData = Status
	} else if Status == "-1" || Status == "1" || Status == "2" || Status == "" {
		ReturnData = "Du lieu sai"
	} else {
		_, err := db.Query("UPDATE `mybistro`.`order_history` SET `Status`='-1' WHERE `OrderID` ='" + order.OrderID + "'")
		if err != nil {
			panic(err.Error())
			ReturnData = "Khong the update du lieu"
		} else {
			ReturnData = "Tu choi don hang thanh cong"
		}
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

func finishOrder(response http.ResponseWriter, request *http.Request) {
	var order Order
	var ReturnData string
	order.OrderID = request.URL.Query().Get("oid")
	Status := checkOrderStatus(order.OrderID)

	if Status == "Khong the lay du lieu" {
		ReturnData = Status
	} else if Status == "-1" || Status == "2" || Status == "0" || Status == "" {
		ReturnData = "Du lieu sai"
	} else {
		_, err := db.Query("UPDATE `mybistro`.`order_history` SET `Status`='2' WHERE `OrderID` ='" + order.OrderID + "'")
		if err != nil {
			panic(err.Error())
			ReturnData = "Khong the update du lieu"
		} else {
			ReturnData = "Hoan thanh don hang"
		}
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

func getDishesByBID(response http.ResponseWriter, request *http.Request) {
	var (
		BID    string
		dish   Dish
		dishes []Dish
	)

	BID = request.URL.Query().Get("bid")

	rows, err := db.Query("SELECT `mybistro`.`dish`.`DID`,`mybistro`.`dish`.`MID`,`mybistro`.`dish`.`Name`,`mybistro`.`dish`.`Price`,`mybistro`.`dish`.`Image` FROM `mybistro`.`dish`,`mybistro`.`menu`,`mybistro`.`bistro` WHERE `mybistro`.`dish`.`MID` = `mybistro`.`menu`.`MID` AND `mybistro`.`menu`.`BID` = `mybistro`.`bistro`.`BID` AND `mybistro`.`bistro`.`BID` = '" + BID + "'")
	if err != nil {
		panic(err.Error())
		jsonResponse, jsonError := json.Marshal("Loi lay data")
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
	} else {
		for rows.Next() {
			rows.Scan(&dish.DID, &dish.MID, &dish.Name, &dish.Price, &dish.Image)
			dishes = append(dishes, dish)
		}
		jsonResponse, jsonError := json.Marshal(dishes)
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

func addOrderDetail(response http.ResponseWriter, request *http.Request) {
	var (
		newDetailOrder OrderDetail
		ReturnData     string
	)

	newDetailOrder.OID = request.URL.Query().Get("oid")
	newDetailOrder.MID = request.URL.Query().Get("mid")
	newDetailOrder.DID = request.URL.Query().Get("did")
	newDetailOrder.Quantity = request.URL.Query().Get("quantity")
	newDetailOrder.Price = request.URL.Query().Get("price")

	_, err := db.Query("INSERT INTO `mybistro`.`detail_order_history` (`OrderID`,`MID`,`DID`,`Quantity`,`Price`) VALUES (" + newDetailOrder.OID + "," + newDetailOrder.MID + "," + newDetailOrder.DID + "," + newDetailOrder.Quantity + "," + newDetailOrder.Price + ")")
	if err != nil {
		panic(err.Error())
		ReturnData = "Them Order Detail that bai"
	} else {
		ReturnData = "Them Order Detail thanh cong"
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

func createCart(response http.ResponseWriter, request *http.Request) {
	var (
		NewCart    Cart
		ReturnData string
	)

	NewCart.UID = request.URL.Query().Get("uid")

	_, err := db.Query("INSERT INTO `mybistro`.`cart` (`UID`,`Status`) VALUES (" + NewCart.UID + ",'1')")

	if err != nil {
		panic(err.Error())
		ReturnData = "Add cart that bai"
	} else {
		ReturnData = "Add cart thanh cong"
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

func deleteCart(response http.ResponseWriter, request *http.Request) {
	var (
		ReturnData string
	)

	CID := request.URL.Query().Get("cid")

	_, err := db.Query("UPDATE `mybistro`.`cart` SET `Status`='0' WHERE `CID` = '" + CID + "'")

	if err != nil {
		panic(err.Error())
		ReturnData = "Xoa cart that bai"
	} else {
		ReturnData = "Xoa cart thanh cong"
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

func addDishInCart(response http.ResponseWriter, request *http.Request) {
	var (
		ReturnData string
	)

	did := request.URL.Query().Get("did")
	cid := request.URL.Query().Get("cid")

	_, err := db.Query("INSERT INTO `mybistro`.`dish_in_cart` (`DishID`,`CID`) VALUES ('" + did + "','" + cid + "')")

	if err != nil {
		panic(err.Error())
		ReturnData = "them dish vao cart that bai"
	} else {
		ReturnData = "them dish vao cart thanh cong"
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

func getDishByCartID(response http.ResponseWriter, request *http.Request) {
	var (
		dishes []Dish
		dish   Dish
	)
	cid := request.URL.Query().Get("cid")

	rows, err := db.Query("SELECT `mybistro`.`dish`.`DID`,`mybistro`.`dish`.`MID`,`mybistro`.`dish`.`Name`,`mybistro`.`dish`.`Price`,`mybistro`.`dish`.`Image` FROM `mybistro`.`dish_in_cart` dic, `mybistro`.`cart` crt,`mybistro`.`dish` dish WHERE CID=" + cid + ";")

	if err != nil {
		panic(err.Error())

	} else {
		for rows.Next() {
			rows.Scan(&dish.DID, &dish.MID, &dish.Name, &dish.Price, &dish.Image)
			dishes = append(dishes, dish)
		}
	}
	jsonResponse, jsonError := json.Marshal(dishes)
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

func getCartID(response http.ResponseWriter, request *http.Request) {
	var ReturnData string
	var storedCart Cart
	storedCart.UID = request.URL.Query().Get("uid")
	rows, err := db.Query("SELECT CartID,Status FROM `mybistro`.`cart` WHERE UID = " + storedCart.UID + ";")
	if err != nil {
		panic(err.Error())

	} else {
		for rows.Next() {
			rows.Scan(&storedCart.CID, &storedCart.Status)
		}
		ReturnData = storedCart.CID
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

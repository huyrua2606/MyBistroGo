package main

type (
	//User struct
	User struct {
		UID      string `json:"UserID"`
		username string `json:"Username"`
		password string `json:"Password"`
	}

	//Menu struct
	Menu struct {
		MID      string `json:"MenuID"`
		BID      string `json:"BistroID"`
		MenuName string `json:"MenuName"`
	}
	//Order struct
	Order struct {
		OrderID   string `json:"OrderID"`
		UID       string `json:"UserID"`
		BID       string `json:"BistroID"`
		Price     string `json:"Price"`
		BeginDate string `json:"BeginDate"`
		EndDate   string `json:"EndDate"`
		Status    string `json:"Status"`
	}

	//Ingredient struct
	Ingredient struct {
		IngreID     string `json:"IngredientID"`
		IngreName   string `json:"IngredientName"`
		MinQuantity string `json:"MinQuantity"`
		Measure     string `json:"Measure"`
	}

	//Dish struct
	Dish struct {
		MID   string `json:"MenuID"`
		DID   string `json:"DishID"`
		Name  string `json:"Name"`
		Price string `json:"Price"`
		Image string `json:"Image"`
	}

	//Bistro_Manager
	Bistro_Manager struct {
		UID   string `json:"UserID"`
		Name  string `json:"FullName"`
		BID   string `json:"BistroID"`
		Image string `json:"Image"`
	}

	//Bistro
	Bistro struct {
		BID        string `json:"BistroID"`
		Name       string `json:"BistroName"`
		Address    string `json:"Address"`
		PhoneNum   string `json:"PhoneNumber"`
		Email      string `json:"Email"`
		Status     string `json:"Status"`
		Rating     string `json:"Rating"`
		Reputation string `json:"Reputation"`
		Image      string `json:"Image"`
	}

	Customer struct {
		UID      string `json:UserID`
		FullName string `json:FullName`
		Email    string `json:Email`
		DoB      string `json:DoB`
		PhoneNum string `json:PhoneNumber`
		Image    string `json:"Image"`
	}

	OrderDetail struct {
		OID      string `json:OrderID`
		MID      string `json:MenuID`
		DID      string `json:DishID`
		Quantity string `json:Quantity`
		Price    string `json:Price`
	}

	Cart struct {
		CID    string `json:CartID`
		UID    string `json:UserID`
		Status string `json:DishID`
	}
)

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
	}

	//Ingredient struct
	Ingredient struct {
		IngreID     string `json:"IngredientID"`
		IngreName   string `json:"IngredientName"`
		MinQuantity string `json:"MinQuantity"`
		Measure     string `json:"Measure"`
	}
)

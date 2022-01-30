package meals

type meal struct {
	Item        string  `json:"item"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

var (
	mealItems = []meal{
		{
			Item:        "Sushi",
			Price:       22.99,
			Description: "Finest fish and veggies",
		},
		{
			Item:        "Schnitzel",
			Price:       16.50,
			Description: "A german specialty!",
		},
		{
			Item:        "Green Bowl",
			Price:       18.99,
			Description: "Healthy...and green...",
		},
	}
)

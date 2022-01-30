package meals

type meal struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

var (
	mealItems = []meal{
		{
			ID:          "m1",
			Name:        "Sushi",
			Price:       22.99,
			Description: "Finest fish and veggies",
		},
		{
			ID:          "m2",
			Name:        "Schnitzel",
			Price:       16.50,
			Description: "A german specialty!",
		},
		{
			ID:          "m3",
			Name:        "Green Bowl",
			Price:       18.99,
			Description: "Healthy...and green...",
		},
	}
)

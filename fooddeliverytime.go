package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	enteredFood := food{}
	switch order {
	case "burger":
		enteredFood.preptime = 15
	case "chips":
		enteredFood.preptime = 10
	case "nuggets":
		enteredFood.preptime = 12
	default:
		return 404
	}
	return enteredFood.preptime
}

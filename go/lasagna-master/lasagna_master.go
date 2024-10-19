package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerPrepTime int) int {
	if avgLayerPrepTime == 0 {
		avgLayerPrepTime = 2
	}
	return len(layers) * avgLayerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.
	noodles := 0
	for _, v := range layers {
		switch v {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		default:
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, nPortions int) []float64 {
	newAmounts := []float64{}
	for _, v := range amounts {
		newAmounts = append(newAmounts, v*(float64(nPortions)/2))
	}
	return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

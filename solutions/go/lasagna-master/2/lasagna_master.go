package lasagna

const (
    defaultPreparationTimePerLayer = 2
    noodleQtyPerLayer 			= 50	
    sauceQtyPerLayer 			= 0.2
    defaultServingInRecipe 		= 2 
)

// PreparationTime estimates the total preparation time based on the number of layers
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
    	avgPreparationTime = defaultPreparationTimePerLayer    
    }
    return len(layers) * avgPreparationTime
}

// Quantities determines the quantity of noodles and sauce needed to make your meal
func Quantities(layers []string) (int, float64) {
	var noodles int
    var sauce float64

    for _, v := range layers {
        switch v {
            case "noodles": noodles += noodleQtyPerLayer
            case "sauce": sauce += sauceQtyPerLayer
        }
    }

    return noodles, sauce
}

// AddSecretIngredient takes the secret ingredient from the friend recipe
func AddSecretIngredient(fromFriend, myRecipe []string) {
    myRecipe[len(myRecipe)-1] = fromFriend[len(fromFriend)-1]
}

// ScaleRecipe makes the magic happen
func ScaleRecipe(amountNeededTwoPortions []float64, numOfPortions int) []float64 {
    var res []float64
    for _, amount := range amountNeededTwoPortions {
        byOne := amount / defaultServingInRecipe
        res = append(res, byOne * float64(numOfPortions))
    }
    return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

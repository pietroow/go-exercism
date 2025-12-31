package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
    	avgPreparationTime = 2    
    }
    return len(layers) * avgPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
    var sauce float64
    
    for _, v := range layers {
        if v == "noodles" {
        	noodles += 50
        }
        if v == "sauce" {
        	sauce += 0.2    
        }
    }

    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fromFriend, myRecipe []string) {
    myRecipe[len(myRecipe)-1] = fromFriend[len(fromFriend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountNeededTwoPortions []float64, numOfPortions int) []float64 {
    var res []float64
    for _, amount := range amountNeededTwoPortions {
        byOne := amount / 2
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

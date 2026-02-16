package lasagna

func PreparationTime(layers []string, avgPrepTimeOfLayer int) int {
    if avgPrepTimeOfLayer == 0 {
        avgPrepTimeOfLayer = 2
    }
    return len(layers) * avgPrepTimeOfLayer
}

func Quantities(layers []string) (int, float64) {
    var sauceReq float64
    var noodReq int
    for _, val := range(layers) {
        if val == "noodles" {
            noodReq ++
        }
        if val == "sauce" {
            sauceReq ++
        }
    }
    return 50 * noodReq, 0.2 * sauceReq
}

func AddSecretIngredient(friendSlice, mySlice []string) {
    mySlice[len(mySlice)-1] = friendSlice[len(friendSlice)-1]
}

func ScaleRecipe(quantities []float64, noOfPortions int) []float64 {
    scaledSlice := make([]float64, len(quantities))
    if len(quantities) == 0 {
        return scaledSlice
    }
    for i := 0; i < len(quantities); i++ {
        scaledSlice[i] = float64(noOfPortions) * (quantities[i] / 2.0)
    }
    return scaledSlice
}



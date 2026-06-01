package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time_per_layer int) int {
    var num_of_layers int = len(layers)
    if time_per_layer == 0{
        return num_of_layers*2
    }else{
        return time_per_layer*num_of_layers
    }
}

// TODO: define the 'Quantities()' function
func Quantities( layers []string) (int, float64){
    var sauce_cnt int = 0
    var noodles_cnt int = 0
    for i := range layers{
        if layers[i] == "sauce"{
            sauce_cnt += 1
        }else if layers[i] == "noodles"{
            noodles_cnt += 1
        }
    }
    var sauce_qty float64 = float64(sauce_cnt) * 0.2
    var noodles_qty int = noodles_cnt * 50
    return noodles_qty,sauce_qty
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend_layers, layers []string){
    var secretIngredient = friend_layers[len(friend_layers) - 1]
    layers[len(layers) - 1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe( quantities []float64, portion int) []float64{
    var scaledQuantities = make([]float64, len(quantities))
    var scaledPortion float64 = float64(portion)/float64(2)
    for i,v := range quantities{
        scaledQuantities[i] = v * scaledPortion
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

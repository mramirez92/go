package lasagna

//PreparationTime() returns prep time for lasagna dependent on amount of layers
func PreparationTime(layers []string, minutes int) int{
    numLayers := len(layers)
    if minutes == 0 {
        return numLayers * 2
    }
	return numLayers * minutes
}

//Quantities() returns amount gramd of noodles and liters of sauce needed 
func Quantities(layers []string) (int, float64){
    noodles := 0
    sauce := 0.0
    for _, layer := range layers{
        if layer == "noodles"{
            noodles += 50
        }else if layer == "sauce"{
            sauce += 0.2
        }
    }
    return noodles, sauce
}

//AddSecretIngredient() modifies recipe and replaces "?" ingredient with newly discovered ingredient
func AddSecretIngredient(newRecipe, ownRecipe []string){
	ownRecipe[len(ownRecipe)-1] = newRecipe[len(newRecipe)-1]
}

//ScaleRecipe() scales amount of ingriedents needed based on portions
func ScaleRecipe(quantities []float64, portions int) []float64{
    quantitiesNeeded := []float64{}
    portionScale := float64(portions) / float64(2)
    for _, q:= range quantities{
        quantitiesNeeded = append(quantitiesNeeded, (q*portionScale))
    }
    return quantitiesNeeded
}

	// layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	// new := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "mushrooms"}



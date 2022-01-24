package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(ar []string, av int) int {
    if av == 0 {
        return len(ar)*2
    } else {
    return len(ar)*av
        }
}
// TODO: define the 'Quantities()' function
func Quantities(ar []string) (int, float64) {
    var n,s1 int
    var s float64
    for i:= range ar {
        if ar[i] == "noodles"{
            n = n+1
        } else if ar[i] == "sauce" {
        s1 = s1+1
        }
    }
s = float64(s1)*0.2
return n*50, s
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fr []string, my []string) {
    s := fr[len(fr)-1]
    my[len(my)-1] = s
    
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, n int) []float64{
    var ar[] float64
    var x float64
    for i := 0; i < len(quantities); i++ {
        x = quantities[i]
        x = x/2
        x = x * float64(n)
        ar = append(ar, x)
    }
return ar
}

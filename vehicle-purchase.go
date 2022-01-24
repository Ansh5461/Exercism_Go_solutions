package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind=="car" || kind=="truck" {
        return true
    } else {
    return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    var s string
	if option1 > option2 {
        s = option2 + " is clearly the better choice."
    } else {
    s = option1 + " is clearly the better choice."
    }
return s
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var p float64
    if age < 3 {
        p = originalPrice*0.8
    } else if age>3 && age < 10 {
        p = originalPrice * 0.7
    } else if age >= 10 {
         p=originalPrice * 0.5
     }
return p  
}

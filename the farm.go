package thefarm

import (
    "fmt"
    "errors"
    )
// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    cow int
}
func(a *SillyNephewError) Error() string {
    var s string 
    s = fmt.Sprintf("silly nephew, there cannot be %d cows", a.cow)
    return s
}


// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    v,e := weightFodder.FodderAmount()
    if e == ErrScaleMalfunction {
        v = v*2
    } else if e!= nil {
        return 0.0, e    
    }
	var res float64
    
    if v < 0 {
        return 0.0, errors.New("negative fodder")
    }
	if cows == 0 {
        return 0.0, errors.New("division by zero")
    }
	if cows < 0 {
        s := SillyNephewError{
            cow : cows,
        }
        return 0.0,errors.New(s.Error())
    }
    res = v/float64(cows)
    return res,nil
}

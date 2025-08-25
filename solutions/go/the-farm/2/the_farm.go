package thefarm
import "errors"
import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    fodAmt, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    factor ,err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return float64((fodAmt * factor) / float64(cows)), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
        return DivideFood(fc, cows)
    } else {
        return 0, errors.New("invalid number of cows")
    }
    
}

type InvalidCowsError struct {
    noOfCows int
    msg string
}

func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.noOfCows, e.msg)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
        return &InvalidCowsError{noOfCows: cows, msg: "there are no negative cows"}
    }
    if cows == 0 {
        return &InvalidCowsError{noOfCows: cows, msg: "no cows don't need food"}
    }
    return nil
}
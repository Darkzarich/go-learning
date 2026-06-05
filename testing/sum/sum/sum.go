package sum

import "errors"

func Sum(a, b int) (int, error) {
    if a == 0 || b ==  0 {
        return 0, errors.New("arg is zero")
    }
    
    if a < 0 || b < 0 {
        return 0, errors.New("arg is negative")
    } 
    return a + b, nil
}


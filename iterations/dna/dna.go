package hamming

import "errors"

func Distance(a, b string) (distance int, e error) {
    if len(a) == len(b){
        for i := 0; i < len(a); i++{
            if a[i] != b[i]{
                distance ++ 
            }
        }
    	return distance, nil
    }
	return -1, errors.New("strings are different sizes, can't compare")
}
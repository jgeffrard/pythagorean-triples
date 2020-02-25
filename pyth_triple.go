package main 

import "fmt"
import "math"

func pyth_triple() {
	a_coef := 0
	b_coef := 0
	c_coef := 0

	a_squared := 0
	b_squared := 0
	c_squared := 0

	_, _, _, _, _, _ = a_coef, b_coef, c_coef, a_squared, b_squared, c_squared

	var input float64
	fmt.Println("Please enter the number up to which you would like to see the Pythagorean triples: ")
	fmt.Scan(&input)

	for a := float64(0); a <= input; a++ {
		a_coef := a
		a_squared := math.Pow(a, 2)
		for b := float64(0); b <= input; b++ {
			b_coef := b
			b_squared := math.Pow(b, 2)
			for c := float64(0); c <= input; c++ {
				c_coef := c
				c_squared := math.Pow(c, 2)
				
				if a_squared + b_squared == c_squared {
					fmt.Printf("The triple (%v, %v, %v) is a Pythagorean triple.\n", a_coef, b_coef, c_coef)
				} else { 
					fmt.Printf("The triple (%v, %v, %v) is not a Pythagorean triple.\n", a_coef, b_coef, c_coef)
				}				
			}
		}
	}
}

func main() {   
    pyth_triple() 
} 
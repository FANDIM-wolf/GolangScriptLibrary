package main 

import (
	"fmt"
)
const (

	Pi=3.4
)

//func Square_of_ball(radius){
	//var result int = 0
	//result = 4*radius * radius *Pi
	//Printf(result)
//}

type Figure struct{
    radius int
    amount_of_sides int
    square int 
    location []int
}

func main(){
	//Square_of_ball(23)
	Ball := Figure{
		radius:12,
		amount_of_sides: 1,
		square: 576,
		location : []int{
			1,
			2,
			3,
		},
		
	}
	fmt.Println(Ball)
}

package main

import (
	"fmt"
)

func main(){
	var min,max int;
	fmt.Print("MIN:")
	fmt.Scan(&min)
	fmt.Print("MAX:")
	fmt.Scan(&max)
	for ;max!=min;{
		if max>min{
			max=max-min
		}else{
			min=min-max
		}
	}
	fmt.Print("ANSWER:",min)
}

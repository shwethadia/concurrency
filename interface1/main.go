package main

import "fmt"


type Dense interface{

		 Density() float64
}

type Metal struct{

		mass float64
		volume float64
}


type Gas struct{

		pressure float64
		temperature float64
		molecularMass float64

}

func (g *Gas) Density() float64 {

		//var density float64

		density := (g.molecularMass * g.pressure) /(0.821 * (g.temperature + 273))
		return density
}

func (m *Metal) Density() float64{

		return m.mass / m.volume
}


func IsDenser(a,b Dense) bool{

	 return a.Density() > b.Density()
}

func main(){


	gold := Metal{478,24}

	silevr := Metal{100,10}


	result := IsDenser(&gold,&silevr)

	if result{
		fmt.Println("Gold is higher density than silver")
	}else{
		fmt.Println("Silver has higher density than gold")
	}


	oxygen := Gas{
		pressure: 5,
		temperature: 27,
		molecularMass: 32,
	}


	hydrogen:= Gas{

		pressure: 1,
		temperature: 0,
		molecularMass: 2,
	}

	resultGas := IsDenser(&oxygen,&hydrogen)

	if resultGas{
		fmt.Println("oxygen is higher density than hydrogen")
	}else{
		fmt.Println("hydrogen has higher density than oxygen")
	}

}
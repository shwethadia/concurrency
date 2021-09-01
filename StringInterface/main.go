package main

import "fmt"




type user struct{

		name string
		email string

}



func (u user) String() string{

		 return fmt.Sprintf("%s <%s>",u.name,u.email)

}

func main(){

	u:= user{

			name:"shwetha",
			email:"shwetha@gmail.com",
	}

	fmt.Println(u)


}
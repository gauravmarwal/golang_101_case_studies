package main

import "fmt"

type Service interface {
	SayHi()
}

type MyService struct{}

func (s MyService) SayHi() {
	fmt.Println("Hi")
}

type SecondService struct{}

func (s SecondService) SayHi() {
	fmt.Println("Hello From the 2nd Service")
}

func main() {

	interfaceMap := make(map[string]Service)

	interfaceMap["SERVICE_1"] = MyService{}
	interfaceMap["SERVICE_2"] = SecondService{}

	interfaceMap["SERVICE_2"].SayHi()

}

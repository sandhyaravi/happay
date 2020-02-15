package main

import (
	"fmt"
)
type Sample struct{
	Name string `json:"name"`
	Id int `json:"id"`
}
func Writer(data []Sample){
	ch1 := make(chan Sample)
	//time.Sleep(2)
	fmt.Println("writer",data)
	for _,val := range data{
		ch1 <-val
	}
	fmt.Println("return channel:",ch1)

}
func Listener(data []Sample){
	//time.Sleep(2)
	ch2 := make(chan struct{})
	<-ch2
	fmt.Println("listener",<-ch2)
}
func main(){

	var test []Sample
	test = append(test, Sample{"happay",3})
	test = append(test,Sample{"happay1",2})
	test = append(test,Sample{"wipro",1})
	var ch = make(chan struct{}) //wg := sync.WaitGroup{}
	//wg.Add(2)
	go Writer(test)

	data := <-ch
	fmt.Println(data)
	go Listener(test)
	<-ch
}
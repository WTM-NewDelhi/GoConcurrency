package main

import "fmt"

func awesome( statement string, c chan string) {
	
	for i := 0;i<10; i++ {

		c <- fmt.Sprintf("%s %d", statement, i)
	}
}

func main( ) {

	c := make(chan string)

	go awesome( " We are awesome ", c)

	fmt.Println( " I don't have to wait for func awesome ")

	for i := 0;i<10; i++ {
		fmt.Println(<-c)
	}	

	fmt.Println( " You are awesome enough, bye! ")

}

package main

import "fmt"
import "time"

func awesome( statement string) {

	for i := 0;i<10; i++ {

		fmt.Println( statement, i )
	}
}
func main( ) {

	go awesome( " We are awesome ")

	fmt.Println( " I don't have to wait for func awesome ")

	time.Sleep( 2*time.Second )

	fmt.Println( " You are awesome enough, bye! ")

}


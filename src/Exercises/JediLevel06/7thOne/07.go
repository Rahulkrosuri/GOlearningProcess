package main
import "fmt"
func main(){
	x := func(){
			fmt.Println("This is an anonymous func assigned to a variable")
		}
	x()
}
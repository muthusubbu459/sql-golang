package main

import("regexp"
	"fmt"
)
func isip(ip string)(value bool) {
	m,_:=regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip)
	return m

}
func main(){
	var ip string
	fmt.Println("Please enter the ip address below in format similar to 192.0.0.1 ")
	fmt.Scanf("%s\t",&ip)
	value:=isip(ip)
	fmt.Println("the ip address is",value)
}

package main
import("regexp"
	"fmt"

)
func digits(newno string)bool{
	var no,err = regexp.Compile(`^[0][1-4]+[0-9]{3}|[1-4]+[0-9]{3}$`)
	if err != nil {
		panic(err)
	}

		result:=no.MatchString(newno)
		//fmt.Printf("%v=%v\n",newno,result)
		return result

}
func main() {

	 newno:=[]string{"04655","1000","5566","ab000","2433"}
	for i:=0;i<len(newno);i++  {
		value:=digits(newno[i])
		fmt.Printf("%v=%v\n",newno[i],value)
	}
	//value:=digits(newno)
	/*for i:= range newno{
		result:=no.MatchString(newno[i])
		fmt.Printf("%v=%v\n",newno[i],result)
	}*/
	/*fmt.Println("enter the number between 0 and 5000 in five digits")
	fmt.Scanf("%s/t",&newno)

	m:= no.MatchString(newno)
	fmt.Println(m)
	if ! m {
		fmt.Println(m)

	}*/
}

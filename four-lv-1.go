package main

import(
	"fmt"
	"time"
)

func main() {

	fmt.Println("input")
	var str [5] int64
	for i:=0;i<len(str);i++{
		fmt.Scanf("%d\n",&str[i])
	}

	for a:=0;a<len(str);a++{
		if str[a] !='0'{
			fmt.Println("input ok!")
		}
	}

fmt.Println("the result are :")
	for j:=0;j<len(str);j++{
		b:=str[j]
		a:=time.Unix(b,1)
		fmt.Println(a)
	}

}

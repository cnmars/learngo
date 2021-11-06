package main
import "fmt"
import "time"

func main(){
	str := "golang你好"
	fmt.Println(len(str))
	//遍
	for i , value := range str {
		fmt.Printf("索引为：%d,具体值为：%c\n",i,value)
		time.Sleep(2000 * time.Millisecond)
	}
}
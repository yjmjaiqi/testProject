package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome exit page!")
	//sum := 0
	//for i := 0; i <= 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	num := rand.Intn(100)
	var ran int
	var a int
	for {
		if a <= 100 {
			fmt.Println("请输入随机数0~100")
			//当程序执行到 fmt.Println("请输入姓名：") ，会阻塞到这里，等待用户输入
			fmt.Scanln(&ran)
			if ran < num {
				fmt.Println("猜小了，嘻嘻，再试试！")
			} else if ran > num {
				fmt.Println("猜大了，哈哈，再试试！")
			} else if ran == num {
				fmt.Println("猜中了，你真棒！")
				break
			}
		}
	}
}

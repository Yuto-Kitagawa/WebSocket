package main

//ターミナル画面表示はfmtパッケージを使う
import "fmt"

type user struct {
	name string
	old  int
}

func main() {
	fmt.Println("1.Hello, playground")

	// const lightspeed = 299792
	// var distance = 56000000
	// var distance, speed = 56000000, 299792
	// fmt.Println(distance/speed, "秒")
	name := "yuto"
	name = getHelloMessage(name)
	fmt.Println("2.", name)

	var a [6]int
	a[2] = 10
	fmt.Println(a[2])
	b := []int{1, 2, 3}          //[0][1][2]
	c := [...]int{4, 5, 6, 7, 8} //not decide
	fmt.Println("b[2]:", b[2])
	fmt.Println("c[2]:", c[2])
	fmt.Println("len(c):", len(c))
	fmt.Println("cap(c):", cap(c))
	s3 := append(b, 8, 2, 10)
	fmt.Println("s3",s3)

	n := 3
	for n < 10 {
		n++
		fmt.Println(n)
	}

	//たぶんi はインデックス
	//bがvに入る
	//bの数だけ回る。拡張for文みたいな感じ
	for i, v := range b {
		fmt.Println(i, v)
	}

	u1 := new(user)
	u1.name = "KITAGAWA"
	u1.old = 20
	fmt.Println("u1:", u1)
}

//名前付きの引数がある関数
func getHelloMessage(name string) (msg string) {
	msg = "Hello " + name
	//retern (msg string)
	return
}

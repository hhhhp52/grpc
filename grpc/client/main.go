package main

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	protobuf "../protobuf"
)

func main() {
	// 連線到遠端 gRPC 伺服器。
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	// 建立新的 Calculator 客戶端，所以等一下就能夠使用 Calculator 的所有方法。
	c := protobuf.NewCalculatorClient(conn)

	fmt.Println("請輸入第一個數字")
	var num1Input int
	fmt.Scan(&num1Input)
	s1 := strconv.Itoa(num1Input)
	num1, _ := strconv.ParseInt(s1, 10, 64)

	fmt.Println("請輸入第二個數字")
	var num2Input int
	fmt.Scan(&num2Input)
	s2 := strconv.Itoa(num2Input)
	num2, _ := strconv.ParseInt(s2, 10, 64)

	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
	r, err := c.Plus(context.Background(), &protobuf.CalcRequest{NumberA: num1, NumberB: num2})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%d", r.Result)
}

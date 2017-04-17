package main

import(
	"net/http"
	"fmt"
	"flag"
	"strconv"
)

var(
	port = flag.Int("p", 80, "Setting Bind Port")
	help = flag.Bool("h", false, "Display Help")
)

func main(){
	// フラグを取得
	flag.Parse()
	args := flag.Args()
	// ヘルプを出力
	if *help || len(args) != 1{
		fmt.Println("textserve [-p=port] TEXT")
		flag.PrintDefaults()
		return
	}
	
	// 受信処理を定義
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(args[0]))
	})
	
	// 受信待機を開始
	fmt.Println("Binding Port =", *port)
	fmt.Println(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}

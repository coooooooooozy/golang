package main

import (
	"fmt"

	"github.com/DannyBen/quandl"
)

func main() {
	quandl.APIKey = "5q9xXAsFNZbL9-s9xoug"

	// v := quandl.Options{} //Option
	//Set registers a key=value pair to be sent in the Quandl request
	// v.Set("trim_start", "2017-01-01")
	// v.Set("trim_end", "2017-02-01")
	//    v.Set("column_index", "4")

	data, _ := quandl.GetSymbol("EOD/AAPL", nil) //第１引数に指定した銘柄のシンボルを取得
	//    data2, _ := quandl.GetSymbolRaw("WIKI/AAPL", "json", v) //第１引数に指定した銘柄のシンボルをjson形式で取得

	for i, item := range data.Data {
		fmt.Println(i, item)
	}

	//ioutil.WriteFileで保存
}

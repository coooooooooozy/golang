package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://alpha-vantage.p.rapidapi.com/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT&outputsize=compact&datatype=json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "f4f75b1a5dmsh465732335f3d394p122bd3jsna4487e467386")
	req.Header.Add("x-rapidapi-host", "alpha-vantage.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

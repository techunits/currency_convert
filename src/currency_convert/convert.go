package currency_convert

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
    "log"
)

func fetchRate(source string, target string) (rate string) {
	endpoint := "http://rate-exchange.herokuapp.com/fetchRate?from=SOURCE_CURRENCY&to=TARGET_CURRENCY"
	endpoint = strings.Replace(endpoint, "SOURCE_CURRENCY", source, 1)
	endpoint = strings.Replace(endpoint, "TARGET_CURRENCY", target, 1)
	//	log.Println(endpoint)
	
	resp, err := http.Get(endpoint)
	if err != nil {
		// handle errors
		log.Println("%s\n", err)
	}
	defer resp.Body.Close()
	
	var jsonStream interface {}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &jsonStream)
	
	rate = fmt.Sprintf("%s", jsonStream.(map[string]interface{})["Rate"])
	
	return rate
}

func Convert(source string, target string, amount float64) (finalAmount float64) {
	rate := fetchRate(source, target);
	frate, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		log.Println(err)
	}
	
	finalAmount = frate * amount
	return finalAmount
}

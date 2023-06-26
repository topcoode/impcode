// must be create in the package
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type (
	Client struct {
		host       string
		httpClient *http.Client
		apiKey     string
		apiSecret  string
	}
) //create the ----------> client struct for function reciever   -----> API function
func NewClient(host string, apiKey, apiSecret string, timeout time.Duration) *Client {   //creating the new client.
	client := &http.Client{                                                              //input:= host,apikey,apisecret,time
		Timeout: timeout,
	}
	return &Client{
		host:       host,
		httpClient: client,
		apiKey:     apiKey,
		apiSecret:  apiSecret,  
	}
}

// ---------------------------------------------------------------------------------------------------->
func (c *Client) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {            //method- http method(post,get,put.....)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

// ----------------------------------------------------------------------------------------------------->
const (
	tickerURL = "api/v3/ticker/price"
)

type (
	TickerResponse struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
)

func (c *Client) GetTickers() (resp []TickerResponse, err error) {
	res, err := c.do(http.MethodGet, tickerURL, nil)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}
	return
}

// ---------------------------------------------------------------------->
func main() {
	defaultTimeout = time.Second * 10
	client := binance.NewClient("https://api.binance.com", os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET"), defaultTimeout)
	//newclient----------->
	
	tickers, err := client.GetTickers()
	//data retrieve --------------------->
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tickers)
}

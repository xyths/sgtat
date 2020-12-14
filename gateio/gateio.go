package gateio

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const DataSource = "http://data.gateio.co/api2/1"

type GateIO struct {
	Key    string
	Secret string
}

func NewGateIO(key, secret string) GateIO {
	return GateIO{Key: key, Secret: secret}
}

// all support pairs
func (g *GateIO) getPairs() string {
	var method string = "GET"
	var url string = DataSource + "/pairs"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Market Info
func (g *GateIO) marketinfo() string {
	var method string = "GET"
	var url string = DataSource + "/marketinfo"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Market Details
func (g *GateIO) marketlist() string {
	var method string = "GET"
	var url string = DataSource + "/marketlist"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// tickers
func (g *GateIO) tickers() string {
	var method string = "GET"
	var url string = DataSource + "/tickers"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// ticker
func (g *GateIO) ticker(ticker string) string {
	var method string = "GET"
	var url string = DataSource + "/ticker" + "/" + ticker
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Depth
func (g *GateIO) orderBooks() string {
	var method string = "GET"
	var url string = DataSource + "/orderBooks"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Depth of pair
func (g *GateIO) orderBook(params string) string {
	var method string = "GET"
	var url string = DataSource + "/orderBook/" + params
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Trade History
func (g *GateIO) tradeHistory(params string) string {
	var method string = "GET"
	var url string = DataSource + "/tradeHistory/" + params
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get account fund balances
func (g *GateIO) balances() string {
	var method string = "POST"
	var url string = DataSource + "/private/balances"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// get deposit address
func (g *GateIO) depositAddress(currency string) string {
	var method string = "POST"
	var url string = DataSource + "/private/depositAddress"
	var param string = "currency=" + currency
	var ret string = g.httpDo(method, url, param)
	return ret
}

// get deposit withdrawal history
func (g *GateIO) depositsWithdrawals(start string, end string) string {
	var method string = "POST"
	var url string = DataSource + "/private/depositsWithdrawals"
	var param string = "start=" + start + "&end=" + end
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Place order buy
func (g *GateIO) buy(currencyPair string, rate string, amount string) string {
	var method string = "POST"
	var url string = DataSource + "/private/buy"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Place order sell
func (g *GateIO) sell(currencyPair string, rate string, amount string) string {
	var method string = "POST"
	var url string = DataSource + "/private/sell"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Cancel order
func (g *GateIO) cancelOrder(orderNumber string, currencyPair string) string {
	var method string = "POST"
	var url string = DataSource + "/private/cancelOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Cancel all orders
func (g *GateIO) cancelAllOrders(types string, currencyPair string) string {
	var method string = "POST"
	var url string = DataSource + "/private/cancelAllOrders"
	var param string = "type=" + types + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get order status
func (g *GateIO) getOrder(orderNumber string, currencyPair string) string {
	var method string = "POST"
	var url string = DataSource + "/private/getOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get my open order list
func (g *GateIO) openOrders() string {
	var method string = "POST"
	var url string = DataSource + "/private/openOrders"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// 获取我的24小时内成交记录
func (g *GateIO) myTradeHistory(currencyPair string, orderNumber string) string {
	var method string = "POST"
	var url string = DataSource + "/private/tradeHistory"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get my last 24h trades
func (g *GateIO) withdraw(currency string, amount string, address string) string {
	var method string = "POST"
	var url string = DataSource + "/private/withdraw"
	var param string = "currency=" + currency + "&amount=" + amount + "&address=" + address
	var ret string = g.httpDo(method, url, param)
	return ret
}

func (g *GateIO) getSign(params string) string {
	key := []byte(g.Secret)
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(params))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

/**
*  http request
*/
func (g *GateIO) httpDo(method string, url string, param string) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		// handle error
	}
	var sign string = g.getSign(param)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", g.Key)
	req.Header.Set("sign", sign)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body);
}

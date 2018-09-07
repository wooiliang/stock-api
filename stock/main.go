package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

const (
	selectorKLSE = "table#stockhdr > tbody > tr:last-child > td:first-child"
	selectorSGX  = "div.stockinfocol1row1 span.value"
)

const (
	urlKLSE = "https://klse.i3investor.com/servlets/stk/%s.jsp"
	urlSGX  = "https://sginvestors.io/sgx/%s/stock-info"
)

// ResponseSuccess struct
type ResponseSuccess struct {
	Price float64 `json:"price"`
}

// ResponseError struct
type ResponseError struct {
	Message string `json:"message"`
}

func scrap(selector string, url string) (float64, error) {
	c := colly.NewCollector()
	var response string
	ch := make(chan error, 1)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		response = e.Text
	})
	c.OnError(func(_ *colly.Response, err error) {
		ch <- err
	})
	c.OnScraped(func(r *colly.Response) {
		if len(response) > 0 {
			ch <- nil
		} else {
			ch <- errors.New("Price Selector Not Found")
		}
	})
	c.Visit(url)
	if err := <-ch; err == nil {
		return strconv.ParseFloat(strings.Replace(response, "\u00a0", "", -1), 64)
	} else {
		return 0, err
	}
}

func formatTicker(ticker string) string {
	return strings.Replace(ticker, "_", "/", 1)
}

func getPrice(market string, ticker string) (float64, error) {
	switch market {
	case "klse":
		return scrap(selectorKLSE, fmt.Sprintf(urlKLSE, ticker))
	case "sgx":
		return scrap(selectorSGX, fmt.Sprintf(urlSGX, formatTicker(ticker)))
	}
	return 0, nil
}

// Handler function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received pathParams: ", request.PathParameters)
	if price, err := getPrice(request.PathParameters["market"], request.PathParameters["ticker"]); err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, err
	} else {
		result, err := json.Marshal(&ResponseSuccess{price})
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, err
		}
		return events.APIGatewayProxyResponse{Body: string(result), StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(Handler)
}

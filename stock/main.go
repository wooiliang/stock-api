package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

const (
	selectorKLSE = "td[class='up big16']"
	selectorSGX  = "div.stockinfocol1row1 span[class='value red']"
)

const (
	urlKLSE = "https://klse.i3investor.com/servlets/stk/%s.jsp"
	urlSGX  = "https://sginvestors.io/sgx/%s/stock-info"
)

func scrap(selector string, url string) (float64, error) {
	c := colly.NewCollector()
	ch := make(chan string, 1)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		ch <- e.Text
	})
	c.Visit(url)
	result := <-ch
	return strconv.ParseFloat(strings.Replace(result, "\u00a0", "", -1), 64)
}

func getPrice(market string, ticker string) (float64, error) {
	switch market {
	case "klse":
		return scrap(selectorKLSE, fmt.Sprintf(urlKLSE, ticker))
	case "sgx":
		return scrap(selectorSGX, fmt.Sprintf(urlSGX, ticker))
	}
	return 0, nil
}

// Handler function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	if price, err := getPrice("klse", "1155"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(price)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

package main

import "testing"

func TestGetPrice(t *testing.T) {
	if _, err := getPrice("sgx", "stock_c6l-sia"); err != nil {
		t.Errorf(`Expected price, got "%v"`, err)
	}
}
func TestFormatTicker(t *testing.T) {
	if output := formatTicker("stock_c6l-sia"); output != "stock/c6l-sia" {
		t.Errorf(`Expected "stock/c6l-sia", got "%v"`, output)
	}
}

package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestInvalidTransactions(_t *testing.T) {
	i, o := []string{"alice,20,800,mtv", "alice,50,100,beijing"}, []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	T(_t, S(i), S(Sort(invalidTransactions(i))), S(Sort(o)))

	i, o = []string{"alice,20,800,mtv", "alice,50,1200,mtv"}, []string{"alice,50,1200,mtv"}
	T(_t, S(i), S(Sort(invalidTransactions(i))), S(Sort(o)))

	i, o = []string{"alice,20,800,mtv", "bob,50,1200,mtv"}, []string{"bob,50,1200,mtv"}
	T(_t, S(i), S(Sort(invalidTransactions(i))), S(Sort(o)))

	i, o = []string{"xnova,261,1949,chicago", "bob,206,1284,chicago", "xnova,420,996,bangkok", "chalicefy,704,1269,chicago", "iris,124,329,bangkok", "xnova,791,700,amsterdam", "chalicefy,572,697,budapest", "chalicefy,231,310,chicago", "chalicefy,763,857,chicago", "maybe,837,198,amsterdam", "lee,99,940,bangkok", "bob,132,1219,barcelona", "lee,69,857,barcelona", "lee,607,275,budapest", "chalicefy,709,1171,amsterdam"}, []string{"xnova,261,1949,chicago", "bob,206,1284,chicago", "chalicefy,704,1269,chicago", "chalicefy,763,857,chicago", "lee,99,940,bangkok", "bob,132,1219,barcelona", "lee,69,857,barcelona", "chalicefy,709,1171,amsterdam"}
	T(_t, S(i), S(Sort(invalidTransactions(i))), S(Sort(o)))
}

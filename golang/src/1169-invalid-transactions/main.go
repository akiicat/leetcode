package main

import (
	"sort"
	"strconv"
	"strings"
)

// T: O(n * log(n))
// M: O(n)
// -- start --

type Transaction struct {
	Id      int
	Name    string
	Time    int
	Amount  int
	City    string
	Invalid bool
}

func invalidTransactions(transactions []string) []string {
	m := map[string][]Transaction{}

	for i, v := range transactions {
		s := strings.Split(v, ",")

		time, _ := strconv.Atoi(s[1])
		amt, _ := strconv.Atoi(s[2])

		m[s[0]] = append(m[s[0]], Transaction{
			Id:      i,
			Name:    s[0],
			Time:    time,
			Amount:  amt,
			City:    s[3],
			Invalid: amt >= 1000,
		})
	}

	res := []string{}

	for _, arr := range m {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Time < arr[j].Time
		})

		l := 0
		for r, _ := range arr {
			for arr[r].Time-arr[l].Time > 60 {
				l++
			}

			for i := l; i < r; i++ {
				if arr[i].Name == arr[r].Name && arr[i].City != arr[r].City {
					arr[i].Invalid = true
					arr[r].Invalid = true
				}
			}
		}

		for _, v := range arr {
			if v.Invalid {
				res = append(res, transactions[v.Id])
			}
		}
	}

	return res
}

// -- end --

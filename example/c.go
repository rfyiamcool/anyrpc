package main

import (
	"fmt"
	"github.com/rfyiamcool/anyrpc"
)

func main() {
	var Add func(int64, int64) (int64, error)
	var SubError func(int64, int64) (int64, error)

	s := anyrpc.NewClient("tcp", "127.0.0.1:7678", 10)
	setLogger()
	err := s.MakeRpc("add", &Add)
	if err != nil {
		panic(err.Error())
	}

	err = s.MakeRpc("subError", &SubError)
	if err != nil {
		panic(err.Error())
	}

	for index := 0; index < 10; index++ {
		a, err := Add(5, 5)
		fmt.Println(a, err)

		b, err := SubError(10, 2)
		fmt.Println(b, err)
	}
}

func setLogger() {
	anyrpc.SetErrorLogger(func(tmpl string, s ...interface{}) {
		if len(s) == 0 {
			fmt.Println(tmpl)
			return
		}

		fmt.Printf(tmpl, s...)
	},
	)

	anyrpc.SetLogger(func(tmpl string, s ...interface{}) {
		if len(s) == 0 {
			fmt.Println(tmpl)
			return
		}

		fmt.Printf(tmpl, s...)
	},
	)
}

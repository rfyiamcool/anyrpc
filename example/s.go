package main

import (
	"errors"
	"fmt"
	"github.com/rfyiamcool/anyrpc"
)

type Service struct {
}

func (serv *Service) Add(a, b int64) (int64, error) {
	fmt.Println("call add")
	return a + b, nil
}

func (serv *Service) AddError(a, b int64) (int64, error) {
	return a + b, errors.New("active raise error")
}

func (serv *Service) Sub(a, b int64) (int64, error) {
	return a - b, nil
}

func (serv *Service) SubError(a, b int64) (int64, error) {
	return a - b, errors.New("active raise error")
}

func main() {
	s := anyrpc.NewServer("tcp", "127.0.0.1:7678")
	serv := Service{}
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
	s.Register("add", serv.Add)
	s.Register("subError", serv.SubError)
	s.Start()
}

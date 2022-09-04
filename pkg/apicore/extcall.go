package apicore

import "fmt"

type extCall func()

var extraCalls []extCall

func (a *app) initExtra() {
	fmt.Println("Init extra")
	for _, init := range extraCalls {
		init()
	}
}

func RegisterExtrCall(e extCall) {
	extraCalls = append(extraCalls, e)
}

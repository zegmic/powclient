package main

import (
	"fmt"
	"github.com/zegmic/powclient/pkg/client"
	"github.com/zegmic/powclient/pkg/solver"
)

func main() {
	challenge := client.GetChallenge()
	fmt.Println("got challange: ", challenge)
	nonce := solver.Solve(challenge)
	data := client.GetQuote(challenge, nonce)
	fmt.Println("got server response: ", data)
}

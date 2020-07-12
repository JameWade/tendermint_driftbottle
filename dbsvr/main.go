package main

import (
	"fmt"

	abciserver "github.com/tendermint/tendermint/abci/server"
)

func main() {

	//var t *testing.T
	app := NewdriftBottleApplication()

	//var app tmtypes.Application
	addr := "localhost:26658"
	transport := "socket"
	svr, err := abciserver.NewServer(addr, transport, app)
	if err != nil {
		return
	}
	svr.Start()
	defer svr.Stop()
	fmt.Println("abci server started.")
	select {}
	//阻塞
}

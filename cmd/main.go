package main

import (
	//	"Challenge/config"
	route "Challenge/internal/adapters/api"
	//	Stream "Challenge/internal/adapters/stream"
	//	Consumer "Challenge/internal/repositories/transactions"
	kafka "Challenge/internal/repositories/transactions"
)

func main() {

	go kafka.Consume()
	route.HandleRequest()

}

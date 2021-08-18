package main

import (
	"intra_payment/router"
)

func main()  {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":3002"))
}

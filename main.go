package main

import (
	"fmt"
	"log"
	"net/http"

	"market-crawl/internal/application"
	"market-crawl/internal/infrastructure"
	"market-crawl/internal/interfaces"
)

func main() {
	// Infrastructure layer
	icbcClient := infrastructure.NewICBCClient()
	cxxxxClient := infrastructure.NewCxxxxClient()
	bocClient := infrastructure.NewBocClient()

	// Application layer
	netValueService := application.NewNetValueService(icbcClient)
	cxxxxProductNavService := application.NewCxxxxProductNavService(cxxxxClient)
	bocRevenueService := application.NewBocRevenueService(bocClient)

	// Interfaces layer
	netValueHandler := interfaces.NewNetValueHandler(netValueService)
	cxxxxProductNavHandler := interfaces.NewCxxxxProductNavHandler(cxxxxProductNavService)
	bocRevenueHandler := interfaces.NewBocRevenueHandler(bocRevenueService)

	// Register routes
	http.HandleFunc("/api/net-value-list", netValueHandler.GetNetValueList)
	http.HandleFunc("/api/citic-product-nav", cxxxxProductNavHandler.GetProductNav)
	http.HandleFunc("/api/boc-revenue-list", bocRevenueHandler.GetRevenueList)

	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

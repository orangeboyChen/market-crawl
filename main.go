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
	citicClient := infrastructure.NewCiticClient()
	bocClient := infrastructure.NewBocClient()

	// Application layer
	netValueService := application.NewNetValueService(icbcClient)
	citicProductNavService := application.NewCiticProductNavService(citicClient)
	bocRevenueService := application.NewBocRevenueService(bocClient)

	// Interfaces layer
	netValueHandler := interfaces.NewNetValueHandler(netValueService)
	citicProductNavHandler := interfaces.NewCiticProductNavHandler(citicProductNavService)
	bocRevenueHandler := interfaces.NewBocRevenueHandler(bocRevenueService)

	// Register routes
	http.HandleFunc("/api/net-value-list", netValueHandler.GetNetValueList)
	http.HandleFunc("/api/citic-product-nav", citicProductNavHandler.GetProductNav)
	http.HandleFunc("/api/boc-revenue-list", bocRevenueHandler.GetRevenueList)

	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

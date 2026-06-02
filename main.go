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

	// Application layer
	netValueService := application.NewNetValueService(icbcClient)
	citicProductNavService := application.NewCiticProductNavService(citicClient)

	// Interfaces layer
	netValueHandler := interfaces.NewNetValueHandler(netValueService)
	citicProductNavHandler := interfaces.NewCiticProductNavHandler(citicProductNavService)

	// Register routes
	http.HandleFunc("/api/net-value-list", netValueHandler.GetNetValueList)
	http.HandleFunc("/api/citic-product-nav", citicProductNavHandler.GetProductNav)

	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

package main

import (
	"fmt"
	"github.com/hurbcom/google-ads-go/ads"
	"github.com/hurbcom/google-ads-go/services"
	"log"
)

func main() {
	// Create a client from credentials file
	client, err := ads.NewClientFromStorage("credentials.json")
	if err != nil {
		panic(err)
	}

	// Load the "GoogleAds" service
	googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())

	// Create a search request
	request := services.SearchGoogleAdsRequest{
		CustomerId: "12345678",
		Query:      "SELECT segments.date, click_view.gclid FROM click_view WHERE segments.date = '2021-07-23'",
	}

	// Get the results
	response, err := googleAdsService.Search(client.Context(), &request)
	if err != nil {
		log.Fatal(err)

	}

	// Printing results
	for _, row := range response.Results {
		fmt.Println(row.ClickView)
	}
}

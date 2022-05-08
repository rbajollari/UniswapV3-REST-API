package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rbajollari/UniswapV3-REST-API/api/models"

	"github.com/machinebox/graphql"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: /")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Uniswap V3 Rest API")
}

func TokenPools(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: /tokenpools")

	w.Header().Add("Content-Type", "application/json")

	// create Uniswap Graph query from api request
	requestJSON := models.TokenPoolsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestJSON); err != nil {
		log.Printf("Failed to decode request body, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := graphql.NewRequest(requestJSON.CreateQueryString())

	response := models.TokensResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.Data.Tokens); err != nil {
		log.Printf("Failed to encode json response %v", err)
	}
}

func TokenVolume(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: /tokenvolume")

	w.Header().Add("Content-Type", "application/json")

	// create Uniswap Graph query from api request
	requestJSON := models.TokenVolumeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestJSON); err != nil {
		log.Printf("Failed to decode request body, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := graphql.NewRequest(requestJSON.CreateQueryString())
	response := models.TokenDayDatasResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}

	var totalVolume float64 = 0
	for _,v := range response.Data.TokenDayDatas {
		fmt.Println(v.VolumeUSD)
		volumeInt,_ := strconv.ParseFloat(v.VolumeUSD, 64)
		fmt.Println(volumeInt)
		totalVolume += volumeInt
	}

	fmt.Println(totalVolume)

	tokenVolumeResponse := models.TokenVolumeResponse {
		TotalVolumeUSD: totalVolume,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tokenVolumeResponse); err != nil {
		log.Printf("Failed to encode json response %v", err)
	}
}

func BlockSwaps(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: /blockswaps")

	w.Header().Add("Content-Type", "application/json")

	// create Uniswap Graph query from api request
	requestJSON := models.BlockInfoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestJSON); err != nil {
		log.Printf("Failed to decode request body, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	query := graphql.NewRequest(requestJSON.CreateQueryString())
	response := models.SwapsResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		log.Printf("Failed to encode json response %v", err)
	}
}
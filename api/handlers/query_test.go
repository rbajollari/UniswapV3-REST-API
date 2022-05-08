package handlers

import (
	"testing"
	"context"
	"strconv"

	"github.com/machinebox/graphql"
)

func TestUniswapV3GraphQuery (t *testing.T) {
	query := graphql.NewRequest(`
	{
		factories(first: 5) {
		  id
		  poolCount
		  txCount
		  totalVolumeUSD
		}
		bundles(first: 5) {
		  id
		  ethPriceUSD
		}
	  }
	`)
	var response interface {}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response); err != nil {
		panic(err)
	}
	t.Log(response)
}

func TestTokenPoolQuery (t *testing.T) {
	request := TokenPoolsRequest{
		ID: "0x6b175474e89094c44da98b954eedeac495271d0f",
	}
	query := graphql.NewRequest(request.createQueryString())
	response := TokensResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}
	t.Log(response.Data.Tokens)
}

func TestTokenVolumeQuery (t *testing.T) {
	request := TokenVolumeRequest{
		ID: "0x6b175474e89094c44da98b954eedeac495271d0f",
		StartingBlockTimestamp: 1651708800,
		EndingBlockTimestamp: 1651881600,
	}
	query := graphql.NewRequest(request.createQueryString())
	response := TokenDayDatasResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}

	var totalVolume float64 = 0
	for _,v := range response.Data.TokenDayDatas {
		volumeInt,_ := strconv.ParseFloat(v.VolumeUSD, 64)
		totalVolume += volumeInt
	}

	t.Log(totalVolume)
}

func TestBlockQuery (t *testing.T) {
	request := BlockInfoRequest{
		BlockNumber: 14732281,
	}
	query := graphql.NewRequest(request.createQueryString())
	response := SwapsResponse{}
	graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")
	if err := graphqlClient.Run(context.Background(), query, &response.Data); err != nil {
		panic(err)
	}

	t.Log(response.Data.Swaps)
}

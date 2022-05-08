package models

import (
	"fmt"
	"strconv"
)

type TokenPoolsRequest struct {
	ID	string `json:"id"`
}

type TokenVolumeRequest struct {
	ID	string `json:"id"`
	StartingBlockTimestamp	int	`json:"startingBlockTimestamp"`
	EndingBlockTimestamp	int	`json:"endingBlockTimestamp"`
}

type BlockInfoRequest struct {
	BlockNumber int `json:"blockNumber"`
}

func (t *TokenPoolsRequest) CreateQueryString() string {
	return fmt.Sprintf(`
	{
		tokens (where: {
			id: "%s"
		}) {
		whitelistPools {
			token0 {
				id
				name
				symbol
			}
			token1 {
				id
				name
				symbol
			}
		  }
		}
	}`, t.ID)
}

func (t *TokenVolumeRequest) CreateQueryString() string {
	gteDate := strconv.Itoa(t.StartingBlockTimestamp/86400)
	lteDate := strconv.Itoa(t.EndingBlockTimestamp/86400)
	id_gte := t.ID + "-" + gteDate
	id_lte := t.ID + "-" + lteDate

	return fmt.Sprintf(`
	{
		tokenDayDatas (where: {
			id_gte: "%s"
			id_lte: "%s"
		}){
			volumeUSD
		}
	  }`, id_gte, id_lte)
}

func (b *BlockInfoRequest) CreateQueryString() string {
	return fmt.Sprintf(`
	{
		swaps (block: {number: %d}) {
			transaction {
				id
			}
			token0 {
				id
				name
				symbol
			}
			token1 {
				id
				name
				symbol
			}
		}
	}`, b.BlockNumber)
}
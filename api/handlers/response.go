package handlers

type TokensResponse struct {
	Data struct {
		Tokens []struct {
			WhitelistPools []struct {
				ID     string `json:"id"`
				Token0 struct {
					ID     string `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
				} `json:"token0"`
				Token1 struct {
					ID     string `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
				} `json:"token1"`
			} `json:"whitelistPools"`
		} `json:"tokens"`
	} `json:"data"`
}

type TokenDayDatasResponse struct {
	Data struct {
		TokenDayDatas []struct {
			VolumeUSD string `json:"volumeUSD"`
		} `json:"tokenDayDatas"`
	} `json:"data"`
}

type TokenVolumeResponse struct {
	TotalVolumeUSD	float64	`json:"totalVolumeUSD"`
}

type SwapsResponse struct {
	Data struct {
		Swaps []struct {
			Transaction struct {
				ID string `json:"id"`
			} `json:"transaction"`
			Token0 struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
			} `json:"token0"`
			Token1 struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
			} `json:"token1"`
		} `json:"swaps"`
	} `json:"data"`
}
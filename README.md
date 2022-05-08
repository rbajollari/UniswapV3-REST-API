# UniswapV3-REST-API
A server side REST application that uses The Graph’s GraphQL API to provide Uniswap v3 information upon user request

## Details
Given an asset ID:
- What pools exist that include it?
- What is the total volume of that asset swapped in a given time range?
Given a block number:
- What swaps occurred during that specific block?
- List all assets swapped during that specific block

# Endpoints
## /tokenpools
### Accepts
```
{
    "id": "Address of token"
}
```
### Returns 
```
[
    {
    "whitelistPools": [
        {
        "id": "Address of Pool",
        "token0": {
            "id": "Address of token0",
            "name": "Name of token0",
            "symbol": "Symbol of token0"
        },
        "token1": {
            "id": "Address of token1",
            "name": "Name of token1",
            "symbol": "Symbol of token1"
        }
        },
        ...
    ]
    }
]
```
### Example request for token pools that include Dai
curl -X GET http://localhost:8080/tokenpools -H "Accept: application/json" -d '{"id":"0x6b175474e89094c44da98b954eedeac495271d0f"}'

## /tokenvolume
### Accepts
```
{
    "id": "Address of token"
    "startingBlockTimestamp": Unix timestamp of startdate
    "endingBlockTimestamp": Unix timestamp of enddate
}
```
### Returns 
```
{
    "totalVolumeUSD": Swap volume of given asset in given range
}
```
### Example request for total volume of Dai swapped in USD between Thursday, May 5, 2022 12:00:00 AM GMT and Saturday, May 7, 2022 12:00:00 AM GMT
curl -X GET http://localhost:8080/tokenvolume -H "Accept: application/json" -d '{"id":"0x6b175474e89094c44da98b954eedeac495271d0f", "startingBlockTimeStamp": 1651708800, "endingBlockTimestamp": 1651881600}'

## /blockswaps
### Accepts
```
{
    "blockNumber": Block number to check swaps in
}
```
### Returns 
```
{
    "swaps": [
      {
        "transaction": {
            "id": "Transaction ID"
        },
        "token0": {
            "id": "Address of token0",
            "name": "Name of token0",
            "symbol": "Symbol of token0"
        },
        "token1": {
            "id": "Address of token1",
            "name": "Name of token1",
            "symbol": "Symbol of token1"
        }
      }
    ]
  }
}
```
### Example request for all swap transactions and assests swapped at block 14732281
curl -X GET http://localhost:8080/blockswaps -H "Accept: application/json" -d '{"blockNumber": 14732281}'

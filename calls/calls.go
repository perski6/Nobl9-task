package calls

import (
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/net/context"
	"perski6.com/nobl9-task/constants"
)

const generateIntegers = "generateIntegers"

type RequestParams struct {
	ApiKey string `json:"apiKey"`
	N      int    `json:"n"`
	Min    int    `json:"min"`
	Max    int    `json:"max"`
}
type Data struct {
	Data []int `json:"data"`
}
type Response struct {
	Random         Data   `json:"random"`
	CompletionTime string `json:"completionTime"`
}

func GetIntegers(client jsonrpc.RPCClient, n int, ctx context.Context) ([]int, error) {

	request := RequestParams{
		ApiKey: constants.GetApiKey(),
		N:      n,
		Min:    constants.GetMin(),
		Max:    constants.GetMax(),
	}

	response, err := client.Call(ctx, generateIntegers, &request)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, response.Error
	}

	var responseData Response

	err = response.GetObject(&responseData)
	if err != nil {
	}
	var randomNumbers []int
	for _, v := range responseData.Random.Data {
		randomNumbers = append(randomNumbers, v)
	}
	return randomNumbers, nil
}

package handlers

import (
	"encoding/json"
	"github.com/montanaflynn/stats"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"perski6.com/nobl9-task/calls"
	"strconv"
	"sync"
	"time"
)

type NumbersWithDeviation struct {
	Sttdev float64 `json:"sttdev"`
	Data   []int   `json:"data"`
}

type Response []NumbersWithDeviation

func RandomsHandler(w http.ResponseWriter, r *http.Request, c jsonrpc.RPCClient) {
	length, _ := strconv.Atoi(r.FormValue("length"))
	requests, _ := strconv.Atoi(r.FormValue("requests"))
	var arr []NumbersWithDeviation
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	randomsChannel := make(chan NumbersWithDeviation, requests)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go getRandoms(&wg, randomsChannel, c, length, ctx)
	}
	wg.Wait()
	close(randomsChannel)

	for v := range randomsChannel {
		arr = append(arr, v)
	}

	arr = calculateAllDeviation(arr)

	jsonResp, err := json.Marshal(arr)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		w.WriteHeader(500)
	}
}

func convertTo64(ar []int) []float64 {
	newar := make([]float64, len(ar))
	var v int
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}

func calculateDeviation(arr []int) NumbersWithDeviation {
	floatArr := convertTo64(arr)
	sttdev, err := stats.StandardDeviation(floatArr)
	if err != nil {

	}
	return NumbersWithDeviation{
		Sttdev: sttdev,
		Data:   arr,
	}
}

func calculateAllDeviation(arr []NumbersWithDeviation) Response {
	var randomNumbers []int
	for _, v := range arr {
		randomNumbers = append(randomNumbers, v.Data...)
	}
	arr = append(arr, calculateDeviation(randomNumbers))
	return arr
}

func getRandoms(wg *sync.WaitGroup, ch chan<- NumbersWithDeviation, c jsonrpc.RPCClient, length int, ctx context.Context) {
	defer wg.Done()
	arr, err := calls.GetIntegers(c, length, ctx)
	if err != nil {
		log.Fatalf("error: %#v\n", err)
	}

	arrWithDeviation := calculateDeviation(arr)
	ch <- arrWithDeviation
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func bootstrapMedian(data []float64, numSamples int) float64 { //Resamples data and computes median for each resample
	medians := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		sample := make([]float64, len(data))
		for j := range sample {
			sample[j] = data[rand.Intn(len(data))]
		}
		medians[i] = median(sample)
	}
	return median(medians)
}

func median(data []float64) float64 { //Calculates median
	sort.Float64s(data) //Sorts slice in ascending order
	n := len(data)
	if n == 0 {
		return 0
	}
	middle := n / 2
	if n%2 == 1 {
		return data[middle]
	}
	return (data[middle-1] + data[middle]) / 2.0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, 100)
	for i := range data {
		data[i] = rand.NormFloat64()*3 + 0 //Generating normal distribution around 0 to emulate the R program as best as possible
	}
	fmt.Println("Original Data: ", data)
	fmt.Println("Bootstrap Median: ", bootstrapMedian(data, 1000))
}

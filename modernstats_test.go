package main

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{"odd number of elements", []float64{1, 3, 2}, 2},
		{"even number of elements", []float64{1, 2, 3, 4}, 2.5},
		{"empty slice", []float64{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.data); got != tt.expected {
				t.Errorf("median(%v) = %v, want %v", tt.data, got, tt.expected)
			}
		})
	}
}

func TestBootstrapMedian(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := bootstrapMedian(data, 1000)
	if math.IsNaN(got) || got == 0 {
		t.Errorf("bootstrapMedian returned an unexpected result: %v", got)
	}
}

func BenchmarkMedian(b *testing.B) {
	b.Run("Benchmark median on 10 elements", func(b *testing.B) {
		data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < b.N; i++ {
			median(data)
		}
	})
}

func BenchmarkBootstrapMedian(b *testing.B) {
	data := make([]float64, 100)
	for i := range data {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bootstrapMedian(data, 100)
	}
}

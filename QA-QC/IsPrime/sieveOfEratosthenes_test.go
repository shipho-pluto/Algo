package main

import (
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	testcases := []struct {
		name   string
		input  int
		output []int
	}{
		{"fo four: 2, 3, 4 => 2, 3", 4, []int{2, 3}},
		{"fo nine: 2, 3, 4, 5, 6, 7, 8, 9 => 2, 3, 5, 7", 9, []int{2, 3, 5, 7}},
		{"fo fourteen: 2, 3, 4 ... 13 => 2, 3, 5, 7, 11, 13", 14, []int{2, 3, 5, 7, 11, 13}},
	}

	for _, tss := range testcases {
		t.Run(tss.name, func(t *testing.T) {
			testingOutput := SieveOfEratosthenes(tss.input)
			if len(testingOutput) != len(tss.output) {
				t.Errorf("Failed in length. Get %d - want %d", len(testingOutput), len(tss.output))
			} else {
				for i := range testingOutput {
					if testingOutput[i] != tss.output[i] {
						t.Errorf("Failed in data. Get %d - want %d", testingOutput[i], tss.output[i])
					}
				}
			}

			t.Log("Testing end")
		})
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	testcases := []struct {
		name   string
		input  int
		output []int
	}{
		{"fo four: 2, 3, 4 => 2, 3", 4, []int{2, 3}},
		{"fo nine: 2, 3, 4, 5, 6, 7, 8, 9 => 2, 3, 5, 7", 9, []int{2, 3, 5, 7}},
		{"fo fourteen: 2, 3, 4 ... 13 => 2, 3, 5, 7, 11, 13", 14, []int{2, 3, 5, 7, 11, 13}},
	}

	for _, tss := range testcases {
		b.Run(tss.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testingOutput := SieveOfEratosthenes(tss.input)
				if len(testingOutput) != len(tss.output) {
					b.Errorf("Failed in length. Get %d - want %d", len(testingOutput), len(tss.output))
				} else {
					for j := range testingOutput {
						if testingOutput[j] != tss.output[j] {
							b.Errorf("Failed in data. Get %d - want %d", testingOutput[j], tss.output[j])
						}
					}
				}

				b.Log("Testing end")
			}
		})
	}
}

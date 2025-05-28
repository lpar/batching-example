package main_test

import (
	"fmt"
	"maps"
	"slices"
)

// This file provides examples of how to use the iter.Seq and iter.Seq2.

func ExampleValues() {
	// An iter.Seq is a function that takes a function and applies it to
	// each element in a slice. If the function you provide returns false,
	// the iteration stops.
	// You get an iter.Seq by calling utility functions like slices.Values
	// or slices.Backward.
	seq := slices.Values([]string{"india", "alpha", "quebec", "bravo"})
	fun := func(value string) bool {
		fmt.Printf("%s ", value)
		return true // continue iteration
	}
	seq(fun)
	// Output:
	// india alpha quebec bravo
}

func ExampleValues_range() {
	// You can also use an iter.Seq in a single-value range loop.
	seq := slices.Values([]string{"india", "alpha", "quebec", "bravo"})
	for value := range seq {
		fmt.Printf("%s ", value)
	}
	// Output:
	// india alpha quebec bravo
}

func ExampleAll() {
	// An iter.Seq2 is like an iter.Seq, but it takes a function that accepts
	// two arguments (key and value).
	// You get an iter.Seq2 by calling utility functions like slices.All.
	seq2 := slices.All([]string{"india", "alpha", "quebec", "bravo"})
	fun := func(idx int, value string) bool {
		fmt.Printf("%d=%s ", idx, value)
		return true // continue iteration
	}
	seq2(fun)
	// Output:
	// 0=india 1=alpha 2=quebec 3=bravo
}

func ExampleAll_range() {
	// You can also use an iter.Seq2 in a two-value range loop.
	seq2 := slices.All([]string{"india", "alpha", "quebec", "bravo"})
	for idx, value := range seq2 {
		fmt.Printf("%d=%s ", idx, value)
	}
	// Output:
	// 0=india 1=alpha 2=quebec 3=bravo
}

func ExampleCollect() {
	// The slices.Collect function turns an iter.Seq back into a slice.
	seq := slices.Values([]string{"india", "alpha", "quebec", "bravo"})
	slice := slices.Collect(seq)
	fmt.Println(slice)
	// Output:
	// [india alpha quebec bravo]
}

func ExampleCollect_seq2() {
	// The maps.Collect function turns an iter.Seq2 into a map.
	seq := slices.All([]string{"india", "alpha", "quebec", "bravo"})
	m := maps.Collect(seq)
	// We need to deal with the fact that maps have no guaranteed order.
	// The maps.Keys function returns a slice of the keys in the map.
	// The slices.Sorted function will take the values from a Seq, sort them,
	// and return them as a slice.
	// The values have to be comparable; see https://go.dev/ref/spec#Comparison_operators
	keys := slices.Sorted(maps.Keys(m))
	for key := range keys {
		fmt.Printf("%d=%s ", key, m[key])
	}
	// Output:
	// 0=india 1=alpha 2=quebec 3=bravo
}

func ExampleAll_map() {
	// The maps.All function returns an iter.Seq2 that iterates over a map.
	m := map[string]int{
		"alpha":   1,
		"bravo":   2,
		"charlie": 3,
		"delta":   4,
	}
	fun := func(key string, value int) bool {
		if key == "charlie" {
			fmt.Printf("%s=%d ", key, value)
			return false // stop iteration when we find the value
		}
		return true
	}
	// Get the iter.Seq2 from maps.All and pass it the function to apply.
	maps.All(m)(fun)
	// Output:
	// charlie=3
}

func ExampleChunk() {
	// slices.Chunk will split a slice into multiple chunks of a given size
	// and return an iter.Seq you can use to range over them.
	chunks := slices.Chunk([]string{"india", "alpha", "quebec", "bravo", "hotel"}, 3)
	for chunk := range chunks {
		fmt.Printf("%v\n", chunk)
	}
	// Output:
	// [india alpha quebec]
	// [bravo hotel]
}

func Example() {
	// Now an example of a custom iter.Seq that generates prime numbers.
	// First define a function to return whether an integer is a prime number.
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	// Now build a function that returns an iter.Seq, and that iter.Seq
	// applies a function to the prime numbers until the function asks
	// it to stop.
	Primes := func(f func(int) bool) {
		for i := 2; ; i++ {
			if isPrime(i) {
				if !f(i) {
					break
				}
			}
		}
	}
	// Define our function to call for each prime number.
	fun := func(n int) bool {
		fmt.Printf("%d ", n)
		// Stop after we get a prime larger than 20.
		if n > 20 {
			return false
		}
		return true
	}
	// Call the Primes function to apply our function to the prime numbers.
	Primes(fun)
	// Output:
	// 2 3 5 7 11 13 17 19 23
}

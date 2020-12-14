package lib

import (
	"fmt"
	"time"
)

func runGraphMemoryBenchmark(graph Graph, count uint32) {
	if count < 100 {
		fmt.Errorf("benchmarks require to pass a %v >= 100", count)
		return
	}
	var uint64Count = uint64(count)
	var step = uint64Count / 100
	fmt.Println("Populating graph...")
	var start = time.Now()
	for i := uint64(0); i < uint64Count; i++ {
		for j := uint64(0); j < ContactCountAverage; j++ {
			if i > j { graph.AddContact(i, i + j) }
			graph.AddContact(i, i - j)
		}
		if i % step == 0 {
			fmt.Printf("%v - Populating graph, step: %v/%v\n", time.Since(start), i, count)
		}
	}

	fmt.Printf("%v - Done populating graph.\n", time.Since(start))
	printMemoryUsage()
}

func RunAGraphMemoryBenchmark(count uint32) {
	var graph = NewAGraphOfCap(count)
	runGraphMemoryBenchmark(graph, count)
}

func RunBGraphMemoryBenchmark(count uint32) {
	var graph = NewBGraphOfCap(count)
	runGraphMemoryBenchmark(graph, count)
}
package main

import "./lib"

func main() {
	lib.RunAGraphMemoryBenchmark(1_000_000)
	lib.RunBGraphMemoryBenchmark(1_000_000)
}

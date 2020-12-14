package lib

import (
	"testing"
)

func createBGraph(s sample) BGraph {
	var graph = BGraph{}
	graph.AddContact(s.x, s.y)
	graph.AddContact(s.x, s.z)
	graph.AddContact(s.y, s.x)
	graph.AddContact(s.y, s.z)
	graph.AddContact(s.p, s.x)
	return graph
}

func createBGraphLookupTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createBGraph(s),
		testCases: createLookupTestCases(s),
	}
}
// We check that we actually return contact phonebook numbers
func TestBGraph_Lookup(t *testing.T) {
	var lookupTestCases = createBGraphLookupTestCases()
	var graph = lookupTestCases.graph

	var errorMessagePattern = "Lookup(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range lookupTestCases.testCases {
		var actual = graph.Lookup(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}

func createBGraphRLookupTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createBGraph(s),
		testCases: createRLookupTestCases(s),
	}
}
// We check that we actually return contacts who have the phone number in their resp. books.
func TestBGraph_RLookup(t *testing.T) {
	var rLookupTestCases = createBGraphRLookupTestCases()
	var graph = rLookupTestCases.graph

	var errorMessagePattern = "RLookup(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range rLookupTestCases.testCases {
		var actual = graph.RLookup(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}

func createBGraphSuggestTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createBGraph(s),
		testCases: createSuggestTestCases(s),
	}
}
// We check that we actually return at most 10 contacts who are 2nd degree contacts.
func TestBGraph_Suggest(t *testing.T) {
	var suggestTestCases = createBGraphSuggestTestCases()
	var graph = suggestTestCases.graph

	var errorMessagePattern = "Suggest(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range suggestTestCases.testCases {
		var actual = graph.Suggest(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}


func BenchmarkBGraph_AddContact_50(b *testing.B) {
	var graph = NewBGraph()
	benchmarkAddContact(graph, 50)
}
func BenchmarkBGraph_AddContact_100(b *testing.B) {
	var graph = NewBGraph()
	benchmarkAddContact(graph, 100)
}
func BenchmarkBGraph_AddContact_500(b *testing.B) {
	var graph = NewBGraph()
	benchmarkAddContact(graph, 500)
}
func BenchmarkBGraph_AddContact_1000(b *testing.B) {
	var graph = NewBGraph()
	benchmarkAddContact(graph, 1000)
}

func BenchmarkBGraph_Lookup_50(b *testing.B) {
	var graph = NewBGraph()
	benchmarkLookupContact(graph, 50, b)
}
func BenchmarkBGraph_Lookup_100(b *testing.B) {
	var graph = NewBGraph()
	benchmarkLookupContact(graph, 100, b)
}
func BenchmarkBGraph_Lookup_500(b *testing.B) {
	var graph = NewBGraph()
	benchmarkLookupContact(graph, 500, b)
}
func BenchmarkBGraph_Lookup_1000(b *testing.B) {
	var graph = NewBGraph()
	benchmarkLookupContact(graph, 1000, b)
}

func BenchmarkBGraph_RLookup_50(b *testing.B) {
	var graph = NewBGraph()
	benchmarkRLookupContact(graph, 50, b)
}
func BenchmarkBGraph_RLookup_100(b *testing.B) {
	var graph = NewBGraph()
	benchmarkRLookupContact(graph, 100, b)
}
func BenchmarkBGraph_RLookup_500(b *testing.B) {
	var graph = NewBGraph()
	benchmarkRLookupContact(graph, 500, b)
}
func BenchmarkBGraph_RLookup_1000(b *testing.B) {
	var graph = NewBGraph()
	benchmarkRLookupContact(graph, 1000, b)
}

func BenchmarkBGraph_Suggest_50(b *testing.B) {
	var graph = NewBGraph()
	benchmarkSuggestContact(graph, 50, b)
}
func BenchmarkBGraph_Suggest_100(b *testing.B) {
	var graph = NewBGraph()
	benchmarkSuggestContact(graph, 100, b)
}
func BenchmarkBGraph_Suggest_500(b *testing.B) {
	var graph = NewBGraph()
	benchmarkSuggestContact(graph, 500, b)
}
func BenchmarkBGraph_Suggest_1000(b *testing.B) {
	var graph = NewBGraph()
	benchmarkSuggestContact(graph, 1000, b)
}

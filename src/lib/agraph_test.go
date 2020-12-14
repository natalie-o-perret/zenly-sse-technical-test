package lib

import (
	"testing"
)

func createAGraph(s sample) AGraph {
	var graph = AGraph{}
	graph.AddContact(s.x, s.y)
	graph.AddContact(s.x, s.z)
	graph.AddContact(s.y, s.x)
	graph.AddContact(s.y, s.z)
	graph.AddContact(s.p, s.x)
	return graph
}

func createAGraphLookupTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createAGraph(s),
		testCases: createLookupTestCases(s),
	}
}
// We check that we actually return contact phonebook numbers
func TestAGraph_Lookup(t *testing.T) {
	var lookupTestCases = createAGraphLookupTestCases()
	var graph = lookupTestCases.graph

	var errorMessagePattern = "Lookup(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range lookupTestCases.testCases {
		var actual = graph.Lookup(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}

func createAGraphRLookupTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createAGraph(s),
		testCases: createRLookupTestCases(s),
	}
}
// We check that we actually return contacts who have the phone number in their resp. books.
func TestAGraph_RLookup(t *testing.T) {
	var rLookupTestCases = createAGraphRLookupTestCases()
	var graph = rLookupTestCases.graph

	var errorMessagePattern = "RLookup(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range rLookupTestCases.testCases {
		var actual = graph.RLookup(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}

func createAGraphSuggestTestCases() graphSampleTestCases {
	var s = createSample()
	return graphSampleTestCases{
		sample:    s,
		graph:     createAGraph(s),
		testCases: createSuggestTestCases(s),
	}
}

// We check that we actually return at most 10 contacts who are 2nd degree contacts.
func TestAGraph_Suggest(t *testing.T) {
	var suggestTestCases = createAGraphSuggestTestCases()
	var graph = suggestTestCases.graph

	var errorMessagePattern = "Suggest(%v) outcome is incorrect\nActual: %v\nExpected: %v."
	for _, testCase := range suggestTestCases.testCases {
		var actual = graph.Suggest(testCase.input)
		checkOutcome(t, errorMessagePattern, actual, testCase)
	}
}


func BenchmarkAGraph_AddContact_50(b *testing.B) {
	var graph = NewAGraph()
	benchmarkAddContact(graph, 50)
}
func BenchmarkAGraph_AddContact_100(b *testing.B) {
	var graph = NewAGraph()
	benchmarkAddContact(graph, 100)
}
func BenchmarkAGraph_AddContact_500(b *testing.B) {
	var graph = NewAGraph()
	benchmarkAddContact(graph, 500)
}
func BenchmarkAGraph_AddContact_1000(b *testing.B) {
	var graph = NewAGraph()
	benchmarkAddContact(graph, 1000)
}

func BenchmarkAGraph_Lookup_50(b *testing.B) {
	var graph = NewAGraph()
	benchmarkLookupContact(graph, 50, b)
}
func BenchmarkAGraph_Lookup_100(b *testing.B) {
	var graph = NewAGraph()
	benchmarkLookupContact(graph, 100, b)
}
func BenchmarkAGraph_Lookup_500(b *testing.B) {
	var graph = NewAGraph()
	benchmarkLookupContact(graph, 500, b)
}
func BenchmarkAGraph_Lookup_1000(b *testing.B) {
	var graph = NewAGraph()
	benchmarkLookupContact(graph, 1000, b)
}

func BenchmarkAGraph_RLookup_50(b *testing.B) {
	var graph = NewAGraph()
	benchmarkRLookupContact(graph, 50, b)
}
func BenchmarkAGraph_RLookup_100(b *testing.B) {
	var graph = NewAGraph()
	benchmarkRLookupContact(graph, 100, b)
}
func BenchmarkAGraph_RLookup_500(b *testing.B) {
	var graph = NewAGraph()
	benchmarkRLookupContact(graph, 500, b)
}
func BenchmarkAGraph_RLookup_1000(b *testing.B) {
	var graph = NewAGraph()
	benchmarkRLookupContact(graph, 1000, b)
}

func BenchmarkAGraph_Suggest_50(b *testing.B) {
	var graph = NewAGraph()
	benchmarkSuggestContact(graph, 50, b)
}
func BenchmarkAGraph_Suggest_100(b *testing.B) {
	var graph = NewAGraph()
	benchmarkSuggestContact(graph, 100, b)
}
func BenchmarkAGraph_Suggest_500(b *testing.B) {
	var graph = NewAGraph()
	benchmarkSuggestContact(graph, 500, b)
}
func BenchmarkAGraph_Suggest_1000(b *testing.B) {
	var graph = NewAGraph()
	benchmarkSuggestContact(graph, 1000, b)
}
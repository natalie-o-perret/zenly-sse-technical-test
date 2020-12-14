package lib

import (
	"reflect"
	"sort"
	"testing"
)

// For the grand of simplicity, letters are easier to reason with
func phoneNumberOf(contact rune) uint64 {
	return uint64(contact)
}
func stringContactOf(phoneNumber uint64) string {
	return string(rune(phoneNumber))
}
func stringContactArrayOf(phoneNumbers []uint64) []string {
	var contacts = make([]string, len(phoneNumbers))
	for i, phoneNumber := range phoneNumbers {
		contacts[i] = stringContactOf(phoneNumber)
	}
	return contacts
}

type sample struct {
	x uint64
	y uint64
	z uint64
	p uint64
}

func createSample() sample {
	return sample{
		x: phoneNumberOf('x'),
		y: phoneNumberOf('y'),
		z: phoneNumberOf('z'),
		p: phoneNumberOf('p'),
	}
}

type testCase struct {
	input uint64
	expected []uint64
}
type graphSampleTestCases struct {
	graph     Graph
	sample    sample
	testCases []testCase
}

func createLookupTestCases(s sample) []testCase {
	return []testCase{
		{input: s.x, expected: []uint64{ s.y, s.z }},
		{input: s.y, expected: []uint64{ s.x, s.z }},
		{input: s.z, expected: []uint64{}},
		{input: s.p, expected: []uint64{ s.x }},
	}
}
func createRLookupTestCases(s sample) []testCase {
	return[]testCase{
		{input: s.x, expected: []uint64{ s.y, s.p }},
		{input: s.y, expected: []uint64{ s.x }},
		{input: s.z, expected: []uint64{ s.x, s.y }},
		{input: s.p, expected: []uint64{}},
	}
}
func createSuggestTestCases(s sample) []testCase {
	return[]testCase{
		{input: s.x, expected: []uint64{}},
		{input: s.y, expected: []uint64{}},
		{input: s.z, expected: []uint64{}},
		{input: s.p, expected: []uint64{ s.y, s.z }},
	}
}

func sortUint64Slice(items []uint64) []uint64 {
	var sortedItems = make([]uint64, len(items))
	copy(sortedItems, items)
	sort.Slice(sortedItems, func(i, j int) bool { return sortedItems[i] < sortedItems[j] })
	return sortedItems
}

func areEquivalent(a []uint64, b []uint64) bool {
	var sortedA = sortUint64Slice(a)
	var sortedB = sortUint64Slice(b)
	return reflect.DeepEqual(sortedA, sortedB)
}

func checkOutcome (t *testing.T, errorMessagePattern string, actual []uint64, testCase testCase) {
	if areEquivalent(actual, testCase.expected) {
		return
	}
	t.Errorf(errorMessagePattern,
		stringContactOf(testCase.input),
		stringContactArrayOf(actual),
		stringContactArrayOf(testCase.expected))
	return
}

func benchmarkAddContact(graph Graph, count uint16) {
	var uint64Count = uint64(count)
	for i := uint64(1); i < uint64Count; i++ {
		graph.AddContact(0, i)
	}
}

func populate(graph Graph, count uint16) {
	var uint64Count = uint64(count)
	for i := uint64(0); i < uint64Count; i++ {
		for j := uint64(0); j < ContactCountAverage; j++ {
			if i > j { graph.AddContact(i, i + j) }
			graph.AddContact(i, i - j)
		}
	}
}

func benchmarkLookupContact(graph Graph, count uint16, b *testing.B) {
	populate(graph, count)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		graph.Lookup(uint64(count / 2))
	}
}
func benchmarkRLookupContact(graph Graph, count uint16, b *testing.B) {
	populate(graph, count)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		graph.RLookup(uint64(count / 2))
	}
}
func benchmarkSuggestContact(graph Graph, count uint16, b *testing.B) {
	populate(graph, count)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		graph.Suggest(uint64(count / 2))
	}
}

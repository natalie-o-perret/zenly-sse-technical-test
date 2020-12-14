package lib

// ANode represents a contact node on a social graph holding resp.:
// - to: the phone numbers of this contact book
// - from: the phone numbers of those who have this contact phone number in their books
type ANode struct {
	to    []uint64
	from  []uint64
}

// Represents a social graph of contact nodes backed by phone number arrays
type AGraph map[uint64]*ANode

// Create a new AGraph of the given capacity
func NewAGraphOfCap(capacity uint32) AGraph {
	return make(AGraph, capacity)
}

// Create a new AGraph of default capacity
func NewAGraph() AGraph {
	return make(AGraph)
}

// Either get an existing contact node given a phone number
// Or add a new one and return it as well.
// Pre-allocated with an initial capacity set to the given AverageContactCount
// => avoid further too many resizing operations
func (aGraph AGraph) getOrAddNode(phoneNumber uint64) *ANode {
	if node, found := aGraph[phoneNumber]; found {
		return node
	}
	node := &ANode{
		from:  make([]uint64, 0, DefaultContactCount),
		to:    make([]uint64, 0, DefaultContactCount),
	}
	aGraph[phoneNumber] = node
	return node
}

// Add the given new contact to the current contact phone directory.
func (aGraph AGraph) AddContact(currentContactPhoneNumber uint64, newContactPhoneNumber uint64) {
	var previousNode = aGraph.getOrAddNode(currentContactPhoneNumber)
	var nextNode = aGraph.getOrAddNode(newContactPhoneNumber)

	previousNode.to = append(previousNode.to, newContactPhoneNumber)
	nextNode.from = append(nextNode.from, currentContactPhoneNumber)
}

// Fetch the given contact 1st degree phone directory entries.
func (aGraph AGraph) Lookup(phoneNumber uint64) []uint64 {
	if node, found := aGraph[phoneNumber]; found {
		return node.to
	}
	return []uint64{}
}

// Fetch the 1st degree contacts who have the given contact phone number in their phone directories.
func (aGraph AGraph) RLookup(phoneNumber uint64) []uint64 {
	if node, found := aGraph[phoneNumber]; found {
		return node.from
	}
	return []uint64{}
}

// Suggest at most 10 phone numbers who are 2nd degree contacts
func (aGraph AGraph) Suggest(phoneNumber uint64) []uint64 {
	var fstDegPnsSet = setOfSlice(aGraph[phoneNumber].to)
	var sndDegPnsSet = set{}
	for _, firstDegPn := range aGraph[phoneNumber].to {
		for _, sndDegPn := range aGraph[firstDegPn].to {
			if len(sndDegPnsSet) >= 10 {
				break
			}
			if !fstDegPnsSet.contains(sndDegPn) && sndDegPn != phoneNumber {
				sndDegPnsSet.addUInt64(sndDegPn)
			}
		}
	}
	var sndDegPns = sndDegPnsSet.toSlice()
	// Could be sorted to systematically guarantee the same order (sets aka maps here aren't)
	return sndDegPns
}
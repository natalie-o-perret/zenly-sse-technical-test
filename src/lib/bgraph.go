package lib

import (
	"github.com/RoaringBitmap/roaring/roaring64"
)

// BNode represents a contact node on a social graph holding resp.:
// - to: the phone numbers of this contact book
// - from: the phone numbers of those who have this contact phone number in their books
type BNode struct {
	to    *roaring64.Bitmap
	from  *roaring64.Bitmap
}

// Represents a social graph of contact nodes backed by phone number bitmaps
type BGraph map[uint64]*BNode

// Create a new BGraph of the given capacity
func NewBGraphOfCap(capacity uint32) BGraph {
	return make(BGraph, capacity)
}

// Create a new BGraph of default capacity
func NewBGraph() BGraph {
	return make(BGraph)
}

func (bGraph BGraph) getOrAddNode(phoneNumber uint64) *BNode {
	if node, found := bGraph[phoneNumber]; found {
		return node
	}
	node := &BNode{
		from:  roaring64.New(),
		to:    roaring64.New(),
	}
	bGraph[phoneNumber] = node
	return node
}

// Add the given new contact to the current contact phone directory.
func (bGraph BGraph) AddContact(currentContactPhoneNumber uint64, newContactPhoneNumber uint64) {
	var previousNode = bGraph.getOrAddNode(currentContactPhoneNumber)
	var nextNode = bGraph.getOrAddNode(newContactPhoneNumber)
	previousNode.to.Add(newContactPhoneNumber)
	nextNode.from.Add(currentContactPhoneNumber)
}

// Fetch the given contact 1st degree phone directory entries.
func (bGraph BGraph) Lookup(phoneNumber uint64) []uint64 {
	if node, found := bGraph[phoneNumber]; found {
		// Could be sorted to systematically guarantee the same order
		// Roaring bitmap does not always guarantee the same order (depends on the underlying container)
		// See D. Lemire presentation for more info
		return node.to.ToArray()
	}
	return []uint64{}
}

// Fetch the given contact 1st degree phone directory entries.
func (bGraph BGraph) RLookup(phoneNumber uint64) []uint64 {
	if node, found := bGraph[phoneNumber]; found {
		// Could be sorted to systematically guarantee the same order
		// Roaring bitmap does not always guarantee the same order (depends on the underlying container)
		// See D. Lemire presentation for more info
		return node.from.ToArray()
	}
	return []uint64{}
}

// Suggest at most 10 phone numbers (sorted to guarantee the same order) who are 2nd degree contacts
func (bGraph BGraph) Suggest(phoneNumber uint64) []uint64 {
	var fstDegPns = bGraph[phoneNumber].to.ToArray()
	var fstDegPnsSet = setOfSlice(fstDegPns)
	var sndDegPnsSet = set{}
	for _, firstDegPn := range fstDegPns {
		for _, sndDegPn := range bGraph[firstDegPn].to.ToArray() {
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
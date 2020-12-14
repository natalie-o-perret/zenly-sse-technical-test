package lib

type Graph interface {
	AddContact(currentContactPhoneNumber uint64, newContactPhoneNumber uint64)
	Lookup(phoneNumber uint64)  []uint64
	RLookup(phoneNumber uint64) []uint64
	Suggest(phoneNumber uint64) []uint64
}

const ContactCountAverage = 50
const ContactCountStd = 10
const DefaultContactCount = ContactCountAverage + ContactCountStd

func setOfSlice(items []uint64) set {
	var s = set{}
	for _, item := range items {
		s.addUInt64(item)
	}
	return s
}
func (s set) addUInt64(item uint64) {
	s[item] = struct{}{}
}
func (s set) contains (item uint64) bool {
	_, found := s[item]
	return found
}
func (s set) toSlice() []uint64 {
	var sl = make([]uint64, len(s))
	var i = 0
	for key := range s {
		sl[i] = key
		i++
	}
	return sl
}
type set map[uint64]struct{}
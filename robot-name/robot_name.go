package robotname

import (
	"math/rand"
)

type Robot struct {
	name string
}

var names = map[string]bool{}

// The first time you boot them up, a random name is generated in the format
// of two uppercase letters followed by three digits, such as RX837 or BC811.

// The names must be random: they should not follow a predictable sequence.
// Random names means a risk of collisions. Your solution should not allow
// the use of the same name twice when avoidable.

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = genRandomStr()
	}
	return r.name
}

func genRandomStr() string {
	s := ""
	b := [5]byte{}
	for {
		// got from
		// http://exercism.io/submissions/f8af11033ba245959c0d877cd54a8457
		rand.Read(b[:]) // Read generates len(p) random bytes
		// the string starts from b[0], little endian
		b[0] = b[0]%26 + 'A'
		b[1] = b[1]%26 + 'A'
		b[2] = b[2]%10 + '0'
		b[3] = b[3]%10 + '0'
		b[4] = b[4]%10 + '0'
		s = string(b[:])
		_, ok := names[s]
		if !ok { // if not found in the map
			names[s] = true
			break
		}
	}
	return s
}

// Every once in a while we need to reset a robot to its factory settings,
// which means that their name gets wiped. The next time you ask, it will
// respond with a new random name.
func (r *Robot) Reset() {
	// name should not be reissued
	// delete(names, r.name)
	r.name = "" // When Name() is called, it will trigger a generation
}

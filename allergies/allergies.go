package allergies

const testVersion = 1

// The list of items (and their value) that were tested are:
// * eggs (1)
// * peanuts (2)
// * shellfish (4)
// * strawberries (8)
// * tomatoes (16)
// * chocolate (32)
// * pollen (64)
// * cats (128)

var table = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// given a score, find all allergens,
// the allergens are scored as multiplers of 2.
// so the question is same as given an integer find the bits which are 1

func Allergies(score uint) (ret []string) {
	for {
		if score != 0 {
			n := score ^ (score & (score - 1))
			v, ok := table[int(n)]
			if ok {
				ret = append(ret, v)
			}
			score = score & (score - 1)

		} else {
			break
		}
	}
	return ret
}

func AllergicTo(i uint, allergen string) bool {
	for {
		if i != 0 {
			n := i ^ (i & (i - 1))
			v, ok := table[int(n)]
			if ok && v == allergen {
				return true
			}
			i = i & (i - 1)

		} else {
			break
		}
	}
	return false
}

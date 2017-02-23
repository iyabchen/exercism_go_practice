// Implement binary search without using standard library.

// If there are duplicate values of the key you are searching for,
// return the first occurance in the slice.

// If the search key is not present, SearchInts returns the index of
// the first value greater than the search key.
// If the key is greater than all values in the slice, SearchInts
// returns the length of the slice.

package binarysearch

import (
	"fmt"
)

const testVersion = 1

func SearchInts(slice []int, key int) int {
	if len(slice) == 0 {
		return 0
	}

	start := 0
	end := len(slice) - 1
	for start <= end {
		mid := start + (end-start)/2
		if slice[mid] == key {
			for mid = mid - 1; mid >= 0; mid-- {
				if slice[mid] != key {
					return mid + 1
				}
			}
			return 0
		} else if slice[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if end < 0 {
		return 0
	}
	if slice[end] > key {
		return end
	} else { // <key
		return end + 1
	}

}

//   k found at beginning of slice.
//   k found at end of slice.
//   k found at index fx.
//   k < all values.
//   k > all n values.
//   k > lv at lx, < gv at gx.
//   slice has no values.
func Message(slice []int, key int) string {
	if len(slice) == 0 {
		return "slice has no values"
	}

	ix := SearchInts(slice, key) // ix: [0,len]
	if ix == len(slice) {
		return fmt.Sprintf("%d > all %d values", key, len(slice))
	}
	val := slice[ix]
	if val == key {
		if ix == 0 {
			return fmt.Sprintf("%d found at beginning of slice", key)
		} else if ix == len(slice)-1 {
			return fmt.Sprintf("%d found at end of slice", key)
		} else {
			return fmt.Sprintf("%d found at index %d", key, ix)
		}
	} else if val < key {
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
			key, val, ix, slice[ix+1], ix+1)

	} else { // val > key
		if ix == 0 {
			return fmt.Sprintf("%d < all values", key)
		} else {
			return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
				key, slice[ix-1], ix-1, val, ix)
		}
	}
}

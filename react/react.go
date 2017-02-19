package react

import "reflect"

const testVersion = 4

type MyReactor struct {
	cells []*MyCell // only store compute cells
}

type MyCell struct {
	val       int
	ref1      *MyCell
	ref2      *MyCell
	function  func(...int) int
	callbacks []func(int)
	reactor   *MyReactor
}

func New() *MyReactor {
	return new(MyReactor)
}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (r *MyReactor) CreateInput(val int) InputCell {
	mycell := &MyCell{val: val, reactor: r}
	// r.cells = append(r.cells, mycell)
	return mycell

}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r *MyReactor) CreateCompute1(in Cell, f func(int) int) ComputeCell {
	foo := func(nums ...int) int {
		return f(nums[0])
	}
	c := in.(*MyCell)
	mycell := &MyCell{val: foo(in.Value()), ref1: c, function: foo, reactor: r}
	r.cells = append(r.cells, mycell)
	return mycell
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r *MyReactor) CreateCompute2(in1 Cell, in2 Cell,
	f func(int, int) int) ComputeCell {
	foo := func(nums ...int) int {
		return f(nums[0], nums[1])
	}
	c1 := in1.(*MyCell)
	c2 := in2.(*MyCell)

	mycell := &MyCell{val: foo(in1.Value(), in2.Value()),
		ref1: c1, ref2: c2, function: foo, reactor: r}
	r.cells = append(r.cells, mycell)
	return mycell
}

func (c *MyCell) SetValue(val int) {
	if c.ref1 != nil {
		panic("SetValue invoked in a compute cell")
	}
	if c.val == val {
		return
	} else {
		c.val = val
		// propagate()
		for _, cell := range c.reactor.cells {
			val := cell.val
			if val != cell.Value() {
				// call all added callbacks
				for _, cb := range cell.callbacks {
					cb(cell.val) // the value has changed after calling Value()
				}
			}
		}
	}

}

// Value returns the current value of the cell.
func (c *MyCell) Value() int {
	if c.ref1 == nil { // input cell
		return c.val
	}
	if c.ref2 != nil {
		c.val = c.function(c.ref1.Value(), c.ref2.Value())
	} else {
		c.val = c.function(c.ref1.Value())
	}
	return c.val

}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *MyCell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, f)
	return f

}

// RemoveCallback removes a previously added callback, if it exists.
func (c *MyCell) RemoveCallback(cb CallbackHandle) {
	for i, v := range c.callbacks {
		if reflect.ValueOf(v).Pointer() == reflect.ValueOf(cb).Pointer() {
			c.callbacks = append(c.callbacks[:i], c.callbacks[i+1:]...)
			return
		}
	}

}

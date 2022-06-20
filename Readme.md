# Num-Stair
Number is consider as a weight of block which can be a positive or negative. 
Stair is built up from each number block. Block is fit into a single step of the stair
and handrail height is depended on the weight of the step. 

![Overview](/doc/overview.png)

Library is capable of,
* Build steps with integer name
* Add block into any step (condition apply)
* Positioning a block

## Example Usage
### Build step
Name of the step should be greater than zero. Step name should be unique.
```go
package main

import (
	ns "github.com/HADLakmal/num-stair"
)

func main() {
	st := ns.NewStair()
	if !st.AddStep(1) {
		panic(`error add step`)
	}
}
```

### Add Block
Name of the step should be greater than zero. Step name should be unique.

```go
package main

import (
	ns "github.com/HADLakmal/num-stair"
	"github.com/senpathi/gofloat"
)

func main() {
	st := ns.NewStair()
	var step uint64  = 1
	if !st.AddStep(step) {
		panic(`error add step`)
	}
	if !st.AddBlock(step, ns.NewBlock(`x`, gofloat.ToFloat(10,2))){
		panic(`error add block`)
    }
}
```


### Position Block
There are two ways to position a block,
* Position block in specific step(above example)
* Suitable step from the beginning of stair
* Position block within given offset

#### Block Position in Stair

```go
package main

import (
	"fmt"
	ns "github.com/HADLakmal/num-stair"
	"github.com/senpathi/gofloat"
)

func main() {
	st := ns.NewStair()
	var step uint64 = 1
	// add three steps
	st.AddStep(step)
	st.AddStep(step + 1)
	st.AddStep(step + 2)

	st.AddBlock(step, ns.NewBlock(`block-1`, gofloat.ToFloat(5, 2)))
	st.AddBlock(step+1, ns.NewBlock(`block-2`, gofloat.ToFloat(20, 2)))
	st.AddBlock(step+2, ns.NewBlock(`block-3`, gofloat.ToFloat(5, 2)))

	// block position in block-2
	if st.PositionBlock(ns.NewBlock(`position`, gofloat.ToFloat(-10, 2))) {
		// print block-2
		fmt.Println(st.Steps[step+1].Inputs[0].ID)
		// print position
		fmt.Println(st.Steps[step+1].Inputs[1].ID)
	}
}
```

#### Block Position within Offset
Block will try to fit into given offset. If block can't fit to desire step then it will stop in best fit position. 
```go
package main

import (
	"fmt"
	ns "github.com/HADLakmal/num-stair"
	"github.com/senpathi/gofloat"
)

func main() {
	st := ns.NewStair()
	var step uint64 = 1
	// add three steps
	st.AddStep(step)
	st.AddStep(step + 2)
	st.AddStep(step + 4)

	st.AddBlock(step, ns.NewBlock(`block-1`, gofloat.ToFloat(20, 2)))
	st.AddBlock(step+2, ns.NewBlock(`block-2`, gofloat.ToFloat(20, 2)))
	st.AddBlock(step+4, ns.NewBlock(`block-3`, gofloat.ToFloat(20, 2)))

	// block position in block-2
	if st.PositionBlock(ns.NewBlock(`position`, gofloat.ToFloat(-10, 2)),ns.Offset(3)) {
		// print block-2
		fmt.Println(st.Steps[step+2].Inputs[0].ID)
		// print position
		fmt.Println(st.Steps[step+2].Inputs[1].ID)
	}
}
```
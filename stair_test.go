package number_stairs

import (
	"fmt"
	"github.com/senpathi/gofloat"
	"testing"
)

func TestNewStair(t *testing.T) {
	st := NewStair()
	st.AddStep(1)
	st.AddStep(3)
	st.AddStep(4)
	fmt.Printf("%t \n", st.AddBlock(1, Block{Value: gofloat.ToFloat(10, 2)}))
	fmt.Printf("%t \n", st.AddBlock(3, Block{Value: gofloat.ToFloat(100, 2)}))
	fmt.Printf("%t \n", st.AddBlock(4, Block{Value: gofloat.ToFloat(10, 2)}))
	st.PositionBlock(Block{Value: gofloat.ToFloat(-80, 2)})
	fmt.Printf("%t \n", st.AddBlock(3, Block{Value: gofloat.ToFloat(-20, 2)}))
}

func TestNewStairWithOffset(t *testing.T) {
	st := NewStair()
	st.AddStep(1)
	st.AddStep(3)
	st.AddStep(4)
	fmt.Printf("%t \n", st.AddBlock(1, Block{Value: gofloat.ToFloat(10, 2)}))
	fmt.Printf("%t \n", st.AddBlock(3, Block{Value: gofloat.ToFloat(100, 2)}))
	fmt.Printf("%t \n", st.AddBlock(4, Block{Value: gofloat.ToFloat(10, 2)}))
	fmt.Printf("%t \n", st.PositionBlock(Block{Value: gofloat.ToFloat(-5, 2)}, Offset(2)))
	fmt.Printf("%t \n", st.AddBlock(3, Block{Value: gofloat.ToFloat(-20, 2)}))
}

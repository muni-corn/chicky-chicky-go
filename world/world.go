package world

import (
	"github.com/municorn/chicky-chicky-go/blocks"
)

// World contains a slice of Plots
type World struct {
	terrain [][][]blocks.Plot
}

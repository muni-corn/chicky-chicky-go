package game

import (
	"fmt"

	"github.com/municorn/chicky-chicky-go/characters"
	"github.com/municorn/chicky-chicky-go/blocks"
	"github.com/municorn/chicky-chicky-go/input"
	"github.com/municorn/chicky-chicky-go/maths"
	"github.com/municorn/chicky-chicky-go/render"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var characterInControl characters.Character
var block = blocks.NewGrassBlock()

func init() {
	input.AddKeyboardListener(&keyListener{})
	chicken := characters.NewChicken()
	characterInControl = chicken
}

// Logic performs logic for the game. This includes movement, physics,
// clocks, animation, etc
func Logic(delta float32) {

}

var cam = render.NewCamera(maths.Vec3{X:0, Y:0, Z:1}, 70, 800.0/600)

// Render renders the game.
func Render() {
	characterInControl.Render(cam)
	// block.Render(cam)
}

type keyListener struct{}

func (k *keyListener) KeyDown(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key down:", int(key))
}

func (k *keyListener) KeyUp(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key up:", int(key))
}

func (k *keyListener) KeyRepeat(key glfw.Key, scancode int, mods glfw.ModifierKey) {
	fmt.Printf("%-20s%d\n", "key repeat:", int(key))
}


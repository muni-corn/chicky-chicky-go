package characters

import (
    "github.com/municorn/chicky-chicky-go/types"
    "github.com/municorn/chicky-chicky-go/textures"
    "github.com/municorn/chicky-chicky-go/world"
    "github.com/municorn/chicky-chicky-go/maths"

    "os"
)

var chickenTextures map[CharacterAction]*textures.Sprite

func init() {
    chickenTextures[ActionNothing] = textures.MustNewSprite("assets/photos/chicken/stand.png", 1, 0)
    chickenTextures[ActionRun] = textures.MustNewSprite("assets/photos/chicken/sprint.png", 4, 0.25)
    chickenTextures[ActionWalk] = textures.MustNewSprite("assets/photos/chicken/walk.png", 4, 0.5)
    chickenTextures[ActionSquat] = textures.MustNewSprite("assets/photos/chicken/squat.png", 1, 0)
    chickenTextures[ActionPush] = textures.MustNewSprite("assets/photos/chicken/push.png", 4, 1)
    chickenTextures[ActionFall] = textures.MustNewSprite("assets/photos/chicken/fall.png", 2, 0.1)
}

// Chicken is the main character of this game. we ain't
// callin it chicky chicky for nothing folks
type Chicken struct {
	*world.PhysicalObject

	backpack *Backpack
	sprite *textures.Sprite
	action CharacterAction
	direction Direction

}

// NewChicken creates and initializes a new Chicken
func NewChicken() *Chicken {
	return &Chicken{sprite: chickenTextures[ActionNothing]}
}

// Move moves the chicken!
func (c *Chicken) Move(direction Direction, super bool)  {
    if super {
        c.action = ActionRun
    } else {
        c.action = ActionWalk
    }
    c.direction = DirectionLeft
}

// Jump jumps the chicken
func (c *Chicken) Jump(super bool) {
	if c.Hitbox != nil {
		c.ApplyForce(maths.Vec2{0, 6})
	}
}

// Down squats the chicken
func (c *Chicken) Down(super bool) {
    c.Stop()
    c.action = ActionSquat
}

// Stop stops the chicken's movement
func (c *Chicken) Stop() {
    c.action = ActionNothing
}

// Hit hits the chicken with the Item and power specified.
func (c *Chicken) Hit(with types.Item, power int) []types.Item {
    return []types.Item{}
}

// Kill kills the chicken, dropping its inventory
func (c *Chicken) Kill() []types.Item {
    tmp := c.inventory
    c.inventory = []types.Item{}
    return tmp
}

// Animate moves and calculates sprite frames for the
// Chicken
func (c *Chicken) Animate(delta float32) {
    
}

// Render renders the chicken onto the screen
func (c *Chicken) Render() {
	
}

package characters

import (
    "github.com/municorn/chicky-chicky-go/types"
    "github.com/municorn/chicky-chicky-go/world"
    "github.com/municorn/chicky-chicky-go/maths"
)

var chickenTextures = make(map[CharacterAction]*world.Sprite)

// InitGL initializes OpenGL-specific functionality for the
// characters package.
func InitGL() {
    chickenTextures[ActionNothing] = world.MustNewSprite("assets/photos/chicken/stand.png", 1, 0)
    chickenTextures[ActionRun] = world.MustNewSprite("assets/photos/chicken/sprint.png", 4, 0.25)
    chickenTextures[ActionWalk] = world.MustNewSprite("assets/photos/chicken/walk.png", 4, 0.5)
    chickenTextures[ActionSquat] = world.MustNewSprite("assets/photos/chicken/squat.png", 1, 0)
    chickenTextures[ActionPush] = world.MustNewSprite("assets/photos/chicken/push.png", 4, 1)
    chickenTextures[ActionFall] = world.MustNewSprite("assets/photos/chicken/fall.png", 2, 0.1)
}

// Chicken is the main character of this game. we ain't
// callin it chicky chicky for nothing folks
type Chicken struct {
	*world.PhysicalObject
	Character

	backpack *Backpack
	sprite *world.Sprite
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
		c.ApplyForce(maths.Vec2{X: 0, Y: 6})
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
func (c *Chicken) Hit(with types.Item, power float32) []types.Item {
    return []types.Item{}
}

// Kill kills the chicken, dropping its inventory
func (c *Chicken) Kill() []types.Item {
    tmp := c.backpack
    c.backpack = new(Backpack)
    return []types.Item(*tmp)
}

// // IsAlive returns true if the chicken is alive
// func (c *Chicken) IsAlive() bool {

// 	return true
// }

// Animate moves and calculates sprite frames for the
// Chicken
func (c *Chicken) Animate(delta float32) {
    
}

// Render renders the chicken onto the screen
func (c *Chicken) Render() {
	c.sprite.Render()
}

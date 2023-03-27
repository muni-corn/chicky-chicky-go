package world

import (
	"os"

	// "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/municorn/chicky-chicky-go/utils"
	"github.com/municorn/chicky-chicky-go/textures"

    mgl "github.com/go-gl/mathgl/mgl32"
)

// Sprite is an image that animates.
type Sprite struct {
    texture uint32
    frames int
    uvCoords []float32

    currentFrame float32
	secondsPerFrame float32

	sizeMatrix mgl.Mat4
	positionMatrix mgl.Mat4
    matrix mgl.Mat4
}



// TODO: initialize plane vbo
var planeVAO, planeVBO uint32

func init() {
	planeVAO, planeVBO = utils.NewTextureVAO(&planeVertices)
}

// NewSprite creates a new sprite and returns it
func NewSprite(spritePath string, frames int, secondsPerFrame float32) (s *Sprite, err error) {
	s = new(Sprite)

	// open the sprite file
	spriteFile, err := os.Open(spritePath)
	if err != nil {
		return
	}

	// assign the sprite texture
	s.texture, err = textures.NewTexture(spriteFile)
	if err != nil {
		return
	}

	// initialize the rest of the fields
	s.frames = frames
	s.secondsPerFrame = secondsPerFrame

	return
}

// MustNewSprite is like NewSprite, but panics if there's an
// error
func MustNewSprite(spritePath string, frames int, secondsPerFrame float32) *Sprite {
	sprite, err := NewSprite(spritePath, frames, secondsPerFrame)
	if err != nil {
		panic(err)
	}

	return sprite
}

// Animate animates the Sprite.
func (s *Sprite) Animate(delta float32) {
    s.currentFrame += delta / s.secondsPerFrame
	for s.currentFrame >= float32(s.frames) {
		s.currentFrame -= float32(s.frames)
	}
}

// SetSize sets the size of the sprite.
func (s *Sprite) SetSize(width, height float32) {
	s.sizeMatrix = mgl.Scale3D(width, height, 1)
	s.updateMatrix()
}

// SetPosition sets the position of the sprite.
func (s *Sprite) SetPosition(x, y, z float32) {
	s.positionMatrix = mgl.Translate3D(x, y, z)
	s.updateMatrix()
}

func (s *Sprite) updateMatrix() {
	s.matrix = s.positionMatrix.Mul4(s.sizeMatrix)
}

// Render renders the sprite onto the screen.
func (s *Sprite) Render() {
	// TODO:
	// - set transformation matrix in vertex shader
	// - uhh
	// - ???
	// - profit
}

var planeVertices = []float32 {
	// first triangle
	-0.5, 0.5, 0, 0, 0,
	-0.5, -0.5, 0, 0, 1,
	0.5, -0.5, 0, 1, 1,

	// second triangle
	-0.5, 0.5, 0, 0, 0,
	0.5, -0.5, 0, 1, 1,
	0.5, 0.5, 0, 1, 0,
}

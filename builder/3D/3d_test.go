package builder

import (
	"log"
	twod "testfile/builder/2D"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test3D(t *testing.T) {
	tube, err := NewTube(10, 3)
	log.Println(tube.Volume())
	require.NoError(t, err)
	require.NotEmpty(t, tube)

	circle, err := twod.GetShape("CIRCLE", 7)
	require.NoError(t, err)
	require.NotEmpty(t, circle)

	diagonal := twod.NewDiagonal(circle)
	log.Println(diagonal.GetDiagonal())
}

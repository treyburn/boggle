package reader

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	fp, err := filepath.Abs("./test.txt")
	require.NoError(t, err)

	data, err := Read(fp)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(data))
	assert.Equal(t, "abc", data[0])
}

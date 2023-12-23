package bitwarden

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/luismayta/envsecrets/v1/internal/app/config"
)

func TestSetFoldersIDs(t *testing.T) {
	conf := config.Initialize()
	bw := NewBW(conf)
	err := bw.SetFoldersIDs([]string{"Folder1", "Folder2"})

	require.NoError(t, err)
	err = bw.FetchItems()
	require.NoError(t, err)
	assert.Len(t, bw.values, 2)

	err = bw.SetFoldersIDs([]string{"NonexistentFolder"})
	require.Error(t, err)
	assert.Len(t, bw.values, 2)

}

func TestFetchItems(t *testing.T) {
	conf := config.Initialize()
	bw := NewBW(conf)

	bw.SetFoldersIDs([]string{"Folder1", "Folder2"})

	err := bw.FetchItems()
	require.NoError(t, err)

	bw.SetFoldersIDs([]string{})
	err = bw.FetchItems()
	require.Error(t, err)
}

func TestGenerateEnv(t *testing.T) {
	values := map[string]string{
		"Key1": "Value1",
		"Key2": "Value2",
	}
	conf := config.Initialize()
	bw := NewBWWithValues(conf, values)

	expected := "export Key1=Value1\nexport Key2=Value2\n"
	result := bw.GenerateEnv()
	assert.Equal(t, expected, result)
}

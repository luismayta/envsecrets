package bitwarden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFoldersIDs(t *testing.T) {
	bw := &BW{}

	err := bw.SetFoldersIDs([]string{"Folder1", "Folder2"})
	assert.NoError(t, err)
	assert.Len(t, bw.FetchItems(), 2)

	err = bw.SetFoldersIDs([]string{"NonexistentFolder"})
	assert.Error(t, err)
	assert.Len(t, bw.FetchItems(), 2)

}

func TestFetchItems(t *testing.T) {
	bw := &BW{}

	bw.SetFoldersIDs([]string{"Folder1", "Folder2"})

	err := bw.FetchItems()
	assert.NoError(t, err)

	bw.SetFoldersIDs([]string{})
	err = bw.FetchItems()
	assert.Error(t, err)
}

func TestGenerateEnv(t *testing.T) {
	bw := &BW{
		values: map[string]string{
			"Key1": "Value1",
			"Key2": "Value2",
		},
	}

	expected := "export Key1=Value1\nexport Key2=Value2\n"
	result := bw.GenerateEnv()
	assert.Equal(t, expected, result)
}
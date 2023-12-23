package bitwarden

import "github.com/luismayta/envsecrets/v1/internal/models"

type IClient interface {
	SetFoldersIDs(foldersNames []string) error

	FetchItems() error

	GenerateEnv() string
}

type IBW interface {
	SetFoldersIDs(foldersNames []string) error
	FetchItems() error
	GenerateEnv() string
	getFolderIDByName(name string) (models.FolderID, error)
	fetchItemsByFolderID(id models.FolderID) ([]models.Item, error)
	updateValues(items []models.Item)
}

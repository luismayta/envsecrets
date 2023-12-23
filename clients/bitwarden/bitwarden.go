package bitwarden

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/luismayta/envsecrets/v1/internal/app/common"
	"github.com/luismayta/envsecrets/v1/internal/app/config"
	"github.com/luismayta/envsecrets/v1/internal/app/models"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

const bwBinary = "bw"

// BW is a Bitwarden client.
type BW struct {
	conf       *config.Config
	foldersIDs []models.FolderID
	values     map[string]string
}

func NewBWWithValues(conf *config.Config, values map[string]string) *BW {
	if values == nil {
		err := errors.New(errors.ErrorInvalidArgument, "values cannot be nil")
		panic(err)

	}
	return &BW{
		conf: conf,
	}

}

// NewBW creates a new instance of BW.
func NewBW(conf *config.Config) *BW {
	return &BW{
		conf: conf,
	}
}

// SetFoldersIDs sets the folder IDs based on folder names.
func (bw *BW) SetFoldersIDs(foldersNames []string) error {
	for _, name := range foldersNames {
		folderID, err := bw.getFolderIDByName(name)
		if err != nil {
			return errors.Wrap(err, errors.ErrorInvalidArgument, "failed to set folder IDs")
		}
		bw.foldersIDs = append(bw.foldersIDs, folderID)
		log.Debugf("Folder ID found: %s\n", folderID)
	}
	return nil
}

// FetchItems fetches items from the specified folders.
func (bw *BW) FetchItems() error {
	if len(bw.foldersIDs) == 0 {
		return errors.New(
			errors.ErrorInvalidArgument,
			"unable to fetch from unknown folder, please set foldersIDs before",
		)
	}

	for _, id := range bw.foldersIDs {
		items, err := bw.fetchItemsByFolderID(id)
		if err != nil {
			return errors.Wrap(err, errors.ErrorInvalidArgument, "failed to fetch items")
		}
		bw.updateValues(items)
	}

	return nil
}

// GenerateEnv generates environment variable strings.
func (bw *BW) GenerateEnv() string {
	var str string
	for key, val := range bw.values {
		str += fmt.Sprintf("export %s=%s\n", key, val)
	}
	return str
}

func (bw *BW) getFolderIDByName(name string) (models.FolderID, error) {
	var results []models.FolderSearch
	err := common.ExecCLI(bwBinary, []string{"list", "folders", "--search", name}, &results)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrorInvalidArgument, "failed to get folder ID by name")
	}

	if len(results) == 0 {
		msg := fmt.Sprintf("no folder found with name %s", name)

		return "", errors.Wrap(
			err,
			errors.ErrorUnknown,
			msg,
		)
	}

	return results[0].ID, nil
}

func (bw *BW) fetchItemsByFolderID(id models.FolderID) ([]models.Item, error) {
	var results []models.Item
	err := common.ExecCLI(bwBinary, []string{"list", "items", "--folderid", string(id)}, &results)
	if err != nil {
		return nil, errors.Wrap(
			err,
			errors.ErrorInvalidArgument,
			"failed to fetch items by folder ID",
		)
	}

	return results, nil
}

func (bw *BW) updateValues(items []models.Item) {
	if bw.values == nil {
		bw.values = make(map[string]string)
	}

	for i := range items {
		item := &items[i]
		bw.values[item.Name] = item.Notes
	}
}

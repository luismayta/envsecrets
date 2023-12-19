package bitwarden

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/luismayta/envsecrets/v1/internal/app/common"
	"github.com/luismayta/envsecrets/v1/internal/app/models"
)

const bwBinary = "bw"

type BW struct {
	foldersIDs []models.FolderID
	values     map[string]string
}

func (bw *BW) SetFoldersIDs(foldersNames []string) error {
	for _, name := range foldersNames {
		folderID, err := bw.getFolderIDByName(name)
		if err != nil {
			return fmt.Errorf("envsecrets error: %v", err)
		}
		bw.foldersIDs = append(bw.foldersIDs, folderID)
		log.Debugf("Folder ID found: %s\n", folderID)
	}
	return nil
}

func (bw *BW) FetchItems() error {
	if len(bw.foldersIDs) == 0 {
		return fmt.Errorf(
			"envsecrets error: unable to fetch from unknown folder, please set foldersIDs before",
		)
	}
	for _, id := range bw.foldersIDs {
		items, err := bw.fetchItemsByFolderID(id)
		if err != nil {
			return fmt.Errorf("envsecrets error: %v", err)
		}
		bw.updateValues(items)
	}
	return nil
}

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
		return "", err
	}
	if len(results) == 0 {
		return "", fmt.Errorf("unable to match folder with name %s", name)
	}
	return results[0].ID, nil
}

func (bw *BW) fetchItemsByFolderID(id models.FolderID) ([]models.Item, error) {
	var results []models.Item
	err := common.ExecCLI(bwBinary, []string{"list", "items", "--folderid", string(id)}, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (bw *BW) updateValues(items []models.Item) {
	if bw.values == nil {
		bw.values = make(map[string]string)
	}
	for _, item := range items {
		bw.values[item.Name] = item.Notes
	}
}
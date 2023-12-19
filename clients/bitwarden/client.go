package bitwarden

type Client interface {
	SetFoldersIDs(foldersNames []string) error

	FetchItems() error

	GenerateEnv() string
}

func NewClient() Client {
	return &BW{}
}

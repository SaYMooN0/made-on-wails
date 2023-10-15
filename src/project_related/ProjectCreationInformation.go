package projectrelated

import src "made/src"

type ProjectCreationInformation struct {
	FolderPath string
	Name       string
	Version    string
	ModLoader  *src.Loader
}

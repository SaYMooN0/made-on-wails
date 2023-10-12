package minecraftrelated

type ProcessingType struct {
	Id                string `json:"id"`
	InGameName        string `json:"inGameName"`
	IsSupportedByMade bool   `json:"isSupportedByMade"`
}

func NewProcessingType(id, inGameName string, isSupportedByMade bool) *ProcessingType {
	return &ProcessingType{
		Id:                id,
		InGameName:        inGameName,
		IsSupportedByMade: isSupportedByMade,
	}
}

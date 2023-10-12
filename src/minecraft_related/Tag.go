package minecraftrelated

import "encoding/json"

type Tag struct {
	Id         string `json:"id"`
	InGameName string `json:"inGameName"`
}

func NewTag(id, inGameName string) *Tag {
	return &Tag{
		Id:         id,
		InGameName: inGameName,
	}
}

func NewTagWithId(id string) *Tag {
	return &Tag{
		Id: id,
	}
}

func (t *Tag) UnmarshalJSON(data []byte) error {
	type Alias Tag
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

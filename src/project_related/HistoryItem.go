package projectrelated

import (
	"fmt"
	src "made/src"
	"time"
)

type HistoryItem struct {
	Arguments  map[string]string
	FilePath   string
	ActionID   string
	ActionType src.ActionType
}

func NewHistoryItem(arguments map[string]string, filePath, actionID string, actionType src.ActionType) *HistoryItem {
	return &HistoryItem{
		Arguments:  arguments,
		FilePath:   filePath,
		ActionID:   actionID,
		ActionType: actionType,
	}
}

func (h *HistoryItem) GetCreationTime() *time.Time {
	t, err := time.Parse("060102150405", h.ActionID)
	if err != nil {
		return nil
	}
	return &t
}

func (h *HistoryItem) HistoryItemLabel() string {
	if input, ok := h.Arguments["input"]; ok {
		if output, ok := h.Arguments["output"]; ok {
			return fmt.Sprintf("'%s' to '%s'", input, output)
		}
	}

	if h.ActionType == src.CraftingTableAdd {
		shape := "shaped"
		if isShapeless, ok := h.Arguments["isShapeless"]; ok && isShapeless == "true" {
			shape = "shapeless"
		}
		if output, ok := h.Arguments["output"]; ok {
			return fmt.Sprintf("'%s' %s crafting recipe", output, shape)
		}
	}

	return ""
}

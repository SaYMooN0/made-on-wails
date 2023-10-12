package projectrelated

import (
	"fmt"
	"time"
)

type ActionType int

// you'll have to define the ActionType enumeration values
const (
	ActionTypeCraftingTableAdd ActionType = iota
	// ... other values
)

type HistoryItem struct {
	Arguments  map[string]string
	FilePath   string
	ActionID   string
	ActionType ActionType
}

func NewHistoryItem(arguments map[string]string, filePath, actionID string, actionType ActionType) *HistoryItem {
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

	if h.ActionType == ActionTypeCraftingTableAdd {
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

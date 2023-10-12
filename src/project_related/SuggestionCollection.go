package projectrelated

import (
	src "made/src"
	minecraft "made/src/minecraft_related"
	"strings"
)

type SuggestionsCollection struct {
	Loader src.Loader
	Mods   []minecraft.Mod
}

const TagChar byte = '#'

func NewSuggestionsCollection(loader src.Loader, mods []minecraft.Mod) *SuggestionsCollection {
	return &SuggestionsCollection{
		Loader: loader,
		Mods:   mods,
	}
}

func (sc *SuggestionsCollection) GetSuggestion(input string) []string {
	if input == "" {
		return []string{}
	}

	if !strings.Contains(input, ":") {
		var suggestions []string
		for _, mod := range sc.Mods {
			if strings.HasPrefix(input, string(TagChar)) {
				if strings.Contains(mod.Id, strings.TrimPrefix(input, string(TagChar))) {
					suggestions = append(suggestions, string(TagChar)+mod.Id)
				}
			} else {
				if strings.Contains(mod.Id, input) {
					suggestions = append(suggestions, mod.Id)
				}
			}
		}
		return suggestions
	}

	mod := sc.getModFromInput(input)
	if mod == nil {
		return []string{}
	}

	afterModIDString := strings.Split(input, ":")[1]

	if input[0] == TagChar {
		var tagSuggestions []string
		for _, tag := range mod.Tags {
			if strings.Contains(tag.Id, afterModIDString) {
				tagSuggestions = append(tagSuggestions, string(TagChar)+mod.Id+":"+tag.Id)
			}
		}
		return tagSuggestions
	}

	var itemSuggestions []string
	for _, item := range mod.Items {
		if strings.Contains(item.Id, afterModIDString) {
			itemSuggestions = append(itemSuggestions, mod.Id+":"+item.Id)
		}
	}
	return itemSuggestions
}

func (sc *SuggestionsCollection) getModFromInput(input string) *minecraft.Mod {
	modString := strings.Split(input, ":")[0]
	if input[0] == TagChar {
		modString = strings.TrimPrefix(modString, string(TagChar))
	}

	for _, mod := range sc.Mods {
		if mod.Id == modString {
			return &mod
		}
	}

	return nil
}

package made

import (
	"encoding/json"
	"os"
)

type ThemeCollection struct {
	Themes           []Theme `json:"Themes"`
	CurrentThemeName string  `json:"_currentThemeName"`
	ThemeNames       []string
	FileName         string `json:"-"`
}

func (tc *ThemeCollection) CurrentTheme() *Theme {
	for _, theme := range tc.Themes {
		if theme.Name == tc.CurrentThemeName {
			return &theme
		}
	}
	if len(tc.Themes) > 0 {
		return &tc.Themes[0]
	}
	defaultTheme := tc.defaultDark()
	tc.Themes = append(tc.Themes, *defaultTheme)
	return defaultTheme
}

func (tc *ThemeCollection) SetCurrentTheme(theme Theme) {
	tc.CurrentThemeName = theme.Name
}

func InitializeThemeCollection() *ThemeCollection {
	if _, err := os.Stat("themes.madeThemes"); os.IsNotExist(err) {
		CreateDefaultThemesFile()
	}
	return LoadFromFile()
}

func (tc *ThemeCollection) defaultDark() *Theme {
	return NewTheme(
		"default dark",
		FromHex("#272527"),
		FromHex("#2F2E37"),
		FromHex("#3B3946"),
		FromHex("#D8D3DC"),
		FromHex("#D4BFFD"),
		FromHex("#A782EA"),
		FromHex("#8862D5"),
		FromHex("#7E54D5"),
		FromHex("#5A39AD"),
		FromHex("#7E032A"),
		FromHex("#930336"),
	)
}

func (tc *ThemeCollection) defaultLight() *Theme {
	return NewTheme(
		"default light",
		FromHex("#E0E1DD"),
		FromHex("#D5BDAF"),
		FromHex("#C9B9B0"),
		FromHex("#457B9D"),
		FromHex("#1F4558"),
		FromHex("#1D3557"),
		FromHex("#171020"),
		FromHex("#222B48"),
		FromHex("#262D4A"),
		FromHex("#7E032A"),
		FromHex("#930336"),
	)
}
func CreateDefaultThemesFile() {
	themeCollection := &ThemeCollection{}
	dark := themeCollection.defaultDark()
	light := themeCollection.defaultLight()
	themeCollection.Themes = append(themeCollection.Themes, *dark, *light)
	themeCollection.ThemeNames = append(themeCollection.ThemeNames, dark.Name, light.Name)
	themeCollection.CurrentThemeName = dark.Name
	themeCollection.SaveToFile()
}

func LoadFromFile() *ThemeCollection {
	data, err := os.ReadFile("themes.madeThemes")
	if err != nil {
		panic(err)
	}
	var themes ThemeCollection
	err = json.Unmarshal(data, &themes)
	if err != nil {
		panic(err)
	}
	return &themes
}

func (tc *ThemeCollection) SaveToFile() {
	data, err := json.MarshalIndent(tc, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("themes.madeThemes", data, 0644)
	if err != nil {
		panic(err)
	}
}

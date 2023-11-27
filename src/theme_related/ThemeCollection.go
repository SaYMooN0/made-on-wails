package themerelated

import (
	"encoding/json"
	"os"
)

type ThemeCollection struct {
	Themes           []Theme `json:"Themes"`
	CurrentThemeName string  `json:"_currentThemeName"`
	FileName         string  `json:"-"`
}

func (tc *ThemeCollection) UpdateTheme(name string, newTheme Theme) string {
	for i, theme := range tc.Themes {
		if theme.Name == name {
			tc.Themes[i] = newTheme
			tc.SaveToFile()
			return ""
		}
	}
	return "Failed to update the theme"
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
	defaultTheme := DefaultDark()
	tc.Themes = append(tc.Themes, *defaultTheme)
	return defaultTheme
}
func (tc *ThemeCollection) GetAllThemesValues() []ThemeHex {
	var themesArray []ThemeHex

	for _, theme := range tc.Themes {
		themeHex := NewThemeHex(theme)
		themesArray = append(themesArray, *themeHex)
	}

	return themesArray
}
func (tc *ThemeCollection) CurrentThemeColors() *ThemeHex {
	current := tc.CurrentTheme()
	if current == nil {
		return nil
	}
	return NewThemeHex(*current)
}

func (tc *ThemeCollection) SetCurrentTheme(themeName string) string {
	for _, theme := range tc.Themes {
		if theme.Name == themeName {
			tc.CurrentThemeName = themeName
			tc.SaveToFile()
			return ""
		}
	}
	return "Theme with the given name does not exist"
}

func InitializeThemeCollection() *ThemeCollection {
	if _, err := os.Stat("themes.madeThemes"); os.IsNotExist(err) {
		CreateDefaultThemesFile()
	}
	return LoadFromFile()
}

func CreateDefaultThemesFile() {
	themeCollection := &ThemeCollection{}
	dark := DefaultDark()
	light := DefaultLight()
	themeCollection.Themes = append(themeCollection.Themes, *dark, *light)
	themeCollection.CurrentThemeName = dark.Name
	err := themeCollection.SaveToFile()
	if err != nil {
		errMessage := "Error in CreateDefaultThemesFile: " + err.Error() + "\n"
		os.WriteFile("error.txt", []byte(errMessage), 0644)
	}
}

func LoadFromFile() *ThemeCollection {
	data, err := os.ReadFile("themes.madeThemes")
	if err != nil {
		errMessage := "Error: LoadFromFile " + err.Error() + "\n"
		os.WriteFile("error.txt", []byte(errMessage), 0644)
		return nil
	}

	var themes ThemeCollection
	err = json.Unmarshal(data, &themes)
	if err != nil {
		errMessage := "Error: LoadFromFile json " + err.Error() + "\n"
		os.WriteFile("error.txt", []byte(errMessage), 0644)
		return nil
	}
	uniqueThemes := make([]Theme, 0)
	seen := make(map[string]bool)

	for _, theme := range themes.Themes {
		if _, found := seen[theme.Name]; !found {
			uniqueThemes = append(uniqueThemes, theme)
			seen[theme.Name] = true
		}
	}

	themes.Themes = uniqueThemes
	return &themes
}

func (tc *ThemeCollection) SaveToFile() error {
	data, err := json.MarshalIndent(tc, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("themes.madeThemes", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (tc *ThemeCollection) ThemeFromHexTheme(th ThemeHex) *Theme {
	return NewTheme(
		th.Name,
		FromHex(th.MainBackColor),
		FromHex(th.SecondBackColor),
		FromHex(th.ThirdBackColor),
		FromHex(th.MainFrontColor),
		FromHex(th.SecondFrontColor),
		FromHex(th.ThirdFrontColor),
		FromHex(th.MainBrightColor),
		FromHex(th.SecondBrightColor),
		FromHex(th.ThirdBrightColor),
		FromHex(th.WarningMainColor),
		FromHex(th.WarningBrightColor),
	)
}
func (tc *ThemeCollection) GetHexDefaultDark() *ThemeHex {
	return NewThemeHex(*DefaultDark())
}
func NewThemeFromHexValues(name, mainBack, secondBack, thirdBack, mainFront, secondFront, thirdFront, mainBright, secondBright, thirdBright, warningMain, warningBright string) *Theme {
	return NewTheme(
		name,
		FromHex(mainBack),
		FromHex(secondBack),
		FromHex(thirdBack),
		FromHex(mainFront),
		FromHex(secondFront),
		FromHex(thirdFront),
		FromHex(mainBright),
		FromHex(secondBright),
		FromHex(thirdBright),
		FromHex(warningMain),
		FromHex(warningBright),
	)
}

func DefaultDark() *Theme {
	return NewThemeFromHexValues(
		"default dark",
		"#272527",
		"#2F2E37",
		"#3B3946",
		"#D8D3DC",
		"#D4BFFD",
		"#A782EA",
		"#8862D5",
		"#7E54D5",
		"#5A39AD",
		"#7E032A",
		"#930336",
	)
}

func DefaultLight() *Theme {
	return NewThemeFromHexValues(
		"default light",
		"#E0E1DD",
		"#D5BDAF",
		"#C9B9B0",
		"#457B9D",
		"#1F4558",
		"#1D3557",
		"#171020",
		"#222B48",
		"#262D4A",
		"#7E032A",
		"#930336",
	)
}

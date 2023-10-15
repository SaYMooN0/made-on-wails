package themerelated

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
	defaultTheme := DefaultDark()
	tc.Themes = append(tc.Themes, *defaultTheme)
	return defaultTheme
}
func (tc *ThemeCollection) CurrentThemeColors() map[string]string {
	current := tc.CurrentTheme()
	if current == nil {
		return nil
	}

	colors := map[string]string{
		"MainBackColor":      current.MainBack.ToHex(),
		"SecondBackColor":    current.SecondBack.ToHex(),
		"ThirdBackColor":     current.ThirdBack.ToHex(),
		"MainFrontColor":     current.MainFront.ToHex(),
		"SecondFrontColor":   current.SecondFront.ToHex(),
		"ThirdFrontColor":    current.ThirdFront.ToHex(),
		"MainBrightColor":    current.MainBright.ToHex(),
		"SecondBrightColor":  current.SecondBright.ToHex(),
		"ThirdBrightColor":   current.ThirdBright.ToHex(),
		"WarningMainColor":   current.WarningMain.ToHex(),
		"WarningBrightColor": current.WarningBright.ToHex(),
	}

	return colors
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

func CreateDefaultThemesFile() {
	themeCollection := &ThemeCollection{}
	dark := DefaultDark()
	light := DefaultLight()
	themeCollection.Themes = append(themeCollection.Themes, *dark, *light)
	themeCollection.ThemeNames = append(themeCollection.ThemeNames, dark.Name, light.Name)
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
	}
	var themes ThemeCollection
	err = json.Unmarshal(data, &themes)
	if err != nil {
		errMessage := "Error: LoadFromFile json " + err.Error() + "\n"
		os.WriteFile("error.txt", []byte(errMessage), 0644)
	}
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
func DefaultDark() *Theme {
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

func DefaultLight() *Theme {
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

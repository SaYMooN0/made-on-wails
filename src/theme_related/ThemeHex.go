package themerelated

type ThemeHex struct {
	Name               string `json:"Name"`
	MainBackColor      string `json:"MainBackColor"`
	SecondBackColor    string `json:"SecondBackColor"`
	ThirdBackColor     string `json:"ThirdBackColor"`
	MainFrontColor     string `json:"MainFrontColor"`
	SecondFrontColor   string `json:"SecondFrontColor"`
	ThirdFrontColor    string `json:"ThirdFrontColor"`
	MainBrightColor    string `json:"MainBrightColor"`
	SecondBrightColor  string `json:"SecondBrightColor"`
	ThirdBrightColor   string `json:"ThirdBrightColor"`
	WarningMainColor   string `json:"WarningMainColor"`
	WarningBrightColor string `json:"WarningBrightColor"`
}

func NewThemeHex(t Theme) *ThemeHex {
	return &ThemeHex{
		Name:               t.Name,
		MainBackColor:      t.MainBack.ToHex(),
		SecondBackColor:    t.SecondBack.ToHex(),
		ThirdBackColor:     t.ThirdBack.ToHex(),
		MainFrontColor:     t.MainFront.ToHex(),
		SecondFrontColor:   t.SecondFront.ToHex(),
		ThirdFrontColor:    t.ThirdFront.ToHex(),
		MainBrightColor:    t.MainBright.ToHex(),
		SecondBrightColor:  t.SecondBright.ToHex(),
		ThirdBrightColor:   t.ThirdBright.ToHex(),
		WarningMainColor:   t.WarningMain.ToHex(),
		WarningBrightColor: t.WarningBright.ToHex(),
	}
}

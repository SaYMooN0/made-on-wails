package themerelated

type ThemeHex struct {
	Name               string `json:"Name"`
	BackMainColor      string `json:"BackMainColor"`
	BackSecondColor    string `json:"BackSecondColor"`
	BackThirdColor     string `json:"BackThirdColor"`
	FrontMainColor     string `json:"FrontMainColor"`
	FrontSecondColor   string `json:"FrontSecondColor"`
	FrontThirdColor    string `json:"FrontThirdColor"`
	BrightMainColor    string `json:"BrightMainColor"`
	BrightSecondColor  string `json:"BrightSecondColor"`
	BrightThirdColor   string `json:"BrightThirdColor"`
	WarningMainColor   string `json:"WarningMainColor"`
	WarningBrightColor string `json:"WarningBrightColor"`
}

func NewThemeHex(t Theme) *ThemeHex {
	return &ThemeHex{
		Name:               t.Name,
		BackMainColor:      t.MainBack.ToHex(),
		BackSecondColor:    t.SecondBack.ToHex(),
		BackThirdColor:     t.ThirdBack.ToHex(),
		FrontMainColor:     t.MainFront.ToHex(),
		FrontSecondColor:   t.SecondFront.ToHex(),
		FrontThirdColor:    t.ThirdFront.ToHex(),
		BrightMainColor:    t.MainBright.ToHex(),
		BrightSecondColor:  t.SecondBright.ToHex(),
		BrightThirdColor:   t.ThirdBright.ToHex(),
		WarningMainColor:   t.WarningMain.ToHex(),
		WarningBrightColor: t.WarningBright.ToHex(),
	}
}

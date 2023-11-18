package themerelated

type Theme struct {
	Name          string `json:"Name"`
	MainBack      Color  `json:"MainBackColor"`
	SecondBack    Color  `json:"SecondBackColor"`
	ThirdBack     Color  `json:"ThirdBackColor"`
	MainFront     Color  `json:"MainFrontColor"`
	SecondFront   Color  `json:"SecondFrontColor"`
	ThirdFront    Color  `json:"ThirdFrontColor"`
	MainBright    Color  `json:"MainBrightColor"`
	SecondBright  Color  `json:"SecondBrightColor"`
	ThirdBright   Color  `json:"ThirdBrightColor"`
	WarningMain   Color  `json:"WarningMainColor"`
	WarningBright Color  `json:"WarningBrightColor"`
}

func NewTheme(name string, mainBack, secondBack, thirdBack, mainFront, secondFront, thirdFront, mainBright, secondBright, thirdBright, warningMain, warningBright Color) *Theme {
	return &Theme{
		Name:          name,
		MainBack:      mainBack,
		SecondBack:    secondBack,
		ThirdBack:     thirdBack,
		MainFront:     mainFront,
		SecondFront:   secondFront,
		ThirdFront:    thirdFront,
		MainBright:    mainBright,
		SecondBright:  secondBright,
		ThirdBright:   thirdBright,
		WarningMain:   warningMain,
		WarningBright: warningBright,
	}
}
func (t *Theme) GetValues() map[string]string {
	return map[string]string{
		"Name":               t.Name,
		"MainBackColor":      t.MainBack.ToHex(),
		"SecondBackColor":    t.SecondBack.ToHex(),
		"ThirdBackColor":     t.ThirdBack.ToHex(),
		"MainFrontColor":     t.MainFront.ToHex(),
		"SecondFrontColor":   t.SecondFront.ToHex(),
		"ThirdFrontColor":    t.ThirdFront.ToHex(),
		"MainBrightColor":    t.MainBright.ToHex(),
		"SecondBrightColor":  t.SecondBright.ToHex(),
		"ThirdBrightColor":   t.ThirdBright.ToHex(),
		"WarningMainColor":   t.WarningMain.ToHex(),
		"WarningBrightColor": t.WarningBright.ToHex(),
	}
}

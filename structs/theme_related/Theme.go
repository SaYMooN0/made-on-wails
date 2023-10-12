package made

type Theme struct {
	Name          string
	mainBack      Color
	secondBack    Color
	thirdBack     Color
	mainFront     Color
	secondFront   Color
	thirdFront    Color
	mainBright    Color
	secondBright  Color
	thirdBright   Color
	warningMain   Color
	warningBright Color
}

func NewTheme(name string, mainBack, secondBack, thirdBack, mainFront, secondFront, thirdFront, mainBright, secondBright, thirdBright, warningMain, warningBright Color) *Theme {
	return &Theme{
		Name:          name,
		mainBack:      mainBack,
		secondBack:    secondBack,
		thirdBack:     thirdBack,
		mainFront:     mainFront,
		secondFront:   secondFront,
		thirdFront:    thirdFront,
		mainBright:    mainBright,
		secondBright:  secondBright,
		thirdBright:   thirdBright,
		warningMain:   warningMain,
		warningBright: warningBright,
	}
}

func (t *Theme) MainBackColor() string {
	return t.mainBack.ToHex()
}
func (t *Theme) SetMainBackColor(hex string) {
	t.mainBack = FromHex(hex)
}
func (t *Theme) SecondBackColor() string {
	return t.secondBack.ToHex()
}
func (t *Theme) SetSecondBackColor(hex string) {
	t.secondBack = FromHex(hex)
}
func (t *Theme) ThirdBackColor() string {
	return t.thirdBack.ToHex()
}
func (t *Theme) SetThirdBackColor(hex string) {
	t.thirdBack = FromHex(hex)
}
func (t *Theme) MainFrontColor() string {
	return t.mainFront.ToHex()
}
func (t *Theme) SetMainFrontColor(hex string) {
	t.mainFront = FromHex(hex)
}
func (t *Theme) SecondFrontColor() string {
	return t.secondFront.ToHex()
}
func (t *Theme) SetSecondFrontColor(hex string) {
	t.secondFront = FromHex(hex)
}
func (t *Theme) ThirdFrontColor() string {
	return t.thirdFront.ToHex()
}
func (t *Theme) SetThirdFrontColor(hex string) {
	t.thirdFront = FromHex(hex)
}
func (t *Theme) MainBrightColor() string {
	return t.mainBright.ToHex()
}
func (t *Theme) SetMainBrightColor(hex string) {
	t.mainBright = FromHex(hex)
}
func (t *Theme) SecondBrightColor() string {
	return t.secondBright.ToHex()
}
func (t *Theme) SetSecondBrightColor(hex string) {
	t.secondBright = FromHex(hex)
}
func (t *Theme) ThirdBrightColor() string {
	return t.thirdBright.ToHex()
}
func (t *Theme) SetThirdBrightColor(hex string) {
	t.thirdBright = FromHex(hex)
}
func (t *Theme) WarningMainColor() string {
	return t.warningMain.ToHex()
}
func (t *Theme) SetWarningMainColor(hex string) {
	t.warningMain = FromHex(hex)
}
func (t *Theme) WarningBrightColor() string {
	return t.warningBright.ToHex()
}
func (t *Theme) SetWarningBrightColor(hex string) {
	t.warningBright = FromHex(hex)
}

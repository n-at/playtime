package storage

import (
	"strings"
	"time"
)

type User struct {
	Id       string `boltholdKey:"Id"`
	Login    string `boltholdIndex:"Login"`
	Password string
	Admin    bool
	Active   bool
}

type Session struct {
	Id      string `boltholdKey:"Id"`
	UserId  string `boltholdIndex:"UserId"`
	Created time.Time
}

type Settings struct {
	UserId           string `boltholdKey:"UserId"`
	Language         string
	EmulatorSettings map[string]EmulatorSettings
}

type Game struct {
	Id                       string `boltholdKey:"Id"`
	UserId                   string `boltholdIndex:"UserId"`
	Name                     string
	OriginalFileName         string
	OriginalFileExtension    string
	Platform                 string `boltholdIndex:"Platform"`
	NetplayEnabled           bool
	NetplaySessionId         string
	DisableCue               bool
	OverrideEmulatorSettings bool
	EmulatorSettings         EmulatorSettings
}

type SaveState struct {
	Id      string `boltholdKey:"Id"`
	UserId  string `boltholdIndex:"UserId"`
	GameId  string `boltholdIndex:"GameId"`
	Created time.Time
	Core    string
}

type UploadBatch struct {
	Id      string `boltholdKey:"Id"`
	UserId  string `boltholdIndex:"UserId"`
	Created time.Time
	GameIds []string
}

type EmulatorSettings struct {
	Core                   string
	Bios                   string
	ColorScheme            string
	ColorBackground        string
	CacheLimit             int64
	Volume                 float64
	Shader                 string
	FPS                    bool
	VirtualGamepadLeftHand bool
	StartFullScreen        bool
	Controls               [4]EmulatorControls
	Buttons                EmulatorButtons
	CoreOptions            map[string]string
	FastForwardRatio       string
	SlowMotionRatio        string
	RewindGranularity      string
	FastForwardMode        bool
	SlowMotionMode         bool
	Rewind                 bool
	Threads                bool
}

type EmulatorControls struct {
	Keyboard EmulatorControlsMapping
	Gamepad  EmulatorControlsMapping
}

type EmulatorControlsMapping struct {
	B               string //0
	Y               string //1
	Select          string //2
	Start           string //3
	Up              string //4
	Down            string //5
	Left            string //6
	Right           string //7
	A               string //8
	X               string //9
	L               string //10
	R               string //11
	L2              string //12
	R2              string //13
	L3              string //14
	R3              string //15
	LStickUp        string //19
	LStickDown      string //18
	LStickLeft      string //17
	LStickRight     string //16
	RStickUp        string //23
	RStickDown      string //22
	RStickLeft      string //21
	RStickRight     string //20
	QuickSaveState  string //24
	QuickLoadState  string //25
	ChangeStateSlot string //26
	FastForward     string //27
	Rewind          string //28
	SlowMotion      string //29
}

type EmulatorButtons struct {
	PlayPause    bool
	Restart      bool
	Mute         bool
	Settings     bool
	FullScreen   bool
	SaveState    bool
	LoadState    bool
	ScreenRecord bool
	Gamepad      bool
	Cheat        bool
	Volume       bool
	SaveSavFiles bool
	LoadSavFiles bool
	QuickSave    bool
	QuickLoad    bool
	Screenshot   bool
	CacheManager bool
}

///////////////////////////////////////////////////////////////////////////////

const (
	GamepadControlB           = "BUTTON_2"
	GamepadControlY           = "BUTTON_4"
	GamepadControlSelect      = "SELECT"
	GamepadControlStart       = "START"
	GamepadControlUp          = "DPAD_UP"
	GamepadControlDown        = "DPAD_DOWN"
	GamepadControlLeft        = "DPAD_LEFT"
	GamepadControlRight       = "DPAD_RIGHT"
	GamepadControlA           = "BUTTON_1"
	GamepadControlX           = "BUTTON_3"
	GamepadControlL           = "LEFT_TOP_SHOULDER"
	GamepadControlR           = "RIGHT_TOP_SHOULDER"
	GamepadControlL2          = "LEFT_BOTTOM_SHOULDER"
	GamepadControlR2          = "RIGHT_BOTTOM_SHOULDER"
	GamepadControlL3          = "LEFT_STICK"
	GamepadControlR3          = "RIGHT_STICK"
	GamepadControlLStickUp    = "LEFT_STICK_Y:-1"
	GamepadControlLStickDown  = "LEFT_STICK_Y:+1"
	GamepadControlLStickLeft  = "LEFT_STICK_X:-1"
	GamepadControlLStickRight = "LEFT_STICK_X:+1"
	GamepadControlRStickUp    = "RIGHT_STICK_Y:-1"
	GamepadControlRStickDown  = "RIGHT_STICK_Y:+1"
	GamepadControlRStickLeft  = "RIGHT_STICK_X:-1"
	GamepadControlRStickRight = "RIGHT_STICK_X:+1"
)

//goland:noinspection GoUnusedConst
const (
	KeyTab            = "9"
	KeyNumpadEqual    = "12"
	KeyEnter          = "13"
	KeyShift          = "16"
	KeyControl        = "17"
	KeyAlt            = "18"
	KeyPause          = "19"
	KeyCapsLock       = "20"
	KeyEscape         = "27"
	KeySpace          = "32"
	KeyPageUp         = "33"
	KeyPageDown       = "34"
	KeyEnd            = "35"
	KeyHome           = "36"
	KeyArrowLeft      = "37"
	KeyArrowUp        = "38"
	KeyArrowRight     = "39"
	KeyArrowDown      = "40"
	KeyPrintScreen    = "44"
	KeyInsert         = "45"
	KeyDelete         = "46"
	KeyDigit0         = "48"
	KeyDigit1         = "49"
	KeyDigit2         = "50"
	KeyDigit3         = "51"
	KeyDigit4         = "52"
	KeyDigit5         = "53"
	KeyDigit6         = "54"
	KeyDigit7         = "55"
	KeyDigit8         = "56"
	KeyDigit9         = "57"
	KeyA              = "65"
	KeyB              = "66"
	KeyC              = "67"
	KeyD              = "68"
	KeyE              = "69"
	KeyF              = "70"
	KeyG              = "71"
	KeyH              = "72"
	KeyI              = "73"
	KeyJ              = "74"
	KeyK              = "75"
	KeyL              = "76"
	KeyM              = "77"
	KeyN              = "78"
	KeyO              = "79"
	KeyP              = "80"
	KeyQ              = "81"
	KeyR              = "82"
	KeyS              = "83"
	KeyT              = "84"
	KeyU              = "85"
	KeyV              = "86"
	KeyW              = "87"
	KeyX              = "88"
	KeyY              = "89"
	KeyZ              = "90"
	KeyMetaLeft       = "91"
	KeyMetaRight      = "92"
	KeyContextMenu    = "93"
	KeyNumpad0        = "96"
	KeyNumpad1        = "97"
	KeyNumpad2        = "98"
	KeyNumpad3        = "99"
	KeyNumpad4        = "100"
	KeyNumpad5        = "101"
	KeyNumpad6        = "102"
	KeyNumpad7        = "103"
	KeyNumpad8        = "104"
	KeyNumpad9        = "105"
	KeyNumpadMultiply = "106"
	KeyNumpadAdd      = "107"
	KeyNumpadSubtract = "109"
	KeyDecimal        = "110"
	KeyNumpadDivide   = "111"
	KeyF1             = "112"
	KeyF2             = "113"
	KeyF3             = "114"
	KeyF4             = "115"
	KeyF5             = "116"
	KeyF6             = "117"
	KeyF7             = "118"
	KeyF8             = "119"
	KeyF9             = "120"
	KeyF10            = "121"
	KeyF11            = "122"
	KeyF12            = "123"
	KeyF13            = "124"
	KeyF14            = "125"
	KeyF15            = "126"
	KeyF16            = "127"
	KeyF17            = "128"
	KeyF18            = "129"
	KeyF19            = "130"
	KeyF20            = "131"
	KeyF21            = "132"
	KeyF22            = "133"
	KeyF23            = "134"
	KeyF24            = "135"
	KeyNumLock        = "144"
	KeyScrollLock     = "145"
	KeySemicolon      = "186"
	KeyEqual          = "187"
	KeyComma          = "188"
	KeyMinus          = "189"
	KeyPeriod         = "190"
	KeyBackquote      = "192"
	KeyIntlRo         = "193"
	KeyNumpadComma    = "194"
	KeyBracketLeft    = "219"
	KeyBackslash      = "220"
	KeyBracketRight   = "221"
	KeyQuote          = "222"
	KeyIntlYen        = "255"
)

type Platform struct {
	Id         string
	Name       string
	Extensions []string
}

type Bios struct {
	Name        string
	Url         string
	Hash        string
	Description string
}

type Shader struct {
	Name  string
	Value string
}

type CoreOption struct {
	Id       string
	Name     string
	Variants string
	Default  string
}

type GameWithData struct {
	Game
	DownloadLink    string
	LatestSaveState SaveStateWithData
}

type SaveStateWithData struct {
	SaveState
	StateFileDownloadLink  string
	ScreenshotDownloadLink string
}

///////////////////////////////////////////////////////////////////////////////

func (u *User) CanControlUsers() bool {
	return u.Admin
}

func (co CoreOption) VariantsList() []string {
	return strings.Split(co.Variants, "|")
}

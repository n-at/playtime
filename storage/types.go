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
	KeyTab            = "tab"
	KeyEnter          = "enter"
	KeyShift          = "shift"
	KeyControl        = "ctrl"
	KeyAlt            = "alt"
	KeyPause          = "pause/break"
	KeyCapsLock       = "caps lock"
	KeyEscape         = "escape"
	KeySpace          = "space"
	KeyPageUp         = "page up"
	KeyPageDown       = "page down"
	KeyEnd            = "end"
	KeyHome           = "home"
	KeyArrowLeft      = "left arrow"
	KeyArrowUp        = "up arrow"
	KeyArrowRight     = "right arrow"
	KeyArrowDown      = "down arrow"
	KeyInsert         = "insert"
	KeyDelete         = "delete"
	KeyDigit0         = "0"
	KeyDigit1         = "1"
	KeyDigit2         = "2"
	KeyDigit3         = "3"
	KeyDigit4         = "4"
	KeyDigit5         = "5"
	KeyDigit6         = "6"
	KeyDigit7         = "7"
	KeyDigit8         = "8"
	KeyDigit9         = "9"
	KeyA              = "a"
	KeyB              = "b"
	KeyC              = "c"
	KeyD              = "d"
	KeyE              = "e"
	KeyF              = "f"
	KeyG              = "g"
	KeyH              = "h"
	KeyI              = "i"
	KeyJ              = "j"
	KeyK              = "k"
	KeyL              = "l"
	KeyM              = "m"
	KeyN              = "n"
	KeyO              = "o"
	KeyP              = "p"
	KeyQ              = "q"
	KeyR              = "r"
	KeyS              = "s"
	KeyT              = "t"
	KeyU              = "u"
	KeyV              = "v"
	KeyW              = "w"
	KeyX              = "x"
	KeyY              = "y"
	KeyZ              = "z"
	KeyMetaLeft       = "left window key"
	KeyMetaRight      = "right window key"
	KeyContextMenu    = "select key"
	KeyNumpad0        = "numpad 0"
	KeyNumpad1        = "numpad 1"
	KeyNumpad2        = "numpad 2"
	KeyNumpad3        = "numpad 3"
	KeyNumpad4        = "numpad 4"
	KeyNumpad5        = "numpad 5"
	KeyNumpad6        = "numpad 6"
	KeyNumpad7        = "numpad 7"
	KeyNumpad8        = "numpad 8"
	KeyNumpad9        = "numpad 9"
	KeyNumpadMultiply = "multiply"
	KeyNumpadAdd      = "add"
	KeyNumpadSubtract = "subtract"
	KeyDecimal        = "decimal point"
	KeyNumpadDivide   = "divide"
	KeyF1             = "f1"
	KeyF2             = "f2"
	KeyF3             = "f3"
	KeyF4             = "f4"
	KeyF5             = "f5"
	KeyF6             = "f6"
	KeyF7             = "f7"
	KeyF8             = "f8"
	KeyF9             = "f9"
	KeyF10            = "f10"
	KeyF11            = "f11"
	KeyF12            = "f12"
	KeyF13            = "f13"
	KeyF14            = "f14"
	KeyF15            = "f15"
	KeyF16            = "f16"
	KeyF17            = "f17"
	KeyF18            = "f18"
	KeyF19            = "f19"
	KeyF20            = "f20"
	KeyF21            = "f21"
	KeyF22            = "f22"
	KeyF23            = "f23"
	KeyF24            = "f24"
	KeyNumLock        = "num lock"
	KeyScrollLock     = "scroll lock"
	KeySemicolon      = "semi-colon"
	KeyEqual          = "equal sign"
	KeyComma          = "comma"
	KeyMinus          = "dash"
	KeyPeriod         = "period"
	KeySlash          = "forward slash"
	KeyBackquote      = "grave accent"
	KeyBracketLeft    = "open bracket"
	KeyBackslash      = "back slash"
	KeyBracketRight   = "close braket"
	KeyQuote          = "single quote"
)

type Platform struct {
	Id         string
	Name       string
	Cores      []string
	Extensions []string
	Bios       []Bios
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

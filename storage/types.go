package storage

import "time"

type User struct {
	Id       string `boltholdKey:"Id"`
	Login    string `boltholdIndex:"Login"`
	Password string
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
	Type                     string `boltholdIndex:"Type"`
	OverrideEmulatorSettings bool
	EmulatorSettings         EmulatorSettings
}

type SaveState struct {
	Id      string `boltholdKey:"Id"`
	UserId  string `boltholdIndex:"UserId"`
	GameId  string `boltholdIndex:"GameId"`
	Created time.Time
}

type LoadBatch struct {
	Id      string `boltholdKey:"Id"`
	UserId  string `boltholdIndex:"UserId"`
	Created time.Time
	GameIds []string
}

type EmulatorSettings struct {
	OldCores               bool
	Core                   string
	Bios                   string
	ColorScheme            string
	CacheLimit             int64
	Volume                 float32
	Shader                 string
	FPS                    bool
	VirtualGamepadLeftHand bool
	Controls               [4]EmulatorControls
	Buttons                EmulatorButtons
}

type EmulatorControls struct {
	Keyboard EmulatorControlsMapping
	Gamepad  EmulatorControlsMapping
}

type EmulatorControlsMapping struct {
	B               string
	Y               string
	Select          string
	Start           string
	Up              string
	Down            string
	Left            string
	Right           string
	A               string
	X               string
	L               string
	R               string
	L2              string
	R2              string
	L3              string
	R3              string
	LStickUp        string
	LStickDown      string
	LStickLeft      string
	LStickRight     string
	RStickUp        string
	RStickDown      string
	RStickLeft      string
	RStickRight     string
	QuickSaveState  string
	QuickLoadState  string
	ChangeStateSlot string
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

type GameSystem struct {
	Name string
}

type Bios struct {
	Name        string
	Url         string
	Hash        string
	Description string
}

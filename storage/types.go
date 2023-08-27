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

package storage

const (
	DefaultColorScheme       = "#1AAFFF"
	DefaultColorBackground   = "#000"
	DefaultCacheLimit        = 1073741824
	DefaultVolume            = 0.5
	DefaultFastForwardRatio  = "3.0"
	DefaultSlowMotionRatio   = "3.0"
	DefaultRewindGranularity = "6"
	DefaultLanguage          = "en-US"
)

var (
	Languages = []string{
		"en-US",
		"ru-RU",
		"af-FR",
		"ar-AR",
		"ben-BEN",
		"de-GER",
		"el-GR",
		"es-ES",
		"hl-HL",
		"ja-JA",
		"jv-JV",
		"ko-KO",
		"pt-BR",
		"zh-CN",
	}

	Shaders = []Shader{
		{
			Name:  "Disabled",
			Value: "disabled",
		}, {
			Name:  "2xScaleHQ",
			Value: "2xScaleHQ.glslp",
		}, {
			Name:  "4xScaleHQ",
			Value: "4xScaleHQ.glslp",
		}, {
			Name:  "CRT easymode",
			Value: "crt-easymode.glslp",
		}, {
			Name:  "CRT aperture",
			Value: "crt-aperture.glslp",
		}, {
			Name:  "CRT geom",
			Value: "crt-geom.glslp",
		}, {
			Name:  "CRT mattias",
			Value: "crt-mattias.glslp",
		},
	}

	FastForwardRatios = []string{
		"1.5", "2.0", "2.5", "3.0", "3.5", "4.0", "4.5", "5.0", "5.5", "6.0", "6.5", "7.0", "7.5", "8.0", "8.5", "9.0", "9.5", "10.0", "unlimited",
	}

	SlowMotionRatios = []string{
		"1.5", "2.0", "2.5", "3.0", "3.5", "4.0", "4.5", "5.0", "5.5", "6.0", "6.5", "7.0", "7.5", "8.0", "8.5", "9.0", "9.5", "10.0",
	}

	RewindGranularities = []string{
		"1", "3", "6", "12", "25", "50", "100",
	}

	PlatformIds = []string{
		"nes",
		"snes",
		"n64",
		"gb",
		"gba",
		"nds",
		"vb",
		"segaMS",
		"segaMD",
		"segaGG",
		"segaCD",
		"segaSaturn",
		"psx",
		"3do",
		"a5200",
		"atari7800",
		"atari2600",
		"jaguar",
		"lynx",
		"mame2003",
		"arcade",
		"pce",
		"pcfx",
		"ngp",
		"ws",
		"coleco",
	}

	Platforms = map[string]Platform{
		"nes": {
			Id:         "nes",
			Name:       "Nintendo Famicom / NES",
			Extensions: []string{"nes", "unf", "unif", "fds"},
		},
		"snes": {
			Id:         "snes",
			Name:       "Nintendo Super Famicom / SNES",
			Extensions: []string{"smc", "sfc", "swc", "fig", "bs", "st"},
		},
		"gb": {
			Id:         "gb",
			Name:       "Nintendo Game Boy (Color)",
			Extensions: []string{"gb", "gbc"},
		},
		"gba": {
			Id:         "gba",
			Name:       "Nintendo Game Boy Advance",
			Extensions: []string{"gba"},
		},
		"vb": {
			Id:         "vb",
			Name:       "Nintendo Virtual Boy",
			Extensions: []string{"vb", "vboy"},
		},
		"nds": {
			Id:         "nds",
			Name:       "Nintendo DS",
			Extensions: []string{"nds"},
		},
		"a5200": {
			Id:         "a5200",
			Name:       "Atari 5200",
			Extensions: []string{"a52"},
		},
		"mame2003": {
			Id:         "mame2003",
			Name:       "MAME 2003",
			Extensions: []string{},
		},
		"arcade": {
			Id:         "arcade",
			Name:       "Arcade",
			Extensions: []string{},
		},
		"psx": {
			Id:         "psx",
			Name:       "Sony PlayStation",
			Extensions: []string{},
		},
		"jaguar": {
			Id:         "jaguar",
			Name:       "Atari Jaguar",
			Extensions: []string{"j64", "jag"},
		},
		"lynx": {
			Id:         "lynx",
			Name:       "Atari Lynx",
			Extensions: []string{"lnx"},
		},
		"segaSaturn": {
			Id:         "segaSaturn",
			Name:       "Sega Saturn",
			Extensions: []string{},
		},
		"segaMS": {
			Id:         "segaMS",
			Name:       "Sega Master System",
			Extensions: []string{"sms"},
		},
		"segaMD": {
			Id:         "segaMD",
			Name:       "Sega Mega Drive / Genesis",
			Extensions: []string{"gen", "smd", "md"},
		},
		"segaGG": {
			Id:         "segaGG",
			Name:       "Sega Game Gear",
			Extensions: []string{"gg"},
		},
		"segaCD": {
			Id:         "segaCD",
			Name:       "Sega CD",
			Extensions: []string{},
		},
		"n64": {
			Id:         "n64",
			Name:       "Nintendo 64",
			Extensions: []string{"n64", "v64", "z64", "u1", "ndd"},
		},
		"3do": {
			Id:         "3do",
			Name:       "3DO",
			Extensions: []string{},
		},
		"atari7800": {
			Id:         "atari7800",
			Name:       "Atari 7800",
			Extensions: []string{"a78"},
		},
		"atari2600": {
			Id:         "atari2600",
			Name:       "Atari 2600",
			Extensions: []string{"a26"},
		},
		"pce": {
			Id:         "pce",
			Name:       "NEC TurboGrafx-16 / SuperGrafx / PC Engine",
			Extensions: []string{"pce"},
		},
		"pcfx": {
			Id:         "pcfx",
			Name:       "NEC PC-FX",
			Extensions: []string{},
		},
		"ngp": {
			Id:         "ngp",
			Name:       "SNK Neo Geo Pocket (Color)",
			Extensions: []string{"ngp", "ngc"},
		},
		"ws": {
			Id:         "ws",
			Name:       "Bandai WonderSwan (Color)",
			Extensions: []string{"ws", "wsc"},
		},
		"coleco": {
			Id:         "coleco",
			Name:       "ColecoVision",
			Extensions: []string{"col"},
		},
		"": {
			Id:         "",
			Name:       "Undefined",
			Extensions: []string{},
		},
	}

	Cores = map[string][]string{
		"nes":        {"fceumm", "nestopia"},
		"snes":       {"snes9x"},
		"gb":         {"gambatte", "mgba"},
		"gba":        {"mgba"},
		"vb":         {"beetle_vb"},
		"nds":        {"melonds", "desmume2015"},
		"a5200":      {"a5200"},
		"mame2003":   {"mame2003"},
		"arcade":     {"fbneo", "fbalpha2012_cps1", "fbalpha2012_cps2"},
		"psx":        {"pcsx_rearmed", "mednafen_psx_hw"},
		"jaguar":     {"virtualjaguar"},
		"lynx":       {"handy"},
		"segaSaturn": {"yabause"},
		"segaMS":     {"genesis_plus_gx", "picodrive"},
		"segaMD":     {"genesis_plus_gx"},
		"segaGG":     {"genesis_plus_gx"},
		"segaCD":     {"genesis_plus_gx"},
		"n64":        {"parallel_n64", "mupen64plus_next"},
		"3do":        {"opera"},
		"atari7800":  {"prosystem"},
		"atari2600":  {"stella2014"},
		"pce":        {"mednafen_pce"},
		"pcfx":       {"mednafen_pcfx"},
		"ngp":        {"mednafen_ngp"},
		"ws":         {"mednafen_wswan"},
		"coleco":     {"gearcoleco"},
	}

	Bioses = map[string][]Bios{
		"nes": {
			Bios{
				Name:        "disksys.rom",
				Url:         "/assets/bios/nes/disksys.rom",
				Hash:        "ca30b50f880eb660a320674ed365ef7a",
				Description: "Family Computer Disk System BIOS",
			},
		},
		"snes": {
			Bios{
				Name:        "BS-X.bin",
				Url:         "/assets/snes/BS-X.bin",
				Hash:        "fed4d8242cfbed61343d53d48432aced",
				Description: "BS-X - Sore wa Namae o Nusumareta Machi no Monogatari (Japan) (Rev 1)",
			},
			Bios{
				Name:        "STBIOS.bin",
				Url:         "/assets/snes/STBIOS.bin",
				Hash:        "d3a44ba7d42a74d3ac58cb9c14c6a5ca",
				Description: "Sufami Turbo (Japan)",
			},
		},
		"gb": {
			Bios{
				Name:        "gb.zip",
				Url:         "/assets/bios/gb/gb.zip",
				Description: "Game Boy (+Color) BIOS",
			},
		},
		"gba": {
			Bios{
				Name:        "gba.zip",
				Url:         "/assets/bios/gba/gba.zip",
				Hash:        "a860e8c0b6d573d191e4ec7db1b1e4f6",
				Description: "Game Boy Advance (+Original, +Color, +Super) BIOS",
			},
		},
		"vb": {},
		"nds": {
			Bios{
				Name:        "ds.zip",
				Url:         "/assets/bios/nds/DS.zip",
				Description: "DS BIOS",
			},
		},
		"a5200":    {},
		"mame2003": {},
		"arcade":   {},
		"psx": {
			Bios{
				Name:        "psp.zip",
				Url:         "/assets/bios/psx/psp.zip",
				Description: "Extracted from a PSP",
			},
			Bios{
				Name:        "psx.zip",
				Url:         "/assets/bios/psx/psx.zip",
				Description: "PS1 BIOS (US, EU, JP)",
			},
		},
		"jaguar": {},
		"lynx": {
			Bios{
				Name:        "lynxboot.img",
				Url:         "/assets/bios/lynx/lynxboot.img",
				Hash:        "fcd403db69f54290b51035d82f835e7b",
				Description: "Lynx Boot Image",
			},
		},
		"segaSaturn": {
			Bios{
				Name:        "saturn_bios.bin",
				Url:         "/assets/bios/segaSaturn/saturn_bios.bin",
				Hash:        "af5828fdff51384f99b3c4926be27762",
				Description: "Saturn BIOS",
			},
		},
		"segaMS": {
			Bios{
				Name:        "segaMS.zip",
				Url:         "/assets/bios/segaMS/segaMS.zip",
				Description: "MasterSystem BIOS (US, EU, JP)",
			},
		},
		"segaMD": {
			Bios{
				Name:        "bios_MD.bin",
				Url:         "/assets/bios/segaMD/bios_MD.bin",
				Hash:        "45e298905a08f9cfb38fd504cd6dbc84",
				Description: "MegaDrive TMSS startup ROM",
			},
		},
		"segaGG": {
			Bios{
				Name:        "bios.gg",
				Url:         "/assets/bios/segaGG/bios.gg",
				Hash:        "672e104c3be3a238301aceffc3b23fd6",
				Description: "GameGear BIOS (bootrom)",
			},
		},
		"segaCD": {
			Bios{
				Name:        "segaCD.zip",
				Url:         "/assets/bios/segaCD/segaCD.zip",
				Description: "SegaCD BIOS (US, EU, JP)",
			},
		},
		"n64": {},
		"3do": {
			Bios{
				Name:        "3do.zip",
				Url:         "/assets/bios/3do/3do.zip",
				Description: "3DO BIOS Collection",
			},
		},
		"atari7800": {
			Bios{
				Name:        "7800_BIOS_U.rom",
				Url:         "/assets/bios/atari7800/7800_BIOS_U.rom",
				Hash:        "0763f1ffb006ddbe32e52d497ee848ae",
				Description: "7800 BIOS",
			},
		},
		"atari2600": {},
		"pce": {
			Bios{
				Name:        "pce.zip",
				Url:         "/assets/bios/pce/pce.zip",
				Description: "PC Engine BIOS Collection",
			},
		},
		"pcfx": {
			Bios{
				Name:        "pcfx.zip",
				Url:         "/assets/bios/pcfx/pcfx.zip",
				Description: "PC-FX BIOS",
			},
		},
		"ngp": {},
		"ws":  {},
		"coleco": {
			Bios{
				Name:        "colecovision.rom",
				Url:         "/assets/bios/coleco/colecovision.rom",
				Hash:        "2c66f5911e5b42b8ebe113403548eee7",
				Description: "ColecoVision BIOS",
			},
		},
	}

	//Configuration reference: https://retropie.org.uk/docs/RetroArch-Configuration

	DefaultControlsNes = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,
			B:      KeyX,
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsGb = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,
			B:      KeyX,
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSnes = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,
			B:      KeyX,
			X:      KeyA,
			Y:      KeyS,
			L:      KeyQ,
			R:      KeyW,
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsN64 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:           KeyZ,
			B:           KeyX,
			Start:       KeyEnter,
			Up:          KeyArrowUp,    //D-Pad Up
			Down:        KeyArrowDown,  //D-Pad Down
			Left:        KeyArrowLeft,  //D-Pad Left
			Right:       KeyArrowRight, //D-Pad Right
			L:           KeyQ,
			R:           KeyW,
			LStickUp:    KeyT, //Stick up
			LStickDown:  KeyG, //Stick Down
			LStickLeft:  KeyF, //Stick Left
			LStickRight: KeyH, //Stick Right
			RStickUp:    KeyI, //C-Pad Up
			RStickDown:  KeyK, //C-Pad Down
			RStickLeft:  KeyJ, //C-Pad Left
			RStickRight: KeyL, //C-Pad Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsGba = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,
			B:      KeyX,
			L:      KeyA,
			R:      KeyS,
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsNds = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,
			B:      KeyX,
			X:      KeyA,
			Y:      KeyS,
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
			L:      KeyQ,
			R:      KeyW,
			L3:     "", //Microphone
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsVb = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:           KeyZ,
			B:           KeyX,
			L:           KeyQ,
			R:           KeyW,
			Select:      KeyV,
			Start:       KeyEnter,
			Up:          KeyArrowUp,    //Left D-Pad Up
			Down:        KeyArrowDown,  //Left D-Pad Down
			Left:        KeyArrowLeft,  //Left D-Pad Left
			Right:       KeyArrowRight, //Left D-Pad Right
			LStickUp:    KeyT,          //Right D-Pad Up
			LStickDown:  KeyG,          //Right D-Pad Down
			LStickLeft:  KeyF,          //Right D-Pad Left
			LStickRight: KeyH,          //Right D-Pad Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaMS = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:     KeyZ, //BUTTON 1 / START
			A:     KeyX, //BUTTON 2
			Up:    KeyArrowUp,
			Down:  KeyArrowDown,
			Left:  KeyArrowLeft,
			Right: KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaMD = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:      KeyZ, //A
			B:      KeyX, //B
			A:      KeyC, //C
			L:      KeyA, //X
			X:      KeyS, //Y
			R:      KeyD, //Z
			Select: KeyV, //Mode
			Start:  KeyEnter,
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaGG = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:     KeyZ, //BUTTON 1
			A:     KeyX, //BUTTON 2
			Start: KeyEnter,
			Up:    KeyArrowUp,
			Down:  KeyArrowDown,
			Left:  KeyArrowLeft,
			Right: KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaSaturn = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:     KeyZ, //A
			B:     KeyX, //B
			A:     KeyC, //C
			X:     KeyA, //X
			L:     KeyS, //Y
			R:     KeyD, //Z
			L2:    KeyQ, //L
			R2:    KeyW, //R
			Start: KeyEnter,
			Up:    KeyArrowUp,
			Down:  KeyArrowDown,
			Left:  KeyArrowLeft,
			Right: KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControls3do = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:      KeyZ,     //A
			B:      KeyX,     //B
			A:      KeyC,     //C
			L:      KeyA,     //L
			R:      KeyS,     //R
			Select: KeyV,     //X
			Start:  KeyEnter, //P
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsAtari2600 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:      KeyX, //Fire
			Select: KeyV, //Select
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsAtari7800 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:      KeyX, //Fire 1
			A:      KeyZ, //Fire 2
			Select: KeyV, //Select
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsLynx = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:     KeyZ,
			B:     KeyX,
			L:     KeyA, //Option 1
			R:     KeyS, //Option 2
			Start: KeyEnter,
			Up:    KeyArrowUp,
			Down:  KeyArrowDown,
			Left:  KeyArrowLeft,
			Right: KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsJaguar = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ,     //A
			B:      KeyX,     //B
			Y:      KeyC,     //C
			Select: KeyV,     //Pause
			Start:  KeyEnter, //Option
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsPCEngine = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyX,     //I
			B:      KeyZ,     //II
			Select: KeyV,     //Select
			Start:  KeyEnter, //Run
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsNeoGeoPocket = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:     KeyX,     //B
			B:     KeyZ,     //A
			Start: KeyEnter, //Option
			Up:    KeyArrowUp,
			Down:  KeyArrowDown,
			Left:  KeyArrowLeft,
			Right: KeyArrowRight,
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsWonderSwan = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ, //A
			B:      KeyX, //B
			Select: KeyV,
			Start:  KeyEnter,
			Up:     KeyArrowUp,    //X Cursor Up
			Down:   KeyArrowDown,  //X Cursor Down
			Left:   KeyArrowLeft,  //X Cursor Left
			Right:  KeyArrowRight, //X Cursor Right
			R2:     KeyT,          //Y Cursor Up
			L2:     KeyG,          //Y Cursor Down
			L:      KeyF,          //Y Cursor Left
			R:      KeyH,          //Y Cursor Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsColecoVision = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyZ, //Left Button
			B:      KeyX, //Right Button
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
			X:      KeyQ, //1
			Y:      KeyW, //2
			R:      KeyE, //3
			L:      KeyR, //4
			R2:     KeyT, //5
			L2:     KeyY, //6
			R3:     KeyU, //7
			L3:     KeyI, //8
			Select: KeyO, //*
			Start:  KeyP, //#
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsPCFX = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      KeyC, //I
			B:      KeyX, //II
			X:      KeyZ, //III
			Y:      KeyD, //IV
			L:      KeyS, //V
			R:      KeyA, //VI
			Select: KeyV,
			Start:  KeyEnter, //Run
			Up:     KeyArrowUp,
			Down:   KeyArrowDown,
			Left:   KeyArrowLeft,
			Right:  KeyArrowRight,
			L2:     KeyQ, //Mode 1
			R2:     KeyW, //Mode 2
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsOther = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:               KeyX,
			Y:               KeyS,
			Select:          KeyV,
			Start:           KeyEnter,
			Up:              KeyArrowUp,
			Down:            KeyArrowDown,
			Left:            KeyArrowLeft,
			Right:           KeyArrowRight,
			A:               KeyZ,
			X:               KeyA,
			L:               KeyQ,
			R:               KeyW,
			L2:              KeyE,
			R2:              KeyR,
			L3:              "",
			R3:              "",
			LStickUp:        KeyT,
			LStickDown:      KeyG,
			LStickLeft:      KeyF,
			LStickRight:     KeyH,
			RStickUp:        KeyI,
			RStickDown:      KeyK,
			RStickLeft:      KeyJ,
			RStickRight:     KeyL,
			QuickSaveState:  "",
			QuickLoadState:  "",
			ChangeStateSlot: "",
		},
		Gamepad: EmulatorControlsMapping{
			B:               GamepadControlB,
			Y:               GamepadControlY,
			Select:          GamepadControlSelect,
			Start:           GamepadControlStart,
			Up:              GamepadControlUp,
			Down:            GamepadControlDown,
			Left:            GamepadControlLeft,
			Right:           GamepadControlRight,
			A:               GamepadControlA,
			X:               GamepadControlX,
			L:               GamepadControlL,
			R:               GamepadControlR,
			L2:              GamepadControlL2,
			R2:              GamepadControlR2,
			L3:              GamepadControlL3,
			R3:              GamepadControlR3,
			LStickUp:        GamepadControlLStickUp,
			LStickDown:      GamepadControlLStickDown,
			LStickLeft:      GamepadControlLStickLeft,
			LStickRight:     GamepadControlLStickRight,
			RStickUp:        GamepadControlRStickUp,
			RStickDown:      GamepadControlRStickDown,
			RStickLeft:      GamepadControlRStickLeft,
			RStickRight:     GamepadControlRStickRight,
			QuickSaveState:  "",
			QuickLoadState:  "",
			ChangeStateSlot: "",
		},
	}

	DefaultButtons = EmulatorButtons{
		PlayPause:    true,
		Restart:      true,
		Mute:         true,
		FullScreen:   true,
		Volume:       true,
		Screenshot:   true,
		SaveState:    false,
		LoadState:    false,
		QuickSave:    false, //only to save in browser
		QuickLoad:    false, //only to save in browser
		ScreenRecord: false, //not implemented in EJS
		Settings:     false, //managed in emulation settings
		Gamepad:      false, //managed in emulation settings
		Cheat:        false, //not implemented in playtime
		SaveSavFiles: false,
		LoadSavFiles: false,
		CacheManager: false, //managed in emulation settings
	}
)

func DefaultEmulatorSettings(systemType string) EmulatorSettings {
	settings := EmulatorSettings{}
	controls := EmulatorControls{}

	switch systemType {

	case "nes":
		settings = EmulatorSettings{
			Core: Cores["nes"][0],
			Bios: "",
		}
		controls = DefaultControlsNes

	case "snes":
		settings = EmulatorSettings{
			Core: Cores["snes"][0],
			Bios: "",
		}
		controls = DefaultControlsSnes

	case "gb":
		settings = EmulatorSettings{
			Core: Cores["gb"][0],
			Bios: Bioses["gb"][0].Name,
		}
		controls = DefaultControlsGb

	case "gba":
		settings = EmulatorSettings{
			Core: Cores["gba"][0],
			Bios: Bioses["gba"][0].Name,
		}
		controls = DefaultControlsGba

	case "vb":
		settings = EmulatorSettings{
			Core: Cores["vb"][0],
			Bios: "",
		}
		controls = DefaultControlsVb

	case "nds":
		settings = EmulatorSettings{
			Core: Cores["nds"][0],
			Bios: Bioses["nds"][0].Name,
		}
		controls = DefaultControlsNds

	case "a5200":
		settings = EmulatorSettings{
			Core: Cores["a5200"][0],
			Bios: "",
		}
		controls = DefaultControlsOther

	case "mame2003":
		settings = EmulatorSettings{
			Core: Cores["mame2003"][0],
			Bios: "",
		}
		controls = DefaultControlsOther

	case "arcade":
		settings = EmulatorSettings{
			Core: Cores["arcade"][0],
			Bios: "",
		}
		controls = DefaultControlsOther

	case "psx":
		settings = EmulatorSettings{
			Core: Cores["psx"][0],
			Bios: Bioses["psx"][0].Name,
		}
		controls = DefaultControlsOther

	case "jaguar":
		settings = EmulatorSettings{
			Core: Cores["jaguar"][0],
			Bios: "",
		}
		controls = DefaultControlsJaguar

	case "lynx":
		settings = EmulatorSettings{
			Core: Cores["lynx"][0],
			Bios: Bioses["lynx"][0].Name,
		}
		controls = DefaultControlsLynx

	case "segaSaturn":
		settings = EmulatorSettings{
			Core: Cores["segaSaturn"][0],
			Bios: Bioses["segaSaturn"][0].Name,
		}
		controls = DefaultControlsSegaSaturn

	case "segaMS":
		settings = EmulatorSettings{
			Core: Cores["segaMS"][0],
			Bios: Bioses["segaMS"][0].Name,
		}
		controls = DefaultControlsSegaMS

	case "segaMD":
		settings = EmulatorSettings{
			Core: Cores["segaMD"][0],
			Bios: "",
		}
		controls = DefaultControlsSegaMD

	case "segaGG":
		settings = EmulatorSettings{
			Core: Cores["segaGG"][0],
			Bios: Bioses["segaGG"][0].Name,
		}
		controls = DefaultControlsSegaGG

	case "segaCD":
		settings = EmulatorSettings{
			Core: Cores["segaCD"][0],
			Bios: Bioses["segaCD"][0].Name,
		}
		controls = DefaultControlsSegaMD

	case "n64":
		settings = EmulatorSettings{
			Core: Cores["n64"][0],
			Bios: "",
		}
		controls = DefaultControlsN64

	case "3do":
		settings = EmulatorSettings{
			Core: Cores["3do"][0],
			Bios: Bioses["3do"][0].Name,
		}
		controls = DefaultControls3do

	case "atari7800":
		settings = EmulatorSettings{
			Core: Cores["atari7800"][0],
			Bios: Bioses["atari7800"][0].Name,
		}
		controls = DefaultControlsAtari7800

	case "atari2600":
		settings = EmulatorSettings{
			Core: Cores["atari2600"][0],
			Bios: "",
		}
		controls = DefaultControlsAtari2600

	case "pce":
		settings = EmulatorSettings{
			Core: Cores["pce"][0],
			Bios: "",
		}
		controls = DefaultControlsPCEngine

	case "pcfx":
		settings = EmulatorSettings{
			Core: Cores["pcfx"][0],
			Bios: Bioses["pcfx"][0].Name,
		}
		controls = DefaultControlsPCFX

	case "ngp":
		settings = EmulatorSettings{
			Core: Cores["ngp"][0],
			Bios: "",
		}
		controls = DefaultControlsNeoGeoPocket

	case "ws":
		settings = EmulatorSettings{
			Core: Cores["ws"][0],
			Bios: "",
		}
		controls = DefaultControlsWonderSwan

	case "coleco":
		settings = EmulatorSettings{
			Core: Cores["coleco"][0],
			Bios: Bioses["coleco"][0].Name,
		}
		controls = DefaultControlsColecoVision
	}

	settings.ColorScheme = DefaultColorScheme
	settings.ColorBackground = DefaultColorBackground
	settings.CacheLimit = DefaultCacheLimit
	settings.Volume = DefaultVolume
	settings.FastForwardRatio = DefaultFastForwardRatio
	settings.SlowMotionRatio = DefaultSlowMotionRatio
	settings.RewindGranularity = DefaultRewindGranularity
	settings.Shader = Shaders[0].Value
	settings.Buttons = DefaultButtons
	settings.Controls = [4]EmulatorControls{controls}

	return settings
}

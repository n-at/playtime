package storage

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

	Shaders = []string{
		"disabled",
		"2xScaleHQ.glslp",
		"4xScaleHQ.glslp",
		"crt-easymode.glslp",
		"crt-aperture.glslp",
		"crt-geom.glslp",
		"crt-mattias.glslp",
	}

	Systems = []string{
		"nes",
		"snes",
		"gb",
		"gba",
		"vb",
		"nds",
		"a5200",
		"mame2003",
		"arcade",
		"psx",
		"jaguar",
		"lynx",
		"segaSaturn",
		"segaMS",
		"segaMD",
		"segaGG",
		"segaCD",
		"sega32x",
		"n64",
		"3do",
		"atari7800",
		"atari2600",
	}

	Platforms = map[string]Platform{
		"nes": {
			Name: "Nintendo Famicom / NES",
		},
		"snes": {
			Name: "Nintendo Super Famicom / SNES",
		},
		"gb": {
			Name: "Nintendo Game Boy",
		},
		"gba": {
			Name: "Nintendo Game Boy Advance",
		},
		"vb": {
			Name: "Nintendo Virtual Boy",
		},
		"nds": {
			"Nintendo DS",
		},
		"a5200": {
			Name: "Atari 5200",
		},
		"mame2003": {
			Name: "MAME 2003",
		},
		"arcade": {
			Name: "Arcade",
		},
		"psx": {
			Name: "Sony PlayStation",
		},
		"jaguar": {
			Name: "Atari Jaguar",
		},
		"lynx": {
			Name: "Atari Lynx",
		},
		"segaSaturn": {
			Name: "Sega Saturn",
		},
		"segaMS": {
			Name: "Sega Master System",
		},
		"segaMD": {
			Name: "Sega Mega Drive / Genesis",
		},
		"segaGG": {
			Name: "Sega Game Gear",
		},
		"segaCD": {
			Name: "Sega CD",
		},
		"sega32x": {
			Name: "Sega 32X",
		},
		"n64": {
			Name: "Nintendo 64",
		},
		"3do": {
			Name: "3DO",
		},
		"atari7800": {
			Name: "Atari 7800",
		},
		"atari2600": {
			Name: "Atari 2600",
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
		"segaMD":     {"genesis_plus_gx", "picodrive"},
		"segaGG":     {"genesis_plus_gx"},
		"segaCD":     {"genesis_plus_gx"},
		"sega32x":    {"picodrive"},
		"n64":        {"mupen64plus_next"},
		"3do":        {"opera"},
		"atari7800":  {"prosystem"},
		"atari2600":  {"stella2014"},
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
				Description: "BS-X - Sore wa Namae o Nusumareta Machi no\nMonogatari (Japan) (Rev 1)",
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
				Name:        "gb_bios.bin",
				Url:         "/assets/bios/gb/gb_bios.bin",
				Hash:        "32fbbd84168d3482956eb3c5051637f5",
				Description: "Game Boy BIOS",
			},
			Bios{
				Name:        "gbc_bios.bin",
				Url:         "/assets/bios/gb/gbc_bios.bin",
				Hash:        "dbfce9db9deaa2567f6a84fde55f9680",
				Description: "Game Boy Color BIOS",
			},
		},
		"gba": {
			Bios{
				Name:        "gba_bios.bin",
				Url:         "/assets/bios/gba/gba_bios.bin",
				Hash:        "a860e8c0b6d573d191e4ec7db1b1e4f6",
				Description: "Game Boy Advance BIOS",
			},
			Bios{
				Name:        "gb_bios.bin",
				Url:         "/assets/bios/gba/gb_bios.bin",
				Hash:        "32fbbd84168d3482956eb3c5051637f5",
				Description: "Game Boy BIOS",
			},
			Bios{
				Name:        "gbc_bios.bin",
				Url:         "/assets/bios/gba/gbc_bios.bin",
				Hash:        "dbfce9db9deaa2567f6a84fde55f9680",
				Description: "Game Boy Color BIOS",
			},
			Bios{
				Name:        "sgb_bios.bin",
				Url:         "/assets/bios/gba/sgb_bios.bin",
				Hash:        "d574d4f9c12f305074798f54c091a8b4",
				Description: "Super Game Boy BIOS",
			},
		},
		"vb": {},
		"nds": {
			Bios{
				Name:        "bios7.bin",
				Url:         "/assets/bios/nds/bios7.bin",
				Hash:        "df692a80a5b1bc90728bc3dfc76cd948",
				Description: "NDS ARM7 BIOS",
			},
			Bios{
				Name:        "bios9.bin",
				Url:         "/assets/bios/nds/bios9.bin",
				Hash:        "a392174eb3e572fed6447e956bde4b25",
				Description: "NDS ARM9 BIOS",
			},
			Bios{
				Name:        "firmware.bin",
				Url:         "/assets/bios/nds/firmware.bin",
				Hash:        "145eaef5bd3037cbc247c213bb3da1b3",
				Description: "NDS Firmware",
			},
			Bios{
				Name:        "dsi_bios7.bin",
				Url:         "/assets/bios/nds/dsi_bios7.bin",
				Hash:        "",
				Description: "DSi ARM7 BIOS",
			},
			Bios{
				Name:        "dsi_bios9.bin",
				Url:         "/assets/bios/nds/dsi_bios9.bin",
				Hash:        "",
				Description: "DSi ARM9 BIOS",
			},
			Bios{
				Name:        "dsi_firmware.bin",
				Url:         "/assets/bios/nds/dsi_firmware.bin",
				Hash:        "",
				Description: "DSi Firmware",
			},
			Bios{
				Name:        "dsi_nand.bin",
				Url:         "/assets/bios/nds/dsi_nand.bin",
				Hash:        "",
				Description: "DSi NAND",
			},
			Bios{
				Name:        "dsi_sd_card.bin",
				Url:         "/assets/bios/nds/dsi_sd_card.bin",
				Hash:        "",
				Description: "DSi SD Card",
			},
		},
		"a5200":    {},
		"mame2003": {},
		"arcade":   {},
		"psx": {
			Bios{
				Name:        "PSXONPSP660.bin",
				Url:         "/assets/bios/psx/PSXONPSP660.bin",
				Hash:        "c53ca5908936d412331790f4426c6c33",
				Description: "Extracted from a PSP",
			},
			Bios{
				Name:        "scph5501.bin",
				Url:         "/assets/bios/psx/scph5501.bin",
				Hash:        "490f666e1afb15b7362b406ed1cea246",
				Description: "PS1 US BIOS",
			},
			Bios{
				Name:        "scph5500.bin",
				Url:         "/assets/bios/psx/scph5500.bin",
				Hash:        "8dd7d5296a650fac7319bce665a6a53c",
				Description: "PS1 JP BIOS",
			},
			Bios{
				Name:        "scph5502.bin",
				Url:         "/assets/bios/psx/scph5502.bin",
				Hash:        "32736f17079d0b2b7024407c39bd3050",
				Description: "PS1 EU BIOS",
			},
			Bios{
				Name:        "scph101.bin",
				Url:         "/assets/bios/psx/scph101.bin",
				Hash:        "6E3735FF4C7DC899EE98981385F6F3D0",
				Description: "Version 4.4 03/24/00 A",
			},
			Bios{
				Name:        "scph7001.bin",
				Url:         "/assets/bios/psx/scph7001.bin",
				Hash:        "1e68c231d0896b7eadcad1d7d8e76129",
				Description: "Version 4.1 12/16/97 A",
			},
			Bios{
				Name:        "scph1001.bin",
				Url:         "/assets/bios/psx/scph1001.bin",
				Hash:        "924e392ed05558ffdb115408c263dccf",
				Description: "Version 2.0 05/07/95 A",
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
				Name:        "bios_U.sms",
				Url:         "/assets/bios/segaMS/bios_U.sms",
				Hash:        "840481177270d5642a14ca71ee72844c",
				Description: "MasterSystem US BIOS",
			},
			Bios{
				Name:        "bios_E.sms",
				Url:         "/assets/bios/segaMS/bios_E.sms",
				Hash:        "840481177270d5642a14ca71ee72844c",
				Description: "MasterSystem EU BIOS",
			},
			Bios{
				Name:        "bios_J.sms",
				Url:         "/assets/bios/segaMS/bios_J.sms",
				Hash:        "24a519c53f67b00640d0048ef7089105",
				Description: "MasterSystem JP BIOS",
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
				Name:        "bios_CD_U.bin",
				Url:         "/assets/bios/segaCD/bios_CD_U.bin",
				Hash:        "2efd74e3232ff260e371b99f84024f7f",
				Description: "SegaCD US BIOS",
			},
			Bios{
				Name:        "bios_CD_E.bin",
				Url:         "/assets/bios/segaCD/bios_CD_E.bin",
				Hash:        "e66fa1dc5820d254611fdcdba0662372",
				Description: "MegaCD EU BIOS",
			},
			Bios{
				Name:        "bios_CD_J.bin",
				Url:         "/assets/bios/segaCD/bios_CD_J.bin",
				Hash:        "278a9397d192149e84e820ac621a8edd",
				Description: "MegaCD JP BIOS",
			},
		},
		"sega32x": {},
		"n64":     {},
		"3do": {
			Bios{
				Name:        "panafz1.bin",
				Url:         "/assets/bios/3do/panafz1.bin",
				Hash:        "f47264dd47fe30f73ab3c010015c155b",
				Description: "Panasonic FZ-1",
			},
			Bios{
				Name:        "panafz10.bin",
				Url:         "/assets/bios/3do/panafz10.bin",
				Hash:        "51f2f43ae2f3508a14d9f56597e2d3ce",
				Description: "Panasonic FZ-10",
			},
			Bios{
				Name:        "panafz10-norsa.bin",
				Url:         "/assets/bios/3do/panafz10-norsa.bin",
				Hash:        "1477bda80dc33731a65468c1f5bcbee9",
				Description: "Panasonic FZ-10 [RSA Patch]",
			},
			Bios{
				Name:        "panafz10e-anvil.bin",
				Url:         "/assets/bios/3do/panafz10e-anvil.bin",
				Hash:        "a48e6746bd7edec0f40cff078f0bb19f",
				Description: "Panasonic FZ-10-E [Anvil]",
			},
			Bios{
				Name:        "panafz10e-anvil-norsa.bin",
				Url:         "/assets/bios/3do/panafz10e-anvil-norsa.bin",
				Hash:        "cf11bbb5a16d7af9875cca9de9a15e09",
				Description: "Panasonic FZ-10-E [Anvil RSA Patch]",
			},
			Bios{
				Name:        "panafz1j.bin",
				Url:         "/assets/bios/3do/panafz1j.bin",
				Hash:        "a496cfdded3da562759be3561317b605",
				Description: "Panasonic FZ-1J",
			},
			Bios{
				Name:        "panafz1j-norsa.bin",
				Url:         "/assets/bios/3do/panafz1j-norsa.bin",
				Hash:        "f6c71de7470d16abe4f71b1444883dc8",
				Description: "Panasonic FZ-1J [RSA Patch]",
			},
			Bios{
				Name:        "goldstar.bin",
				Url:         "/assets/bios/3do/goldstar.bin",
				Hash:        "8639fd5e549bd6238cfee79e3e749114",
				Description: "Goldstar GDO-101M",
			},
			Bios{
				Name:        "sanyotry.bin",
				Url:         "/assets/bios/3do/sanyotry.bin",
				Hash:        "35fa1a1ebaaeea286dc5cd15487c13ea",
				Description: "Sanyo IMP-21J TRY",
			},
			Bios{
				Name:        "3do_arcade_saot.bin",
				Url:         "/assets/bios/3do/3do_arcade_saot.bin",
				Hash:        "8970fc987ab89a7f64da9f8a8c4333ff",
				Description: "Shootout At Old Tucson",
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
	}

	DefaultControlsNes = EmulatorControlsMapping{
		A:      "z",
		B:      "x",
		Select: "v",
		Start:  "enter",
		Up:     "arrowup",
		Down:   "arrowdown",
		Left:   "arrowleft",
		Right:  "arrowright",
	}

	DefaultControlsSnes = EmulatorControlsMapping{
		A:      "z",
		B:      "x",
		X:      "a",
		Y:      "s",
		Select: "v",
		Start:  "enter",
		Up:     "arrowup",
		Down:   "arrowdown",
		Left:   "arrowleft",
		Right:  "arrowright",
		L:      "q",
		R:      "w",
	}

	DefaultControlsNds = EmulatorControlsMapping{
		A:      "z",
		B:      "x",
		X:      "a",
		Y:      "s",
		Select: "v",
		Start:  "enter",
		Up:     "arrowup",
		Down:   "arrowdown",
		Left:   "arrowleft",
		Right:  "arrowright",
		L:      "q",
		R:      "w",
		L3:     "",
	}

	DefaultControlsSegaMD = EmulatorControlsMapping{
		Y:     "z", //A
		B:     "x", //B
		A:     "c", //C
		L:     "a", //X
		X:     "s", //Y
		R:     "d", //Z
		Start: "enter",
		Up:    "arrowup",
		Down:  "arrowdown",
		Left:  "arrowleft",
		Right: "arrowright",
	}

	DefaultControlsOther = EmulatorControlsMapping{
		B:               "x",
		Y:               "s",
		Select:          "v",
		Start:           "enter",
		Up:              "arrowup",
		Down:            "arrowdown",
		Left:            "arrowleft",
		Right:           "arowright",
		A:               "z",
		X:               "a",
		L:               "q",
		R:               "w",
		L2:              "e",
		R2:              "r",
		L3:              "",
		R3:              "",
		LStickUp:        "t",
		LStickDown:      "g",
		LStickLeft:      "f",
		LStickRight:     "h",
		RStickUp:        "i",
		RStickDown:      "k",
		RStickLeft:      "j",
		RStickRight:     "l",
		QuickSaveState:  "",
		QuickLoadState:  "",
		ChangeStateSlot: "",
	}

	DefaultButtons = EmulatorButtons{
		PlayPause:    true,
		Restart:      true,
		Mute:         true,
		Settings:     true,
		FullScreen:   true,
		SaveState:    true,
		LoadState:    true,
		ScreenRecord: true,
		Gamepad:      true,
		Cheat:        false,
		Volume:       true,
		SaveSavFiles: true,
		LoadSavFiles: true,
		QuickSave:    true,
		QuickLoad:    true,
		Screenshot:   true,
		CacheManager: true,
	}
)

const (
	DefaultColorScheme = "#1AAFFF"
	DefaultCacheLimit  = 1073741824
	DefaultVolume      = 0.5
)

func DefaultEmulatorSettings(systemType string) EmulatorSettings {
	switch systemType {

	case "nes":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["nes"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNes,
				},
			},
		}

	case "snes":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["snes"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsSnes,
				},
			},
		}

	case "gb":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["gb"][0],
			Bios:                   Bioses["gb"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNes,
				},
			},
		}

	case "gba":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["gba"][0],
			Bios:                   Bioses["gba"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNes,
				},
			},
		}

	case "vb":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["vb"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "nds":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["nds"][0],
			Bios:                   Bioses["nds"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNds,
				},
			},
		}

	case "a5200":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["a5200"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "mame2003":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["mame2003"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "arcade":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["arcade"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "psx":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["psx"][0],
			Bios:                   Bioses["psx"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "jaguar":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["jaguar"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "lynx":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["lynx"][0],
			Bios:                   Bioses["lynx"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "segaSaturn":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["segaSaturn"][0],
			Bios:                   Bioses["segaSaturn"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsSegaMD,
				},
			},
		}

	case "segaMS":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["segaMS"][0],
			Bios:                   Bioses["segaMS"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNes,
				},
			},
		}

	case "segaMD":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["segaMD"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsSegaMD,
				},
			},
		}

	case "segaGG":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["segaGG"][0],
			Bios:                   Bioses["segaGG"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsNes,
				},
			},
		}

	case "segaCD":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["segaCD"][0],
			Bios:                   Bioses["segaCD"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsSegaMD,
				},
			},
		}

	case "sega32x":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["sega32x"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsSegaMD,
				},
			},
		}

	case "n64":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["n64"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "3do":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["3do"][0],
			Bios:                   Bioses["3do"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "atari7800":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["atari7800"][0],
			Bios:                   Bioses["atari7800"][0].Name,
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}

	case "atari2600":
		return EmulatorSettings{
			OldCores:               false,
			Core:                   Cores["atari2600"][0],
			Bios:                   "",
			ColorScheme:            DefaultColorScheme,
			CacheLimit:             DefaultCacheLimit,
			Volume:                 DefaultVolume,
			Shader:                 Shaders[0],
			FPS:                    false,
			VirtualGamepadLeftHand: false,
			Buttons:                DefaultButtons,
			Controls: [4]EmulatorControls{
				{
					Keyboard: DefaultControlsOther,
				},
			},
		}
	}

	return EmulatorSettings{}
}

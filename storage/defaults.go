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
		},
		"a5200":    {},
		"mame2003": {},
		"arcade":   {},
		"psx": {
			Bios{
				Name:        "PSXONPSP660_as_scph5501",
				Url:         "/assets/bios/psx/PSXONPSP660/scph5501.bin",
				Hash:        "c53ca5908936d412331790f4426c6c33",
				Description: "Extracted from a PSP as US",
			},
			Bios{
				Name:        "PSXONPSP660_as_scph5500",
				Url:         "/assets/bios/psx/PSXONPSP660/scph5500.bin",
				Hash:        "c53ca5908936d412331790f4426c6c33",
				Description: "Extracted from a PSP as JP",
			},
			Bios{
				Name:        "PSXONPSP660_as_scph5502",
				Url:         "/assets/bios/psx/PSXONPSP660/scph5502.bin",
				Hash:        "c53ca5908936d412331790f4426c6c33",
				Description: "Extracted from a PSP as EU",
			},
			Bios{
				Name:        "PSXONPSP660.bin",
				Url:         "/assets/bios/psx/PSXONPSP660.bin",
				Hash:        "c53ca5908936d412331790f4426c6c33",
				Description: "Extracted from a PSP",
			},
			Bios{
				Name:        "scph5500.bin",
				Url:         "/assets/bios/psx/scph5500.bin",
				Hash:        "8dd7d5296a650fac7319bce665a6a53c",
				Description: "PS1 JP BIOS",
			},
			Bios{
				Name:        "scph5501.bin",
				Url:         "/assets/bios/psx/scph5501.bin",
				Hash:        "490f666e1afb15b7362b406ed1cea246",
				Description: "PS1 US BIOS",
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
		"n64": {},
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
		"pce": {
			Bios{
				Name:        "syscard3.pce",
				Url:         "/assets/bios/pce/syscard3.pce",
				Hash:        "79f5ff55dd10187c7fd7b8daab0b3ffbd1f56a2c",
				Description: "Super CD-ROM2 System V3.xx",
			},
			Bios{
				Name:        "syscard2.pce",
				Url:         "/assets/bios/pce/syscard2.pce",
				Hash:        "88da02e2503f7c32810f5d93a34849d470742b6d",
				Description: "CD-ROM System V2.xx",
			},
			Bios{
				Name:        "syscard1.pce",
				Url:         "/assets/bios/pce/syscard2.pce",
				Hash:        "a39a66da7de6ba94ab84d04eef7afeec7d4ee66a",
				Description: "CD-ROM System V1.xx",
			},
			Bios{
				Name:        "gexpress.pce",
				Url:         "/assets/bios/pce/gexpress.pce",
				Hash:        "014881a959e045e00f4db8f52955200865d40280",
				Description: "Game Express CD Card",
			},
		},
		"pcfx": {
			Bios{
				Name:        "pcfx.rom",
				Url:         "/assets/bios/pcfx/pcfx.rom",
				Hash:        "08e36edbea28a017f79f8d4f7ff9b6d7",
				Description: "PC-FX BIOS v1.00",
			},
			Bios{
				Name:        "pcfxbios.bin",
				Url:         "/assets/bios/pcfx/pcfxbios.bin",
				Hash:        "08e36edbea28a017f79f8d4f7ff9b6d7",
				Description: "pcfxbios.bin",
			},
			Bios{
				Name:        "pcfxv101.bin",
				Url:         "/assets/bios/pcfx/pcfxv101.bin",
				Hash:        "e2fb7c7220e3a7838c2dd7e401a7f3d8",
				Description: "pcfxv101.bin",
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
			A:      "z",
			B:      "x",
			Select: "v",
			Start:  "enter",
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsGb = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z",
			B:      "x",
			Select: "v",
			Start:  "enter",
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSnes = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z",
			B:      "x",
			X:      "a",
			Y:      "s",
			L:      "q",
			R:      "w",
			Select: "v",
			Start:  "enter",
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsN64 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:           "z",
			B:           "x",
			Start:       "enter",
			Up:          "arrowup",    //D-Pad Up
			Down:        "arrowdown",  //D-Pad Down
			Left:        "arrowleft",  //D-Pad Left
			Right:       "arrowright", //D-Pad Right
			L:           "q",
			R:           "w",
			LStickUp:    "t", //Stick up
			LStickDown:  "g", //Stick Down
			LStickLeft:  "f", //Stick Left
			LStickRight: "h", //Stick Right
			RStickUp:    "i", //C-Pad Up
			RStickDown:  "k", //C-Pad Down
			RStickLeft:  "j", //C-Pad Left
			RStickRight: "l", //C-Pad Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsGba = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z",
			B:      "x",
			L:      "a",
			R:      "s",
			Select: "v",
			Start:  "enter",
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsNds = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
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
			L3:     "", //Microphone
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsVb = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:           "z",
			B:           "x",
			L:           "q",
			R:           "w",
			Select:      "v",
			Start:       "enter",
			Up:          "arrowup",    //Left D-Pad Up
			Down:        "arrowdown",  //Left D-Pad Down
			Left:        "arrowleft",  //Left D-Pad Left
			Right:       "arrowright", //Left D-Pad Right
			LStickUp:    "t",          //Right D-Pad Up
			LStickDown:  "g",          //Right D-Pad Down
			LStickLeft:  "f",          //Right D-Pad Left
			LStickRight: "h",          //Right D-Pad Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaMS = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:     "z", //BUTTON 1 / START
			A:     "x", //BUTTON 2
			Up:    "arrowup",
			Down:  "arrowdown",
			Left:  "arrowleft",
			Right: "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaMD = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:      "z", //A
			B:      "x", //B
			A:      "c", //C
			L:      "a", //X
			X:      "s", //Y
			R:      "d", //Z
			Select: "v", //Mode
			Start:  "enter",
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaGG = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:     "z", //BUTTON 1
			A:     "x", //BUTTON 2
			Start: "enter",
			Up:    "arrowup",
			Down:  "arrowdown",
			Left:  "arrowleft",
			Right: "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsSegaSaturn = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:     "z", //A
			B:     "x", //B
			A:     "c", //C
			X:     "a", //X
			L:     "s", //Y
			R:     "d", //Z
			L2:    "q", //L
			R2:    "w", //R
			Start: "enter",
			Up:    "arrowup",
			Down:  "arrowdown",
			Left:  "arrowleft",
			Right: "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControls3do = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			Y:      "z",     //A
			B:      "x",     //B
			A:      "c",     //C
			L:      "a",     //L
			R:      "s",     //R
			Select: "v",     //X
			Start:  "enter", //P
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsAtari2600 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:      "x", //Fire
			Select: "v", //Select
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsAtari7800 = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:      "x", //Fire 1
			A:      "z", //Fire 2
			Select: "v", //Select
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsLynx = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:     "z",
			B:     "x",
			L:     "a", //Option 1
			R:     "s", //Option 2
			Start: "enter",
			Up:    "arrowup",
			Down:  "arrowdown",
			Left:  "arrowleft",
			Right: "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsJaguar = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z",     //A
			B:      "x",     //B
			Y:      "c",     //C
			Select: "v",     //Pause
			Start:  "enter", //Option
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsPCEngine = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "x",     //I
			B:      "z",     //II
			Select: "v",     //Select
			Start:  "enter", //Run
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsNeoGeoPocket = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:     "x",     //B
			B:     "z",     //A
			Start: "enter", //Option
			Up:    "arrowup",
			Down:  "arrowdown",
			Left:  "arrowleft",
			Right: "arrowright",
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsWonderSwan = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z", //A
			B:      "x", //B
			Select: "v",
			Start:  "enter",
			Up:     "arrowup",    //X Cursor Up
			Down:   "arrowdown",  //X Cursor Down
			Left:   "arrowleft",  //X Cursor Left
			Right:  "arrowright", //X Cursor Right
			R2:     "t",          //Y Cursor Up
			L2:     "g",          //Y Cursor Down
			L:      "f",          //Y Cursor Left
			R:      "h",          //Y Cursor Right
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsColecoVision = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "z", //Left Button
			B:      "x", //Right Button
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
			X:      "q", //1
			Y:      "w", //2
			R:      "e", //3
			L:      "r", //4
			R2:     "t", //5
			L2:     "y", //6
			R3:     "u", //7
			L3:     "i", //8
			Select: "o", //*
			Start:  "p", //#
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsPCFX = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			A:      "c", //I
			B:      "x", //II
			X:      "z", //III
			Y:      "d", //IV
			L:      "s", //V
			R:      "a", //VI
			Select: "v",
			Start:  "enter", //Run
			Up:     "arrowup",
			Down:   "arrowdown",
			Left:   "arrowleft",
			Right:  "arrowright",
			L2:     "q", //Mode 1
			R2:     "w", //Mode 2
		},
		Gamepad: DefaultControlsOther.Gamepad,
	}

	DefaultControlsOther = EmulatorControls{
		Keyboard: EmulatorControlsMapping{
			B:               "x",
			Y:               "s",
			Select:          "v",
			Start:           "enter",
			Up:              "arrowup",
			Down:            "arrowdown",
			Left:            "arrowleft",
			Right:           "arrowright",
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

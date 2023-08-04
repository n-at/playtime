package storage

// CoreOptionsFCEUmm https://docs.libretro.com/library/fceumm/
var CoreOptionsFCEUmm = []CoreOption{
	{
		Id:       "fceumm_region",
		Name:     "Region",
		Variants: "Auto|NTSC|PAL|Dendy",
		Default:  "Auto",
	}, {
		Id:       "fceumm_aspect",
		Name:     "Preferred aspect ratio",
		Variants: "8:7 PAR|4:3",
		Default:  "8:7 PAR",
	}, {
		Id:       "fceumm_palette",
		Name:     "Color Palette",
		Variants: "default|asqrealc|nintendo-vc|rgb|yuv-v3|unsaturated-final|sony-cxa2025as-us|pal|bmf-final2|bmf-final3|smooth-fbx|composite-direct-fbx|pvm-style-d93-fbx|ntsc-hardware-fbx|nes-classic-fbx-fs|nescap|wavebeam|raw|custom",
		Default:  "default",
	}, {
		Id:       "fceumm_up_down_allowed",
		Name:     "Allow Opposing Directions",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "fceumm_overscan_h",
		Name:     "Crop Overscan (Horizontal)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "fceumm_overscan_v",
		Name:     "Crop Overscan (Vertical)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "fceumm_nospritelimit",
		Name:     "No Sprite Limit",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "fceumm_sndvolume",
		Name:     "Sound Volume",
		Variants: "0|1|2|3|4|5|6|7|8|9|10",
		Default:  "7",
	}, {
		Id:       "fceumm_sndquality",
		Name:     "Sound Quality",
		Variants: "Low|High|Very High",
		Default:  "Low",
	}, {
		Id:       "fceumm_swapduty",
		Name:     "Swap Duty Cycles",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "fceumm_turbo_enable",
		Name:     "Turbo Enable",
		Variants: "None|Player 1|Player 2|Both",
		Default:  "None",
	}, {
		Id:       "fceumm_turbo_delay",
		Name:     "Turbo Delay (in frames)",
		Variants: "3|5|10|15|30|60|1|2",
		Default:  "3",
	}, {
		Id:       "fceumm_zapper_mode",
		Name:     "Zapper Mode",
		Variants: "lightgun|touchscreen|mouse",
		Default:  "lightgun",
	}, {
		Id:       "fceumm_show_crosshair",
		Name:     "Show Crosshair",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "fceumm_overclocking",
		Name:     "Overclocking",
		Variants: "disabled|2x-Postrender|2x-VBlank",
		Default:  "disabled",
	}, {
		Id:       "fceumm_ramstate",
		Name:     "RAM power up state",
		Variants: "FF|00|random",
		Default:  "FF",
	}, {
		Id:       "fceumm_ntsc_filter",
		Name:     "NTSC Filter",
		Variants: "disabled|composite|svideo|rgb|monochrome",
		Default:  "disabled",
	},
}

// CoreOptionsNestopia https://docs.libretro.com/library/nestopia_ue/
var CoreOptionsNestopia = []CoreOption{
	{
		Id:       "nestopia_blargg_ntsc_filter",
		Name:     "Blargg NTSC filter",
		Variants: "disabled|composite|svideo|rgb|monochrome",
		Default:  "disabled",
	}, {
		Id:       "nestopia_palette",
		Name:     "Palette",
		Variants: "cxa2025as|consumer|canonical|alternative|rgb|pal|composite-direct-fbx|pvm-style-d93-fbx|ntsc-hardware-fbx|nes-classic-fbx-fs|raw|custom",
		Default:  "cxa2025as",
	}, {
		Id:       "nestopia_nospritelimit",
		Name:     "Remove Sprite Limit",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "nestopia_overclock",
		Name:     "CPU Speed (Overclock)",
		Variants: "1x|2x",
		Default:  "1x",
	}, {
		Id:       "nestopia_select_adapter",
		Name:     "4 Player Adapter",
		Variants: "auto|ntsc|famicom",
		Default:  "auto",
	}, {
		Id:       "nestopia_fds_auto_insert",
		Name:     "FDS Auto Insntert",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "nestopia_overscan_v",
		Name:     "Mask Overscan (Vertical)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "nestopia_overscan_h",
		Name:     "Mask Overscan (Horizontal)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "nestopia_aspect",
		Name:     "Preferred aspect ratio",
		Variants: "auto|ntsc|pal|4:3",
		Default:  "auto",
	}, {
		Id:       "nestopia_genie_distortion",
		Name:     "Game Genie Sound Distortion",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "nestopia_favored_system",
		Name:     "System Region",
		Variants: "auto|ntsc|pal|famicom|dendy",
		Default:  "auto",
	}, {
		Id:       "nestopia_ram_power_state",
		Name:     "RAM Power-on State",
		Variants: "0x00|0xFF|random",
		Default:  "0x00",
	}, {
		Id:       "nestopia_turbo_pulse",
		Name:     "Turbo Pulse Speed",
		Variants: "2|3|4|5|6|7|8|9",
		Default:  "2",
	},
}

// CoreOptionsA5200 ? TODO
var CoreOptionsA5200 = []CoreOption{}

// CoreOptionsBeetleVB https://docs.libretro.com/library/beetle_vb/
var CoreOptionsBeetleVB = []CoreOption{
	{
		Id:       "vb_anaglyph_preset",
		Name:     "Anaglyph preset",
		Variants: "Off/red & blue/red & cyan/red & electric cyan/red & green/green & magenta/yellow & blue",
		Default:  "Off",
	}, {
		Id:       "vb_color_mode",
		Name:     "Palette",
		Variants: "black & red/black & white",
		Default:  "black & red",
	}, {
		Id:       "vb_right_analog_to_digital",
		Name:     "Right Analog to Digital",
		Variants: "Off/On/invert x/invert y/invert both",
		Default:  "Off",
	},
}

// CoreOptionsDesmume2015 ? TODO
var CoreOptionsDesmume2015 = []CoreOption{}

// CoreOptionsFBAlpha2012CPS1 ? TODO
var CoreOptionsFBAlpha2012CPS1 = []CoreOption{}

// CoreOptionsFBAlpha2012CPS2 ? TODO
var CoreOptionsFBAlpha2012CPS2 = []CoreOption{}

// CoreOptionsFBNeo https://docs.libretro.com/library/fbneo/ TODO
var CoreOptionsFBNeo = []CoreOption{}

// CoreOptionsGambatte https://docs.libretro.com/library/gambatte/ TODO
var CoreOptionsGambatte = []CoreOption{}

// CoreOptionsGenesisPlusGX https://docs.libretro.com/library/genesis_plus_gx/ TODO
var CoreOptionsGenesisPlusGX = []CoreOption{}

// CoreOptionsHandy https://docs.libretro.com/library/handy/ TODO
var CoreOptionsHandy = []CoreOption{}

// CoreOptionsMAME2003 https://docs.libretro.com/library/mame_2003/ TODO
var CoreOptionsMAME2003 = []CoreOption{}

// CoreOptionsMednafenPSXHW ? TODO
var CoreOptionsMednafenPSXHW = []CoreOption{}

// CoreOptionsMelonDS https://docs.libretro.com/library/melonds/ TODO
var CoreOptionsMelonDS = []CoreOption{}

// ? TODO
var CoreOptionsMGBA = []CoreOption{}

// CoreOptionsMupen64Plus https://docs.libretro.com/library/mupen64plus/ TODO
var CoreOptionsMupen64Plus = []CoreOption{}

// CoreOptionsOpera https://docs.libretro.com/library/opera/ TODO
var CoreOptionsOpera = []CoreOption{}

// CoreOptionsPCSXRearmed https://docs.libretro.com/library/pcsx_rearmed/ TODO
var CoreOptionsPCSXRearmed = []CoreOption{}

// CoreOptionsPicoDrive https://docs.libretro.com/library/picodrive/ TODO
var CoreOptionsPicoDrive = []CoreOption{}

// CoreOptionsProSystem https://docs.libretro.com/library/prosystem/
var CoreOptionsProSystem = []CoreOption{}

// CoreOptionsSnes9x https://docs.libretro.com/library/snes9x/ TODO
var CoreOptionsSnes9x = []CoreOption{}

// CoreOptionsStella2014 https://docs.libretro.com/library/stella/
var CoreOptionsStella2014 = []CoreOption{}

// CoreOptionsVirtualJaguar https://docs.libretro.com/library/virtual_jaguar/ TODO
var CoreOptionsVirtualJaguar = []CoreOption{}

// CoreOptionsYabause https://docs.libretro.com/library/yabause/ TODO
var CoreOptionsYabause = []CoreOption{}

// CoreOptionsPPSSPP https://docs.libretro.com/library/ppsspp/ TODO
var CoreOptionsPPSSPP = []CoreOption{}

/*

{
		Id:       "",
		Name:     "",
		Variants: "",
		Default:  "",
	},

*/

func CoreOptionsByCore(core string) []CoreOption {
	switch core {
	case "fceumm":
		return CoreOptionsFCEUmm
	case "nestopia":
		return CoreOptionsNestopia
	case "snes9x":
		return CoreOptionsSnes9x
	case "gambatte":
		return CoreOptionsGambatte
	case "mgba":
		return CoreOptionsMGBA
	case "beetle_vb":
		return CoreOptionsBeetleVB
	case "melonds":
		return CoreOptionsMelonDS
	case "desmume2015":
		return CoreOptionsDesmume2015
	case "a5200":
		return CoreOptionsA5200
	case "mame2003":
		return CoreOptionsMAME2003
	case "fbneo":
		return CoreOptionsFBNeo
	case "fbalpha2012_cps1":
		return CoreOptionsFBAlpha2012CPS1
	case "fbalpha2012_cps2":
		return CoreOptionsFBAlpha2012CPS2
	case "pcsx_rearmed":
		return CoreOptionsPCSXRearmed
	case "mednafen_psx_hw":
		return CoreOptionsMednafenPSXHW
	case "virtualjaguar":
		return CoreOptionsVirtualJaguar
	case "handy":
		return CoreOptionsHandy
	case "yabause":
		return CoreOptionsYabause
	case "genesis_plus_gx":
		return CoreOptionsGenesisPlusGX
	case "picodrive":
		return CoreOptionsPicoDrive
	case "mupen64plus_next":
		return CoreOptionsMupen64Plus
	case "opera":
		return CoreOptionsOpera
	case "prosystem":
		return CoreOptionsProSystem
	case "stella2014":
		return CoreOptionsStella2014
	}
	return []CoreOption{}
}

func CoreOptionsByPlatform(platform string) map[string][]CoreOption {
	result := make(map[string][]CoreOption)

	for _, core := range Cores[platform] {
		result[core] = CoreOptionsByCore(core)
	}

	return result
}

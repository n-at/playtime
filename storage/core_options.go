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

// CoreOptionsDesmume2015 https://docs.libretro.com/library/desmume_2015/ TODO
var CoreOptionsDesmume2015 = []CoreOption{
	{
		Id:       "desmume_internal_resolution",
		Name:     "Internal resolution",
		Variants: "256x192|512x384|768x576|1024x768|1280x960|1536x1152|1792x1344|2048x1536|2304x1728|2560x1920",
		Default:  "256x192",
	}, {
		Id:       "desmume_num_cores",
		Name:     "CPU cores",
		Variants: "1|2|3|4",
		Default:  "1",
	}, {
		Id:       "desmume_cpu_mode",
		Name:     "CPU mode",
		Variants: "jit|interpreter",
		Default:  "jit",
	}, {
		Id:       "desmume_jit_block_size",
		Name:     "JIT block size",
		Variants: "1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100",
		Default:  "12",
	}, {
		Id:       "desmume_screens_layout",
		Name:     "Screen layout",
		Variants: "top/bottom|bottom/top|left/right|right/left|top only|bottom only|quick switch|hybrid/top|hybrid/bottom",
		Default:  "top/bottom",
	}, {
		Id:       "desmume_hybrid_layout_scale",
		Name:     "Hybrid layout scale",
		Variants: "1|3",
		Default:  "1",
	}, {
		Id:       "desmume_hybrid_showboth_screens",
		Name:     "Hybrid layout show both screen",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_hybrid_cursor_always_smallscreen",
		Name:     "Hybrid layout cursor always on small screen",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_pointer_mouse",
		Name:     "Enable mouse/pointer",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_pointer_type",
		Name:     "Pointer type",
		Variants: "mouse|touch",
		Default:  "mouse",
	}, {
		Id:       "desmume_mouse_speed",
		Name:     "Mouse Speed",
		Variants: "1.0|1.5|2.0|0.125|0.25|0.5",
		Default:  "1.0",
	}, {
		Id:       "desmume_pointer_colour",
		Name:     "Pointer Colour",
		Variants: "white|black|red|blue|yellow",
		Default:  "white",
	}, {
		Id:       "desmume_pointer_device_l",
		Name:     "Pointer mode l-analog",
		Variants: "none|emulated|absolute|pressed",
		Default:  "none",
	}, {
		Id:       "desmume_pointer_device_r",
		Name:     "Pointer mode r-analog",
		Variants: "none|emulated|absolute|pressed",
		Default:  "none",
	}, {
		Id:       "desmume_pointer_device_deadzone",
		Name:     "Emulated pointer deadzone percent",
		Variants: "15|20|25|30|35|0|5|10",
		Default:  "15",
	}, {
		Id:       "desmume_pointer_device_acceleration_mod",
		Name:     "Emulated pointer acceleration modifier percent",
		Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100",
		Default:  "0",
	}, {
		Id:       "desmume_pointer_stylus_pressure",
		Name:     "Emulated stylus pressure modifier percent",
		Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100",
		Default:  "50",
	}, {
		Id:       "desmume_pointer_stylus_jitter",
		Name:     "Enable emulated stylus jitter",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "desmume_load_to_memory",
		Name:     "Load Game into Memory",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "desmume_advanced_timing",
		Name:     "Enable Advanced Bus-Level Timing",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_firmware_language",
		Name:     "Firmware language",
		Variants: "Auto|English|Japanese|French|German|Italian|Spanish",
		Default:  "Auto",
	}, {
		Id:       "desmume_frameskip",
		Name:     "Frameskip",
		Variants: "0|1|2|3|4|5|6|7|8|9",
		Default:  "0",
	}, {
		Id:       "desmume_screens_gap",
		Name:     "Screen Gap",
		Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100",
		Default:  "0",
	}, {
		Id:       "desmume_gfx_edgemark",
		Name:     "Enable Edgemark",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_gfx_linehack",
		Name:     "Enable Line Hack",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "desmume_gfx_txthack",
		Name:     "Enable TXT Hack",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "desmume_mic_force_enable",
		Name:     "Force Microphone Enable",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "desmume_mic_mode",
		Name:     "Microphone Simulation Settings",
		Variants: "internal|sample|random|physical",
		Default:  "internal",
	},
}

// CoreOptionsFBAlpha2012CPS1
var CoreOptionsFBAlpha2012CPS1 = []CoreOption{}

// CoreOptionsFBAlpha2012CPS2
var CoreOptionsFBAlpha2012CPS2 = []CoreOption{}

// CoreOptionsFBNeo https://docs.libretro.com/library/fbneo/
var CoreOptionsFBNeo = []CoreOption{}

// CoreOptionsGambatte https://docs.libretro.com/library/gambatte/
var CoreOptionsGambatte = []CoreOption{
	{
		Id:       "gambatte_up_down_allowed",
		Name:     "Allow Opposing Directions",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "gambatte_gb_colorization",
		Name:     "GB Colorization",
		Variants: "disabled|auto|GBC|SGB|internal|custom",
		Default:  "disabled",
	}, {
		Id:       "gambatte_gb_internal_palette",
		Name:     "Internal Palette",
		Variants: "GB - DMG|GB - Pocket|GB - Light|GBC - Blue|GBC - Brown|GBC - Dark Blue|GBC - Dark Brown|GBC - Dark Green|GBC - Grayscale|GBC - Green|GBC - Inverted|GBC - Orange|GBC - Pastel Mix|GBC - Red|GBC - Yellow|SGB - 1A|SGB - 1B|SGB - 1C|SGB - 1D|SGB - 1E|SGB - 1F|SGB - 1G|SGB - 1H|SGB - 2A|SGB - 2B|SGB - 2C|SGB - 2D|SGB - 2E|SGB - 2F|SGB - 2G|SGB - 2H|SGB - 3A|SGB - 3B|SGB - 3C|SGB - 3D|SGB - 3E|SGB - 3F|SGB - 3G|SGB - 3H|SGB - 4A|SGB - 4B|SGB - 4C|SGB - 4D|SGB - 4E|SGB - 4F|SGB - 4G|SGB - 4H|Special 1|Special 2|Special 3",
		Default:  "GB - DMG",
	}, {
		Id:       "gambatte_gbc_color_correction",
		Name:     "Color correction",
		Variants: "GBC only|always|disabled",
		Default:  "GBC only",
	}, {
		Id:       "gambatte_gbc_color_correction_mode",
		Name:     "Color correction mode",
		Variants: "accurate|fast",
		Default:  "accurate",
	}, {
		Id:       "gambatte_gbc_frontlight_position",
		Name:     "Color correction - frontlight position",
		Variants: "central|above screen|below screen",
		Default:  "central",
	}, {
		Id:       "gambatte_dark_filter_level",
		Name:     "Dark Filter Level (percent)",
		Variants: "0|5|10|15|20|25|30|35|40|45|50",
		Default:  "0",
	}, {
		Id:       "gambatte_gb_hwmode",
		Name:     "Emulated hardware",
		Variants: "Auto|GB|GBC|GBA",
		Default:  "Auto",
	}, {
		Id:       "gambatte_gb_bootloader",
		Name:     "Use official bootloader",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "gambatte_mix_frames",
		Name:     "Mix frames",
		Variants: "disabled|accurate|fast",
		Default:  "disabled",
	},
}

// CoreOptionsGenesisPlusGX https://docs.libretro.com/library/genesis_plus_gx/ TODO
var CoreOptionsGenesisPlusGX = []CoreOption{}

// CoreOptionsHandy https://docs.libretro.com/library/handy/
var CoreOptionsHandy = []CoreOption{
	{
		Id:       "handy_rot",
		Name:     "Display rotation",
		Variants: "0|90|240",
		Default:  "0",
	},
}

// CoreOptionsMAME2003 https://docs.libretro.com/library/mame_2003/ TODO
var CoreOptionsMAME2003 = []CoreOption{}

// CoreOptionsMednafenPSXHW ? TODO
var CoreOptionsMednafenPSXHW = []CoreOption{}

// CoreOptionsMelonDS https://docs.libretro.com/library/melonds/ TODO
var CoreOptionsMelonDS = []CoreOption{}

// CoreOptionsMGBA https://docs.libretro.com/library/mgba/
var CoreOptionsMGBA = []CoreOption{
	{
		Id:       "mgba_solar_sensor_level",
		Name:     "Solar sensor level",
		Variants: "0|1|2|3|4|5|6|7|8|9|10",
		Default:  "0",
	}, {
		Id:       "mgba_allow_opposing_directions",
		Name:     "Allow opposing directional input",
		Variants: "OFF|ON",
		Default:  "OFF",
	}, {
		Id:       "mgba_gb_model",
		Name:     "Game Boy model",
		Variants: "Autodetect|Game Boy|Super Game Boy|Game Boy Color|Game Boy Advance",
		Default:  "Autodetect",
	}, {
		Id:       "mgba_use_bios",
		Name:     "Use BIOS file if found",
		Variants: "ON|OFF",
		Default:  "ON",
	}, {
		Id:       "mgba_skip_bios",
		Name:     "Skip BIOS intro",
		Variants: "OFF|ON",
		Default:  "OFF",
	}, {
		Id:       "mgba_sgb_borders",
		Name:     "Use Super Game Boy borders",
		Variants: "ON|OFF",
		Default:  "ON",
	}, {
		Id:       "mgba_idle_optimization",
		Name:     "Idle loop removal",
		Variants: "Remove Known|Detect and Remove|Don't Remove",
		Default:  "Remove Known",
	}, {
		Id:       "mgba_frameskip",
		Name:     "Frameskip",
		Variants: "0|1|2|3|4|5|6|7|8|9|10",
		Default:  "0",
	},
}

// CoreOptionsMupen64Plus https://docs.libretro.com/library/mupen64plus/ TODO
var CoreOptionsMupen64Plus = []CoreOption{}

// CoreOptionsOpera https://docs.libretro.com/library/opera/
var CoreOptionsOpera = []CoreOption{
	{
		Id:       "opera_cpu_overclock",
		Name:     "CPU overclock",
		Variants: "1.0x (12.50Mhz)|1.1x (13.75Mhz)|1.2x (15.00Mhz)|1.5x (18.75Mhz)|1.6x (20.00Mhz)|1.8x (22.50Mhz)|2.0x (25.00Mhz)",
		Default:  "1.0x (12.50Mhz)",
	}, {
		Id:       "opera_high_resolution",
		Name:     "High Resolution",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "opera_nvram_storage",
		Name:     "NVRAM Storage",
		Variants: "per game|shared",
		Default:  "per game",
	}, {
		Id:       "opera_active_devices",
		Name:     "Active Devices",
		Variants: "1|2|3|4|5|6|7|8|0",
		Default:  "1",
	}, {
		Id:       "opera_hack_timing_1",
		Name:     "Timing Hack 1 (Crash 'n Burn)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "opera_hack_timing_3",
		Name:     "Timing Hack 3 (Dinopark Tycoon)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "opera_hack_timing_5",
		Name:     "Timing Hack 5 (Microcosm)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "opera_hack_timing_6",
		Name:     "Timing Hack 6 (Alone in the Dark)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "opera_hack_graphics_step_y",
		Name:     "Graphics Step Y Hack (Samurai Shodown)",
		Variants: "disabled|enabled",
		Default:  "disabled",
	},
}

// CoreOptionsPCSXRearmed https://docs.libretro.com/library/pcsx_rearmed/ TODO
var CoreOptionsPCSXRearmed = []CoreOption{}

// CoreOptionsPicoDrive https://docs.libretro.com/library/picodrive/
var CoreOptionsPicoDrive = []CoreOption{
	{
		Id:       "picodrive_input1",
		Name:     "Input device 1",
		Variants: "3 button pad|6 button pad|None",
		Default:  "3 button pad",
	}, {
		Id:       "picodrive_input2",
		Name:     "Input device 2",
		Variants: "3 button pad|6 button pad|None",
		Default:  "3 button pad",
	}, {
		Id:       "picodrive_sprlim",
		Name:     "No sprite limit",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "picodrive_ramcart",
		Name:     "MegaCD RAM cart",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "picodrive_region",
		Name:     "Region",
		Variants: "Auto|Japan NTSC|Japan PAL|US|Europe",
		Default:  "Auto",
	}, {
		Id:       "picodrive_aspect",
		Name:     "Core-provided aspect ratio",
		Variants: "PAR|4|3|CRT",
		Default:  "PAR",
	}, {
		Id:       "picodrive_overscan",
		Name:     "Show Overscan",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "picodrive_overclk68k",
		Name:     "68k overclock",
		Variants: "disabled|+25%|+50%|+75%|+100%|+200%|+400%",
		Default:  "disabled",
	}, {
		Id:       "picodrive_drc",
		Name:     "Dynamic recompilers",
		Variants: "enabled|disabled",
		Default:  "enabled",
	}, {
		Id:       "picodrive_audio_filter",
		Name:     "Audio filter",
		Variants: "disabled|low-pass",
		Default:  "disabled",
	}, {
		Id:       "picodrive_lowpass_range",
		Name:     "Low-pass filter %",
		Variants: "60|65|70|75|80|85|90|95|5|10|15|20|25|30|35|40|45|50|55",
		Default:  "60",
	},
}

// CoreOptionsProSystem https://docs.libretro.com/library/prosystem/
var CoreOptionsProSystem = []CoreOption{}

// CoreOptionsSnes9x https://docs.libretro.com/library/snes9x/ TODO
var CoreOptionsSnes9x = []CoreOption{}

// CoreOptionsStella2014 https://docs.libretro.com/library/stella/
var CoreOptionsStella2014 = []CoreOption{}

// CoreOptionsVirtualJaguar https://docs.libretro.com/library/virtual_jaguar/
var CoreOptionsVirtualJaguar = []CoreOption{
	{
		Id:       "virtualjaguar_usefastblitter",
		Name:     "Fast Blitter",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "virtualjaguar_doom_res_hack",
		Name:     "Doom Res Hack",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "virtualjaguar_bios",
		Name:     "Bios",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "virtualjaguar_pal",
		Name:     "PAL",
		Variants: "disabled|enabled",
		Default:  "disabled",
	},
}

// CoreOptionsYabause https://docs.libretro.com/library/yabause/
var CoreOptionsYabause = []CoreOption{
	{
		Id:       "yabause_frameskip",
		Name:     "Frameskip",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "yabause_force_hle_bios",
		Name:     "Force HLE BIOS",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "yabause_addon_cart",
		Name:     "Addon Cartridge",
		Variants: "none|1M_ram|4M_ram",
		Default:  "none",
	}, {
		Id:       "yabause_multitap_port1",
		Name:     "6Player Adaptor on Port 1",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "yabause_multitap_port2",
		Name:     "6Player Adaptor on Port 2",
		Variants: "disabled|enabled",
		Default:  "disabled",
	}, {
		Id:       "yabause_numthreads",
		Name:     "Number of Threads",
		Variants: "1|2|4|8|16|32",
		Default:  "4",
	},
}

// CoreOptionsPPSSPP https://docs.libretro.com/library/ppsspp/ TODO
var CoreOptionsPPSSPP = []CoreOption{}

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

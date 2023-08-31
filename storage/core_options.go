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

// CoreOptionsDesmume2015 https://docs.libretro.com/library/desmume_2015/
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

// CoreOptionsGenesisPlusGX https://docs.libretro.com/library/genesis_plus_gx/ (extracted from emulator)
var CoreOptionsGenesisPlusGX = []CoreOption{
	{Id: "genesis_plus_gx_frameskip_threshold", Name: "genesis_plus_gx_frameskip_threshold", Variants: "15|18|21|24|27|30|33|36|39|42|45|48|51|54|57|60", Default: "33"},
	{Id: "genesis_plus_gx_lowpass_range", Name: "genesis_plus_gx_lowpass_range", Variants: "5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95", Default: "60"},
	{Id: "genesis_plus_gx_psg_preamp", Name: "genesis_plus_gx_psg_preamp", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100|105|110|115|120|125|130|135|140|145|150|155|160|165|170|175|180|185|190|195|200", Default: "150"},
	{Id: "genesis_plus_gx_fm_preamp", Name: "genesis_plus_gx_fm_preamp", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100|105|110|115|120|125|130|135|140|145|150|155|160|165|170|175|180|185|190|195|200", Default: "100"},
	{Id: "genesis_plus_gx_cdda_volume", Name: "genesis_plus_gx_cdda_volume", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100", Default: "100"},
	{Id: "genesis_plus_gx_pcm_volume", Name: "genesis_plus_gx_pcm_volume", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100", Default: "100"},
	{Id: "genesis_plus_gx_audio_eq_low", Name: "genesis_plus_gx_audio_eq_low", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100", Default: "100"},
	{Id: "genesis_plus_gx_audio_eq_mid", Name: "genesis_plus_gx_audio_eq_mid", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100", Default: "100"},
	{Id: "genesis_plus_gx_audio_eq_high", Name: "genesis_plus_gx_audio_eq_high", Variants: "0|5|10|15|20|25|30|35|40|45|50|55|60|65|70|75|80|85|90|95|100", Default: "100"},
	{Id: "genesis_plus_gx_enhanced_vscroll_limit", Name: "genesis_plus_gx_enhanced_vscroll_limit", Variants: "2|3|4|5|6|7|8|9|10|11|12|13|14|15|16", Default: "8"},
	{Id: "genesis_plus_gx_psg_channel_0_volume", Name: "genesis_plus_gx_psg_channel_0_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_psg_channel_1_volume", Name: "genesis_plus_gx_psg_channel_1_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_psg_channel_2_volume", Name: "genesis_plus_gx_psg_channel_2_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_psg_channel_3_volume", Name: "genesis_plus_gx_psg_channel_3_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_0_volume", Name: "genesis_plus_gx_md_channel_0_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_1_volume", Name: "genesis_plus_gx_md_channel_1_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_2_volume", Name: "genesis_plus_gx_md_channel_2_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_3_volume", Name: "genesis_plus_gx_md_channel_3_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_4_volume", Name: "genesis_plus_gx_md_channel_4_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_md_channel_5_volume", Name: "genesis_plus_gx_md_channel_5_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_0_volume", Name: "genesis_plus_gx_sms_fm_channel_0_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_1_volume", Name: "genesis_plus_gx_sms_fm_channel_1_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_2_volume", Name: "genesis_plus_gx_sms_fm_channel_2_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_3_volume", Name: "genesis_plus_gx_sms_fm_channel_3_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_4_volume", Name: "genesis_plus_gx_sms_fm_channel_4_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_5_volume", Name: "genesis_plus_gx_sms_fm_channel_5_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_6_volume", Name: "genesis_plus_gx_sms_fm_channel_6_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_7_volume", Name: "genesis_plus_gx_sms_fm_channel_7_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
	{Id: "genesis_plus_gx_sms_fm_channel_8_volume", Name: "genesis_plus_gx_sms_fm_channel_8_volume", Variants: "0|10|20|30|40|50|60|70|80|90|100", Default: "100"},
}

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

// CoreOptionsMednafenPSXHW (extracted from emulator)
var CoreOptionsMednafenPSXHW = []CoreOption{
	{Id: "beetle_psx_hw_negcon_deadzone", Name: "beetle_psx_hw_negcon_deadzone", Variants: "0%|5%|10%|15%|20%|25%|30%", Default: "0%"},
	{Id: "beetle_psx_hw_memcard_left_index", Name: "beetle_psx_hw_memcard_left_index", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63", Default: "0"},
	{Id: "beetle_psx_hw_widescreen_hack_aspect_ratio", Name: "beetle_psx_hw_widescreen_hack_aspect_ratio", Variants: "16:9|16:10|18:9|19:9|20:9|21:9|32:9", Default: "16:9"},
	{Id: "beetle_psx_hw_initial_scanline", Name: "beetle_psx_hw_initial_scanline", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40", Default: "0"},
	{Id: "beetle_psx_hw_initial_scanline_pal", Name: "beetle_psx_hw_initial_scanline_pal", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40", Default: "0"},
}

// CoreOptionsMelonDS https://docs.libretro.com/library/melonds/ (extracted from emulator)
var CoreOptionsMelonDS = []CoreOption{
	{Id: "melonds_console_mode", Name: "melonds_console_mode", Variants: "DS|DSi", Default: "DS"},
	{Id: "melonds_language", Name: "melonds_language", Variants: "Japanese|English|French|German|Italian|Spanish", Default: "English"},
	{Id: "melonds_mic_input", Name: "melonds_mic_input", Variants: "Blow Noise|White Noise", Default: "Blow Noise"},
	{Id: "melonds_audio_bitrate", Name: "melonds_audio_bitrate", Variants: "Automatic|10-bit|16-bit", Default: "Automatic"},
	{Id: "melonds_audio_interpolation", Name: "melonds_audio_interpolation", Variants: "None|Linear|Cosine|Cubic", Default: "None"},
	{Id: "melonds_touch_mode", Name: "melonds_touch_mode", Variants: "Mouse|Touch|Joystick", Default: "Mouse"},
	{Id: "melonds_swapscreen_mode", Name: "melonds_swapscreen_mode", Variants: "Toggle|Hold", Default: "Toggle"},
	{Id: "melonds_screen_layout", Name: "melonds_screen_layout", Variants: "Top/Bottom|Bottom/Top|Left/Right|Right/Left|Top Only|Bottom Only|Hybrid Top|Hybrid Bottom", Default: "Top/Bottom"},
	{Id: "melonds_screen_gap", Name: "melonds_screen_gap", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100|101|102|103|104|105|106|107|108|109|110|111|112|113|114|115|116|117|118|119|120|121|122|123|124|125|126", Default: "0"},
	{Id: "melonds_hybrid_small_screen", Name: "melonds_hybrid_small_screen", Variants: "Bottom|Top|Duplicate", Default: "Bottom"},
}

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

// CoreOptionsMupen64Plus https://docs.libretro.com/library/mupen64plus/ (extracted from emulator)
var CoreOptionsMupen64Plus = []CoreOption{
	{Id: "mupen64plus-43screensize", Name: "mupen64plus-43screensize", Variants: "320x240|640x480|960x720|1280x960|1440x1080|1600x1200|1920x1440|2240x1680|2560x1920|2880x2160|3200x2400|3520x2640|3840x2880", Default: "640x480"},
	{Id: "mupen64plus-BilinearMode", Name: "mupen64plus-BilinearMode", Variants: "3point|standard", Default: "standard"},
	{Id: "mupen64plus-MultiSampling", Name: "mupen64plus-MultiSampling", Variants: "0|2|4|8|16", Default: "0"},
	{Id: "mupen64plus-FXAA", Name: "mupen64plus-FXAA", Variants: "0|1", Default: "0"},
	{Id: "mupen64plus-EnableCopyDepthToRDRAM", Name: "mupen64plus-EnableCopyDepthToRDRAM", Variants: "Software|FromMem", Default: "Software"},
	{Id: "mupen64plus-BackgroundMode", Name: "mupen64plus-BackgroundMode", Variants: "Stripped|OnePiece", Default: "OnePiece"},
	{Id: "mupen64plus-OverscanTop", Name: "mupen64plus-OverscanTop", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50", Default: "0"},
	{Id: "mupen64plus-OverscanLeft", Name: "mupen64plus-OverscanLeft", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50", Default: "0"},
	{Id: "mupen64plus-OverscanRight", Name: "mupen64plus-OverscanRight", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50", Default: "0"},
	{Id: "mupen64plus-OverscanBottom", Name: "mupen64plus-OverscanBottom", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50", Default: "0"},
	{Id: "mupen64plus-MaxHiResTxVramLimit", Name: "mupen64plus-MaxHiResTxVramLimit", Variants: "0|500|1000|1500|2000|2500|3000|3500|4000", Default: "0"},
	{Id: "mupen64plus-MaxTxCacheSize", Name: "mupen64plus-MaxTxCacheSize", Variants: "1500|4000|8000", Default: "8000"},
	{Id: "mupen64plus-txFilterMode", Name: "mupen64plus-txFilterMode", Variants: "None|Smooth filtering 1|Smooth filtering 2|Smooth filtering 3|Smooth filtering 4|Sharp filtering 1|Sharp filtering 2", Default: "None"},
	{Id: "mupen64plus-txEnhancementMode", Name: "mupen64plus-txEnhancementMode", Variants: "None|As Is|X2|X2SAI|HQ2X|HQ2XS|LQ2X|LQ2XS|HQ4X|2xBRZ|3xBRZ|4xBRZ|5xBRZ|6xBRZ", Default: "None"},
	{Id: "mupen64plus-Framerate", Name: "mupen64plus-Framerate", Variants: "Original|Fullspeed", Default: "Original"},
	{Id: "mupen64plus-virefresh", Name: "mupen64plus-virefresh", Variants: "Auto|1500|2200", Default: "Auto"},
	{Id: "mupen64plus-CountPerOp", Name: "mupen64plus-CountPerOp", Variants: "0|1|2|3|4|5", Default: "0"},
	{Id: "mupen64plus-CountPerOpDenomPot", Name: "mupen64plus-CountPerOpDenomPot", Variants: "0|1|2|3|4|5|6|7|8|9|10|11", Default: "0"},
	{Id: "mupen64plus-astick-deadzone", Name: "mupen64plus-astick-deadzone", Variants: "0|5|10|15|20|25|30", Default: "15"},
	{Id: "mupen64plus-astick-sensitivity", Name: "mupen64plus-astick-sensitivity", Variants: "50|55|60|65|70|75|80|85|90|95|100|105|110|115|120|125|130|135|140|145|150", Default: "100"},
	{Id: "mupen64plus-r-cbutton", Name: "mupen64plus-r-cbutton", Variants: "C1|C2|C3|C4", Default: "C1"},
	{Id: "mupen64plus-l-cbutton", Name: "mupen64plus-l-cbutton", Variants: "C1|C2|C3|C4", Default: "C2"},
	{Id: "mupen64plus-d-cbutton", Name: "mupen64plus-d-cbutton", Variants: "C1|C2|C3|C4", Default: "C3"},
	{Id: "mupen64plus-u-cbutton", Name: "mupen64plus-u-cbutton", Variants: "C1|C2|C3|C4", Default: "C4"},
	{Id: "mupen64plus-pak1", Name: "mupen64plus-pak1", Variants: "none|memory|rumble|transfer", Default: "memory"},
	{Id: "mupen64plus-pak2", Name: "mupen64plus-pak2", Variants: "none|memory|rumble|transfer", Default: "none"},
	{Id: "mupen64plus-pak3", Name: "mupen64plus-pak3", Variants: "none|memory|rumble|transfer", Default: "none"},
	{Id: "mupen64plus-pak4", Name: "mupen64plus-pak4", Variants: "none|memory|rumble|transfer", Default: "none"},
}

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

// CoreOptionsPCSXRearmed https://docs.libretro.com/library/pcsx_rearmed/ (extracted from emulator)
var CoreOptionsPCSXRearmed = []CoreOption{
	{Id: "pcsx_rearmed_psxclock", Name: "pcsx_rearmed_psxclock", Variants: "30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63|64|65|66|67|68|69|70|71|72|73|74|75|76|77|78|79|80|81|82|83|84|85|86|87|88|89|90|91|92|93|94|95|96|97|98|99|100", Default: "57"},
	{Id: "pcsx_rearmed_frameskip_threshold", Name: "pcsx_rearmed_frameskip_threshold", Variants: "15|18|21|24|27|30|33|36|39|42|45|48|51|54|57|60", Default: "33"},
	{Id: "pcsx_rearmed_frameskip_interval", Name: "pcsx_rearmed_frameskip_interval", Variants: "1|2|3|4|5|6|7|8|9|10", Default: "3"},
	{Id: "pcsx_rearmed_gpu_slow_llists", Name: "pcsx_rearmed_gpu_slow_llists", Variants: "auto", Default: "auto"},
	{Id: "pcsx_rearmed_input_sensitivity", Name: "pcsx_rearmed_input_sensitivity", Variants: "0.05|0.10|0.15|0.20|0.25|0.30|0.35|0.40|0.45|0.50|0.55|0.60|0.65|0.70|0.75|0.80|0.85|0.90|0.95|1.00|1.05|1.10|1.15|1.20|1.25|1.30|1.35|1.40|1.45|1.50|1.55|1.60|1.65|1.70|1.75|1.80|1.85|1.90|1.95|2.00", Default: "1.00"},
	{Id: "pcsx_rearmed_gunconadjustx", Name: "pcsx_rearmed_gunconadjustx", Variants: "-25|-24|-23|-22|-21|-20|-19|-18|-17|-16|-15|-14|-13|-12|-11|-10|-9|-8|-7|-6|-5|-4|-3|-2|-1|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25", Default: "0"},
	{Id: "pcsx_rearmed_gunconadjusty", Name: "pcsx_rearmed_gunconadjusty", Variants: "-25|-24|-23|-22|-21|-20|-19|-18|-17|-16|-15|-14|-13|-12|-11|-10|-9|-8|-7|-6|-5|-4|-3|-2|-1|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25", Default: "0"},
	{Id: "pcsx_rearmed_gunconadjustratiox", Name: "pcsx_rearmed_gunconadjustratiox", Variants: "0.75|0.76|0.77|0.78|0.79|0.80|0.81|0.82|0.83|0.84|0.85|0.86|0.87|0.88|0.89|0.90|0.91|0.92|0.93|0.94|0.95|0.96|0.97|0.98|0.99|1.00|1.01|1.02|1.03|1.04|1.05|1.06|1.07|1.08|1.09|1.10|1.11|1.12|1.13|1.14|1.15|1.16|1.17|1.18|1.19|1.20|1.21|1.22|1.23|1.24|1.25", Default: "1.00"},
	{Id: "pcsx_rearmed_gunconadjustratioy", Name: "pcsx_rearmed_gunconadjustratioy", Variants: "0.75|0.76|0.77|0.78|0.79|0.80|0.81|0.82|0.83|0.84|0.85|0.86|0.87|0.88|0.89|0.90|0.91|0.92|0.93|0.94|0.95|0.96|0.97|0.98|0.99|1.00|1.01|1.02|1.03|1.04|1.05|1.06|1.07|1.08|1.09|1.10|1.11|1.12|1.13|1.14|1.15|1.16|1.17|1.18|1.19|1.20|1.21|1.22|1.23|1.24|1.25", Default: "1.00"},
}

// CoreOptionsProSystem https://docs.libretro.com/library/prosystem/
var CoreOptionsProSystem = []CoreOption{}

// CoreOptionsSnes9x https://docs.libretro.com/library/snes9x/ (extracted from emulator)
var CoreOptionsSnes9x = []CoreOption{
	{Id: "snes9x_aspect", Name: "snes9x_aspect", Variants: "4:3", Default: "4:3"},
	{Id: "snes9x_overclock_superfx", Name: "snes9x_overclock_superfx", Variants: "50%|60%|70%|80%|90%|100%|150%|200%|250%|300%|350%|400%|450%|500%", Default: "100%"},
	{Id: "snes9x_superscope_crosshair", Name: "snes9x_superscope_crosshair", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16", Default: "2"},
	{Id: "snes9x_superscope_color", Name: "snes9x_superscope_color", Variants: "White|White (blend)|Red|Red (blend)|Orange|Orange (blend)|Yellow|Yellow (blend)|Green|Green (blend)|Cyan|Cyan (blend)|Sky|Sky (blend)|Blue|Blue (blend)|Violet|Violet (blend)|Pink|Pink (blend)|Purple|Purple (blend)|Black|Black (blend)|25% Grey|25% Grey (blend)|50% Grey|50% Grey (blend)|75% Grey|75% Grey (blend)", Default: "White"},
	{Id: "snes9x_justifier1_crosshair", Name: "snes9x_justifier1_crosshair", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16", Default: "4"},
	{Id: "snes9x_justifier1_color", Name: "snes9x_justifier1_color", Variants: "Blue|Blue (blend)|Violet|Violet (blend)|Pink|Pink (blend)|Purple|Purple (blend)|Black|Black (blend)|25% Grey|25% Grey (blend)|50% Grey|50% Grey (blend)|75% Grey|75% Grey (blend)|White|White (blend)|Red|Red (blend)|Orange|Orange (blend)|Yellow|Yellow (blend)|Green|Green (blend)|Cyan|Cyan (blend)|Sky|Sky (blend)", Default: "Blue"},
	{Id: "snes9x_justifier2_crosshair", Name: "snes9x_justifier2_crosshair", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16", Default: "4"},
	{Id: "snes9x_justifier2_color", Name: "snes9x_justifier2_color", Variants: "Pink|Pink (blend)|Purple|Purple (blend)|Black|Black (blend)|25% Grey|25% Grey (blend)|50% Grey|50% Grey (blend)|75% Grey|75% Grey (blend)|White|White (blend)|Red|Red (blend)|Orange|Orange (blend)|Yellow|Yellow (blend)|Green|Green (blend)|Cyan|Cyan (blend)|Sky|Sky (blend)|Blue|Blue (blend)|Violet|Violet (blend)", Default: "Pink"},
	{Id: "snes9x_rifle_crosshair", Name: "snes9x_rifle_crosshair", Variants: "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16", Default: "2"},
	{Id: "snes9x_rifle_color", Name: "snes9x_rifle_color", Variants: "White|White (blend)|Red|Red (blend)|Orange|Orange (blend)|Yellow|Yellow (blend)|Green|Green (blend)|Cyan|Cyan (blend)|Sky|Sky (blend)|Blue|Blue (blend)|Violet|Violet (blend)|Pink|Pink (blend)|Purple|Purple (blend)|Black|Black (blend)|25% Grey|25% Grey (blend)|50% Grey|50% Grey (blend)|75% Grey|75% Grey (blend)", Default: "White"},
}

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

// CoreOptionsParallelN64 (extracted from emulator)
var CoreOptionsParallelN64 = []CoreOption{
	{Id: "parallel-n64-cpucore", Name: "parallel-n64-cpucore", Variants: "cached_interpreter|pure_interpreter", Default: "cached_interpreter"},
	{Id: "parallel-n64-audio-buffer-size", Name: "parallel-n64-audio-buffer-size", Variants: "2048|1024", Default: "2048"},
	{Id: "parallel-n64-astick-deadzone", Name: "parallel-n64-astick-deadzone", Variants: "15|20|25|30|0|5|10", Default: "15"},
	{Id: "parallel-n64-astick-sensitivity", Name: "parallel-n64-astick-sensitivity", Variants: "100|105|110|115|120|125|130|135|140|145|150|200|50|55|60|65|70|75|80|85|90|95", Default: "100"},
	{Id: "parallel-n64-pak1", Name: "parallel-n64-pak1", Variants: "none|memory|rumble", Default: "none"},
	{Id: "parallel-n64-pak2", Name: "parallel-n64-pak2", Variants: "none|memory|rumble", Default: "none"},
	{Id: "parallel-n64-pak3", Name: "parallel-n64-pak3", Variants: "none|memory|rumble", Default: "none"},
	{Id: "parallel-n64-pak4", Name: "parallel-n64-pak4", Variants: "none|memory|rumble", Default: "none"},
	{Id: "parallel-n64-gfxplugin-accuracy", Name: "parallel-n64-gfxplugin-accuracy", Variants: "veryhigh|high|medium|low", Default: "veryhigh"},
	{Id: "parallel-n64-gfxplugin", Name: "parallel-n64-gfxplugin", Variants: "auto|glide64|gln64|rice|angrylion", Default: "auto"},
	{Id: "parallel-n64-rspplugin", Name: "parallel-n64-rspplugin", Variants: "auto|hle|cxd4", Default: "auto"},
	{Id: "parallel-n64-screensize", Name: "parallel-n64-screensize", Variants: "640x480|960x720|1280x960|1440x1080|1600x1200|1920x1440|2240x1680|2880x2160|5760x4320|320x240", Default: "640x480"},
	{Id: "parallel-n64-aspectratiohint", Name: "parallel-n64-aspectratiohint", Variants: "normal|widescreen", Default: "normal"},
	{Id: "parallel-n64-filtering", Name: "parallel-n64-filtering", Variants: "automatic|N64 3-point|bilinear|nearest", Default: "automatic"},
	{Id: "parallel-n64-polyoffset-factor", Name: "parallel-n64-polyoffset-factor", Variants: "-3.0|-2.5|-2.0|-1.5|-1.0|-0.5|0.0|0.5|1.0|1.5|2.0|2.5|3.0|3.5|4.0|4.5|5.0|-3.5|-4.0|-4.5|-5.0", Default: "-3.0"},
	{Id: "parallel-n64-polyoffset-units", Name: "parallel-n64-polyoffset-units", Variants: "-3.0|-2.5|-2.0|-1.5|-1.0|-0.5|0.0|0.5|1.0|1.5|2.0|2.5|3.0|3.5|4.0|4.5|5.0|-3.5|-4.0|-4.5|-5.0", Default: "-3.0"},
	{Id: "parallel-n64-angrylion-vioverlay", Name: "parallel-n64-angrylion-vioverlay", Variants: "Filtered|AA+Blur|AA+Dedither|AA only|Unfiltered|Depth|Coverage", Default: "Filtered"},
	{Id: "parallel-n64-angrylion-sync", Name: "parallel-n64-angrylion-sync", Variants: "Low|Medium|High", Default: "Low"},
	{Id: "parallel-n64-angrylion-multithread", Name: "parallel-n64-angrylion-multithread", Variants: "all threads|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30|31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46|47|48|49|50|51|52|53|54|55|56|57|58|59|60|61|62|63", Default: "all threads"},
	{Id: "parallel-n64-virefresh", Name: "parallel-n64-virefresh", Variants: "auto|1500|2200", Default: "auto"},
	{Id: "parallel-n64-framerate", Name: "parallel-n64-framerate", Variants: "original|fullspeed", Default: "original"},
	{Id: "parallel-n64-boot-device", Name: "parallel-n64-boot-device", Variants: "Default|64DD IPL", Default: "Default"},
}

var CoreOptionsBeetlePCE = []CoreOption{
	{Id: "pce_palette", Name: "pce_palette", Variants: "RGB|Composite", Default: "RGB"},
	{Id: "pce_psgrevision", Name: "pce_psgrevision", Variants: "HuC6280|HuC6280A", Default: "HuC6280A"},
	{Id: "pce_mouse_sensitivity", Name: "pce_mouse_sensitivity", Variants: "0.125|0.250|0.375|0.500|0.625|0.750|0.875|1.000|1.125|1.25|1.50|1.75|2.00|2.25|2.50|2.75|3.00|3.25|3.50|3.75|4.00|4.25|4.50|4.75|5.00", Default: "1.25"},
	{Id: "pce_default_joypad_type_p1", Name: "pce_default_joypad_type_p1", Variants: "2 Buttons|6 Buttons", Default: "2 Buttons"},
	{Id: "pce_default_joypad_type_p2", Name: "pce_default_joypad_type_p2", Variants: "2 Buttons|6 Buttons", Default: "2 Buttons"},
	{Id: "pce_default_joypad_type_p3", Name: "pce_default_joypad_type_p3", Variants: "2 Buttons|6 Buttons", Default: "2 Buttons"},
	{Id: "pce_default_joypad_type_p4", Name: "pce_default_joypad_type_p4", Variants: "2 Buttons|6 Buttons", Default: "2 Buttons"},
	{Id: "pce_default_joypad_type_p5", Name: "pce_default_joypad_type_p5", Variants: "2 Buttons|6 Buttons", Default: "2 Buttons"},
	{Id: "pce_Turbo_Delay", Name: "pce_Turbo_Delay", Variants: "Fast|Medium|Slow", Default: "Fast"},
	{Id: "pce_cdbios", Name: "pce_cdbios", Variants: "Games Express|System Card 1|System Card 2|System Card 3|System Card 2 US|System Card 3 US", Default: "System Card 3"},
	{Id: "pce_cdspeed", Name: "pce_cdspeed", Variants: "1|2|4|8", Default: "1"}, {Id: "pce_adpcmextraprec", Name: "pce_adpcmextraprec", Variants: "10-bit|12-bit", Default: "10-bit"},
	{Id: "pce_adpcmvolume", Name: "pce_adpcmvolume", Variants: "0|10|20|30|40|50|60|70|80|90|100|110|120|130|140|150|160|170|180|190|200", Default: "100"},
	{Id: "pce_cddavolume", Name: "pce_cddavolume", Variants: "0|10|20|30|40|50|60|70|80|90|100|110|120|130|140|150|160|170|180|190|200", Default: "100"},
	{Id: "pce_cdpsgvolume", Name: "pce_cdpsgvolume", Variants: "0|10|20|30|40|50|60|70|80|90|100|110|120|130|140|150|160|170|180|190|200", Default: "100"},
	{Id: "pce_ocmultiplier", Name: "pce_ocmultiplier", Variants: "1|2|3|4|5|6|7|8|9|10|20|30|40|50", Default: "1"},
}

var CoreOptionsBeetlePCFX = []CoreOption{
	{Id: "pcfx_high_dotclock_width", Name: "pcfx_high_dotclock_width", Variants: "256|341|1024", Default: "1024"},
	{Id: "pcfx_resamp_quality", Name: "pcfx_resamp_quality", Variants: "0|1|2|3|4|5", Default: "3"},
	{Id: "pcfx_mouse_sensitivity", Name: "pcfx_mouse_sensitivity", Variants: "1.00|1.25|1.50|1.75|2.00|2.25|2.50|2.75|3.00|3.25|3.50|3.75|4.00|4.25|4.50|4.75|5.00", Default: "1.25"},
}

var CoreOptionsBeetleNGP = []CoreOption{
	{Id: "ngp_language", Name: "ngp_language", Variants: "english|japanese", Default: "english"},
}

var CoreOptionsBeetleWSwan = []CoreOption{
	{Id: "wswan_frameskip_threshold", Name: "wswan_frameskip_threshold", Variants: "15|18|21|24|27|30|33|36|39|42|45|48|51|54|57|60", Default: "33"},
	{Id: "wswan_sound_sample_rate", Name: "wswan_sound_sample_rate", Variants: "11025|22050|44100|48000", Default: "44100"},
}

var CoreOptionsGearcoleco = []CoreOption{
	{Id: "gearcoleco_timing", Name: "gearcoleco_timing", Variants: "Auto|NTSC (60 Hz)|PAL (50 Hz)", Default: "Auto"},
}

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
	case "mupen64plus_next":
		return CoreOptionsMupen64Plus
	case "opera":
		return CoreOptionsOpera
	case "prosystem":
		return CoreOptionsProSystem
	case "stella2014":
		return CoreOptionsStella2014
	case "parallel_n64":
		return CoreOptionsParallelN64
	case "mednafen_pce":
		return CoreOptionsBeetlePCE
	case "mednafen_pcfx":
		return CoreOptionsBeetlePCFX
	case "mednafen_ngp":
		return CoreOptionsBeetleNGP
	case "mednafen_wswan":
		return CoreOptionsBeetleWSwan
	case "gearcoleco":
		return CoreOptionsGearcoleco
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

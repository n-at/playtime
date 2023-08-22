#!/bin/bash

cd assets

npm install

mkdir emulatorjs
git clone https://github.com/EmulatorJS/EmulatorJS _tmp
mv _tmp/data/* emulatorjs
rm -rf _tmp

#download BIOS

mkdir bios
cd bios

mkdir nes
wget -O "nes/disksys.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Famicom%20Disk%20System/disksys.rom"

mkdir snes
wget -O "snes/BS-X.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Satellaview/BS-X.bin"
wget -O "snes/STBIOS.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20SuFami%20Turbo/STBIOS.bin"

mkdir gb
wget -O "gb/gb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy/gb_bios.bin"
wget -O "gb/gbc_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy%20Color/gbc_bios.bin"

mkdir gba
wget -O "gba/gb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy/gb_bios.bin"
wget -O "gba/gbc_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy%20Color/gbc_bios.bin"
wget -O "gba/gba_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Game%20Boy%20Advance/gba_bios.bin"
wget -O "gba/sgb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Super%20Game%20Boy/sgb_bios.bin"

mkdir nds
wget -O "nds/bios7.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios7.bin"
wget -O "nds/bios9.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios9.bin"
wget -O "nds/firmware.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/firmware.bin"

mkdir psx
wget -O "psx/scph5500.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5500.bin"
wget -O "psx/scph5501.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5501.bin"
wget -O "psx/scph5502.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5502.bin"
wget -O "psx/PSXONPSP660.BIN" "https://github.com/Abdess/retroarch_system/raw/Other/Sony%20-%20PlayStation/PSXONPSP660.BIN"
wget -O "psx/scph101.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph101.bin"
wget -O "psx/scph7001.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph7001.bin"
wget -O "psx/scph1001.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph1001.bin"

mkdir lynx
wget -O "lynx/lynxboot.img" "https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%20Lynx/lynxboot.img"

mkdir segaSaturn
wget -O "segaSaturn/saturn_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Saturn/saturn_bios.bin"

mkdir segaMS
wget -O "segaMS/bios_E.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_E.sms"
wget -O "segaMS/bios_U.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_U.sms"
wget -O "segaMS/bios_J.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_J.sms"

mkdir segaMD
wget -O "segaMD/bios_MD.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20Drive%20-%20Genesis/bios_MD.bin"

mkdir segaGG
wget -O "segaGG/bios.gg" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Game%20Gear/bios.gg"

mkdir segaCD
wget -O "segaCD/bios_CD_E.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_E.bin"
wget -O "segaCD/bios_CD_U.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_U.bin"
wget -O "segaCD/bios_CD_J.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_J.bin"

mkdir 3do
wget -O "3do/panafz1.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1.bin"
wget -O "3do/panafz10.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10.bin"
wget -O "3do/panafz10-norsa.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10-norsa.bin"
wget -O "3do/panafz10e-anvil.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10e-anvil.bin"
wget -O "3do/panafz10e-anvil-norsa.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10e-anvil-norsa.bin"
wget -O "3do/panafz1j.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1j.bin"
wget -O "3do/panafz1j-norsa.bi" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1j-norsa.bin"
wget -O "3do/goldstar.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/goldstar.bin"
wget -O "3do/sanyotry.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/sanyotry.bin"
wget -O "3do/3do_arcade_saot.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/3do_arcade_saot.bin"

mkdir atari7800
wget -O "atari7800/7800_BIOS_U.rom" 'https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%207800/7800%20BIOS%20(U).rom'

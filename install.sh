#!/bin/bash

cd assets

npm install

#download EmulatorJS

git clone https://github.com/EmulatorJS/EmulatorJS _tmp
cd _tmp
git checkout "806623af696f625618f9a073739cb1b6e4742e7c"
cd data
mkdir cores
cd minify
npm install
node index.js
rm -rf node_modules
cd ../../..
mkdir emulatorjs
mv _tmp/data/* emulatorjs
rm -rf _tmp

cd emulatorjs/cores
wget "https://cdn.emulatorjs.org/latest/data/cores/a5200-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/beetle_vb-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/desmume-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/desmume2015-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/fbalpha2012_cps1-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/fbalpha2012_cps2-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/fbneo-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/fceumm-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/gambatte-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/gearcoleco-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/genesis_plus_gx-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/handy-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mame2003_plus-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mame2003-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_ngp-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_pce-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_pcfx-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_psx_hw-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_psx_hw-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mednafen_wswan-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/melonds-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/melonds-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mgba-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mgba-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mupen64plus_next-legacy-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mupen64plus_next-thread-legacy-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mupen64plus_next-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/mupen64plus_next-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/nestopia-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/opera-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/opera-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/parallel_n64-legacy-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/parallel_n64-thread-legacy-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/parallel_n64-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/parallel_n64-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/pcsx_rearmed-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/pcsx_rearmed-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/picodrive-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/prosystem-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/puae-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/snes9x-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/stella2014-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/vice_x64-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/virtualjaguar-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/yabause-thread-wasm.data"
wget "https://cdn.emulatorjs.org/latest/data/cores/yabause-wasm.data"
cd ../..

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

cd gb
zip gb.zip *.bin
rm *.bin
cd ..

mkdir gba
wget -O "gba/gb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy/gb_bios.bin"
wget -O "gba/gbc_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy%20Color/gbc_bios.bin"
wget -O "gba/gba_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Game%20Boy%20Advance/gba_bios.bin"
wget -O "gba/sgb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Super%20Game%20Boy/sgb_bios.bin"

cd gba
zip gba.zip *.bin
rm *.bin
cd ..

mkdir nds
wget -O "nds/bios7.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios7.bin"
wget -O "nds/bios9.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios9.bin"
wget -O "nds/firmware.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/firmware.bin"

cd nds
zip ds.zip firmware.bin bios7.bin bios9.bin
rm firmware.bin bios7.bin bios9.bin
cd ..

mkdir psx
wget -O "psx/scph5500.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5500.bin"
wget -O "psx/scph5501.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5501.bin"
wget -O "psx/scph5502.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph5502.bin"
wget -O "psx/scph101.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph101.bin"
wget -O "psx/scph7001.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph7001.bin"
wget -O "psx/scph1001.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sony%20-%20PlayStation/scph1001.bin"
wget -O "psx/psp.bin" "https://github.com/Abdess/retroarch_system/raw/Other/Sony%20-%20PlayStation/PSXONPSP660.BIN"

cd psx
mkdir tmp
cp "psp.bin" "tmp/scph5500.bin"
cp "psp.bin" "tmp/scph5501.bin"
cp "psp.bin" "tmp/scph5502.bin"
cd tmp
zip ../psp.zip *.bin
cd ..
zip psx.zip scph5500.bin scph5501.bin scph5502.bin scph101.bin scph7001.bin scph1001.bin
rm -rf tmp *.bin
cd ..

mkdir lynx
wget -O "lynx/lynxboot.img" "https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%20Lynx/lynxboot.img"

mkdir segaSaturn
wget -O "segaSaturn/saturn_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Saturn/saturn_bios.bin"

mkdir segaMS
wget -O "segaMS/bios_E.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_E.sms"
wget -O "segaMS/bios_U.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_U.sms"
wget -O "segaMS/bios_J.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_J.sms"

cd segaMS
zip segaMS.zip *.sms
rm *.sms
cd ..

mkdir segaMD
wget -O "segaMD/bios_MD.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20Drive%20-%20Genesis/bios_MD.bin"

mkdir segaGG
wget -O "segaGG/bios.gg" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Game%20Gear/bios.gg"

mkdir segaCD
wget -O "segaCD/bios_CD_E.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_E.bin"
wget -O "segaCD/bios_CD_U.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_U.bin"
wget -O "segaCD/bios_CD_J.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_J.bin"

cd segaCD
zip segaCD.zip *.bin
rm *.bin
cd ..

mkdir 3do
wget -O "3do/panafz1.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1.bin"
wget -O "3do/panafz10.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10.bin"
wget -O "3do/panafz10-norsa.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10-norsa.bin"
wget -O "3do/panafz10e-anvil.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10e-anvil.bin"
wget -O "3do/panafz10e-anvil-norsa.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz10e-anvil-norsa.bin"
wget -O "3do/panafz1j.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1j.bin"
wget -O "3do/panafz1j-norsa.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/panafz1j-norsa.bin"
wget -O "3do/goldstar.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/goldstar.bin"
wget -O "3do/sanyotry.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/sanyotry.bin"
wget -O "3do/3do_arcade_saot.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/3DO%20Company%2C%20The%20-%203DO/3do_arcade_saot.bin"

cd 3do
zip 3do.zip *.bin
rm *.bin
cd ..

mkdir atari7800
wget -O "atari7800/7800_BIOS_U.rom" 'https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%207800/7800%20BIOS%20(U).rom'

mkdir pce
wget -O "pce/syscard3.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard3.pce"
wget -O "pce/syscard2.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard2.pce"
wget -O "pce/syscard1.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard1.pce"
wget -O "pce/gexpress.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/gexpress.pce"

cd pce
zip pce.zip *.pce
rm *.pce
cd ..

mkdir coleco
wget -O "coleco/colecovision.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/Coleco%20-%20ColecoVision/colecovision.rom"

mkdir pcfx
wget -O "pcfx/pcfx.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfx.rom"
wget -O "pcfx/pcfxbios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfxbios.bin"
wget -O "pcfx/pcfxv101.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfxv101.bin"

cd pcfx
zip pcfx.zip pcfx.rom pcfxbios.bin pcfxv101.bin
rm pcfx.rom pcfxbios.bin pcfxv101.bin
cd ..

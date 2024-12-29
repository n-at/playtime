#!/bin/bash

PLAYTIME_EJS_REPO_URL=${PLAYTIME_EJS_REPO_URL:-"https://github.com/EmulatorJS/EmulatorJS"}
PLAYTIME_EJS_REVISION=${PLAYTIME_EJS_REVISION:-"be29680d57015482612f3b492b7952455494be5a"}
PLAYTIME_EJS_CORES_URL=${PLAYTIME_EJS_CORES_URL:-"https://cdn.emulatorjs.org/stable/data/cores"}

#

cd assets

npm install --no-fund --ignore-scripts

#download EmulatorJS

git clone "${PLAYTIME_EJS_REPO_URL}" _tmp
cd _tmp
git checkout "${PLAYTIME_EJS_REVISION}"
cd data/minify
npm install --ignore-scripts
node index.js
rm -rf node_modules
cd ../../..
mkdir emulatorjs
mv _tmp/data/* emulatorjs
rm -rf _tmp

cd emulatorjs
mkdir cores
cd cores
wget "${PLAYTIME_EJS_CORES_URL}/81-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/81-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/81-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/81-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/a5200-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/a5200-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/a5200-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/a5200-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/beetle_vb-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/beetle_vb-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/beetle_vb-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/beetle_vb-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/cap32-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/cap32-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/cap32-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/cap32-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/crocods-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/crocods-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/crocods-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/crocods-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume2015-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume2015-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume2015-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/desmume2015-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps1-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps1-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps1-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps1-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps2-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps2-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps2-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbalpha2012_cps2-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbneo-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbneo-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbneo-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fbneo-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fceumm-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fceumm-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fceumm-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fceumm-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fuse-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fuse-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fuse-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/fuse-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gambatte-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gambatte-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gambatte-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gambatte-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gearcoleco-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gearcoleco-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gearcoleco-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/gearcoleco-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/genesis_plus_gx-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/genesis_plus_gx-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/genesis_plus_gx-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/genesis_plus_gx-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/handy-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/handy-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/handy-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/handy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003_plus-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003_plus-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003_plus-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003_plus-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mame2003-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_ngp-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_ngp-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_ngp-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_ngp-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pce-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pce-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pce-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pce-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pcfx-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pcfx-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pcfx-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_pcfx-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_psx_hw-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_psx_hw-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_psx_hw-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_psx_hw-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_wswan-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_wswan-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_wswan-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mednafen_wswan-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/melonds-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/melonds-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/melonds-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/melonds-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mgba-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mgba-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mgba-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mgba-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mupen64plus_next-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mupen64plus_next-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mupen64plus_next-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/mupen64plus_next-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/nestopia-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/nestopia-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/nestopia-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/nestopia-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/opera-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/opera-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/opera-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/opera-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/parallel_n64-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/parallel_n64-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/parallel_n64-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/parallel_n64-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/pcsx_rearmed-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/pcsx_rearmed-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/pcsx_rearmed-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/pcsx_rearmed-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/picodrive-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/picodrive-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/picodrive-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/picodrive-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/ppsspp-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/ppsspp-assets.zip"
wget "${PLAYTIME_EJS_CORES_URL}/prboom-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prboom-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prboom-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prboom-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prosystem-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prosystem-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prosystem-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/prosystem-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/puae-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/puae-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/puae-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/puae-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/smsplus-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/smsplus-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/smsplus-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/smsplus-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/snes9x-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/snes9x-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/snes9x-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/snes9x-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/stella2014-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/stella2014-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/stella2014-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/stella2014-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x128-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x128-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x128-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x128-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64sc-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64sc-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64sc-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_x64sc-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xpet-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xpet-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xpet-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xpet-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xplus4-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xplus4-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xplus4-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xplus4-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xvic-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xvic-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xvic-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/vice_xvic-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/virtualjaguar-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/virtualjaguar-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/virtualjaguar-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/virtualjaguar-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/yabause-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/yabause-thread-legacy-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/yabause-thread-wasm.data"
wget "${PLAYTIME_EJS_CORES_URL}/yabause-wasm.data"
mkdir reports && cd reports
wget "${PLAYTIME_EJS_CORES_URL}/reports/81.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/a5200.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/beetle_vb.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/cap32.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/crocods.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/desmume.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/desmume2015.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/fbalpha2012_cps1.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/fbalpha2012_cps2.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/fbneo.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/fceumm.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/fuse.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/gambatte.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/gearcoleco.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/genesis_plus_gx.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/handy.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mame2003_plus.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mame2003.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mednafen_ngp.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mednafen_pce.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mednafen_pcfx.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mednafen_psx_hw.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mednafen_wswan.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/melonds.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mgba.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/mupen64plus_next.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/nestopia.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/opera.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/parallel_n64.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/pcsx_rearmed.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/picodrive.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/ppsspp.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/prboom.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/prosystem.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/puae.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/smsplus.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/snes9x.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/stella2014.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_x128.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_x64.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_x64sc.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_xpet.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_xplus4.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/vice_xvic.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/virtualjaguar.json"
wget "${PLAYTIME_EJS_CORES_URL}/reports/yabause.json"
cd ../../..

#download BIOS

mkdir bios
cd bios

#

mkdir nes
wget -O "nes/disksys.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Famicom%20Disk%20System/disksys.rom"

#

mkdir snes
wget -O "snes/BS-X.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Satellaview/BS-X.bin"
wget -O "snes/STBIOS.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20SuFami%20Turbo/STBIOS.bin"

#

mkdir gb
wget -O "gb/gb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy/gb_bios.bin"
wget -O "gb/gbc_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy%20Color/gbc_bios.bin"

cd gb
zip gb.zip *.bin
rm *.bin
cd ..

#

mkdir gba
wget -O "gba/gb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy/gb_bios.bin"
wget -O "gba/gbc_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Gameboy%20Color/gbc_bios.bin"
wget -O "gba/gba_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Game%20Boy%20Advance/gba_bios.bin"
wget -O "gba/sgb_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Super%20Game%20Boy/sgb_bios.bin"

cd gba
zip gba.zip *.bin
rm *.bin
cd ..

#

mkdir nds
wget -O "nds/bios7.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios7.bin"
wget -O "nds/bios9.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/bios9.bin"
wget -O "nds/firmware.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Nintendo%20-%20Nintendo%20DS/firmware.bin"

cd nds
zip ds.zip firmware.bin bios7.bin bios9.bin
rm firmware.bin bios7.bin bios9.bin
cd ..

#

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

#

mkdir lynx
wget -O "lynx/lynxboot.img" "https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%20Lynx/lynxboot.img"

#

mkdir segaSaturn
wget -O "segaSaturn/saturn_bios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Saturn/saturn_bios.bin"

#

mkdir segaMS
wget -O "segaMS/bios_E.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_E.sms"
wget -O "segaMS/bios_U.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_U.sms"
wget -O "segaMS/bios_J.sms" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Master%20System%20-%20Mark%20III/bios_J.sms"

cd segaMS
zip segaMS.zip *.sms
rm *.sms
cd ..

#

mkdir segaMD
wget -O "segaMD/bios_MD.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20Drive%20-%20Genesis/bios_MD.bin"

#

mkdir segaGG
wget -O "segaGG/bios.gg" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Game%20Gear/bios.gg"

#

mkdir segaCD
wget -O "segaCD/bios_CD_E.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_E.bin"
wget -O "segaCD/bios_CD_U.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_U.bin"
wget -O "segaCD/bios_CD_J.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_J.bin"

cd segaCD
zip segaCD.zip *.bin
rm *.bin
cd ..

#

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

#

mkdir atari7800
wget -O "atari7800/7800_BIOS_U.rom" 'https://github.com/Abdess/retroarch_system/raw/libretro/Atari%20-%207800/7800%20BIOS%20(U).rom'

#

mkdir pce
wget -O "pce/syscard3.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard3.pce"
wget -O "pce/syscard2.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard2.pce"
wget -O "pce/syscard1.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/syscard1.pce"
wget -O "pce/gexpress.pce" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC%20Engine%20-%20TurboGrafx%2016%20-%20SuperGrafx/gexpress.pce"

cd pce
zip pce.zip *.pce
rm *.pce
cd ..

#

mkdir coleco
wget -O "coleco/colecovision.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/Coleco%20-%20ColecoVision/colecovision.rom"

#

mkdir pcfx
wget -O "pcfx/pcfx.rom" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfx.rom"
wget -O "pcfx/pcfxbios.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfxbios.bin"
wget -O "pcfx/pcfxv101.bin" "https://github.com/Abdess/retroarch_system/raw/libretro/NEC%20-%20PC-FX/pcfxv101.bin"

cd pcfx
zip pcfx.zip pcfx.rom pcfxbios.bin pcfxv101.bin
rm pcfx.rom pcfxbios.bin pcfxv101.bin
cd ..

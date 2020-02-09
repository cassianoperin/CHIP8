# CHIP-8 / SCHIP Emulator

CHIP-8 / SCHIP Emulator writen in GO with simple code to be easy to be studied and understood.

<img width="430" alt="invaders" src="https://github.com/cassianoperin/CHIP-8_GO/blob/master/images/invaders.png">

## Features
* Pause and resume emulation
* Reset emulation
* Step Forward CPU Cycles for Debug
* Step Back (Rewind) CPU Cycles for Debug
* Online Debug mode

## Requirements
* GO
* go get github.com/faiface/pixel/pixelgl
* go get github.com/faiface/beep
* go get github.com/faiface/beep/mp3
* go get github.com/faiface/beep/speaker

## Usage

1. Run:
	`$ go run chip8.go ROM_NAME`

2. Keys
- Original COSMAC Keyboard Layout:

	`1` `2` `3` `C`

	`4` `5` `6` `D`

	`7` `8` `9` `E`

	`A` `0` `B` `F`

- **Keys used in this emulator:**

	`1` `2` `3` `4`

	`Q` `W` `E` `R`

	`A` `S` `D` `F`

	`Z` `X` `C` `V`

	`P`: Pause and Resume emulation

	`[`: Step back (rewind) one CPU cycle **in Pause Mode** (for debug and study purposes)

	`]`: Step forward one CPU cycle in **Pause Mode** (for debug and study purposes)

	`9`: Enable / Disable Debug Mode

	`0`: Reset

	`ESC`: Exit emulator


## Documentation
[Cowgod's Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#0.0)

[How to write an emulator (CHIP-8 interpreter) — Multigesture.net](http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/)

[Wikipedia - CHIP-8](https://en.wikipedia.org/wiki/CHIP-8)

[HP48 Superchip](https://github.com/Chromatophore/HP48-Superchip)

[SCHIP](http://devernay.free.fr/hacks/chip8/schip.txt)

[trapexit chip-8 documentation](https://github.com/trapexit/chip-8_documentation)

[CHIP‐8-Extensions-Reference](https://github.com/mattmikolay/chip-8/wiki/CHIP%E2%80%908-Extensions-Reference)




## TODO LIST

1. CHIP8 - Equalize game speed (some games runs too fast, other slow)
2. CHIP8 - Key pressing cause slowness
3. CHIP8 - Improve draw method (Rewrite graphics mode to just draw the differences from each frame)
4. CHIP8 - Implement a correct 60 FPS control
5. SCHIP - Document README with now supported SCHIP games
6. ALL - Test on Windows and Linux
7. ALL - Rewind mode make emulation slow due to arrays and graphics processing
8. SCHIP - Emulation really slow at this moment
9. SCHIP - IMPLEMENT different handling opcodes in schip mode (and chip modern games)
10. Map and identify these modern chip8 games that uses schip opcodes 

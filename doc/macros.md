# Macros for FPGA Synthesis

Macro                    | Target  |  Usage
-------------------------|---------|----------------------
JTFRAME_180SHIFT         | MiSTer  | Use DDIO cell instead of PLL to create the SDRAM phase shift
JTFRAME_4PLAYERS         |         | Extra inputs for 4 players
JTFRAME_ANALOG           |         | Enables analog sticks
JTFRAME_ANALOG_DUAL      |         | Enables second analog stick, requires JTFRAME_ANALOG too
JTFRAME_ARX              | MiSTer  | Defines aspect ratio
JTFRAME_ARY              | MiSTer  | Defines aspect ratio
JTFRAME_AUTOFIRE0        | MiSTer  | Button 0 will autofire when pressed ([OSD](osd.md) option)
JTFRAME_AVATARS          |         | Enables avatars on credits screen
JTFRAME_BUTTONS          |         | Sets the number of action buttons used (2 by default)
JTFRAME_CLK24            |         | Adds an additional clock input
JTFRAME_CLK48            |         | Adds an additional clock input
JTFRAME_CLK6             |         | Adds an additional clock input
JTFRAME_CLK96            |         | Adds an additional clock input
JTFRAME_CHEAT            |         | Enables the [cheat engine](cheat.md)
JTFRAME_CHEAT_SCRAMBLE   |         | Enables cheat firmware encryption
JTFRAME_DEBUG            |         | Enables the debug_bus signal connection to the game instance
JTFRAME_DONTSIM_SCAN2X   |         | Internal. Do not define externally
JTFRAME_DUAL_RAM_DUMP    |         | Enables dumping of RAM contents in simulation
JTFRAME_DWNLD_PROM_ONLY  |         | Quick download sim with only PROM contents
JTFRAME_FLIP_RESET       |         | Varying the flip DIP setting causes a reset
JTFRAME_HEIGHT           |         | Sets the video height
JTFRAME_INTERLACED       |         | Support for interlaced games
JTFRAME_INTERPOL2        |         | Enables a x2 FIR interpolation filter for audio. The game 'sample' output must be well defined
JTFRAME_IOCTL_RD         |         | Enables saving to SD card via NVRAM interface. Set it to the number of bytes to save on MiST. Any value will work for MiSTer
JTFRAME_LITE_KEYBOARD    |         | Disables automatic MAME keys mapping
JTFRAME_MIST_DIPBASE     | MiST    | Starting base in status word for MiST dip switches
JTFRAME_MIST_DIRECT      | MiST    | On by default. Define as 0 to disable. Fast ROM load
JTFRAME_MR_FASTIO        | MiSTer  | 16-bit ROM load in MiSTer. Set by default if CLK96 is set
JTFRAME_MR_DDRLOAD       | MiSTer  | ROM download process uses the DDR as proxy
JTFRAME_MR_DDR           | MiSTer  | Defined internally. Do not define manually.
JTFRAME_MR_ROTCCW        | MiSTer  | Rotate the frame counter clockwise
JTFRAME_MRA_DIP          |         | DIPs are in an MRA file
JTFRAME_NOHOLDBUS        |         | Reduces bus noise (non-interleaved SDRAM controller)
JTFRAME_NOHQ2X           | MiSTer  | Disables HQ2X filter in MiSTer
JTFRAME_NO_DB15          | MiSTer  | Disables DB15 controller modules
JTFRAME_NO_DB15_OSD      | MiSTer  | Disables OSD control via DB15 controller
JTFRAME_OSD_FLIP         |         | flip option on OSD
JTFRAME_OSD_NOCREDITS    |         | No credits option on OSD
JTFRAME_OSD_NOLOAD       |         | No load option on OSD
JTFRAME_OSD_NOLOGO       |         | Disables the JT logo as OSD background
JTFRAME_OSD_SND_EN       |         | OSD option to enable/disable FX and FM channels
JTFRAME_OSD_TEST         |         | Test option on OSD
JTFRAME_OSD_VOL          |         | Show FX volume control on OSD
JTFRAME_OSDCOLOR         |         | Sets the OSD colour
JTFRAME_PLL              |         | PLL module name to be used. Defaults to jtframe_pll0
JTFRAME_RELEASE          |         | Disables debug control via keyboard
JTFRAME_ROTATE           |         | Enables more rotate options in the OSD
JTFRAME_SCAN2X_NOBLEND   | MiST    | Disables pixel blending
JTFRAME_SDRAM96          |         | SDRAM is clocked at 96MHz and the clk input of game is 96MHz
JTFRAME_SDRAM_BANKS      |         | Game module ports will support interleaved bank access
JTFRAME_SHADOW           | MiSTer  | Start address for SDRAM shadowing and dump as NVRAM
JTFRAME_SHADOW_LEN       | MiSTer  | Length in bits of the shadowing. See [sdram.md](sdram.md)
JTFRAME_SIGNED_SND       | MiSTer  | Set to 0 if the game only uses unsigned sound sources
JTFRAME_STATUS           |         | Game module will receive an 8-bit address and can output 8-bit data in response
JTFRAME_SUPPORT_4WAY     |         | Enables support for 4-way joysticks if the MRA sets it
JTFRAME_VERTICAL         |         | Enables support for vertical games
JTFRAME_WIDTH            |         | Sets the video width

# Core-specific OSD Items

Macro                    | Target  |  Usage
-------------------------|---------|----------------------
CORE_OSD                 |         | Adds an option to the OSD

Example from the JTCPS core:

`CORE_OSD=O5,Turbo,Off,On;`

Follow the character coding documented in [osd.md](osd.md)


# Device Selection

The wrappers jtframe_m68k and jtframe_z80 offer an uniform interface for
different underlying modules.

Macro                    | Target  |  Usage
-------------------------|---------|----------------------
JTFRAME_J68              |         | Selects J68_CPU as M68000 module (default fx68k)
VHDLZ80                  |         | Selects VHDL version of T80 core (default for synthesis)
TV80S                    |         | Selects verilog version of T80 core

# SDRAM Banks

Macro                    | Target  |  Usage
-------------------------|---------|----------------------
JTFRAME_SDRAM_ADQM       | MiSTer  | A12 and A11 are equal to DQMH/L
JTFRAME_SDRAM_BWAIT      |         | Adds a wait cycle in the SDRAM
JTFRAME_SDRAM_CHECK      |         | Double check SDRAM data through modules (slow)
JTFRAME_SDRAM_DEBUG      |         | Outputs debug messages for SDRAM during simulation
JTFRAME_SDRAM_LARGE      | MiSTer  | Enables 64MB access to SDRAM modules
JTFRAME_SDRAM_MUXLATCH   |         | Extra latch for SDRAM mux for <64MHz operation
JTFRAME_SDRAM_NO_DWNRFSH |         | No refresh during download (non-interleaved SDRAM controller)
JTFRAME_SDRAM_REPACK     |         | Extra latch stage at SDRAM mux output

# SDRAM64

Macro                    | Target  |  Usage
-------------------------|---------|----------------------
JTFRAME_BAx_AUTOPRECH    |         | Enables auto precharge on bank X (0,1,2,3)
JTFRAME_BAx_LEN          |         | Sets length of bank x, valid values 16, 32 or 64

# Simulation-only Macros

The following macros only have an effect if SIMULATION is defined.

Macro                    | Target  |  Usage
-------------------------|---------|---------------------------------------------
DUMP_VIDEO               |         | Enables video dump to a file
DUMP_VIDEO_FNAME         |         | Internal. Do not assign.
JTFRAME_SAVESDRAM        |         | Saves SDRAM contents at the end of each frame (slow)
JTFRAME_SDRAM_STATS      |         | Produce SDRAM usage data during simulation
JTFRAME_SIM_DIPS         |         | Define DIP switch values during simulation
JTFRAME_SIM_ROMRQ_NOCHECK|         | Disable protocol checking of romrq
JTFRAME_SIM_SCAN2X       |         | Enables scan doubler simulation
JTFRAME_SIM_GFXEN        |         | Sets the gfx_en value (4 bits). See [debug.md](debug.md)
SIMULATION               |         | Enables simulation features
VIDEO_START              |         | First frame for which video output is provided use it to prevent a split first frame

# ROM Downloading

Macro                    | Target  |  Usage
-------------------------|---------|---------------------------------------------
LOADROM                  |         | Sends ROM data via serial interface
JTFRAME_DWNLD_PROM_ONLY  |         | Skip the regular download and go directly to the PROM section
JTFRAME_SIM_LOAD_EXTRA   |         | Extra wait time when transferring ROM in simulation

# Credits

JTFRAME_CREDITS is always enabled if JTFRAME_CHEAT is defined

Macro                    | Target  |  Usage
-------------------------|---------|---------------------------------------------
JTFRAME_CREDITS          |         | Adds credits screen
JTFRAME_CREDITS_AON      |         | credits screen is always on
JTFRAME_CREDITS_HSTART   |         | Horizontal offset for the 256-pxl wide credits
JTFRAME_CREDITS_PAGES    |         | number of pages of credits text
JTFRAME_CREDITS_NOROTATE |         | Always display the credits horizontally

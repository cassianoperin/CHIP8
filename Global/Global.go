// Global Variables used to to avoid circular dependencies
package Global

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

// Fullscreen / Video Modes
type Setting struct {
	Mode	*pixelgl.VideoMode
	Monitor	*pixelgl.Monitor
}

var (
	// ------------------------ Global Variables ------------------------ //
	Game_signature		string	= ""	// Game Signature (identify games that needs legacy opcodes)

	// ----------------------- Graphics Variables ----------------------- //
	Win			*pixelgl.Window
	WindowTitle		string = "Chip-8"
	Color_theme		= 2
	// Fullscreen / Video Modes
	Texts			[]*text.Text
	StaticText		*text.Text
	Settings		[]Setting
	ActiveSetting		*Setting
	IsFullScreen		= false		// Fullscrenn flag
	ResolutionCounter	int = 0		// Index of the available video resolution supported
	// FPS
	ShowFPS			bool		// Show or hide FPS counter flag
	// On screen messages
	ShowMessage		bool
	TextMessageStr		string
	// Draw operation executed, sinalize Graphics to update the screen
	DrawFlag		bool = false	// True if the screen must be drawn
	// Draw Mode
	// True  = Refresh screen (draw) every time DrawFlag is set
	// False = Refresh screen at 60Hz
	OriginalDrawMode	= false
	DrawModeMessage	string = ""
	// Input Commands that need a Draw
	InputDrawFlag		bool		// Force draw, necessary in some emulator rewind and forward status
	// Screen Size
	SizeX			float64		// Number of Columns in Graphics
	SizeY			float64		// Number of Lines in Graphics
	SizeYused		float64	= 1.0	// Percentage of the Screen Heigh used by the emulator // 1.0 = 100%, 0.0 = 0%
	// Monitor Size (to center Window)
	MonitorWidth	float64
	MonitorHeight	float64

	// ----------------------- SaveStates Variables ----------------------- //
	SavestateFolder		string = "Savestates"

	// ------------------------ Sound Variables ------------------------- //
	SpeakerPlaying		bool = false
	SpeakerStopped		bool = false

	// ---------------------------- Hybrids ----------------------------- //
	Hybrid_ETI_660_HW	bool = false

)

// Center Window Function
func CenterWindow() {
	winPos := Win.GetPos()
	winPos.X = (MonitorWidth  - float64(ActiveSetting.Mode.Width) ) / 2
	winPos.Y = (MonitorHeight - float64(ActiveSetting.Mode.Height) ) / 2
	Win.SetPos(winPos)
}

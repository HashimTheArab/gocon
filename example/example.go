package example

import (
	"fmt"
	"github.com/hashimthearab/gocon"
)

// Settings is the config structure for our program.
// This will hold values that we want to persist across restarts.
type Settings struct {
	Theme    string
	DarkMode bool
}

// Path ...
func (Settings) Path() string {
	// You can save it to your systems config directory, or simply return "settings.json"
	// to save it in the current directory.
	// dir, _ := os.UserConfigDir()
	// return filepath.Join(dir, "settings.json")
	return "settings.json"
}

func main() {
	// Load the settings
	settings, err := gocon.Load[Settings]()
	if err != nil {
		// If the file fails to load, or it was not found, an error will be returned.
		panic(err)
	}

	// Print the saved theme.
	fmt.Println(settings.Theme)

	// Save the new settings.
	gocon.Save(Settings{Theme: "blue", DarkMode: true})
}

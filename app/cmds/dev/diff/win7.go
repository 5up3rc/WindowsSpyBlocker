package diff

import (
	"github.com/crazy-max/WindowsSpyBlocker/app/menu"
	"github.com/crazy-max/WindowsSpyBlocker/app/utils/data"
	"github.com/fatih/color"
)

func menuWin7(args ...string) (err error) {
	menuCommands := []menu.CommandOption{
		{
			Description: "All",
			Color:       color.FgHiYellow,
			Function:    allWin7,
		},
		{
			Description: "Proxifier",
			Color:       color.FgHiYellow,
			Function:    proxifierWin7,
		},
		{
			Description: "Sysmon",
			Color:       color.FgHiYellow,
			Function:    sysmonWin7,
		},
		{
			Description: "Wireshark",
			Color:       color.FgHiYellow,
			Function:    wiresharkWin7,
		},
	}

	menuOptions := menu.NewOptions("Diff for Windows 7", "'menu' for help [dev-diff-win7]> ", 0, "")

	menuN := menu.NewMenu(menuCommands, menuOptions)
	menuN.Start()
	return
}

func allWin7(args ...string) error {
	all(data.OS_WIN7)
	return nil
}

func proxifierWin7(args ...string) error {
	prog(data.OS_WIN7, "proxifier")
	return nil
}

func sysmonWin7(args ...string) error {
	prog(data.OS_WIN7, "sysmon")
	return nil
}

func wiresharkWin7(args ...string) error {
	prog(data.OS_WIN7, "sysmon")
	return nil
}

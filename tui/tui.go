package tui

type TUI interface {
	Show()
	Result() uint8
	Refresh()
}

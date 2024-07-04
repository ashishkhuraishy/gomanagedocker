package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type navigationKeymap struct {
	Enter    key.Binding
	Back     key.Binding
	Quit     key.Binding
	NextTab  key.Binding
	PrevTab  key.Binding
	NextItem key.Binding
	PrevItem key.Binding
	PrevPage key.Binding
	NextPage key.Binding
}

type imgKeymap struct {
	Run    key.Binding
	Rename key.Binding
	Scout  key.Binding
	// Pull        key.Binding
	Prune       key.Binding
	Delete      key.Binding
	DeleteForce key.Binding
	CopyId      key.Binding
	RunAndExec  key.Binding
}

type contKeymap struct {
	ToggleListAll   key.Binding
	ToggleStartStop key.Binding
	TogglePause     key.Binding
	Restart         key.Binding
	Delete          key.Binding
	DeleteForce     key.Binding
	Exec            key.Binding
	Prune           key.Binding
	CopyId          key.Binding
}

type volKeymap struct {
	Delete key.Binding
	Prune  key.Binding
	CopyId key.Binding
}

var ImageKeymap = imgKeymap{
	Run: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "run"),
	),
	Rename: key.NewBinding(
		key.WithKeys("R"),
		key.WithHelp("r", "rename"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	DeleteForce: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete (force)"),
	),
	Scout: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Scout"),
	),
	// Pull: key.NewBinding(
	// 	key.WithKeys("o"),
	// 	key.WithHelp("o", "Pull new Image"),
	// ),
	Prune: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "Prune images"),
	),
	CopyId: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Copy Image ID"),
	),
	RunAndExec: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "Run and exec"),
	),
}

func (m imgKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{m.Run,
			m.Delete,
			m.DeleteForce,
			m.Prune},
	}
}

func (m imgKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		m.Run,
		m.Delete,
		m.DeleteForce,
		m.Prune,
		m.Scout,
		m.CopyId,
		m.RunAndExec,
	}

}

var ContainerKeymap = contKeymap{
	ToggleListAll: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Toggle list all"),
	),
	ToggleStartStop: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Toggle Start/Stop"),
	),
	TogglePause: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "Toggle Pause/unPause"),
	),
	Restart: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Restart"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	DeleteForce: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete (force)"),
	),
	Prune: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune"),
	),
	Exec: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "exec"),
	),
	CopyId: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Copy Container ID"),
	),
}

func (m contKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (m contKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		m.ToggleListAll,
		m.ToggleStartStop,
		m.Restart,
		m.TogglePause,
		m.Delete,
		m.DeleteForce,
		m.Prune,
		m.Exec,
		m.CopyId,
	}
}

var VolumeKeymap = volKeymap{
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Prune: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune"),
	),
	CopyId: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Copy Name"),
	),
}

func (m volKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (m volKeymap) ShortHelp() []key.Binding {
	return []key.Binding{m.Delete, m.Prune, m.CopyId}
}

var NavKeymap = navigationKeymap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("right", "l", "tab"),
		key.WithHelp("->/l/tab", "next"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("left", "h", "shift+tab"),
		key.WithHelp("<-/h/shift+tab", "prev"),
	),
	NextItem: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "next item"),
	),
	PrevItem: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/h", "prev item"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("["),
		key.WithHelp("[", "prev page"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("]"),
		key.WithHelp("]", "next page"),
	),
}

func (m navigationKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (m navigationKeymap) ShortHelp() []key.Binding {
	return []key.Binding{m.NextItem, m.PrevItem, m.NextTab, m.PrevTab, m.PrevPage, m.NextPage, m.Enter, m.Quit}
}

package main_menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var state common.SessionState

		if i, ok := m.SelectedItem().(MenuItem); ok {
			state = i.state
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				var cmds []tea.Cmd
				cmds = append(cmds, func() tea.Msg {
					return common.SwitchToStateMsg{State: state, ConstructNew: true}
				})
				return tea.Batch(cmds...)
			}
		}

		return nil
	}

	d.ShortHelpFunc = keys.ShortHelp
	d.FullHelpFunc = keys.FullHelp

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

// ShortHelp Additional short help entries. This satisfies
// the help.KeyMap interface and is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{d.choose}
}

// FullHelp - Additional full help entries. This satisfies
// the help.KeyMap interface and is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{d.choose}}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter", " "),
			key.WithHelp("enter/space", "choose"),
		),
	}
}

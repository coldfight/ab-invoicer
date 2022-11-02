package invoice_list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
)

// The delegate is used if you want to manipulate or do something
// with a specific list item

// newItemDelegate creates a new item delegate
func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var (
			invoiceNumber models.InvoiceNumber
		)

		if i, ok := m.SelectedItem().(MenuItem); ok {
			invoiceNumber = i.invoiceNumber
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				var cmds []tea.Cmd
				cmds = append(cmds, func() tea.Msg {
					return common.SwitchToStateMsg{
						State: common.InvoiceItemView,
						Data:  invoiceNumber,
					}
				})
				return tea.Batch(cmds...)
			case key.Matches(msg, keys.edit):
				var cmds []tea.Cmd
				cmds = append(cmds, func() tea.Msg {
					return common.SwitchToStateMsg{
						State: common.EditInvoiceFormView,
						Data:  invoiceNumber,
					}
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
	edit   key.Binding
}

// ShortHelp Additional short help entries. This satisfies
// the help.KeyMap interface and is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
	}
}

// FullHelp - Additional full help entries. This satisfies
// the help.KeyMap interface and is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{
		d.choose,
		d.edit,
	}}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter", "v"),
			key.WithHelp("enter/v", "view invoice"),
		),
		edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit invoice"),
		),
	}
}

package main_menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
)

type MenuItem struct {
	title       string
	description string
	state       common.SessionState
}

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.description }
func (i MenuItem) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleHelpMenu key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type Model struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func New() Model {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	// Make initial list of items
	menuItems := []list.Item{
		MenuItem{
			title:       "Create A New Invoice",
			description: "A form will be displayed that will allow you to generate an invoice.",
			state:       common.InvoiceFormView,
		},
		MenuItem{
			title:       "View All Invoices",
			description: "A list of past invoices will be displayed. You can select one to view or edit it.",
			state:       common.InvoiceListView,
		},
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)

	menuList := list.New(menuItems, delegate, 0, 0)
	menuList.SetShowTitle(true)
	menuList.SetShowFilter(true)
	menuList.SetFilteringEnabled(true)
	menuList.SetShowStatusBar(false)
	menuList.SetShowPagination(true)
	menuList.Title = "AB Invoicer"
	menuList.Styles.Title = common.TitleStyle
	menuList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleHelpMenu,
		}
	}

	return Model{
		list:         menuList,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		//top, right, bottom, left := common.AppStyle.GetPadding()
		//m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
		h, v := common.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return common.AppStyle.Render(m.list.View())
}

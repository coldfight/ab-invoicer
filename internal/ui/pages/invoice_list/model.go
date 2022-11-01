package invoice_list

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_service"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
)

type MenuItem struct {
	title         string
	description   string
	state         common.SessionState
	invoiceNumber models.InvoiceNumber
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
	windowSize   tea.WindowSizeMsg
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func New(windowSize tea.WindowSizeMsg) Model {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	// Make initial list of items
	invoices := invoice_service.GetInvoices()
	var menuItems []list.Item
	for _, inv := range invoices {
		menuItems = append(menuItems, MenuItem{
			title:         inv.InvoiceDate.Format("Jan 02, 2006"),
			description:   fmt.Sprintf("%s - %f", inv.BilledTo.Name, inv.Total()),
			state:         common.InvoiceItemView,
			invoiceNumber: inv.InvoiceNumber,
		})
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)

	menuList := list.New(menuItems, delegate, 0, 0)
	menuList.SetShowTitle(true)
	menuList.SetShowFilter(true)
	menuList.SetFilteringEnabled(true)
	menuList.SetShowStatusBar(false)
	menuList.SetShowPagination(true)
	h, v := common.AppStyle.GetFrameSize()
	menuList.SetSize(windowSize.Width-h, windowSize.Height-v)
	menuList.Title = "All Invoices"
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

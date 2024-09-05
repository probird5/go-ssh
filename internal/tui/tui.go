package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/probird5/go-ssh/config"
	"github.com/probird5/go-ssh/internal/ssh"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// Updated item struct with title, address, and description fields
type item struct {
	title, address, description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.address + "\n" + i.description }  // Return both address and description
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
	cfg  config.Config
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			// Handle the item selection
			selectedItem, ok := m.list.SelectedItem().(item)
			if ok {
				serverAddress := getServerAddress(m.cfg, selectedItem.title)
				ssh.ConnectToServer(selectedItem.title, serverAddress)
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

// StartTUI initializes and starts the text user interface
func StartTUI(cfg config.Config) error {
	// Populate the list with servers from the config file
	var items []list.Item
	for _, server := range cfg.Servers {
		// Ensure description is properly included if present
		desc := fmt.Sprintf("address: %s", server.Address)
		if server.Description != "" {
			desc += fmt.Sprintf("\ndescription: %s", server.Description) // Ensure description is added
		}
		items = append(items, item{title: server.Name, address: desc, description: server.Description})
	}

	// Set up the model
	delegate := list.NewDefaultDelegate()
	itemList := list.New(items, delegate, 50, 20) // Set default width/height
	itemList.Title = "SSH Servers"

	m := model{list: itemList, cfg: cfg}

	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	return err
}

// Function to get the server address based on the selected item
func getServerAddress(cfg config.Config, serverName string) string {
	for _, server := range cfg.Servers {
		if server.Name == serverName {
			return server.Address
		}
	}
	return ""
}


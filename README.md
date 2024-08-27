# Go SSH

This is a terminal-based interactive tool built using Go, which allows users to select from a list of servers and establish an SSH connection either in a new `alacritty` terminal window or a new `tmux` tab.

## Features

- Interactive user interface built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).
- Provides a list of servers to connect via SSH.
- Prompts the user to choose between opening the SSH session in a new `alacritty` terminal window or a new `tmux` tab.
- Easy to customize with additional servers.

## Prerequisites

- Go 1.16 or later installed on your system.
- `tmux` installed and configured.
- `alacritty` installed and configured.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/ssh-connection-tool.git
    cd ssh-connection-tool
    ```

2. **Build the tool:**

    ```bash
    go build -o ssh-tool main.go
    ```

3. **Run the tool:**

    ```bash
    ./ssh-tool
    ```

## Usage

1. Run the tool from your terminal:

    ```bash
    ./ssh-tool
    ```

2. Use the arrow keys to navigate the list of servers.

3. Press `Enter` to select a server.

4. When prompted, choose:
   - `a` to open the SSH connection in a new `alacritty` terminal window.
   - `t` to open the SSH connection in a new `tmux` tab.

5. The SSH session will start in the selected environment.

## Configuration

**will be adding a config file so you do not need to modify the `` **

To add or remove servers from the list:

1. Open the `main.go` file in your preferred text editor.
2. Modify the `items` slice in the `main` function with your desired server names and descriptions.
3. Update the `getServerAddress` function to include the SSH connection details for each server.

```go
items := []list.Item{
    item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
    item{title: "Linux Server", desc: "My main development server"},
    // Add more items as needed
}

func getServerAddress(serverName string) string {
    switch serverName {
    case "Raspberry Pi’s":
        return "pi@192.168.1.10"
    case "Linux Server":
        return "user@linux-server.com"
    // Add more cases for other items if needed
    default:
        return ""
    }
}
```

4 . Save your changes and rebuild the tool:

```
go build -o ssh-tool main.go
```

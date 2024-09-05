package ssh

import (
  "os"
  "fmt"
  "os/exec"
)

func ConnectToServer(serverName, serverAddress string) {
  if serverAddress == "" {
    fmt.Println("No server associated with this item")
    return
  }

  cmd := exec.Command("alacritty", "-e", "ssh", serverAddress)

  cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start() // Start the terminal command in a new process
	if err != nil {
		fmt.Printf("Error starting new alacritty session: %v\n", err)
}
}

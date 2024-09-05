package main

import( 
  "fmt"
  "os"
  "github.com/probird5/go-ssh/config"
  "github.com/probird5/go-ssh/internal/tui"
)

func main() {
  //loads the configuration file
  cfg, err := config.LoadConfig()
  if err != nil {
    fmt.Println("error loading the configuration file")
    os.Exit(1)
  }
  
  // Start the TUI applicaiton with the loaded config
  if err := tui.StartTUI(cfg); err !=nil {
    fmt.Println("error running the TUI:", err)
    os.Exit(1)
  }

}

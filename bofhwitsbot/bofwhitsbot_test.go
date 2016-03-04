package bofhwitsbot

import (
	"os"
	"testing"
)

// tests don't work.  sorry

func TestMain(m *testing.M) {
	// Initialize so we can actually run tests
	botSetup()
	retVal := m.Run()
	botTeardown()
	os.Exit(retVal)

}

func botSetup() {
	// bot := BofhwitsBot{ConfigFilePath: "../config/bofhwits.yaml"}
}

func botTeardown() {

}

// test that we can handle the default config file
func Test_Config(t *testing.T) {
	t.Skip("Test not implemented.")
	// if _, err := os.Stat("../config/bofhwits.yaml"); os.IsNotExist(err) {
	// 	t.Fatalf("../config/bofhwits.yaml not found")
	// }
	//
	// err := bot.LoadConfig()
	// if err != nil {
	// 	t.Errorf("Configuration: %v\n", err)
	// }
}

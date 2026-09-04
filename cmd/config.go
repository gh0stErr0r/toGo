package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"toGo/utils"

	"github.com/spf13/viper"
)

var serverFile string // Path to the configuration file

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Server   string `mapstructure:"server"`   // Field to store the server URL
	Username string `mapstructure:"username"` // Field to store the username
	Password string `mapstructure:"password"` // Field to store the password
}

// Initialize the server file path based on the operating system.
func Init() {
	switch runtime.GOOS {
	case "linux":
		serverFile = utils.ExpandHome("~/.config/toGo/server.json")
	case "darwin": // macOS
		serverFile = utils.ExpandHome("~/.config/toGo/server.json")
	case "windows":
		serverFile = filepath.Join(os.Getenv("USERPROFILE"), "go", "bin", "server.json")
	default:
		fmt.Println("Unknown operating system")
		return
	}

	// Create the directory for the server file if it does not exist
	if err := os.MkdirAll(filepath.Dir(serverFile), os.ModePerm); err != nil {
		utils.Fatal("Failed to create directory for server configuration:", err)
	}

	// Set up Viper
	viper.SetConfigFile(serverFile)
	viper.SetConfigType("json") // Specify the config file type
	viper.AutomaticEnv()        // Automatically read environment variables that match

	// Read the configuration file if it exists
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		utils.Fatal("Failed to read configuration file:", err)
	}
}

// LoadServerConfig loads the server configuration from the configuration file.
func LoadServerConfig() (ServerConfig, error) {
	var config ServerConfig
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}

// SaveServerConfig saves the server configuration to the configuration file.
func SaveServerConfig(config ServerConfig) error {
	viper.Set("server", config.Server)
	viper.Set("username", config.Username)
	viper.Set("password", config.Password)

	// Write the configuration to the file
	return viper.WriteConfigAs(serverFile) // Create or overwrite the JSON file
}

// GetServer returns the server URL from the configuration.
func GetServer() (string, error) {
	config, err := LoadServerConfig()
	if err != nil {
		return "", err
	}
	return config.Server, nil
}

// GetCredentials returns the username and password from the configuration.
func GetCredentials() (string, string, error) {
	config, err := LoadServerConfig()
	if err != nil {
		return "", "", err
	}
	return config.Username, config.Password, nil
}

// SetServer sets the server URL in the configuration.
func SetServer(server string) error {
	config, err := LoadServerConfig()
	if err != nil {
		return err
	}
	config.Server = server
	return SaveServerConfig(config)
}

// SetCredentials sets the username and password in the configuration.
func SetCredentials(username, password string) error {
	config, err := LoadServerConfig()
	if err != nil {
		return err
	}
	config.Username = username
	config.Password = password
	return SaveServerConfig(config)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	config := &DBConfig{}
	err := UpdateConfig("config.json", config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config: %+v\n", config)

	done := make(chan bool)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			panic(err)
		}
		defer watcher.Close()
		fmt.Println("Watching for changes in config.json")
		err = watcher.Add("config.json")
		if err != nil {
			panic(err)
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("Event: %v\n", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Config file modified")
					err := UpdateConfig("config.json", config)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
					}
					fmt.Printf("Config: %+v\n", config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Error: %v -> not ok\n", err)
			}

		}
	}()

	<-done
}

func UpdateConfig(filename string, config *DBConfig) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, config)
	if err != nil {
		return err
	}

	return nil
}

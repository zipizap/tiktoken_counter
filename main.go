package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkoukk/tiktoken-go"
)

func main() {
	// Check if filename and encoding name were provided
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage:   tiktoken_counter <'gpt-3.5-turbo' or 'gpt-4'> <filename>")
		os.Exit(1)
	}
	modelName := os.Args[1]
	filename := os.Args[2]
	encodingName := tiktoken.MODEL_TO_ENCODING[modelName]

	// Get the encoding
	tiktoken, err := tiktoken.GetEncoding(encodingName)
	if err != nil {
		panic(err)
	}

	// Read text from file
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Encode text and count tokens
	tokens := tiktoken.Encode(string(text), nil, nil)
	count := len(tokens)

	fmt.Println(count)

	// Watch file for modifications
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <-sig:
				fmt.Println("Exiting...")
				os.Exit(0)
			default:
				// Check file modification time
				info, err := os.Stat(filename)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error checking file:", err)
					time.Sleep(1 * time.Second)
					continue
				}

				// If file has been modified, recalculate token count
				if info.ModTime().After(time.Now().Add(-1 * time.Second)) {
					text, err := ioutil.ReadFile(filename)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error reading file:", err)
						time.Sleep(1 * time.Second)
						continue
					}

					tokens := tiktoken.Encode(string(text), nil, nil)
					count := len(tokens)

					fmt.Println(count, "tokens")
				}

				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Wait for program to exit
	select {}
}

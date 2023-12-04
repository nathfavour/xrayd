package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"gopkg.in/ini.v1"
)

func main() {

	art := `                                .___
	___  _______________  ___.__. __| _/
	\  \/  /\_  __ \__  \<   |  |/ __ | 
	 >    <  |  | \// __ \\___  / /_/ | 
	/__/\_ \ |__|  (____  / ____\____ | 
		  \/            \/\/         \/ `

	fmt.Println(art)

	if len(os.Args) < 2 {
		processInputFolder()
	} else {
		email := os.Args[1]
		runComplexGo(email)
	}
}

func processInputFolder() {
	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Println("Failed to read settings.ini:", err)
		return
	}

	multithreading := cfg.Section("DEFAULT").Key("multithreading").MustInt(0)
	multivalue := cfg.Section("DEFAULT").Key("multivalue").MustInt(3)
	keep := cfg.Section("DEFAULT").Key("keep").MustInt(0)

	if multithreading == 1 {
		fmt.Println("multithreading enabled.")
		var wg sync.WaitGroup
		sem := make(chan bool, multivalue)

		err := filepath.Walk("INPUT", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					wg.Add(1)
					sem <- true
					go func(email string) {
						defer wg.Done()
						runComplexGo(email)
						<-sem
					}(scanner.Text())
				}

				if err := scanner.Err(); err != nil {
					return err
				}

				if keep == 0 {
					ioutil.WriteFile(path, []byte(""), 0644)
				}
			}

			return nil
		})

		if err != nil {
			fmt.Println("Failed to process INPUT folder:", err)
		}

		wg.Wait()
	} else {
		// Process files sequentially...
		err := filepath.Walk("INPUT", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					runComplexGo(scanner.Text())
				}

				if err := scanner.Err(); err != nil {
					return err
				}

				if keep == 0 {
					ioutil.WriteFile(path, []byte(""), 0644)
				}
			}

			return nil
		})

		if err != nil {
			fmt.Println("Failed to process INPUT folder:", err)
		}
	}
}

func runComplexGo(email string) {
	cmd := exec.Command("go", "run", "cli/complex.go", email)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to run complex.go:", err)
		return
	}

	fmt.Println(string(out))
}

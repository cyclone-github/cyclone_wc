package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"encoding/base64"
)

// version history
// vbeta1-2023-1-2.1000; initial release

// TODO: 
// optimize code for faster performance

// INFO:
// reads file from -i flag and returns word count
// currently cyclone word count is not as fast as gnu wc, but provides an alternative that is cross compiled for Linux, Windows & Mac
// writen by cyclone

func versionFunc() {
	fmt.Println("Cyclone Word Count")
    funcBase64Decode("dmJldGExLTIwMjMtMS0yLjEwMDA7IGluaXRpYWwgcmVsZWFzZQo=")
}

// help function
func helpFunc() {
    versionFunc()
    str := "Example Usage:\n"+
	"\n./wc -i input.txt\n"
	fmt.Println(str)
	os.Exit(0)
}

// main function
func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "Input file to process")
	version := flag.Bool("version", false, "Program version:")
    cyclone := flag.Bool("cyclone", false, "wc")
    help := flag.Bool("help", false, "Prints help:")
	flag.Parse()

	// run sanity checks for -version & -help
    if *version == true {
        versionFunc()
        os.Exit(0)
    } else if *cyclone == true {
        funcBase64Decode("Q29kZWQgYnkgY3ljbG9uZSA7KQo=")
        os.Exit(0)
    } else if *help == true {
        helpFunc()
    }
	// run sanity checks on wordlist input (-i)
    if len(inputFile) < 1 {
        fmt.Println("--> missing '-i input.txt' <--\n")
        helpFunc()
        os.Exit(0)
    }
	// open input file
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("--> Error opening input file: %v <--\n", err)
		os.Exit(1)
	}
	defer input.Close()
	
	// create read buffers
	// input buffer
	inputBuffer := bufio.NewScanner(input)
	// set buffer capacity to 10mb
	const maxCapacity = 100*1024*1024
	buf := make([]byte, maxCapacity)
	inputBuffer.Buffer(buf, maxCapacity)
	
	// create WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	
	// create goroutine bool channel
	done := make(chan bool)
	// start hashgen goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Counting lines:\t", inputFile)
		startTime := time.Now()
		linesCount := 0
		for inputBuffer.Scan() {
			linesCount++
		}
		elapsedTime := time.Since(startTime)
		linesPerSecond := float64(linesCount) / elapsedTime.Seconds() *0.001 // convert to thousand hashes per second
		fmt.Printf("Finished in:\t %v (%.0f lines/sec)\n", elapsedTime, linesPerSecond)
		fmt.Printf("Total lines:\t %d\n", linesCount)
		done <- true
	}()

	// Wait for goroutine to finish
	<-done

	// Wait for sync wait groupto finish <-- doesn't need sync / wait group yet due to having channels, but implimented for testing
	wg.Wait()
}

// base64 decode function used for displaying encoded messages
func funcBase64Decode(line string) {
    str, err := base64.StdEncoding.DecodeString(line)
    if err != nil {
		log.Println("--> Text doesn't appear to be base64 encoded. <--")
        os.Exit(0)
    }
    fmt.Printf("%s\n", str)
}

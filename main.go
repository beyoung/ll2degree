package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/youth/ll2degres/dms"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ProcessTxt(filename, outfilename string) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {

	}

	outFile, err := os.OpenFile(outfilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

	}
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(outFile)
	buf := bytes.NewBuffer(file)
	counter := 1
	for {
		line, err := buf.ReadString('\n')

		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}

			}
		}
		if err != nil && err != io.EOF {
		}
		s := strings.Split(line, ",")
		lon, err := strconv.ParseFloat(s[1], 8)
		lat, err := strconv.ParseFloat(s[2], 8)
		dLat, dLon, err := dms.NewDMS(lat, lon)

		datawriter.WriteString(fmt.Sprintf("%d,%s,%s,%s", counter, dLon.String(), dLat.String(), s[3]))
		counter += 1
	}
	datawriter.Flush()
}

func process(baseDir string) []string {
	var files []string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if path != baseDir && filepath.Ext(path) != ".csv" {
			files = append(files, path)
		}
		if filepath.Ext(path) == ".csv" {
			err := os.Remove(path)
			if err != nil {
				log.Println(err)
			}

		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return files
	}
}

var dir = flag.String("d", "data", "目录参数")

func main() {
	fmt.Println(`
      ___       ___       ___           ___           ___           ___           ___           ___     
     /\__\     /\__\     /\  \         /\  \         /\  \         /\  \         /\  \         /\  \    
    /:/  /    /:/  /    /::\  \       /::\  \       /::\  \       /::\  \       /::\  \       /::\  \   
   /:/  /    /:/  /    /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\:\  \  
  /:/  /    /:/  /    /:/  \:\__\   /::\~\:\  \   /:/  \:\  \   /::\~\:\  \   /::\~\:\  \   /::\~\:\  \ 
 /:/__/    /:/__/    /:/__/ \:|__| /:/\:\ \:\__\ /:/__/_\:\__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/\:\ \:\__\
 \:\  \    \:\  \    \:\  \ /:/  / \:\~\:\ \/__/ \:\  /\ \/__/ \/_|::\/:/  / \:\~\:\ \/__/ \:\~\:\ \/__/
  \:\  \    \:\  \    \:\  /:/  /   \:\ \:\__\    \:\ \:\__\      |:|::/  /   \:\ \:\__\    \:\ \:\__\  
   \:\  \    \:\  \    \:\/:/  /     \:\ \/__/     \:\/:/  /      |:|\/__/     \:\ \/__/     \:\ \/__/  
    \:\__\    \:\__\    \::/__/       \:\__\        \::/  /       |:|  |        \:\__\        \:\__\    
     \/__/     \/__/     ~~            \/__/         \/__/         \|__|         \/__/         \/__/    `)
	flag.Parse()
	start := time.Now()
	files := process(*dir)
	for idx := range files {
		log.Println("process file", files[idx])
		ProcessTxt(files[idx], strings.Split(files[idx], ".")[0]+"-degree.csv")
	}
	log.Println(fmt.Sprintf("Process %d files  cost %s", len(files), time.Since(start).String()))
}

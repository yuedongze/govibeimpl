package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/yuedongze/govibeimpl"
)

func main() {
	name := flag.String("name", "", "")
	flag.Parse()

	pkg := os.Getenv("GOPACKAGE")
	cmd := exec.Command("go", "doc", "-src", pkg+"."+*name)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Description for interface to be implemented:")
	fmt.Println(string(out))

	resp := govibeimpl.InvokeGenAI(string(out))
	if resp == "" {
		return
	}
	resp = "// AI generated implementation. PROCEED WITH CAUTION.\n" + resp

	file := os.Getenv("GOFILE")
	ext := path.Ext(file)
	newFile := strings.TrimSuffix(file, ext) + "_ai_gen" + ext
	fmt.Printf("Writing to file: %s\n", newFile)
	os.WriteFile(newFile, []byte(resp), 0644)
}

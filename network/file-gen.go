package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func fn_gen_file(f_path string, f_size_min int, f_size_max int) {

	f_size := fn_random(f_size_min, f_size_max)

	var cmd_file_count bytes.Buffer
	cmd_file_count.WriteString("count=")
	cmd_file_count.WriteString(strconv.Itoa(f_size))

	var cmd_file_of bytes.Buffer
	cmd_file_of.WriteString("of=")
	cmd_file_of.WriteString(f_path)
	cmd_file_of.WriteString("upload.dat")

	cmd := exec.Command("/bin/dd", "bs=1024000", cmd_file_count.String(), "if=/dev/urandom", cmd_file_of.String())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("/bin/dd", "bs=1024000", cmd_file_count.String(), "if=/dev/urandom", cmd_file_of.String())
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Fatal(err)
	}
	log.Printf("Creating File: upload.dat with size: %vM", strconv.Itoa(f_size))

}

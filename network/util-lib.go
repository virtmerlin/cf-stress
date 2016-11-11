package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func fn_gen_uuid() (uuid string) {
	f_uuid, err := exec.Command("bash", "-c", "for i in {1..1};do uuidgen; done").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return string(f_uuid[:])
}

func fn_gen_uuid_bank() (uuid_bank []string) {
	const count = 5
	var returnvar [count]string
	for i := 0; i < count; i++ {
		returnvar[i] = fn_gen_uuid()
	}
	return returnvar[:]
}

func fn_random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func fn_delete_file(file string) {
	var cmd_file_name bytes.Buffer
	cmd_file_name.WriteString("*")
	cmd_file_name.WriteString(file[:len(file)-1])
	cmd_file_name.WriteString("*")
	cmd := exec.Command("rm", cmd_file_name.String())
	err := cmd.Run()
	if err != nil {
		log.Println("rm %v", cmd.Args[1])
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Fatal(err)
	}
}

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
	f_uuid, err := exec.Command("bash", "-c", "r=( $(openssl rand 100000 | sha1sum) ); printf \"%s${r[0]:0:13}\n\"").Output()
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

func fn_delete_file(filepath string) {

	cmd := exec.Command("/bin/rm", filepath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("rm ", cmd.Args[1])
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Fatal(err)
	}
}

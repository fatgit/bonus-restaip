package main

import (
	"encoding/base64"
	"log"
	"os/exec"
)

func (c Client) lsHelm() []byte {
	cmdHelm := exec.Command("helm" , "ls", "-q")
	out, err := cmdHelm.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func (c Client) installHelm() []byte {
	cmd := exec.Command("helm" , "install", "bonus/bonus", "--name="+ c.Name, "--set", "clientName=" + c.Name,
		"--set", "mysql.password=" + base64.StdEncoding.EncodeToString([]byte(c.clientMysqlPassword())),
		"--set", "ingress.host=" + c.URL)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func (c Client) delHelm() []byte {
	cmd := exec.Command("helm" , "del", "--purge", c.Name)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
package main

import (
	"fmt"
	"github.com/m-tarasov/launcher/menu"
	"os/user"
	"log"
	"path"
	"os/exec"
	"os"
	"syscall"
)

func main() {
	usr, err := user.Current()

	if err != nil {
		log.Fatalln("Failed to determine current user")
	}

	list := menu.Load(path.Join(usr.HomeDir, ".menu.json"))

	if len(list.List) == 0{
		fmt.Println("no menu entries found")
		return
	}


	g := menu.InitGui()

	item, selected := g.Run(list)
	if err != nil {
		log.Panicln(err)
	}

	if !selected {
		return
	}


	cmd := exec.Command(item.Cmd[0], item.Cmd[1:]...)
	fmt.Println("Run", item.Cmd[0])
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		os.Exit(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
	}
}

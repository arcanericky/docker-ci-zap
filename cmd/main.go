// +build windows

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arcanericky/dockercizap"
	"golang.org/x/sys/windows"
)

var version string = "development"
var zapper func(string) error = dockercizap.Zap

type config struct {
	folder string
}

// isAdmin implements code from https://coolaj86.com/articles/golang-and-windows-and-admins-oh-my/
func isAdmin() bool {
	var sid *windows.SID

	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}

	token := windows.Token(0)

	member, err := token.IsMember(sid)
	if err != nil {
		return false
	}

	return member
}

func parseFlags(exe string, args []string) (config, error) {
	var cfg config

	flags := flag.NewFlagSet(exe, flag.ContinueOnError)
	flags.StringVar(&cfg.folder, "folder", "", "Folder to zap")

	if err := flags.Parse(args); err != nil {
		return config{}, err
	}

	return cfg, nil
}

func execute(args []string) int {
	fmt.Printf("docker-ci-zap version %s\n", version)

	if !isAdmin() {
		fmt.Fprintln(os.Stderr, "This program must be ran with administrative privileges.")
		return 1
	}

	cfg, err := parseFlags(args[0], args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing command line: %s\n", err)
		return 1
	}

	if cfg.folder == "" {
		fmt.Fprintln(os.Stderr, "Folder not specified. Use option \"-folder\".")
		return 1
	}

	if err := zapper(cfg.folder); err != nil {
		fmt.Fprintf(os.Stderr, "Error zapping folder \"%s\"m: %s\n", cfg.folder, err)
		return 1
	}

	fmt.Printf("Folder \"%s\" zapped successfully", cfg.folder)
	return 0
}

func main() {
	os.Exit(execute(os.Args))
}

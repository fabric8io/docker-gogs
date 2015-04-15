package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/docker/libcontainer/system"
	"github.com/docker/libcontainer/user"
)

// this function comes from libcontainer/namespaces/init.go
// we don't use that directly because we don't want the whole namespaces package imported here

// SetupUser changes the groups, gid, and uid for the user inside the container
func SetupUser(u string) error {
	// Set up defaults.
	defaultExecUser := user.ExecUser{
		Uid:  syscall.Getuid(),
		Gid:  syscall.Getgid(),
		Home: "/",
	}

	execUser, err := getUser(u, &defaultExecUser)
	if err != nil {
		return err
	}

	if err := syscall.Setgroups(execUser.Sgids); err != nil {
		return fmt.Errorf("setgroups %s", err)
	}

	if err := system.Setgid(execUser.Gid); err != nil {
		return fmt.Errorf("setgid %s", err)
	}

	if err := system.Setuid(execUser.Uid); err != nil {
		return fmt.Errorf("setuid %s", err)
	}

	// if we didn't get HOME already, set it based on the user's HOME
	if envHome := os.Getenv("HOME"); envHome == "" {
		if err := os.Setenv("HOME", execUser.Home); err != nil {
			return fmt.Errorf("set HOME %s", err)
		}
	}

	os.Setenv("USER", u)

	return nil
}

func getUser(u string, defaultExecUser *user.ExecUser) (*user.ExecUser, error) {
	passwdFile, err := user.GetPasswdFile()
	if err != nil {
		return nil, err
	}

	groupFile, err := user.GetGroupFile()
	if err != nil {
		return nil, err
	}

	execUser, err := user.GetExecUserFile(u, defaultExecUser, passwdFile, groupFile)
	if err != nil {
		return nil, fmt.Errorf("get supplementary groups %s", err)
	}
	return execUser, nil
}

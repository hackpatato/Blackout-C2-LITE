package Whoami

import (
	"log"
	"os/user"
)

// ahhh
type SystemUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Uid      string `json:"uid"`
	Gid      string `json:"gid"`
	HomeDir  string `json:"gid"`
}

func Whoami() *SystemUser {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("ERROR , no user knowlage: %v", err)
	}

	sysUser := &SystemUser{
		Username: currentUser.Username,
		Name:     currentUser.Name,
		Uid:      currentUser.Uid,
		Gid:      currentUser.Gid,
		HomeDir:  currentUser.HomeDir,
	}

	return sysUser
}

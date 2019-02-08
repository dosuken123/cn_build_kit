package service

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
)

func (s Service) EnsureUser() error {
	var err error

	if s.User.Name == "" {
		return nil
	}

	// Add current user to the group
	groupName := s.addGroup()
	userName := s.addUser(groupName)
	s.addCurrentUserToGroup(groupName)
	s.correctPermission(groupName, userName)

	return err
}

func (s Service) addGroup() string {
	var groupName string
	var gid string

	if s.User.Group != "" {
		groupName = s.User.Group
	} else {
		groupName = s.User.Name
	}

	if g, error := user.LookupGroup(groupName); error == nil {
		return g.Name
	}

	if s.User.GID != "" {
		gid = fmt.Sprintf("--gid %s", s.User.GID)
	} else {
		gid = ""
	}

	cmdName := fmt.Sprintf("sudo addgroup %s %s", gid, groupName)

	s.ExecuteCommandWithLog("ensure_user", cmdName)

	return groupName
}

func (s Service) addUser(groupName string) string {
	if _, error := user.Lookup(s.User.Name); error == nil {
		return s.User.Name
	}

	var homeDir string
	var uid string

	if s.User.HomeDir != "" {
		homeDir = fmt.Sprintf("--home %s", s.User.HomeDir)
	} else {
		homeDir = ""
	}

	if s.User.UID != "" {
		uid = fmt.Sprintf("--uid %s", s.User.UID)
	} else {
		uid = ""
	}

	g, _ := user.LookupGroup(groupName)

	cmdName := fmt.Sprintf("sudo adduser %s --gid %s %s --shell /bin/bash --gecos \"\" --disabled-password %s", uid, g.Gid, homeDir, s.User.Name)

	s.ExecuteCommandWithLog("ensure_user", cmdName)

	return s.User.Name
}

func (s Service) correctPermission(groupName string, userName string) string {
	cmdName := fmt.Sprintf("sudo chown -R %s:%s %s", groupName, userName, s.GetServiceDir())
	s.ExecuteCommandWithLog("ensure_user", cmdName)

	// cmdName = fmt.Sprintf("sudo chmod -R g+wrx %s", s.GetServiceDir())
	cmdName = fmt.Sprintf("sudo chmod -R 777 %s", s.GetServiceDir())
	s.ExecuteCommandWithLog("ensure_user", cmdName)

	return s.User.Name
}

func (s Service) addCurrentUserToGroup(groupName string) {
	u, error := user.Current()

	if error != nil {
		log.Fatal(error)
	}

	gids, error := u.GroupIds()

	if error != nil {
		log.Fatal(error)
	}

	g, error := user.LookupGroup(groupName)

	if error != nil {
		log.Fatal(error)
	}

	for _, gid := range gids {
		if gid == g.Gid {
			return
		}
	}

	cmdName := fmt.Sprintf("sudo usermod --append --groups %s %s", groupName, u.Name)
	s.ExecuteCommandWithLog("ensure_user", cmdName)
}

func (s Service) GetUserName() string {
	if s.User.Name != "" {
		return s.User.Name
	}

	user, error := user.Current()

	if error != nil {
		return ""
	}

	return user.Name
}

func (s Service) GetUserGroup() string {
	if s.User.Group != "" {
		return s.User.Group
	}

	g, error := s.getGroup()

	if error != nil {
		return ""
	}

	return g.Name
}

func (s Service) GetUID() string {
	if s.User.UID != "" {
		return s.User.UID
	}

	user, error := user.Lookup(s.GetUserName())

	if error != nil {
		return ""
	}

	return user.Uid
}

func (s Service) GetGID() string {
	if s.User.GID != "" {
		return s.User.GID
	}

	g, error := s.getGroup()

	if error != nil {
		return ""
	}

	return g.Gid
}

func (s Service) GetUIDUint32() uint32 {
	u, _ := strconv.ParseUint(s.GetUID(), 10, 32)

	return uint32(u)
}

func (s Service) GetGIDUint32() uint32 {
	u, _ := strconv.ParseUint(s.GetGID(), 10, 32)

	return uint32(u)
}

func (s Service) GetUIDInt() int {
	u, _ := strconv.Atoi(s.GetUID())

	return u
}

func (s Service) GetGIDInt() int {
	u, _ := strconv.Atoi(s.GetGID())

	return u
}

func (s Service) getGroup() (*user.Group, error) {
	u, error := user.Lookup(s.GetUserName())

	if error != nil {
		return nil, error
	}

	gids, error := u.GroupIds()

	if error != nil {
		return nil, error
	}

	group, error := user.LookupGroupId(gids[0])

	if error != nil {
		return nil, error
	}

	return group, nil
}

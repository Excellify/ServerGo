package datastructure

import (
	"github.com/SevenTV/ServerGo/src/configure"
)

// The default role.
// It grants permissions for users without a defined role
var DefaultRole *Role = &Role{
	Allowed: configure.Config.GetInt64("default_permissions"),
	Denied:  0,
	Default: true,
}

var DeletedUser *User = &User{
	Login:       "*deleteduser",
	DisplayName: "Deleted User",
}

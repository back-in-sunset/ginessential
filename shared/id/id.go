package id

import "fmt"

// ObjID 对象ID
type ObjID string

// ToUserID objid to userid
func (o ObjID) ToUserID() UserID {
	return UserID(o)
}

// ToDemoID objid to demoid
func (o ObjID) ToDemoID() DemoID {
	return DemoID(o)
}

// ToRoleID objid to roleid
func (o ObjID) ToRoleID() RoleID {
	return RoleID(o)
}

// FromID gen ObjID
func FromID(eid fmt.Stringer) ObjID {
	return ObjID(eid.String())
}

// DemoID demo id
type DemoID string

func (d *DemoID) String() string {
	if d == nil {
		return ""
	}

	return string(*d)
}

// RoleID role id
type RoleID string

func (i *RoleID) String() string {
	if i == nil {
		return ""
	}

	return string(*i)
}

// UserID id filed
type UserID string

func (i *UserID) String() string {
	if i == nil {
		return ""
	}

	return string(*i)
}

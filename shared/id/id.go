package id

import "ginessential/pkg/utils"

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

func (o ObjID) String() string {
	return string(o)
}

// NewObjID new objid
func NewObjID() ObjID {
	return ObjID(utils.NanoNumbID())
}

// IDer IDer
type IDer interface {
	String() string
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

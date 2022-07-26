package datamodel

import (
	"strings"

	"gorm.io/gorm"
)

// TODO: not null, not dup

type Organization struct {
	gorm.Model

	Address string `gorm:"type:varchar(128);unique" json:"address"`
	Name    string `gorm:"type:varchar(40);unique" json:"name"`
	Admin   string `gorm:"type:varchar(256)"; json:"admin"`
	Member  string `gorm:"type:varchar(256)"; json:"member"`
}

func (o Organization) ToMassage() OrganizationMassage {
	adminAddresses := strings.Split(o.Admin, ",")
	memberAddresses := strings.Split(o.Member, ",")

	msg := OrganizationMassage{
		Address:         o.Address,
		Name:            o.Name,
		AdminAddresses:  adminAddresses,
		MemberAddresses: memberAddresses,
		NPeople:         len(adminAddresses) + len(memberAddresses),
	}
	return msg
}

type OrganizationMassage struct {
	Address         string   `json:"address"`
	Name            string   `json:"name"`
	AdminAddresses  []string `json:"admin"`
	MemberAddresses []string `json:"member"`
	NPeople         int      `json:"number_of_people"`
}

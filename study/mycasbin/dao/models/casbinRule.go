package models

// CasbinRule  .
type CasbinRule struct {
	PType string `json "ptype" xorm:"varchar(100) index not null default ''"`
	V0    string `json "v0" xorm:"varchar(100) index not null default ''"`
	V1    string `json "v1" xorm:"varchar(100) index not null default ''"`
	V2    string `json "v2" xorm:"varchar(100) index not null default ''"`
	V3    string `json "v3" xorm:"varchar(100) index not null default ''"`
	V4    string `json "v4" xorm:"varchar(100) index not null default ''"`
	V5    string `json "v5" xorm:"varchar(100) index not null default ''"`
}


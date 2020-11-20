package dao

import (
	"fmt"
	"github.com/casbin/casbin"
)

type CasbinDao struct{
	Enforcer *casbin.Enforcer
}

func (dao CasbinDao) PolicyAdd(ptype string, params ...interface{}) bool {
	if ptype=="p"{
		return dao.Enforcer.AddNamedPolicy(ptype, params...)
	}else{
		has := dao.Enforcer.HasGroupingPolicy(params...)
		fmt.Println("has=",has)
		if !has {
			return dao.Enforcer.AddNamedGroupingPolicy(ptype, params...)
		}
		return !has
	}
}

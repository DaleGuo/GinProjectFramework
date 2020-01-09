package service

import (
	"github.com/casbin/casbin"
)

var Enforcer *casbin.Enforcer

func AddPolicy(sub string,obj string,act string) (bool,string) {
	result := Enforcer.AddPolicy(sub,obj,act)
	if !result {
		return result,"policy has existed"
	} else {
		Enforcer.SavePolicy()
		return result,"policy add successfully"
	}
}

func RemovePolicy(sub string,obj string,act string) (bool,string) {
	result := Enforcer.RemovePolicy(sub,obj,act)
	if !result {
		return result,"policy is not existed"
	} else {
		Enforcer.SavePolicy()
		return result,"policy remove successfully"
	}
}
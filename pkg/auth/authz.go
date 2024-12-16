// Copyright 2024 Innkeeper cheng <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"time"
)

const (
	// casbin 访问控制模型.
	aclModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

// Authz 定义了一个授权器，提供授权功能
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz 创建一个使用casbin完成授权的授权器
func NewAuthz(db *gorm.DB) (*Authz, error) {
	adapter, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	m, _ := model.NewModelFromString(aclModel)

	//	Init the enforcer
	enforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	enforcer.StartAutoLoadPolicy(5 * time.Second)
	a := &Authz{enforcer}
	return a, nil
}

// Authorize 用来进行授权
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}

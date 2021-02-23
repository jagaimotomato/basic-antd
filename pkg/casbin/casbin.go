package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB) *casbin.SyncedEnforcer {
	apter, err := gormAdapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	log.SetLogger(&log.DefaultLogger{})
	e.EnableLog(true)
	return e
}

package dao

import (
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"mycasbin/dao/models"
	"testing"
)

func TestCasbin(t *testing.T) {
	a:= xormadapter.NewAdapter(DBNAME, DNS,true)
	e := casbin.NewEnforcer("../conf/rbac_models.conf", a)

	var policy1 = &models.CasbinRule{PType: "p",V0:"admin",V1:"/api/hello",V2:"GET",V3:"",V4:"",V5:""}
	var policy2 = &models.CasbinRule{PType: "p",V0:"admin",V1:"/api/job/add",V2:"GET",V3:"",V4:"",V5:""}
	var policy3 = &models.CasbinRule{PType: "g",V0:"admin",V1:"rooter",V2:"",V3:"",V4:"",V5:""}
	var policy4 = &models.CasbinRule{PType: "g",V0:"admin",V1:"developer",V2:"",V3:"",V4:"",V5:""}

	tests := map[string]struct {
		policy       *models.CasbinRule
	}{
		"AddpolicyCase1": {policy: policy1},
		"AddpolicyCase2": {policy: policy2},
		"AddpolicyCase3": {policy: policy3},
		"AddpolicyCase4": {policy: policy4},
	}

	var dao CasbinDao
	dao.Enforcer = e
	for testname, item := range tests {
		//ok := Enforcer.AddNamedPolicy(item.policy.PType, item.policy.V0, item.policy.V1, item.policy.V2)
		ok := dao.PolicyAdd(item.policy.PType, item.policy.V0, item.policy.V1, item.policy.V2)
		t.Log(testname," ",ok)
		if ok {
			t.Log(testname,"添加成功")
		} else {
			t.Log(testname,"添加失败")
		}
	}
}

func TestCasbin2(t *testing.T) {
	a:= xormadapter.NewAdapter(DBNAME, DNS,true)
	e := casbin.NewEnforcer("../conf/rbac_models.conf", a)

	var dao CasbinDao
	dao.Enforcer = e

	added := e.AddNamedPolicy("p1", "eve", "data3", "read")
	t.Log("added",added)
	t.Log("----------------------------------------")
}

func TestCasbinMethods(t *testing.T) {
	a := xormadapter.NewAdapter(DBNAME, DNS, true)
	e := casbin.NewEnforcer("../conf/rbac_models.conf", a)
	var dao CasbinDao
	dao.Enforcer = e

	//获取当前策略中显示的主题列表
	allSubjects := e.GetAllSubjects()
	t.Log("allSubjects",allSubjects)
	t.Log("----------------------------------------")

	//获取当前命名策略中显示的主题列表
	allNamedSubjects := e.GetAllNamedSubjects("p")
	t.Log("allNamedSubjects",allNamedSubjects)
	t.Log("----------------------------------------")

	//获取当前策略中显示的对象列表
	allObjects := e.GetAllObjects()
	t.Log("allSubjects",allObjects)
	t.Log("----------------------------------------")

	//获取当前命名策略中显示的对象列表
	allNamedObjects := e.GetAllNamedObjects("p")
	t.Log("allNamedObjects",allNamedObjects)
	t.Log("----------------------------------------")

	//获取当前策略中显示的操作列表
	allActions := e.GetAllActions()
	t.Log("allSubjects",allActions )
	t.Log("----------------------------------------")

	//获取当前命名策略中显示的操作列表
	allNamedActions := e.GetAllNamedActions("p")
	t.Log("allNamedActions",allNamedActions)
	t.Log("----------------------------------------")

	//获取当前策略中显示的角色列表
	allRoles := e.GetAllRoles()
	t.Log("allRoles",allRoles)
	t.Log("----------------------------------------")

	//获取当前命名策略中显示的角色列表
	allNamedRoles := e.GetAllNamedRoles("g")
	t.Log("allNamedRoles",allNamedRoles)
	t.Log("----------------------------------------")

	//获取策略中的所有授权规则
	policy := e.GetPolicy()
	t.Log("policy",policy)
	t.Log("----------------------------------------")

	//获取策略中的所有授权规则，可以指定字段筛选器
	filteredPolicy := e.GetFilteredPolicy(0, "alice")
	t.Log("filteredPolicy",filteredPolicy)
	t.Log("----------------------------------------")

	//获取命名策略中的所有授权规则
	namedPolicy := e.GetNamedPolicy("p")
	t.Log("namedPolicy",namedPolicy)
	t.Log("----------------------------------------")

	//获取命名策略中的所有授权规则，可以指定字段过滤器
	filteredNamedPolicy := e.GetFilteredNamedPolicy("p", 0, "bob")
	t.Log("filteredNamedPolicy",filteredNamedPolicy)
	t.Log("----------------------------------------")

	//获取策略中的所有角色继承规则
	groupingPolicy := e.GetGroupingPolicy()
	t.Log("groupingPolicy",groupingPolicy)
	t.Log("----------------------------------------")

	//获取策略中的所有角色继承规则，可以指定字段筛选器
	filteredGroupingPolicy := e.GetFilteredGroupingPolicy(0, "alice")
	t.Log("filteredGroupingPolicy",filteredGroupingPolicy)
	t.Log("----------------------------------------")

	//获取策略中的所有角色继承规则
	namedGroupingPolicy := e.GetNamedGroupingPolicy("g")
	t.Log("namedGroupingPolicy",namedGroupingPolicy)
	t.Log("----------------------------------------")

	//获取策略中的所有角色继承规则
	namedGroupingPolicy = e.GetFilteredNamedGroupingPolicy("g", 0, "alice")
	t.Log("namedGroupingPolicy",namedGroupingPolicy)
	t.Log("----------------------------------------")

	// 确定是否存在授权规则
	hasPolicy := e.HasPolicy("data2_admin", "data2", "read")
	t.Log("hasPolicy",hasPolicy )
	t.Log("----------------------------------------")

	//确定是否存在命名授权规则
	hasNamedPolicy := e.HasNamedPolicy("p", "data2_admin", "data2", "read")
	t.Log("hasNamedPolicy",hasNamedPolicy)
	t.Log("----------------------------------------")

	//向当前策略添加授权规则。 如果规则已经存在，函数返回false，并且不会添加规则。 否则，函数通过添加新规则并返回true
	added := e.AddPolicy("eve", "data3", "read")
	t.Log("added",added)
	t.Log("----------------------------------------")

	// 向当前命名策略添加授权规则。 如果规则已经存在，函数返回false，并且不会添加规则。 否则，函数通过添加新规则并返回true
	added = e.AddNamedPolicy("p", "eve", "data3", "read")
	t.Log("added",added)
	t.Log("----------------------------------------")

	// 从当前策略中删除授权规则
	removed := e.RemovePolicy("alice", "data1", "read")
	t.Log("removed",removed)
	t.Log("----------------------------------------")

	// 移除当前策略中的授权规则，可以指定字段筛选器。 RemovePolicy 从当前策略中删除授权规则
	removed = e.RemoveFilteredPolicy(0, "alice", "data1", "read")
	t.Log("removed",removed)
	t.Log("----------------------------------------")

	//从当前命名策略中删除授权规则
	removed = e.RemoveNamedPolicy("p", "alice", "data1", "read")
	t.Log("removed",removed)
	t.Log("----------------------------------------")

	//从当前命名策略中移除授权规则，可以指定字段筛选器
	removed = e.RemoveFilteredNamedPolicy("p", 0, "alice", "data1", "read")
	t.Log("removed",removed)
	t.Log("----------------------------------------")

	//确定是否存在角色继承规则
	has := e.HasGroupingPolicy("alice", "data2_admin")
	t.Log("has",has)
	t.Log("----------------------------------------")

	//确定是否存在命名角色继承规则
	has = e.HasNamedGroupingPolicy("g", "alice", "data2_admin")
	t.Log("has",has)
	t.Log("----------------------------------------")

	// 向当前策略添加角色继承规则。 如果规则已经存在，函数返回false，并且不会添加规则。 如果规则已经存在，函数返回false，并且不会添加规则
	added = e.AddGroupingPolicy("group1", "data2_admin")
	t.Log("added",added)
	t.Log("----------------------------------------")

	//将命名角色继承规则添加到当前策略。 如果规则已经存在，函数返回false，并且不会添加规则。 否则，函数通过添加新规则并返回true
	added = e.AddNamedGroupingPolicy("g", "group1", "data2_admin")
	t.Log("added",added)
	t.Log("----------------------------------------")

	// 从当前策略中删除角色继承规则
	removed = e.RemoveGroupingPolicy("alice", "data2_admin")
	t.Log("removed ",removed)
	t.Log("----------------------------------------")

	//从当前策略中移除角色继承规则，可以指定字段筛选器
	removed = e.RemoveFilteredGroupingPolicy(0, "alice")
	t.Log("removed ",removed )
	t.Log("----------------------------------------")

	//从当前命名策略中移除角色继承规则
	removed = e.RemoveNamedGroupingPolicy("g", "alice")
	t.Log("removed ",removed)
	t.Log("----------------------------------------")

	//当前命名策略中移除角色继承规则，可以指定字段筛选器
	removed = e.RemoveFilteredNamedGroupingPolicy("g", 0, "alice")
	t.Log("removed ",removed )
	t.Log("----------------------------------------")
}


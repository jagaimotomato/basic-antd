package model

import (
	orm "basic-antd/init/global"
)

type Permission struct {
	PermissionId        int           `gorm:"primaryKey;autoIncrement" json:"permissionId"`
	Title               string        `json:"title"`
	Name                string        `gorm:"size:128" json:"name"`
	Actions             string        `json:"actions"`
	Target              string        `json:"target"`
	HideChildren        int           `gorm:"size:2" json:"hideChildren"`
	HiddenHeaderContent int           `gorm:"size:2" json:"hiddenHeaderContent"` // 隐藏 PageHeader 组件中的页面带的 面包屑和页面标题栏
	Show                int           `gorm:"size:2" json:"show"`
	Redirect            string        `json:"redirect"`
	Describe            string        `json:"describe"`
	ParentId            int           `json:"parentId"`
	Component           string        `json:"component"`
	Icon                string        `gorm:"size:30" json:"icon"`
	Path                string        `json:"path"`
	Sort                int           `json:"sort"`
	Children            []Permission  `gorm:"-" json:"children"`
	ActionList          []interface{} `gorm:"-" json:"actionList"`
	Action              []string      `gorm:"-" json:"action"` // 查询action用
	ActionUpload        []Action      `gorm:"-" json:"actionUpload"`
	By
	Model
}

type Action struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (*Permission) TableName() string {
	return "permission"
}

func (p *Permission) GetPermission() (permission Permission, err error) {
	permissions, err := p.GetList()
	if err != nil {
		return
	}
	for i := 0; i < len(permissions); i++ {
		if permissions[i].PermissionId == p.PermissionId {
			for j := 0; j < len(permissions); j++ {
				if permissions[i].PermissionId != permissions[j].ParentId {
					continue
				}
				if permissions[j].Actions != "" {
					permissions[i].Action = append(permissions[i].Action, permissions[j].Actions)
				}
			}
			permission = permissions[i]
		}
	}
	/*if err = orm.Eloquent.Table(p.TableName()).First(&permission, p.PermissionId).Error; err != nil {
		return
	}*/
	return
}

func (p *Permission) GetPermissionsByRoleId(roleId int) (permissions []Permission, err error) {
	if err = orm.Eloquent.Table(p.TableName()).Joins("left join role_permission on role_permission.permission_id = permission.permission_id").Where("role_permission.role_id = ?", roleId).Order("sort").Find(&permissions).Error; err != nil {
		return
	}
	return
}

func (p *Permission) GetPermissionsByRoleIds(roleIds []int) (permissions []Permission, err error) {
	if err = orm.Eloquent.Select("role_permission.*, permission.*").Joins("left join role_permission on role_permission.permission_id = permission.permission_id").Where("role_permission.role_id in (?)", roleIds).Order("sort").Find(&permissions).Error; err != nil {
		return
	}
	/* todo 将去重方法变为公共方法 interface */
	if len(permissions) < 1024 {
		permissions = RemoveRepByLoop(permissions)
	} else {
		permissions = RemoveRepByMap(permissions)
	}
	return
}

func RemoveRepByLoop(arr []Permission) (result []Permission) {
	for i := range arr {
		flag := true
		for j := range result {
			if arr[i].PermissionId == result[j].PermissionId {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, arr[i])
		}
	}
	return
}

func RemoveRepByMap(arr []Permission) (result []Permission) {
	tempMap := map[interface{}]byte{}
	for _, v := range arr {
		l := len(tempMap)
		tempMap[v] = 0
		if len(tempMap) != l {
			result = append(result, v)
		}
	}
	return
}

func (p *Permission) GetMenus(roleIds []int) (menus []Permission, err error) {
	list, err := p.GetPermissionsByRoleIds(roleIds)
	if err != nil {
		return
	}
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		children := RecursionPermission(&list, list[i])
		menus = append(menus, children)
	}
	return
}

func (p *Permission) GetPermissionsByPid(pageIndex, pageSize int) (permissions []Permission, count int64, err error) {
	table := orm.Eloquent.Table(p.TableName())
	if p.Title != "" {
		table = table.Where("title like ?", "%"+p.Title+"%")
	}
	if err = table.Where("parent_id = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("sort").Find(&permissions).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

func (p *Permission) GetList() (permissions []Permission, err error) {
	table := orm.Eloquent.Table(p.TableName())
	if p.Title != "" {
		table = table.Where("title like ?", "%"+p.Title+"%")
	}
	if p.Path != "" {
		table = table.Where("path = ?", p.Path)
	}
	if p.Actions != "" {
		table = table.Where("actions = ?", p.Actions)
	}
	if err = table.Find(&permissions).Order("sort").Error; err != nil {
		return
	}
	return
}

func (p *Permission) GetRecursionPermission(pageIndex, pageSize int) (permissions []Permission, count int64, err error) {
	list, count, err := p.GetPermissionsByPid(pageIndex, pageSize)
	if err != nil {
		return
	}
	page, err := p.GetList()
	for i := 0; i < len(list); i++ {
		/*if list[i].ParentId != 0 {
			continue
		}*/
		children := RecursionActionListPermission(&page, list[i])
		permissions = append(permissions, children)
	}
	return
}

func RecursionPermission(menus *[]Permission, menu Permission) Permission {
	list := *menus
	children := make([]Permission, 0)
	for i := 0; i < len(list); i++ {
		if menu.PermissionId != list[i].ParentId {
			continue
		}
		m := list[i]
		//fmt.Println(m, "mmmmmmm")
		if m.Actions == "" {
			ms := RecursionPermission(menus, m)
			children = append(children, ms)
		}
	}
	menu.Children = children
	return menu
}

/* 比上面多个获取 actionList */
func RecursionActionListPermission(permissions *[]Permission, permission Permission) Permission {
	list := *permissions
	children := make([]Permission, 0)
	for i := 0; i < len(list); i++ {
		if permission.PermissionId != list[i].ParentId {
			continue
		}
		m := list[i]
		actionList := GetActionList(permissions, permission)
		permission.ActionList = actionList
		if m.Actions == "" {
			ms := RecursionActionListPermission(permissions, m)
			children = append(children, ms)
		}
	}
	if len(children) > 0 {
		permission.Children = children
	} else {
		permission.Children = nil
	}
	return permission
}

func RecursionAllPermission(permissions *[]Permission, permission Permission) Permission {
	list := *permissions
	children := make([]Permission, 0)
	for i := 0; i < len(list); i++ {
		if permission.PermissionId != list[i].ParentId {
			continue
		}
		m := list[i]
		if m.Actions == "" {
			ms := RecursionAllPermission(permissions, m)
			children = append(children, ms)
		} else {
			children = append(children, m)
		}
	}
	if len(children) > 0 {
		permission.Children = children
	} else {
		permission.Children = nil
	}
	return permission
}

func (p *Permission) GetAllPermission() (permissions []Permission, err error) {
	permissionList, err := p.GetList()
	for i := 0; i < len(permissionList); i++ {
		if permissionList[i].ParentId != 0 {
			continue
		}
		children := RecursionAllPermission(&permissionList, permissionList[i])
		permissions = append(permissions, children)
	}
	return
}

func GetActionList(permissions *[]Permission, permission Permission) (actionList []interface{}) {
	list := *permissions
	for i := 0; i < len(list); i++ {
		if permission.PermissionId == list[i].ParentId && list[i].Actions != "" {
			actionList = append(actionList, list[i])
		}
	}
	return
}

func (p *Permission) SetPermission() (permissions []Permission, err error) {
	page, err := p.GetList()
	if err != nil {
		return
	}
	for i := 0; i < len(page); i++ {
		if page[i].ParentId != 0 {
			continue
		}
		children := RecursionPermission(&page, page[i])
		permissions = append(permissions, children)
	}
	return
}

func (p *Permission) Insert() (id int, err error) {
	if err = orm.Eloquent.Table(p.TableName()).Create(&p).Error; err != nil {
		return
	}
	id = p.PermissionId
	return
}

func (p *Permission) BatchInsert(ps *[]Permission) (err error) {
	if err = orm.Eloquent.Table(p.TableName()).Create(ps).Error; err != nil {
		return
	}
	return
}

func (p *Permission) Update() (err error) {
	var permission Permission
	if err = orm.Eloquent.Table(p.TableName()).First(&permission, p.PermissionId).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Table(p.TableName()).Model(&permission).Updates(&p).Error; err != nil {
		return
	}
	return
}

func (p *Permission) DeleteActionByPid() (err error) {
	if err = orm.Eloquent.Table(p.TableName()).Where("parent_id = ?", p.PermissionId).Not("actions = ?", "").Delete(&Permission{}).Error; err != nil {
		return
	}
	return
}

func (p *Permission) Delete() (err error) {
	if err = orm.Eloquent.Table(p.TableName()).Where("permission_id = ?", p.PermissionId).Delete(&Permission{}).Error; err != nil {
		return
	}
	return
}

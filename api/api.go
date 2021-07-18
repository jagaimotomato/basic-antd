package api

import (
	"basic-antd/model"
	"basic-antd/pkg/response"
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateApi(c *gin.Context) {
	var (
		api model.Api
		err error
	)
	if err = c.ShouldBind(&api); err != nil {
		response.Error(c, response.CodeParamErr, nil, "参数绑定错误")
		return
	}
	err = api.Insert()
	if err != nil {
		response.Error(c, response.CodeDBError, nil, err.Error())
		return
	}
	response.Success(c, nil, "新增api成功")
}

func GetApiPage(c *gin.Context) {
	var (
		api       model.Api
		err       error
		pageIndex = 1
		pageSize  = 10
	)
	if index := c.Request.FormValue("page_index"); index != "" {
		pageIndex, _ = strconv.Atoi(index)
	}

	if size := c.Request.FormValue("page_size"); size != "" {
		pageSize, _ = strconv.Atoi(size)
	}

	api.Path = c.Request.FormValue("path")
	api.Method = c.Request.FormValue("method")
	api.Desc = c.Request.FormValue("desc")
	api.Group = c.Request.FormValue("group")

	data, count, err := api.GetPage(pageIndex, pageSize)
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "获取api分页数据出错")
		return
	}
	response.Paginate(c, data, count, pageIndex, pageSize, "获取api分页数据成功")
}

func GetApiList(c *gin.Context) {
	var (
		api model.Api
		err error
	)
	data, err := api.GetList()
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "获取api列表失败")
		return
	}
	response.Success(c, data, "获取列表成功")
}

func GetCasbinApiList(c *gin.Context) {
	var (
		casbin model.CasbinRule
		err    error
	)
	roleKey := c.Param("roleKey")
	data, err := casbin.GetList(roleKey)
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "获取casbin列表失败")
		return
	}
	response.Success(c, data, "获取casbin列表成功")
}

func GetApi(c *gin.Context) {
	var (
		api model.Api
		err error
	)
	id, _ := strconv.Atoi(c.Param("id"))
	api.ID = id
	data, err := api.Get()
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "获取api数据出错")
		return
	}
	response.Success(c, data, "获取api数据成功")
}

func UpdateApi(c *gin.Context) {
	var (
		api model.Api
		err error
	)
	if err = c.ShouldBind(&api); err != nil {
		response.Error(c, response.CodeParamErr, nil, "参数错误")
		return
	}

	err = api.Update()
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

func DeleteApi(c *gin.Context) {
	var (
		api model.Api
		err error
	)
	ids := tools.IdStrToIdsGroup("id", c)
	err = api.Delete(ids)
	if err != nil {
		response.Error(c, response.CodeDBError, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

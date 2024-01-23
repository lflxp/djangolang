package service

import (
	"fmt"
	"strconv"

	"github.com/lflxp/djangolang/examples/allinone/demo/controller"
	"github.com/lflxp/djangolang/examples/allinone/demo/model"

	"github.com/gin-gonic/gin"
	log "github.com/go-eden/slf4go"
	"github.com/lflxp/djangolang/utils"
)

// @Summary  查询所有数据
// @Description 查询所有数据
// @Tags Demo
// @Param page query string false "页数"
// @Param pagesize query string false "每页大小"
// @Success 200 {object} []model.Demo "success"
// @Security ApiKeyAuth
// @Router /apis/v1/demo [get]
func GetAllDemo(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("pagesize")
	if page != "" && size != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			log.Error("数据转换错误", err.Error())
			utils.SendErrorMessage(c, 500, "DemoError", err.Error())
			return
		}
		pagesizeInt, err := strconv.Atoi(size)
		if err != nil {
			log.Error("数据转换错误", err.Error())
			utils.SendErrorMessage(c, 500, "DemoError", err.Error())
			return
		}
		data, err := controller.PageDemo(pageInt, pagesizeInt)
		if err != nil {
			log.Error("获取数据失败", err.Error())
			utils.SendErrorMessage(c, 500, "DemoError", err.Error())
			return
		}
		utils.SendSuccessMessage(c, 200, data)
		return
	}
	list, err := controller.ListDemo()
	if err != nil {
		log.Error("获取数据失败", err.Error())
		utils.SendErrorMessage(c, 500, "DemoError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, list)
}

// @Summary  根据id查询数据
// @Description 根据id查询数据
// @Tags Demo
// @Param id path string true "要删除的数据id"
// @Success 200 {object} model.Demo "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/demo [get]
func GetDemoById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "DemoGetByIdError", "id is empty")
		return
	}
	data, exist, err := controller.GetByUUIDDemo(id)
	if err != nil {
		log.Error("获取数据失败", err.Error())
		utils.SendErrorMessage(c, 500, "DemoError", err.Error())
		return
	}

	if !exist {
		log.Errorf("id %s 数据不存在", id)
		utils.SendErrorMessage(c, 500, "DemoError", fmt.Sprintf("id %s 数据不存在", id))
		return
	}

	utils.SendSuccessMessage(c, 200, data)
}

// @Summary  新增数据
// @Description 新增数据
// @Tags Demo
// @Param request body model.Demo true "工作空间数据 注意命名规范: Demo-ide-* | Demo-devfile-*"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/demo [post]
func CreateDemo(c *gin.Context) {
	var data model.Demo
	if err := c.BindJSON(&data); err != nil {
		log.Error(err.Error())
		utils.SendErrorMessage(c, 400, "DemoPostError", err.Error())
		return
	}

	num, err := controller.AddDemo(&data)
	if err != nil {
		utils.SendErrorMessage(c, 500, "DemoPostError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, fmt.Sprintf("success add %d", num))
}

// @Summary  删除数据
// @Description 删除数据
// @Tags Demo
// @Param id path string true "要删除的数据id"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/demo [delete]
func DeleteDemo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "DemoPutError", "id is empty")
		return
	}

	num, err := controller.DelDemo(id)
	if err != nil {
		utils.SendErrorMessage(c, 500, "DemoDeleteError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, fmt.Sprintf("success delete %d", num))
}

// @Summary  修改数据
// @Description 修改数据
// @Tags Demo
// @Param id path string true "数据id"
// @Param request body model.Demo true "修改数据"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/demo [put]
func PutDemo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "DemoPutError", "id is empty")
		return
	}
	var data model.Demo
	if err := c.BindJSON(&data); err != nil {
		log.Error(err.Error())
		utils.SendErrorMessage(c, 400, "DemoPutError", err.Error())
		return
	}

	num, err := controller.UpdateDemo(id, &data)
	if err != nil {
		utils.SendErrorMessage(c, 500, "DemoPutError", err.Error())
		return
	}
	utils.SendSuccessMessage(c, 200, fmt.Sprintf("update %d success", num))
}

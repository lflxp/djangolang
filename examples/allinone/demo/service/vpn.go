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
// @Tags Vpn
// @Param page query string false "页数"
// @Param pagesize query string false "每页大小"
// @Success 200 {object} []model.Vpn "success"
// @Security ApiKeyAuth
// @Router /apis/v1/vpn [get]
func GetAllVpn(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("pagesize")
	if page != "" && size != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			log.Error("数据转换错误", err.Error())
			utils.SendErrorMessage(c, 500, "VpnError", err.Error())
			return
		}
		pagesizeInt, err := strconv.Atoi(size)
		if err != nil {
			log.Error("数据转换错误", err.Error())
			utils.SendErrorMessage(c, 500, "VpnError", err.Error())
			return
		}
		data, err := controller.PageVpn(pageInt, pagesizeInt)
		if err != nil {
			log.Error("获取数据失败", err.Error())
			utils.SendErrorMessage(c, 500, "VpnError", err.Error())
			return
		}
		utils.SendSuccessMessage(c, 200, data)
		return
	}
	list, err := controller.ListVpn()
	if err != nil {
		log.Error("获取数据失败", err.Error())
		utils.SendErrorMessage(c, 500, "VpnError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, list)
}

// @Summary  根据id查询数据
// @Description 根据id查询数据
// @Tags Vpn
// @Param id path string true "要删除的数据id"
// @Success 200 {object} model.Vpn "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/vpn [get]
func GetVpnById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "VpnGetByIdError", "id is empty")
		return
	}
	data, exist, err := controller.GetByUUIDVpn(id)
	if err != nil {
		log.Error("获取数据失败", err.Error())
		utils.SendErrorMessage(c, 500, "VpnError", err.Error())
		return
	}

	if !exist {
		log.Errorf("id %s 数据不存在", id)
		utils.SendErrorMessage(c, 500, "VpnError", fmt.Sprintf("id %s 数据不存在", id))
		return
	}

	utils.SendSuccessMessage(c, 200, data)
}

// @Summary  新增数据
// @Description 新增数据
// @Tags Vpn
// @Param request body model.Vpn true "新增数据"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/vpn [post]
func CreateVpn(c *gin.Context) {
	var data model.Vpn
	if err := c.BindJSON(&data); err != nil {
		log.Error(err.Error())
		utils.SendErrorMessage(c, 400, "VpnPostError", err.Error())
		return
	}

	num, err := controller.AddVpn(&data)
	if err != nil {
		utils.SendErrorMessage(c, 500, "VpnPostError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, fmt.Sprintf("success add %d", num))
}

// @Summary  删除数据
// @Description 删除数据
// @Tags Vpn
// @Param id path string true "要删除的数据id"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/vpn [delete]
func DeleteVpn(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "VpnPutError", "id is empty")
		return
	}

	num, err := controller.DelVpn(id)
	if err != nil {
		utils.SendErrorMessage(c, 500, "VpnDeleteError", err.Error())
		return
	}

	utils.SendSuccessMessage(c, 200, fmt.Sprintf("success delete %d", num))
}

// @Summary  修改数据
// @Description 修改数据
// @Tags Vpn
// @Param id path string true "数据id"
// @Param request body model.Vpn true "修改数据"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /apis/v1/{id}/vpn [put]
func PutVpn(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.SendErrorMessage(c, 400, "VpnPutError", "id is empty")
		return
	}
	var data model.Vpn
	if err := c.BindJSON(&data); err != nil {
		log.Error(err.Error())
		utils.SendErrorMessage(c, 400, "VpnPutError", err.Error())
		return
	}

	num, err := controller.UpdateVpn(id, &data)
	if err != nil {
		utils.SendErrorMessage(c, 500, "VpnPutError", err.Error())
		return
	}
	utils.SendSuccessMessage(c, 200, fmt.Sprintf("update %d success", num))
}

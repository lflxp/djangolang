package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/lflxp/djangolang/middlewares"
	"github.com/lflxp/djangolang/model/test"
	"github.com/lflxp/djangolang/utils"

	"github.com/lflxp/djangolang/utils/orm/sqlite"

	"github.com/lflxp/djangolang"

	"github.com/gin-gonic/gin"
)

func init() {
	djangolang.RegisterAdmin(new(test.Demotest))
}

func Registertest(router *gin.Engine) {
	var authMiddleware = middlewares.NewGinJwtMiddlewares(middlewares.AllUserAuthorizator)
	testGroup := router.Group("/test")
	testGroup.Use(authMiddleware.MiddlewareFunc())
	{
		testGroup.Any("/:type", test_process)
	}
}

func test_process(c *gin.Context) {
	typed := c.Params.ByName("type")
	if c.Request.Method == "GET" {
		if typed == "index" {
			tmp := make([]test.Demotest, 0)
			err := sqlite.NewOrm().Limit(10, 0).Find(&tmp)
			if err != nil {
				c.String(400, err.Error())
			} else {
				c.HTML(200, "test/index.html", gin.H{
					"Data":     tmp,
					"Register": djangolang.GetRegistered(),
					"Other":    "other",
				})
			}
		} else if typed == "show" {
			id := c.Query("id")
			var tmp test.Demotest
			has, err := sqlite.NewOrm().Where("id = ?", id).Get(&tmp)
			if err != nil {
				c.String(400, err.Error())
			} else {
				if !has {
					c.String(404, "not found")
				} else {
					// c.HTML(200, "test/show.html", gin.H{
					// 	"Data": tmp,
					// })
					utils.SendSuccessMessage(c, 200, tmp)
				}
			}
		}
	} else if c.Request.Method == "POST" {
		// func Add(data *demo) (int64, error) {
		// 	affected, err := utils.Engine.Insert(data)
		// 	return affected, err
		// }
		if typed == "new" {
			var data test.Demotest
			info, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.String(500, err.Error())
				return
			}

			err = json.Unmarshal(info, &data)
			if err != nil {
				c.String(500, err.Error())
				return
			}

			num, err := test.Add(&data)
			if err != nil {
				c.JSONP(500, gin.H{
					"status": false,
					"msg":    err.Error(),
				})
			} else {
				c.JSONP(200, gin.H{
					"status":  true,
					"message": fmt.Sprintf("success add %d", num),
				})
			}
			// TODO: new
		} else if typed == "create" {
			// TODO: create
		}
	} else if c.Request.Method == "PUT" {
		// func Update(id string, data *demo) (int64, error) {
		// 	affected, err := utils.Engine.Table(new(demo)).ID(id).Update(data)
		// 	return affected, err
		// }

		if typed == "edit" {
			// TODO: edit
			id := c.Params.ByName("id")
			if id == "" {
				c.String(200, "id is none")
				return
			}

			var data test.Demotest
			info, _ := io.ReadAll(c.Request.Body)
			err := json.Unmarshal(info, &data)
			if err != nil {
				c.String(500, err.Error())
				return
			}

			num, err := test.Update(id, &data)
			if err != nil {
				c.JSONP(200, gin.H{
					"status": false,
					"msg":    err.Error(),
				})
			} else {
				c.JSONP(200, gin.H{
					"status":  true,
					"message": fmt.Sprintf("success put %d", num),
				})
			}
		}
	} else if c.Request.Method == "DELETE" {
		// func Del(id string) (int64, error) {
		// 	data := new(demo)
		// 	affected, err := utils.Engine.ID(id).Delete(data)
		// 	return affected, err
		// }
		if typed == "destroy" {
			// TODO: destroy
			id := c.Params.ByName("id")
			if id == "" {
				c.String(200, "id is none")
				return
			}

			num, err := test.Del(id)
			if err != nil {
				c.JSONP(200, gin.H{
					"status": false,
					"msg":    err.Error(),
				})
			} else {
				c.JSONP(200, gin.H{
					"status":  true,
					"message": fmt.Sprintf("success del %d", num),
				})
			}
		}
	} else if c.Request.Method == "PATCH" {
		if typed == "patch" {
			// TODO: patch
		}
	}
}

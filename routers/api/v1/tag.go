package v1

import (
	"github.com/gin-gonic/gin"
	"gin-blog/models"
	"github.com/Unknwon/com"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/pkg/setting"
	"net/http"
	"github.com/astaxie/beego/validation"
	"github.com/gpmgo/gopm/modules/log"
)

type Tag struct {
	models.Model
}

func GetTags(context *gin.Context) {
	name := context.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(context), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTags(context *gin.Context) {
	name := context.Query("name")

	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()

	createdBy := context.Query("created_by")

	valid := validation.Validation{}

	valid.Required(name, "name")
	valid.MaxSize(name, 100, "name").Message("名字不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	name := context.Query("name")
	modifiedBy := context.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1

	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy

			if name != "" {
				data[name] = name
			}

			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS

		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

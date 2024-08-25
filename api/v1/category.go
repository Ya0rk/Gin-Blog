package v1

import (
	"Gin-Blog/model"
	"Gin-Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
	Category模块的api接口， 对category表单进行处理
*/

// 添加分类
func AddCategory(c *gin.Context) {
	// 解析 json 并绑定在 data 中
	var cate model.Category
	if err := c.ShouldBindJSON(&cate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&cate)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
		// 并将data设置为空，避免返回其他用户的信息
		cate = model.Category{}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cate)
	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &cate)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

package v1

import (
	"MiMengCore/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Content struct {
	gorm.Model
	Type    string
	Content string
}

func GetNotice(c *gin.Context) {
	var content Content
	result := global.DB.Where("type = ?", "notice").First(&content)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "获取公告失败"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "获取公告成功", "data": content})
}

func UpdateNotice(c *gin.Context) {
	var content Content
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求数据绑定失败"})
		return
	}
	// 使用Where条件来查找type为"notice"的记录
	result := global.DB.Model(&content).Where("type = ?", "notice").Updates(Content{Type: "notice", Content: content.Content})
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "更新公告失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "更新公告成功"})
}

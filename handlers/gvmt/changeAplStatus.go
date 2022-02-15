package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary      修改申诉列表状态
// @Description  修改申诉列表状态
// @Tags         gvmt(admin)
// @Accept       json
// @Produce      json
// @Param         object body models.AplStatusDto true "status=1表示不通过,status=2表示等待,status=3表示通过"
// @Success 200 object models.SuccessResponse
//@Router       /gvmt/changeAplStatus [post]
func ChangeAplStatus(c *gin.Context) {
	postDto := &models.AplStatusDto{}
	err := c.BindJSON(postDto)
	if err != nil {
		log.Println(err)
		return
	}

	post := &models.Appeal{}
	drivers.MysqlDb.Table("appeal").Where("id = ?", postDto.Id).First(post)
	post.Status = postDto.Status
	drivers.MysqlDb.Table("appeal").Save(post)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}

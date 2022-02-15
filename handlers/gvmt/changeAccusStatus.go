package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary      修改待审核的排污实地图列表状态
// @Description  修改待审核的排污实地图列表状态
// @Tags         gvmt(admin)
// @Accept       json
// @Produce      json
// @Param         object body models.AccusStatusDto true "status=1表示不通过,status=2表示等待,status=3表示通过"
// @Success 200 object models.SuccessResponse
//@Router       /gvmt/changeAccusStatus [post]
func ChangeAccusStatus(c *gin.Context) {
	postDto := &models.AccusStatusDto{}
	err := c.BindJSON(postDto)
	if err != nil {
		log.Println(err)
		return
	}

	post := &models.Accusation{}
	drivers.MysqlDb.Table("accusation").Where("id = ?", postDto.Id).First(post)
	post.Status = postDto.Status
	drivers.MysqlDb.Table("accusation").Save(post)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}

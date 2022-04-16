package handlers

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
)

// @Summary      获得企业的port_id和polution_id
// @Description  获得企业的port_id和polution_id
// @Accept       json
// @Produce      json
// @Param         object body models.InfoReq true "企业id"
// @Success 200 object models.InfoResp
// @Router       /info [post]
func InfoHandler(c *gin.Context) {
	var req models.InfoReq
	//fmt.Println("info")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "error!!"})
		return
	}

	var alldatas []models.AllData
	//time1, _ = time.ParseInLocation(layout, req.Date1, loc)
	//fmt.Println("----------------------------------------------")
	drivers.MysqlDb.Table("all_data").Where("company_id = ?", req.CompanyId).Find(&alldatas)
	//fmt.Println("----------------------------------------------")
	//fmt.Println("req.CompanyId: ", req.CompanyId)

	//fmt.Println(len(alldatas))
	//for i := range alldatas {
	//	fmt.Println(alldatas[i].CompanyId, alldatas[i].PolutionId, alldatas[i].PortId, alldatas[i].Date)
	//}

	var resp models.InfoResp

	respPortIdSet := map[string]bool{}
	respPolutionIdSet := map[string]bool{}

	for i := range alldatas {
		respPortIdSet[alldatas[i].PortId] = true
		respPolutionIdSet[alldatas[i].PolutionId] = true
	}

	for portId := range respPortIdSet {
		resp.PortIds = append(resp.PortIds, portId)
	}
	for polutionId := range respPolutionIdSet {
		resp.PolutionIds = append(resp.PolutionIds, polutionId)
	}

	//time.Now().Format("1/2/2006 15:04:05")
	//fmt.Println(time.Now().Format("02/01/2006 15:04:05"))

	c.JSON(200, resp)

}

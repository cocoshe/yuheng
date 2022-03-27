package routers

import (
	_ "backend/docs"
	"backend/handlers"
	"backend/handlers/company"
	"backend/handlers/gvmt"
	"backend/handlers/visitor"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {

	r.POST("/display", handlers.DisplayHandler)
	r.POST("/displayBad", handlers.DisplayBadHandler)
	r.POST("/rank", handlers.DisplayRank)
	r.POST("/insight", handlers.InsightHandler)
	r.POST("/addAtten", handlers.AddAttenHandler)
	r.POST("/delAtten", handlers.DelAttenHandler)
	r.GET("/attenInfo", handlers.AttenInfoHandler)
	r.POST("/run", handlers.RunHandler) // export GOROOT=/d/sdk/go1.17.2
	r.POST("/info", handlers.InfoHandler)
	r.POST("/selfcheck", handlers.SelfCheckHandler)
	r.GET("/getFeatures", handlers.GetFeaturesHandler)

	r.GET("/index", handlers.IndexHandler)
	r.GET("/ping", handlers.PingHandler)
	r.POST("/login", handlers.LoginHandler)
	gvmtGroup := r.Group("/gvmt").Use(middleware.GvmtMiddleware())
	{
		gvmtGroup.GET("/getAppealList", gvmt.GetAppealList)
		gvmtGroup.POST("/changeStatus", gvmt.ChangeStatus)
		gvmtGroup.POST("/changeAccusStatus", gvmt.ChangeAccusStatus)
		gvmtGroup.POST("/changeAplStatus", gvmt.ChangeAplStatus)
		gvmtGroup.POST("/changeThreshold", gvmt.ChangeThreshold)
		// 查看群众上传的举报视图
		gvmtGroup.GET("/getAccusationList", gvmt.GetAccusationList)
	}
	visitorGroup := r.Group("/visitor").Use(middleware.VisitorMiddleware())
	{
		// 举报历史视图
		visitorGroup.GET("/history", visitor.History)
		// 举报信息上传
		visitorGroup.POST("/upload", visitor.Upload)
	}
	companyGroup := r.Group("/company").Use(middleware.CompanyMiddleware())
	{
		// 公司申诉
		companyGroup.POST("/appeal", company.Appeal)
		companyGroup.GET("/history", company.History)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

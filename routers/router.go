package routers

import (
	"backend/handlers"
	"backend/handlers/company"
	"backend/handlers/gvmt"
	"backend/handlers/visitor"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/index", handlers.IndexHandler)
	r.GET("/ping", handlers.PingHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/info", middleware.GvmtMiddleware(), handlers.Info)
	r.POST("/test", handlers.TestHandler)
	gvmtGroup := r.Group("/gvmt").Use(middleware.GvmtMiddleware())
	{
		// 审核算法输出视图
		gvmtGroup.GET("/getWarningList", gvmt.GetWarningList)
		gvmtGroup.GET("/appeal", gvmt.Appeal)
		gvmtGroup.POST("/changeStatus", gvmt.ChangeStatus)

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
	}

}

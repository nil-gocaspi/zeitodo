package routes

import (
	"example-project/model"
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HandlerInterface
type HandlerInterface interface {
	GetUserHandler(c *gin.Context)
	GetAllUserHandler(c *gin.Context)
	CreateUserHandler(c *gin.Context)
	GetTeamMemberHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	PermissionMiddleware(c *gin.Context)
	LoginUserHandler(c *gin.Context)
	//	LogoutUserHandler(c *gin.Context)
	RefreshTokenHandler(c *gin.Context)
	GetProposalsById(c *gin.Context)
	CreateProposalsHandler(c *gin.Context)
	DeleteProposalHandler(c *gin.Context)
	UpdateProposalsHandler(c *gin.Context)
	CreatTimeEntry(c *gin.Context)
	UpdateTimeEntry(c *gin.Context)
	GetTimeEntry(c *gin.Context)
	DeleteTimeEntry(c *gin.Context)
	GetAllTimeEntry(c *gin.Context)
	CalcultimeEntry(c *gin.Context)
	GetUserIdHandler(c *gin.Context)
}

var Handler HandlerInterface
var PermissionList model.PermissionList

func CreateRoutes(group *gin.RouterGroup) {
	PermissionList.Permissions = append(PermissionList.Permissions, model.Permission{Uri: "/user/", Methods: []string{"GET", "OPTIONS"}, GetSameUser: true, Group: "user"})

	PermissionList.Permissions = append(PermissionList.Permissions, model.Permission{Uri: "/user/", Methods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, GetSameUser: false, Group: "admin"})

	group.POST("/login", Handler.LoginUserHandler)
	//	group.POST("/logout", Handler.LogoutUserHandler)
	group.POST("/refresh", Handler.RefreshTokenHandler)
	group.GET("/uId/:username", Handler.GetUserIdHandler)
	user := group.Group("/user")

	user.GET("/:id", Handler.PermissionMiddleware, Handler.GetUserHandler)
	user.GET("/", Handler.PermissionMiddleware, Handler.GetAllUserHandler)
	user.POST("/", Handler.CreateUserHandler)
	user.GET("/team", Handler.PermissionMiddleware, Handler.GetTeamMemberHandler)
	user.PATCH("/", Handler.PermissionMiddleware, Handler.UpdateUserHandler)
	user.DELETE("/:id", Handler.PermissionMiddleware, Handler.DeleteUserHandler)
	route := group.Group("/proposals")
	//	route.Use(CORS)
	route.GET("/:id", Handler.GetProposalsById)
	route.POST("/:id", Handler.CreateProposalsHandler)
	route.DELETE("/:id", Handler.DeleteProposalHandler)
	route.PATCH("/", Handler.UpdateProposalsHandler)
	timeentry := group.Group("/timeentry")
	timeentry.POST("/createtime", Handler.CreatTimeEntry)
	timeentry.PUT("/:id/update", Handler.UpdateTimeEntry)
	timeentry.GET("/:id/gettime", Handler.GetTimeEntry)
	timeentry.DELETE("/:id/delete", Handler.DeleteTimeEntry)
	timeentry.GET("/", Handler.GetAllTimeEntry)
	timeentry.GET("/:id/calcul", Handler.CalcultimeEntry)
}

/*
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

*/

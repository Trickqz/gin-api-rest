package routes

import (
	"github.com/trickqz/api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAllStudents)
	r.POST("/alunos", controllers.CreateStudent)
	r.GET("/alunos/:id", controllers.GetStudentById)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)
	r.GET("/alunos/cpf/:cpf", controllers.GetStudentByCPF)
	r.Run()
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hrqcds/api-go-gin/controllers"
)

var R *gin.Engine

func HandleRequests() {

	R = gin.Default()

	R.GET("/alunos", controllers.AllAlunos)
	R.GET("/alunos/:id", controllers.GetAlunoById)
	R.GET("/alunos/cpf/:cpf", controllers.GetAlunoByCPF)
	R.POST("/alunos", controllers.CriarNovoAluno)
	R.DELETE("/alunos/:id", controllers.DeleterAluno)
	R.PATCH("/alunos/:id", controllers.UpdateAluno)

	R.Run(":8000")
}

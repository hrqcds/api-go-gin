package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrqcds/api-go-gin/database"
	"github.com/hrqcds/api-go-gin/models"
)

func AllAlunos(c *gin.Context) {

	alunos := models.BuscarAlunos(database.DB)

	c.JSON(200, alunos)
}

func GetAlunoById(c *gin.Context) {
	id := c.Params.ByName("id")

	aluno := models.BuscarAlunoPorId(database.DB, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	// Buscando no body
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	models.CriarAluno(database.DB, &aluno)

	c.JSON(http.StatusOK, aluno)

}

func DeleterAluno(c *gin.Context) {
	id := c.Params.ByName("id")

	// var aluno models.Aluno

	// database.DB.First(&aluno, id)
	// database.DB.Delete(&aluno)

	aluno := models.DeletarAluno(database.DB, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Aluno não encontrado",
		})

		return

	}

	c.JSON(http.StatusOK, aluno)
}

func UpdateAluno(c *gin.Context) {

	id := c.Params.ByName("id")

	aluno := models.BuscarAlunoPorId(database.DB, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
			"status":  "Not Found",
		})
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "Not Found",
		})

		return
	}

	models.UpdateAluno(database.DB, &aluno)

	c.JSON(http.StatusOK, aluno)

}

func GetAlunoByCPF(c *gin.Context) {

	cpf := c.Param("cpf")

	aluno := models.GetAlunoByCPF(database.DB, cpf)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

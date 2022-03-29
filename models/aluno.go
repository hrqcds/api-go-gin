package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

func BuscarAlunos(db *gorm.DB) []Aluno {

	var alunos []Aluno

	db.Find(&alunos)

	return alunos

}

func BuscarAlunoPorId(db *gorm.DB, id string) Aluno {

	var aluno Aluno

	db.First(&aluno, id)

	return aluno

}

func CriarAluno(db *gorm.DB, a *Aluno) {

	db.Create(&a)

}

func DeletarAluno(db *gorm.DB, id string) Aluno {

	var aluno Aluno

	db.First(&aluno, id)

	if aluno.ID == 0 {
		return aluno
	}

	db.Delete(&aluno)

	return aluno

}

func UpdateAluno(db *gorm.DB, update *Aluno) {

	db.Save(&update)

}

func GetAlunoByCPF(db *gorm.DB, cpf string) Aluno {

	var aluno Aluno

	db.Where(Aluno{CPF: cpf}).First(&aluno)

	return aluno
}

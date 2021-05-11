package template

var RepoTemplate = `package repository

import (
	"{{.PackageName}}/internal/app/{{.LowerName}}/model"

	"gorm.io/gorm"
)

const (
	invalidID = -1
)

//{{.UperName}}Repo {{.UperName}}Repo
type {{.UperName}}Repo interface {
	GetByID(id int64) (*{{.UperName}}, error)
	Create(m *{{.UperName}}) (id int64, err error)
	DeleteByID(id int64) error
	Updates(id int64, data map[string]interface{}) (rowsAffected int64, err error)
	Save(m *{{.UperName}}) (rowsAffected int64, err error)
	GetList(offset, limit int) (results []*{{.UperName}}, err error)
	Count() (num int64, err error)
}

type {{.LowerName}}Repo struct {
	db *gorm.DB
}

//New{{.UperName}}Repo New{{.UperName}}Repo
func New{{.UperName}}Repo(db *gorm.DB) {{.UperName}}Repo {
	return &{{.LowerName}}Repo{db: db}
}

func (repo *{{.LowerName}}Repo) GetByID(id int64) (*{{.UperName}}, error) {
	var result {{.UperName}}
	if err := repo.db.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *{{.LowerName}}Repo) Create(m *{{.UperName}}) (id int64, err error) {
	db := repo.db.Create(m)
	if err = db.Error; err != nil {
		return invalidID, err
	}
	return m.ID, nil
}

func (repo *{{.LowerName}}Repo) DeleteByID(id int64) error {
	return repo.db.Delete({{.UperName}}{}, "id = ?", id).Error
}

func (repo *{{.LowerName}}Repo) Updates(id int64, data map[string]interface{}) (rowsAffected int64, err error) {
	db := repo.db.Model({{.UperName}}{}).Where("id = ?", id).Updates(data)
	return db.RowsAffected, db.Error
}

func (repo *{{.LowerName}}Repo) Save(m *{{.UperName}}) (rowsAffected int64, err error) {
	db := repo.db.Save(m)
	return db.RowsAffected, db.Error
}

func (repo *{{.LowerName}}Repo) GetList(offset, limit int) (results []*{{.UperName}}, err error) {
	if err := repo.db.Limit(limit).Offset(offset).Find(&results).Error; err != nil {
		return nil, err
	}
	return
}

func (repo *{{.LowerName}}Repo) Count() (num int64, err error) {
	if err = repo.db.Model({{.UperName}}{}).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}
`

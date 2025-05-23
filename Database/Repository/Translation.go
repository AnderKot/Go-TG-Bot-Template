package Repository

import (
	Table "Bot/Database/Repository/Table"
	Interface "Bot/Interface"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"gorm.io/gorm"
)

type Translation struct {
	database *gorm.DB
}

func NewTemplateRepository(_database *gorm.DB) Translation {
	return Translation{
		database: _database,
	}
}

func (r Translation) Init() Translation {
	r.database.AutoMigrate(&Table.Translation{})

	file, err := os.Open(filepath.Join("", "templates.csv"))
	if err == nil {
		defer file.Close()
		var templates []Table.Translation
		err = gocsv.Unmarshal(file, &templates)
		if err != nil {
			panic(err)
		}
		r.database.Create(templates)
	}
	return r
}

func (r *Translation) GetTemplateText(template Interface.ITemplate, langCode string) string {
	if template.IsTranslated() {
		templates := new(Table.Translation)
		templates.Code = template.GetTemplateCode()
		templates.LangCode = langCode
		r.database.Where(&templates).Take(templates)
		if templates.Text != "" {
			return templates.Text
		}
		templates.LangCode = "en"
		r.database.Where(&templates).Take(templates)
		if templates.Text == "" {
			templates.Text = "Template name not specificate"
			r.database.Save(templates)
			templates.LangCode = langCode
			r.database.Save(templates)
		}
		return templates.Text
	} else {
		return template.GetTemplateText()
	}
}

func (r *Translation) ExportTemplates() {

}

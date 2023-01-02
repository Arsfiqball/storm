package book

import (
	"app/pkg/restpl"
	"database/sql"
	"time"
)

type QueryBook struct {
	//
}

func (q *QueryBook) GetSqlWhereStatement() string {
	return ""
}

func (q *QueryBook) GetSqlWhereValues() []interface{} {
	var values []interface{}
	return values
}

type Mutation []string

type PayloadBook struct {
	restpl.Mutation

	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Series      string    `json:"series"`
	Volume      int       `json:"volume"`
	FileUrl     string    `json:"fileUrl"`
	CoverUrl    string    `json:"coverUrl"`
	PublishDate time.Time `json:"publishDate"`
}

func (p *PayloadBook) GetSqlSelectFields() []interface{} {
	fields := []interface{}{}

	if p.Mutation.HasField("title") {
		fields = append(fields, "Title")
	}

	if p.Mutation.HasField("author") {
		fields = append(fields, "Author")
	}

	if p.Mutation.HasField("series") {
		fields = append(fields, "Series")
	}

	if p.Mutation.HasField("volume") {
		fields = append(fields, "Volume")
	}

	if p.Mutation.HasField("fileUrl") {
		fields = append(fields, "FileUrl")
	}

	if p.Mutation.HasField("coverUrl") {
		fields = append(fields, "CoverUrl")
	}

	if p.Mutation.HasField("publishDate") {
		fields = append(fields, "PublishDate")
	}

	return fields
}

func (p *PayloadBook) ToEntity() EntityBook {
	var entity EntityBook

	if p.Mutation.HasField("title") {
		entity.Title = p.Title
	}

	if p.Mutation.HasField("author") {
		entity.Author = p.Author
	}

	if p.Mutation.HasField("series") {
		entity.Series = p.Series
	}

	if p.Mutation.HasField("volume") {
		entity.Volume = p.Volume
	}

	if p.Mutation.HasField("fileUrl") {
		fileUrl := sql.NullString{String: p.FileUrl, Valid: true}

		if p.Mutation.IsNullField("fileUrl") {
			fileUrl = sql.NullString{}
		}

		entity.FileUrl = fileUrl
	}

	if p.Mutation.HasField("coverUrl") {
		coverUrl := sql.NullString{String: p.CoverUrl, Valid: true}

		if p.Mutation.IsNullField("coverUrl") {
			coverUrl = sql.NullString{}
		}

		entity.CoverUrl = coverUrl
	}

	if p.Mutation.HasField("publishDate") {
		publishDate := sql.NullTime{Time: p.PublishDate, Valid: true}

		if p.Mutation.IsNullField("publishDate") {
			publishDate = sql.NullTime{}
		}

		entity.PublishDate = publishDate
	}

	return entity
}

type EntityBook struct {
	ID          uint           `gorm:"column:id"`
	Title       string         `gorm:"column:title"`
	Author      string         `gorm:"column:author"`
	Series      string         `gorm:"column:series"`
	Volume      int            `gorm:"column:volume"`
	FileUrl     sql.NullString `gorm:"column:file_url"`
	CoverUrl    sql.NullString `gorm:"column:cover_url"`
	PublishDate sql.NullTime   `gorm:"column:publish_date"`
}

func (EntityBook) TableName() string {
	return "books"
}

type AggregateBook struct {
	Docs []EntityBook
}

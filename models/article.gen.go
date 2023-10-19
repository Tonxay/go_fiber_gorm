// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AuthorID int32  `gorm:"column:author_id;not null" json:"author_id"`
	Title    string `gorm:"column:title;not null" json:"title"`
	Content  string `gorm:"column:content;not null" json:"content"`
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}
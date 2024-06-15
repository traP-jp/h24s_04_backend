package model


import (
	"database/sql"
	"time"
)

// image.pngを参考にslides structとgenre structを書いて欲しいです(nullとかも反映させてくださると:tasukaru:)

type Slide struct {
	Id          string         `json:"id,omitempty"  db:"id"`
	DL_url      string         `json:"dl_url,omitempty"  db:"dl_url"`
	Thumb_url   sql.NullString `json:"thumb_url,omitempty"  db:"thumb"`
	Title       string         `json:"title,omitempty"  db:"title"`
	Genre_id    string         `json:"genre_id,omitempty"  db:"genre_id"`
	Posted_at   time.Time      `json:"posted_at,omitempty"  db:"posted_at"`
	Description sql.NullString `json:"description,omitempty"  db:"description"`
}

type Genre struct {
	Id        string `json:"id,omitempty"  db:"id"`
	Genrename string `json:"genrename,omitempty"  db:"genrename"`
}

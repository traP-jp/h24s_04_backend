package model


import (
	"database/sql"
	"time"
)

// image.pngを参考にslides structとgenre structを書いて欲しいです(nullとかも反映させてくださると:tasukaru:)

type Slide struct {
	id          string         `json:"id,omitempty"  db:"id"`
	dl_url      string         `json:"dl_url,omitempty"  db:"dl_url"`
	thumb_url   sql.NullString `json:"thumb_url,omitempty"  db:"thumb"`
	title       string         `json:"title,omitempty"  db:"title"`
	genre_id    string         `json:"genre_id,omitempty"  db:"genre_id"`
	posted_at   time.Time      `json:"posted_at,omitempty"  db:"posted_at"`
	description sql.NullString `json:"description,omitempty"  db:"description"`
}

type Genre struct {
	id        string `json:"id,omitempty"  db:"id"`
	genrename string `json:"genrename,omitempty"  db:"genrename"`
}

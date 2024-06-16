package model

import (
	"time"

	"github.com/gofrs/uuid"
)

// image.pngを参考にslides structとgenre structを書いて欲しいです(nullとかも反映させてくださると:tasukaru:)

type Slide struct {
	Id             uuid.UUID `json:"id,omitempty"  db:"id"`
	DL_url         string    `json:"dl_url,omitempty"  db:"dl_url"`
	Thumb_url      string    `json:"thumb_url,omitempty"  db:"thumb_url"`
	Title          string    `json:"title,omitempty"  db:"title"`
	Genre_id       uuid.UUID `json:"genre_id,omitempty"  db:"genre_id"`
	Posted_at      time.Time `json:"posted_at,omitempty"  db:"posted_at"`
	Description    string    `json:"description,omitempty"  db:"description"`
	URL_updated_at time.Time `json:"url_updated_at,omitempty"  db:"url_updated_at"`
	Filepath       string    `json:"filepath,omitempty"  db:"filepath"`
}

type Genre struct {
	Id         uuid.UUID `json:"id,omitempty"  db:"id"`
	Genre_name string    `json:"genrename,omitempty"  db:"genrename"`
}

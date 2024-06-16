package transfer

import (
	"time"

	"log"
)

func (s *transferFileService) Urlupdate() {
	var paths []string

	err := s.db.Select(&paths, "SELECT `filepath` FROM `Slide` WHERE `url_updated_at` < TIMESTAMPADD(HOUR,-144,UTC_TIMESTAMP()) ")
	if err != nil {
		log.Printf("failed to get url should be updated from DB: %v", err)
		return
	}
	for _, path := range paths {
		filepath := "files/" + path
		thumbpath := "files/thumb_" + path
		newdlurl, err := s.uu.UpdateURL(filepath)
		if err != nil {
			log.Printf("failed to update dlurl from Firebase Storage: %v", err)
			return
		}
		newthumburl, err := s.uu.UpdateURL(thumbpath)
		if err != nil {
			log.Printf("failed to update thumburl from Firebase Storage: %v", err)
			return
		}
		_, err = s.db.Exec("Update `Slide` SET `dl_url`=?,`thumb_url`=?,`url_updated_at`=? WHERE `filepath`=?", newdlurl, newthumburl, time.Now(), path)
		if err != nil {
			log.Printf("failed to insert newurl to DB: %v", err)
			return
		}
	}
}

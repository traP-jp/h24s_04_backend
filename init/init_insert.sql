USE h24s_04

INSERT INTO Genre (id, genrename) VALUES ("5ebb3054-7a0b-ad30-d02c-4bbf1f11d48e","noya2");
INSERT INTO Slide (id, dl_url, thumb_url, title, genre_id, posted_at, description) VALUES ("628","dl_url628","thumb_url628","noya2-slide","5ebb3054-7a0b-ad30-d02c-4bbf1f11d48e", NOW(),"traq-id");

INSERT INTO Genre (id, genrename) VALUES ("52821c98-2aba-11ef-82c2-0242ac130003","kayama");
INSERT INTO Slide (id, dl_url, thumb_url, title, genre_id, posted_at, description) VALUES (uuid(),"dl_url65536","thumb_url65536","kayama-slide","52821c98-2aba-11ef-82c2-0242ac130003", NOW(),"kayamakayamakayama");

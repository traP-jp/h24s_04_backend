USE h24s_04
INSERT INTO
  Genre (id, genrename)
VALUES
  ("5ebb3054-7a0b-ad30-d02c-4bbf1f11d48e", "noya2");

INSERT INTO
  Slide (
    id,
    dl_url,
    thumb_url,
    title,
    genre_id,
    posted_at,
    description,
    url_lastupdated
  )
VALUES
  (
    "66bbba55-f0db-5d0a-ff40-5f7ca24b9ece",
    "dl_url628",
    "thumb_url628",
    "noya2-slide",
    "5ebb3054-7a0b-ad30-d02c-4bbf1f11d48e",
    NOW (),
    "traq-id",
    NOW ()
  );

INSERT INTO
  Genre (id, genrename)
VALUES
  ("52821c98-2aba-11ef-82c2-0242ac130003", "kayama");

INSERT INTO
  Slide (
    id,
    dl_url,
    thumb_url,
    title,
    genre_id,
    posted_at,
    description,
    url_lastupdated
  )
VALUES
  (
    "e44eca0e-f3fc-0f50-37dc-adcd89c37334",
    "dl_url65536",
    "thumb_url65536",
    "kayama-slide",
    "52821c98-2aba-11ef-82c2-0242ac130003",
    NOW (),
    "kayamakayamakayama",
    NOW ()
  );

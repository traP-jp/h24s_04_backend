# h24s_04_backend

## api

### endpoints

- slides
  - GET:スライドの取得
    - queryparam:ジャンル毎取得
    - respose:
      - `slide`の配列を返す
  - POST:スライドのアップロード
- slides/{slideid}
  - GET:スライドidでデータ取得
    - response:
      - `slide`を返す
  - PATCH:スライドデータの修正(再up,ジャンルなど)
    - require:
      - title
      - genreid
    - respose:
      - 編集した`slide`を返す
  - DELETE:該当スライドの削除(ストレージからの削除も行う)
    - respose:
      - 削除した`slide`を返す
- genres
  - GET:ジャンル全取得
    - `genre`の配列を返す
  - POST:ジャンルの登録
    - require:
      - genrename
    - 登録したジャンルの`genre`を返す
- genres/{genreid}
  - PATCH:ジャンルデータの修正(ジャンル名の変更)
    - require:
      - genrename
  - DELETE:ジャンルの削除
    - 削除したジャンルの`genre`を返す

### schema :done:

- slide
  - uuid
    - string
    - format:uuid
  - dlurl
    - string
  - thumburl
    - stringNULL
  - title
    - string
  - genreid
    - string
    - format:uuid
  - posted_at
    - string
    - format:timestamp
  - description
    - stringNULL

- genre
  - uuid
    - string
    - format:uuid
  - name
    - string

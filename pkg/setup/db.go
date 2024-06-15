package setup

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBsetup() *sqlx.DB {

	log.Println(":::Start DBsetup:::")

	NS, def := os.LookupEnv("NS")

	if !def {
		NS = "false"
	}
	var user, pass, addr, host, port, dbname string

	if NS == "true" {
		user, def = os.LookupEnv("NS_MARIADB_USER")
		if !def {
			user = "root"
		}
		pass, def = os.LookupEnv("NS_MARIADB_PASSWORD")
		if !def {
			pass = "password"
		}
		host, def = os.LookupEnv("NS_MARIADB_HOSTNAME")
		if !def {
			host = "localhost"
		}
		port, def = os.LookupEnv("NS_MARIADB_PORT")
		if !def {
			port = "3306"
		}
		addr = host + ":" + port
		dbname, def = os.LookupEnv("NS_MARIADB_DATABASE")
		if !def {
			dbname = "h24s_04"
		}
	} else {
		user, def = os.LookupEnv("LOCAL_MYSQL_USER")
		if !def {
			user = "root"
		}
		pass, def = os.LookupEnv("LOCAL_MYSQL_PASSWORD")
		if !def {
			pass = "password"
		}
		host, def = os.LookupEnv("LOCAL_MYSQL_HOSTNAME")
		if !def {
			host = "localhost"
		}
		port, def = os.LookupEnv("LOCAL_MYSQL_PORT")
		if !def {
			port = "3306"
		}
		addr = host + ":" + port

		dbname, def = os.LookupEnv("LOCAL_MYSQL_DATABASE")
		if !def {
			dbname = "h24s_04"
		}
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("ERROR_DBsetup:", err)
	}

	conf := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal("ERROR_DBsetup:", err)
	}
	err = _db.Ping()
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	log.Println(":::Complete DBsetup:::")

	return _db

}

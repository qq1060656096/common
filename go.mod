module github.com/qq1060656096/go-common

go 1.13

require (
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/qq1060656096/go-gorm-manager v0.0.0-20200304033243-c7cddf14a207
	github.com/qq1060656096/go-redis-manager v0.0.0-20200304033259-e6c809ef5d95
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/qq1060656096/go-gorm-manager v0.0.0-20200304033243-c7cddf14a207 => ../go-gorm-manager
	github.com/qq1060656096/go-redis-manager v0.0.0-20200304033259-e6c809ef5d95 => ../go-redis-manager
)

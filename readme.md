## blogbackend
use GoFiber/v2
#### setp1
```
go mod init github.com/wz/blogbackend
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
go get -u github.com/dgrijalva/jwt-go
```
#### step2
create .env
```
DSN="root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
```

#### setp3
create ./database/connect.go
```
package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = database
}

```
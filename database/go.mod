module github.com/better-go/pkg/database

go 1.18

//replace github.com/better-go/pkg/time => ../time
//
//replace github.com/better-go/pkg/log => ../log
//
//replace github.com/better-go/pkg/x/go-zero => ../x/go-zero

require (
	github.com/better-go/pkg/log v0.0.0-20220923023940-c922e8210ef0
	github.com/better-go/pkg/time v0.0.0-20220923022650-d97906983f30
	github.com/better-go/pkg/x/go-zero v0.0.0-20220923023940-c922e8210ef0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.10
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220919173607-35f4265a4bc0 // indirect
)

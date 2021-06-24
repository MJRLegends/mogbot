module github.com/ChrisMcDearman/mogbot

go 1.16

require (
	github.com/bwmarrin/discordgo v0.23.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jackc/pgx/v4 v4.11.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.0 // indirect
	github.com/mattn/go-shellwords v1.0.11
	golang.org/x/text v0.3.6 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.8
)

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210619182440-23a197a98005

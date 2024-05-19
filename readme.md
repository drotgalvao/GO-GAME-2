docker-compose down
docker system prune -v


docker-compose up -d


psql -U admin -d go_game
psql -U postgres -d go-game

psql -h localhost -p 5432 -U admin -d go_game -W

\du
\l

\c go-game


\?  -- \q

\dt

\d todos;

docker system prune -v

docker system prune -a -v

package db

import (
	"database/sql"
	"fmt"

	"github.com/drotgalvao/GO-GAME/configs"
	_ "github.com/lib/pq"
)

func buildConnectionString(conf *configs.DBConfig) string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)
}

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := buildConnectionString(&conf)

	fmt.Println(sc)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

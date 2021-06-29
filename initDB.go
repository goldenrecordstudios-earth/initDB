package initDB

import (
    "database/sql"
    "time"
    "net"
    "log"

    "github.com/go-sql-driver/mysql"
)

// https://github.com/go-sql-driver/mysql/blob/v1.6.0/dsn.go#L36

func init(dbUser, dbPass, dbHost, dbPort, dbName, dbType string) (*sql.DB, error) {
    db := &sql.DB{}

    DSNConfig := mysql.NewConfig()
    DSNConfig.User    = dbUser
    DSNConfig.Passwd  = dbPass
    DSNConfig.Net     = "tcp"
    DSNConfig.Addr    = net.JoinHostPort(dbHost, dbPort)
    DSNConfig.DBName  = dbName

    DSNStr := DSNConfig.FormatDSN()

    parsedDSN, err := mysql.ParseDSN(DSNStr)
    if err != nil {
        log.Println(parsedDSN)
        log.Println(err)
        return nil, err
    }

    db, err = sql.Open(dbType, DSNStr)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    err = db.Ping()
    if err != nil {
        log.Println(err)
        return nil, err
    }

    log.Println("db connection established")

    return db, nil
}
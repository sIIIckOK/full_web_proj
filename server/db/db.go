package db

import (
	"database/sql"

	tp "github.com/siiickok/full-web-proj/types"
)

type PostgresDB struct {
    db *sql.DB
}

func PostgresOpen(dbString string) (PostgresDB, error) {
    pq, err := sql.Open("postgres", dbString)
    return PostgresDB{
        db: pq,
    }, err
}

func (p PostgresDB) Close() error {
    return p.db.Close()
}

func (p PostgresDB) Ping() error {
    return p.db.Ping()
}

func (p PostgresDB) Init() error {
    query :=
    `
    CREATE TABLE IF NOT EXISTS item(
        id SERIAL PRIMARY KEY,
        title TEXT,
        price DECIMAL)
    `
    _, err := p.db.Exec(query)
    return err
}

func (p PostgresDB) GetItem(id int) (tp.Item, error) {
    query := 
    `
    SELECT title, price
        FROM item
            WHERE id = $1
    `
    var it tp.Item
    it.Id = id
    if err := p.db.QueryRow(query, id).Scan(&it.Title, &it.Price); err != nil {
        return it, err
    }
    return it, nil
}

func (p PostgresDB) GetItems(amount int) ([]tp.Item, error) {
    query := 
    `
    SELECT (id, title, price)
        FROM item
    `
    rows, err := p.db.Query(query)
    if err != nil { return nil, err }
    
    var its []tp.Item
    var it tp.Item

    i := amount
    for rows.Next() && i > 0 {
        if err := rows.Err(); err != nil { return nil, err }
        rows.Scan(&it.Id, &it.Title, &it.Price)
        its = append(its, it)
        i--
    }
    return its, nil
}

func (p PostgresDB) PostItem(item tp.Item) error {
    query := 
    `
    INSERT INTO item (title, price)
        VALUES ($1, $2)
    `
    _, err := p.db.Exec(query, item.Title, item.Price)
    return err
}




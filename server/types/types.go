package types

type DB interface {
    Close() error 
    Ping()  error

    Init() error

    GetItem(id int) (Item, error)
    GetItems(number int) ([]Item, error)
    PostItem(Item) error
}

type Item struct {
    Id    int     `json:"id"` 
    Title string  `json:"title"` 
    Price float32 `json:"price"` 
}



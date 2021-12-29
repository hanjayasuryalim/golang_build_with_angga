package entity

type Store struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
}

func (s Store) StockStatus() string {
	if s.Stock < 5 {
		return "Stock hampir habis"
	} else if s.Stock < 10 {
		return "Stock terbatas"
	}

	return ""
}

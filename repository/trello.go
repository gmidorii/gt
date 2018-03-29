package repository

type TrelloRepository struct {
	Token string
	Key   string
	Board string
}

func NewTrello(token, key, board string) *TrelloRepository {
	return &TrelloRepository{
		Token: token,
		Key:   key,
		Board: board,
	}
}

func (tr *TrelloRepository) FetchCards() {
}

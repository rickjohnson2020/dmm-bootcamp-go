package object

type (
	StatusID = int64

	Status struct {
		ID        StatusID  `json:"id"`
		AccountID AccountID `json:"-" db:"account_id"`
		Content   string    `json:"content"`
		CreateAt  DateTime  `json:"create_at" db:"create_at"`
		Account   *Account  `json:"account"`
	}
)

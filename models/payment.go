package models

type Payment struct {
	ID        int    `json:"id"`
	IdUser    int    `json:"id_user"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type ListPayment struct {
	Payments []Payment `json:"payments"`
}

package db

import (
	"github.com/aulyarahman/bucketeer/models"
)

func (db Database) GetAllPayment() (*models.ListPayment, error) {
	payment := &models.ListPayment{}
	rows, err := db.Conn.Query("SELECT * FROM payment ORDER BY ID DESC")
	//rows, err := db.Conn.Query("SELECT id_user,amount,u.id_user FROM payment FULL OUTER JOIN users AS u ON id_user = u.id_user")

	if err != nil {
		return payment, err
	}

	for rows.Next() {
		var payments models.Payment
		err := rows.Scan(&payments.ID, &payments.Amount)

		if err != nil {
			return payment, err
		}
		payment.Payments = append(payment.Payments, payments)
	}

	return payment, nil
}

func (db Database) CreatePayment(payment *models.Payment) error {
	var id int
	var createdAt string
	query := `INSERT INTO payment (id_user, amount) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, payment.IdUser, payment.Amount).Scan(&id, &createdAt)

	if err != nil {
		return err
	}

	payment.ID = id
	payment.CreatedAt = createdAt
	return nil

}

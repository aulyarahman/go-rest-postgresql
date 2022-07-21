package db

import (
	"bytes"
	"encoding/json"
	"github.com/aulyarahman/bucketeer/models"
	"log"
	"net/http"
	"os"
)

const (
	URL_SEND  = "https://sendtalk-api.taptalk.io/api/v1/message/send_whatsapp"
	URL_CHECK = "https://sendtalk-api.taptalk.io/api/v1/message/get_status"
)

func (db Database) Login(phone *models.Auth) error {
	query := `SELECT * FROM users WHERE phone_number = $1`

	row := db.Conn.QueryRow(query, phone.PhoneNumber)
	err := row.Scan(phone)

	if err != nil {
		return ErrNoMatch
	}

	postBody, _ := json.Marshal(map[string]string{
		"phone":       phone.PhoneNumber,
		"messageType": "otp",
	})
	resBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(URL_SEND, "application/json", resBody)

	resp.Header = http.Header{
		"Conten-Type": {"application/json"},
		"API-key":     {os.Getenv("TAPTALK_KEY")},
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (db Database) VerifyOTP(otp *models.OtpVerify) (models.OtpVerify, error) {
	otpCheck := models.OtpVerify{}

	query := `SELECT * FROM otp WHERE id = $1`
	row := db.Conn.QueryRow(query, otp.Id)
	if err := row.Scan(&otpCheck.Id, &otpCheck.Otp, &otpCheck.CreatedAt); err != nil {
		return otpCheck, ErrNoMatch
	}

	postBody, _ := json.Marshal(map[string]string{
		"id": otp.Id,
	})
	resBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(URL_CHECK, "application/json", resBody)
	resp.Header = http.Header{
		"Conten-Type": {"application/json"},
		"API-key":     {os.Getenv("TAPTALK_KEY")},
	}

	if err != nil {
		panic("Something Wrong.")
	}
	defer resp.Body.Close()

	tapResp := models.TapTalkStatus{}
	if err := json.NewDecoder(resp.Body).Decode(&tapResp); err != nil {
		log.Fatal("Error Occurred")
	}

	if tapResp.Status != 200 {
		log.Fatalf("Something Wrong, %s", tapResp.Error.Message)
	}

	if otpCheck.Otp != otp.Otp {
		log.Fatalf("Otp Not Match: %v", otp)
		return otpCheck, err
	}

	return otpCheck, nil
}

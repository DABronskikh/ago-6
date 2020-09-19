package app

import (
	"context"
	"errors"
	fineV1Pb "github.com/DABronskikh/ago-6/pkg/fine/v1"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrDB = errors.New("error db")
	ErrAutopayNotFound    = errors.New("autopay not found")
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

//получение списка
func (s *Service) GetAutopayUser(ctx context.Context, userID int64) (*fineV1Pb.AutopayResponse, error) {
	var autopayDB []*fineV1Pb.Autopay
	rows, err := s.pool.Query(ctx, `
		SELECT id, name, phone
		FROM autopay_templates
		LIMIT 50
	`)
	defer rows.Close()

	for rows.Next() {
		autopayEl := &fineV1Pb.Autopay{}
		err = rows.Scan(&autopayEl.Id, &autopayEl.Name, &autopayEl.Phone)
		if err != nil {
			return nil, ErrDB
		}
		autopayDB = append(autopayDB, autopayEl)
	}

	if err != nil {
		return nil, ErrDB
	}

	return &fineV1Pb.AutopayResponse{
		Items: autopayDB,
	}, nil
}

//получение по id
func (s *Service) GetAutopayById(ctx context.Context, id int64) (*fineV1Pb.Autopay, error) {
	var autopayDB fineV1Pb.Autopay
	err := s.pool.QueryRow(ctx, `
		SELECT id, name, phone
		FROM autopay_templates
		WHERE id = $1
	`, id).Scan(&autopayDB.Id, &autopayDB.Name, &autopayDB.Phone)

	if err == pgx.ErrNoRows {
		return nil, ErrAutopayNotFound
	}
	if err != nil {
		return nil, err
	}

	return &autopayDB, nil
}

//создание
//редактирование по id
func (s *Service) SaveAutopay(ctx context.Context, autopay *fineV1Pb.Autopay) (*fineV1Pb.Autopay, error) {
	var id int64
	// создадим запись
	if autopay.Id == 0 {
		err := s.pool.QueryRow(ctx, `
		INSERT INTO autopay_templates (name, phone) VALUES($1, $2) RETURNING id
		`, autopay.Name, autopay.Phone).Scan(&id)

		if err == pgx.ErrNoRows {
			return nil, ErrDB
		}
		if err != nil {
			return nil, err
		}

		autopay.Id = id
		return autopay, nil
	}

	// обновим по id
	err := s.pool.QueryRow(ctx, `
		UPDATE autopay_templates SET name = $1, phone = $2 WHERE id = $3 RETURNING id
		`, autopay.Name, autopay.Phone, autopay.Id).Scan(&id)

	if err == pgx.ErrNoRows {
		return nil, ErrAutopayNotFound
	}
	if err != nil {
		return nil, err
	}

	return autopay, nil
}

//удаление по id
func (s *Service) DeleteAutopay(ctx context.Context, autopay *fineV1Pb.Autopay) (*fineV1Pb.Autopay, error) {
	var id int64
	err := s.pool.QueryRow(ctx, `
		DELETE FROM autopay_templates  WHERE id = $1  RETURNING id
		`, autopay.Id).Scan(&id)

	if err == pgx.ErrNoRows {
		return nil, ErrAutopayNotFound
	}
	if err != nil {
		return nil, err
	}

	return autopay, nil
}
package db

import (
	"context"
	"fmt"
	"stockfetcher/models"
	"strings"
)

func GetFilteredStocks(ticker, company, sort string, page, limit, month int) ([]models.Stock, error) {
	query := `
		SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time
		FROM stocks
	`
	var conditions []string
	var params []interface{}
	i := 1

	if ticker != "" {
		conditions = append(conditions, fmt.Sprintf("ticker ILIKE $%d", i))
		params = append(params, "%"+ticker+"%")
		i++
	}
	if company != "" {
		conditions = append(conditions, fmt.Sprintf("company ILIKE $%d", i))
		params = append(params, "%"+company+"%")
		i++
	}
	if month > 0 {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(MONTH FROM time) = $%d", i))
		params = append(params, month)
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Ordenamiento predeterminado (recientes primero)
	query += " ORDER BY time DESC"

	offset := (page - 1) * limit
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := DB.Query(context.Background(), query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var s models.Stock
		if err := rows.Scan(
			&s.Ticker,
			&s.Company,
			&s.Brokerage,
			&s.Action,
			&s.RatingFrom,
			&s.RatingTo,
			&s.TargetFrom,
			&s.TargetTo,
			&s.Time,
		); err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}

	return stocks, nil
}

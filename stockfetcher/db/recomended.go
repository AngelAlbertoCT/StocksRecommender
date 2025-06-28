package db

import (
	"context"
	"fmt"
	"stockfetcher/models"
	"strconv"
	"strings"
)

func GetRecommendedStocks(profile string, page, limit int) ([]models.Stock, error) {
	query := `
		SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time
		FROM stocks
		WHERE
		CASE
	`

	switch profile {
	case "conservative":
	query += `
		WHEN
			(target_from ~ '^[\\$]?[0-9.]+$') AND (target_to ~ '^[\\$]?[0-9.]+$')
		THEN
			(((CAST(REPLACE(target_to, '$', '') AS FLOAT) - CAST(REPLACE(target_from, '$', '') AS FLOAT)) / CAST(REPLACE(target_from, '$', '') AS FLOAT)) * 100) > 0
			AND
			(((CAST(REPLACE(target_to, '$', '') AS FLOAT) - CAST(REPLACE(target_from, '$', '') AS FLOAT)) / CAST(REPLACE(target_from, '$', '') AS FLOAT)) * 100) <= 5
	`
	case "moderate":
		query += `
			WHEN
				(target_from ~ '^[\\$]?[0-9.]+$') AND (target_to ~ '^[\\$]?[0-9.]+$')
			THEN
				(((CAST(REPLACE(target_to, '$', '') AS FLOAT) - CAST(REPLACE(target_from, '$', '') AS FLOAT)) / CAST(REPLACE(target_from, '$', '') AS FLOAT)) * 100) > 5
				AND
				(((CAST(REPLACE(target_to, '$', '') AS FLOAT) - CAST(REPLACE(target_from, '$', '') AS FLOAT)) / CAST(REPLACE(target_from, '$', '') AS FLOAT)) * 100) <= 10
		`
	case "aggressive":
		query += `
			WHEN
				(target_from ~ '^[\\$]?[0-9.]+$') AND (target_to ~ '^[\\$]?[0-9.]+$')
			THEN
				(((CAST(REPLACE(target_to, '$', '') AS FLOAT) - CAST(REPLACE(target_from, '$', '') AS FLOAT)) / CAST(REPLACE(target_from, '$', '') AS FLOAT)) * 100) > 10
		`
	default:
		return nil, fmt.Errorf("invalid profile: %s", profile)
	}

	query += `
		ELSE false
		END
		ORDER BY time DESC
		LIMIT 50
	`

	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.Stock
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

		fromStr := strings.ReplaceAll(s.TargetFrom, "$", "")
		toStr := strings.ReplaceAll(s.TargetTo, "$", "")

		from, err1 := strconv.ParseFloat(fromStr, 64)
		to, err2 := strconv.ParseFloat(toStr, 64)

		if err1 == nil && err2 == nil && from > 0 {
			s.GainPercent = ((to - from) / from) * 100
		} else {
			s.GainPercent = 0.0
		}

		all = append(all, s)
	}


	start := (page - 1) * limit
	end := start + limit
	if start > len(all) {
		return []models.Stock{}, nil
	}
	if end > len(all) {
		end = len(all)
	}

	return all[start:end], nil
}

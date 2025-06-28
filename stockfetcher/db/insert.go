package db

import (
	"context"
	"fmt"
	"stockfetcher/models"
)

func InsertStocks(stocks []models.Stock) error {
	ctx := context.Background()

	for _, stock := range stocks {
		_, err := DB.Exec(ctx, `
			INSERT INTO stocks (
				ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT DO NOTHING
		`,
			stock.Ticker,
			stock.Company,
			stock.Brokerage,
			stock.Action,
			stock.RatingFrom,
			stock.RatingTo,
			stock.TargetFrom,
			stock.TargetTo,
			stock.Time,
		)

		if err != nil {
			fmt.Printf("❌ Error inserting stock %s: %v\n", stock.Ticker, err)
			continue
		}
	}

	return nil
}

package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID
	ExtID         string
	CreatedAt     time.Time
	Name          string
	Price         string
	OriginalPrice string
	Rating        int8
	RatingAverage float64
	Url           string
	ImageUrl      string
}

func (db *Postgres) InsertItemToProduct(item Product) error {
	_, err := db.DB.Exec(`
		INSERT INTO products (
			id,
			created_at,
			name, 
			price, 
			original_price, 
			rating, 
			rating_average, 
			url, 
			url_img,
			ext_id
		) VALUES (
			$1,	$2, $3,	$4,	$5, $6, $7, $8,	$9, $10
		) ON CONFLICT DO NOTHING`,
		item.ID,
		item.CreatedAt,
		item.Name,
		item.Price,
		item.OriginalPrice,
		item.Rating,
		item.RatingAverage,
		item.Url,
		item.ImageUrl,
		item.ExtID,
	)
	if err != nil {
		return err
	}
	return nil
}

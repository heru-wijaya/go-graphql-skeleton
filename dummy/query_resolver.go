package dummy

import (
	"context"
	"time"
)

type queryResolver struct {
}

func (r *queryResolver) Products(ctx context.Context, id *string) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var products []*Product
	product := &Product{
		ID:   "1",
		Name: "testing",
	}
	products = append(products, product)

	return products, nil
}

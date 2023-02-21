package config

import (
	"sqlc-rest-api/graph/generated"
	"sqlc-rest-api/graph/resolvers"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/services"
)

const COMPLEXITY_POINT = 999999999

func GraphConfig(service services.Service) generated.Config {
	resolver := resolvers.NewResolver(service)
	config := generated.Config{
		Resolvers: resolver,
	}

	config.Complexity.Product.User = func(childComplexity int, input *requests.BindUriID) int {
		if childComplexity > 4 {
			return COMPLEXITY_POINT
		}

		return childComplexity + 1
	}

	config.Complexity.User.Products = func(childComplexity int, input *requests.GetUserProductsRequest) int {
		if childComplexity > 12 {
			return COMPLEXITY_POINT
		}

		first := 5
		if input != nil {
			if input.First != nil {
				first = *input.First
			}
		}

		return (childComplexity * first) + 1
	}

	config.Complexity.ProductEdge.Node = func(childComplexity int) int {
		if childComplexity > 5 {
			return COMPLEXITY_POINT
		}

		return childComplexity + 1
	}

	return config
}

package ports

import (
	"context"
	"git.snapp.ninja/search-and-discovery/framework/pkg/adapters/searchEngine/oliverElastic"
	"github.com/olivere/elastic"
)

type SearchEngine interface {
	//GetConnection() *elasticsearch.Client
	GetConnection() *elastic.Client
}

type Source interface {
	SingleFieldSearch(query string, field string, selectedSearchMethods map[string]float64) *oliverElastic.ExtendedSearchSource
	MultiFieldSearch(query string, fields []string, selectedSearchMethods map[string]float64) *oliverElastic.ExtendedSearchSource
	SortByField(fieldSortOrder map[string]string) *oliverElastic.ExtendedSearchSource
	Pagination(from, size int) *oliverElastic.ExtendedSearchSource
	PerformSearch(ctx context.Context, result interface{}) (interface{}, error)
	CreateQuery() *oliverElastic.ExtendedSearchSource
}

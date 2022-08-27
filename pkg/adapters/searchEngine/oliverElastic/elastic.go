package oliverElastic

import (
	"context"
	"fmt"
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/olivere/elastic"
)

type ElasticSearch struct {
	url    string
	user   string
	secret string
}

func New(url, user, secret string) ports.SearchEngine {
	return &ElasticSearch{
		url:    url,
		user:   user,
		secret: secret,
	}
}

type ExtendedSearchSource struct {
	*elastic.SearchSource
	searchCondition []elastic.Query
}

func (e *ElasticSearch) GetConnection() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(e.url))
	elastic.SetBasicAuth(e.user, e.secret)
	if err != nil {
		panic(err)
	}
	return client
}

//type SearchSource struct {
//	*elastic.SearchSource
//}

type unmarshalResult []interface{}

const (
	FUZZY_TYPE                          = "FUZZY"
	FUZZY_AND_CONCATENATED_TYPE         = "FUZZY_CONCATENATED"
	SIMPLE_TYPE                         = "SIMPLE"
	CONCATENATED_DOCUMENT_TYPE          = "CONCATENATED_DOC"
	FUZZY_ON_CONCATENATED_DOCUMENT_TYPE = "FUZZY_ON_CONCATENATED_DOCUMENT"
	TOKEN_BASED_TYPE                    = "TOKEN_BASED"
)

func (s *ExtendedSearchSource) SingleFieldSearch(query string, field string, selectedSearchMethods map[string]float64) *ExtendedSearchSource {
	//s.searchCondition := make([]elastic.Query, 0)
	for searchMethod, boost := range selectedSearchMethods {
		switch searchMethod {
		case FUZZY_TYPE:
			fuzzySearchQuery := fuzzySearch(query, field, boost)
			s.searchCondition = append(s.searchCondition, fuzzySearchQuery)
		case FUZZY_AND_CONCATENATED_TYPE:
			fuzzyAndConcatenatedSearchQuery := fuzzyAndConcatenatedSearch(query, field, boost)
			s.searchCondition = append(s.searchCondition, fuzzyAndConcatenatedSearchQuery)
		case TOKEN_BASED_TYPE:
			simpleSearchQuery := tokenBasedSearch(query, field, boost)
			s.searchCondition = append(s.searchCondition, simpleSearchQuery)
		case CONCATENATED_DOCUMENT_TYPE:
			concatenatedDocumentSearchQuery := concatenatedDocumentSearch(query, field, boost)
			s.searchCondition = append(s.searchCondition, concatenatedDocumentSearchQuery)
		case FUZZY_ON_CONCATENATED_DOCUMENT_TYPE:
			fuzzyOnConcatenatedDocumentSearchQuery := fuzzyOnConcatenatedDocumentSearch(query, field, boost)
			s.searchCondition = append(s.searchCondition, fuzzyOnConcatenatedDocumentSearchQuery)
		default:
			fmt.Println("invalid search type")
		}
	}
	return s
}

func fuzzySearch(query, field string, boost float64) *elastic.MatchQuery {
	return elastic.NewMatchQuery(field, query).Fuzziness("1").Boost(boost)
}

func fuzzyAndConcatenatedSearch(query, field string, boost float64) *elastic.MatchQuery {
	return elastic.NewMatchQuery(field, query).Fuzziness("1").Boost(boost).
		MinimumShouldMatch("2<70%")
}

func tokenBasedSearch(query, field string, boost float64) *elastic.MatchQuery {
	return elastic.NewMatchQuery(field, query).Boost(boost)
}

func concatenatedDocumentSearch(query, field string, boost float64) *elastic.MatchQuery {
	return elastic.NewMatchQuery(field, query).Boost(boost).Analyzer("concatenated_shingle_analyzer")
}

func fuzzyOnConcatenatedDocumentSearch(query, field string, boost float64) *elastic.MatchQuery {
	return elastic.NewMatchQuery(field, query).Fuzziness("1").Boost(boost).
		Analyzer("concatenated_shingle_analyzer").MinimumShouldMatch("1")
}

func (s *ExtendedSearchSource) MultiFieldSearch(query string, fields []string, selectedSearchMethods map[string]float64) *ExtendedSearchSource {
	//queryParts := make([]elastic.Query, 0)
	for searchMethod, boost := range selectedSearchMethods {
		switch searchMethod {
		case TOKEN_BASED_TYPE:
			tokenBasedSearchQuery := multiFieldTokenBasedSearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, tokenBasedSearchQuery)
		case FUZZY_TYPE:
			fuzzySearchQuery := multiFieldFuzzySearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, fuzzySearchQuery)
		case FUZZY_AND_CONCATENATED_TYPE:
			fuzzyAndConcatenatedSearchQuery := multiFieldFuzzyAndConcatenatedSearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, fuzzyAndConcatenatedSearchQuery)
		case SIMPLE_TYPE:
			simpleSearchQuery := multiFieldSimpleSearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, simpleSearchQuery)
		case CONCATENATED_DOCUMENT_TYPE:
			concatenatedDocumentSearchQuery := multiFieldConcatenatedDocumentSearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, concatenatedDocumentSearchQuery)
		case FUZZY_ON_CONCATENATED_DOCUMENT_TYPE:
			fuzzyOnConcatenatedDocumentSearchQuery := multiFieldFuzzyOnConcatenatedDocumentSearch(query, fields, boost)
			s.searchCondition = append(s.searchCondition, fuzzyOnConcatenatedDocumentSearchQuery)
		}
	}
	//finalSearchQuery := elastic.NewDisMaxQuery().Query(queryParts...)
	//s.Query(finalSearchQuery)
	return s
}

func multiFieldSimpleSearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Operator("AND").Boost(boost).
		Type("cross_fields")
}

func multiFieldFuzzySearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Boost(boost).Type("best_fields").
		Fuzziness("1")
}

func multiFieldFuzzyAndConcatenatedSearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Boost(boost).
		Type("best_fields").Fuzziness("1").MinimumShouldMatch("2<70%")
}

func multiFieldTokenBasedSearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Boost(boost).Type("cross_fields")
}

func multiFieldConcatenatedDocumentSearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Boost(boost).Type("cross_fields").
		Analyzer("concatenated_shingle_analyzer")
}

func multiFieldFuzzyOnConcatenatedDocumentSearch(query string, fields []string, boost float64) *elastic.MultiMatchQuery {
	return elastic.NewMultiMatchQuery(query, fields...).Boost(boost).
		Type("best_fields").Analyzer("concatenated_shingle_analyzer").Fuzziness("1").
		MinimumShouldMatch("1")
}

//limitation: we cannot support other features of sortByField as its
//struct is not exported
func (s *ExtendedSearchSource) SortByField(fieldSortOrder map[string]string) *ExtendedSearchSource {
	for field, order := range fieldSortOrder {
		if order == "ASC" {
			s.Sort(field, true)
		} else if order == "DESC" {

			s.Sort(field, false)
		}
	}
	return s
}

func (s *ExtendedSearchSource) Pagination(from, size int) *ExtendedSearchSource {
	s.From(from).Size(size)
	return s
}

func (s *ExtendedSearchSource) PerformSearch(ctx context.Context) (interface{}, error) {
	s.addSearchConditionsToQuery()
	client := ctx.Value("client").(elastic.Client)
	index := ctx.Value("index").(string)
	searchResult, err := client.Search().Index(index).SearchSource(s.SearchSource).Do(ctx)
	if err != nil {
		return searchResult, err
	}

	return searchResult, err
}

func (s *ExtendedSearchSource) addSearchConditionsToQuery() {
	if s.searchCondition != nil {
		finalSearchQuery := elastic.NewDisMaxQuery().Query(s.searchCondition...)
		s.Query(finalSearchQuery)
	}
}

func (s *unmarshalResult) ToJson(searchResult elastic.SearchResult) {
	//	for _, hit := range searchResult.Hits.Hits {
	//		var searcheResult result
	//		err := json.Unmarshal(*hit.Source, &searcheResult)
	//		if err != nil {
	//			// Deserialization failed
	//		}
	//
	//		// Work with tweet
	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//	}
}

func CreateQuery() *ExtendedSearchSource {
	e := &ExtendedSearchSource{SearchSource: elastic.NewSearchSource()}
	return e
}

// limitation on sort
// how to parse elastic output to json

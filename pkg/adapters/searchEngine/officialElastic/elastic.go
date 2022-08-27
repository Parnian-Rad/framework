package officialElastic

//
//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"fmt"
//	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
//	"github.com/elastic/go-elasticsearch/v8"
//	"github.com/elastic/go-elasticsearch/v8/esapi"
//	"log"
//)
//
//type ElasticSearch struct {
//	url    string
//	user   string
//	secret string
//}
//
//type SearchSource struct {
//	query           map[string]interface{}
//	sort            []map[string]interface{}
//	from            int
//	size            int
//	searchCondition []map[string]interface{}
//}
//
//const (
//	FUZZY_TYPE                          = "FUZZY"
//	FUZZY_AND_CONCATENATED_TYPE         = "FUZZY_CONCATENATED"
//	SIMPLE_TYPE                         = "SIMPLE"
//	CONCATENATED_DOCUMENT_TYPE          = "CONCATENATED_DOC"
//	FUZZY_ON_CONCATENATED_DOCUMENT_TYPE = "FUZZY_ON_CONCATENATED_DOCUMENT"
//	TOKEN_BASED_TYPE                    = "TOKEN_BASED"
//)
//
//func New(url, user, secret string) ports.SearchEngine {
//	return &ElasticSearch{
//		url:    url,
//		user:   user,
//		secret: secret,
//	}
//}
//
//func (e *ElasticSearch) GetConnection() *elasticsearch.Client {
//	cfg := elasticsearch.Config{
//		Addresses: []string{e.url},
//		Username:  e.user,
//		Password:  e.secret,
//	}
//	client, err := elasticsearch.NewClient(cfg)
//	if err != nil {
//		panic(err)
//	}
//	return client
//}
//
//func CreateQuery(client *elasticsearch.Client) *SearchSource {
//	return &SearchSource{
//		query: map[string]interface{}{
//			"query": map[string]interface{}{},
//		},
//	}
//}
//
//func (s *SearchSource) SortByField(fieldSortOrder map[string]string) *SearchSource {
//	var sortField map[string]interface{}
//	for field, order := range fieldSortOrder {
//		if order == "ASC" {
//			sortField = map[string]interface{}{
//				"sort": field,
//			}
//		} else if order == "DESC" {
//			sortField = map[string]interface{}{
//				"sort": field,
//			}
//		}
//		s.sort = append(s.sort, sortField)
//	}
//	return s
//}
//
//func (s *SearchSource) Pagination(from, size int) *SearchSource {
//	s.size = size
//	s.from = from
//	return s
//}
//
//func (s *SearchSource) SingleFieldSearch(field string, query string, selectedSearchMethods map[string]float64) *SearchSource {
//	for searchMethod, boost := range selectedSearchMethods {
//		switch searchMethod {
//		case FUZZY_TYPE:
//			fuzzySearchQuery := fuzzySearch(query, field, boost)
//			s.searchCondition = append(s.searchCondition, fuzzySearchQuery)
//		case FUZZY_AND_CONCATENATED_TYPE:
//			fuzzyAndConcatenatedSearchQuery := fuzzyAndConcatenatedSearch(query, field, boost)
//			s.searchCondition = append(s.searchCondition, fuzzyAndConcatenatedSearchQuery)
//		case TOKEN_BASED_TYPE:
//			simpleSearchQuery := tokenBasedSearch(query, field, boost)
//			s.searchCondition = append(s.searchCondition, simpleSearchQuery)
//		case CONCATENATED_DOCUMENT_TYPE:
//			concatenatedDocumentSearchQuery := concatenatedDocumentSearch(query, field, boost)
//			s.searchCondition = append(s.searchCondition, concatenatedDocumentSearchQuery)
//		case FUZZY_ON_CONCATENATED_DOCUMENT_TYPE:
//			fuzzyOnConcatenatedDocumentSearchQuery := fuzzyOnConcatenatedDocumentSearch(query, field, boost)
//			s.searchCondition = append(s.searchCondition, fuzzyOnConcatenatedDocumentSearchQuery)
//		}
//	}
//	return s
//}
//
////type MatchQuery struct {
////	field              string
////	query              interface{}
////	operator           string // or / and
////	analyzer           string
////	boost              *float64
////	fuzziness          string
////	maxExpansions      *int
////	minimumShouldMatch string
////}
////
////func NewMatchQuery(field, query string) *MatchQuery {
////	return &MatchQuery{field: field, query: query}
////}
////func (q *MatchQuery) Fuzziness(fuzziness string) *MatchQuery {
////	q.fuzziness = fuzziness
////	return q
////}
////func (q *MatchQuery) Analyzer(analyzer string) *MatchQuery {
////	q.analyzer = analyzer
////	return q
////}
////func (q *MatchQuery) Operator(operator string) *MatchQuery {
////	q.operator = operator
////	return q
////}
////func (q *MatchQuery) MinimumShouldMatch(minimumShouldMatch string) *MatchQuery {
////	q.minimumShouldMatch = minimumShouldMatch
////	return q
////}
////func (q *MatchQuery) Boost(boost float64) *MatchQuery {
////	q.boost = &boost
////	return q
////}
//
//func fuzzySearch(query, field string, boost float64) map[string]interface{} {
//	//return NewMatchQuery(field, query).Fuzziness("1").Boost(boost)
//	//OR
//	fuzzyQuery := map[string]interface{}{
//		"match": map[string]interface{}{
//			field: map[string]interface{}{
//				"query":     query,
//				"fuzziness": "1",
//				"boost":     boost,
//			},
//		},
//	}
//	return fuzzyQuery
//}
//
//func fuzzyAndConcatenatedSearch(query, field string, boost float64) map[string]interface{} {
//	//return NewMatchQuery(field, query).Fuzziness("1").Boost(boost).MinimumShouldMatch("2<70%")
//	//OR
//	fuzzyAndConcatenatedQuery := map[string]interface{}{
//		"match": map[string]interface{}{
//			field: map[string]interface{}{
//				"query":                query,
//				"boost":                boost,
//				"minimum_should_match": "2<70%",
//			},
//		},
//	}
//	return fuzzyAndConcatenatedQuery
//}
//
//func tokenBasedSearch(query, field string, boost float64) map[string]interface{} {
//	//return NewMatchQuery(field, query).Boost(boost)
//	tokenBasedQuery := map[string]interface{}{
//		"match": map[string]interface{}{
//			field: map[string]interface{}{
//				"query": query,
//				"boost": boost,
//			},
//		},
//	}
//	return tokenBasedQuery
//}
//
//func concatenatedDocumentSearch(query, field string, boost float64) map[string]interface{} {
//	concatenatedDocumentQuery := map[string]interface{}{
//		"match": map[string]interface{}{
//			field: map[string]interface{}{
//				"boost":    boost,
//				"analyzer": "concatenated_shingle_analyzer",
//			},
//		},
//	}
//	return concatenatedDocumentQuery
//}
//
//func fuzzyOnConcatenatedDocumentSearch(query, field string, boost float64) map[string]interface{} {
//	fuzzyOnConcatenatedDocumentQuery := map[string]interface{}{
//		"match": map[string]interface{}{
//			field: map[string]interface{}{
//				"fuzziness":            "1",
//				"boost":                boost,
//				"minimum_should_match": "1",
//				"analyzer":             "concatenated_shingle_analyzer",
//			},
//		},
//	}
//	return fuzzyOnConcatenatedDocumentQuery
//}
//
//func (s *SearchSource) MultiFieldSearch(query string, fields []string, selectedSearchMethods map[string]float64) *SearchSource {
//	for searchMethod, boost := range selectedSearchMethods {
//		switch searchMethod {
//		case TOKEN_BASED_TYPE:
//			tokenBasedSearchQuery := multiFieldTokenBasedSearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, tokenBasedSearchQuery)
//		case FUZZY_TYPE:
//			fuzzySearchQuery := multiFieldFuzzySearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, fuzzySearchQuery)
//		case FUZZY_AND_CONCATENATED_TYPE:
//			fuzzyAndConcatenatedSearchQuery := multiFieldFuzzyAndConcatenatedSearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, fuzzyAndConcatenatedSearchQuery)
//		case SIMPLE_TYPE:
//			simpleSearchQuery := multiFieldSimpleSearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, simpleSearchQuery)
//		case CONCATENATED_DOCUMENT_TYPE:
//			concatenatedDocumentSearchQuery := multiFieldConcatenatedDocumentSearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, concatenatedDocumentSearchQuery)
//		case FUZZY_ON_CONCATENATED_DOCUMENT_TYPE:
//			fuzzyOnConcatenatedDocumentSearchQuery := multiFieldFuzzyOnConcatenatedDocumentSearch(query, fields, boost)
//			s.searchCondition = append(s.searchCondition, fuzzyOnConcatenatedDocumentSearchQuery)
//		default:
//			fmt.Println("invalid search type")
//		}
//	}
//	return s
//}
//
//func multiFieldSimpleSearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldSimpleQuery := map[string]interface{}{
//		"multi_match": map[string]interface{}{
//			"fields":   fields,
//			"query":    query,
//			"operator": "AND",
//			"boost":    boost,
//			"type":     "cross_fields",
//		},
//	}
//	return multiFieldSimpleQuery
//}
//
//func multiFieldFuzzySearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldFuzzyQuery := map[string]interface{}{
//		"fields":    fields,
//		"query":     query,
//		"boost":     boost,
//		"type":      "best_fields",
//		"fuzziness": "1",
//	}
//	return multiFieldFuzzyQuery
//}
//
//func multiFieldFuzzyAndConcatenatedSearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldFuzzyAndConcatenatedQuery := map[string]interface{}{
//		"multi_match": map[string]interface{}{
//			"fields":               fields,
//			"query":                query,
//			"boost":                boost,
//			"type":                 "best_fields",
//			"fuzziness":            "1",
//			"minimum_should_match": "2<70%",
//		},
//	}
//	return multiFieldFuzzyAndConcatenatedQuery
//}
//
//func multiFieldTokenBasedSearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldTokenBasedQuery := map[string]interface{}{
//		"multi_match": map[string]interface{}{
//			"fields": fields,
//			"query":  query,
//			"boost":  boost,
//			"type":   "cross_field",
//		},
//	}
//	return multiFieldTokenBasedQuery
//}
//
//func multiFieldConcatenatedDocumentSearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldConcatenatedDocumentQuery := map[string]interface{}{
//		"multi_match": map[string]interface{}{
//			"query":    query,
//			"fields":   fields,
//			"boost":    boost,
//			"type":     "cross_field",
//			"analyzer": "concatenated_shingle_analyzer",
//		},
//	}
//	return multiFieldConcatenatedDocumentQuery
//}
//
//func multiFieldFuzzyOnConcatenatedDocumentSearch(query string, fields []string, boost float64) map[string]interface{} {
//	multiFieldFuzzyOnConcatenatedDocumentQuery := map[string]interface{}{
//		"multi_match": map[string]interface{}{
//			"fields":               fields,
//			"query":                query,
//			"boost":                boost,
//			"type":                 "best_fields",
//			"analyzer":             "concatenated_shingle_analyzer",
//			"fuzziness":            "1",
//			"minimum_should_match": "1",
//		},
//	}
//	return multiFieldFuzzyOnConcatenatedDocumentQuery
//}
//
//func (s *SearchSource) PerformSearch(client *elasticsearch.Client, index string) (response *esapi.Response, err error) {
//	var buf bytes.Buffer
//	s.GenerateFinalSearchQuery()
//
//	if err := json.NewEncoder(&buf).Encode(s.query); err != nil {
//		panic(err)
//		return response, err
//	}
//	res, err := client.Search(
//		client.Search.WithContext(context.Background()),
//		client.Search.WithIndex(index),
//		client.Search.WithBody(&buf),
//		client.Search.WithTrackTotalHits(true),
//		client.Search.WithPretty(),
//	)00
//	defer res.Body.Close()
//
//	if err != nil {
//		panic(err)
//		return res, err
//	}
//
//	if res.IsError() {
//		var e map[string]interface{}
//		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
//			log.Printf("Error parsing the response body: %s", err)
//		} else {
//			// Print the response status and error information.
//			log.Printf("[%s] %s: %s",
//				res.Status(),
//				e["error"].(map[string]interface{})["type"],
//				e["error"].(map[string]interface{})["reason"],
//			)
//		}
//	}
//
//	return res, err
//}
//
//func (s *SearchSource) GenerateFinalSearchQuery() {
//
//}

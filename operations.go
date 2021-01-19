package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/geojson"
	"net/http"

	// "reflect"
	"errors"
	"log"
	"net/url"
	"sort"

	//"reflect"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Workiva/go-datastructures/bitarray"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/geojson"
)

type filterFuncc func(*Item, string) bool
type registerFuncType map[string]filterFuncc

type bitsetFuncc func(string) bitarray.BitArray
type registerBitSetType map[string]bitsetFuncc

type filterType map[string][]string

func (ft filterType) CacheKey() string {
	filterlist := []string{}
	for k, v := range ft {
		filterlist = append(filterlist, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(filterlist)
	return strings.Join(filterlist, "-")
}

type filterType map[string][]string

func (ft filterType) CacheKey() string {
	filterlist := []string{}
	for k, v := range ft {
		filterlist = append(filterlist, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(filterlist)
	return strings.Join(filterlist, "-")
}

type formatRespFunc func(w http.ResponseWriter, r *http.Request, items Items)
type registerFormatMap map[string]formatRespFunc

type Query struct {
	Filters   filterType
	Excludes  filterType
	Anys      filterType
	BitArrays filterType

	GroupBy string
	Reduce  string

	GroupBy string
	Reduce  string

	Limit         int
	LimitGiven    bool
	Page          int
	PageGiven     bool
	PageSize      int
	PageSizeGiven bool

	SortBy      []string
	SortByGiven bool

	IndexQuery string
	IndexGiven bool

	Geometry      geom.Geometry
	GeometryGiven bool

	ReturnFormat string
}

func (q Query) EarlyExit() bool {
	return q.LimitGiven && !q.PageGiven && !q.SortByGiven
}

// return cachable key for query
func (q Query) CacheKey() (string, error) {

	if SETTINGS.Get("groupbycache") != "yes" {
		return "", errors.New("cache disabled")
	}

	if q.EarlyExit() {
		return "", errors.New("not cached")
	}
	if q.GeometryGiven {
		return "", errors.New("geo not cached")
	}

	for k := range RegisterBitArray {
		_, filterFound := q.Filters[k]
		if filterFound {
			return "", errors.New("bitarrays filters do not need to be cached")
		}
	}

	if q.EarlyExit() {
		return "", errors.New("not cached")
	}

	keys := []string{
		q.Filters.CacheKey(),
		q.Excludes.CacheKey(),
		q.Anys.CacheKey(),
		q.GroupBy,
		q.Reduce,
		q.ReturnFormat,
	}

	return strings.Join(keys, "-"), nil

}

func decodeUrl(s string) string {
	newS, err := url.QueryUnescape(s)
	if err != nil {
		fmt.Println("oh no error", err)
		return s
	}

	keys := []string{
		q.Filters.CacheKey(),
		q.Excludes.CacheKey(),
		q.Anys.CacheKey(),
		q.GroupBy,
		q.Reduce,
		q.ReturnFormat,
	}

	return strings.Join(keys, "-"), nil
}

// parseURLParameters checks parameters and builds a query to be run.
func parseURLParameters(r *http.Request) (Query, error) {
	filterMap := make(filterType)
	excludeMap := make(filterType)
	anyMap := make(filterType)
	groupBy := ""
	reduce := ""

	groupBy := ""
	reduce := ""

	//TODO change query to be based on input.

	// parse params and body posts // (geo)json data
	r.ParseForm()

	if SETTINGS.Get("debug") == "yes" {
		for key, value := range r.Form {
			fmt.Printf("F %s = %s\n", key, value)
		}
	}

	// we can post gejson data

	urlItems := r.URL.Query()
	// parse post geojson data
	r.ParseForm()

	if SETTINGS.Get("debug") == "yes" {

		for key, value := range r.Form {
			fmt.Printf("F %s = %s\n", key, value)
		}
		for key, value := range urlItems {

			fmt.Printf("P %s = %s\n", key, value)
		}
	}

	for k := range RegisterFuncMap {
		parameter, parameterFound := r.Form[k]
		if parameterFound && parameter[0] != "" {
			filterMap[k] = parameter
		}
		parameter, parameterFound = r.Form["!"+k]
		if parameterFound && parameter[0] != "" {
			excludeMap[k] = parameter
		}
		parameter, parameterFound = r.Form["any_"+k]
		if parameterFound && parameter[0] != "" {
			anyMap[k] = parameter
		}
	}

	// Check and validate groupby parameter
	parameter, found := r.Form["groupby"]
	if found && parameter[0] != "" {
		_, funcFound1 := RegisterGroupBy[parameter[0]]
		_, funcFound2 := RegisterGroupByCustom[parameter[0]]
		if !funcFound1 && !funcFound2 {
			return Query{}, errors.New("invalid groupby parameter")
		}
		groupBy = parameter[0]
	}

	// Check and validate reduce parameter
	parameter, found = r.Form["reduce"]

	parameter, found := urlItems["groupby"]
	if found && parameter[0] != "" {
		_, funcFound := RegisterGroupBy[parameter[0]]
		if !funcFound {
			return Query{}, errors.New("Invalid groupby parameter")
		}
		groupBy = parameter[0]

	}

	// Check and validate reduce parameter
	parameter, found = urlItems["reduce"]

	if found && parameter[0] != "" {
		_, funcFound := RegisterReduce[parameter[0]]
		if !funcFound {
			return Query{}, errors.New("Invalid reduce parameter")
		}
		reduce = parameter[0]
	}

	// TODO there must be better way
	page := 1
	pageStr, pageGiven := r.Form["page"]
	if pageGiven {
		page = intMoreDefault(pageStr[0], 1)
	}

	pageSize := 10
	pageSizeStr, pageSizeGiven := r.Form["pagesize"]
	if pageSizeGiven {
		pageSize = intMoreDefault(pageSizeStr[0], 1)
	}

	limit := 0
	limitStr, limitGiven := r.Form["limit"]
	if limitGiven {
		limit = intMoreDefault(limitStr[0], 1)
	}

	format := "json"
	formatStr, formatGiven := r.Form["format"]

	if formatGiven {
		if formatStr[0] == "csv" {
			format = "csv"
		}
	}

	sortingL, sortingGiven := r.Form["sortby"]

	index := ""
	indexL, indexGiven := r.Form["search"]
	indexUsed := indexGiven && indexL[0] != ""

	if indexUsed {
		index = indexL[0]
	}

	// check for geojson geometry stuff.
	geometryS, geometryGiven := r.Form["geojson"]
	var geoinput geojson.Geometry
	if geometryGiven && geometryS[0] != "" {
		err := json.Unmarshal([]byte(geometryS[0]), &geoinput)
		if err != nil {
			fmt.Println("parsing geojson error")
			fmt.Println(err)
			geometryGiven = false
			return Query{}, errors.New("failed to parse geojson")
		}
	}

	return Query{
		Filters:  filterMap,
		Excludes: excludeMap,
		Anys:     anyMap,
		GroupBy:  groupBy,
		Reduce:   reduce,

		GroupBy: groupBy,
		Reduce:  reduce,

		Limit:      limit,
		LimitGiven: limitGiven,

		Page:          page,
		PageGiven:     pageGiven,
		PageSize:      pageSize,
		PageSizeGiven: pageSizeGiven,

		SortBy:      sortingL,
		SortByGiven: sortingGiven,

		IndexQuery: index,
		IndexGiven: indexUsed,

		Geometry: geoinput.Geometry,

		GeometryGiven: geometryGiven,

		ReturnFormat: format,
	}, nil
}

func groupByRunner(items Items, groupByParameter string) ItemsGroupedBy {
	grouping := make(ItemsGroupedBy)
	groupingFunc := RegisterGroupBy[groupByParameter]

	customGrouping := RegisterGroupByCustom[groupByParameter]

	if groupingFunc == nil && customGrouping == nil {
		return grouping
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, item := range items {
		if customGrouping == nil {
			GroupingKey := groupingFunc(item)
			grouping[GroupingKey] = append(grouping[GroupingKey], item)
		} else {
			customGrouping(item, grouping)
		}
	}
	return grouping
}

//Runner of filter functions, Item Should pass all
func all(item *Item, filters filterType, registerFuncs registerFuncType) bool {
	for funcName, args := range filters {
		filterFunc := registerFuncs[funcName]
		if filterFunc == nil {
			continue
		}
		for _, arg := range args {
			if !filterFunc(item, arg) {
				return false
			}
		}
	}
	return true
}

//Runner of filter functions, Item Should pass all
func any(item *Item, filters filterType, registerFuncs registerFuncType) bool {
	if len(filters) == 0 {
		return true
	}
	for funcName, args := range filters {
		filterFunc := registerFuncs[funcName]
		if filterFunc == nil {
			continue
		}
		for _, arg := range args {
			if filterFunc(item, arg) {
				return true
			}
		}
	}
	return false
}

//Runner of exlude functions, Item Should pass all
func exclude(item *Item, excludes filterType, registerFuncs registerFuncType) bool {
	for funcName, args := range excludes {
		excludeFunc := registerFuncs[funcName]
		if excludeFunc == nil {
			continue
		}
		for _, arg := range args {
			if excludeFunc(item, arg) {
				return false
			}
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func filteredEarlyExit(items *labeledItems, operations GroupedOperations, query Query) Items {

	registerFuncs := operations.Funcs
	filteredItems := make(Items, 0, len(*items)/4)
	excludes := query.Excludes
	filters := query.Filters
	anys := query.Anys

	limit := query.Limit
	start := (query.Page - 1) * query.PageSize
	end := start + query.PageSize
	stop := end
	if query.LimitGiven {
		stop = limit
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, item := range items {
		if !any(item, anys, registerFuncs) {
			continue
		}
		if !all(item, filters, registerFuncs) {
			continue
		}
		if !exclude(item, excludes, registerFuncs) {
			continue
		}
		filteredItems = append(filteredItems, item)

		if len(filteredItems) == stop {
			break
		}
	}

	return filteredItems
}

func filteredEarlyExitSingle(items *labeledItems, column string, operations GroupedOperations, query Query) []string {
	registerFuncs := operations.Funcs
	filteredItemsSet := make(map[string]bool)
	excludes := query.Excludes
	filters := query.Filters
	anys := query.Anys

	limit := query.Limit
	start := (query.Page - 1) * query.PageSize
	end := start + query.PageSize
	stop := end
	if query.LimitGiven {
		stop = limit
	}

	lock.RLock()
	defer lock.RUnlock()

	for _, item := range items {
		if !any(item, anys, registerFuncs) {
			continue
		}
		if !all(item, filters, registerFuncs) {
			continue
		}
		if !exclude(item, excludes, registerFuncs) {
			continue
		}

		// return single example value for search field
		if f, ok := operations.Getters[column]; ok {
			single := f(item)
			filteredItemsSet[single] = true
		} else {
			fmt.Println(column)
			fmt.Println("missing getter?")
		}

		if len(filteredItemsSet) == stop {
			break
		}
	}
	results := []string{}
	for k := range filteredItemsSet {
		// empty keys crashes frontend.
		// should be fixed in frontend then below can go.
		// NOTE: add a special field so we can filter on 'nil' / empty values.
		if len(k) > 0 {
			results = append(results, k)
		}
	}
	return results
}

// bit Array Filter.
// for columns with not so unique values it makes sense te create bitarrays.
// to do fast bitwise operations.
func bitArrayFilter(
	items *labeledItems,
	operations GroupedOperations,
	query Query) (labeledItems, error) {

	balock.RLock()
	defer balock.RUnlock()

	lock.RLock()
	defer lock.RUnlock()

	combinedBitArrays := make([]bitarray.BitArray, 0)

	for k := range operations.BitArrays {
		parameter, foundkey := query.Filters["match-"+k]

		if len(parameter) == 0 {
			continue
		}
		if !foundkey {
			continue
		}
		ba, err := operations.BitArrays[k](parameter[0])
		if err != nil {
			continue
		}
		combinedBitArrays = append(combinedBitArrays, ba)

	}

	var bitArrayResult bitarray.BitArray

	if len(combinedBitArrays) > 0 {
		bitArrayResult = combinedBitArrays[0]
	} else {
		log.Println("no bitarrays found")
		return nil, errors.New("no bitarray found")
	}

	// combine AND bitarrays
	if len(combinedBitArrays) > 1 {
		for i := range combinedBitArrays[1:] {
			bitArrayResult = bitArrayResult.And(combinedBitArrays[i])
		}
	}

	// TODO OR
	// TODO EXCLUDE

	if bitArrayResult == nil {
		log.Fatal("something went wrong with bitarray..")
	}

	newItems := make(labeledItems, 0)
	labels := bitArrayResult.ToNums()

	for _, l := range labels {
		newItems = append(newItems, (*items)[int(l)])
	}

	return newItems, nil
}

func runQuery(items *labeledItems, query Query, operations GroupedOperations) (Items, int64) {
	start := time.Now()
	var newItems Items

	if query.GeometryGiven {
		cu := CoverDefault(query.Geometry)
		if len(cu) == 0 {
			fmt.Println("covering cell union not created")
		} else {
			geoitems := SearchGeoItems(cu)
			items = &geoitems
			fmt.Println(len(geoitems))
		}
	}

	var nextItems *labeledItems
	filteredItems, err := bitArrayFilter(items, operations, query)

	if err != nil {
		nextItems = items
	} else {
		nextItems = &filteredItems
	}

	if query.IndexGiven && len(STR_INDEX) > 0 {
		items = make(Items, 0)
		indices := INDEX.Lookup([]byte(query.IndexQuery), -1)
		seen := make(map[string]bool)
		for _, idx := range indices {
			key := getStringFromIndex(STR_INDEX, idx)
			if !seen[key] {
				seen[key] = true
				items = append(items, LOOKUP[key]...)
			}
		}
	}

	if query.EarlyExit() {
		newItems = filteredEarlyExit(nextItems, operations, query)
	} else {
		newItems = filtered(nextItems, operations, query)
	}

	if query.GeometryGiven {
		cu := CoverDefault(query.Geometry)
		if len(cu) == 0 {
			fmt.Println("covering cell union not created")
		} else {
			geoitems := SearchOverlapItems(items, cu)
			items = &geoitems
			fmt.Println(len(geoitems))
		}
	}

	if query.EarlyExit() {
		newItems = filteredEarlyExit(items, operations, query)
	} else {
		newItems = filtered(items, operations, query)
	}

	diff := time.Since(start)
	return newItems, int64(diff) / int64(1000000)
}

func runTypeAheadQuery(
	items *labeledItems, column string, query Query,
	operations GroupedOperations) ([]string, int64) {
	start := time.Now()
	results := filteredEarlyExitSingle(items, column, operations, query)
	diff := time.Since(start)
	return results, int64(diff) / int64(1000000)
}

func filtered(items *labeledItems, operations GroupedOperations, query Query) Items {
	registerFuncs := operations.Funcs
	filteredItems := make(Items, 0)
	excludes := query.Excludes
	filters := query.Filters
	anys := query.Anys

	lock.RLock()
	defer lock.RUnlock()

	for _, item := range *items {
		if !any(item, anys, registerFuncs) {
			continue
		}
		if !all(item, filters, registerFuncs) {
			continue
		}
		if !exclude(item, excludes, registerFuncs) {
			continue
		}
		filteredItems = append(filteredItems, item)
	}
	return filteredItems
}

func mapIndex(items Items, indexes []int) Items {
	o := Items{}
	for _, index := range indexes {
		o = append(o, items[index])
	}
	return o
}

type HeaderData map[string]string

func getHeaderData(items Items, query Query, queryDuration int64) HeaderData {
	headerData := make(HeaderData)

	if query.LimitGiven {
		headerData["Limit"] = strconv.Itoa(query.Limit)
	}

	if query.PageGiven {
		headerData["Page"] = strconv.Itoa(query.Page)
		headerData["Page-Size"] = strconv.Itoa(query.PageSize)
		headerData["Total-Pages"] = strconv.Itoa((len(items) / query.PageSize) + 1)
	}

	headerData["Cache-Control"] = "public, max-age=300"
	headerData["Total-Items"] = strconv.Itoa(len(items))
	headerData["Query-Duration"] = strconv.FormatInt(queryDuration, 10) + "ms"
	bytesQuery, _ := json.Marshal(query)
	headerData["query"] = string(bytesQuery)

	return headerData
}

//getHeaderDataSlice extract from header information with data slice we want
func getHeaderDataSlice(items []string, query Query, queryDuration int64) HeaderData {
	headerData := make(HeaderData)

	if query.LimitGiven {
		headerData["Limit"] = strconv.Itoa(query.Limit)
	}

	if query.PageGiven {
		headerData["Page"] = strconv.Itoa(query.Page)
		headerData["Page-Size"] = strconv.Itoa(query.PageSize)
		headerData["Total-Pages"] = strconv.Itoa((len(items) / query.PageSize) + 1)
	}

	headerData["Total-Items"] = strconv.Itoa(len(items))
	headerData["Query-Duration"] = strconv.FormatInt(queryDuration, 10) + "ms"
	bytesQuery, _ := json.Marshal(query)
	headerData["query"] = string(bytesQuery)

	return headerData
}

func sortLimit(items Items, query Query) Items {
	count := len(items)
	if count == 0 {
		return items
	}

	if query.SortByGiven {
		items, _ = sortBy(items, query.SortBy)
	}

	if !query.LimitGiven && !query.PageGiven {
		return items
	}

	//TODO there should be nicer way
	start := (query.Page - 1) * query.PageSize
	end := start + query.PageSize

	items = items[min(start, count):min(end, count)]
	if !query.LimitGiven {
		return items
	}

	// Note the slice built on array, slicing a slice larger then the the slice adds array items
	if len(items) < query.Limit {
		return items
	}
	return items[:query.Limit]
}

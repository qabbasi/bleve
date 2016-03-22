package bleve

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/blevesearch/bleve/document"
	"github.com/blevesearch/bleve/index"
	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/search"
)

func TestIndexAliasSingle(t *testing.T) {
	expectedError := fmt.Errorf("expected")
	ei1 := &stubIndex{
		err: expectedError,
	}

	alias := NewIndexAlias(ei1)

	err := alias.Index("a", "a")
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	err = alias.Delete("a")
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	batch := alias.NewBatch()
	err = alias.Batch(batch)
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	_, err = alias.Document("a")
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	_, err = alias.Fields()
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	res := alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping := alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat := alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	sr := NewSearchRequest(NewTermQuery("test"))
	_, err = alias.Search(sr)
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	count, err := alias.DocCount()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}

	// now change the def using add/remove
	expectedError2 := fmt.Errorf("expected2")
	ei2 := &stubIndex{
		err: expectedError2,
	}

	alias.Add(ei2)
	alias.Remove(ei1)

	err = alias.Index("a", "a")
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	err = alias.Delete("a")
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	err = alias.Batch(batch)
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	_, err = alias.Document("a")
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	_, err = alias.Fields()
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	res = alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping = alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat = alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	_, err = alias.Search(sr)
	if err != expectedError2 {
		t.Errorf("expected %v, got %v", expectedError2, err)
	}

	count, err = alias.DocCount()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}

	// now change the def using swap
	expectedError3 := fmt.Errorf("expected3")
	ei3 := &stubIndex{
		err: expectedError3,
	}

	alias.Swap([]Index{ei3}, []Index{ei2})

	err = alias.Index("a", "a")
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	err = alias.Delete("a")
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	err = alias.Batch(batch)
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	_, err = alias.Document("a")
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	_, err = alias.Fields()
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	res = alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping = alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat = alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	_, err = alias.Search(sr)
	if err != expectedError3 {
		t.Errorf("expected %v, got %v", expectedError3, err)
	}

	count, err = alias.DocCount()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}
}

func TestIndexAliasClosed(t *testing.T) {
	alias := NewIndexAlias()
	err := alias.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = alias.Index("a", "a")
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	err = alias.Delete("a")
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	batch := alias.NewBatch()
	err = alias.Batch(batch)
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	_, err = alias.Document("a")
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	_, err = alias.Fields()
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	res := alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping := alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat := alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	sr := NewSearchRequest(NewTermQuery("test"))
	_, err = alias.Search(sr)
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}

	_, err = alias.DocCount()
	if err != ErrorIndexClosed {
		t.Errorf("expected %v, got %v", ErrorIndexClosed, err)
	}
}

func TestIndexAliasEmpty(t *testing.T) {
	alias := NewIndexAlias()

	err := alias.Index("a", "a")
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	err = alias.Delete("a")
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	batch := alias.NewBatch()
	err = alias.Batch(batch)
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	_, err = alias.Document("a")
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	_, err = alias.Fields()
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	res := alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping := alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat := alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	sr := NewSearchRequest(NewTermQuery("test"))
	_, err = alias.Search(sr)
	if err != ErrorAliasEmpty {
		t.Errorf("expected %v, got %v", ErrorAliasEmpty, err)
	}

	count, err := alias.DocCount()
	if count != 0 {
		t.Errorf("expected %d, got %d", 0, count)
	}
}

func TestIndexAliasMulti(t *testing.T) {
	ei1Count := uint64(7)
	ei1 := &stubIndex{
		err:            nil,
		docCountResult: &ei1Count,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: search.DocumentMatchCollection{
				&search.DocumentMatch{
					ID:    "a",
					Score: 1.0,
				},
			},
			MaxScore: 1.0,
		}}
	ei2Count := uint64(8)
	ei2 := &stubIndex{
		err:            nil,
		docCountResult: &ei2Count,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: search.DocumentMatchCollection{
				&search.DocumentMatch{
					ID:    "b",
					Score: 2.0,
				},
			},
			MaxScore: 2.0,
		}}

	alias := NewIndexAlias(ei1, ei2)

	err := alias.Index("a", "a")
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	err = alias.Delete("a")
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	batch := alias.NewBatch()
	err = alias.Batch(batch)
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	_, err = alias.Document("a")
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	_, err = alias.Fields()
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	_, err = alias.GetInternal([]byte("a"))
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	err = alias.SetInternal([]byte("a"), []byte("a"))
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	err = alias.DeleteInternal([]byte("a"))
	if err != ErrorAliasMulti {
		t.Errorf("expected %v, got %v", ErrorAliasMulti, err)
	}

	res := alias.DumpAll()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpDoc("a")
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	res = alias.DumpFields()
	if res != nil {
		t.Errorf("expected nil, got %v", res)
	}

	mapping := alias.Mapping()
	if mapping != nil {
		t.Errorf("expected nil, got %v", res)
	}

	indexStat := alias.Stats()
	if indexStat != nil {
		t.Errorf("expected nil, got %v", res)
	}

	// now a few things that should work
	sr := NewSearchRequest(NewTermQuery("test"))
	expected := &SearchResult{
		Status: &SearchStatus{
			Total:      2,
			Successful: 2,
			Errors:     make(map[string]error),
		},
		Request: sr,
		Total:   2,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				ID:    "b",
				Score: 2.0,
			},
			&search.DocumentMatch{
				ID:    "a",
				Score: 1.0,
			},
		},
		MaxScore: 2.0,
	}
	results, err := alias.Search(sr)
	if err != nil {
		t.Error(err)
	}
	// cheat and ensure that Took field matches since it invovles time
	expected.Took = results.Took
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %#v, got %#v", expected, results)
	}

	count, err := alias.DocCount()
	if count != (*ei1.docCountResult + *ei2.docCountResult) {
		t.Errorf("expected %d, got %d", (*ei1.docCountResult + *ei2.docCountResult), count)
	}
}

// TestMultiSearchNoError
func TestMultiSearchNoError(t *testing.T) {
	ei1 := &stubIndex{err: nil, searchResult: &SearchResult{
		Status: &SearchStatus{
			Total:      1,
			Successful: 1,
			Errors:     make(map[string]error),
		},
		Total: 1,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				Index: "1",
				ID:    "a",
				Score: 1.0,
			},
		},
		MaxScore: 1.0,
	}}
	ei2 := &stubIndex{err: nil, searchResult: &SearchResult{
		Status: &SearchStatus{
			Total:      1,
			Successful: 1,
			Errors:     make(map[string]error),
		},
		Total: 1,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				Index: "2",
				ID:    "b",
				Score: 2.0,
			},
		},
		MaxScore: 2.0,
	}}

	sr := NewSearchRequest(NewTermQuery("test"))
	expected := &SearchResult{
		Status: &SearchStatus{
			Total:      2,
			Successful: 2,
			Errors:     make(map[string]error),
		},
		Request: sr,
		Total:   2,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				Index: "2",
				ID:    "b",
				Score: 2.0,
			},
			&search.DocumentMatch{
				Index: "1",
				ID:    "a",
				Score: 1.0,
			},
		},
		MaxScore: 2.0,
	}

	results, err := MultiSearch(context.Background(), sr, ei1, ei2)
	if err != nil {
		t.Error(err)
	}
	// cheat and ensure that Took field matches since it invovles time
	expected.Took = results.Took
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %#v, got %#v", expected, results)
	}
}

// TestMultiSearchSomeError
func TestMultiSearchSomeError(t *testing.T) {
	ei1 := &stubIndex{name: "ei1", err: nil, searchResult: &SearchResult{
		Status: &SearchStatus{
			Total:      1,
			Successful: 1,
			Errors:     make(map[string]error),
		},
		Total: 1,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				ID:    "a",
				Score: 1.0,
			},
		},
		Took:     1 * time.Second,
		MaxScore: 1.0,
	}}
	ei2 := &stubIndex{name: "ei2", err: fmt.Errorf("deliberate error")}
	sr := NewSearchRequest(NewTermQuery("test"))
	res, err := MultiSearch(context.Background(), sr, ei1, ei2)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Status.Total != 2 {
		t.Errorf("expected 2 indexes to be queried, got %d", res.Status.Total)
	}
	if res.Status.Failed != 1 {
		t.Errorf("expected 1 index to fail, got %d", res.Status.Failed)
	}
	if res.Status.Successful != 1 {
		t.Errorf("expected 1 index to be successful, got %d", res.Status.Successful)
	}
	if len(res.Status.Errors) != 1 {
		t.Fatalf("expected 1 status error message, got %d", len(res.Status.Errors))
	}
	if res.Status.Errors["ei2"].Error() != "deliberate error" {
		t.Errorf("expected ei2 index error message 'deliberate error', got '%s'", res.Status.Errors["ei2"])
	}
}

// TestMultiSearchAllError
// reproduces https://github.com/blevesearch/bleve/issues/126
func TestMultiSearchAllError(t *testing.T) {
	ei1 := &stubIndex{name: "ei1", err: fmt.Errorf("deliberate error")}
	ei2 := &stubIndex{name: "ei2", err: fmt.Errorf("deliberate error")}
	sr := NewSearchRequest(NewTermQuery("test"))
	res, err := MultiSearch(context.Background(), sr, ei1, ei2)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Status.Total != 2 {
		t.Errorf("expected 2 indexes to be queried, got %d", res.Status.Total)
	}
	if res.Status.Failed != 2 {
		t.Errorf("expected 2 indexes to fail, got %d", res.Status.Failed)
	}
	if res.Status.Successful != 0 {
		t.Errorf("expected 0 indexes to be successful, got %d", res.Status.Successful)
	}
	if len(res.Status.Errors) != 2 {
		t.Fatalf("expected 2 status error messages, got %d", len(res.Status.Errors))
	}
	if res.Status.Errors["ei1"].Error() != "deliberate error" {
		t.Errorf("expected ei1 index error message 'deliberate error', got '%s'", res.Status.Errors["ei1"])
	}
	if res.Status.Errors["ei2"].Error() != "deliberate error" {
		t.Errorf("expected ei2 index error message 'deliberate error', got '%s'", res.Status.Errors["ei2"])
	}
}

func TestMultiSearchSecondPage(t *testing.T) {
	checkRequest := func(sr *SearchRequest) error {
		if sr.From != 0 {
			return fmt.Errorf("child request from should be 0")
		}
		if sr.Size != 20 {
			return fmt.Errorf("child request size should be 20")
		}
		return nil
	}

	ei1 := &stubIndex{
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
		},
		checkRequest: checkRequest,
	}
	ei2 := &stubIndex{
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
		},
		checkRequest: checkRequest,
	}
	sr := NewSearchRequestOptions(NewTermQuery("test"), 10, 10, false)
	_, err := MultiSearch(context.Background(), sr, ei1, ei2)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

}

// TestMultiSearchTimeout tests simple timeout cases
// 1. all searches finish successfully before timeout
// 2. no searchers finish before the timeout
// 3. no searches finish before cancellation
func TestMultiSearchTimeout(t *testing.T) {
	ei1 := &stubIndex{
		name: "ei1",
		checkRequest: func(req *SearchRequest) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		err: nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "1",
					ID:    "a",
					Score: 1.0,
				},
			},
			MaxScore: 1.0,
		}}
	ei2 := &stubIndex{
		name: "ei2",
		checkRequest: func(req *SearchRequest) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		err: nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "2",
					ID:    "b",
					Score: 2.0,
				},
			},
			MaxScore: 2.0,
		}}

	// first run with absurdly long time out, should succeed
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	query := NewTermQuery("test")
	sr := NewSearchRequest(query)
	res, err := MultiSearch(ctx, sr, ei1, ei2)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Status.Total != 2 {
		t.Errorf("expected 2 total, got %d", res.Status.Failed)
	}
	if res.Status.Successful != 2 {
		t.Errorf("expected 0 success, got %d", res.Status.Successful)
	}
	if res.Status.Failed != 0 {
		t.Errorf("expected 2 failed, got %d", res.Status.Failed)
	}
	if len(res.Status.Errors) != 0 {
		t.Errorf("expected 0 errors, got %v", res.Status.Errors)
	}

	// now run a search again with an absurdly low timeout (should timeout)
	ctx, _ = context.WithTimeout(context.Background(), 1*time.Microsecond)
	res, err = MultiSearch(ctx, sr, ei1, ei2)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Status.Total != 2 {
		t.Errorf("expected 2 failed, got %d", res.Status.Failed)
	}
	if res.Status.Successful != 0 {
		t.Errorf("expected 0 success, got %d", res.Status.Successful)
	}
	if res.Status.Failed != 2 {
		t.Errorf("expected 2 failed, got %d", res.Status.Failed)
	}
	if len(res.Status.Errors) != 2 {
		t.Errorf("expected 2 errors, got %v", res.Status.Errors)
	} else {
		if res.Status.Errors["ei1"].Error() != context.DeadlineExceeded.Error() {
			t.Errorf("expected err for 'ei1' to be '%s' got '%s'", context.DeadlineExceeded.Error(), res.Status.Errors["ei1"])
		}
		if res.Status.Errors["ei2"].Error() != context.DeadlineExceeded.Error() {
			t.Errorf("expected err for 'ei2' to be '%s' got '%s'", context.DeadlineExceeded.Error(), res.Status.Errors["ei2"])
		}
	}

	// now run a search again with a normal timeout, but cancel it first
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cancel()
	res, err = MultiSearch(ctx, sr, ei1, ei2)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Status.Total != 2 {
		t.Errorf("expected 2 failed, got %d", res.Status.Failed)
	}
	if res.Status.Successful != 0 {
		t.Errorf("expected 0 success, got %d", res.Status.Successful)
	}
	if res.Status.Failed != 2 {
		t.Errorf("expected 2 failed, got %d", res.Status.Failed)
	}
	if len(res.Status.Errors) != 2 {
		t.Errorf("expected 2 errors, got %v", res.Status.Errors)
	} else {
		if res.Status.Errors["ei1"].Error() != context.Canceled.Error() {
			t.Errorf("expected err for 'ei1' to be '%s' got '%s'", context.Canceled.Error(), res.Status.Errors["ei1"])
		}
		if res.Status.Errors["ei2"].Error() != context.Canceled.Error() {
			t.Errorf("expected err for 'ei2' to be '%s' got '%s'", context.Canceled.Error(), res.Status.Errors["ei2"])
		}
	}
}

// TestMultiSearchTimeoutPartial tests the case where some indexes exceed
// the timeout, while others complete successfully
func TestMultiSearchTimeoutPartial(t *testing.T) {
	ei1 := &stubIndex{
		name: "ei1",
		err:  nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "1",
					ID:    "a",
					Score: 1.0,
				},
			},
			MaxScore: 1.0,
		}}
	ei2 := &stubIndex{
		name: "ei2",
		err:  nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "2",
					ID:    "b",
					Score: 2.0,
				},
			},
			MaxScore: 2.0,
		}}

	ei3 := &stubIndex{
		name: "ei3",
		checkRequest: func(req *SearchRequest) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		err: nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "3",
					ID:    "c",
					Score: 3.0,
				},
			},
			MaxScore: 3.0,
		}}

	// ei3 is set to take >50ms, so run search with timeout less than
	// this, this should return partial results
	ctx, _ := context.WithTimeout(context.Background(), 25*time.Millisecond)
	query := NewTermQuery("test")
	sr := NewSearchRequest(query)
	expected := &SearchResult{
		Status: &SearchStatus{
			Total:      3,
			Successful: 2,
			Failed:     1,
			Errors: map[string]error{
				"ei3": context.DeadlineExceeded,
			},
		},
		Request: sr,
		Total:   2,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				Index: "2",
				ID:    "b",
				Score: 2.0,
			},
			&search.DocumentMatch{
				Index: "1",
				ID:    "a",
				Score: 1.0,
			},
		},
		MaxScore: 2.0,
	}

	res, err := MultiSearch(ctx, sr, ei1, ei2, ei3)
	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}
	expected.Took = res.Took
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %#v, got %#v", expected, res)
	}
}

func TestIndexAliasMultipleLayer(t *testing.T) {
	ei1 := &stubIndex{
		name: "ei1",
		err:  nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "1",
					ID:    "a",
					Score: 1.0,
				},
			},
			MaxScore: 1.0,
		}}
	ei2 := &stubIndex{
		name: "ei2",
		checkRequest: func(req *SearchRequest) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		err: nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "2",
					ID:    "b",
					Score: 2.0,
				},
			},
			MaxScore: 2.0,
		}}

	ei3 := &stubIndex{
		name: "ei3",
		checkRequest: func(req *SearchRequest) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		err: nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "3",
					ID:    "c",
					Score: 3.0,
				},
			},
			MaxScore: 3.0,
		}}

	ei4 := &stubIndex{
		name: "ei4",
		err:  nil,
		searchResult: &SearchResult{
			Status: &SearchStatus{
				Total:      1,
				Successful: 1,
				Errors:     make(map[string]error),
			},
			Total: 1,
			Hits: []*search.DocumentMatch{
				&search.DocumentMatch{
					Index: "4",
					ID:    "d",
					Score: 4.0,
				},
			},
			MaxScore: 4.0,
		}}

	alias1 := NewIndexAlias(ei1, ei2)
	alias2 := NewIndexAlias(ei3, ei4)
	aliasTop := NewIndexAlias(alias1, alias2)

	// ei2 and ei3 have 50ms delay
	// search across aliasTop should still get results from ei1 and ei4
	// total should still be 4

	ctx, _ := context.WithTimeout(context.Background(), 25*time.Millisecond)
	query := NewTermQuery("test")
	sr := NewSearchRequest(query)
	expected := &SearchResult{
		Status: &SearchStatus{
			Total:      4,
			Successful: 2,
			Failed:     2,
			Errors: map[string]error{
				"ei2": context.DeadlineExceeded,
				"ei3": context.DeadlineExceeded,
			},
		},
		Request: sr,
		Total:   2,
		Hits: search.DocumentMatchCollection{
			&search.DocumentMatch{
				Index: "4",
				ID:    "d",
				Score: 4.0,
			},
			&search.DocumentMatch{
				Index: "1",
				ID:    "a",
				Score: 1.0,
			},
		},
		MaxScore: 4.0,
	}

	res, err := aliasTop.SearchInContext(ctx, sr)
	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}
	expected.Took = res.Took
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %#v, got %#v", expected, res)
	}
}

// stubIndex is an Index impl for which all operations
// return the configured error value, unless the
// corresponding operation result value has been
// set, in which case that is returned instead
type stubIndex struct {
	name           string
	err            error
	searchResult   *SearchResult
	documentResult *document.Document
	docCountResult *uint64
	checkRequest   func(*SearchRequest) error
}

func (i *stubIndex) Index(id string, data interface{}) error {
	return i.err
}

func (i *stubIndex) Delete(id string) error {
	return i.err
}

func (i *stubIndex) Batch(b *Batch) error {
	return i.err
}

func (i *stubIndex) Document(id string) (*document.Document, error) {
	if i.documentResult != nil {
		return i.documentResult, nil
	}
	return nil, i.err
}

func (i *stubIndex) DocCount() (uint64, error) {
	if i.docCountResult != nil {
		return *i.docCountResult, nil
	}
	return 0, i.err
}

func (i *stubIndex) Search(req *SearchRequest) (*SearchResult, error) {
	return i.SearchInContext(context.Background(), req)
}

func (i *stubIndex) SearchInContext(ctx context.Context, req *SearchRequest) (*SearchResult, error) {
	if i.checkRequest != nil {
		err := i.checkRequest(req)
		if err != nil {
			return nil, err
		}
	}
	if i.searchResult != nil {
		return i.searchResult, nil
	}
	return nil, i.err
}

func (i *stubIndex) Fields() ([]string, error) {
	return nil, i.err
}

func (i *stubIndex) FieldDict(field string) (index.FieldDict, error) {
	return nil, i.err
}

func (i *stubIndex) FieldDictRange(field string, startTerm []byte, endTerm []byte) (index.FieldDict, error) {
	return nil, i.err
}

func (i *stubIndex) FieldDictPrefix(field string, termPrefix []byte) (index.FieldDict, error) {
	return nil, i.err
}

func (i *stubIndex) DumpAll() chan interface{} {
	return nil
}

func (i *stubIndex) DumpDoc(id string) chan interface{} {
	return nil
}

func (i *stubIndex) DumpFields() chan interface{} {
	return nil
}

func (i *stubIndex) Close() error {
	return i.err
}

func (i *stubIndex) Mapping() *IndexMapping {
	return nil
}

func (i *stubIndex) Stats() *IndexStat {
	return nil
}

func (i *stubIndex) StatsMap() map[string]interface{} {
	return nil
}

func (i *stubIndex) GetInternal(key []byte) ([]byte, error) {
	return nil, i.err
}

func (i *stubIndex) SetInternal(key, val []byte) error {
	return i.err
}

func (i *stubIndex) DeleteInternal(key []byte) error {
	return i.err
}

func (i *stubIndex) Advanced() (index.Index, store.KVStore, error) {
	return nil, nil, nil
}

func (i *stubIndex) NewBatch() *Batch {
	return &Batch{}
}

func (i *stubIndex) Name() string {
	return i.name
}

func (i *stubIndex) SetName(name string) {
	i.name = name
}

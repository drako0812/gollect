package gollect

import (
	"math/rand"
	"testing"
	"time"
)

const nVectorItemsCount = 50000000

var nVectorSearchIndex int
var nVectorVec NVector[int]

func BenchmarkNVectorSearchFunctions(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(12345)
	nVectorSearchIndex = rand.Intn(nVectorItemsCount)
	//nVectorSearchIndex = int(nVectorItemsCount * 0.75)
	nVectorVec = NewNVector[int]()
	for i := 0; i < nVectorItemsCount; i++ {
		nVectorVec.PushBack(rand.Int())
		//nVectorVec.PushBack(i)
	}

	b.Run("OrderedSearch", SubBenchmarkNVectorOrderedSearch)
	b.Run("OrderedRefSearch", SubBenchmarkNVectorOrderedRefSearch)
	b.Run("OrderedSearchRef", SubBenchmarkNVectorOrderedSearchRef)
	b.Run("OrderedRefSearchRef", SubBenchmarkNVectorOrderedRefSearchRef)
	b.Run("Search", SubBenchmarkNVectorSearch)
	b.Run("RefSearch", SubBenchmarkNVectorRefSearch)
	b.Run("SearchRef", SubBenchmarkNVectorSearchRef)
	b.Run("RefSearchRef", SubBenchmarkNVectorRefSearchRef)
}

func SubBenchmarkNVectorOrderedSearch(b *testing.B) {
	searched_for := nVectorVec.At(nVectorSearchIndex)
	nVectorVec.OrderedSearch(searched_for)
	// if found {
	// 	b.Logf("Found %v at %v (expected %v)", searched_for, idx, search_index)
	// } else {
	// 	b.Logf("Did not find %v at %v", searched_for, search_index)
	// }
}

func SubBenchmarkNVectorOrderedRefSearch(b *testing.B) {
	searched_for := nVectorVec.AtRef(nVectorSearchIndex)
	nVectorVec.OrderedRefSearch(searched_for)
	// if found {
	// 	b.Logf("Found %v(%v) at %v (expected %v)", *searched_for, searched_for, idx, search_index)
	// } else {
	// 	b.Logf("Did not find %v(%v) at %v", searched_for, *searched_for, search_index)
	// }
}

func SubBenchmarkNVectorOrderedSearchRef(b *testing.B) {
	searched_for := nVectorVec.At(nVectorSearchIndex)
	nVectorVec.OrderedSearchRef(searched_for)
	// if found != nil {
	// 	b.Logf("Found %v(%v)", *found, found)
	// } else {
	// 	b.Logf("Did not find %v", searched_for)
	// }
}

func SubBenchmarkNVectorOrderedRefSearchRef(b *testing.B) {
	searched_for := nVectorVec.AtRef(nVectorSearchIndex)
	nVectorVec.OrderedRefSearchRef(searched_for)
	// if found != nil {
	// 	b.Logf("Found %v(%v)", *found, found)
	// } else {
	// 	b.Logf("Did not find %v(%v)", *searched_for, searched_for)
	// }
}

func SubBenchmarkNVectorSearch(b *testing.B) {
	searched_for := nVectorVec.At(nVectorSearchIndex)
	nVectorVec.Search(searched_for)
	// if found {
	// 	b.Logf("Found %v at %v (expected %v)", searched_for, idx, search_index)
	// } else {
	// 	b.Logf("Did not find %v at %v", searched_for, search_index)
	// }
}

func SubBenchmarkNVectorRefSearch(b *testing.B) {
	searched_for := nVectorVec.AtRef(nVectorSearchIndex)
	nVectorVec.RefSearch(searched_for)
	// if found {
	// 	b.Logf("Found %v(%v) at %v (expected %v)", *searched_for, searched_for, idx, search_index)
	// } else {
	// 	b.Logf("Did not find %v(%v) at %v", searched_for, *searched_for, search_index)
	// }
}

func SubBenchmarkNVectorSearchRef(b *testing.B) {
	searched_for := nVectorVec.At(nVectorSearchIndex)
	nVectorVec.SearchRef(searched_for)
	// if found != nil {
	// 	b.Logf("Found %v(%v)", *found, found)
	// } else {
	// 	b.Logf("Did not find %v", searched_for)
	// }
}

func SubBenchmarkNVectorRefSearchRef(b *testing.B) {
	searched_for := nVectorVec.AtRef(nVectorSearchIndex)
	nVectorVec.RefSearchRef(searched_for)
	// if found != nil {
	// 	b.Logf("Found %v(%v)", *found, found)
	// } else {
	// 	b.Logf("Did not find %v(%v)", *searched_for, searched_for)
	// }
}

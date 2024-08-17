package noasm_test

import (
	"testing"

	"github.com/kshard/vector/internal/noasm"
	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/vtest"
)

func TestEuclidean(t *testing.T) {
	sut := noasm.Euclidean(0)

	vtest.TestEqual(t, vtest.ID, sut)
	vtest.TestDistance(t, vtest.ID, pure.Euclidean(0), sut)
}

func TestCosine(t *testing.T) {
	sut := noasm.Cosine(0)

	vtest.TestEqual(t, vtest.ID, sut)
	vtest.TestDistance(t, vtest.ID, pure.Cosine(0), sut)
}

package flights

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNormal(t *testing.T) {
	var list [][]string
	var expected []string

	list = [][]string{
		[]string{"SFO", "EWR"},
	}
	expected = []string{
		"SFO", "EWR",
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))

	list = [][]string{
		[]string{"SFO", "ATL"},
		[]string{"ATL", "EWR"},
	}
	expected = []string{
		"SFO", "EWR",
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))

	list = [][]string{
		[]string{"IND", "EWR"},
		[]string{"SFO", "ATL"},
		[]string{"GSO", "IND"},
		[]string{"ATL", "GSO"},
	}
	expected = []string{
		"SFO", "EWR",
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))

	list = [][]string{
		[]string{"B", "C"},
		[]string{"D", "E"},
		[]string{"C", "D"},
		[]string{"A", "B"},
		[]string{"E", "F"},
	}
	expected = []string{
		"A", "F",
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
}

func TestBadInput(t *testing.T) {
	var list [][]string
	var expected []string = []string{}

	// bad input
	list = nil
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
	// bad input
	list = [][]string{}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
	// bad input cycle list
	list = [][]string{
		[]string{"A", "A"},
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
	// bad input cycle list
	list = [][]string{
		[]string{"A", "B"},
		[]string{"B", "A"},
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
	// bad input cycle list
	list = [][]string{
		[]string{"A", "B"},
		[]string{"B", "C"},
		[]string{"C", "A"},
	}
	assert.Equal(t, expected, GetSrcDestFromFlights(list))
}

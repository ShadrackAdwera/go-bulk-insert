package db

import (
	"context"
	"testing"

	"github.com/ShadrackAdwera/go-bulk-insert/utils"
	"github.com/stretchr/testify/require"
)

func createEcde(t *testing.T) EcdeCase {
	newEcde := CreateCaseParams{
		DateRep:                 utils.RandomString(12),
		Day:                     int32(utils.RandomInteger(1, 31)),
		Month:                   int32(utils.RandomInteger(1, 12)),
		Year:                    int32(utils.RandomInteger(2019, 2021)),
		Cases:                   utils.RandomInteger(1, 1000),
		Deaths:                  utils.RandomInteger(1, 1000),
		CountriesAndTerritories: utils.RandomString(2),
		GeoID:                   utils.RandomString(3),
		CountryTerritoryCode:    utils.RandomString(4),
		ContinentExp:            utils.RandomString(3),
		LoadDate:                utils.RandomString(12),
		IsoCountry:              utils.RandomString(3),
	}

	ecde, err := testQuery.CreateCase(context.Background(), newEcde)

	require.NoError(t, err)
	require.NotEmpty(t, ecde)
	require.NotZero(t, ecde.ID)

	return ecde
}

func TestCreateEcde(t *testing.T) {
	createEcde(t)
}

func TestListEcdes(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		createEcde(t)
	}

	ecdes, err := testQuery.ListCases(context.Background(), ListCasesParams{
		Limit:  5,
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, ecdes)
	require.Equal(t, len(ecdes), n)
}

package address

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestAddress(testingT *testing.T) {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	address := New(coreRand, localized)

	tC := gt.CreateTest(testingT)
	tC.Describe("Country", func() {
		tC.It("should return a country", func() {
			r := address.Country()
			gt.Expect(tC, &r).ToBeIn(en_US.Countries)
		})
	})

	tP := gt.CreateTest(testingT)
	tP.Describe("Postcode", func() {
		tP.It("should return a postcode", func() {
			r := address.Postcode()
			testutil.Output("Address.Postcode", r)
		})
	})

	tStateAbbr := gt.CreateTest(testingT)
	tStateAbbr.Describe("StateAbbr", func() {
		tStateAbbr.It("should return a state abbreviation", func() {
			r := address.StateAbbr()
			gt.Expect(tStateAbbr, &r).ToBeIn(en_US.StateAbbrs)
		})
	})

	tState := gt.CreateTest(testingT)
	tState.Describe("State", func() {
		tState.It("should return a state", func() {
			r := address.State()
			gt.Expect(tState, &r).ToBeIn(en_US.States)
		})
	})

	tCity := gt.CreateTest(testingT)
	tCity.Describe("City", func() {
		tCity.Test("CityName should return a city name", func() {
			r := address.CityName()
			gt.Expect(tCity, &r).ToBeIn(en_US.CityNames)
		})

		tCity.Test("CitySuffix should return a city suffix", func() {
			r := address.CitySuffix()
			gt.Expect(tCity, &r).ToBeIn(en_US.CitySuffixes)
		})

		tCity.Test("CityPrefix should return a city prefix", func() {
			r := address.CityPrefix()
			gt.Expect(tCity, &r).ToBeIn(en_US.CityPrefixes)
		})

		tCity.Test("City should return a city", func() {
			r := address.City()
			testutil.Output("Address.City", r)
		})
	})

	tStreet := gt.CreateTest(testingT)
	tStreet.Describe("Street", func() {
		tStreet.Test("StreetName should return a street name", func() {
			r := address.StreetName()
			gt.Expect(tStreet, &r).ToBeIn(en_US.StreetNames)
		})

		tStreet.Test("StreetSuffix should return a street suffix", func() {
			r := address.StreetSuffix()
			gt.Expect(tStreet, &r).ToBeIn(en_US.StreetSuffixes)
		})

		tStreet.Test("Street should return a street name", func() {
			r := address.Street()
			testutil.Output("Address.Street", r)
		})
	})

	tSec := gt.CreateTest(testingT)
	tSec.Describe("SecondaryAddress", func() {
		tSec.Test("BuildingNumber should return a building number", func() {
			r := address.BuildingNumber()
			testutil.Output("Address.BuildingNumber", r)
		})

		tSec.Test("SecondaryAddress should return a secondary address", func() {
			r := address.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	tStAddress := gt.CreateTest(testingT)
	tStAddress.Describe("StreetAddress", func() {
		tStAddress.It("should return a street address", func() {
			r := address.StreetAddress()
			testutil.Output("Address.StreetAddress", r)
		})
	})

	tAddr := gt.CreateTest(testingT)
	tAddr.Describe("Address", func() {
		tAddr.It("should return an address", func() {
			r := address.Address()
			testutil.Output("Address.Address", r)
		})
	})

	tLat := gt.CreateTest(testingT)
	tLat.Describe("Latitude", func() {
		tLat.It("should return a latitude", func() {
			r := address.Latitude()

			gt.Expect(tLat, &r).ToBe_(gt.Between(-90.0, 90.0))
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// length := testutil.GetDecimalLength(r)
			// gt.Expect(t15, &length).ToBe(6)
		})
	})

	tLon := gt.CreateTest(testingT)
	tLon.Describe("Longitude", func() {
		tLon.It("should return a longitude", func() {
			r := address.Longitude()

			gt.Expect(tLon, &r).ToBe_(gt.Between(-180.0, 180.0))
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// length := testutil.GetDecimalLength(r)
			// gt.Expect(t16, &length).ToBe(6)
		})
	})

	tCoord := gt.CreateTest(testingT)
	tCoord.Describe("LocalCoordinates", func() {
		tCoord.It("should return a coordinate", func() {
			lat, lon := address.LocalCoordinates()

			gt.Expect(tCoord, &lat).ToBe_(gt.Between(-90.0, 90.0))
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// latLength := testutil.GetDecimalLength(lat)
			// gt.Expect(t17, &latLength).ToBe(6)

			gt.Expect(tCoord, &lon).ToBe_(gt.Between(-180.0, 180.0))
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// lonLength := testutil.GetDecimalLength(lon)
			// gt.Expect(t17, &lonLength).ToBe(6)

		})
	})

	localizedJaJp := ja_JP.New()
	addressJaJp := New(coreRand, localizedJaJp)

	tPref := gt.CreateTest(testingT)
	tPref.Describe("Prefecture", func() {
		tPref.It("should return a prefecture", func() {
			r := addressJaJp.Prefecture()
			gt.Expect(tPref, &r).ToBeIn(ja_JP.Prefectures)
		})
	})

	tWard := gt.CreateTest(testingT)
	tWard.Describe("Ward", func() {
		tWard.Test("WardSuffix should return a ward suffix", func() {
			r := addressJaJp.WardSuffix()
			gt.Expect(tWard, &r).ToBeIn(ja_JP.WardSuffixes)
		})

		tWard.Test("WardName should return a ward name", func() {
			r := addressJaJp.WardName()
			gt.Expect(tWard, &r).ToBeIn(ja_JP.WardNames)
		})

		tWard.Test("Ward should return a ward", func() {
			r := addressJaJp.Ward()
			testutil.Output("Address.Ward", r)
		})
	})

	tArea := gt.CreateTest(testingT)
	tArea.Describe("Area", func() {
		tArea.Test("AreaNumber should return an area number", func() {
			r := addressJaJp.AreaNumber()
			testutil.Output("Address.AreaNumber", r)
		})

		tArea.Test("AreaName should return an area name", func() {
			r := addressJaJp.AreaName()
			gt.Expect(tArea, &r).ToBeIn(ja_JP.AreaNames)
		})

		tArea.Test("Area should return an area", func() {
			r := addressJaJp.Area()
			testutil.Output("Address.Area", r)
		})
	})

	tSecJaJp := gt.CreateTest(testingT)
	tSecJaJp.Describe("SecondaryAddress for ja_JP", func() {
		tSecJaJp.Test("BuildingName should return a building name", func() {
			r := addressJaJp.BuildingName()
			gt.Expect(tSecJaJp, &r).ToBeIn(ja_JP.BuildingNames)
		})
		tSecJaJp.Test("RoomNumber should return a room number", func() {
			r := addressJaJp.RoomNumber()
			testutil.Output("Address.RoomNumber", r)
		})
		tSecJaJp.Test("SecondaryAddress for ja_JP should return a secondary address", func() {
			r := addressJaJp.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	tAddrJaJp := gt.CreateTest(testingT)
	tAddrJaJp.Describe("Address for ja_JP", func() {
		tAddrJaJp.It("should return an address", func() {
			r := addressJaJp.Address()
			testutil.Output("Address.Address", r)
		})
	})
}

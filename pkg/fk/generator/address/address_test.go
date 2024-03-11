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

	t := gt.CreateTest(testingT)
	t.Describe("CitySuffix", func() {
		t.It("should return a city suffix", func() {
			r := address.CitySuffix()
			gt.Expect(t, &r).ToBeIn(en_US.CitySuffixes)
		})
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("CityPrefix", func() {
		t2.It("should return a city prefix", func() {
			r := address.CityPrefix()
			gt.Expect(t2, &r).ToBeIn(en_US.CityPrefixes)
		})
	})

	t3 := gt.CreateTest(testingT)
	t3.Describe("CityName", func() {
		t3.It("should return a city name", func() {
			r := address.CityName()
			gt.Expect(t3, &r).ToBeIn(en_US.CityNames)
		})
	})

	t6 := gt.CreateTest(testingT)
	t6.Describe("City", func() {
		t6.It("should return a city", func() {
			r := address.City()
			testutil.Output("Address.City", r)
		})
	})

	t4 := gt.CreateTest(testingT)
	t4.Describe("StreetSuffix", func() {
		t4.It("should return a street suffix", func() {
			r := address.StreetSuffix()
			gt.Expect(t4, &r).ToBeIn(en_US.StreetSuffixes)
		})
	})

	t7 := gt.CreateTest(testingT)
	t7.Describe("StreetName", func() {
		t7.It("should return a street name", func() {
			r := address.StreetName()
			testutil.Output("Address.StreetName", r)
		})
	})

	t5 := gt.CreateTest(testingT)
	t5.Describe("BuildingNumber", func() {
		t5.It("should return a building number", func() {
			r := address.BuildingNumber()
			testutil.Output("Address.BuildingNumber", r)
		})
	})

	t8 := gt.CreateTest(testingT)
	t8.Describe("SecondaryAddress", func() {
		t8.It("should return a secondary address", func() {
			r := address.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	t9 := gt.CreateTest(testingT)
	t9.Describe("StreetAddress", func() {
		t9.It("should return a street address", func() {
			r := address.StreetAddress()
			testutil.Output("Address.StreetAddress", r)
		})
	})

	t10 := gt.CreateTest(testingT)
	t10.Describe("Postcode", func() {
		t10.It("should return a postcode", func() {
			r := address.Postcode()
			testutil.Output("Address.Postcode", r)
		})
	})

	t11 := gt.CreateTest(testingT)
	t11.Describe("StateAbbr", func() {
		t11.It("should return a state abbreviation", func() {
			r := address.StateAbbr()
			gt.Expect(t11, &r).ToBeIn(en_US.StateAbbrs)
		})
	})

	t12 := gt.CreateTest(testingT)
	t12.Describe("State", func() {
		t12.It("should return a state", func() {
			r := address.State()
			gt.Expect(t12, &r).ToBeIn(en_US.States)
		})
	})

	t13 := gt.CreateTest(testingT)
	t13.Describe("Address", func() {
		t13.It("should return an address", func() {
			r := address.Address()
			testutil.Output("Address.Address", r)
		})
	})

	t14 := gt.CreateTest(testingT)
	t14.Describe("Country", func() {
		t14.It("should return a country", func() {
			r := address.Country()
			gt.Expect(t14, &r).ToBeIn(en_US.Countries)
		})
	})

	t15 := gt.CreateTest(testingT)
	t15.Describe("Latitude", func() {
		t15.It("should return a latitude", func() {
			r := address.Latitude()

			gt.Expect(t15, &r).ToBe_(gt.Between(-90.0), 90.0)
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// length := testutil.GetDecimalLength(r)
			// gt.Expect(t15, &length).ToBe(6)
		})
	})

	t16 := gt.CreateTest(testingT)
	t16.Describe("Longitude", func() {
		t16.It("should return a longitude", func() {
			r := address.Longitude()

			gt.Expect(t16, &r).ToBe_(gt.Between(-180.0), 180.0)
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// length := testutil.GetDecimalLength(r)
			// gt.Expect(t16, &length).ToBe(6)
		})
	})

	t17 := gt.CreateTest(testingT)
	t17.Describe("LocalCoordinates", func() {
		t17.It("should return a coordinate", func() {
			lat, lon := address.LocalCoordinates()

			gt.Expect(t17, &lat).ToBe_(gt.Between(-90.0), 90.0)
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// latLength := testutil.GetDecimalLength(lat)
			// gt.Expect(t17, &latLength).ToBe(6)

			gt.Expect(t17, &lon).ToBe_(gt.Between(-180.0), 180.0)
			// FIXME: testutil.GetDecimalLength is not accurate when the end of the number is 0
			// lonLength := testutil.GetDecimalLength(lon)
			// gt.Expect(t17, &lonLength).ToBe(6)

		})
	})

	localizedJaJp := ja_JP.New()
	addressJaJp := New(coreRand, localizedJaJp)
	t18 := gt.CreateTest(testingT)
	t18.Describe("Prefecture", func() {
		t18.It("should return a prefecture", func() {
			r := addressJaJp.Prefecture()
			gt.Expect(t18, &r).ToBeIn(ja_JP.Prefectures)
		})
	})

	t19 := gt.CreateTest(testingT)
	t19.Describe("WardSuffix", func() {
		t19.It("should return a ward suffix", func() {
			r := addressJaJp.WardSuffix()
			gt.Expect(t19, &r).ToBeIn(ja_JP.WardSuffixes)
		})
	})

	t20 := gt.CreateTest(testingT)
	t20.Describe("WardName", func() {
		t20.It("should return a ward name", func() {
			r := addressJaJp.WardName()
			gt.Expect(t20, &r).ToBeIn(ja_JP.WardNames)
		})
	})

	t21 := gt.CreateTest(testingT)
	t21.Describe("AreaName", func() {
		t21.It("should return an area name", func() {
			r := addressJaJp.AreaName()
			gt.Expect(t21, &r).ToBeIn(ja_JP.AreaNames)
		})
	})

	t22 := gt.CreateTest(testingT)
	t22.Describe("AreaNumber", func() {
		t22.It("should return an area number", func() {
			r := addressJaJp.AreaNumber()
			testutil.Output("Address.AreaNumber", r)
		})
	})

	t23 := gt.CreateTest(testingT)
	t23.Describe("BuildingName", func() {
		t23.It("should return a building name", func() {
			r := addressJaJp.BuildingName()
			gt.Expect(t23, &r).ToBeIn(ja_JP.BuildingNames)
		})
	})

	t26 := gt.CreateTest(testingT)
	t26.Describe("RoomNumber", func() {
		t26.It("should return a room number", func() {
			r := addressJaJp.RoomNumber()
			testutil.Output("Address.RoomNumber", r)
		})
	})

	t27 := gt.CreateTest(testingT)
	t27.Describe("SecondaryAddress for ja_JP", func() {
		t27.It("should return a secondary address", func() {
			r := addressJaJp.SecondaryAddress()
			testutil.Output("Address.SecondaryAddress", r)
		})
	})

	t28 := gt.CreateTest(testingT)
	t28.Describe("Address for ja_JP", func() {
		t28.It("should return an address", func() {
			r := addressJaJp.Address()
			testutil.Output("Address.Address", r)
		})
	})

	t24 := gt.CreateTest(testingT)
	t24.Describe("Ward", func() {
		t24.It("should return a ward", func() {
			r := addressJaJp.Ward()
			testutil.Output("Address.Ward", r)
		})
	})

	t25 := gt.CreateTest(testingT)
	t25.Describe("Area", func() {
		t25.It("should return an area", func() {
			r := addressJaJp.Area()
			testutil.Output("Address.Area", r)
		})
	})
}

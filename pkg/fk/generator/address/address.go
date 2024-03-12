package address

import (
	"github.com/yrichika/gfaker/pkg/fk/common/log"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Address struct {
	rand *core.Rand
	data *provider.Addresses
}

func New(rand *core.Rand, local *provider.Localized) *Address {
	return &Address{
		rand,
		local.Addresses,
	}
}

// example: 'town'
func (a *Address) CitySuffix() string {
	if len(a.data.CitySuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CitySuffixes)
}

func (a *Address) CityPrefix() string {
	if len(a.data.CityPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CityPrefixes)
}

func (a *Address) CityName() string {
	if len(a.data.CityNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.CityNames)
}

// example: 'Shieldsfurt'
func (a *Address) City() string {
	if len(a.data.CityFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.CityFormats)
	cityNameData := a.data.CreateCity(a)
	return util.RenderTemplate(format, cityNameData)
}

// example: 'Avenue'
func (a *Address) StreetSuffix() string {
	if len(a.data.StreetSuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StreetSuffixes)
}

func (a *Address) StreetName() string {
	if len(a.data.StreetNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StreetNames)
}

// example: 'Lindgren Dam'
func (a *Address) Street() string {
	if len(a.data.StreetFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.StreetFormats)
	streetNameData := a.data.CreateStreet(a)
	return util.RenderTemplate(format, streetNameData)
}

// example: '791'
func (a *Address) BuildingNumber() string {
	if len(a.data.BuildingNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.BuildingNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

// example: 'Apt. 160'
func (a *Address) SecondaryAddress() string {
	if len(a.data.SecondaryAddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.SecondaryAddressFormats)
	data := a.data.CreateSecondaryAddress(a)
	return util.RenderTemplate(format, data)
}

func (a *Address) StreetAddress() string {
	if len(a.data.StreetAddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.StreetAddressFormats)
	streetAddressData := a.data.CreateStreetAddress(a)
	return util.RenderTemplate(format, streetAddressData)
}

// example: '87678'
func (a *Address) Postcode() string {
	if len(a.data.Postcodes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.Postcodes)
	return a.rand.Str.AlphaDigitsLike(format)
}

// example: 'CA'
func (a *Address) StateAbbr() string {
	if len(a.data.StateAbbrs) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.StateAbbrs)
}

// example: 'California'
func (a *Address) State() string {
	if len(a.data.States) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.States)
}

// example: '026 Rolfson Summit\nPollichfurt, AZ 34986'
func (a *Address) Address() string {
	if len(a.data.AddressFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.AddressFormats)
	addressData := a.data.CreateAddress(a)
	return util.RenderTemplate(format, addressData)
}

// example: 'United States of America'
func (a *Address) Country() string {
	if len(a.data.Countries) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.Countries)
}

func (a *Address) Prefecture() string {
	if len(a.data.Prefectures) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.Prefectures)
}

func (a *Address) WardSuffix() string {
	if len(a.data.WardSuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.WardSuffixes)
}

func (a *Address) WardName() string {
	if len(a.data.WardNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.WardNames)
}

func (a *Address) Ward() string {
	if len(a.data.WardFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.WardFormats)
	wardData := a.data.CreateWard(a)
	return util.RenderTemplate(format, wardData)
}

func (a *Address) AreaName() string {
	if len(a.data.AreaNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.AreaNames)
}

func (a *Address) AreaNumber() string {
	if len(a.data.AreaNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	// TODO: 0丁目0番地0号のような表記も含まれてしまうが問題ないか?
	format := a.rand.Slice.StrElem(a.data.AreaNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

func (a *Address) Area() string {
	if len(a.data.AreaFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.AreaFormats)
	areaData := a.data.CreateArea(a)
	return util.RenderTemplate(format, areaData)

}

func (a *Address) BuildingName() string {
	if len(a.data.BuildingNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return a.rand.Slice.StrElem(a.data.BuildingNames)
}

func (a *Address) RoomNumber() string {
	if len(a.data.RoomNumbers) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	format := a.rand.Slice.StrElem(a.data.RoomNumbers)
	return a.rand.Str.AlphaDigitsLike(format)
}

// example: 35.785163
func (a *Address) Latitude() float64 {
	val := a.rand.Num.Float64Bt(-90, 90)
	return util.TruncateToPrecision(val, 6)
}

// example: -71.462048
func (a *Address) Longitude() float64 {
	val := a.rand.Num.Float64Bt(-180, 180)
	return util.TruncateToPrecision(val, 6)
}

// example: 35.785163, -71.462048
func (a *Address) LocalCoordinates() (float64, float64) {
	return a.Latitude(), a.Longitude()
}

package internet

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Internet struct {
	rand *core.Rand
	data *provider.Internets
}

func New(rand *core.Rand, global *provider.Global) *Internet {
	return &Internet{
		rand,
		global.Internets,
	}
}

func (i *Internet) FirstName() string {
	fName := i.rand.Slice.StrElem(i.data.FirstNames)
	without1Quote := strings.ReplaceAll(fName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) LastName() string {
	lName := i.rand.Slice.StrElem(i.data.LastNames)
	without1Quote := strings.ReplaceAll(lName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) UserName() string {
	baseFormat := i.rand.Slice.StrElem(i.data.UserNameFormats)
	format := i.rand.Str.AlphaDigitsLike(baseFormat)
	userName := i.data.CreateUserName(i)
	return util.RenderTemplate(format, userName)
}

func (i *Internet) DomainWord() string {
	lastName := i.rand.Slice.StrElem(i.data.LastNames)
	word := strings.ToLower(lastName)
	return word
}

func (i *Internet) Tld() string {
	return i.rand.Slice.StrElem(i.data.Tld)
}

// example: "howell.com"
func (i *Internet) DomainName() string {
	return i.DomainWord() + "." + i.Tld()
}

// example: "jude.borer@oberbrunner.com"
func (i *Internet) Email() string {
	format := i.rand.Slice.StrElem(i.data.EmailFormats)
	data := i.data.CreateEmail(i)
	return util.RenderTemplate(format, data)
}

// example: "18w50q2412G5Iky60QL"
func (i *Internet) Password() string {
	min := 8
	max := 20
	length := i.rand.Num.IntBt(min, max)
	like := strings.Repeat("*", length)
	return i.rand.Str.AlphaDigitsLike(like)
}

// TODO: Loremができてから戻って来る
func (i *Internet) Slug() string {
	//
	return ""
}

// TODO: Slugが必要なので、Loremができてから戻って来る
func (i *Internet) Url() string {
	//
	return ""
}

func (i *Internet) Ipv4() net.IP {
	var ipNum int
	if i.rand.Bool.Evenly() {
		ipNum = i.rand.Num.IntBt(-2147483648, -2)
	} else {
		ipNum = i.rand.Num.IntBt(16777216, 2147483647)
	}

	return long2ip(uint32(ipNum))
}

func (i *Internet) Ipv6() string {
	var res []string

	for index := 0; index < 8; index++ {
		res = append(res, fmt.Sprintf("%x", i.rand.Num.Intn(65536)))
	}

	return strings.Join(res, ":")

}

func (i *Internet) LocalIpv4() net.IP {
	lenIpBlocks := len(i.data.LocalIpBlocks)
	ipBlock := i.data.LocalIpBlocks[i.rand.Num.Intn(lenIpBlocks)]
	ipBlock0, _ := ip2long(ipBlock[0])
	ipBlock1, _ := ip2long(ipBlock[1])
	num := i.rand.Num.Int32Bt(int32(ipBlock0), int32(ipBlock1))
	return long2ip(uint32(num))
}

func (i *Internet) MacAddress() string {
	var mac []string

	for index := 0; index < 6; index++ {
		mac = append(mac, fmt.Sprintf("%02X", i.rand.Num.Intn(256)))
	}

	return strings.Join(mac, ":")
}

// REFACTOR: 名前を変える
func long2ip(long uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, long)
	return ip
}

// REFACTOR: 名前を変える
func ip2long(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ipStr)
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

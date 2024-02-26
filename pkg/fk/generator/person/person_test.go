package person

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestPerson(testingT *testing.T) {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	person := New(coreRand, localized)

	t1 := gt.CreateTest(testingT)
	t1.Describe("FirstNameMale", func() {
		t1.It("should return a male first name", func() {
			r := person.FirstNameMale()
			testutil.Output("Person.FirstNameMale", r)
		})
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("FirstNameFemale", func() {
		t2.It("should return a female first name", func() {
			r := person.FirstNameFemale()
			testutil.Output("Person.FirstNameFemale", r)
		})
	})

	t3 := gt.CreateTest(testingT)
	t3.Describe("FirstName", func() {
		t3.It("should return a first name", func() {
			r := person.FirstName()
			testutil.Output("Person.FirstName", r)
		})
	})

	t4 := gt.CreateTest(testingT)
	t4.Describe("LastName", func() {
		t4.It("should return a last name", func() {
			r := person.LastName()
			testutil.Output("Person.LastName", r)
		})
	})

	t5 := gt.CreateTest(testingT)
	t5.Describe("TitleMale", func() {
		t5.It("should return a male title", func() {
			r := person.TitleMale()
			testutil.Output("Person.TitleMale", r)
		})
	})

	t6 := gt.CreateTest(testingT)
	t6.Describe("TitleFemale", func() {
		t6.It("should return a female title", func() {
			r := person.TitleFemale()
			testutil.Output("Person.TitleFemale", r)
		})
	})

	t7 := gt.CreateTest(testingT)
	t7.Describe("Title", func() {
		t7.It("should return a title", func() {
			r := person.Title()
			testutil.Output("Person.Title", r)
		})
	})

	t8 := gt.CreateTest(testingT)
	t8.Describe("Suffix", func() {
		t8.It("should return a suffix", func() {
			r := person.Suffix()
			testutil.Output("Person.Suffix", r)
		})
	})

	t9 := gt.CreateTest(testingT)
	t9.Describe("MaleName", func() {
		t9.It("should return a male name", func() {
			r := person.MaleName()
			testutil.Output("Person.MaleName", r)
		})
	})

	t10 := gt.CreateTest(testingT)
	t10.Describe("FemaleName", func() {
		t10.It("should return a female name", func() {
			r := person.FemaleName()
			testutil.Output("Person.FemaleName", r)
		})
	})

	t11 := gt.CreateTest(testingT)
	t11.Describe("Name", func() {
		t11.It("should return a person's full name", func() {
			r := person.Name()
			testutil.Output("Person.Name", r)
		})
	})

	t12 := gt.CreateTest(testingT)
	t12.Describe("Ssn", func() {
		t12.It("should return a person's SSN", func() {
			r := person.Ssn()
			gt.Expect(t12, &r).ToMatchRegex(`^\d{3}-\d{2}-\d{4}$`)
		})
	})

	jaJp := ja_JP.New()
	personJaJp := New(coreRand, jaJp)
	t13 := gt.CreateTest(testingT)
	t13.Describe("FirstKanaNameMale", func() {
		t13.It("should return a male first kana name", func() {
			r := personJaJp.FirstKanaNameMale()
			testutil.Output("Person.FirstKanaNameMale", r)
		})
	})

	t14 := gt.CreateTest(testingT)
	t14.Describe("FirstKanaNameFemale", func() {
		t14.It("should return a female first kana name", func() {
			r := personJaJp.FirstKanaNameFemale()
			testutil.Output("Person.FirstKanaNameFemale", r)
		})
	})

	t15 := gt.CreateTest(testingT)
	t15.Describe("FirstKanaName", func() {
		t15.It("should return a first kana name", func() {
			r := personJaJp.FirstKanaName()
			testutil.Output("Person.FirstKanaName", r)
		})
	})

	t16 := gt.CreateTest(testingT)
	t16.Describe("LastKanaName", func() {
		t16.It("should return a last kana name", func() {
			r := personJaJp.LastKanaName()
			testutil.Output("Person.LastKanaName", r)
		})
	})

	t17 := gt.CreateTest(testingT)
	t17.Describe("MaleKanaName", func() {
		t17.It("should return a male kana name", func() {
			r := personJaJp.MaleKanaName()
			testutil.Output("Person.MaleKanaName", r)
		})
	})

	t18 := gt.CreateTest(testingT)
	t18.Describe("FemaleKanaName", func() {
		t18.It("should return a female kana name", func() {
			r := personJaJp.FemaleKanaName()
			testutil.Output("Person.FemaleKanaName", r)
		})
	})

	tKn := gt.CreateTest(testingT)
	tKn.Describe("KanaName", func() {
		tKn.It("should return a kana name", func() {
			r := personJaJp.KanaName()
			testutil.Output("Person.KanaName", r)
		})
	})
}

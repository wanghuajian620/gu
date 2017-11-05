package fo_FO

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fo_FO struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'fo_FO' locale
func New() locales.Translator {
	return &fo_FO{
		locale:                 "fo_FO",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "−",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mar.", "apr.", "mai", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "mars", "apríl", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
		daysAbbreviated:        []string{"sun.", "mán.", "týs.", "mik.", "hós.", "frí.", "ley."},
		daysNarrow:             []string{"S", "M", "T", "M", "H", "F", "L"},
		daysShort:              []string{"su.", "má.", "tý.", "mi.", "hó.", "fr.", "le."},
		daysWide:               []string{"sunnudagur", "mánadagur", "týsdagur", "mikudagur", "hósdagur", "fríggjadagur", "leygardagur"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"f.Kr.", "e.Kr."},
		erasNarrow:             []string{"fKr", "eKr"},
		erasWide:               []string{"fyri Krist", "eftir Krist"},
		timezones:              map[string]string{"NZDT": "Nýsæland summartíð", "WART": "Vestur Argentina vanlig tíð", "CHADT": "Chatham summartíð", "BOT": "Bolivia tíð", "HADT": "Hawaii-Aleutian summartíð", "HNNOMX": "Northwest Mexico vanlig tíð", "HEEG": "Eystur grønlendsk summartíð", "SGT": "Singapor tíð", "CHAST": "Chatham vanlig tíð", "BT": "Butan tíð", "WITA": "Mið Indonesia tíð", "AST": "Atlantic vanlig tíð", "COST": "Kolombia summartíð", "GYT": "Gujana tíð", "ACST": "mið Avstralia vanlig tíð", "WIB": "Vestur Indonesia tíð", "HEPM": "St. Pierre & Miquelon summartíð", "HECU": "Cuba summartíð", "MST": "Mountain vanlig tíð", "VET": "Venesuela tíð", "∅∅∅": "Azorurnar summartíð", "AEDT": "eystur Avstralia summartíð", "WAST": "Vesturafrika summartíð", "COT": "Kolombia vanlig tíð", "GMT": "Greenwich Mean tíð", "AWST": "vestur Avstralia vanlig tíð", "UYST": "Uruguai summartíð", "MESZ": "Miðevropa summartíð", "HAST": "Hawaii-Aleutian vanlig tíð", "HNOG": "Vestur grønlendsk vanlig tíð", "CLT": "Kili vanlig tíð", "HNPMX": "Mexican Pacific vanlig tíð", "PDT": "Pacific summartíð", "TMST": "Turkmenistan summartíð", "JDT": "Japan summartíð", "EST": "Eastern vanlig tíð", "EDT": "Eastern summartíð", "CST": "Central vanlig tíð", "ADT": "Atlantic summartíð", "ACDT": "mið Avstralia summartíð", "ACWDT": "miðvestur Avstralia summartíð", "ARST": "Argentina summartíð", "EAT": "Eysturafrika tíð", "HAT": "Newfoundland summartíð", "HKST": "Hong Kong summartíð", "HNPM": "St. Pierre & Miquelon vanlig tíð", "OEZ": "Eysturevropa vanlig tíð", "LHDT": "Lord Howe summartíð", "WESZ": "Vesturevropa summartíð", "CDT": "Central summartíð", "UYT": "Uruguai vanlig tíð", "MEZ": "Miðevropa vanlig tíð", "MDT": "Mountain summartíð", "WARST": "Vestur Argentina summartíð", "ART": "Argentina vanlig tíð", "HNEG": "Eystur grønlendsk vanlig tíð", "GFT": "Franska Gujana tíð", "CAT": "Miðafrika tíð", "TMT": "Turkmenistan vanlig tíð", "LHST": "Lord Howe vanlig tíð", "HENOMX": "Northwest Mexico summartíð", "JST": "Japan vanlig tíð", "HEOG": "Vestur grønlendsk summartíð", "SRT": "Surinam tíð", "MYT": "Malaisia tíð", "NZST": "Nýsæland vanlig tíð", "WAT": "Vesturafrika vanlig tíð", "HKT": "Hong Kong vanlig tíð", "AKST": "Alaska vanlig tíð", "AKDT": "Alaska summartíð", "ECT": "Ekvador tíð", "WEZ": "Vesturevropa vanlig tíð", "WIT": "Eystur Indonesia tíð", "HNCU": "Cuba vanlig tíð", "ACWST": "miðvestur Avstralia vanlig tíð", "SAST": "Suðurafrika vanlig tíð", "HNT": "Newfoundland vanlig tíð", "AWDT": "vestur Avstralia summartíð", "OESZ": "Eysturevropa summartíð", "IST": "India tíð", "CLST": "Kili summartíð", "ChST": "Chamorro vanlig tíð", "AEST": "eystur Avstralia vanlig tíð", "HEPMX": "Mexican Pacific summartíð", "PST": "Pacific vanlig tíð"},
	}
}

// Locale returns the current translators string locale
func (fo *fo_FO) Locale() string {
	return fo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fo_FO'
func (fo *fo_FO) PluralsCardinal() []locales.PluralRule {
	return fo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fo_FO'
func (fo *fo_FO) PluralsOrdinal() []locales.PluralRule {
	return fo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fo_FO'
func (fo *fo_FO) PluralsRange() []locales.PluralRule {
	return fo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fo_FO'
func (fo *fo_FO) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fo_FO'
func (fo *fo_FO) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fo_FO'
func (fo *fo_FO) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fo *fo_FO) MonthAbbreviated(month time.Month) string {
	return fo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fo *fo_FO) MonthsAbbreviated() []string {
	return fo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fo *fo_FO) MonthNarrow(month time.Month) string {
	return fo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fo *fo_FO) MonthsNarrow() []string {
	return fo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fo *fo_FO) MonthWide(month time.Month) string {
	return fo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fo *fo_FO) MonthsWide() []string {
	return fo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fo *fo_FO) WeekdayAbbreviated(weekday time.Weekday) string {
	return fo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fo *fo_FO) WeekdaysAbbreviated() []string {
	return fo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fo *fo_FO) WeekdayNarrow(weekday time.Weekday) string {
	return fo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fo *fo_FO) WeekdaysNarrow() []string {
	return fo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fo *fo_FO) WeekdayShort(weekday time.Weekday) string {
	return fo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fo *fo_FO) WeekdaysShort() []string {
	return fo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fo *fo_FO) WeekdayWide(weekday time.Weekday) string {
	return fo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fo *fo_FO) WeekdaysWide() []string {
	return fo.daysWide
}

// Decimal returns the decimal point of number
func (fo *fo_FO) Decimal() string {
	return fo.decimal
}

// Group returns the group of number
func (fo *fo_FO) Group() string {
	return fo.group
}

// Group returns the minus sign of number
func (fo *fo_FO) Minus() string {
	return fo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fo_FO' and handles both Whole and Real numbers based on 'v'
func (fo *fo_FO) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fo.minus) - 1; j >= 0; j-- {
			b = append(b, fo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fo_FO' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fo *fo_FO) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fo.minus) - 1; j >= 0; j-- {
			b = append(b, fo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fo.percentSuffix...)

	b = append(b, fo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fo_FO'
func (fo *fo_FO) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fo.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fo.minus) - 1; j >= 0; j-- {
			b = append(b, fo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fo.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fo_FO'
// in accounting notation.
func (fo *fo_FO) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fo.currencies[currency]
	l := len(s) + len(symbol) + 8 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, fo.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fo.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fo.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fo_FO'
func (fo *fo_FO) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}

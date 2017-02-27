// Refactor a ledger printer.

package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const testVersion = 4

type Entry struct {
	Date        string // "Y-m-date"
	Description string
	Change      int // in cents
}

var currencyMap = map[string]string{
	"EUR": "€", "USD": "$",
}

var localeMap = map[string]struct {
	header, dateFmt, numFmt, thouSep, posFmt, negFmt string
}{
	"nl-NL": {fmt.Sprintf("%-10s | %-25s | Verandering\n", "Datum",
		"Omschrijving"), "%s-%s-%s", // day-mont-year
		"%s %s,%s", ".", "%s ", "%s-"},

	"en-US": {fmt.Sprintf("%-10s | %-25s | Change\n", "Date", "Description"),
		"%[2]s/%[1]s/%[3]s", // month/day/year
		"%s%s.%s", ",", "%s ", "(%s)"},
}

func FormatChange(currency string, locale string, change int) string {
	currencySymbol := currencyMap[currency]
	negative := false
	cents := change
	if cents < 0 {
		cents = -cents
		negative = true
	}
	centsStr := fmt.Sprintf("%03d", cents)
	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append([]string{rest[len(rest)-3:]}, parts...)
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append([]string{rest}, parts...)
	}

	format := localeMap[locale]
	number := ""

	number = fmt.Sprintf(format.numFmt,
		currencySymbol, strings.Join(parts, format.thouSep),
		centsStr[len(centsStr)-2:])

	if negative {
		number = fmt.Sprintf(format.negFmt, number)
	} else {
		number = fmt.Sprintf(format.posFmt, number)
	}
	return number
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	// Error check
	if _, ok := currencyMap[currency]; !ok {
		return "", errors.New("Unsupported currency")
	}
	if _, ok := localeMap[locale]; !ok {
		return "", errors.New("Unsupported locale")
	}
	for _, entry := range entries {
		if len(entry.Date) != 10 ||
			entry.Date[4] != '-' || entry.Date[7] != '-' {
			return "", errors.New("Incorrect date format")
		}
	}

	if len(entries) == 0 {
		locale = "en-US"
	}

	// sort by date, description, change, ascending
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Sort(entryList(entriesCopy))

	// Parallelism is always a great idea, and it provies to be faster
	// at local env. Even tho eventually the results are collceted sequentially
	// which meaning the formatting takes long
	co := make(chan struct {
		i int // to maintain the sequence
		s string
	})

	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			year, month, day := entry.Date[0:4], entry.Date[5:7], entry.Date[8:10]
			var date string
			date = fmt.Sprintf(localeMap[locale].dateFmt, day, month, year)

			desc := entry.Description
			if len(desc) > 25 {
				desc = desc[:22] + "..."
			} else {
				desc = fmt.Sprintf("%-25s", desc)
			}

			money := FormatChange(currency, locale, entry.Change)

			co <- struct {
				i int
				s string
			}{i: i, s: fmt.Sprintf("%-10s | %s | %13s\n", date, desc, money)}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		ss[v.i] = v.s
	}

	s := localeMap[locale].header // table head
	s = s + strings.Join(ss, "")

	return s, nil
}

type entryList []Entry

func (e entryList) Len() int      { return len(e) }
func (e entryList) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e entryList) Less(i, j int) bool {
	switch {
	case e[i].Date < e[j].Date:
		return true
	case e[i].Date > e[j].Date:
		return false
	case e[i].Description < e[j].Description:
		return true
	case e[i].Description > e[j].Description:
		return false
	case e[i].Change < e[j].Change:
		return true
	case e[i].Change < e[j].Change:
		return false
	default:
		return false
	}
	return false
}

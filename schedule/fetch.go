package schedule

import (
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

const url = "https://www.keltas.lt/tvarkarastis/"

func Fetch() (ferries []Ferry, err error) {
	c := colly.NewCollector()
	ferries = make([]Ferry, 0)

	c.OnHTML("div.schedule-bottom", func(e *colly.HTMLElement) {
		f := trim(e.ChildText("div.schedule-info-block h2"))
		ferry := Ferry{
			Title:     strings.ToUpper(f),
			Schedules: make([]Schedule, 0, 2),
		}
		defer func() { ferries = append(ferries, ferry) }()

		e.ForEach("div.col-md-4", func(i int, e *colly.HTMLElement) {
			title := trim(e.ChildText("h4"))
			schedule := Schedule{
				Title: strings.Join(strings.SplitN(title, " ", 3)[:2], " "),
				Table: make([]Time, 0),
			}
			defer func() { ferry.Schedules = append(ferry.Schedules, schedule) }()

			e.ForEach("p.schedule-time", func(i int, e *colly.HTMLElement) {
				h := trim(e.ChildText("span.hour-class"))
				m := trim(e.ChildText("span.minute-class"))
				bits := regexp.MustCompile("[^0-9*]+").Split(m, -1)
				m = strings.Join(bits, " ")

				schedule.Table = append(schedule.Table, Time{
					Hour:    h,
					Minutes: m,
				})
			})
		})
	})

	err = c.Visit(url)

	return
}

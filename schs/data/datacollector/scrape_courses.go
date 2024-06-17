package main

import (
	"fmt"
	"net/url"
	"projects/saelections/pkg/sysout"

	"github.com/gocolly/colly"
)

var (
	c *colly.Collector
)

func init() {
	c = NewCollector()

	d := NewCollector()

	cc := CourseCollector{}

	cc.OnCourseFound(c, d)

	cc.OnCourseDurationFound(d)

	cc.OnCourseLongReqFound(d)
}

func ScrapeCourses() {
	xurl, err := url.Parse("https://www.findamasters.com/masters-degrees/computer-science/?1002")
	if err != nil {
		sysout.Print(err)
	}

	for i := 1; i < 2; i++ {
		q := xurl.Query()
		// q.Set("Keywords", strings.ToLower(_query))
		q.Set("PG", fmt.Sprint(i))
		xurl.RawQuery = q.Encode()
		c.Visit(xurl.String())
	}
}

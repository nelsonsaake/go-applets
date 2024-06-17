package main

import (
	"fmt"
	"projects/saelections/pkg/sysout"

	"github.com/gocolly/colly"
)

const (
	QuerySelectorCourse = "#SearchResults > div.courses.col-24.px-0 > div > div > div.row.m-0 > div.col-md-18.col-lg-19.order-md-1.col-24.p-0.pr-md-3.course-content-info"

	QuerySelectorCourseURL = "h3 > a.courseLink.text-dark"

	QuerySelectorCourseName = "h3 > a.courseLink.text-dark > u"

	QuerySelectorCourseInst = "h4 > a.instLink.text-secondary"

	QuerySelectorCourseDept = "h4 > a.deptLink.text-secondary"

	QuerySelectorCourseClassStyle = "div.course-icon-area.row.mx-n1 > a.hoverTitle.subButton.grey.dist.badge.badge-light.card-badge.p-2.m-1.font-weight-light > span"

	QuerySelectorCourseStyle = "div.course-icon-area.row.mx-n1 > a.hoverTitle.subButton.grey.full.badge.badge-light.card-badge.p-2.m-1.font-weight-light > span"

	QuerySelectorCourseDuration = "#MainArea > div.full-width-container.container-fluid.full-width-white.full-width-main > div > div > div.page-content.col-24.col-md-16.col-lg-17.px-0.px-md-auto.pr-lg-3 > div > div.key-info__outer.row.pt-3 > div > span:nth-child(4)"

	QuerySelectorCourseLongReq = "#MainArea > div.full-width-container.container-fluid.full-width-white.full-width-main > div > div > div.page-content.col-24.col-md-16.col-lg-17.px-0.px-md-auto.pr-lg-3 > div > div.course-sections.course-sections__entry-requirements.tight.col-xs-24 > div > p"

	QuerySelectorNextPage = "#SearchResults > div.searchOpts.bot > div > div > ul > li:nth-child(2) > a"

	QuerySelectorCourseDurationPartime = "#SearchResults > div.courses.col-24.px-0 > div:nth-child > div > div.row.m-0 > div.col-md-18.col-lg-19.order-md-1.col-24.p-0.pr-md-3.course-content-info > div.course-icon-area.row.mx-n1 > a.hoverTitle.subButton.grey.part.badge.badge-light.card-badge.p-2.m-1.font-weight-light > span"
)

type CourseCollector struct{}

func (c *CourseCollector) Domain(e *colly.HTMLElement) string {
	return fmt.Sprint(e.Request.URL.Scheme, "://", e.Request.URL.Host)
}

func (c *CourseCollector) CourseURL(e *colly.HTMLElement) string {
	return c.Domain(e) + e.ChildAttr(QuerySelectorCourseURL, "href")
}

func (c *CourseCollector) School(e *colly.HTMLElement) string {
	return e.ChildText(QuerySelectorCourseInst) + ", " + e.ChildText(QuerySelectorCourseDept)
}

func (c *CourseCollector) Name(e *colly.HTMLElement) string {
	return e.ChildText(QuerySelectorCourseName)
}

func (c *CourseCollector) ClassStyle(e *colly.HTMLElement) string {
	return e.ChildText(QuerySelectorCourseClassStyle)
}

func (c *CourseCollector) CourseStyle(e *colly.HTMLElement) string {
	return e.ChildText(QuerySelectorCourseStyle)
}

func (c *CourseCollector) ElementURL(e *colly.HTMLElement) string {
	return e.Request.URL.String()
}

func (c *CourseCollector) OnCourseFound(cc *colly.Collector, d *colly.Collector) {
	cc.OnHTML(QuerySelectorCourse, func(e *colly.HTMLElement) {
		_, err := RepoCourseSave(c.CourseURL(e), Name(c.Name(e)), School(c.School(e)),
			ClassStyle(c.ClassStyle(e)), ProgramStyle(c.CourseStyle(e)), DataSourceURL(e.Request.URL.String()))
		if err != nil {
			sysout.Print(err)
		}
		d.Visit(c.CourseURL(e))
	})
	cc.OnHTML(QuerySelectorCourseDurationPartime, func(e *colly.HTMLElement) {
		if _, err := RepoCourseSave(c.CourseURL(e), ProgramStyle(e.Text)); err != nil {
			sysout.Print(err)
		}
	})
}

func (c *CourseCollector) OnCourseDurationFound(cc *colly.Collector) {
	cc.OnHTML(QuerySelectorCourseDuration, func(h *colly.HTMLElement) {
		_, err := RepoCourseSave(c.ElementURL(h), Duration(h.Text))
		if err != nil {
			sysout.Print(err)
		}
	})
}

func (c *CourseCollector) OnCourseLongReqFound(cc *colly.Collector) {
	cc.OnHTML(QuerySelectorCourseLongReq, func(h *colly.HTMLElement) {
		_, err := RepoCourseSave(c.ElementURL(h), LongReq(h.Text))
		if err != nil {
			sysout.Print(err)
		}
	})
}

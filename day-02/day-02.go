package main

import (
	"github.com/samber/lo"
	"slices"
	"strings"
	. "utils"
)

func puzzle1(input []string) (result int) {
	reports := getReports(input)

	for _, report := range reports {
		increasing := report[1] > report[0]

		include := true
		for j := 0; j < len(report)-1; j++ {
			if increasing {
				if report[j+1] <= report[j] || report[j+1]-report[j] > 3 {
					include = false
					break
				}

				continue
			}

			if report[j+1] >= report[j] || report[j]-report[j+1] > 3 {
				include = false
				break
			}
		}

		if include {
			result++
		}
	}

	return
}

func puzzle2(input []string) (result int) {
	reports := lo.Map(getReports(input), func(report Report, _ int) ReportWithDampener {
		return ReportWithDampener{
			r:          report,
			isDampened: false,
		}
	})

	for _, report := range reports {
		if checkReport(report) {
			result++
		}
	}

	return
}

func checkReport(report ReportWithDampener) bool {
	increasing := report.r[1] > report.r[0]

	for j := 0; j < len(report.r)-1; j++ {
		if increasing {
			if report.r[j+1] <= report.r[j] || report.r[j+1]-report.r[j] > 3 {
				return checkDampenedReport(report, j)
			}
		} else if report.r[j+1] >= report.r[j] || report.r[j]-report.r[j+1] > 3 {
			return checkDampenedReport(report, j)
		}
	}

	return true
}

func checkDampenedReport(report ReportWithDampener, j int) bool {
	if report.isDampened {
		return false
	}

	r := make(Report, len(report.r))

	copy(r, report.r)
	if checkReport(ReportWithDampener{
		r:          slices.Delete(r, j, j+1),
		isDampened: true,
	}) {
		return true
	}

	copy(r, report.r)
	if checkReport(ReportWithDampener{
		r:          slices.Delete(r, j+1, j+2),
		isDampened: true,
	}) {
		return true
	}

	if j != 0 {
		copy(r, report.r)
		if checkReport(ReportWithDampener{
			r:          slices.Delete(r, j-1, j),
			isDampened: true,
		}) {
			return true
		}
	}

	return false
}

type Report []int

type ReportWithDampener struct {
	r          Report
	isDampened bool
}

func getReports(input []string) (reports []Report) {
	for _, line := range input {
		split := strings.Split(line, " ")

		reports = append(reports, MapToInts(split))
	}

	return
}

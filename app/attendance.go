package app

import (
	"bot-for-academic-attendance/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetAttendance(client *http.Client) {
	absenURL := "https://akademik.polban.ac.id/ajar/absen"
	req, err := http.NewRequest("GET", absenURL, nil)
	utils.CheckError(err)

	resp, err := client.Do(req)
	utils.HandleResponse(resp, err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	utils.CheckError(err)

	attendanceData := ParseAttendance(doc)
	PerformAttendance(client, attendanceData)

}

func ParseAttendance(doc *goquery.Document) []string {
	var array []string
	doc.Find("#jadwal tbody tr").Each(func(i int, s *goquery.Selection) {
		kls, _ := s.Find("input").Attr("value")
		dsn := strings.Split(s.Find("td").Eq(1).Text(), "-")[0]
		mk := s.Find("td").Eq(2).Text()
		namamk := s.Find("td").Eq(3).Text()
		tp := s.Find("td").Eq(4).Text()
		ja := s.Find("td").Eq(5).Text()
		jb := s.Find("td").Eq(6).Text()
		jam := s.Find("td").Eq(7).Text()
		stat, _ := s.Find("td a").Attr("class")
		status := "Unknown"
		if stat == "btn btn-success" {
			status = "Already attended (Green)"
		} else if stat == "btn btn-warning" {
			status = "Not able to attend yet (Yellow)"
		} else if stat == "btn btn-danger" {
			status = "Not able to attend yet (Red)"
		} else if stat == "btn btn-info simpan_awal" {
			status = "Not attended yet (Blue)"
			obj := map[string]string{
				"kls":    kls,
				"dsn":    dsn,
				"mk":     mk,
				"namamk": namamk,
				"tp":     tp,
				"ja":     ja,
				"jb":     jb,
				"jam":    jam,
			}
			array = append(array, fmt.Sprint(obj))
		}

		fmt.Printf("%-31s | %-6s | %s\n", namamk, jam, status)
	})

	return array
}

func PerformAttendance(client *http.Client, attendanceData []string) {
	for _, row := range attendanceData {
		var obj map[string]string
		err := json.Unmarshal([]byte(row), &obj)
		utils.CheckError(err)
		fmt.Printf("Performing attendance for the course %s at %s o'clock\n", obj["namamk"], obj["jam"])
		absenURL := "https://akademik.polban.ac.id/ajar/absen/absensi_awal"
		absenData := url.Values{
			"ja":  {obj["ja"]},
			"jb":  {obj["jb"]},
			"mk":  {obj["mk"]},
			"dsn": {obj["dsn"]},
			"tp":  {obj["tp"]},
			"kls": {obj["kls"]},
		}
		req, err := CreateAttendanceRequest(absenURL, absenData)
		utils.CheckError(err)

		_, err = client.Do(req)
		utils.CheckError(err)

	}
}

func CreateAttendanceRequest(absenURL string, absenData url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", absenURL, strings.NewReader(absenData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

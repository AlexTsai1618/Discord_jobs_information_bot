package services

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	// "io/ioutil"

	// "time"

	"github.com/gocolly/colly"
)

func Scrapper() string{
	
	c := colly.NewCollector(
		colly.AllowedDomains("www.riversmall.com"),
	)
	// jobs := []Job{}
	// detailCollector := c.Clone()
	jobs := make([]Linkedin_data,0,600)

	c.OnHTML("ul.list-group", func(e *colly.HTMLElement) {
		
		e.ForEach("a[href]", func(_ int, el *colly.HTMLElement){
			job := Linkedin_data{
				Company_name: el.ChildText("p.companyName"),
				Job_url: el.Attr("href"),
				Job_title: el.ChildText("div.justify-content-between > p.jobTitle"),
				Job_location: el.ChildText("div > small.location"),
				Posted_date: el.ChildText("div > small.postDate"),
			}
			jobs = append(jobs,job)
		
		})				
	})
	

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})


	// Dump json to the standard output
	c.Visit("http://www.riversmall.com/intern.html")

	output_file,file_err := json.MarshalIndent(jobs, "", " ")
	
	if file_err != nil {
		fmt.Println("Unable to create json file")
		return ""
	}
	time_stemps:= time.Now().Format("2006-01-02")
	file_path := "data/clean_data/reiversmall/" + time_stemps +".json"	
	err :=ioutil.WriteFile(file_path,output_file, 0644)
	if err == nil{
		fmt.Println("Data written to file successfully")
	}

	return file_path
}
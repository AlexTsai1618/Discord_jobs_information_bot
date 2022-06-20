package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"time"
	"os"
)

func Job_search(job_title string,host_title string,location string)string{

	url := "https://" + host_title + "/"
	data := map[string]string{
		"search_terms": job_title,
		"location": location,
		"page": "1",
		"fetch_full_text": "no",
	}
	json_data, err := json.Marshal(data)

	if err != nil{
		fmt.Println("Json file encoding error!")
		log.Fatal(err)
	}
	api_key := os.Getenv("API_KEY")
	client := &http.Client{}
	request, _ := http.NewRequest("POST",url,bytes.NewBuffer(json_data))
	request.Header.Add("content-type","application/json")
	request.Header.Add("X-RapidAPI-Key",api_key)
	request.Header.Add("X-RapidAPI-Host",host_title)

	response, err := client.Do(request)
	if err != nil{
		fmt.Println("Client error!")
		log.Fatal(err)
	}
	
	body,_ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body",string(body))
	
	// json_file := json.MarshalIndent(body)
	
	time_stemps:= time.Now().Format("2006-01-02")
//  
	
	file_name := "data/raw_data/linkedin/" + time_stemps +".json"
 	ioutil.WriteFile(file_name,body, 0644)
	return string(file_name)
}

func Linkedin_parser(file_name string)string{
	file,error := ioutil.ReadFile(file_name)
	if error != nil{
		fmt.Println("Can't read file (line 56)",error)
		log.Fatal(error)
	}
	
	// decode data
	var data []Linkedin_data
	error_message := json.Unmarshal(file, &data)
	if error_message != nil{
		log.Fatal(error_message)
	}

	// turn date to interger
	re := regexp.MustCompile("[0-9]+")
	for i:= 0; i < len(data);i++{
		temp_date := re.FindAllString(data[i].Posted_date,-1)[0] + re.FindAllString(data[i].Posted_date,-1)[1] + re.FindAllString(data[i].Posted_date,-1)[2]
		interger_date,_ := strconv.Atoi(temp_date)
		data[i].Posted_date_int = interger_date
	}

	// sort by date time
	sort.SliceStable(data, func(index1,index2 int) bool{
		return data[index1].Posted_date_int > data[index2].Posted_date_int
	})

	output_file,_ := json.MarshalIndent(data, "","")
	time_stemps:= time.Now().Format("2006-01-02")
	file_path := "data/clean_data/linkedin/" + time_stemps +".json"
	_ = ioutil.WriteFile(file_path,output_file, 0644)
	fmt.Printf(file_path+" successfully writed")
	return string(file_path)
}

type local_data struct{
	Jobs_url []string `json:"jobs_url"`
}

func Linkedin_bot_update(file_name string){
	file,error_message := ioutil.ReadFile(file_name)
	if error_message != nil{
		log.Fatal(error_message)
	}
	var data []Linkedin_data
	error_message_2 := json.Unmarshal(file, &data)
	if error_message_2 != nil{
		log.Fatal("line 97",error_message_2)
	}
	visited_file_path := "data/posted_data/visited.json"
	job_visited, _ := ioutil.ReadFile(visited_file_path)
	
	var local local_data
	json.Unmarshal(job_visited, &local)

	for i:=0; i < len(data); i++{
		company_name := data[i].Company_name
		job_title := data[i].Job_title
		job_url := data[i].Job_url
		job_location := data[i].Job_location
		if !contains(local.Jobs_url,job_url){
			local.Jobs_url = append(local.Jobs_url,job_url)
			bot(company_name,job_title,job_url,job_location)
		}
		continue
	}

	output_visited_file,_ := json.MarshalIndent(local, "","")
	_ = ioutil.WriteFile(visited_file_path,output_visited_file, 0644)
	if error_message != nil{
		log.Fatal(error_message)
	}
	fmt.Println("success!")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
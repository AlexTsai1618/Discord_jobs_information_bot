package main

import (
	"Discord_jobs_information_bot/services"
)


func main(){
	Linkedin_jobs("2024 summer software engineer intern")
	Riversmall_scrap()
	
}
func Riversmall_scrap(){
	clean_data_path := services.Scrapper()
	services.Linkedin_bot_update(clean_data_path)
}
func Linkedin_jobs(job_title string){
	file_name := services.Job_search(job_title,"linkedin-jobs-search.p.rapidapi.com","3031")
	clean_data_path := services.Linkedin_parser(file_name)
	services.Linkedin_bot_update(clean_data_path)
}
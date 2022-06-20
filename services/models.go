package services
// type models interface{
// 	Linkedin_data
// }
type Linkedin_data struct{
	
	Company_name    string    `json:"company_name"`
	Job_title       string    `json:"job_title"`	
	Job_url         string    `json:"linkedin_job_url_cleaned"`
	Job_location    string    `json:"job_location"`
	Posted_date     string    `json:"posted_date"`
	Posted_date_int int
	Readme_update	string
}
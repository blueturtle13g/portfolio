package main

import (
	"net/http"
	"fmt"
	"regexp"
)

type Msg struct{
	Id int
	Name, Email, Text, CreatedOn string
}


func (msg Msg) Validate() (warnings []string) {
	if len(msg.Name) < 3 || len(msg.Name) > 20 {
		warnings = append(warnings, "The length of your name can't be less than 3 or more than 20 characters.")
	}
	if len(msg.Text) == 0 {
		warnings = append(warnings, "You Are Sending Us An Empty Message, Please Check.")
	}
	if er := mailValidation(msg.Email); er != "" {
		warnings = append(warnings, er)
	}
	return warnings
}

func insertMsg(msg Msg) (ItsId int) {
	err := DB.QueryRow(
		"INSERT INTO msgs(name,email,text,createdOn) VALUES($1,$2,$3,$4) returning id;",
		msg.Name, msg.Email, msg.Text, msg.CreatedOn).Scan(&ItsId)
	if err != nil {
		fmt.Println(err)
	}
	return ItsId
}


func mailValidation(email string) (er string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email) {
		return ""
	}
	return "you've entered an invalid email address, please check it."
}

func Index(w http.ResponseWriter, r *http.Request){
	data := map[string]string{"Title": "Reza Shoja"}
	if err := tpl.ExecuteTemplate(w, "index.gohtml", data); err != nil{
		fmt.Println(err)
	}
}
func About(w http.ResponseWriter, r *http.Request){
	data := map[string]string{"Title": "About"}
	if err := tpl.ExecuteTemplate(w, "about.gohtml", data); err != nil{
		fmt.Println(err)
	}
}
func Contact(w http.ResponseWriter, r *http.Request){
	data := map[string]interface{}{"Title": "Contact"}
	if r.Method == "POST"{
		name := r.FormValue("name")
		email := r.FormValue("email")
		text := r.FormValue("text")

		msg := Msg{0, name, email, text, getNow()}
		if warnings := msg.Validate(); len(warnings) > 0{
			data["Errs"] = warnings
			data["Name"] = name
			data["Email"] = email
			data["Text"] = text
			if err := tpl.ExecuteTemplate(w, "contact.gohtml", data); err != nil{
				fmt.Println(err)
			}
			return
		}

		if msgId := insertMsg(msg); msgId < 1{
			fmt.Println("inserting msgs failed.")
		}
		data["Cong"] = "Congratulations dear "+name+", your message has been sent."
	}
	if err := tpl.ExecuteTemplate(w, "contact.gohtml", data); err != nil{
		fmt.Println(err)
	}
}

func Tutorials(w http.ResponseWriter, r *http.Request){
	data := map[string]string{"Title": "Tutorials"}

	if err := tpl.ExecuteTemplate(w, "tutorials.gohtml", data); err != nil{
		fmt.Println(err)
	}
}
func Projects(w http.ResponseWriter, r *http.Request){
	data := map[string]string{"Title": "Projects"}

	if err := tpl.ExecuteTemplate(w, "projects.gohtml", data); err != nil{
		fmt.Println(err)
	}
}

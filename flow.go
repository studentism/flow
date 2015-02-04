package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
  "strconv"
 )

/*
type Module struct {
  
}
*/

type Workflow struct {
  UID int
}

type User struct {
  UID int
  Username []byte
}

type Page struct {
  Workflow Workflow
  User User
}

// dummy function for loading workflow info from flat files
func loadPage(id int) (*Page, error) {
  // TODO: add input checks
  filename := "flows/" + strconv.Itoa(id) + ".flo"
  uname, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Workflow: Workflow{UID: id}, User: User{UID: 3, Username: uname}}, nil
 }

func viewHandler(w http.ResponseWriter, r *http.Request) {
  // TODO: implement error handling for Atoi
  id, _ := strconv.Atoi(r.URL.Path[len("/view/"):])
  // TODO: implement error handling for loadPage
  p, _ := loadPage(id)
  renderTemplate(w, "view", p)
}

// renders a HTML template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  // TODO: implement error handling for ParseFiles
  t, _ := template.ParseFiles("tmpl/" + tmpl + ".html")
  t.Execute(w, p)
}

func main() {
  // TODO: add input checking(?) and error handling
  http.Handle("/static/", http.FileServer(http.Dir("/Users/jpierce/src/flow/web")))
  http.HandleFunc("/view/", viewHandler)
  http.ListenAndServe(":8888", nil)
}

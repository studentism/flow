package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
  "strconv"
  "encoding/json"
 )

/*
type Module struct {
  
}
*/

type Workflow struct {
  UID int
  Diagram map[string]interface{}
}

type User struct {
  UID int
  Username string
}

type Page struct {
  Workflow Workflow
  User User
}

// dummy function for loading workflow info from flat files
func loadPage(id int) (*Page, error) {
  // TODO: add input checks
  filename := "flows/" + strconv.Itoa(id) + ".json"
  diag, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  var dat map[string]interface{}
  if err := json.Unmarshal(diag, &dat); err != nil {
    return nil, err
  }
  return &Page{Workflow: Workflow{UID: id, Diagram: dat}, User: User{UID: 1372, Username: "jpierce"}}, nil
 }

 /*func saveDiagram(res http.ResponseWriter, req *http.Request) error {
    // TODO: implement error handling for Atoi
    id, _ := strconv.Atoi(r.URL.Path[len("/save/"):])
    filename := "flows/" + strconv.Itoa(id) + ".json"
    // TODO: implement error handling for WriteFile
    err := ioutil.WriteFile(filename, req.Body, 0600)

    data, _ := json.Marshal("{'flow':'saved'}")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
 }*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
  // TODO: implement error handling for Atoi
  id, _ := strconv.Atoi(r.URL.Path[len("/view/"):])
  // TODO: implement error handling for loadPage
  p, err := loadPage(id)
  if err != nil {
    /*http.Error(w, err.Error(), http.StatusInternalServerError)
    return*/
    panic(err)
  }
  renderTemplate(w, "view", p)
}

// renders a HTML template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, err := template.ParseFiles("tmpl/" + tmpl + ".html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  // TODO: add input checking(?) and error handling
  http.Handle("/static/", http.FileServer(http.Dir("/Users/jpierce/src/flow/web")))
  http.HandleFunc("/view/", viewHandler)
  http.ListenAndServe(":8888", nil)
}

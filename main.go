package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	
	fmt.Println("Starting server...")
	fmt.Println("The server is now runnning")
	http.HandleFunc("/enter_data", EntryHandler)
	http.HandleFunc("/get_data", OutputHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

type Country struct {
	Name string `json:"name"`
	Capital string `json:"capital"`
	Population int `json:"pop"`

}
func EntryHandler(w http.ResponseWriter, r *http.Request) {
	c :=Country{}
	json.NewDecoder(r.Body).Decode(&c)
	fmt.Fprint(w, c)
	fmt.Fprint(w, "\n")
	countryData, err := json.Marshal(c)

	file,err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		  fmt.Fprint(w,"Could not open database\n")
      return
	  }

	  defer file.Close()
	 
    _, err2 := file.WriteString(string(countryData))

	  if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		  fmt.Fprint(w, "Could not write the data to database\n")
      
	  }else{
      fmt.Fprint(w, "Operation successful! Data has been appended to the database\n")
    }
		
}


func OutputHandler(w http.ResponseWriter, r *http.Request) {

	fi, err := os.ReadFile("data.txt")
    if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "There was some error opening the database")
	}
	w.Header().Set("Content-Type","Application/Json")
	fmt.Fprint(w, string(fi))

	fmt.Fprint(w,"\n")

}
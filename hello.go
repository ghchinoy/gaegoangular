package hello

import (
	"net/http"
	"fmt"
	"appengine"
	"appengine/datastore"
	"encoding/json"
)

func init() {
	http.HandleFunc("/getWords", getWordsHandler)
	http.HandleFunc("/addWords", addWordsHandler)
}

type Word struct {
	Item string
	Meaning string
}

var words = []Word{
	{"banana", "fruit"},
	{"bee", "insect"},
}

func addWordsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	for _, oneWord := range words {
		datastore.Put(c, datastore.NewIncompleteKey(c, "Word", nil), &oneWord)
	}
}

func getWordsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Word")
	var getWords []Word
	q.GetAll(c, &getWords)

	b, _ := json.Marshal(getWords);
	fmt.Fprintf(w, "%v", string(b))
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello. GAE go app is working!")
//}
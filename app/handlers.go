package main
import (
	"net/http"
	"strconv"
	"encoding/json"
	"io"
	"io/ioutil"
)

func SpellGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	number, ok := r.URL.Query()["number"]
	if ok {
		number, err := strconv.Atoi(number[0])
		if err == nil {
			resp := ResponseToString{Status:"OK", Text:NumeralSpeller(number)}
			res, _ := json.MarshalIndent(resp, "", "  ")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(res)
		}
	}
}

func SpellPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
	}

	var data Text
	if err := json.Unmarshal(body, &data); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("STATUS CODE 406\nData harus merupakan teks"))
		panic(err)
	}

	if (data.Text == "") {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("STATUS CODE 406\nData harus merupakan teks"))
		return
	}

	spell := SpellToNumeral(data.Text)
	if spell==0 {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("STATUS CODE 406\nData harus merupakan teks"))
		return
	}

	resp := ResponseToNumber{Status:"OK", Number:spell}
	res, _ := json.MarshalIndent(resp, "", "  ")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(res)
	// fmt.Fprintf(w, "%s" , spell)
}
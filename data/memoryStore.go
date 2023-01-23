package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/samcho07/getir-test-case-MS/model"
)

// set the datas into the memory
func (d *Store) SetInMemory(w http.ResponseWriter, r *http.Request) {
	var input model.DataInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Logger.Error().Err(err)
		return
	}

	_ = d.SetKey(input.Key, input.Value)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&input)
	if err != nil {
		log.Logger.Error().Err(err)
		return
	}
}

//get the datas from memory.

func (d *Store) GetInMemory(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query()
	queryParam := key.Get("key")
	value, err := d.GetKey(queryParam)
	if err != nil {
		_, err := fmt.Fprintf(w, "%+v", err.Error())
		if err != nil {
			util.Error(err)
			return
		}
	} else {
		out := model.DataInput{Key: queryParam, Value: value}
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			util.Error(err)
			return
		}
	}
}

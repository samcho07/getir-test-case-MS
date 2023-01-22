package search

import (
	"encoding/json"
	"net/http"
	//github linkleri....
)

type mongodb struct{}

func (m *mongodb) SearchMongo(rw http.ResponseWriter, request *http.Request) {

	var Res interface{}

	if request.URL.Path == "/records" && request.Method == http.MethodPost {

		var record_quer model.Request

		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&record_quer)

		Res, _ = dataBase.MongoManager().Retrieve(record_quer)

		jsonData, _ = json.Marshal(Res)

		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonData)

		return
	}
}

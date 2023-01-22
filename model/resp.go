package model

// this is model of the request.
import "go.mongodb.org/mongo-driver/bson"

type Response struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []bson.M `json:"records"`
}

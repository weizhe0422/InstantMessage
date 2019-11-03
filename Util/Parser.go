package Util

import (
	"encoding/json"
	"log"
	"net/http"
)

type returnJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:":data,omitempty"`
}

func ParseOKResult (w http.ResponseWriter, data interface{}){
	parseAsReturnObject(w, 0, "", data)
}

func ParseFailResult (w http.ResponseWriter, msg string){
	parseAsReturnObject(w, -1, msg, nil)
}

func parseAsReturnObject(writer http.ResponseWriter, statusCode int, msg string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	returnMsg := &returnJson{
		Code: statusCode,
		Msg:  msg,
		Data: data,
	}
	jsonMsg, err := json.Marshal(returnMsg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	writer.Write(jsonMsg)
}

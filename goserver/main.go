package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Success bool `json:"success"`
	Data map[string]string `json:"data"`
}

type RequestBody struct {
	SleepTime int `json:"sleepTime"`
}

func handError(writer *http.ResponseWriter, err error) {
	if err != nil {
		fmt.Errorf("failed to marshal json with error - %+v", err)
		(*writer).WriteHeader(500)
		(*writer).Write([]byte(fmt.Sprintf("failed json marshal with error - %+v", err)))
		return
	}
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		resp := Response{
			Success: true,
			Data:    nil,
		}
		switch request.Method {
		case http.MethodGet:
			data := map[string]string{"message": "go server working !!!"}
			resp.Data = data
			buf, err := json.Marshal(resp)
			handError(&writer, err)
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(buf)
			return
		case http.MethodPost:
			iof, err := ioutil.ReadAll(request.Body)
			defer request.Body.Close()
			handError(&writer, err)
			req := RequestBody{}
			err = json.Unmarshal(iof, &req)
			handError(&writer, err)
			data := map[string]string{"message": fmt.Sprintf("responding after %d milliseconds.", req.SleepTime)}
			resp.Data = data
			buf, err := json.Marshal(resp)
			handError(&writer, err)
			time.Sleep(time.Duration(req.SleepTime) * time.Millisecond)
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(buf)
			return
		}
	})
	port := "5000"
	if p := os.Getenv("PORT"); p != ""{
		port = p
	}
	fmt.Println("go server is starting on port: ", port)
	http.ListenAndServe(":"+port, nil)
}
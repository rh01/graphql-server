package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Params struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

// https://stackoverflow.com/questions/46948050/how-to-read-request-body-twice-in-golang-middleware

func AuthorizationMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var params Params
		// convert request body to bytes then unmarshal the json
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(bodyBytes, &params); err != nil {
			log.Fatal(err)
		}
		// construct a new ReadCloser
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// if createOrder, let it pass
		if strings.Contains(params.OperationName, "CreateOrder") {
			fmt.Printf("got %s mutation, OK to pass", params.OperationName)
			return
		}
		// verify request is a mutation
		if strings.Contains(params.Query, "mutation") {
			fmt.Println("got mutation request")
			// check for valid JWT in authorization header
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer")
			if len(splitToken) != 2 {
				log.Fatal("Bearer token not in proper format")
			}
			reqToken = strings.TrimSpace(splitToken[1])
			fmt.Println("token is %s", reqToken)
			if reqToken != "" {
				h.ServeHTTP(w, r)
				return
			}
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

package api_server

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wwqgtxx/mfp/find_pair"
	"github.com/wwqgtxx/mfp/mfp_web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type event struct {
	ID     string `json:"id"`
	Word   string `json:"word"`
	Result string `json:"result"`
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome visit mfp api api_server!")
	if err != nil {
		log.Println(err)
	}
}

func apiHandle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("%+v\n", r)
			_, err := fmt.Fprintln(w, r)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}
	}()
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		panic(err)
	}

	newEvent.Result, err = find_pair.FindPairFromString(newEvent.Word)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(newEvent)
	if err != nil {
		panic(err)
	}
}

func ApiServer() {
	var listen string
	commandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	commandLine.StringVar(&listen, "listen", "0.0.0.0:8000", "监听地址")
	commandLine.StringVar(&find_pair.FilePath, "res", find_pair.FilePath, "数据文件地址")

	err := commandLine.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}

	find_pair.Init()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", apiHandle)
	router.PathPrefix("/").Handler(http.FileServer(http.FS(mfp_web.Dist)))
	log.Printf("启动API服务器，监听地址：%s\n", listen)
	log.Fatal(http.ListenAndServe(listen, router))
}

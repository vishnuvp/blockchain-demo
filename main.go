package main

import (
	"blockchain/pkg/block"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
type Message struct{
	Data int
}
func getBlockChain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(block.BlockChain, "", "  ")
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
	return
}
func sendResponse(data interface{}, w http.ResponseWriter, statusCode int) {
	response, err := json.MarshalIndent(data, "", " ")
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)


}

func postBlock(w http.ResponseWriter, r*http.Request) {
	log.Println("Trying to create new block")
	var m Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err!=nil {
		sendResponse(r.Body, w, http.StatusBadRequest)
	}
	defer r.Body.Close()

	newBlock := block.CreateBlock(block.BlockChain[len(block.BlockChain)-1], m.Data)

	if block.IsValid(newBlock, block.BlockChain[len(block.BlockChain) -1]) {
		newBlockChain := append(block.BlockChain, newBlock)
		block.ReplaceChain(newBlockChain)
		spew.Dump(block.BlockChain)
	}

	sendResponse(newBlock, w, http.StatusCreated)


}

func makeMuxRouter() http.Handler{
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", getBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", postBlock).Methods("POST")
	return muxRouter

}
func run() error {
	mux := makeMuxRouter()

	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", httpAddr)

	s := &http.Server{
		Addr:              ":"+httpAddr,
		Handler:           mux,
		ReadTimeout:       10*time.Second,
		WriteTimeout:      10*time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	if err := s.ListenAndServe(); err!= nil {
		return err
	}
	return nil
}

func main() {

	if err := godotenv.Load(); err!=nil {
		log.Fatal(err)
	}

	go func() {
		genesisBlock := block.Block{
			Index:     0,
			Timestamp: time.Now().String(),
			Data:      0,
			Hash:      "",
			PrevHash:  "",
		}
		block.BlockChain = append(block.BlockChain, genesisBlock)
		spew.Dump(genesisBlock)

	}()

	log.Fatal(run())
}

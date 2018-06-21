package main

import (
	"net/http"
	"fmt"
	"log"
	"io"
	"strconv"
	"os"
	"math/big"

	"github.com/gorilla/mux"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

var auth *bind.TransactOpts
var backend *backends.SimulatedBackend

func main() {
	connectToEthClient()

	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	//router.Host("www.example.com")
	router.Methods("GET", "POST")

	router.HandleFunc("/", Index)
	router.HandleFunc("/createContract/{}", CreateContract)
	router.HandleFunc("/health", HealthCheckHandler)
	router.HandleFunc("/hello/{name}", Hello)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Println(r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %v\n", vars["name"])
}

func connectToEthClient() {
	key, _ := crypto.GenerateKey()
	auth = bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(100000000)}
	backend = backends.NewSimulatedBackend(alloc)
}

func CreateContract(w http.ResponseWriter, r *http.Request) {
	vals := make(map[string]string)
	vals["name"], vals["id"], vals["birthDate"] = r.FormValue("name"), r.FormValue("id"), r.FormValue("birthDate")

	fmt.Fprintln(w, "Connecting to Ethereum client!")

	DeployContract(auth, backend, vals)
}

func DeployContract(auth *bind.TransactOpts, backend *backends.SimulatedBackend, vals map[string]string) (addr common.Address, contract *VisaApplicationContract) {
	var name [32]byte
	copy(name[:], vals["name"])

	idInt, err1 := strconv.ParseInt(vals["id"], 10, 64)
	birthDateInt, err2 := strconv.ParseInt(vals["birthDate"], 10, 64)

	if err1 != nil && err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		os.Exit(2)
	}

	id := big.NewInt(idInt)
	birthDate := big.NewInt(birthDateInt)

	addr, _, contract, err := DeployVisaApplicationContract(auth, backend, name, id, birthDate)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}

	return
}
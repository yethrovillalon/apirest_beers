package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var _status = StatusC{}
var _ret = Retorno{}
var _retError = RetornoError{}
var _apiret = apiCashResponse{}
var _retBox = RetornoBox{}

var collection = getSession().DB("api_beer").C("beers")

// getSession ... < Some lines that describe your function>
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

// Index ... < Some lines that describe your function>
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta API esta diseñada para ser una prueba para los nuevos candidatos al equipo.")
}

// Beers ... < Some lines that describe your function>
func beers(w http.ResponseWriter, r *http.Request) {

	/*b, err := json.MarshalIndent(_beers, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)*/

	var results []Beer

	err := collection.Find(nil).Sort("-id").All(&results)

	var count = len(results)
	var code int

	if err != nil {
		log.Fatal(err)
		code = 500
		_status = StatusC{
			500, "Error contactarse con el adminsitrador"}
		_retError = RetornoError{
			_status}
	} else if count == 0 {
		code = 400
		_status = StatusC{
			400, "Sin reusltados"}
		_retError = RetornoError{
			_status}
	} else {
		fmt.Println("Resultados: ", results)
		_status = StatusC{
			200, "Operacion exitosa"}
		_ret = Retorno{
			_status,
			results}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(_ret)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(_retError)

}

// beersIng ... < Some lines that describe your function>
func beersIng(w http.ResponseWriter, r *http.Request) {

	/*b, err := json.MarshalIndent(_beers, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)*/
	decoder := json.NewDecoder(r.Body)
	var _beerdata Beer
	var code int = 0
	err4 := decoder.Decode(&_beerdata)
	if err4 != nil {
		log.Fatal(err4)
		code = 500
		_status = StatusC{
			code, "Request invalida"}
		_retError = RetornoError{
			_status}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(_retError)
		return
	}

	var results []BeerPrice
	err3 := collection.Find(bson.M{"id": _beerdata.ID}).All(&results)

	if err3 != nil || len(results) > 0 {
		_status = StatusC{
			409, "El ID de la cerveza ya existe"}
		_retError = RetornoError{
			_status}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(_retError)
		return
	}

	defer r.Body.Close()

	err3 = collection.Insert(_beerdata)
	if err3 != nil {
		code = 400
		_status = StatusC{
			code, "Request invalida"}
		_retError = RetornoError{
			_status}
	}

	if code == 0 {
		code = 200
		_status = StatusC{
			201, "Cerveza creada"}
		_retError = RetornoError{
			_status}
	}

	//_beers = append(_beers, _beerdata)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(_retError)

}

// beersBoxprice ... < Some lines that describe your function>
func beersBoxprice(w http.ResponseWriter, r *http.Request) {

	/*b, err := json.MarshalIndent(_beers, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)*/
	params := mux.Vars(r)

	_beerid := params["beerID"]

	IDBeer, err1 := strconv.Atoi(_beerid)
	if err1 != nil {
		log.Fatal(err1)
	}

	var results []BeerPrice
	err := collection.Find(bson.M{"id": IDBeer}).All(&results)

	if err != nil || len(results) == 0 {
		_status = StatusC{
			404, "El Id de la cerveza no existe"}
		_retError = RetornoError{
			_status}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(_retError)
		return
	}

	moneda := string(results[0].Currency)

	APIURL := "http://www.apilayer.net/api/live?access_key=381020e1688034c5ce48ff5150775baf&currencies=" + moneda + "&format=1"
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	respAPI := _apiret
	json.Unmarshal(body, &respAPI)

	var cantidad float64 = 6
	total_precio := (cantidad * float64(respAPI.Quotes["USD"+moneda]))

	results[0].Quantity = cantidad
	results[0].PriceBox = total_precio

	_status = StatusC{
		200, "Operacion exitosa"}
	_retBox = RetornoBox{
		_status,
		results}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(_retBox)

	/*var data map[string]interface{}
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		panic(err2)
	}
	//fmt.Println(data)
	//fmt.Printf("%#v", data)

	fetchValue(data)*/
}

func fetchValue(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Printf("%v is an interface \n ", value)
	case bool:
		fmt.Printf("%v is bool \n ", value)
	case float64:
		fmt.Printf("%v is float64 \n ", value)
	case []interface{}:
		fmt.Printf("%v is a slice of interface \n ", value)
		for _, v := range value.([]interface{}) { // use type assertion to loop over []interface{}
			fetchValue(v)
		}
	case map[string]interface{}:
		fmt.Printf("%v is a map \n ", value)
		for _, v := range value.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			fetchValue(v)
		}
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
}

// BeerDetail ... < Some lines that describe your function>
func BeerDetail(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	_beerid := params["beerID"]

	/*if !bson.IsObjectIdHex(_beerid) {
		_status = StatusC{
			404, "El Id de la cerveza no existe"}
		_ret = Retorno{
			_status,
			nil}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(_ret)
		return
	}

	//oid := bson.ObjectIdHex(_beerid)*/
	var results []Beer

	IDBeer, err1 := strconv.Atoi(_beerid)
	if err1 != nil {
		log.Fatal(err1)
	}

	err := collection.Find(bson.M{"id": IDBeer}).All(&results)

	if err != nil || len(results) == 0 {
		_status = StatusC{
			404, "El Id de la cerveza no existe"}
		_retError = RetornoError{
			_status}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(_retError)
		return
	}

	_status = StatusC{
		200, "Operacion exitosa"}
	_ret = Retorno{
		_status,
		results}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(_ret)

}

package main

import (
	"fmt"
	"net/http" 
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

type AutoGenerated struct {
	Rajaongkir Rajaongkir `json:"rajaongkir"`
}
type Query struct {
	ID string `json:"id"`
}
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
type Results struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}
type Rajaongkir struct {
	Query   Query   `json:"query"`
	Status  Status  `json:"status"`
	Results Results `json:"results"`
}

func getAPI() (string,string,string,string){
	api_key := "8b06b6a059bd1e215464fb7c0d57f74c"

	url := "https://api.rajaongkir.com/starter/city?id=25"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", api_key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	var ag AutoGenerated
	json.Unmarshal(body, &ag)
	var provinceid = ag.Rajaongkir.Results.ProvinceID
	var province = ag.Rajaongkir.Results.Province
	var cityname = ag.Rajaongkir.Results.CityName
	var cityid = ag.Rajaongkir.Results.CityID
	return cityid, cityname,provinceid,province
}

type Ongkir struct {
	ProvinceID string    
	Province string    
	CityID string    
	CityName string       
}

func getOngkir(w http.ResponseWriter, r *http.Request){
	var ongkir Ongkir // variable untuk memetakan data product yang terbagi menjadi 3 field
	
	cityid,cityname,provinceid,provincename := getAPI()

	fmt.Println(cityid)
	fmt.Println(cityname)
	fmt.Println(provinceid)
	fmt.Println(provincename)

	ongkir.CityID = cityid
	ongkir.CityName = cityname
	ongkir.ProvinceID = provinceid
	ongkir.Province = provincename
	json.NewEncoder(w).Encode(ongkir)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getOngkir",getOngkir).Methods("GET") // menjalurkan URL untuk dapat mengkases data JSON API product ke /getproducts
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":4321",router))
}
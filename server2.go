package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Products struct {
	Sku          string
	Product_name string 
	Stocks       int
}

func getProduct(w http.ResponseWriter, r *http.Request){
    var products Products // variable untuk memetakan data product yang terbagi menjadi 3 field
    
    var arr_products []Products
    db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/go_coba")
    defer db.Close()
                    
    if(err != nil) {
        log.Fatal(err)
    }
    rows, err := db.Query("Select sku,product_name,stocks from products ORDER BY sku DESC")
    if err!= nil {
        log.Print(err)
    }
    count:=0
    for rows.Next(){
        if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
            log.Fatal(err.Error())
        
        }else{
            arr_products = append(arr_products, products)
            fmt.Println(arr_products[count])
        }
        count++
    }

    json.NewEncoder(w).Encode(arr_products)

}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/getProduct",getProduct).Methods("GET") // menjalurkan URL untuk dapat mengkases data JSON API product ke /getproducts
    http.Handle("/",router)
    log.Fatal(http.ListenAndServe(":5432",router))
}
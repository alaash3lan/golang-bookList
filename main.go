package main

import (

	// "encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	"log"
    "database/sql"
   _"github.com/go-sql-driver/mysql"

)

type Book struct {
    
     ID     int   `json:id`
     Title  string   `json:title`
     Author string   `json:author`
     Year   string   `json:year`
   
}




func db () *sql.DB{


	db, err := sql.Open("mysql", "root:01060902788@tcp(127.0.0.1:3306)/gotest")
	// db, err := sql.Open("mysql","root:01060902788@tcp(127.0.0.1:3306/gotest)")

	if err != nil {
	
       panic(err.Error())

	} 

	// defer db.Close()
	return db

}






func main(){

  
	


	


    

    

	fmt.Println("system work")

   // router :=mux.NewRouter()
   // router.HandleFunc("/books",getBooks).Methods("GET")
   // router.HandleFunc("/books/{id}",getBook).Methods("GET")
   // router.HandleFunc("/books",addBook).Methods("POST")
   // router.HandleFunc("/books",updateBook).Methods("PUT")
   // router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")
   
   // log.Fatal(http.ListenAndServe(":8000", router))


}

func getBooks(w http.ResponseWriter, r *http.Request){

	db := db()

	defer db.Close()
	selectAll(db)

log.Println("Gets all Books")

}

func getBook(w http.ResponseWriter, r *http.Request){
    db := db() 
    defer db.Close()
 
    selectOne(db,5)

    log.Println("Gets one Books")

}

func addBook(w http.ResponseWriter, r *http.Request){

	 db := db()
	 defer db.Close()

	 update(db,4)


     


     log.Println("add  Books")

}

func updateBook(w http.ResponseWriter, r *http.Request){
	 db := db()
	 defer db.Close()

	 update(db,4)

log.Println("update  Book")

}
func removeBook(w http.ResponseWriter, r *http.Request){

    db := db()

	defer db.Close()

    delete(db,3)

log.Println("remove  Book")

}

func selectAll(db *sql.DB){
    var book Book
    sqlStatement := "SELECT * FROM books"
    rows, err := db.Query(sqlStatement)
    if err != nil {
       panic(err)
    }
    defer rows.Close()
    var bookSlice []Book
    for rows.Next(){
        rows.Scan(&book.ID, &book.Title, &book.Author, 
                  &book.Year)
         bookSlice = append(bookSlice, book)
    }
    fmt.Println(bookSlice)
}

func selectOne(db *sql.DB,id int){
    var book Book
    sqlStatement := `SELECT * FROM books WHERE id=?`
    rows, err := db.Query(sqlStatement,id)
    if err != nil {
       panic(err)
    }
    defer rows.Close()
    var bookSlice []Book
    for rows.Next(){
        rows.Scan(&book.ID, &book.Title, &book.Author, 
                  &book.Year)
         bookSlice = append(bookSlice, book)
    }
    fmt.Println(bookSlice)
}

func create(db *sql.DB) {

	sqlStatement :="INSERT INTO books (title, author, year) VALUES ('the trial', 'kafka', '1918')"

     insert, err := db.Query(sqlStatement);

     if err != nil{

           
         panic(err.Error())  
        
     }

     defer insert.Close()


}

func delete(db *sql.DB,id int) {

	sqlStatement :=`DELETE FROM books WHERE id=?`

     result, err := db.Query(sqlStatement,id);

     if err != nil{
           
         panic(err.Error())  
        
     }

     defer result.Close()


}

func update(db *sql.DB,id int) {
    

	sqlStatement :=`UPDATE books SET title='Juan',author='Juan', year='1999' WHERE id=?`

    result, err := db.Query(sqlStatement,id);

    fmt.Println(err)

    if err != nil{
    	if err == sql.ErrNoRows {
        fmt.Println("Zero rows found")
    } else {
        panic(err)
    }


    }

     defer result.Close()


}

package main

import (
    "encoding/csv"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/gorilla/mux"
)

// Record represents a single row in the dataset
type Record struct {
    ID    int
    Name  string
    Age   int
    Email string
}

var dataset []Record

func main() {
    // Read dataset file into memory
    file, err := os.Open("Book1.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = 4

    rows, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    for i, row := range rows {
        if i == 0 {
            // Skip header row
            continue
        }

        id, _ := strconv.Atoi(row[0])
        age, _ := strconv.Atoi(row[2])

        record := Record{
            ID:    id,
            Name:  row[1],
            Age:   age,
            Email: row[3],
        }

        dataset = append(dataset, record)
    }

    // Create new Gorilla Mux router
    r := mux.NewRouter()

    // Define API endpoints
    r.HandleFunc("/dataset", getDataset).Methods("GET")
    r.HandleFunc("/dataset", createRecord).Methods("POST")
    r.HandleFunc("/dataset/{id}", updateRecord).Methods("PUT")
    r.HandleFunc("/dataset/{id}", deleteRecord).Methods("DELETE")
    r.HandleFunc("/dataset/filter", filterDataset).Methods("GET")

    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", r))
}

// getDataset retrieves all records in the dataset
func getDataset(w http.ResponseWriter, r *http.Request) {
    writeCSV(w, dataset)
}

// createRecord creates a new record in the dataset with the provided column values
func createRecord(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    id, _ := strconv.Atoi(r.FormValue("id"))
    name := r.FormValue("name")
    age, _ := strconv.Atoi(r.FormValue("age"))
    email := r.FormValue("email")

    // Create new record
    record := Record{
        ID:    id,
        Name:  name,
        Age:   age,
        Email: email,
    }

    // Add record to dataset
    dataset = append(dataset, record)

}

// updateRecord updates an existing record in the dataset with the provided column values
func updateRecord(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    name := r.FormValue("name")
    age, _ := strconv.Atoi(r.FormValue("age"))
    email := r.FormValue("email")

    // Find record by ID
    var record *Record
    for i, r := range dataset {
        if r.ID == id {
            record = &dataset[i]
            break
        }
    }

    // Update record
}
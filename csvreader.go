package main

import (
   "encoding/csv"
   "fmt"
   "os"
)

func main() {

   path := "/tmp/test.csv"
   file, err := os.Open(path)

   if err != nil {
      panic(err)
   }

   defer file.Close()

   reader := csv.NewReader(file)
   rows, err := reader.ReadAll()
   if err != nil {
      panic(err)
   }

   for _, row := range rows {
      fmt.Println(row)
   }
}

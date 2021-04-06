package main 
// module takes file with coin hash and write it to data base and in system file

import (
	"fmt"
  "os"
  "io"
)

func main(){
   var text_of_hash  string
   
   file , err := os.Create("hash.txt")
   //checking
   if err != nil {
     fmt.Println("Unable to get information",err)
     os.Exit(1)
   }
   defer file.Close()
   text_of_hash = "hash"
   file.WriteString(text_of_hash) //to write data
   fmt.Println("Done.")
   data := make([]byte, 64) // type of data is byte to read file
   for {
    n, err := file.Read(data)
    //if end of the file go out , break loop
    if err == io.EOF{
      break
    }
    
    fmt.Print(data[:n]) //print data array before n
   }  
}
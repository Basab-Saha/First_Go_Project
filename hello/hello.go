package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

   names:= []string{"Basab","Anish","Bhumika"}

   messages,err :=greetings.Hellos(names)

   if err!=nil{
         log.Fatal(err)
   }
   for _,message:=range messages{
     //key
        fmt.Println(message)
     //value

   }
}
package greetings

import (
    "errors"
    "fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string,error){
    //first make an empty map
    messages:=make(map[string]string)
    //loop through the names slice
    for _,name := range names{
        message,err:=Hello(name)
        //if there is  error
        if err!=nil{
            return nil,err
        }
        //if there is no error make the key value pair
        messages[name]=message

    }
    return messages,nil
}

func randomFormat() string {
	//we will decleare a slice (vector in c++) of strings
	formats := []string{
		"Hi , %v welcome!",
		"Great to see you , %v",
		"Hail , %v! Well met!",
	}

	//we will return a random string from the slice
	return formats[rand.Intn(len(formats))]
}
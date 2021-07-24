package commands

import (
    "net/http"
    "log"
    "encoding/json"
)



func Joke() string {
    joke, err := getJoke()
    if(err != nil){
        log.Printf(joke)
        log.Panic("Error al traer /joke")
    }
    return joke
}

func getJoke() (string, error){
    resp, err := http.Get("http://api.icndb.com/jokes/random")
    c := &joke{}
    if err != nil {
        return "", err
    }
    err = json.NewDecoder(resp.Body).Decode(c)
    log.Printf(c.Value.Joke)
    return c.Value.Joke, err
}

type joke struct {
    Value struct {
        Joke string `json:"joke"`
    } `json:"value"`
}

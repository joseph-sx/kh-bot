package commands

import (
    "net/http"
    "log"
)



func jokeFetcher() (string, error) {
    resp, err := http.Get("http://api.icndb.com/jokes/random")
    c := &joke{}
    if err != nil {
        return "", err
    }
    err = json.NewDecoder(resp.Body).Decode(c)
    return c.Value.Joke, err
}

type joke struct {
    Value struct {
        Joke string `json:"joke"`
    } `json:"value"`
}

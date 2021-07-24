package commands

import (
	"fmt"
    "net/http"
	"net/url"
    "log"
	
	"encoding/json"
    
)



func Pokemon(Pokemon string) string{
	pokemon := url.QueryEscape(Pokemon)
	poke_api := fmt.Sprintf("https://app.pokemon-api.xyz/pokemon/%s",pokemon)
	req, err := http.NewRequest("GET", poke_api, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	var poke_dat PokeData

	if err := json.NewDecoder(resp.Body).Decode(&poke_dat); err != nil {
		log.Println(err)
	}
	return poke_dat.Hires
	
}

type PokeData struct {
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`
	Type []string `json:"type"`
	Base struct {
		Hp        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
	Species     string `json:"species"`
	Description string `json:"description"`
	Evolution   struct {
		Prev []string `json:"prev"`
	} `json:"evolution"`
	Profile struct {
		Height  string     `json:"height"`
		Weight  string     `json:"weight"`
		Egg     []string   `json:"egg"`
		Ability [][]string `json:"ability"`
		Gender  string     `json:"gender"`
	} `json:"profile"`
	Image struct {
	} `json:"image"`
	Sprite    string `json:"sprite"`
	Thumbnail string `json:"thumbnail"`
	Hires     string `json:"hires"`
}

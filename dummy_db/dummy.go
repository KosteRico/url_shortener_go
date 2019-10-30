package dummy_db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

type Link struct {
	ID        int    `json:"ID"`
	RealURL   string `json:"RealURL"`
	ShortName string `json:"ShortName"`
}

type AllLinks []Link

var Links AllLinks

func (all AllLinks) FindByName(name string) (*Link, error) {
	for _, item := range all {
		if name == item.ShortName {
			return &item, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Link with name \"%s\" not found", name))
}

func (all AllLinks) FindById(id int) (*Link, error) {

	if id <= len(all) && id >= 1 {
		link := all[id-1]
		return &link, nil
	}

	return nil, errors.New(fmt.Sprintf("Link with name \"%d\" not found", id))
}

func InitDB() {

	var temp AllLinks

	jsonFile, err := ioutil.ReadFile("dummy_db/db.json")

	if err != nil {
		os.Create("dummy_db/db.json")
		jsonFile, err = ioutil.ReadFile("dummy_db/db.json")
	}

	json.Unmarshal(jsonFile, &temp)

	if len(temp) == 0 {
		Links = AllLinks{}
	} else {
		Links = temp
	}

	log.Println(fmt.Sprintf("Links array: %v", Links))

}

func isUrlValid(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}
	return true
}

func NewLink(body []byte) error {

	var link = Link{}

	var id int

	if len(Links) > 0 {
		id = Links[len(Links)-1].ID + 1
	} else {
		id = 1
	}
	link.ID = id

	err := json.Unmarshal(body, &link)

	if err != nil {
		return err
	}

	if isUrlValid(link.RealURL) == false {
		return errors.New("URL is invalid!!")
	}

	Links = append(Links, link)

	log.Println(fmt.Sprintf("Link with id \"%d\" was created!", id))

	WriteToDB()

	return nil
}

func WriteToDB() {
	file, err := json.MarshalIndent(Links, "", " ")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("dummy_db/db.json", file, 0644)

	if err != nil {
		panic(err)
	}
}

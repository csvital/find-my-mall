package consumer

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	. "github.com/works-forces/find-my-mall/config"
	. "github.com/works-forces/find-my-mall/dao"
	models "github.com/works-forces/find-my-mall/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = ShoppingMallsDAO{}

func consume() {
	config.Read("../config.toml")

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mall := models.ShoppingMall{ID: bson.NewObjectId()}
		err := json.Unmarshal([]byte(scanner.Text()), &mall)
		if err != nil {
			log.Fatal(err)
		}
		if err := dao.Insert(mall); err != nil {
			log.Fatalln(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {

}

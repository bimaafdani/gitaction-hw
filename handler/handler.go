package handler

import (
	"gitaction_hw/model"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	response, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]model.User{}
	json.Unmarshal(data, &user)
	var result string
	SearchUsername := ctx.Params("username")

	for _, v := range user {
		if v.Username == SearchUsername {
			result = "followers:" + strconv.Itoa(v.Follower)
			break
		}
	}
	return ctx.SendString(result)
	//return c.SendString("All Data Show")
}

func GetByUserId(ctx *fiber.Ctx) error {
	response, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]model.User{}
	json.Unmarshal(data, &user)
	userid := ctx.Params("userid")
	result := user[userid]

	return ctx.JSON(result)
	//return c.SendString("By Id Data Show")
}

package function

import (
	"log"
	"../utils"
	"../api"
	"fmt"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

type BaseTicketBean struct {
	validateMessagesShowId string
	status                 bool
	httpstatus             int
	data                   []interface{}
}

//noinspection GoUnusedExportedFunction
func LeftTicket(train_date, from_station, to_station string) []interface{}{

	client := utils.NewClient()

	url := fmt.Sprintf(api.LeftTicketURL+"leftTicketDTO.train_date=%s"+"&leftTicketDTO.from_station=%s"+"&leftTicketDTO.to_station=%s"+"&purpose_codes=ADULT", train_date, from_station, to_station)

	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
		//return "null"
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		//return "null"
	}

	js, _ := simplejson.NewJson(body)
	dataArr, _ := js.Get("data").Array()

/*	for _, data := range dataArr {
		//var stb = &BaseTicketBean{}

		//str, _ := data.(string)
		//json.Unmarshal([]byte(str), &stb)
		fmt.Println(data)
	}*/
	return dataArr
}

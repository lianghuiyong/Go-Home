package api

//noinspection GoUnusedConst
const (

	Host = "https://kyfw.12306.cn/"
	//获取车站名连接
	StationNameURL = Host+"otn/resources/js/framework/station_name.js"

	// /otn/leftTicket/query?leftTicketDTO.train_date=2017-04-25&leftTicketDTO.from_station=SZQ&leftTicketDTO.to_station=VAG&purpose_codes=ADULT
	//车站查找
	LeftTicketURL =Host+"otn/leftTicket/query?"
)

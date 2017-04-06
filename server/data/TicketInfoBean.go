package data

type BaseTicketBean struct {
	queryLeftNewDTO TicketInfoBean
	secretStr       string
	buttonTextInfo  string
}

type TicketInfoBean struct {
	train_no                 string
	station_train_code       string
	start_station_telecode   string
	start_station_name       string
	end_station_telecode     string
	end_station_name         string
	from_station_telecode    string
	from_station_name        string
	to_station_telecode      string
	to_station_name          string
	start_time               string
	arrive_time              string
	day_difference           string
	train_class_name         string
	lishi                    string
	canWebBuy                string
	lishiValue               string
	yp_info                  string
	control_train_day        string
	start_train_date         string
	seat_feature             string
	yp_ex                    string
	train_seat_feature       string
	train_type_code          string
	start_province_code      string
	start_city_code          string
	end_province_code        string
	end_city_code            string
	seat_types               string
	location_code            string
	from_station_no          string
	to_station_no            string
	control_day              int
	sale_time                string
	is_support_card          string
	controlled_train_flag    string
	controlled_train_message string
	gg_num                   string
	gr_num                   string
	qt_num                   string
	rw_num                   string
	rz_num                   string
	tz_num                   string
	wz_num                   string
	yb_num                   string
	yw_num                   string
	yz_num                   string
	ze_num                   string
	zy_num                   string
	swz_num                  string
}

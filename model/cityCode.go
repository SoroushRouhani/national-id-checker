package model

var cityCode map[string]string = map[string]string{  
	"008": "Tehran-Center","007" : "Tehran-Center","006" : "Tehran-Center","005" : "Tehran-Center","004" : "Tehran-Center",
	"003" : "Tehran-Center","002" : "Tehran-Center","001" : "Tehran-Center","011" : "Tehran-south","020" : "Tehran-north",
	"258": "Rasht","259": "Rasht","407": "Khoramabad",

}

func getCityName(nid string) string { 

	val, ok := cityCode[nid] 
	if ok {
		return val
	}
	return "Sorry, City is not defined!"
}

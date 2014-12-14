package data

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/influxdb/influxdb/client"
	"gopkg.in/mgo.v2/bson"
)

var hashKey = []byte("0ea20ffa64fb127136be5bb9fb0814289267f6d162c81f06af237a09e375d182bada41b9ba14eb4d34f73c9c8b792fc0")
var s = securecookie.New(hashKey, nil)

type Data struct {
	IncluxDB *client.Client
}

func (d *Data) CreateUserId(w http.ResponseWriter) bson.ObjectId {
	id := bson.NewObjectId()

	if encoded, err := s.Encode("onepixel", map[string]string{"id": id.Hex()}); err == nil {
		cookie := &http.Cookie{
			Name:  "onepixel",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	} else {
		log.Println("cannot encrypt uuid")
	}
	return id
}

func (d *Data) GetUserIdInCookie(r *http.Request) bson.ObjectId {
	if cookie, err := r.Cookie("onepixel"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("onepixel", cookie.Value, &value); err == nil {
			if bson.IsObjectIdHex(value["id"]) {
				return bson.ObjectIdHex(value["id"])
			}
		}
	}
	return ""
}

func (d *Data) GetAndSetId(w http.ResponseWriter, request *http.Request) bson.ObjectId {
	if id := d.GetUserIdInCookie(request); id.Valid() {
		return id
	} else {
		return d.CreateUserId(w)
	}
}

func (d *Data) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	ip, _, _ := net.SplitHostPort(request.RemoteAddr)

	userAgent := request.UserAgent()
	useId := d.GetAndSetId(w, request)
	url := vars["url"]
	serieName := vars["serie_name"]

	d.AddDataPoint(serieName, url, userAgent, ip, useId)
}

func (d *Data) CreateSeries(serieName, url, userAgent, ip string, userId bson.ObjectId) []*client.Series {
	series := &client.Series{
		Name:    serieName,
		Columns: []string{"url", "user_id", "user_agent", "ip"},
		Points: [][]interface{}{
			{url, userId, userAgent, ip},
		},
	}
	return []*client.Series{series}
}

func (d *Data) AddDataPoint(serieName, url, userAgent, ip string, userId bson.ObjectId) {
	fmt.Println("serieName  :", serieName)
	fmt.Println("url        :", url)
	fmt.Println("id         :", userId.Hex())
	fmt.Println("userAgent  :", userAgent)
	fmt.Println("")

	series := d.CreateSeries(serieName, url, userAgent, ip, userId)
	if err := d.IncluxDB.WriteSeries(series); err != nil {
		log.Println(err)
	}

}

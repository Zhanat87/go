package pack

import (
	"encoding/xml"
	"fmt"
	"net/http"
)
// weather api
func PrintWeather() {
	fmt.Println("start")
	apiUrl := "http://api.openweathermap.org/data/2.5/weather?appid=8b745f60c241e69915fbe7a9b4ab96b9&units=imperial&q=Almaty&mode=xml"
	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err)
		return
	}
	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)

	var weather Weather

	// https://golang.org/src/encoding/xml/example_test.go
	err = xml.Unmarshal(data, &weather)

	if err != nil {
		fmt.Println("error 2")
		fmt.Println(err)
		return
	}

	fmt.Println(weather)
	fmt.Println("celcius: ")
	fmt.Println(weather.Temperature.getTempInCelcius())
	fmt.Println("end")
}

type City struct {
	ID      int    `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Coord   Coord  `xml:"coord"`
	Country string `xml:"country"`
	Sun     Sun    `xml:"sun"`
}

type Coord struct {
	Lon float32 `xml:"lon,attr"`
	Lat float32 `xml:"lat,attr"`
}

type Sun struct {
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}

type Temperature struct {
	Value float32 `xml:"value,attr"`
	Min   float32 `xml:"min,attr"`
	Max   float32 `xml:"max,attr"`
	Unit  string  `xml:"unit,attr"`
}

type Weather struct {
	XMLName     xml.Name    `xml:"current"`
	City        City        `xml:"city"`
	Temperature Temperature `xml:"temperature"`

}

func (t Temperature) getTempInCelcius() float32 {
	return (t.Value - 32) * 5 / 9
}

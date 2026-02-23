package geo_test

import (
	"http/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - подготовка, excpected результат, данные для функции
	city := "London"
	excpected := geo.GeoData{
		City: "London",
	}

	//Act - выполнение функции
	got, err := geo.GetMyLocation(city)

	
	//Assert - првоерка резульатов с excpected
	if err != nil{
		t.Error("Ошибка получения города")
	}
	if got.City != excpected.City{
		t.Errorf("Ожидалось %v, получено %v", excpected, got)
	}
}

// func TestGetMyLocationNoCity(t *testing.T) { //тест как пример, потому что сервис по чеку города не работает
// 	city := "Qweasd"
// 	_, err := geo.GetMyLocation(city)
// 	if err != geo.ErrorNoCity{
// 		t.Errorf("Ожидалось %v, получено %v", geo.ErrorNoCity, err)
// 	}
// }
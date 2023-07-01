package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	Currency "main.go/Currency/Model"
	appCfg "main.go/cfg"
	postgres "main.go/pkg/client"
	"net/http"
	"reflect"
	"strconv"
)

func main() {
	dbCfg := appCfg.GetDbConfig()
	conn, _ := postgres.NewClient()
	defer conn.Close()

	// Создание GET-запроса
	date := "27.07.2018"
	resp, err := http.Get("https://www.cnb.cz/en/financial_markets/foreign_exchange_market/exchange_rate_fixing/daily.txt?date=" + date)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Чтение тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Получаем тип содержимого из заголовков ответа
	contentType := resp.Header.Get("Content-Type")
	fmt.Println("Тип содержимого:", contentType)
	fmt.Println(reflect.TypeOf(body))

	var currencys []Currency.Currency

	lines := bytes.Split(body, []byte("\n"))
	for index, line := range lines {
		if index == 0 || index == 1 {
			continue
		}

		fields := bytes.Split(line, []byte("|"))
		if len(fields) < 5 {
			break
		}

		amount, err := strconv.Atoi(string(fields[2]))
		if err != nil {
			log.Fatalf("Ошибка при чтении Amount")
		}

		rate, err := strconv.ParseFloat(string(fields[4]), 64)
		if err != nil {
			log.Fatalf("Ошибка при чтении Rate")
		}

		currencys = append(currencys, Currency.Currency{
			Country:      string(fields[0]),
			CurrencyName: string(fields[1]),
			Amount:       amount,
			Code:         string(fields[3]),
			Rate:         rate,
		})
	}

	// todo сохранить данные в таблицу

	currencyJson, _ := json.Marshal(currencys)

	println(currencyJson)

	defer resp.Body.Close()

	println(dbCfg.DatabaseName)
}

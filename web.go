package main

import (
	// "fmt"
	"net/http"
	"strconv"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

const (
	OK       = iota
	ERRMYSQL
	ERRAUTH
)

type HttpServer struct {
	port  int
	Mysql *MySql
}

func CreateHttpServer(port int, mysql *MySql) (*HttpServer, error) {
	server := &HttpServer{ port, mysql }
	sport := strconv.Itoa(port)
	server.Register()
	err := http.ListenAndServe(":" + sport, nil)
	return server, err
}


func (s *HttpServer) Register() {

	http.HandleFunc("/userbots/list", func(w http.ResponseWriter, r *http.Request) {

		type Data struct{
			Code int		  `code`
			Msg  string		  `msg`
			Data interface{}  `data`
		}

		data := Data{ OK, "OK", "" }

		rows, err := s.Mysql.Query("SELECT * FROM userbots")
		if err != nil {
			data.Code = ERRMYSQL
			data.Msg = "MySql Error: " + err.Error()
		} else {
			type RowData struct {
				Ownerid string  `ownerid`
				Botid   string  `botid`
				Prefix  string  `prefix`
			}
			var rowDataSet []*RowData
			var dummy string
			for rows.Next() {
				rowData := RowData{}
				rows.Scan(&rowData.Ownerid, &rowData.Botid, &rowData.Prefix, &dummy, &dummy, &dummy)
				rowDataSet = append(rowDataSet, &rowData)
			}
			data.Data = rowDataSet
		}

		w.Header().Set("Content-Type", "application/json")
		bdata, _ := json.MarshalIndent(data, "", "    ")
		w.Write(bdata)
	})

}
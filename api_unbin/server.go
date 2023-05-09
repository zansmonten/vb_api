package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type booking struct {
	ID      uint   `json:"id"`
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Alamat  string `json:"alamat"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbs_vapi")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Mahasiswa
	e.GET("/tbl_booking", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_booking")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []booking
		for res.Next() {
			var m booking
			_ = res.Scan(&m.ID, &m.Nama, &m.Tanggal, &m.Alamat)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/tbl_booking", func(c echo.Context) error {
		var mahasiswa booking
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO mahasiswa (npm,nama, nohp,alamat)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Alamat)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/mahasiswa/:npm", func(c echo.Context) error {
		var mahasiswa booking
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_booking SET nama=?,tanggal=?,alamat=? WHERE id=?"
		res, err := db.Query(sqlStatement, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Alamat, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/mahasiswa/:npm", func(c echo.Context) error {
		var mahasiswa booking
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM mahasiswa WHERE npm=?"
		res, err := db.Query(sqlStatement, c.Param("npm"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Mahasiswa

	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	//create echo new instance
	e := echo.New()

	// serve static files from public directory(untuk mengakses direktory file)
	e.Static("/public", "public")

	//Routing
	e.GET("/hello-world", helloworld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)

	e.POST("/add-blog", addBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloworld(c echo.Context) error {
	// return c.String(http.StatusOK, "Helloworld")
	var tmpl, _ = template.ParseFiles("views/hello-world.html")
	//dibawah adalah menghilangkan tugas dari err dan mengabaikan error handling
	//jadi jika didlm Parsefiles ada yg salah maka error tidak akan berjalan
	// if err != nil {
	// 	// fmt.Println("Tidak ada datanya")
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }

	return tmpl.Execute(c.Response(), nil)
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	//dibawah adalah tugas dari err yaitu error handling
	if err != nil {
		// fmt.Println("Tidak ada datanya")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact-me.html")
	if err != nil {
		// fmt.Println("Tidak ada datanya")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		// fmt.Println("Tidak ada datanya")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))  //123 string => 123 int
	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		// fmt.Println("Tidak ada datanya")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Id": id,
		"Title": "Go to Beach",
		"Content": "Jalan jalan ke Pantai itu asik loh",
		"Author": "Chahyo Purnomo Aji",
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error{
	title := c.FormValue("title")
	content := c.FormValue("content")

	fmt.Println(title)
	fmt.Println(content)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
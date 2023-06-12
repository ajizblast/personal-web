package main

import (
	"context"
	"fmt"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

// kenapa struct tidak di deklarasi di dalam function main tapi di luar?
// agar nanti struct yg kita deklarasikan bisa dipakai di function lain

type Blog struct {
	ID 			int
	Title      string
	StartDate  time.Time
	EndDate    time.Time
	Duration   string
	Content    string
	Author     string
	NodeJs     bool
	ReactJs    bool
	NextJs     bool
	TypeScript bool
	Image 		string
	FormatStartDate 	string
	FormatEndDate 		string
}

// membuat tempat peampung slice, yang berguna sebagai data dummy yg akan muncul di web
// kenapa kita menggunakan slice? karena di slice bisa menggunakan pengoperasian di slice
var dataBlog = []Blog{
	// {
	// 	Title:      "Cara menjadi programmer yang baik dan benar",
	// 	Content:    "Untuk menjadi seorang programmer yang baik dan benar, penting memiliki pemahaman kuat tentang konsep dasar pemrograman, keterampilan teknis yang terus berkembang, kemampuan menganalisis masalah, kerja tim yang baik, kode yang bersih, dan etika profesional. Dengan mengikuti prinsip-prinsip ini, seseorang dapat menjadi programmer yang baik dan benar.",
	// 	Author:     "Chahyo Purnomo Aji",
	// 	StartDate:  "07/06/2023",
	// 	EndDate:    "08/06/2023",
	// 	Duration:   "1 hari",
	// 	NodeJs:     true,
	// 	ReactJs:    false,
	// 	NextJs:     true,
	// 	TypeScript: true,
	// },
	// {
	// 	Title:      "Cara membuat Javascript Menjadi Pacarmu",
	// 	Content:    "Untuk membuat JavaScript menjadi pacarmu yaitu memahami konsep dasar, berlatih secara aktif, manfaatkan sumber daya online, ikuti komunitas pengembang, dan terus pantau perkembangan terbaru. Dengan kesabaran dan konsistensi, kamu dapat menguasai JavaScript dengan baik dalam pengembangan web.",
	// 	Author:     "Chahyo Purnomo Aji",
	// 	StartDate:  "07/06/2023",
	// 	EndDate:    "09/06/2023",
	// 	Duration:   "2 hari",
	// 	NodeJs:     true,
	// 	ReactJs:    true,
	// 	NextJs:     false,
	// 	TypeScript: false,
	// },
}

func main() {
	
	connection.DatabaseConnection()
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
	e.GET("/edit-blog/:id", editBlog)

	e.POST("/add-blog", addBlog)
	// e.POST("/edit-blog/:id", submitEditBlog)
	e.POST("/delete-blog/:id", deleteBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

//Func GET //Func GET //Func Get

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
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, start_date, end_date, content, nodejs, nextjs, reactjs, typescript, image FROM tb_blog ORDER BY id ASC")

	var result []Blog
	for data.Next(){
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Content, &each.NodeJs, &each.NextJs, &each.ReactJs, &each.TypeScript, &each.Image)
		if err != nil{
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Author = "Chahyo"
		each.FormatStartDate = each.StartDate.Format("2 January 2006")
		each.FormatEndDate = each.EndDate.Format("2 January 2006")

		result = append(result, each)
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}


	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)
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
	id, _ := strconv.Atoi(c.Param("id")) //123 string => 123 int

	var blogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, content, nodejs, nextjs, reactjs, typescript, image FROM tb_blog WHERE id=$1", id).Scan(&blogDetail.ID, &blogDetail.Title, &blogDetail.StartDate, &blogDetail.EndDate, &blogDetail.Content, &blogDetail.NodeJs, &blogDetail.NextJs, &blogDetail.ReactJs, &blogDetail.TypeScript,  &blogDetail.Image)

	if err != nil {
	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogDetail.Author = "Chahyo Purnomo Aji"
	blogDetail.Duration = "0 Month"
	blogDetail.FormatStartDate = blogDetail.StartDate.Format("2 January 2006")
	blogDetail.FormatEndDate = blogDetail.EndDate.Format("2 January 2006")
		
	data := map[string]interface{}{
		"Blog":      blogDetail,
		"StartDate": getDateString(blogDetail.StartDate.Format("2006-01-02")),
		"EndDate":   getDateString(blogDetail.EndDate.Format("2006-01-02")),
	}

	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")

	if errTemplate != nil {
		// fmt.Println("Tidak ada datanya")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func editBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var blogDetail = Blog{}

	for index, data := range dataBlog {
		if id == index {
			blogDetail = Blog{
				Title:      data.Title,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				Duration:   data.Duration,
				Content:    data.Content,
				NodeJs:     data.NodeJs,
				ReactJs:    data.ReactJs,
				NextJs:     data.NextJs,
				TypeScript: data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": blogDetail,
		"Id":   id,
	}

	var tmpl, err = template.ParseFiles("views/edit-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

//Func POST //Func POST //Func POST

// func calculateDuration(startDate, endDate string) string {
// 	startTime, _ := time.Parse("2006-01-02", startDate)
// 	endTime, _ := time.Parse("2006-01-02", endDate)

// 	durationTime := int(endTime.Sub(startTime).Hours())
// 	durationDays := durationTime / 24
// 	durationWeeks := durationDays / 7
// 	durationMonths := durationWeeks / 4
// 	durationYears := durationMonths / 12

// 	var duration string

// 	if durationYears > 1 {
// 		duration = strconv.Itoa(durationYears) + " years"
// 	} else if durationYears > 0 {
// 		duration = strconv.Itoa(durationYears) + " year"
// 	} else {
// 		if durationMonths > 1 {
// 			duration = strconv.Itoa(durationMonths) + " months"
// 		} else if durationMonths > 0 {
// 			duration = strconv.Itoa(durationMonths) + " month"
// 		} else {
// 			if durationWeeks > 1 {
// 				duration = strconv.Itoa(durationWeeks) + " weeks"
// 			} else if durationWeeks > 0 {
// 				duration = strconv.Itoa(durationWeeks) + " week"
// 			} else {
// 				if durationDays > 1 {
// 					duration = strconv.Itoa(durationDays) + " days"
// 				} else {
// 					duration = strconv.Itoa(durationDays) + " day"
// 				}
// 			}
// 		}
// 	}

// 	return duration
// }

func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	nodeJs := c.FormValue("nodeJs")
	reactJs := c.FormValue("reactJs")
	nextJs := c.FormValue("nextJs")
	typescript := c.FormValue("typescript")
	image := c.FormValue("input-image")

	fmt.Println(title, content, nodeJs, reactJs, nextJs, typescript, image)

	_, err:= connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, start_date, end_date, nodejs, reactjs, nextjs, typescript, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", title, content, startDate, endDate, nodeJs, reactJs, nextJs, typescript, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// func submitEditBlog(c echo.Context) error {
	// id := getBlogIndex(c.Response(), c.Request())
	
	// title := c.FormValue("title")
	// content := c.FormValue("content")
	// startDate := c.FormValue("startDate")
	// endDate := c.FormValue("endDate")
	// nodeJs := c.FormValue("nodeJs")
	// reactJs := c.FormValue("reactJs")
	// nextJs := c.FormValue("nextJs")
	// typescript := c.FormValue("typescript")

	// startTime, _ := time.Parse("2006-01-02", startDate)
	// endTime, _ := time.Parse("2006-01-02", endDate)

	// var editedBlog = Blog{
	// 	Title:   title,
	// 	Content: content,
	// 	// Author:     "Chahyo Purnomo Aji",
	// 	StartDate:  startTime,
	// 	EndDate:    endTime,
	// 	NodeJs:     (nodeJs == "nodejs"),
	// 	ReactJs:    (reactJs == "reactjs"),
	// 	NextJs:     (nextJs == "nextjs"),
	// 	TypeScript: (typescript == "typescript"),
	// }

	// i, _ := strconv.Atoi(id)
	// dataBlog[i] = editedBlog
	// return c.Redirect(http.StatusMovedPermanently, "/blog-detail/"+id)
// }

// func getBlogIndex(w http.ResponseWriter, r *http.Request) string {
// 	// to call: getBlogIndex(c.Response(), c.Request())
// 	// value of url: /edit-project/0?
// 	// returned value: 0
// 	url := r.URL.String()
// 	lastSegment := path.Base(url)
// 	re := regexp.MustCompile("[0-9]+")
// 	return strings.Join((re.FindAllString(lastSegment, -1))[:], "")
// }

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID: ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// func deleteBlog(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	//fungsi strconv.Atoi untuk merubah string menjadi integer
// 	//id untuk menampung hasil konversi string ke integer

// 	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

// 	//diulang di video mas dandi karena masih kurang memahami 1:50:00

// 	return c.Redirect(http.StatusMovedPermanently, "/")
// }

func getDateString(date string) string {
	y := date[0:4]
	m, _ := strconv.Atoi(date[5:7])
	d := date[8:10]
	if string(d[0]) == "0" {
		d = string(d[1])
	}

	mon := ""
	switch m {
	case 1:
		mon = "Jan"
	case 2:
		mon = "Feb"
	case 3:
		mon = "Mar"
	case 4:
		mon = "Apr"
	case 5:
		mon = "Mei"
	case 6:
		mon = "Jun"
	case 7:
		mon = "Jul"
	case 8:
		mon = "Agu"
	case 9:
		mon = "Sep"
	case 10:
		mon = "Okt"
	case 11:
		mon = "Nov"
	case 12:
		mon = "Des"
	}

	return d + " " + mon + " " + y
}


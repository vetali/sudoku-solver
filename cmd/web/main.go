package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct {
	Grid  []int
	Level string
}

func newPage() Page {
	return Page{
		Grid:  make([]int, 81),
		Level: "Easy",
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/images", "assets/images")
	e.Static("/css", "assets/css")
	page := newPage()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})
	e.POST("/solve", func(c echo.Context) error {
		cell1Val := c.Param("cell-1")
		fmt.Printf("Cell submitted %v", cell1Val)
		return c.Render(200, "index", page)
	})
	e.POST("/contacts", func(c echo.Context) error {
		//name := c.FormValue("name")
		//email := c.FormValue("email")
		//if page.Data.hasEmail(email) {
		//	page.Form.Values["name"] = name
		//	page.Form.Values["email"] = email
		//	page.Form.Errors["email"] = "Email Already Exists"
		//	return c.Render(422, "createContact", page)
		//}
		//
		//contact := newContact(name, email)
		//page.Data.Contacts = append(page.Data.Contacts, contact)
		//c.Render(200, "createContact", newFormData())
		//return c.Render(200, "oob-contact", contact)
		return c.String(200, "")
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		//time.Sleep(1 * time.Second)
		//idStr := c.Param("id")
		//id, err := strconv.Atoi(idStr)
		//if err != nil {
		//	return c.String(400, "Invalid ID")
		//}
		//
		//index := page.Data.indexOf(id)
		//if index == -1 {
		//	return c.String(404, "Contact not found")
		//}
		//
		//page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)
		//
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":42069"))
}

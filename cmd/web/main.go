package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vetali/sudoku/pkg"
	"html/template"
	"io"
	"strconv"
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

type CellData struct {
	Value   int
	Options []int
}

func NewCellData() CellData {
	return CellData{0, make([]int, 0)}
}

func NewGrid() map[int]CellData {
	var grid = make(map[int]CellData)
	for i := 0; i < 81; i++ {
		grid[i] = NewCellData()
	}
	return grid
}

type Page struct {
	Grid  map[int]CellData
	Level string
	Valid string
}

func newPage() Page {
	return Page{
		Grid:  NewGrid(),
		Level: "Multiple Solutions",
		Valid: "Valid",
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/images", "assets/images")
	e.Static("/css", "assets/css")

	e.GET("/", func(c echo.Context) error {
		page := newPage()
		return c.Render(200, "index", page)
	})

	e.POST("/update", func(c echo.Context) error {
		page := newPage()
		cellValues := CellValuesFromForm(c)
		//fmt.Printf("Cell values: %+v", cellValues)
		sudoku := pkg.NewSudoku(cellValues)
		sudoku.UpdateOptions()

		for i, cell := range page.Grid {
			//fmt.Printf("Cell value at idx: %v is %v\n", i, sudoku.GetCellValue(i))
			cell.Value = sudoku.GetCellValue(i)
			for o := range sudoku.GetCellOptions(i) {
				cell.Options = append(cell.Options, o)
			}
			page.Grid[i] = cell
		}
		page.Level = sudoku.Level
		if sudoku.IsValid() {
			page.Valid = "Valid"
		} else {
			page.Valid = "Invalid"
		}

		return c.Render(200, "sudokuGrid", page)
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

func CellValuesFromForm(c echo.Context) [81]int {
	var cellValues = make([]int, 81)
	values, _ := c.FormParams()

	for i := 0; i < 81; i++ {
		cellInputName := fmt.Sprintf("cell_%d", i)
		if intVal, err := strconv.Atoi(values[cellInputName][0]); err == nil {
			cellValues[i] = intVal
		}
	}

	return [81]int(cellValues)
}

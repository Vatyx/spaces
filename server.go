package main

import  (
    "os"
    "database/sql"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    _ "github.com/lib/pq"
)

func main() {
    connection := os.Getenv("DATABASE_URL")
    _, err := sql.Open("postgres", connection)
    if err != nil {
        panic(err)
    }

    m := martini.Classic()

    m.Use(render.Renderer())
    m.Use(martini.Static("assets"))

    m.Get("/", func(r render.Render) {
      r.HTML(200, "index", nil)
    })

    m.Get("/buy", func(r render.Render) {
      r.HTML(200, "buy", nil)
    })

    m.Run()
}

package main

func main() {


	e := echo.New()


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

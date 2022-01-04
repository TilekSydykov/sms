package main

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"solar-faza/env"
	"solar-faza/repository/database"
	"solar-faza/routes"
	"solar-faza/utils/templateFunc"
	"strings"
	"syscall"
)

type globalRoutes struct {
	router *gin.Engine
}

func main() {
	port := env.GetEnvVars().Port
	// serve
	database.DB = database.Connect()
	database.SetupDB(database.DB)
	r := globalRoutes{
		router: gin.Default(),
	}
	r.router.SetFuncMap(templateFunc.GetFunctions())
	r.router.Static("/assets", "./assets")
	r.router.MaxMultipartMemory = 8 << 20
	r.router.LoadHTMLGlob("templates/**/*.gohtml")
	r.router.Use(location.Default())

	// web routes
	routes.MainRouter(r.router.Group(""))

	r.router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(fmt.Sprintf("%s", c.Request.URL), "/api/") {
			c.JSON(404, gin.H{
				"error": 404,
			})
			return
		}
		c.HTML(404, "404.gohtml", gin.H{})
	})
	log.Printf("Server up on port '%s'", port)
	r.Run(":" + port)
}

func (r globalRoutes) Run(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: r.router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//// catching ctx.Done(). timeout of 5 seconds.
	//select {
	//case <-ctx.Done():
	//	log.Println("timeout of 3 seconds.")
	//}

	log.Println("Server exiting")

}

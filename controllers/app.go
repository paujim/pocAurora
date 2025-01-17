package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/paujim/pocAurora/repositories"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

type App struct {
	repository repositories.SQLRepository
	router     *gin.Engine
}

func NewApp(repo repositories.SQLRepository, router *gin.Engine) *App {
	return &App{
		repository: repo,
		router:     router,
	}
}

func getArrayFromQuery(query string) []string {
	query = strings.ReplaceAll(query, "\"", "")
	query = strings.ReplaceAll(query, "[", "")
	query = strings.ReplaceAll(query, "]", "")
	return strings.Split(query, ",")
}

func validateIds(ids []string) error {
	for _, id := range ids {
		_, err := uuid.FromString(id)
		if err != nil {
			return fmt.Errorf("the %s is not valid", id)
		}
	}
	return nil
}

func (app *App) SetupServer() *gin.Engine {
	rest := app.router.Group("/rest")
	{
		v1 := rest.Group("/v1")
		{
			v1.GET("/racing", app.RaicingEndPoint)
		}
	}
	return app.router
}

func (app *App) RaicingEndPoint(c *gin.Context) {
	method := c.Query("method")
	count := c.Query("count")
	categories := getArrayFromQuery(c.Query("include_categories"))

	if err := validateIds(categories); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cnt, err := strconv.Atoi(count)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid parameter count",
		})
		return
	}

	if method != "nextraces-categorygroup" {
		c.JSON(400, gin.H{
			"error": "Invalid parameter method",
		})
		return
	}
	categoryRaceMap, raceSummaries, err := app.repository.GetNextRacesByCategory(cnt, categories)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"error": "Internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": map[string]interface{}{
			"category_race_map": categoryRaceMap,
			"race_summaries":    raceSummaries,
		},
	})
}

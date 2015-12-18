package main

import (
	//"time"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"github.com/goibibo/worktree"
	"fmt"
)

var dbmap = initDb()

type Article struct {
	Id int64 `db:"article_id"`
	Created int64
	Title string
	Content string
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func ArticleList(c *gin.Context) {
	var articles []Article
	_, err := dbmap.Select(&articles, "select * from articles order by article_id")
	checkErr(err, "select failed")
	content := gin.H{}
	for k, v := range articles {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

func ArticleDetail(c *gin.Context) {
	article_id := c.Params.ByName("id")
	a_id, _ := strconv.Atoi(article_id)
	article := getArticle(a_id)
	content := gin.H{"title" : article.Title, "content": article.Content}
	c.JSON(200, content)
}

func ArticlePost(c *gin.Context) {
	var json Article
	c.Bind(&json)
	article := createArticle(json.Title, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result": "Success",
			"title": article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occurred"})
	}
}

func createArticle(title, body string) Article {
	article := Article{
		Created:	10,
		Title:		title,
		Content: 	body,
	}
	err := dbmap.Insert(&article)
	checkErr(err, "Insert failed")
	return article
}

func getArticle(article_id int) Article {
	article := Article{}
	err := dbmap.SelectOne(&article, "select * from articles where article_id=?", article_id)
	checkErr(err, "SelectOne failed")
	return article
}


/*
** Code below is to test routes grouping
**
*/

func searchHotelInCity(c *gin.Context) {
	c.String(200, "Searching for hotels in selected city")
}

func getHotelReview(c *gin.Context) {
	c.String(200, "Please find below review of your selected hotel")
}

func searchFlightInCity(c *gin.Context) {
	c.String(200, "Searching for flight in selected city")
}

func getFlightReview(c *gin.Context) {
	c.String(200, "Please find below review of your selected flight")
}


func AddHotelRoutes(group *gin.RouterGroup) {
	group.GET("/search", searchHotelInCity)
	group.GET("/review", getHotelReview)
}

func AddFlightRoutes(group *gin.RouterGroup) {
	group.GET("/search", searchFlightInCity)
	group.GET("/review", getFlightReview)
}

/*
** Testing goibibo/worktree. Just want to check its usefulness
*/

type TwoArgs struct {
	X int
	Y int
}

func leaf1(i interface{}) interface{} {
	args := i.(TwoArgs)
	return args.X + args.Y
}

func leaf2(i interface{}) interface{} {
	args := i.(TwoArgs)
	return args.X * args.Y
}

func merge2(results []interface{}) interface{} {
	var sum int
	for _, x := range results {
		sum += x.(int)
	}
	return sum
}

func main() {
	fmt.Println("inside main")

	//Testing goibibo/worktree. Just want to check its usefulness
	// TWO Level work tree
	l2 := worktree.CommandTree{}
	l2.AddMapper(leaf2, TwoArgs{2, 3})
	l2.AddMapper(leaf2, TwoArgs{2, 2})
	l2.AddReducer(merge2)

	l1 := worktree.CommandTree{}
	l1.AddMapper(l2.Run, nil) // When nesting use nil for Run
    l1.AddMapper(leaf2, TwoArgs{2, 2})
    l1.AddReducer(merge2)

    fmt.Println("printing the result: ",l1.Run(nil).(int))


    //Routers and grouping
	app := gin.Default()
	v1 := app.Group("/v1")

	v1_hotels := v1.Group("/hotels")
	AddHotelRoutes(v1_hotels)

	v1_flight := v1.Group("/flight")
	AddFlightRoutes(v1_flight)

	app.GET("/", index)
	app.GET("/articles", ArticleList)
	app.POST("/articles", ArticlePost)
	app.GET("/articles/:Article_id", ArticleDetail)
	app.Run(":8080")
}
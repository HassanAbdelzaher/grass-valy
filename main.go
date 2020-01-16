package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	BasePath string `json="BasePath"`
}
var appConfig Config
func init(){
	dat, err := ioutil.ReadFile("config.json")
	if err!=nil {
		log.Println("can not load config file")
		log.Println(err)
		time.Sleep(3*time.Second)
		panic(err)
	}
	err=json.Unmarshal(dat,&appConfig)
	if err!=nil {
		log.Println("can not parse config file")
		log.Println(err)
		time.Sleep(3*time.Second)
		panic(err)
	}
}
type Movie struct {
	BinId     string `db:"binId"`
	Id        string `db:"id"`
	Name      string `db:"name"`
	KeyName   string `db:"keyName"`
	AssetType string `db:"assetType"`
	SourceId  string `db:"sourceId"`
	Broken bool `db:_`
}
type MovieInfo struct {
	Movie
	BinName string
	ParentName string
	ParentType string
	ParentBinName string
	SubMovie []*Movie
}

type Bin struct {
	BinId    string `db:"binId"`
	BinName  string `db:"binName"`
	ParentId string `db:"parentId"`
}

type Reference struct {
	AssetId string `db:"assetId"`
	RefId   string `db:"refId"`
}

func main() {
	//loading config dile

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com","*"}
	router.Use(cors.New(config))
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/api/bin/:binId", func(c *gin.Context) {
		binId := c.Param("binId")
		bins, err := GetBinMovies(binId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.GET("/api/bin", func(c *gin.Context) {
		bins, err := GetBins()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.GET("/api/movie/:movieId", func(c *gin.Context) {
		movieId := c.Param("movieId")
		movie,err:=GetMovie(movieId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, movie)
	})
	router.GET("/api/list/:listId", func(c *gin.Context) {
		listId := c.Param("listId")
		bins, err := GetListMovies(listId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.GET("/api/list", func(c *gin.Context) {
		bins, err := GetLists()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.StaticFS("/static", http.Dir("public/static"))
	router.StaticFile("/", "./public/index.html")
	router.Run(":8080")
}

func GetBins() ([]Bin, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var bin []Bin
	err = db.Select(&bin, "select binId,binName,parentId from MovieGroup")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bin, err
}

func GetLists() ([]Movie, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var bin []Movie
	err = db.Select(&bin, "select binId,id,name,assetType,sourceId from Movie where assetType='List'")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bin, err
}

func GetListMovies(id string) ([]Movie, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var movies []Movie
	quer := fmt.Sprintf("select a.binId,a.id,a.name,a.assetType,sourceId from Movie a,Reference b where a.id=b.refid and b.assetid='%s'", id)
	err = db.Select(&movies, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return movies, err
}

func GetBinMovies(binId string) ([]*MovieInfo, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var movies []string
	var bin Bin
	err=db.Get(&bin,fmt.Sprintf("select binId,binName from MovieGroup where binId='%s'",binId))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	quer := fmt.Sprintf("select id from Movie where binId='%s'", binId)
	err = db.Select(&movies, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	info:=make([]*MovieInfo,0)
	for _,mov:=range movies {
		mf,err:=GetMovie(mov)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		info=append(info,mf)
	}
	return info, err
}

func GetAssetRefs(assetId string) ([]string, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var refs []string
	quer := fmt.Sprintf("select refId from Reference where assetId='%s'", assetId)
	err = db.Select(&refs, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return refs, err
}

func GetRefAssets(refId string) ([]string, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var assets []string
	quer := fmt.Sprintf("select assetId from Reference where refId='%s'", refId)
	err = db.Select(&assets, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return assets, err
}

func ListTables() ([]string, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var tables []string
	err = db.Select(&tables, "select name from sqlite_master WHERE type='table'")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return tables, err
}

func GetMovie(movieId string) (*MovieInfo, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var mov Movie
	quer := fmt.Sprintf("select binId,id,name,assetType,sourceId from Movie where Id='%s'", movieId)
	err = db.Get(&mov, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	binId:=mov.BinId
	var bin Bin
	err=db.Get(&bin,fmt.Sprintf("select binId,binName from MovieGroup where binId='%s'",binId))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pth:=appConfig.BasePath+"/"+strings.TrimSpace(bin.BinName)+"/"+strings.TrimSpace(mov.Name)
	_,err=os.Stat(pth)
	if os.IsNotExist(err) {
		mov.Broken=true
	}
	var info MovieInfo
	info.Movie=mov
	info.BinName=bin.BinName
	if info.AssetType=="SubClip"{
		refId,err:=GetAssetRefs(info.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if len(refId)==1{
			pInfo,err:=GetMovie(refId[0])
			if err != nil {
				log.Println(err)
				return nil, err
			}
			info.ParentName=pInfo.Name
			info.ParentType=pInfo.AssetType
			info.ParentBinName=pInfo.BinName
			if pInfo.Broken{
				info.Broken=pInfo.Broken
			}
		}
	}
	if info.AssetType=="List"{
		assets,err:=GetAssetRefs(info.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if assets!=nil {
			info.SubMovie=make([]*Movie,0)
			for _,ass:=range assets{
				mif,err:=GetMovie(ass)
				if err==nil {
					info.SubMovie=append(info.SubMovie,&mif.Movie)
				}
			}
		}

	}
	return &info, nil
}

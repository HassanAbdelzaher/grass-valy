package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	BinId     string `db:"binId"`
	Id        string `db:"id"`
	Name      string `db:"name"`
	KeyName   string `db:"keyName"`
	AssetType string `db:"assetType"`
	SourceId  string `db:"sourceId"`
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
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/bin/:binId", func(c *gin.Context) {
		binId := c.Param("binId")
		bins, err := GetBinMovies(binId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.GET("/bin", func(c *gin.Context) {
		bins, err := GetBins()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})

	router.GET("/list/:listId", func(c *gin.Context) {
		listId := c.Param("listId")
		bins, err := GetListMovies(listId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})
	router.GET("/list", func(c *gin.Context) {
		bins, err := GetLists()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err.Error())
			return
		}
		c.JSON(http.StatusAccepted, bins)
	})

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

func GetBinMovies(binId string) ([]Movie, error) {
	db, err := sqlx.Connect("sqlite3", "media.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var movies []Movie
	quer := fmt.Sprintf("select binId,id,name,assetType,sourceId from Movie where binId='%s'", binId)
	err = db.Select(&movies, quer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return movies, err
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
	err = db.Select(&refId, quer)
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

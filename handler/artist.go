package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgkg/project/model"
	"github.com/gin-gonic/gin"
)

func GetAllArtist(ctx *gin.Context) {

	var listArtists []model.Artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := RequestGet(url, &listArtists)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	contentType := ctx.Request.Header.Get("Content-type")
	if strings.ContainsAny(contentType, "application/json") {
		ctx.JSON(http.StatusOK, listArtists)
		return
	}

	ctx.HTML(http.StatusOK, "list-artist", gin.H{
		"title":   "List of artists",
		"artists": listArtists,
	})
}

const (
	apiURL = "https://groupietrackers.herokuapp.com/api/artists/"
)

func GetArtist(ctx *gin.Context) {

	idstring := ctx.Param("id")
	url := apiURL + idstring
	var artist model.Artist
	err := RequestGet(url, &artist)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var location model.Location
	if len(artist.Locations) != 0 {
		err = RequestGet(artist.Locations, &location)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}

	var date model.Date
	if len(artist.Concertdates) != 0 {
		err = RequestGet(artist.Concertdates, &date)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}

	var relation model.Relation
	if len(artist.Relations) != 0 {
		err = RequestGet(artist.Relations, &relation)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}

	contentType := ctx.Request.Header.Get("Content-type") // header(': application/json; charset=utf-8');
	if strings.ContainsAny(contentType, "application/json") {
		ctx.JSON(http.StatusOK, artist)
		return
	}

	ctx.HTML(http.StatusOK, "artist", gin.H{
		"title":        artist.Name,
		"image":        artist.Image,
		"name":         artist.Name,
		"members":      artist.Members,
		"creationDate": artist.Creationdate,
		"firstAlbum":   artist.Firstalbum,
		"locations":    location,
		"concertDates": date,
		"relations":    relation,
	})
}

func RequestGet(url string, bind interface{}) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, bind)
}

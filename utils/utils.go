package utils

import (
	"errors"
	"fifa-telegram-bot/models"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetTeamsFromSoFifa gets the teams from the SoFifa website.
func GetTeamsFromSoFifa() ([]models.Team, error) {
	var teams []models.Team
	var team models.Team

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("#body > div.center > div > div.col.col-12 > div > table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			team.Name = el.ChildText("td.col-name-wide > a:first-child > div")
			team.League = el.ChildText("td.col-name-wide > a:nth-child(2) > div")
			team.Overall, _ = strconv.Atoi(el.ChildText("td.col.col-oa > span"))
			team.Attack, _ = strconv.Atoi(el.ChildText("td.col.col-at > span"))
			team.Midfield, _ = strconv.Atoi(el.ChildText("td.col.col-md > span"))
			team.Defence, _ = strconv.Atoi(el.ChildText("td.col.col-df > span"))
			flagLink := el.ChildAttr("td.col-avatar > figure.avatar > img", "data-src")
			err := downloadFile(flagLink, team.Name+".png")
			if err != nil {
				fmt.Println(err)
			}

			team.Image = team.Name + ".png"
			teams = append(teams, team)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for i := 0; i <= 180; i = i + 60 {
		err := c.Visit("https://sofifa.com/teams?type=club&oal=70&oah=99&offset=" + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
	}

	return teams, nil

}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	// create a folder if it doesn't exist
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		fmt.Println("Creating images folder")
		os.Mkdir("images", 0755)
	}
	//Create a empty file
	if strings.Contains(fileName, "/") {
		fileName = strings.ReplaceAll(fileName, "/", "\\")
	}
	file, err := os.Create("images/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// GenerateTeams generates 2 random teams from the list of teams
func GenerateTeams(teams []models.Team) []models.Team {

	// generate a random number between 0 and len(teams)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randNum := rand.Intn(len(teams))

	fmt.Println("randNum: ", randNum)
	//select 1 random team
	team1 := teams[randNum]

	//select 1 random team
	team2 := teams[randNum]

	//if team1 == team2, select another team
	for team1.Name == team2.Name {
		team2 = teams[randNum+1]
	}

	var generatedTeams []models.Team

	generatedTeams = append(generatedTeams, team1)
	generatedTeams = append(generatedTeams, team2)

	return generatedTeams
}

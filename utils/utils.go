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
		err := c.Visit("https://sofifa.com/teams?type=club&oal=75&oah=99&offset=" + strconv.Itoa(i))
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

	// generate 2 random numbers
	randNum1, randNum2 := GenerateTwoRandomNumbers(len(teams))

	//select first random team with the first random number
	team1 := teams[randNum1]

	//select second random team with the second random number
	team2 := teams[randNum2]

	// add the 2 teams to a slice
	var generatedTeams []models.Team
	generatedTeams = append(generatedTeams, team1)
	generatedTeams = append(generatedTeams, team2)

	return generatedTeams
}

// GenerateTwoRandomNumbers generates 2 random numbers within the range of the length of the teams
func GenerateTwoRandomNumbers(len int) (int, int) {
	// generate a random number between 0 and len(teams)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randNum1 := rand.Intn(len)
	randNum2 := rand.Intn(len)

	randNum2 = randNum1
	// if the 2 random numbers are the same, generate a new random number for the second number
	for randNum1 == randNum2 {
		randNum2 = rand.Intn(len)
	}

	return randNum1, randNum2
}

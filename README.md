#   FIFA-RANDOM-TEAM-GENERATOR-BOT

## Description
This is a Telegram bot that allows you to randomly select a team from the FIFA 23 game. The bot is written in Go and uses the [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api/tree/master) library for the Telegram Bot API.
And also it uses [colly](https://github.com/gocolly/colly) for scraping teams data from the [SoFifa](https://sofifa.com/teams) website.

You can use this bot to play with your friends in FIFA 23 when you don't know which team to choose or you want a little more excitement :)

## Usage
To use the bot, you need to add it to your Telegram account. You can do this by following the link: [Fifa Random Team Generator Bot](https://t.me/fifa23_random_team_generator_bot)

## Commands

### The bot has the following commands:

- `/start` - start the bot and get a list of commands
- `/help` - get help

- `/random` - generates 2 random teams with random leagues and ratings
- `/update` - updates the list of teams


## Todo
- [ ] `/unique` - generates 2 random teams with unique leagues and ratings
- [ ] `/random <number>` - get a random number of teams
- [ ] `/random <rating>` - get a random team with the selected rating
- [ ] `/random <country>` - get a random team from the selected country
- [ ] `/random <league>` - get a random team from the selected league
- [ ] `/random <number> <league>` - get a random number of teams from the selected league
- [ ] `/random <rating> <league>` - get a random team with the selected rating and league
- [ ] `/random <number> <country>` - get a random number of teams from the selected country
- [ ] `/random <number> <rating>` - get a random number of teams with the selected rating
- [ ] `/random <number> <league>` - get a random number of teams from the selected league
- [ ] `/random <rating> <league> <country>` - get a random team with the selected rating, league and country
- [ ] `/random <number> <league> <rating>` - get a random number of teams from the selected league and rating
- [ ] `/random <number> <league> <country>` - get a random number of teams from the selected league and country
- [ ] `/random <number> <country> <rating>` - get a random number of teams from the selected country and rating
- [ ] `/random <number> <league> <country> <rating>` - get a random number of teams from the selected league, country and rating
- 

## License
[MIT](https://choosealicense.com/licenses/mit/)

# Discord Webhook for CROUS

This is a simple script that sends a message to a Discord webhook with the menu of the day from the CROUS website.

This uses [HackTheCrous](https://hackthecrous.com)'s api to get the menu.

## Installation

### From sources

1. Clone the repository
2. Build with go with `go build`
3. Enjoy !

### From releases

1. Download the latest release from the [releases page](https://github.com/remi-espie/discord-webhook-for-crous-meals/releases)
2. Enjoy !

## Usage

Configure the webhook however you want by renaming `.env.example` to `.env` (or use environment variable !) and customize it with your own webhook.

The available environment variables are:
- `WEBHOOK_URL`: URL of your Discord webhook
- `AVATAR_URL`: URL of an image to use as the webhook's avatar
- `CROUS_RESTAURANT_ID`: ID of the CROUS restaurant you want to track, per HackTheCrous
- `WITH_CALENDAR`: optional, a link to a .ics file. If provided, the bot will only post if there is an event today in the calendar

Then, run the script with `./discord-webhook-for-crous` or `go run main.go`.

Personally, I use a cron job to run this script every day at 11:05 AM, as it is the latest time the menu will be updated.

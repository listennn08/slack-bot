# Slack Bot

## Prerequisite

### Install the Slack Bot

1. Click Click `Create App` button

![Screen Shot 2023-07-30 at 18 16 59](https://github.com/listennn08/slack-bot/assets/52522402/fddd08d6-714a-4027-ba8e-86a0935b0eee)

2. Click `From scratch`

![Screen Shot 2023-07-30 at 18 17 04](https://github.com/listennn08/slack-bot/assets/52522402/24d43b01-370d-44a5-a487-1a522fd74b37)

3. Click `Bots`

![Screen Shot 2023-07-30 at 18 18 16](https://github.com/listennn08/slack-bot/assets/52522402/bfe03720-4a30-48bd-a537-f7b0931b278c)

4. Click `Review Scopes to Add`

![Screen Shot 2023-07-30 at 18 18 22](https://github.com/listennn08/slack-bot/assets/52522402/19b46826-fb9f-4c08-9157-70aa759e6d0c)

5. Click `Add an OAuth Scope` then choose the below scope

![Screen Shot 2023-07-30 at 18 18 27](https://github.com/listennn08/slack-bot/assets/52522402/f8e18613-7e59-488c-9b15-494b9ba9c9f7)
![Screen Shot 2023-07-30 at 18 19 02](https://github.com/listennn08/slack-bot/assets/52522402/aa6f0713-0328-4f60-b2b0-dcddcc4bf0cc)

6. Generate the App-level token with the below scope in `Basic Information`, and copy token

![Screen Shot 2023-07-30 at 18 13 22](https://github.com/listennn08/slack-bot/assets/52522402/1f8ba0dc-e8f2-4fc7-898f-78e45aa30200)

7. Enable Socket Mode

![Screen Shot 2023-07-30 at 18 14 01](https://github.com/listennn08/slack-bot/assets/52522402/186810f7-266d-4f59-8104-d1cd1f19209d)

8. Click `Install app to workspace`, then choose which workspace you want to install

![Screen Shot 2023-07-30 at 18 19 30](https://github.com/listennn08/slack-bot/assets/52522402/c124be34-3052-4c91-bc72-51e333ee9ba4)

9. Enable Event Events, and add the below scope

![Screen Shot 2023-07-30 at 22 41 39](https://github.com/listennn08/slack-bot/assets/52522402/f0b4c3d4-0d2c-4783-ac36-6851873baf6c)


### Set the service config

1. Copy the `.env.example` as `.env`
2. Paste your app token to `SLACK_APP_TOKEN`
3. From `Basic Information` to copy your app's `Bot User OAuth Token` to `SLACK_AUTH_TOKEN`

### Create Sheet Database

1. Create a Google sheet and public this file to the internet with `csv`
2. Copy the URL to `SHEET_URL` in `.env`

## Development

```make
# Run this command to open local service
make dev

# Build the project
make build

# Clean the build
make clean
```

## Docker

```bash
# Build the image
docker build -t slack-bot:lastest .

# Run the docker container
docker run -d --env-file .env slack-bot
# or
docker run -d \                                                                                                                                                                                                                                         22:51:26
           -e SLACK_AUTH_TOKEN={token} \
           -e SLACK_APP_TOKEN={app token} \
           -e SHEET_URL={fully sheet url} \
           slack-bot
```

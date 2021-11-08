# League teamfinder API

<!-- Section for your links, references, etc. --->

[//]: # "References"
[logo]: https://via.placeholder.com/900x300/000000/FFFFFF/?text=project+logo
[shields-badge]: https://img.shields.io/badge/make%20your%20own%20badges-on%20shields.io-brightgreen.svg
[issue-tracker]: #
[contributor-one-img]: https://via.placeholder.com/150?text=profile+image
[contributor-one-link]: #
[contributor-two-img]: https://via.placeholder.com/150?text=profile+image
[contributor-two-link]: #
[contributor-three-img]: https://via.placeholder.com/150?text=profile+image
[contributor-three-link]: #
[license]: #
[sphinx]: https://www.sphinx-doc.org/en/master/
[mkdocs]: https://www.mkdocs.org/
[gitbook]: https://www.gitbook.com/
[bibtex-wikipedia]: https://en.wikipedia.org/wiki/BibTeX
[go-official-site]: https://golang.org/doc/install
[docker-install]: https://docs.docker.com/install/
[open-api2-link]: https://swagger.io/specification/v2/
[golang-install]: https://golang.org/doc/install
[vscode-go]: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
[vscode-prettier]: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode

<!-- Your project's logo --->

![Your project's logo][logo]

<!-- Your badges --->

[![shields.io badge][shields-badge]](https://shields.io)

<!-- One liner about your project --->

Golang api for a league of legends teamfinder project

## Table of Contents

- [League teamfinder API](#league-teamfinder-api)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
      - [Run locally](#run-locally)
      - [Run with docker](#run-with-docker)
    - [Callable routes](#callable-routes)
  - [License](#license)
  - [Citation](#citation)
  - [Contact](#contact)

## Introduction

The following documentation assumes previous knowledge of:

- Docker containers
- General API usage (call a route, differences between get, post, patch operations, send a request with body or with data on headers, process a response...)
- Understanding of the purpose of the API; being able to answer the questions:
  - What the API has ben built for?
  - What resources are managed by it

## Getting Started

### Prerequisites

- Install golang from the reference [guide][golang-install]
- Install [docker][docker-install] from the official site
- Optional: vscode addons for golang
  - [Vscode][vscode-go] Golang addon
  - [Prettier][vscode-prettier] for code formatting

### Installation

- Open a terminal and navigate to your golang working directory

  ```sh
  cd PATH/TO/YOUR/GOLANG/WORK/DIRECTORY
  ```

- Run the command on the terminal to clone the project from the gitlab repository

   ```sh
    git clone https://gitlab.com/lol_fashion/lol-api.git
   ```

#### Run locally

1. You will find a folder named lol-api, navigate inside the folder with the command

   ```sh
   cd ./lol-api`
   ```

2. Run this command to install all the necessary golang packages

   ```sh
   go get -v ./...
   ```

3. Paste the following data into the .env file inside the folder and update the variables

    ```yml
    version: "3.7"
    services:

    lol-api:
        build: ./lol-api
        image: lol-api:latest
        restart: always
        container_name: lol-api
        depends_on:
        environment:
            PORT: your_db_port
            GO_ENV: development
            SERVER_LOGS_PATH: logs/server_info.log
            FRONT_LOGS_PATH: logs/front_info.log
            CORS_ALLOWED_ORIGINS: your_allowed_url_1, your_allowed_url_2
            ACCESS_CONTROL_ALLOW_CREDENTIALS: your_allowed_url_1, your_allowed_url_2
            CORS_ADMIN_ORIGINS: your_allowed_url_1, your_allowed_url_2
            DB_ENV: dev
            DB_PORT: your_rethinkdb_port
            DB_HOST: your_rethinkdb_host
            MAIL_SEND_TO: your_email
            MAIL_USER: your_email_user
            MAIL_PASS: your_email_pass
        volumes:
        - ./mounted/logs:/go/src/lol-api/logs
        ports:
        - "8082:81"
    ```

4. After completing the 3 previous steps, the project should be configured and ready to run, in order to launch it execute the following command, which will start the server and will show the logs on the terminal

    ```sh
    go run app.go
    ```

#### Run with docker

1. Create a file called `docker-compose.yml` in the same directory of the lol-api folder

2. Paste the following content into the `docker-compose.yml` file:

    ```yml
    version: "3.7"
    services:

    lol-api:
        build: ./lol-api
        image: lol-api:latest
        restart: always
        container_name: lol-api
        depends_on:
        environment:
            PORT: your_db_port
            GO_ENV: development
            SERVER_LOGS_PATH: logs/server_info.log
            FRONT_LOGS_PATH: logs/front_info.log
            CORS_ALLOWED_ORIGINS: your_allowed_url_1, your_allowed_url_2
            ACCESS_CONTROL_ALLOW_CREDENTIALS: your_allowed_url_1, your_allowed_url_2
            CORS_ADMIN_ORIGINS: your_allowed_url_1, your_allowed_url_2
            DB_ENV: dev
            DB_PORT: your_rethinkdb_port
            DB_HOST: your_rethinkdb_host
            MAIL_SEND_TO: your_email
            MAIL_USER: your_email_user
            MAIL_PASS: your_email_pass
        volumes:
        - ./mounted/logs:/go/src/lol-api/logs
        ports:
        - "8082:81"
    ```

3. To launch the container and have the system running, execute the following command on the terminal:

    ```sh
    sudo docker-compose up --build -d
    ```

    This command will build the project, copy the resulting executable into the docker container and launch it

4. To monitor the system is running, execute the following command which will display the logs

    ```sh
    sudo docker-compose logs -t -f
    ```

### Callable routes

Those are the core resources of the REST API and the routes used to modify them

| Resource    | Route                             | Method | Information                               | Params                |
| :---------- | :-------------------------------- | :----: | :---------------------------------------- | :-------------------- |
| collections | `/api/v1/GtS2Fa`                  |  GET   | Returns the collection names.             |                       |
| collections | `/api/v1/5JUSUk`                  |  GET   | Returns all items collections.            | collection_id         |
| collections | `/api/v1/5JUSUk/l9ew8v`           |  GET   | Items collections overview.               | collection_id         |
| collections | `/api/v1/5JUSUk/3fj48f`           |  GET   | Returns three items fron new collections. |                       |
| ideas       | `/api/v1/AcD9pq`                  |  GET   | Returns styng ideas of a user.            |                       |
| ideas       | `/api/v1/AcD9pq/8g73jg`           |  GET   | Returns liked styling ideas.              |                       |
| ideas       | `/api/v1/AcD9pq/j3f84f`           |  GET   | Returns disliked styling ideas.           |                       |
| ideas       | `/api/v1/AcD9pq/MtcZ29`           |  GET   | Returns styling ideas properties.         | idea                  |
| ideas       | `/api/v1/AcD9pq`                  | PATCH  | Updates idea to liked state.              | ideaName & ideaStatus |
| ideas       | `/api/v1/AcD9pq/PgwS23`           | PATCH  | Updates idea to seen state.               | ideaName              |
| ideas       | `/api/v1/AcD9pq/gmr9qj`           |  GET   | Returns items of ideas.                   |                       |
| picks       | `/api/v1/nv2sdo/82fj93`           |  GET   | Returns new picks.                        |                       |
| picks       | `/api/v1/nv2sdo/zw4w84`           |  GET   | Returns old picks.                        |                       |
| wishlist    | `/api/v1/BiT7gF/{ITEM_ID}`        | DELETE | Deletes item from wishlist.               |                       |
| wishlist    | `/api/v1/BiT7gF/{ITEM_ID}`        |  GET   | Returns wishlist.                         |                       |
| Wishlist    | `/api/v1/BiT7gF/j3f84f`           |  GET   | Returns disliked items.                   |                       |
| wishlist    | `/api/v1/BiT7gF`                  | PATCH  | Adds item to wishlist.                    | itemid struct         |
| items       | `/api/v1/5JUSUk/{ITEM_ID}`        | DELETE | Deletes item from wishlist.               |                       |
| items       | `/api/v1/5JUSUk/gmr9qj/{ITEM_ID}` |  GET   | Returns single item.                      |                       |
| user        | `/api/v1/Pq3Gm9`                  | PATCH  | Updates userData.                         |                       |
| user        | `/api/v1/Pq3Gm9/j484hs`           |  GET   | Returns  user preferences.                |                       |
| login       | `/api/v1/G35df2/Pdfe32`           |  POST  | post login.                               | code & state          |
| mailchimp   | `/api/v1/FIYnCI/t3Sf9a`           |  POST  | mailchimp subscription.                   |                       |
| survey      | `/api/v1/nZiejS`                  |  POST  | post survey.                              |                       |
| wardrobe    | `/api/v1/jfm2gw`                  |  GET   | Returns wardrobe.                         |                       |
| stylist     | `/api/v1/ZQdG8p`                  |  GET   | Returns stylist data.                     |                       |
| creator     | `/api/v1/abc`                     |  GET   | Returns arrivals from about you.          | aylimit & aypage      |


## License

Provide your license information here and link to your [LICENSE.md][license] file.

## Citation

Explain how other people can cite your project.

Make sure you provide a [BibTeX formatted reference][bibtex-wikipedia] e.g.:

```
@misc{yourlabel,
 author    = {Name of the authors},
 title     = {Your project's title},
 publisher = {Your publisher},
 year      = {2019},
}
```

## Contact

Put here your contact details so contributors and team members can reach out to you if they have open questions (or want to compliment you).

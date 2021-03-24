# minecraft-mobs
Minecraft Mob Wiki written in Go. Our app is hosted [here](https://minecraft-mobs.herokuapp.com/)

## Background
This project is designed to take a step back from all the common programming languages we have learned about and explore a non-traditional programming language. We choose Go as ours and will implement a basic web application that will be hosted on Heroku, so as to utilize the Cloud.
We are developing a Mob Wiki for the classic game Minecraft. This wiki will show basic information such as health, spawn info, and items dropped, related to various mobs in the game. 

## Developer Set-up
### Useful Resources
- [Install Golang](https://golang.org/doc/install)

- [Getting Started with Golang on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go)

- [Setting Up App for Heroku](https://elements.heroku.com/buildpacks/heroku/heroku-buildpack-go)

- [Deploying Own App to Heroku](https://devcenter.heroku.com/articles/preparing-a-codebase-for-heroku-deployment)

## Developer Notes

### Development Pipeline
General development should look like: Local -> Heroku Deployment -> Git Push to Github

### Local Development
To test your code, you run the following in the top most directory:
>`go build main.go`

>`./main`

### Adding new modules in Go
Each new module should be in its respective directory within the `src` directory of the project. Each new module also needs its respective `go.mod` file that can be generated with:
>`go mod init example.com/module`

Whenever one module requires other *local* module(s), you must replace the path of the module to its relative local path (see existing `go.mod` files for examples).

After assigning the replacement paths, perform a `go mod tidy` in the directory of the `go.mod` file you would like to clean up.

### Deploying to Heroku
Since our Github codebase is not directly linked with Heroku, Heroku and Github contains two different copies of the project.

In order to update to Heroku, you must first connect to Heroku remote like this:
>`git remote: -a tranquil-taiga-90293`

Then everytime you'd like to push, you can:
>`git push heroku`

The code on the main branch will be automatically deployed.

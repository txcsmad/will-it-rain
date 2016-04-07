# will-it-rain

Find out if it will rain in Austin today! This is part 8 of the [spring 2016 web workshops](https://github.com/txcsmad/s16-web).

![](http://i.imgur.com/Annd1bi.png)

## What we built

Web app that tells you if it will rain in Austin today. When a user makes a request, the Go server gets the forecast information as JSON, decodes it, and uses the precipitation probability to come up with a "yes" or "no" answer. The answer is rendered to the user using Go's built-in html template package.

## Topics

- setting up a web server using the `net/http` package
- using the `html/template` package to make simple HTML pages
- error handling
- making HTTP requests, unmarshaling JSON using `encoding/json`

## Local development

Install [Go](https://golang.org/dl/). 

Clone this repo and switch into it:

```
$ go get https://github.com/txcsmad/will-it-rain
$ cd $GOPATH/src/github.com/txcsmad/will-it-rain
```

Get a [Forecast IO API key](https://developer.forecast.io/). Then set the environment variable:

```
$ export FORECASTIO_API_KEY=<your-api-key>
```

Run local server.

```
$ go build
$ ./will-it-rain
```

Then visit `http://localhost:8080`

## Questions

* Nishanth Shanmugham [email](mailto:nishanths@utexas.edu)
* Brian Cui

## License

[MIT](http://nishanths.mit-license.org)

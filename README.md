# Indonesian Numeral Spellers

Indonesian Numeral Spellers is a REST API that created using Go Programming Language for dealing with numeral translation to its spelled out in Bahasa and vice versa.

## Requirements

In order to run you should install :
1. Go Programming Language
2. Node JS (optional)
3. ReactJS (optional)

## Usage

After you installed the requirements, clone this repo in $GOPATH/src/github.com. 
```bash
    go install github.com/Indonesian-Numeral-Spellers/app
```

run app.exe in bin directory file created in $GOPATH
the server will run at http://localhost:8080

I make the webapp that used this API. To run this webapp you should install all optional requirments.

```bash
    cd webapp
    npm install
    npm start
```

## API Endpoints
#### GET /spell?number={num}
Example Response
```json
    {
        "status":"OK",
        "text": "seratus",
    }
```

#### POST /read
Example Response
```json
    {
        "status":"OK",
        "number": 100,
    }
```

## References
1. https://golang.org/doc/install
2. https://github.com/golang/go/wiki/SettingGOPATH
3. https://golang.org/doc/code.html#Introduction
4. https://tour.golang.org/welcome/1
5. https://openclassrooms.com/en/courses/3432056-build-your-web-projects-with-rest-apis/3496011-identify-examples-of-rest-apis
6. https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
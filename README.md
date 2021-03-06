
[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community)
 [![CircleCI](https://circleci.com/gh/ledongthuc/hemnetparser.svg?style=svg)](https://circleci.com/gh/ledongthuc/hemnetparser) [![GoDoc](https://godoc.org/github.com/ledongthuc/hemnetparser?status.svg)](https://godoc.org/github.com/ledongthuc/hemnetparser) [![Go Report Card](https://goreportcard.com/badge/github.com/ledongthuc/hemnetparser)](https://goreportcard.com/report/github.com/ledongthuc/hemnetparser)

# Hemnet Parser

Library and a simple commandline to parse information from Hemnet URL to JSON, XML or update it to Google sheet.

## Library

### Install

```
go get -u github.com/ledongthuc/hemnetparser
```

### Usage

#### Parse data to JSON/XML:

```
func main() {
	output, err := hemnetparser.Parse("https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891")
	if err != nil {
		panic(err)
	}

	fmt.Println(output.ToJSON())
	fmt.Println(output.ToXML())
}
```

#### Parse and update to google sheet:

```
type Configuration struct {
}
	googlesheet.Process(googlesheet.Configuration{
	GoogleCredentialPath: "credentials.json" ,
	SpreadsheetID      : "1S0Xzpl_a3SFGnkFWCw7M1t_NXa2VZV42Lxh70G9PTlo" ,

})
```

You can check full configuration at: https://github.com/ledongthuc/hemnetparser/blob/master/googlesheet/sheet.go#L32

## Commandline

### Install

```
go get -u github.com/ledongthuc/hemnetparser/cmd/hemnet-parser
```

### Usage

#### Parse data to JSON

```
hemnet-parser --format=json https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891
```

#### Parse data to XML

```
hemnet-parser --format=xml https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891
```

#### Parse data and update to Google sheet

```
hemnet-parser --format=googlesheet --credential=credentials.json --sheet-id=1S0Xzpl_a3SFGnkFWCw7M1t_NXa2VZV42Lxh70G9PTlo --skipheader=1``
```

credential: Turn on Google Sheet APIs and download credentials.json from https://developers.google.com/sheets/api/quickstart/go

sheet-id: https://developers.google.com/sheets/api/guides/concepts

Before run:
<img width="900" alt="Before" src="https://user-images.githubusercontent.com/1828895/77475363-037d4300-6e19-11ea-9633-03217b3437c0.png">

After run:
<img width="900" alt="After" src="https://user-images.githubusercontent.com/1828895/77475382-0841f700-6e19-11ea-9f48-12f84a1e6bd6.png">

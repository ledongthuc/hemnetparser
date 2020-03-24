[![CircleCI](https://circleci.com/gh/ledongthuc/hemnetparser.svg?style=svg)](https://circleci.com/gh/ledongthuc/hemnetparser)

# Hemnet Parser

Library and a simple commandline to parse information from Hemnet URL to JSON, XML or update it to Google sheet.

## Library

### Install

```
go get -u github.com/ledongthuc/hemnetparser
```

### Usage

Parse data to JSON/XML:

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

Parse and update to google sheet:

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

Parse data to JSON

```
hemnet-parser --format=json https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891
```

Parse data to XML

```
hemnet-parser --format=xml https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891
```

Parse data and update to Google sheet

```
hemnet-parser --format=googlesheet --credential=credentials.json --sheet-id=1S0Xzpl_a3SFGnkFWCw7M1t_NXa2VZV42Lxh70G9PTlo --skipheader=1``
```

credential: Turn on Google Sheet APIs and download credentials.json from https://developers.google.com/sheets/api/quickstart/go

sheet-id: https://developers.google.com/sheets/api/guides/concepts

# goFetchNasaEpicImage

Fetching images from the Nasa EPIC Api

## How To

1. Clone or download the repository
2. Use a terminal and the commands below

Create crossplatform binaries

```bash
sh install.sh
```

Type in your terminal (first binary, then static):

```bash
nasa-epic-downloader --target ~/nasa/images/ --date 2023-03-14
```

```bash
go run cmd/main.go --target ~/nasa/images/ --date 2023-03-14
```

## To Do's

* Fix 'Failed to get latest date: parsing time "2024-06-11 00:31:45" as "20060102": cannot parse "-06-11 00:31:45" as "01"
* Fix or test downloaded images, which seems to be in a wrong format?

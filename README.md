# **gkvDB**
Transactional key-value store DB in Go that sits on top of LevelDB.

## Installation
### 1 - Executable
Go to Releases and download the latest binary. Then run it locally:
```sh
./gkvDB
```

### 2 - go get
```sh
go get -u github.com/gusandrioli/gkvDB
```
### 3 - Run it locally
```sh
git@github.com:gusandrioli/gkvDB.git
go mod tidy
go build
./gkvDB
```

## Usage/Commands
| Command  | Description                                      | Args                 |
|----------|--------------------------------------------------|----------------------|
| BEGIN    | Begins a transaction                             | -                    |
| COMMIT   | Commits a transaction                            | -                    |
| COUNT    | Retrieves the number of records/databases stored | DATABASES/RECORDS    |
| DELETE   | Deletes a record based on a key                  | DATABASE/RECORD + KEY|
| END      | Ends a transaction                               | -                    |
| EXIT     | Exits the console                                | -                    |
| GET      | Gets a record based on a key                     | KEY                  |
| LIST     | Lists all databases/records stored               | DATABASES/RECORDS    |
| NEW      | Creates new database                             | DATABASE + DB_NAME   |
| ROLLBACK | Rolls back a transaction                         |                      |
| SET      | Sets a key to a certain value                    | KEY + VALUE          |
| USE      | Use a specific database                          | DB_NAME              |


## Bugs
Bugs or suggestions? Open an issue [here](https://github.com/gusandrioli/gkvDB/issues/new).

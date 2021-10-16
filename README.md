# **gkvDB**
Transactional key-value store DB in Go that sits on top of LevelDB.

## Installation
### 1 - Download package
```sh
go install github.com/gusandrioli/gkvDB
```

### 2 - Executable
Go to Releases and download the latest binary. Then run it locally:
```sh
./gkvDB
```

### 3 - Run it locally
```sh
git@github.com:gusandrioli/gkvDB.git
go mod tidy
go install
gkvDB
```

## Quick Start
1. Open the gkvDB console:
```sh
gkvDB
```

2. Set a value:
```sh
gkvDB >>> SET Name Joe
```

3. List all records
```sh
gkvDB >>> LIST RECORDS
Name: Joe
```

## Usage/Commands
| Command  | Description                                      | Args                           |
|----------|--------------------------------------------------|--------------------------------|
| BEGIN    | Begins a transaction                             | -                              |
| COMMIT   | Commits a transaction                            | -                              |
| COUNT    | Retrieves the number of records/databases stored | DATABASES/RECORDS              |
| DELETE   | Deletes a record based on a key                  | DATABASE/RECORD + KEY          |
| END      | Ends a transaction                               | -                              |
| EXIT     | Exits the console                                | -                              |
| GET      | Gets a record based on a key                     | KEY                            |
| LIST     | Lists all databases/records/transactions stored  | DATABASES/RECORDS/TRANSACTIONS |
| NEW      | Creates new database                             | DATABASE + DB_NAME             |
| ROLLBACK | Rolls back a transaction                         |                                |
| SET      | Sets a key to a certain value                    | KEY + VALUE                    |
| USE      | Uses a specific database                         | DB_NAME                        |

## Bugs
Bugs or suggestions? Open an issue [here](https://github.com/gusandrioli/gkvDB/issues/new).

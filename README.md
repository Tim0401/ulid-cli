# ulid-cli

## Motivation

Interconversion between ULID strings and [16]bytes.  

## Install

```
go install github.com/Tim0401/ulid-cli/cmd/ulid-cli@{version}
```

## Usage

string->byte
ex) 01FM73V2R6A8G3T20KQ9SQDM6D
```
ulid-cli -sb ${ULID}
```

byte->string  
ex) 0x17d0e3d8b0652203d0813ba7376d0cd
```
ulid-cli -bs ${ULID}
```

## TODO

- テストを書く
- ロジックのmain.goからの分離

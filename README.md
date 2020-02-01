# lambdadg

# Build 
```
go get github.com/LambdaIM/lambda-debugtool
cd $GOPATH/src/github.com/LambdaIM/lambda-debugtool
make
```

# Run

### show real pubkey of validator pubkey
```
./lambdadg show --pubkey lambdavalconspub1zcjduepq77q44pfk5aucvzmrkxzcrg6fsf4e9urru2ln5xs4jxm6xj494rgqj5dd9j --bech pub
```
output 
```
ACD35F73BD94C9AF03F88C34D2D139498C6D28D0
```
### repair keys db 
ERROR: couldn't create db: Error initializing DB: leveldb: manifest corrupted (field 'comparer'): missing [file=MANIFEST-00000xx]
```
./lambdadg repair 
```
output 
```bash
repair successful, keys db: /Users/xxx/.lambdacli/keys/keys.db
```

# Binary

### linux

```
wget https://github.com/LambdaIM/lambda-debugtool/releases/download/v0.0.2/lambda-debugtool.tar.gz
```

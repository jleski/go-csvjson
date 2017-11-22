# CSV-JSON

Simple GO HTTP server with TLS encryption for parsing input query string and outputting it to CSV or JSON data. Just a simple experiment with Golang :)

## Installation

This server supports HTTPS / TLS only so make sure you create SSL private key and certificate first:

```sh
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## Usage

```sh
# Convert POST parameters to JSON (/csv2json)
$ curl -k -XPOST "https://localhost:8443/csv2json?key1=value1&key2=value2&key3=value3"
{"key1":["value1"],"key2":["value2"],"key3":["value3"]}

# Convert POST parameters to CSV (/json2csv)
$ curl -k -XPOST "https://localhost:8443/json2csv?key1=value1&key2=value2&key3=value3"
key2,key3,key1
value2,value3,value1
```


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

2017-11-22 Initial commit

## Credits

Jaakko Leskinen

## License

MIT License
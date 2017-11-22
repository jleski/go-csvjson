# CSV-JSON

Simple GO HTTP server for parsing input query string and outputting it to CSV or JSON data.

## Installation

This server supports HTTPS / TLS only so make sure you create SSL private key and certificate first:

```sh
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## Usage



## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

TODO: Write history

## Credits

TODO: Write credits

## License

TODO: Write license
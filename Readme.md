# X-role

XRole is an utility to assume aws roles from command line and aws profiles

## Installation

```
go get -u github.com/cthulhu/x-role
```

## Usage

List assumable profiles:

```bash
x-role -l
```

Assume profile production:

```bash
x-role production
```

Assume profile production with verbose info:

```bash
x-role -i production
```

Cleanup identity:

```bash
x-role -u
```

Display version info
```bash
x-role -v
```

Display help
```bash
x-role -v
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
# Tarzan - Compress files

*Note: This is a pre-release version of this software.*

## What is Tarzan?

Tarzan is a basic file compressor built using Go.
Tarzan appends the `.tarzan` extension to files when compressing and removes the `.tarzan` extension when decompressing.

Tarzan is still in development and so, currently, it cannot reduce file sizes.

## Usage
```
tarzan --help/-h | --decompress/d <path> | <path>
```
`<path>` is compressed if no options are used

### Examples  
- `tarzan file.txt` will compress `file.txt` and create a new compressed file `file.txt.tarzan` in the same directory as `file.txt`

- `tarzan -d file.txt.tarzan` will decompress `file.txt.tarzan` and recreate the original file in the same directory as `file.txt.tarzan`

- `tarzan --help` will print a short help message


## Installing

You can grab an executable from [releases](https://github.com/quackduck/tarzan/releases) or follow the install steps below.

### macOS

```sh
brew install quackduck/tap/tarzan
```
or
```sh
curl -sSL https://github.com/quackduck/tarzan/blob/master/dist/tarzan_darwin_amd64/tarzan\?raw=true -o /usr/local/bin/tarzan # GNU/linux users might need to use sudo
chmod +x /usr/local/bin/tarzan
```
### Other Unix-like OS
```sh
brew install quackduck/tap/tarzan
```
or
```sh
curl -sSL https://github.com/quackduck/tarzan/blob/master/dist/tarzan_linux_amd64/tarzan\?raw=true -o /usr/local/bin/tarzan # you might need to use sudo
chmod +x /usr/local/bin/tarzan
```
*Note: you might need to change the architecture/os in the url above. The other choice for a Unix OS is `freebsd`. Other architectures available are `arm64`, `arm_6` and `386`*
### Windows install
Download the latest [release](https://github.com/quackduck/tarzan/releases) 


## Uninstalling
Tarzan can be uninstalled completely just by deleting the executable or by following the steps below

### macOS or other Unix-like OS
```sh
rm /usr/local/bin/uniclip
```

### Windows
Just delete the executable

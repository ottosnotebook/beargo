# BearGo  

Author blog content in Bear. Generate a static site using Hugo.
 
![](./assets/beargo.png)
  
## About
BearGo is a tool that allows you to create [Hugo](https://gohugo.io/) blog content in the [Bear](http://www.bear-writer.com/) notes app. 


  
## Install
You need a valid [Go environment](https://golang.org/doc/install) before installing. Once Go is installed: 
```
go get github.com/ottosnotebook/beargo
```

`cd` into the `beargo` repository in your `GOPATH` and run
  
```
go install -v github.com/ottosnotebook/beargo
```  
  
## Usage

BearGo's root command, which converts and builds your Bear notes, can be run by calling `beargo` in your [Hugo project's](https://gohugo.io/getting-started/directory-structure/) directory.
```
$ beargo
```  
  
All commands have help text, just run `--help` after any command: 

```
$ beargo --help
Author blog content in Bear. Generate a static site using Hugo

Usage:
  beargo [flags]
  beargo [command]

Available Commands:
  help        Help about any command
  watch       Watch for changes to Bear notes, regenerate the site using Hugo, and serve the contents

Flags:
      --bear-db string           Path to Bear SQLite database
      --content-dir string       Name of Hugo content directory (default "content")
      --content-section string   Name of content section directory (default "post")
  -h, --help                     help for beargo
      --hugo-exec                Call the version of Hugo installed on the host machine rather than use the latest version of Hugo as a library

Use "beargo [command] --help" for more information about a command.
```  

## Known Issues  

| Issue      | Status |
|-------------------|--------|
| Title generation |   ✅   |
| Filename generation |   ✅   |
| Frontmatter support |   ❌   |
| Image support    |   ❌   |
| Local file support |   ❌   |
| Note year parsing |   ❌   |  
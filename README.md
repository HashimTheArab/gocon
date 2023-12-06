# GoCon
GoCon is the simplest way to store and load a config in your Go program.

## Installation

To use GoCon in your Go project, you can install it by using `go get`:

```bash
go get github.com/HashimTheArab/gocon
```

## Example

```go
type Settings struct {
    Theme string
    DarkMode bool
}

func (Settings) Path() string {
    return "settings.json"
}

func main(){
    // Load the settings
    settings, err := gocon.Load[Settings]()
    if err != nil {
        // If the file fails to load, or it was not found, an error will be returned.
        panic(err)
    }

    // Save settings
    gocon.Save(Settings{Theme: "blue", DarkMode: true})
}
```

## Explanation

Define a config struct, you can define multiple and each one will be saved to a separate file.

1. **Define Your Settings Struct:**<br>
In your settings.go file, define a struct representing your application settings. Implement the Path method to specify the path where the configuration file will be stored.
```go
type Settings struct {
    Theme string
    DarkMode bool
}

func (Settings) Path() string {
    return "settings.json"
}
```

2. **Load and save settings:**<br>
After creating your settings struct, you can simply use it. 

You can save a config:
```go
gocon.Save(Settings{Theme: "blue", DarkMode: true})
```

And then you can load it in the future:
```go
settings, _ := gocon.Load[Settings]()
```

You can have multiple configs / settings, each struct will be saved to its own path.

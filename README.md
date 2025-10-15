
# PipeSim

**PipeSim** is a command-line tool for simulating CI/CD pipelines locally using YAML configuration files. 
It allows developers to test pipeline steps, validate structure, and optionally perform dry runs — all without pushing to a remote CI/CD system.

## 📦 Installation
```
git clone https://github.com/HugoUlm/pipesim.git
cd pipesim
```

## 🚀 Usage
Since **PipeSim** isn't published yet, you'll have to run it locally.
```
go build -o pipesim
./pipesim pipesim run --file path/to/your/file --dry-run
```

### Available commands
```
./pipesim pipesim --help
```
| Flag | Short flag | Description |
| ------- | ----- | ----------- |
| --dry-run | No short flag | Print steps without executing |
| --file | -f | Path to your yaml file (required) |
| --project | -p | Path to your project |
| --use-cache | No short flag | Remove downloaded packages after run |

## 🤝 Contributing
Feel free to open issues or submit pull requests to improve functionality or add new pipeline simulation features.

## 📄 License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/HugoUlm/pipesim/blob/main/LICENSE) file for details.

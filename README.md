# System Status Monitor CLI
The System Status Monitor CLI is a command-line tool written in Go that provides real-time monitoring and display of various system status information. This tool offers insights into memory usage, CPU performance, disk space, operating system details, and more.

## Features
- View system status information such as memory usage, CPU statistics, disk space, and more.
- Choose from a range of available status options using an interactive menu.
- Display status information in human-readable formats for easy understanding.
- Utilize color-coded output to highlight different types of information.
- Access system status information using an intuitive and user-friendly interface.

## Installation
1. Make sure you have Go installed on your system. If not, you can download and install it from the official Go website.

2. Clone this repository to your local machine:

```bash
git clone https://github.com/sz9751210/go-system-status.git
```

3. Navigate to the cloned directory:
```bash
cd go-system-status
```

4. Run the following command to build and execute the CLI tool:
```bash
make run
```

## Usage

Upon running the CLI tool, you'll be presented with an interactive menu that allows you to select the type of system status information you want to view. Follow the prompts and choose the desired option by typing the corresponding number. You can exit the tool by selecting the "Quit" option or pressing q at any time.

## Acknowledgments

This project was inspired by the need to quickly access and monitor system status information in a user-friendly way. Special thanks to the open-source community for providing the libraries that make this tool possible.

## License
This project is licensed under the [MIT License](LICENSE).

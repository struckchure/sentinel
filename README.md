# Sentinel API Gateway

![Go Version](https://img.shields.io/badge/go-1.24.1-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/struckchure/sentinel)](https://goreportcard.com/report/github.com/struckchure/sentinel)

Sentinel is a lightweight, configuration-driven API gateway and reverse proxy built with Go. It is designed to be a simple, performant, and extensible solution for managing and securing backend services through a single entry point.

## ✨ Key Features

-   **Configuration-Driven**: Define all routing, load balancing, and plugin behavior in a single YAML or JSON file.
-   **Multiple Load Balancing Strategies**: Choose from various algorithms like `round-robin`, `random`, `least-connections`, `ip-hash`, and `weighted` to distribute traffic effectively.
-   **Extensible Plugin Architecture**: Enhance gateway functionality with custom logic. The configuration supports plugins for concerns like authentication and rate-limiting.
-   **Command-Line Interface**: Manage the gateway with a clean and simple CLI, built using Cobra.
-   **Schema Generation**: Automatically generate a JSON schema for your configuration file to enable validation and autocompletion in your editor.

## 🛠️ Technologies Used

| Technology                                               | Description                                        |
| -------------------------------------------------------- | -------------------------------------------------- |
| [**Go**](https://golang.org/)                            | The core programming language used for development. |
| [**Cobra**](https://github.com/spf13/cobra)              | A powerful library for creating modern CLI applications. |
| [**Echo**](https://github.com/labstack/echo)             | A high-performance, minimalist Go web framework.   |

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/struckchure/sentinel.git
    cd sentinel
    ```

2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Build the application**:
    ```bash
    go build -o sentinel ./cmd
    ```
    This will create a `sentinel` executable in the root directory.

## ⚙️ Usage

Sentinel is operated via its command-line interface and configured using a YAML file.

### 1. Configuration

Create a `sentinel.yaml` file to define your gateway's behavior. The gateway routes requests based on URL patterns to a set of backend services, applying load balancing and plugins as configured.

Here is an example configuration from `example/sentinel.yaml`:

```yaml
# sentinel.yaml
host: localhost
port: 3000
backends:
  - load_balancer: round-robin
    methods: ["*"]
    patterns:
      - /todo/
      - /todo/*
    plugins:
      - name: rate-limiter
        config:
          limit: 20s
          expires: 3m
      - name: auth-n
        config:
          alg: RS256
          jwk_url: string
          jwt_secret: string
          propagate_claims:
            - from: sub
              to: X-User
    services:
      - url: localhost:8010
        weight: 1
      - url: localhost:8020
        weight: 1
      - url: localhost:8030
        weight: 1
```

### 2. Running the Gateway

Start the gateway using the `run` command, specifying your configuration file with the `-c` flag.

```bash
./sentinel run -c example/sentinel.yaml
```

The gateway will now listen for requests on `localhost:3000` and forward traffic matching the `/todo/*` pattern to the backend services `localhost:8010`, `localhost:8020`, and `localhost:8030` using a round-robin strategy.

### 3. CLI Commands

Sentinel provides several commands to help with development and management.

-   **Run the Gateway**
    ```bash
    ./sentinel run [flags]
    ```
    Flags:
    - `-c`, `--config`: Path to the sentinel config file (default: `sentinel.yaml`).

-   **Generate Config Schema**
    Generate a JSON schema to validate your configuration files.
    ```bash
    ./sentinel schema [flags]
    ```
    Flags:
    - `-i`, `--indentation`: Indentation size for the output (default: `2`).
    - `-s`, `--save`: Save the schema to a file instead of printing to stdout.
    - `-o`, `--output`: The output file path (default: `sentinel.schema.json`).

-   **Check Version**
    Print the current version, commit, and build date.
    ```bash
    ./sentinel version
    ```

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  **Fork the Project**
2.  **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4.  **Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **Open a Pull Request**

## 📄 License

This project does not have a specified license. Please refer to the `LICENSE` file for more details.

## 📬 Contact

Your Name - [@struckchure](https://twitter.com/struckchure)

Project Link: [https://github.com/struckchure/sentinel](https://github.com/struckchure/sentinel)

<br/>
<p align="center">
  <a href="https://www.npmjs.com/package/dokugen">
    <img src="https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen" alt="Dokugen">
  </a>
</p>
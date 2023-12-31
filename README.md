# Release Hunter

Release Hunter is a Go-based command-line tool that helps you discover latest GitHub releases for any software. You can use it to search for GitHub repositories by keyword and access direct download URLs for release assets.

![Release Hunter LOGO](/extras/rh_logo.png)

## Command-Line Flags and Aliases

- `-help, -h`: Show usage and examples.

- `-version, -v`: Show the version of the CLI.

- `-token, -t <token>`: GitHub personal access token.

- `-repo, -r <repository>`: GitHub repository in the format `user/repo`.

- `-find, -f <keyword>`: Search GitHub repositories by keyword to find correct `user/repo`.

- `-keyword, -k <keyword>`: Filter links by an optional keyword, can be used with the -f or -r flags.

## Prerequisites

Before using Release Hunter, you need to set up a GitHub personal access token. You can create the token and set it either as an environment variable or use it as a flag when running the tool.

### Creating a GitHub Personal Access Token

1. Go to [GitHub Settings → New Token](https://github.com/settings/tokens/new) in your GitHub account.
2. Set "name", "expiration date" and grant permissions: "repo - public_repo" and "user - read:user".
3. Use the token:
- as an environment variable:
```sh
export GITHUB_TOKEN=<your-token>
```
- or as a flag:
```sh
rh -t <your-token>
```

## Installation

All available versions for different operating systems can be found in releases [here](https://github.com/sv222/release-hunter/releases)

To install Release Hunter, use the following commands (linux_amd64 example):

```sh
curl -L -o rh https://github.com/sv222/release-hunter/releases/download/v0.1.0/release_hunter_0.1.0_linux_amd64 && chmod +x rh && sudo mv rh /usr/local/bin/rh
```

## Usage

### Search GitHub Repositories

Search for GitHub repositories using a keyword:

```sh
rh -f <keyword>
```

For example, to search for repositories related to Podman:

```sh
rh -f podman
```

Filter a search result by the keyword "desktop":

```sh
rh -find podman -k desktop
```
Copy <user/repo> for found repo "containers/podman-desktop".

### Get the Latest Release for a Repository

Retrieve the latest release for a specific GitHub repository:

```sh
rh -repo <user/repo>
```

For example, to get the latest release for "podman-desktop":

```sh
rh -repo containers/podman-desktop
```

Filter the resulting links using the keyword "arm":

```sh
rh -repo containers/podman-desktop -k arm
```
![How to use GIF](/extras/rh_usage.gif)

## Docker Usage

To run the Release Hunter CLI within a Docker container, follow these steps:

### Step 1: Build the Docker Image

1. Make sure you have Docker installed on your system.

2. Clone a project locally:

    ```sh
    git clone https://github.com/sv222/release-hunter.git
    ```

3. Open a terminal and navigate to the root directory of the project where the Dockerfile is located.

4. Build the Docker image using the following command:

```sh
docker build -t rh .
```

### Step 2: Run the Docker Container

Once the Docker image is built, you can run the Release Hunter CLI inside a Docker container with the desired flags and arguments.

Example command to run the container with flags and arguments:

```sh
docker run --rm rh -t <your-github-token> -r <user/repo>
```

## Contributing

We welcome contributions from the community. If you have ideas, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

<h1>GOPPORTUNITIES</h1>
<p align="left">
	<!-- Shields.io badges disabled, using skill icons. --></p>
<p align="left">Built with the tools and technologies:</p>
<p align="left">
	<a href="https://skillicons.dev">
		<img src="https://skillicons.dev/icons?i=go,md,sqlite">
	</a></p>
</div>
<br clear="right">

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

<code>❯ GOPPORTUNITIES is an API in GO to post job opportunities</code>

---

##  Features

<code>❯ CRUD Operations for posting a job opportunitie</code>

---

##  Project Structure

```sh
└── gopportunities/
    ├── README.md
    ├── config
    │   ├── config.go
    │   ├── logger.go
    │   └── sqlite.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── go.mod
    ├── go.sum
    ├── handler
    │   ├── createOpening.go
    │   ├── deleteOpening.go
    │   ├── handler.go
    │   ├── listOpenings.go
    │   ├── request.go
    │   ├── response.go
    │   ├── showOpening.go
    │   └── updateOpening.go
    ├── main.go
    ├── makefile
    ├── router
    │   ├── router.go
    │   └── routes.go
    └── schemas
        └── opening.go
```


###  Project Index
<details open>
	<summary><b><code>GOPPORTUNITIES/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/makefile'>makefile</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/main.go'>main.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/go.mod'>go.mod</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/go.sum'>go.sum</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- schemas Submodule -->
		<summary><b>schemas</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/schemas/opening.go'>opening.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- config Submodule -->
		<summary><b>config</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/config/config.go'>config.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/config/logger.go'>logger.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/config/sqlite.go'>sqlite.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- router Submodule -->
		<summary><b>router</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/router/router.go'>router.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/router/routes.go'>routes.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- handler Submodule -->
		<summary><b>handler</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/createOpening.go'>createOpening.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/showOpening.go'>showOpening.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/response.go'>response.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/updateOpening.go'>updateOpening.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/handler.go'>handler.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/listOpenings.go'>listOpenings.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/deleteOpening.go'>deleteOpening.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/gopportunities/blob/master/handler/request.go'>request.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with gopportunities, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules


###  Installation

Install gopportunities using one of the following methods:

**Build from source:**

1. Clone the gopportunities repository:
```sh
❯ git clone https://github.com/diegobbrito/gopportunities
```

2. Navigate to the project directory:
```sh
❯ cd gopportunities
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```




###  Usage
Run gopportunities using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```

---

##  Contributing

- **💬 [Join the Discussions](https://github.com/diegobbrito/gopportunities/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/diegobbrito/gopportunities/issues)**: Submit bugs found or log feature requests for the `gopportunities` project.
- **💡 [Submit Pull Requests](https://github.com/diegobbrito/gopportunities/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/diegobbrito/gopportunities
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/diegobbrito/gopportunities/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=diegobbrito/gopportunities">
   </a>
</p>
</details>

---

##  License

This project is protected under the MIT License. For more details, refer to the [MIT License](https://choosealicense.com/licenses/mit/) file.

---

##  Acknowledgments

-[Youtube: Gopportunities](https://www.youtube.com/watch?v=wyEYpX5U4Vg)

---
#  Development Containers
**Development container** is a container in which a user can develop an application, read here.


## Pre-requisites
- Docker installed and has at least 4GB of free resources
- GoLand 2023.2 is installed
- You have a project, that contains .devcontainer folder with devcontainer.json in the root
- You have access to GitHub
- Your Git version should be 2.25 or later

## How to run a development container
1. Clone this repository to your local filesystem.
2. Open the `goDevContainer` directory in GoLand 2023.2 or higher.
3. In the `Project` tool window, open the `.devcontainer` folder and double-click the  `devcontainer.json` file.
4. Click the gutter icon and select `Create Dev Container (mount sources)`.
5. After the development container is created, select `GoLand` from the drop-down list of IDE instances (for example, `GoLand 2023.2 (232.8660.185) | download latest`).
6. Click the `Continue` button.  The IDE will be downloaded as a backend and the Client is opened. The Client is connected to the backend.
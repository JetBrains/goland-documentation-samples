#  Development Containers
**Development container** is a container in which a user can develop an application, read here.


## Pre-requisites
- The up-to-date Docker version should be installed with at least 4GB of free resources
- GoLand 2023.2 or later
- You have a project, that contains `.devcontainer` folder with `devcontainer.json` in the root
- You have access to GitHub
- Your Git version should be 2.25 or later

## How to run a development container
1. Ensure that the Docker service is up and running.
2. Clone this repository to your local filesystem.
3. Open the `goDevContainer` directory in GoLand 2023.2 or higher.
4. In the `Project` tool window, open the `.devcontainer` folder and double-click the  `devcontainer.json` file.
5. Click the gutter icon and select `Create Dev Container and Mount Sources`.
6. After the development container is created, select `GoLand` from the drop-down list of IDE instances (for example, `GoLand 2023.2 (232.8660.185) | download latest`).
7. Click the `Continue` button.  The IDE will be downloaded as a backend and the Client is opened. The Client is connected to the backend.
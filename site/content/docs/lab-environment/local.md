## 📌 Overview

To get off the ground running with a local copy of this projects, start with making sure that you have the proper tools instead head of time.

---

## Local Requirements

| Technology | Install Guide                                  |
| ---------- | ---------------------------------------------- |
| Docker     | [install](https://docs.docker.com/get-docker/) |
| Poetry     | [install](https://python-poetry.org/)          |

Please make sure you have correct permissions to execute

---

## Setting up the enviornment

Once you're ready for flight, start by [cloning](https://rogerdudler.github.io/git-guide/) this repository to your workstation.

```bash
git clone https://github.com/cdot65/apstra-edge-ansible.git
```

![git clone](https://raw.githubusercontent.com/cdot65/juniper-apstra-edge-design/main/site/content/assets/images/clone.png)

After changing into the project's directory, create and activate a virtual environment with Poetry

```bash
cd apstra-edge-ansible.git
poetry install
poetry shell
```

![poetry install](https://raw.githubusercontent.com/cdot65/juniper-apstra-edge-design/main/site/content/assets/images/poetry_install.png)

![poetry shell](https://raw.githubusercontent.com/cdot65/juniper-apstra-edge-design/main/site/content/assets/images/poetry_shell.png)

> **Note: If you have EVE-NG running within your home, we have provided a [sample EVE-NG lab file](https://github.com/cdot65/juniper-apstra-edge-design/tree/main/files/eve-ng).**

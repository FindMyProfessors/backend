# FindMyProfessor's Backend

## What is this?

This repository is a centralized overview of all the backend's microservices/cronjobs and also should be the root directory of
all the microservices in your development environment. We
utilize [git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules) which allows us to create a more
containerized development flow.

If you open any of the submodule repositories from here it will send you to the repository with a detached head, in
simple terms, you are viewing the repository from the context of some in commit that happened. This repository should
act as a single source of truth and the afformentioned referenced commit should always be pointing to what is in
production currently, for the purpose of clarity.

## Cloning

Since this repository utilizes git submodules, you must clone the repository recursively use the following command.

```bash
eval `ssh-agent`
ssh-add
git clone --recurse-submodules -j8 git@github.com:FindMyProfessors/backend.git
```

The ssh-agent will save you from typing your passphrase SSH in many times 
The `-j8` flag parallelizes the cloning to save time in the cloning process.


Following the previous command I suggest running the following command inside of the backend repository folder. It will checkout all submdules onto the main branch and pull down any recent changes. 
```bash
bash checkoutall.sh main
```

## Requirements

### Golang 1.18+

We require version 1.18 in the development environment because
of [workspace support](https://go.dev/doc/tutorial/workspaces). Workspaces allow you to seemlessly utilize our
module-based workspace without having to push the changes of a dependency (
like [shared](https://github.com/FindMyProfessors/shared)), and then make use the change on the dependent module.
You can see live-changes to your dependencies from your dependents in your IDE when using go workspaces.

### Something that can run a shell script

Whether you are on Linux, MacOS, [Windows using WSL](https://docs.microsoft.com/en-us/windows/wsl/about),
or [Git Bash (not recommended)](https://gitforwindows.org/), we don't care. If you want to be able to run the bash
scripts you must be on one of these platforms. This is not a must, however it will make your development flow a lot more
seemless. No one wants to remember some long command that you use once in a blue moon.

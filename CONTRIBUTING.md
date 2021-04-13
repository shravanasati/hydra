# Contributing to hydra

👍🎉 First off, thanks for taking the time to contribute! 🎉👍

The following is a set of guidelines for contributing to *hydra*, which is hosted on GitHub. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.


## Project Structure
```
├── .github
|  ├── ISSUE_TEMPLATE           # issue templeates
|  |  ├── bug_report.md
|  |  ├── custom.md
|  |  └── feature_request.md
|  └── workflows                # ci workflow
|     └── go.yml
├── .gitignore
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── README.md
├── assets                      # media assets for readme
|  ├── python_dir.PNG
|  └── python_init.PNG
├── config.go                   # config command code
├── gitignores                  # all gitignores
|  ├── go.gitignore
|  └── python.gitignore
├── go.mod
├── go.sum
├── hydra.go                    # main code of cli
├── init.go                     # init command code
└── licenses                    # all licenses
   ├── APACHE
   ├── BSD
   ├── EPL
   ├── GPL
   ├── MIT
   └── MPL
```

## Setup Development Environment
This section shows how you can setup your development environment to contribute to hydra.

- Fork the repository.
- Clone it using Git (`git clone https://github.com/<YOUR USERNAME>hydra.git`).
- Create a new git branch (`git checkout -b "BRANCH NAME"`).
- Install `commando` module using the command `go get github.com/thatisuday/commando`.
- Make changes.
- Stage and commit (`git add .` and `git commit -m "COMMIT MESSAGE"`).
- Push it your remote repository (`git push`).
- Open a pull request by clicking [here](https://github.com/Shravan-1908/hydra/compare).

## Reporting Issues
If you know a bug in the code or you want to file a feature request, open an issue.
Choose the correct issue template from [here](https://github.com/Shravan-1908/hydra/issues/new/choose).
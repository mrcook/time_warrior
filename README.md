# TimeWarrior

A command line time tracking tool for developers and freelance workers who need to track time worked on their personal and client projects.


## Usage

When you start TimeWarrior (`tw`) for the first time, a data directory will be created in `$HOME/time_warrior`. This is where all your project data files will be stored.

You can view the help by typing `tw -h`, or by prefixing each sub command like so: `tw help done`.


### Start New Timeslip

    $ tw start MyProject.SetupTask

This command will start a new _pending_ timeslip with a _project_ name of `MyProject` and a _task_ name of `SetupTask`. This timeslip will be saved in the `$HOME/time_warrior/.pending` file, and the details printed to the command line:

    $ MyProject.SetupTask | Started: 2017-12-11 13:35 | Worked: 0 seconds | Status: started

There are several points to take note of regarding starting new timeslips:  

- Use ASCII characters for the `Project.Task` names - upper and lower case alphanumeric characters.
- Spaces are **not allowed** - it is recommended you use CamelCase.
- The _task_ name is optional.
- Only one timeslip can be started at a time.


### Pause Timeslip

    $ tw pause

As expected, this will pause a currently running timeslip, and print the details on the command line:

    $ MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 2 minutes | Status: paused


### Resume Timeslip

    $ tw resume

As expected, this will resume a paused timeslip, and print the details on the command line:

    $ MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 2 minutes | Status: started


### Done! Complete Timeslip

    $ tw done "Basic project setup with a nice README"

When you've finished your current task you can mark the slip as done, giving a short description of the work done. You should use either 'single' or "double" quotes.

The details will again be printed on the command line:

    $ MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 9m 23s | Status: completed

A new project data file will be created using the project _name_ you gave (`MyProject`) and saved as `$HOME/time_warrior/my_project.json`. Each timeslip created for this project will be save on a separate line in this file.


### Delete Timeslip

    $ tw delete

If you make a mistake when starting a new timeslip, you can delete it easily with this command.

**Caution! This action can not be undone!**


### NOTES

* Running the program (typing just `tw`) while you have a _pending_ timeslip will show the current details for the that slip.
* `Worked` time format is displayed in several styles: `hours`, `minutes`, `seconds`, along with two abbreviated combinations: `1h 23m` and `10m 14s`.
* All commands have a _short_ alias (except `delete`), as follows:
  - `s`, `start`
  - `p`, `pause`
  - `r`, `resume`
  - `d`, `done`


## Installation

    go get -u -v github.com/mrcook/time_warrior/...

It's recommend to have your `$GOPATH/bin` directory exported in your shell so you can then just type `tw` to get started.


## Contributing

To contribute to the source code or documentation, you should [fork the TimeWarrior GitHub project](https://github.com/mrcook/time_warrior) and clone it to your local machine. Then making a PR (Pull Request) for review.


## Reporting Issues

If you believe you have found a defect in TimeWarrior or its documentation, use the GitHub issue tracker to report the problem to the TimeWarrior maintainers.

When reporting the issue, please provide the version of TimeWarrior in use (`tw --version`).

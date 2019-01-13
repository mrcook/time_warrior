# TimeWarrior

A command line time tracking tool for people who need to track the time they work on their personal and client projects.


## Usage

Simply type `tw` in your terminal to get started. You can view the help by typing `tw -h`, or by prefixing each sub command with `help`, like so: `tw help start`.

By default, running the application will display the help page, however, if you have a _pending_ timeslip in progress then those details will be printed instead:

    MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 22 minutes | Status: started

The `Worked` time format is displayed using `hours`, `minutes`, `seconds`, along with two abbreviated combinations: `1h 23m` and `10m 14s`.

On the very first time you run TimeWarrior on your system, a directory will be created in `$HOME/time_warrior`. This is where all your project data files will be stored - it can be useful to make this into a `git` repository.


### Shortcuts

Common commands have a _short_ alias:

  - `s`, `start`
  - `p`, `pause`
  - `r`, `resume`
  - `d`, `done`


### Start New Timeslip

    $ tw start MyProject.SetupTask

This command will start a new _pending_ timeslip with a _project_ name of `MyProject` and a _task_ name of `SetupTask`. This timeslip will be saved in the `$HOME/time_warrior/.pending` file, and the details printed to the command line:

    MyProject.SetupTask | Started: 2017-12-11 13:35 | Worked: 0 seconds | Status: started

Recommendations for starting new timeslips:

- Use CamelCase for _Project_ and _Task_ names.
- Always including a _Task_ name, this will improve the generated reports.
- Tasks can span several work sessions, so it's okay to use the same _Project.Task_ name multiple times.

There are several points to take note of regarding starting new timeslips:  

- Use only ASCII characters for the `Project.Task` names - upper and lower case alphanumeric characters.
- Spaces are **not allowed**.
- The _task_ name is optional, but recommended.
- Only one timeslip can be started at a time.
 

### Pause Timeslip

    $ tw pause

This will pause the currently running timeslip, and print the details to the terminal:

    MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 2 minutes | Status: paused


### Resume Timeslip

    $ tw resume

This will resume a paused timeslip and print the details to the terminal:

    MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 2 minutes | Status: started


### Done! Complete Timeslip

    $ tw done "Basic project setup with a nice README"

When you've finished your current task you can mark the slip as done, giving a short description of the work made. You should use either 'single' or "double" quotes.

The details will again be printed on the command line:

    MyProject.SetupTask | Started: 2017-12-11 13:37 | Worked: 9m 23s | Status: completed

If this is a new project a data file will be created using the project _name_ you gave (`MyProject`) and saved as `$HOME/time_warrior/my_project.json`. Each timeslip created for this project will be save on a separate line in this file.


### Adjust Timeslip

If you forget to `start`, `pause`, or `resume` your current timeslip, you can use the `adjust` command to add/subtract a time duration to the `worked` time.

Adjustments are made using a simple duration string in the format of a decimal number followed by a single time unit character (e.g. `10m`). Allowed units are `h`, `m`, and `s`, for hours, minutes, and seconds, respectively.

Here's some examples adding worked time:

    $ tw adjust 2h
    $ tw adjust 30m
    $ tw adjust 45s

    $ tw adjust 72m
    $ tw adjust 720s

To subtract time you need to specify the `-n` (negative) flag. Examples:

    $ tw adjust -n 1h
    $ tw adjust -n 20m
    $ tw adjust -n 90s

Here is a full example of adding `1 hour` and `15 minutes` to a running timeslip:

    $ tw
    MyProject.SetupTask | Started: 2017-12-11 18:01 | Worked: 15m | Status: started

    $ tw adjust 75m
    MyProject.SetupTask | Started: 2017-12-11 18:01 | Worked: 1h 30m | Status: started


### Delete Timeslip

**Caution! This action can not be undone!**

    $ tw delete

If you make a mistake when starting a new timeslip, perhaps using an incorrect project name, you can delete it easily with this command.


## Reports

Reports can be generated using the `report` command, with the listing printed to the terminal:

    $ tw report
       3h  01m : Project A
       1h  22m : Project B
    ===========
       4h  23m

The time worked is specified in the `h` and `m` columns (hours, minutes), which is followed by the project name. At the end of each listing, the total time worked is displayed.


### Projects Overview

When no project name has been given, all projects found in the `$HOME/time_warrior` directory will be included, and displayed in alphabetical order.

    $ tw report
       3h  01m : Project A
      11h  22m : Project B
      70h  52m : TimeWarrior
    ===========
      85h  05m


### Tasks Overview

When specifying a project name, all tasks within that project will be included, and displayed in alphabetical order.

    $ tw report TimeWarrior
    Project Name: TimeWarrior
    
    Task List
      10h   5m : .
       8h  19m : AdjustCommand
       .
       .
       .
       4h  11m : Timeslip
    ===========
      70h  52m

Note the `.` task name above. If a timeslip is recorded without a task name being used then these will be grouped here.

To have a better breakdown in your reports it's recommended that task names be used when starting new timeslips.


### Report Time Periods

Although it is interesting to know the total time worked on a project, it's more common for a specific period of time, such as _today_ or _last month_, to be requested.

The `report` command takes a `-p` flag to specify a time period:

    $ tw report -p w TimeWarrior
    Project Name: TimeWarrior
    Time Period:  This Week (Jan 7, 2019 to Jan 13, 2019)
    
    Task List
       1h  32m : Documentation
      10h  20m : Reports
    ===========
      11h  52m

The following time periods are available:

* `t` - Today
* `w` - This Week
* `m` - This Month
* `y` - This Year

* `1d` - Yesterday
* `1w` - Last Week
* `1m` - Last Month
* `1y` - Last Year

A time period of `1d` can be described as _one day previous_, otherwise known as _yesterday_, and `1m` would be _one month previous_ (_last month_).

I've tried to follow the same pattern as with the _adjust_ command, hopefully this nomenclature is clear.


## Installation

    $ go get -u -v github.com/mrcook/time_warrior/...

To install the app after manually cloning the repository you must first `cd` into the `tw` directory:

    $ cd $GOPATH/src/github.com/mrcook/time_warrior/tw
    $ go install

If you've added the `$GOPATH/bin` directory to your `$PATH`, you can then just type `tw` to get started.


## Contributing

To contribute to the source code or documentation, you should [fork the TimeWarrior GitHub project](https://github.com/mrcook/time_warrior) and clone it to your local machine. Then making a PR (Pull Request) for review.


## Reporting Issues

If you believe you have found a defect in TimeWarrior or its documentation, use the GitHub issue tracker to report the problem to the TimeWarrior maintainers.

When reporting the issue, please provide the version of TimeWarrior in use (`tw --version`).

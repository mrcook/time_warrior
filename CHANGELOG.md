# TimeWarrior changes

## HEAD


## 1.4.2 (2026-01-24)

- Bugfix: `done` cmd correctly calculates worked time for resumed timeslips.
- Update `go.mod` dependencies: `uuid` and `cobra`.


## 1.4.1 (2025-12-27)

- Fix `TotalTimeWorked` calculation on completing a paused timeslip.


## 1.4.0 (2025-11-12)

- Add a "resumed" status for when a paused timeslip is resumed.
- Resumed timeslips show the last modified timestamp - like when paused.
- Refactor and improve the timeslip tests.


## 1.3.3 (2023-08-12)

- Bugfix: when resuming a non-existent timeslip, a better error is displayed.
- Updates dependencies to their latest versions.


## 1.3.2 (2020-11-09)

- display a message (not error) when pausing a non-existing timeslip.


## 1.3.1 (2020-02-22)

- Fixes a bug in reports related to paused timeslips.


## 1.3.0 (2020-02-22)

- Append the timestamp to the `tw` output for when a timeslip was paused.
- Reports now include any worked time for a pending timeslip.
- Update app dependencies
- Use `go mod`


## 1.2.1 (2019-05-12)

- Fix: do not change `modified` time when completing **paused** timeslips.
- Add missing entries to CHANGELOG.


## 1.2.0 (2019-01-13)

- Add basic reporting.
- Fix `toSnakeCase()` to split on numbers correctly.
- More refactoring.


## 1.1.1 (2018-11-01)

- Small fix to `adjust` command to handle missing time unit.
- Small refactor of `manager` package.
- Add some basic code documentation.


## 1.1.0 (2017-12-25)

- `adjust` command will now automatically pause/resume a running timeslip.
- Bugfix: `adjust` command does not allow the `worked` time to be less than `0`.


## 1.0.0 (2017-12-15)

First release of my TimeWarrior time tracker tool, ported from Ruby to Go.

The following commands are available:

- `start`:  Start a new timeslip
- `pause`:  Pause a started timeslip
- `resume`: Resume a paused timeslip
- `adjust`: Adjust +/- the time worked on a timeslip
- `delete`: Delete an in progress timeslip
- `done`:   Mark current timeslip as completed

# TimeWarrior changes

## HEAD


## 1.1.0 (2017-12-25)

- `adjust` command will now automatically pause/resume a running timeslip.
- Bugfix: `adjust` command does not allow the `worked` time to be less than `0`.


## 1.0.0 (2017-12-15)

First release of my TimeWarrior time tracker tool, ported from Ruby to Go.

The following commands are avalable:

- `start`:  Start a new timeslip
- `pause`:  Pause a started timeslip
- `resume`: Resume a paused timeslip
- `adjust`: Adjust +/- the time worked on a timeslip
- `delete`: Delete an in progress timeslip
- `done`:   Mark current timeslip as completed


// Robot simulator, using channel to communicate between the room and the
// robots, and report final status to the test program
package robot

const testVersion = 3

const (
	N Dir = iota
	E
	S
	W
)

func stepforward(p Pos, d Dir) Pos {
	switch d {
	case N:
		p.Northing++
	case S:
		p.Northing--
	case W:
		p.Easting--
	case E:
		p.Easting++
	}
	return p
}

func turnright(d Dir) Dir {
	switch d {
	case N:
		return E
	case S:
		return W
	case W:
		return N
	case E:
		return S
	default:
		return d
	}
}

func turnleft(d Dir) Dir {
	switch d {
	case N:
		return W
	case S:
		return E
	case W:
		return S
	case E:
		return N
	default:
		return d
	}
}

// Step1Robot advance one step based on its current direction
func Advance() {
	p := Pos{RU(Step1Robot.X), RU(Step1Robot.Y)}
	p = stepforward(p, Step1Robot.Dir)
	Step1Robot.X = int(p.Easting)
	Step1Robot.Y = int(p.Northing)
}

// Step1Robot turning right
func Right() {
	Step1Robot.Dir = turnright(Step1Robot.Dir)
}

// Step1Robot turning left
func Left() {
	Step1Robot.Dir = turnleft(Step1Robot.Dir)
}

// Print the directoin
func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case W:
		return "W"
	case E:
		return "E"
	default:
		return "Invalid direction"
	}
}

type Action byte

// Step2
// Test program send commands to robot
// Robot accept commands and inform Room of actions it is attempting.
// When it senses the command channel closing, it shut down itself.
func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)

}

// Step2
// The room interpret the physical consequences of the robot actions.
// When it senses the robot shutting down, it sends a final report back
// to the test program, telling the robot's final position and direction.
func Room(extent Rect, robot Step2Robot, act chan Action,
	rep chan Step2Robot) {
	for {
		select {
		case a, ok := <-act:
			if !ok {
				act = nil
			} else {
				switch a {
				case 'A':
					attempt := stepforward(robot.Pos, robot.Dir)
					if IsInsideRoom(extent, attempt) {
						robot.Pos = attempt
					} // else traverse into the wall, not moving
				case 'L':
					robot.Dir = turnleft(robot.Dir)
				case 'R':
					robot.Dir = turnright(robot.Dir)
				}
			}
		}
		if act == nil {
			// send to rep
			rep <- robot
			break
		}
	}
}

// Judege whether the point is inside of Rect
// on the wall is also OK based on the test case
func IsInsideRoom(extent Rect, p Pos) bool {
	// suppose in Rect, min always < max
	if extent.Min.Easting <= p.Easting && p.Easting <= extent.Max.Easting &&
		extent.Min.Northing <= p.Northing && p.Northing <= extent.Max.Northing {
		return true
	}
	return false

}

type Action3 struct {
	a    Action
	name string
}

// Robots run scripts containing bytes of commands
// A log channel allows robots and the room to log messages.
func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() {
		action <- Action3{'D', name}
	}()

	for i := 0; i < len(script); i++ {
		switch script[i] {
		case 'R':
			action <- Action3{'R', name}
		case 'L':
			action <- Action3{'L', name}
		case 'A':
			action <- Action3{'A', name}
		default:
			log <- "An undefined command in a script"
			return // can't break, since it breaks only the inner switch
			// can use tag Loop, break Loop instead
		}
	}
}

// For the final position report sent from StartRobot3, you can return
// the same slice received from the robots channel, just with
// updated positions and directions.
func Room3(extent Rect, robots []Step3Robot, action chan Action3,
	report chan []Step3Robot, log chan string) {
	defer func() {
		report <- robots
	}()

	var step3RobotMap = make(map[string]int) // name, index
	var posMap = make(map[Pos]bool)
	for i, r := range robots {
		if len(r.Name) == 0 {
			log <- "A robot without a name"
			return
		}
		if _, ok := step3RobotMap[r.Name]; !ok {
			step3RobotMap[r.Name] = i
		} else {
			log <- "Duplicate robot names"
			return
		}
		if !IsInsideRoom(extent, r.Pos) {
			log <- "A robot placed outside of the room"
			return
		}
		if !posMap[r.Pos] {
			posMap[r.Pos] = true
		} else {
			log <- "Robots placed at the same place"
			return
		}
	}
	// It should quit + close the chan if the above error shows?

	cnt := 0
	for act := range action {
		if _, ok := step3RobotMap[act.name]; !ok {
			log <- "An action from an unknown robot"
			break // based on the test case
		}
		robot := &robots[step3RobotMap[act.name]]
		switch act.a {
		case 'A':
			attempt := stepforward(robot.Pos, robot.Dir)
			if !IsInsideRoom(extent, attempt) {
				log <- "A robot attempting to advance into a wall"
				continue
			}
			if posMap[attempt] {
				log <- "A robot attempting to advance into another robot"
				continue
			}
			delete(posMap, robot.Pos)
			robot.Pos = attempt
			posMap[attempt] = true

		case 'L':
			robot.Dir = turnleft(robot.Dir)
		case 'R':
			robot.Dir = turnright(robot.Dir)
		case 'D':
			cnt++
		}
		if cnt == len(robots) {
			break
		}
	}

}

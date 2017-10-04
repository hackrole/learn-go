package foo

type Game struct {
	rolls   []int
	current int
}

func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, maxThrowsPerGame)
	return game
}

func (self *Game) Roll(pins int) {
	self.rolls[self.current] = pins
	self.current++
}

func (self *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if self.isStrike(throw) {
			sum += self.strikeBonusFor(throw)
			throw += 1
		} else if self.isSpare(throw) {
			sum += self.spareBonusFor(throw)
			throw += 2
		}
	}
	return sum
}

func (self *Game) isStrike(throw int) bool {
	return self.rolls[throw] == allPins
}

func (self *Game) strikeBonusFor(throw int) int {
	return allPins + self.framePointsAt(throw+1)
}

func (self *Game) isSpare(throw int) bool {
	return self.framePointsAt(throw) == allPins
}

func (self *Game) framePointsAt(throw int) int {
	return self.rolls[throw] + self.rolls[throw+1]
}

func (self *Game) framePointAt(throw int) int {
	return self.rolls[throw] + self.rolls[throw+1]
}

// testing utilities

func (self *Game) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		self.Roll(pins)
	}
}

func (self *Game) rollSpare() {
	self.Roll(5)
	self.Roll(5)
}

func (self *Game) rollStrike() {
	self.Roll(10)
}

const (
	allPins = 10

	framesPerGame = 10

	maxThrowPerGame = 21
)

package foo

import (
	. "github.com/smartystreets/convey"
	"testing"
)

func TestBowlingGameScoring(t *testing.T) {
	Convey("Given a fresh score card", t, func() {
		game := NewGame()

		Convey("The score should be zero", func() {
			So(game.Score(), ShouldEqual, 0)
		})
	})

	Convey("when all throws knock down one pin", func() {
		game.rollMany(20, 1)
		Convey("the score should be 20", func() {
			So(game.Score(), ShouldEqual, 20)
		})
	})

	Convey("when a spare is thrown", func() {
		game.rollSpare()
		game.Roll(3)
		game.rollMany(17, 0)

		convery("the score should be 20", func() {
			So(game.Score(), ShouldEqual, 20)
		})
	})

}

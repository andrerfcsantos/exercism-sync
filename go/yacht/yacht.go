package yacht

type Dice struct {
	rolls      []int
	faceToFreq map[int]int
	freqToFace map[int][]int
}

func (d *Dice) Total() int {
	sum := 0
	for _, d := range d.rolls {
		sum += d
	}
	return sum
}

func NewDice(rolls []int) *Dice {
	d := Dice{
		rolls:      rolls,
		faceToFreq: make(map[int]int),
		freqToFace: make(map[int][]int),
	}

	for _, r := range rolls {
		d.faceToFreq[r]++
	}

	for face, freq := range d.faceToFreq {
		d.freqToFace[freq] = append(d.freqToFace[freq], face)
	}

	return &d
}

type scoreFunction func(dice *Dice) int

var categories = map[string]scoreFunction{
	"ones": func(dice *Dice) int {
		return dice.faceToFreq[1]
	},
	"twos": func(dice *Dice) int {
		return dice.faceToFreq[2] * 2
	},
	"threes": func(dice *Dice) int {
		return dice.faceToFreq[3] * 3
	},
	"fours": func(dice *Dice) int {
		return dice.faceToFreq[4] * 4
	},
	"fives": func(dice *Dice) int {
		return dice.faceToFreq[5] * 5
	},
	"sixes": func(dice *Dice) int {
		return dice.faceToFreq[6] * 6
	},
	"full house": func(dice *Dice) int {
		_, threeOk := dice.freqToFace[3]
		if !threeOk {
			return 0
		}

		_, twoOk := dice.freqToFace[2]
		if !twoOk {
			return 0
		}

		return dice.Total()
	},
	"four of a kind": func(dice *Dice) int {

		fours, fourOk := dice.freqToFace[4]
		if fourOk {
			return fours[0] * 4
		}
		fives, fiveOk := dice.freqToFace[5]
		if fiveOk {
			return fives[0] * 4
		}

		return 0
	},
	"little straight": func(dice *Dice) int {
		if dice.faceToFreq[1] == 1 &&
			dice.faceToFreq[2] == 1 &&
			dice.faceToFreq[3] == 1 &&
			dice.faceToFreq[4] == 1 &&
			dice.faceToFreq[5] == 1 {
			return 30
		}
		return 0
	},
	"big straight": func(dice *Dice) int {
		if dice.faceToFreq[2] == 1 &&
			dice.faceToFreq[3] == 1 &&
			dice.faceToFreq[4] == 1 &&
			dice.faceToFreq[5] == 1 &&
			dice.faceToFreq[6] == 1 {
			return 30
		}
		return 0
	},
	"choice": func(dice *Dice) int {
		return dice.Total()
	},
	"yacht": func(dice *Dice) int {
		_, fiveOk := dice.freqToFace[5]
		if !fiveOk {
			return 0
		}
		return 50
	},
}

func Score(dice []int, category string) int {

	scoreFn, ok := categories[category]
	if !ok {
		return 0
	}

	return scoreFn(NewDice(dice))
}

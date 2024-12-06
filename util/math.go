package util

// There is no math.Abs for ints its kinda fucked up
func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

// Vec2 stuff
type Vec2 struct { X, Y int }

var VEC2_UP		 = Vec2{  0, -1 }
var VEC2_DOWN	 = Vec2{  0,  1	}
var VEC2_RIGHT = Vec2{  1,  0	}
var VEC2_LEFT	 = Vec2{ -1,  0	}

func (this Vec2) Add(that Vec2) Vec2 {
	return Vec2{this.X + that.X, this.Y + that.Y}
}

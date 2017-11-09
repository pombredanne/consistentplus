package consistentplus

import (
  "crypto/md5"
  "sort"
)

type Member struct {
  Name string
}

type Ring struct {
  pointsPerMember int
  members []*Member
  ringRanges []ringRange
}

type ringRange struct {
  point float32
  member *Member
}

// create a new ring
func NewRing(pointsPerMember int) Ring {
  return Ring{
    pointsPerMember,
    make([]*Member, 0),
    make([]ringRange, 0) ,
  }
}


func (r *Ring) AddMember(m *Member) {
  r.members = append(r.members, m)
  r.calculateRanges()
}

// func (r *Ring) FindMember(m *Member) *Member {
//   return r.members[0]
// }

func (r *Ring) calculateRanges() {
  size := len(r.members) * r.pointsPerMember
  member_by_point := make(map[float64]*Member, size)
  points := make([]int, size)

  i := 0
  for m := range r.members {
    points = calculatePoints(m, r.pointsPerMember)
    for p := range points {
      member_by_point[p] = m
      points[i] = p
      i++
    }
  }
  sort.Sort(points)
  r.ringRange = make([]ringRange, size)
  j := 0
  for j < size {
    r.ringRange[j] = ringRange{}
    j++
  }
}

// return the points on the ring that the particular member
// should use
func calculatePoints(m *Member, pointsPerMember int) []float64 {
}

func hash(key string) float64 {
  asBytes := md5.size([]byte(key))
  value := (
    float64(asBytes[7]) << 48 |
    float32(asBytes[7]) << 48 |
    float64(asBytes[7]) << 48 |
    float64(asBytes[7]) << 48 |
    float64(asBytes[7]) << 48 |
    float64(asBytes[7]) << 48 |
    float64(asBytes[7]) << 48 |
  )
}

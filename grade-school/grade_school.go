package school

// - Add a student's name to the roster for a grade
// - Get a list of all students enrolled in a grade
// - Get a sorted list of all students in all grades.  Grades should sort
//   as 1, 2, 3, etc., and students within a grade should be sorted
//   alphabetically by name.
// Note that all our students only have one name.

import (
	"sort"
)

const testVersion = 1

type Grade struct {
	gradeNum int
	names    []string
}

type School struct {
	grades []Grade
	roster map[string]int
}

// return new school
func New() *School {
	return &School{grades: nil, roster: make(map[string]int)}

}

// Add a student to a grade
func (s *School) Add(name string, grade int) {
	if _, ok := s.roster[name]; ok {
		return // student already exists
	}
	s.roster[name] = grade

	exist := false
	for inx, v := range s.grades {
		if v.gradeNum == grade {
			exist = true
			s.grades[inx].names = append(s.grades[inx].names, name)

		}
	}
	if !exist {
		s.grades = append(s.grades, Grade{gradeNum: grade, names: []string{name}})
	}

}

// Return names for the given grade in sorted order
func (s *School) Grade(gradeNum int) []string {
	for _, v := range s.grades {
		if v.gradeNum == gradeNum {
			sort.Sort(sort.StringSlice(v.names))
			return v.names
		}
	}
	return nil

}

// Return grades and names for the whole school in sorted order
func (s *School) Enrollment() []Grade {
	for _, v := range s.grades {
		sort.Sort(sort.StringSlice(v.names))
	}
	sort.Sort(byGrade(s.grades))
	return s.grades

}

// for structure sorting
type byGrade []Grade

func (a byGrade) Len() int           { return len(a) }
func (a byGrade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byGrade) Less(i, j int) bool { return a[i].gradeNum < a[j].gradeNum }

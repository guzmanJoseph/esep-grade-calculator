package esepunittests

type GradeCalculator struct {
	grades []Grade
	mode   GradingMode
}

type GradingMode int

const (
	LetterGradeMode GradingMode = iota
	PassFailMode
)

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(mode ...GradingMode) *GradeCalculator {
	gradingMode := LetterGradeMode
	if len(mode) > 0 {
		gradingMode = mode[0]
	}

	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   gradingMode,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.mode == PassFailMode {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAverage := computeAverageOfType(gc.grades, Assignment)
	examAverage := computeAverageOfType(gc.grades, Exam)
	essayAverage := computeAverageOfType(gc.grades, Essay)

	weightedGrade := assignmentAverage*0.50 +
		examAverage*0.35 +
		essayAverage*0.15

	return int(weightedGrade)
}

func computeAverageOfType(grades []Grade, gradeType GradeType) float64 {
	matchingGrades := make([]Grade, 0)
	for _, grade := range grades {
		if grade.Type == gradeType {
			matchingGrades = append(matchingGrades, grade)
		}
	}
	return computeAverage(matchingGrades)
}

func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0
	}

	sum := 0
	for _, grade := range grades {
		sum += grade.Grade
	}

	return float64(sum) / float64(len(grades))
}

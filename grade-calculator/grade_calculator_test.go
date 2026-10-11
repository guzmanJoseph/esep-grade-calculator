package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)
	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)
	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)
	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "C")
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 60, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "D")
	}
}

func TestGetGradeFBelow60(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 50, Assignment)
	gc.AddGrade("exam", 50, Exam)
	gc.AddGrade("essay", 50, Essay)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "F")
	}
}

func TestGetGradeWithNoGrades(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "F")
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		want      string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}
	for _, test := range tests {
		if got := test.gradeType.String(); got != test.want {
			t.Errorf("GradeType.String() = %q, want %q", got, test.want)
		}
	}
}

func TestAverageMultipleGrades(t *testing.T) {
	grades := []Grade{
		{Name: "one", Grade: 80},
		{Name: "two", Grade: 100},
	}
	if got := computeAverage(grades); got != 90 {
		t.Errorf("computeAverage() = %v, want 90", got)
	}
}


func TestPassFailPassAt70(t *testing.T) {
	gc := NewGradeCalculator(PassFailMode)
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)

	if got := gc.GetFinalGrade(); got != "Pass" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "Pass")
	}
}

func TestPassFailFailBelow70(t *testing.T) {
	gc := NewGradeCalculator(PassFailMode)
	gc.AddGrade("assignment", 69, Assignment)
	gc.AddGrade("exam", 69, Exam)
	gc.AddGrade("essay", 69, Essay)

	if got := gc.GetFinalGrade(); got != "Fail" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "Fail")
	}
}

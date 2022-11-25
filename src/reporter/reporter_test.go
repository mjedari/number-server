package reporter

import "testing"

func TestNewReporter(t *testing.T) {

	reporter := NewReporter()
	if reporter == nil {
		t.Error("failed to create reporter instance")
	}
}

func TestReporter_IncDuplicateNumber(t *testing.T) {
	// arrange
	reporter := NewReporter()
	num := uint(10)

	//act
	reporter.IncDuplicateNumber(10)

	//assert
	if reporter.DuplicateNumbers != num {
		t.Errorf("got %v insted of %v", reporter.DuplicateNumbers, num)
	}
}

func TestReporter_IncUniqueNumber(t *testing.T) {
	// arrange
	reporter := NewReporter()
	num := uint(10)

	//act
	reporter.IncUniqueNumber(10)

	//assert
	if reporter.UniqueNumbers != num {
		t.Errorf("got %v insted of %v", reporter.UniqueNumbers, num)
	}
}

func TestReporter_Report(t *testing.T) {
	//
}

func TestReporter_RestNumbers(t *testing.T) {
	// arrange
	reporter := NewReporter()
	uNum := uint(10)
	dNum := uint(12)

	//act
	reporter.IncUniqueNumber(uNum)
	reporter.IncDuplicateNumber(dNum)
	reporter.resetNumbers()

	// assert
	if reporter.DuplicateNumbers == dNum && reporter.UniqueNumbers == uNum {
		t.Errorf("expected 0 got %v:%v", reporter.UniqueNumbers, reporter.DuplicateNumbers)
	}

}

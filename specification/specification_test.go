package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isTrue(b bool) bool {
	return b == true
}

func TestNotSpecification(t *testing.T) {
	isTrue := Predicate(isTrue)
	notIsTrue := Not(isTrue)

	assert.Equal(
		t, true,
		notIsTrue.IsSatisfiedBy(false),
		"should return true when called with false")

	assert.Equal(
		t, false,
		notIsTrue.IsSatisfiedBy(true),
		"should return false when called with true")
}

func TestOrSpecification(t *testing.T) {
	isTrue := Predicate(isTrue)
	notIsTrue := Not(isTrue)
	isTrueOrNot := Or(notIsTrue, isTrue)
	isTrueOrIsTrue := Or(isTrue, isTrue)
	notIsTrueOrNotIsTrue := Or(notIsTrue, notIsTrue)

	assert.Equal(
		t, true,
		isTrueOrNot.IsSatisfiedBy(false),
		"should return true when called with false")

	assert.Equal(
		t, true,
		isTrueOrNot.IsSatisfiedBy(true),
		"should return true when called with true")

	assert.Equal(
		t, false,
		isTrueOrIsTrue.IsSatisfiedBy(false),
		"should return false when called with false")

	assert.Equal(
		t, true,
		isTrueOrIsTrue.IsSatisfiedBy(true),
		"should return true when called with true")

	assert.Equal(
		t, true,
		notIsTrueOrNotIsTrue.IsSatisfiedBy(false),
		"should return true when called with false")

	assert.Equal(
		t, false,
		notIsTrueOrNotIsTrue.IsSatisfiedBy(true),
		"should return false when called with true")
}

func TestAndSpecification(t *testing.T) {
	isTrue := Predicate(isTrue)
	notIsTrue := Not(isTrue)
	isTrueAndNot := And(notIsTrue, isTrue)
	isTrueAndIsTrue := And(isTrue, isTrue)
	notIsTrueAndNotIsTrue := And(notIsTrue, notIsTrue)

	assert.Equal(
		t, false,
		isTrueAndNot.IsSatisfiedBy(false),
		"should return false when called with false")

	assert.Equal(
		t, false,
		isTrueAndNot.IsSatisfiedBy(true),
		"should return false when called with true")

	assert.Equal(
		t, false,
		isTrueAndIsTrue.IsSatisfiedBy(false),
		"should return false when called with false")

	assert.Equal(
		t, true,
		isTrueAndIsTrue.IsSatisfiedBy(true),
		"should return true when called with true")

	assert.Equal(
		t, true,
		notIsTrueAndNotIsTrue.IsSatisfiedBy(false),
		"should return true when called with false")

	assert.Equal(
		t, false,
		notIsTrueAndNotIsTrue.IsSatisfiedBy(true),
		"should return false when called with true")
}

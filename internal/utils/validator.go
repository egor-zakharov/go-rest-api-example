package utils

import "library/internal/models"

func ValidateAddableBook(in models.Book) error {
	if !in.IsDraft() {
		return models.ErrorIncorrectId
	}
	if !in.IsValidAuthor() {
		return models.ErrorIncorrectAuthor
	}
	if !in.IsValidTitle() {
		return models.ErrorIncorrectTitle
	}
	if !in.IsValidReleasedYear() {
		return models.ErrorIncorrectYear
	}
	return nil
}

func ValidateUpdatableBook(in models.Book) error {
	if !in.IsValidId() {
		return models.ErrorNegativeId
	}
	if !in.IsValidAuthor() {
		return models.ErrorIncorrectAuthor
	}
	if !in.IsValidTitle() {
		return models.ErrorIncorrectTitle
	}
	if !in.IsValidReleasedYear() {
		return models.ErrorIncorrectYear
	}
	return nil
}

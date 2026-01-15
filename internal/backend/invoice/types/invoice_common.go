package types

import "-invoice_manager/internal/backend/invoice/payload"

func IncomeCategoryExists(classificationitem payload.ClassificationItem, summary []payload.ClassificationItem) (int, bool) {
	for index, category := range summary {
		if classificationitem.ClassificationCategory == category.ClassificationCategory && classificationitem.ClassificationType == category.ClassificationType {
			return index, true
		}
	}

	return 0, false
}

func ExpenseCategoryExists(classificationitem payload.ExpensesClassificationItem, summary []payload.ExpensesClassificationItem) (int, bool) {
	for index, category := range summary {
		if classificationitem.ClassificationCategory == category.ClassificationCategory && classificationitem.ClassificationType == category.ClassificationType {
			return index, true
		}
	}

	return 0, false
}

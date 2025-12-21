import { vatCategories, measurementUnitCodes } from "./data.js"
import { attachAutocomplete } from "./autocompletions.js"


const form = document.getElementById('create-product-form');



attachAutocomplete('vat_category_input', vatCategories, 'vatCategoriesSuggestions');
attachAutocomplete('measurementUnitInput', measurementUnitCodes, 'measurementUnitsSuggestions');



form.addEventListener('submit', async (e) => {
	e.preventDefault();
	const formData = new FormData(form);
	console.log(form.action);

	console.log(formData);
	try {
		const response = await fetch(form.action, {
			method: 'POST',
			body: formData
		});

		const data = await response.json();

		if (data.success) {
			console.log('success');
			window.location.reload();
		} else {
			console.log('error');
		}
	} catch (error) {
		console.log(error);
	}
});

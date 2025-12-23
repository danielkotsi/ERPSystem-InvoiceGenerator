import { vatCategories, measurementUnitCodes } from "./data.js"
import { attachAutocomplete } from "./autocompletions.js"


const form = document.getElementById('create-product-form');
const invoiceLink = document.querySelector(".invoice");
const suggestionsBox = document.querySelector(".invoice-suggestions");
const suggestions = [
	{ label: "Τιμολόγιο Πώλησης", href: "/makeaninvoice?invoice_type=1.1" },
	{ label: "Τιμολογιο Αγοράς", href: "/makeaninvoice?invoice_type=13.1" },
	{ label: "Απόδειξη Εισπραξης", href: "/makeaninvoice?invoice_type=8.1" },
	{ label: "Δελτιο Αποστολής", href: "/makeaninvoice?invoice_type=9.3" }
];

invoiceLink.addEventListener("click", (e) => {
	e.preventDefault(); // prevent navigation

	// Clear existing suggestions
	suggestionsBox.innerHTML = "";

	// Populate
	suggestions.forEach(item => {
		const a = document.createElement("a");
		a.href = item.href;
		a.textContent = item.label;
		suggestionsBox.appendChild(a);
	});

	// Show box
	suggestionsBox.style.display = "block";
});

// Hide when clicking outside
document.addEventListener("click", (e) => {
	if (!e.target.closest(".invoice-wrapper")) {
		suggestionsBox.style.display = "none";
	}
});



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

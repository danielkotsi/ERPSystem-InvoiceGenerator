import { invoiceTypes, incomeClassificationTypes, incomeClassificationCategories, paymentMethodCodes, paymentDueCodes } from "./data.js"
import { addAutocompletion, attachAutocomplete } from "./autocompletions.js"


const customersNameInput = document.getElementById('customersName')
const customer_suggestionsDiv = document.getElementById("customers-suggestions");

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
const customers_fields = {
	name: document.getElementById('customersName'),
	codeNumber: document.getElementById('customersCode'),
	doi: document.getElementById('customersDOI'),
	gemi: document.getElementById('customersGEMI'),
	phone: document.getElementById('customersPhone'),
	mobile_phone: document.getElementById('customersMobile_Phone'),
	email: document.getElementById('customersEmail'),
	'address.street': document.getElementById('customersAddressStreet'),
	'address.number': document.getElementById('customersAddressNumber'),
	'address.postalCode': document.getElementById('customersAddressPostalCode'),
	'address.city': document.getElementById('customersAddressCity'),
	vatNumber: document.getElementById('customersVatNumber'),
	country: document.getElementById('customersCountry'),
	branch: document.getElementById('customersBranch'),
	//this is a fiels from invoice details
	discount: document.getElementById('customersDiscount'),
};



addAutocompletion(customersNameInput, customer_suggestionsDiv, 'suggestions/customers?search=', customers_fields);


attachAutocomplete('paymentdue-input', paymentDueCodes, 'paymentdue-suggestions');

attachAutocomplete('paymentmethods-input', paymentMethodCodes, 'paymentmethods-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');



























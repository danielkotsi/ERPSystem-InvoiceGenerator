import { descriptions } from "./data.js"
import { invoiceTypes } from "./data.js"
import { vatCategories } from "./data.js"
import { incomeClassificationTypes } from "./data.js"
import { incomeClassificationCategories } from "./data.js"
const customersNameInput = document.getElementById('customersName')

async function fetchDB(fetchurl) {
	const response = await fetch(`${fetchurl}`, {
	})
	const data = await response.json();
	return data;
};

customersNameInput.addEventListener('input', (e) => {
	console.log(e.target.value);
	const customersuggestions = fetchDB('http://localhost:8080/suggestions/customers?search=' + e.target.value)
	console.log(customersuggestions);
});


customersNameInput.addEventListener('focus', (e) => {
	console.log(e.target.value);
	const customersuggestions = fetchDB('http://localhost:8080/suggestions/customers?search=' + e.target.value)
	console.log(customersuggestions);
});



function attachAutocomplete(inputId, items, whichsuggestions) {
	const input = document.getElementById(inputId);
	const suggestionsBox = document.getElementById(whichsuggestions);

	function renderSuggestions(list) {
		suggestionsBox.innerHTML = "";

		list.forEach(item => {
			const div = document.createElement("div");
			div.className = "suggestion-item";
			div.textContent = item.label;

			div.addEventListener("mousedown", () => {
				input.value = item.value;
				suggestionsBox.innerHTML = "";
			});

			suggestionsBox.appendChild(div);
		});
	}

	input.addEventListener("focus", () => {
		renderSuggestions(items);
	});

	input.addEventListener("input", () => {
		const value = input.value.toLowerCase();

		const filtered = items.filter(item =>
			item.label.toLowerCase().includes(value)
		);

		renderSuggestions(filtered);
	});

	document.addEventListener("click", (e) => {
		if (!e.target.closest("#" + inputId)) {
			suggestionsBox.innerHTML = "";
		}
	});
}

attachAutocomplete('descriptioninput', descriptions, 'description-suggestions');
attachAutocomplete('vatCategory', vatCategories, 'vatCategory-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');


let lineItemIndex = 1;
function addLineItem() {
	const div = document.createElement('div');
	div.classList.add('line-item');
	div.innerHTML = `
    <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>
    <label>Unit Price: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].unitPrice"></label><br>
    <label>VAT Category: <input type="text" id="vatCategory" name="invoiceDetails[${lineItemIndex}].vatCategory"></label><br>
  `;
	document.getElementById('invoiceDetails').appendChild(div);
	lineItemIndex++;
}

let paymentMethodIndex = 1;
function addPaymentMethod() {
	const div = document.createElement('div');
	div.classList.add('payment-method');
	div.innerHTML = `
    <label>Type (1=Bank, 2=Credit Card): <input type="number" name="paymentMethods.paymentdetails[${paymentMethodIndex}].type"></label><br>
    <label>Amount: <input type="number" step="0.01" name="paymentMethods.paymentdetails[${paymentMethodIndex}].amount"></label><br>
  `;
	document.getElementById('paymentMethods').appendChild(div);
	paymentMethodIndex++;
}


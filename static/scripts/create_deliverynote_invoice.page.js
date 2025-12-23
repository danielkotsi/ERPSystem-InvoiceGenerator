import { descriptions, invoiceTypes, vatCategories, incomeClassificationTypes, incomeClassificationCategories, paymentMethodCodes } from "./data.js"
import { addAutocompletion, attachAutocomplete, addBranchCompletion } from "./autocompletions.js"


const customersNameInput = document.getElementById('customersName')
const productNameInput = document.getElementById('product_name_input-0')
const branchesCodeInput = document.getElementById('branchCode')
const customer_suggestionsDiv = document.getElementById("customers-suggestions");
const product_suggestionsDiv = document.getElementById("product-suggestions-0");
const branches_suggestionsDiv = document.getElementById("branchcompany-suggestions");

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


const product_fields = {
	description: document.getElementById('product_description-0'),
	measurementUnit: document.getElementById('product_measurementUnit-0'),
	measurementUnitCode: document.getElementById('product_measurementUnitCode-0'),
	unitNetPrice: document.getElementById('product_unit_net_price-0'),
};

const branches_fieldsmap = {
	branchCode: document.getElementById('branchCode'),
	name: document.getElementById('branchName'),
	'address.street': document.getElementById('branchAddStreet'),
	'address.number': document.getElementById('branchAddNumber'),
	'address.postalCode': document.getElementById('branchAddPostalCode'),
	'address.city': document.getElementById('branchCity'),

};

const addproductButton = document.getElementById('addrowbutton');
function addLineItem() {
	const div = document.createElement('div');
	const container = document.getElementById('invoiceDetails');
	const lineItemIndex = container.querySelectorAll('.line-item').length;
	div.classList.add('line-item');
	div.innerHTML = `
	<button type="button" class="remove-line-item">Remove</button><br>
        <label>Product Name: <input type="text" id="product_name_input-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].name"></label><br>
	<div id="product-suggestions-${lineItemIndex}" class="suggestions"></div>
        <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>
        <label>Measurement_Unit: <input type="text" id="product_measurementUnit-${lineItemIndex}" step="1" name="product.measurementUnit"></label><br>
        <label>Measurement_Unit_Code: <input type="number" id="product_measurementUnitCode-${lineItemIndex}" step="1" name="invoiceDetails[${lineItemIndex}].measurementUnit"></label><br>
        <label>Unit Net Price: <input type="number" id="product_unit_net_price-${lineItemIndex}" step="0.01" name="invoiceDetails[${lineItemIndex}].unitNetPrice"></label><br>
        <label>Discount: <input type="number" id="customersDiscount-${lineItemIndex}" class="discount" step="1" name="buyer.discount"></label><br>
        <label>VAT Category: <input type="text" id="product_vatCategory-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].vatCategory" value="8"></label><br>
	<div id="vatCategory-suggestions" class="suggestions"></div>
        <label>Product Description: <input type="text" id="product_description-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].description"></label><br>
	<div id="description-suggestions" class="suggestions"></div>
	    <!-- IncomeClassification -->
        <label>Income Classification Category <input type="text" id="income_classification_category" name="invoiceDetails[${lineItemIndex}].incomeClassification.classificationCategory" value="category3"></label><br>
	<div id="income-classification-category-suggestions" class="suggestions"></div>
        <label>Income Classification Amount: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].incomeClassification.amount"></label><br>
  `;
	document.getElementById('invoiceDetails').appendChild(div);

	updateDiscount();

	const removeButton = div.querySelector('.remove-line-item');
	removeButton.addEventListener('click', () => {
		div.remove();
		reIndexLineItems();
	});
}

function updateDiscount() {
	const firstDiscount = document.getElementById('customersDiscount');
	if (!firstDiscount) return;

	const value = firstDiscount.value;

	const discountInputs = document.querySelectorAll('#invoiceDetails .discount');

	discountInputs.forEach(input => {
		input.value = value;
	});
}

function reIndexLineItems() {
	const lineItems = document.querySelectorAll('#invoiceDetails .line-item');

	lineItems.forEach((div, newIndex) => {
		// Update inputs inside this line item
		const inputs = div.querySelectorAll('input');

		inputs.forEach(input => {
			// Update name attributes that use the old index
			if (input.name.includes('invoiceDetails[')) {
				// Replace the old index with newIndex
				input.name = input.name.replace(/invoiceDetails\[\d+\]/, `invoiceDetails[${newIndex}]`);
			}

			// Update id attributes if they contain an index
			if (input.id && input.id.match(/-\d+$/)) {
				input.id = input.id.replace(/-\d+$/, `-${newIndex}`);
			}
		});

		// Update suggestion div ids if they exist
		const suggestionDivs = div.querySelectorAll('.suggestions');
		suggestionDivs.forEach(sug => {
			if (sug.id && sug.id.match(/-\d+$/)) {
				sug.id = sug.id.replace(/-\d+$/, `-${newIndex}`);
			}
		});
	});
}

addproductButton.addEventListener('click', () => {
	addLineItem();
	const container = document.getElementById('invoiceDetails');
	const lineItemIndex = container.querySelectorAll('.line-item').length;
	console.log('product_name_input-' + (lineItemIndex - 1));
	const productwithIndexNameInput = document.getElementById('product_name_input-' + (lineItemIndex - 1));
	const productwithIndexsuggestionsDiv = document.getElementById('product-suggestions-' + (lineItemIndex - 1));
	const productwithIndexfields = {
		description: document.getElementById('product_description-' + (lineItemIndex - 1)),
		measurementUnit: document.getElementById('product_measurementUnit-' + (lineItemIndex - 1)),
		measurementUnitCode: document.getElementById('product_measurementUnitCode-' + (lineItemIndex - 1)),
		unitNetPrice: document.getElementById('product_unit_net_price-' + (lineItemIndex - 1)),
		vatCategory: document.getElementById('product_vatCategory-' + (lineItemIndex - 1)),
	};
	addAutocompletion(productwithIndexNameInput, productwithIndexsuggestionsDiv, 'suggestions/products?search=', productwithIndexfields);

});



addBranchCompletion(branchesCodeInput, branches_suggestionsDiv, "suggestions/branchcompanies", branches_fieldsmap);
addAutocompletion(customersNameInput, customer_suggestionsDiv, 'suggestions/customers?search=', customers_fields);
addAutocompletion(productNameInput, product_suggestionsDiv, 'suggestions/products?search=', product_fields);



attachAutocomplete('paymentmethods-input', paymentMethodCodes, 'paymentmethods-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');

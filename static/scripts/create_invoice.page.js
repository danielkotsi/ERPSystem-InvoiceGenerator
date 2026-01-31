import { descriptions, invoiceTypes, vatCategories, incomeClassificationTypes, incomeClassificationCategories, paymentMethodCodes, paymentDueCodes } from "./data.js"
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
	name: document.getElementById('product_name_input-0'),
	codeNumber: document.getElementById('product_code_input-0'),
	description: document.getElementById('product_description-0'),
	measurementUnit: document.getElementById('product_measurementUnit-0'),
	measurementUnitCode: document.getElementById('product_measurementUnitCode-0'),
	unitNetPrice: document.getElementById('product_unit_net_price-0'),
	vatCategory: document.getElementById('product_vatCategory-0'),
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
//this is where it starts
function addLineItem() {
	const div = document.createElement('div');
	const container = document.getElementById('invoiceDetails');
	const lineItemIndex = container.querySelectorAll('.line-item').length;
	div.classList.add('line-item');
	div.innerHTML = `
	<button type="button" class="remove-line-item">Remove</button><br>
        <label>Όνομα Προϊόντος: <input type="text" id="product_name_input-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].name" required></label><br>
	<div id="product-suggestions-${lineItemIndex}" class="suggestions"></div>
	<label>Κωδικός Προϊόντος: <input type="text" autocomplete="off" id="product_code_input-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].codeNumber" required></label><br>
        <label>Περιγραφή Προϊόντος: <input type="text" id="product_description-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].itemDescr" required></label><br>
        <label>Επιλογή Έκπτωσης: <input type="text" id="discount-option-${lineItemIndex}" class="discount-option" name="invoiceDetails[${lineItemIndex}].discountOption"></label><br>
        <label>Ποσότητα: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity" required></label><br>
        <label>Μονάδα Μέτρησης: <input type="text" id="product_measurementUnit-${lineItemIndex}" step="1" name="invoiceDetails[${lineItemIndex}].measurementUnitName" required></label><br>
        <label>Κωδικός Μον.Μέτρησης: <input type="number" id="product_measurementUnitCode-${lineItemIndex}" step="1" name="invoiceDetails[${lineItemIndex}].measurementUnit" required></label><br>
        <label>Καθαρή Αξία Μονάδας: <input type="number" id="product_unit_net_price-${lineItemIndex}" step="0.01" name="invoiceDetails[${lineItemIndex}].unitNetPrice" required></label><br>
	<label>Έκπτωση: <input type="number" id="customersDiscount-${lineItemIndex}" class="discount" step="1" name="buyer.discount"></label><br>
        <label>Κατηγορία Φορολόγησης: <input type="text" id="product_vatCategory-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].vatCategory" required></label><br>
	<div id="vatCategory-suggestions" class="suggestions"></div>
	    <!-- IncomeClassification -->
        <label>Τύπος Κατάταξης Εσόδων <input type="text" id="income_classification_type" name="invoiceDetails[${lineItemIndex}].incomeClassification.classificationType" value="E3_561_001" required></label><br>
	<div id="income-classification-type-suggestions" class="suggestions"></div>
        <label>Κατηγορία Κατάταξης Εσόδων <input type="text" id="income_classification_category" name="invoiceDetails[${lineItemIndex}].incomeClassification.classificationCategory" value="category1_2" required></label><br>
	<div id="income-classification-category-suggestions" class="suggestions"></div>
        <label>Ποσότητα Κατάταξης Εσόδων : <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].incomeClassification.amount"></label><br>
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
//this is where it ends

addproductButton.addEventListener('click', () => {
	addLineItem();
	const container = document.getElementById('invoiceDetails');
	const lineItemIndex = container.querySelectorAll('.line-item').length;
	console.log('product_name_input-' + (lineItemIndex - 1));
	const productwithIndexNameInput = document.getElementById('product_name_input-' + (lineItemIndex - 1));
	const productwithIndexsuggestionsDiv = document.getElementById('product-suggestions-' + (lineItemIndex - 1));
	const productwithIndexfields = {
		name: document.getElementById('product_name_input-' + (lineItemIndex - 1)),
		codeNumber: document.getElementById('product_code_input-' + (lineItemIndex - 1)),
		description: document.getElementById('product_description-' + (lineItemIndex - 1)),
		measurementUnit: document.getElementById('product_measurementUnit-' + (lineItemIndex - 1)),
		measurementUnitCode: document.getElementById('product_measurementUnitCode-' + (lineItemIndex - 1)),
		unitNetPrice: document.getElementById('product_unit_net_price-' + (lineItemIndex - 1)),
		vatCategory: document.getElementById('product_vatCategory-' + (lineItemIndex - 1)),
	};
	addAutocompletion(productwithIndexNameInput, productwithIndexsuggestionsDiv, 'suggestions/products?search=', productwithIndexfields, 'suggestions/full/product?search=');

});



addBranchCompletion(branchesCodeInput, branches_suggestionsDiv, "suggestions/branchcompanies", branches_fieldsmap, 'suggestions/full/branchcompany?company=');
addAutocompletion(customersNameInput, customer_suggestionsDiv, 'suggestions/customers?search=', customers_fields, 'suggestions/full/customer?search=');
addAutocompletion(productNameInput, product_suggestionsDiv, 'suggestions/products?search=', product_fields, 'suggestions/full/product?search=');


attachAutocomplete('paymentdue-input', paymentDueCodes, 'paymentdue-suggestions');

attachAutocomplete('paymentmethods-input', paymentMethodCodes, 'paymentmethods-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');


const discountelement = document.getElementById('customersDiscount');

discountelement.addEventListener("change", () => {
	const value = discountelement.value
	console.log(value)
	const discountoptions = document.querySelectorAll('#invoiceDetails .discount-option');
	discountoptions.forEach(option => {
		if (Number(value) > 0) {
			option.value = 'true'
		} else {
			option.value = 'false'
		}
	});
});

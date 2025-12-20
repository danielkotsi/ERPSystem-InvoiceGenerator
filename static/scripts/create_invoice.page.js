import { descriptions, invoiceTypes, vatCategories, incomeClassificationTypes, incomeClassificationCategories } from "./data.js"
import { addAutocompletion, attachAutocomplete, addLineItem } from "./autocompletions.js"


const customersNameInput = document.getElementById('customersName')
const productNameInput = document.getElementById('product_name_input-0')
const customer_suggestionsDiv = document.getElementById("customers-suggestions");
const product_suggestionsDiv = document.getElementById("product-suggestions-0");

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
	vatCategory: document.getElementById('product_vatCategory-0'),
};


const addproductButton = document.getElementById('addrowbutton');

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




addAutocompletion(customersNameInput, customer_suggestionsDiv, 'suggestions/customers?search=', customers_fields);
addAutocompletion(productNameInput, product_suggestionsDiv, 'suggestions/products?search=', product_fields);



attachAutocomplete('descriptioninput', descriptions, 'description-suggestions');
attachAutocomplete('vatCategory', vatCategories, 'vatCategory-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');

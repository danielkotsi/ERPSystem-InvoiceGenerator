import { descriptions, invoiceTypes, vatCategories, incomeClassificationTypes, incomeClassificationCategories } from "./data.js"
import { addAutocompletion, attachAutocomplete } from "./autocompletions.js"


const customersNameInput = document.getElementById('customersName')
const productNameInput = document.getElementById('product_name_input-0')
const customer_suggestionsDiv = document.getElementById("customers-suggestions");
const product_suggestionsDiv = document.getElementById("product-suggestions");

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





addAutocompletion(customersNameInput, customer_suggestionsDiv, 'suggestions/customers?search=', customers_fields);
addAutocompletion(productNameInput, product_suggestionsDiv, 'suggestions/products?search=', product_fields);



attachAutocomplete('descriptioninput', descriptions, 'description-suggestions');
attachAutocomplete('vatCategory', vatCategories, 'vatCategory-suggestions');
attachAutocomplete('income_classification_type', incomeClassificationTypes, 'income-classification-type-suggestions');
attachAutocomplete('income_classification_category', incomeClassificationCategories, 'income-classification-category-suggestions');
attachAutocomplete('invoiceType', invoiceTypes, 'invoiceType-suggestions');

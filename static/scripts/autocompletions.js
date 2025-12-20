const baseURL = 'http://localhost:8080/'


async function fetchDB(fetchurl) {
	try {
		const response = await fetch(`${fetchurl}`, {
		})
		const data = await response.json();
		return data;
	} catch (error) {
		console.error("Fetch error:", error);
	}
};

export function addAutocompletion(element, elementsuggestions, endpoint, fieldsMap = {}) {
	element.addEventListener('input', async (e) => {
		console.log(e.target.value);
		const resultsuggestions = await fetchDB(baseURL + endpoint + e.target.value)
		console.log(resultsuggestions, Array.isArray(resultsuggestions))
		console.log(resultsuggestions);
		showSuggestions(resultsuggestions, element, elementsuggestions, fieldsMap);
	});


	element.addEventListener('focus', async (e) => {
		console.log(e.target.value);
		const resultsuggestions = await fetchDB(baseURL + endpoint + e.target.value)
		console.log(resultsuggestions);
		showSuggestions(resultsuggestions, element, elementsuggestions, fieldsMap);
	});

	document.addEventListener("click", (e) => {
		if (!e.target.closest("#" + element.id)) {
			elementsuggestions.innerHTML = "";
		}
	});
};


function showSuggestions(results, input, element, fieldsMap = {}) {
	clearSuggestions(element);

	results.forEach(item => {
		const div = document.createElement("div");
		div.className = "suggestion-item";
		div.textContent = item.name;

		div.addEventListener("click", () => {
			selectSuggestion(item, input, element, fieldsMap);
		});

		element.appendChild(div);
	});
}


function clearSuggestions(element) {
	element.innerHTML = "";
}
function selectSuggestion(item, input, element, fieldsMap = {}) {
	input.value = item.name;
	clearSuggestions(element);

	for (const key in fieldsMap) {
		const el = fieldsMap[key];
		const value = getNestedValue(item, key);
		if (value !== undefined && el instanceof HTMLElement) {
			el.value = value;
		}
	}


	console.log("Selected:", item);
}

function getNestedValue(item, fieldspath) {
	return fieldspath
		.replace(/\[(\d+)\]/g, ".$1") // convert [0] to .0
		.split(".")
		.reduce((acc, key) => {
			if (acc && acc[key] !== undefined) return acc[key];
			return undefined;
		}, item);
}


export function attachAutocomplete(inputId, items, whichsuggestions) {
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





// <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>

export function addLineItem() {
	const div = document.createElement('div');
	const container = document.getElementById('invoiceDetails');
	const lineItemIndex = container.querySelectorAll('.line-item').length;
	div.classList.add('line-item');
	div.innerHTML = `
<!-- this is the new addition  -->
	<button type="button" class="remove-line-item">Remove</button><br>
<!-- it ends here  -->
        <label>Product Name: <input type="text" id="product_name_input-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].name"></label><br>
	<div id="product-suggestions-${lineItemIndex}" class="suggestions"></div>
        <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>
        <label>Measurement_Unit: <input type="text" id="product_measurementUnit-${lineItemIndex}" step="1" name="product.measurementUnit"></label><br>
        <label>Measurement_Unit_Code: <input type="number" id="product_measurementUnitCode-${lineItemIndex}" step="1" name="invoiceDetails[${lineItemIndex}].measurementUnit"></label><br>
        <label>Unit Net Price: <input type="number" id="product_unit_net_price-${lineItemIndex}" step="0.01" name="invoiceDetails[${lineItemIndex}].unitNetPrice"></label><br>
        <label>Discount: <input type="number" id="customersDiscount" step="1" name="buyer.discount"></label><br>
        <label>VAT Category: <input type="text" id="product_vatCategory-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].vatCategory"></label><br>
	<div id="vatCategory-suggestions" class="suggestions"></div>
        <label>Product Description: <input type="text" id="product_description-${lineItemIndex}" name="invoiceDetails[${lineItemIndex}].description"></label><br>
	<div id="description-suggestions" class="suggestions"></div>
	    <!-- IncomeClassification -->
        <label>Income Classification Type <input type="text" id="income_classification_type" name="invoiceDetails[${lineItemIndex}].incomeClassification.classificationType"></label><br>
	<div id="income-classification-type-suggestions" class="suggestions"></div>
        <label>Income Classification Category <input type="text" id="income_classification_category" name="invoiceDetails[${lineItemIndex}].incomeClassification.classificationCategory"></label><br>
	<div id="income-classification-category-suggestions" class="suggestions"></div>
        <label>Income Classification Amount: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].incomeClassification.amount"></label><br>
	    <!-- expensesClassification -->
        <label>Expenses Classification Type <input type="text" id="expenses_classification_type" name="invoiceDetails[${lineItemIndex}].expensesClassification.classificationType"></label><br>
	<div id="expenses-classification-type-suggestions" class="suggestions"></div>
        <label>Expenses Classification Category <input type="text" id="expenses_classification_category" name="invoiceDetails[${lineItemIndex}].expensesClassification.classificationCategory"></label><br>
	<div id="expenses-classification-category-suggestions" class="suggestions"></div>
        <label>Expenses Classification Amount: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].expensesClassification.amount"></label><br>
  `;
	document.getElementById('invoiceDetails').appendChild(div);
	// Add event listener to the remove button
	const removeButton = div.querySelector('.remove-line-item');
	removeButton.addEventListener('click', () => {
		div.remove(); // removes this line item from the DOM
		reIndexLineItems();
	});
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


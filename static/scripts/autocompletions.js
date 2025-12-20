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






// let lineItemIndex = 1;
// function addLineItem() {
// 	const div = document.createElement('div');
// 	div.classList.add('line-item');
// 	div.innerHTML = `
//     <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>
//     <label>Unit Price: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].unitPrice"></label><br>
//     <label>VAT Category: <input type="text" id="vatCategory" name="invoiceDetails[${lineItemIndex}].vatCategory"></label><br>
//   `;
// 	document.getElementById('invoiceDetails').appendChild(div);
// 	lineItemIndex++;
// }
//
// let paymentMethodIndex = 1;
// function addPaymentMethod() {
// 	const div = document.createElement('div');
// 	div.classList.add('payment-method');
// 	div.innerHTML = `
//     <label>Type (1=Bank, 2=Credit Card): <input type="number" name="paymentMethods.paymentdetails[${paymentMethodIndex}].type"></label><br>
//     <label>Amount: <input type="number" step="0.01" name="paymentMethods.paymentdetails[${paymentMethodIndex}].amount"></label><br>
//   `;
// 	document.getElementById('paymentMethods').appendChild(div);
// 	paymentMethodIndex++;
// }
//

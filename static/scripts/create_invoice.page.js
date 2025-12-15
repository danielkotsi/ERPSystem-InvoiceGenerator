const descriptions = [
	"something",
	"new",
	"old",
];
function attachAutocomplete(inputId, items) {
	const input = document.getElementById(inputId);
	const suggestionsBox = document.getElementById("suggestions");

	function renderSuggestions(list) {
		suggestionsBox.innerHTML = "";

		list.forEach(item => {
			const div = document.createElement("div");
			div.className = "suggestion-item";
			div.textContent = item;

			div.addEventListener("mousedown", () => {
				input.value = item;
				suggestionsBox.innerHTML = "";
			});

			suggestionsBox.appendChild(div);
		});
	}

	input.addEventListener("focus", () => {
		renderSuggestions(items);
	});

	// Filter items while typing
	input.addEventListener("input", () => {
		const value = input.value.toLowerCase();

		const filtered = items.filter(item =>
			item.toLowerCase().includes(value)
		);

		renderSuggestions(filtered);
	});

	// Close suggestions when clicking outside
	document.addEventListener("click", (e) => {
		if (!e.target.closest("#" + inputId)) {
			suggestionsBox.innerHTML = "";
		}
	});
}
attachAutocomplete('descriptioninput', descriptions);


let lineItemIndex = 1;
function addLineItem() {
	const div = document.createElement('div');
	div.classList.add('line-item');
	div.innerHTML = `
    <label>Quantity: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].quantity"></label><br>
    <label>Unit Price: <input type="number" step="0.01" name="invoiceDetails[${lineItemIndex}].unitPrice"></label><br>
    <label>VAT Category: <input type="number" name="invoiceDetails[${lineItemIndex}].vatCategory"></label><br>
  `;
	document.getElementById('ivoiceDetails').appendChild(div);
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

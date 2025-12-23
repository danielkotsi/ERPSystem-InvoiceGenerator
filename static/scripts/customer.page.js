const form = document.getElementById('create-customer-form');
const baseURL = 'http://localhost:8080/'

const searchInput = document.getElementById('customersearch')
const searchInputSuggestions = document.getElementById('customer-search-suggestions')
const invoiceLink = document.querySelector(".invoice");
const suggestionsBox = document.querySelector(".invoice-suggestions");

// Example suggestions
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


function addAutocompletion(element, elementsuggestions, endpoint) {
	element.addEventListener('input', async (e) => {
		console.log(e.target.value);
		const resultsuggestions = await fetchDB(baseURL + endpoint + e.target.value)
		console.log(resultsuggestions, Array.isArray(resultsuggestions))
		console.log(resultsuggestions);
		showSuggestions(resultsuggestions, elementsuggestions);
	});


	element.addEventListener('focus', async (e) => {
		console.log(e.target.value);
		const resultsuggestions = await fetchDB(baseURL + endpoint + e.target.value)
		console.log(resultsuggestions);
		showSuggestions(resultsuggestions, elementsuggestions);
	});

	document.addEventListener("click", (e) => {
		if (!e.target.closest("#" + element.id)) {
			elementsuggestions.innerHTML = "";
		}
	});
};


function showSuggestions(results, element) {
	clearSuggestions(element);

	results.forEach(item => {
		const a = document.createElement("a");
		a.className = "suggestion-item";
		a.textContent = item.name;

		// build the destination URL
		a.href = `/customers/byid/${encodeURIComponent(item.codeNumber)}`;

		element.appendChild(a);
	});
}

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

function clearSuggestions(element) {
	element.innerHTML = "";
}


addAutocompletion(searchInput, searchInputSuggestions, 'suggestions/customers?search=');

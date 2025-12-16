const descriptions = [
	{ value: "APL", label: "Apple – Fresh green apples" },
	{ value: "BAN", label: "Banana – Organic yellow bananas" },
	{ value: "ORG", label: "Orange – Sweet Valencia oranges" }
];

const invoiceTypes = [
	{ value: "1.1", label: "1.1 Τιμολόγιο Πώλησης" },
	{ value: "1.2", label: "1.2 Τιμολόγιο Πώλησης / Ενδοκοινοτικές Παραδόσεις" },
	{ value: "1.3", label: "1.3 Τιμολόγιο Πώλησης / Παραδόσεις Τρίτων Χωρών" },
	{ value: "1.4", label: "1.4 Τιμολόγιο Πώλησης / Πώληση για Λογαριασμό Τρίτων" },
	{ value: "1.5", label: "1.5 Τιμολόγιο Πώλησης / Εκκαθάριση Πωλήσεων Τρίτων - Αμοιβή από Πωλήσεις Τρίτων" },
	{ value: "1.6", label: "1.6 Τιμολόγιο Πώλησης / Συμπληρωματικό Παραστατικό" }
];

const vatCategories = [
	{ value: "1", label: "1-ΦΠΑ συντελεστής 24% " },
	{ value: "2", label: "2-ΦΠΑ συντελεστής 13%" },
	{ value: "3", label: "3-ΦΠΑ συντελεστής 6% " },
	{ value: "4", label: "4-ΦΠΑ συντελεστής 17% " },
	{ value: "5", label: "5-ΦΠΑ συντελεστής 9% " },
	{ value: "6", label: "6-ΦΠΑ συντελεστής 4%" },
	{ value: "7", label: "7-Άνευ Φ.Π.Α." },
	{ value: "8", label: "8-Εγγραφές χωρίς ΦΠΑ(πχ Μισθοδοσία, Αποσβέσεις)" },
	{ value: "9", label: "9-ΦΠΑ συντελεστής 3% (αρ.31 ν.5057/2023) " },
	{ value: "10", label: "10-ΦΠΑ συντελεστής 4% (αρ.31 ν.5057/2023)" }
];
const incomeClassificationTypes = [
	{ value: "E3_106", label: "E3_106 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Εμπορεύματα" },
	{ value: "E3_205", label: "E3_205 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά" },
	{ value: "E3_210", label: "E3_210 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη" },
	{ value: "E3_305", label: "E3_305 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις – Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά" },
	{ value: "E3_310", label: "E3_310 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη" },
	{ value: "E3_318", label: "E3_318 Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Έξοδα παραγωγής" },
	{ value: "E3_561_001", label: "E3_561_001 Πωλήσεις αγαθών και υπηρεσιών Χονδρικές - Επιτηδευματιών" },
	{ value: "E3_561_002", label: "E3_561_002 Πωλήσεις αγαθών και υπηρεσιών Χονδρικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)" },
	{ value: "E3_561_003", label: "E3_561_003 Πωλήσεις αγαθών και υπηρεσιών Λιανικές - Ιδιωτική Πελατεία" },
	{ value: "E3_561_004", label: "E3_561_004 Πωλήσεις αγαθών και υπηρεσιών Λιανικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)" },
	{ value: "E3_561_005", label: "E3_561_005 Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Ενδοκοινοτικές" },
	{ value: "E3_561_006", label: "E3_561_006 Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Τρίτες Χώρες" },
	{ value: "E3_561_007", label: "E3_561_007 Πωλήσεις αγαθών και υπηρεσιών Λοιπά" },
	{ value: "E3_562", label: "E3_562 Λοιπά συνήθη έσοδα" },
	{ value: "E3_563", label: "E3_563 Πιστωτικοί τόκοι και συναφή έσοδα" },
	{ value: "E3_564", label: "E3_564 Πιστωτικές συναλλαγματικές διαφορές" },
	{ value: "E3_565", label: "E3_565 Έσοδα συμμετοχών" },
	{ value: "E3_566", label: "E3_566 Κέρδη από διάθεση μη κυκλοφορούντων περιουσιακών στοιχείων" },
	{ value: "E3_567", label: "E3_567 Κέρδη από αναστροφή προβλέψεων και απομειώσεων" },
	{ value: "E3_568", label: "E3_568 Κέρδη από επιμέτρηση στην εύλογη αξία" },
	{ value: "E3_570", label: "E3_570 Ασυνήθη έσοδα και κέρδη" },
	{ value: "E3_595", label: "E3_595 Έξοδα σε ιδιοπαραγωγή" },
	{ value: "E3_596", label: "E3_596 Επιδοτήσεις - Επιχορηγήσεις" },
	{ value: "E3_597", label: "E3_597 Επιδοτήσεις - Επιχορηγήσεις για επενδυτικούς σκοπούς - κάλυψη δαπανών" },
	{ value: "E3_880_001", label: "E3_880_001 Πωλήσεις Παγίων Χονδρικές" },
	{ value: "E3_880_002", label: "E3_880_002 Πωλήσεις Παγίων Λιανικές" },
	{ value: "E3_880_003", label: "E3_880_003 Πωλήσεις Παγίων Εξωτερικού Ενδοκοινοτικές" },
	{ value: "E3_880_004", label: "E3_880_004 Πωλήσεις Παγίων Εξωτερικού Τρίτες Χώρες" },
	{ value: "E3_881_001", label: "E3_881_001 Πωλήσεις για λογ/σμο Τρίτων Χονδρικές" },
	{ value: "E3_881_002", label: "E3_881_002 Πωλήσεις για λογ/σμο Τρίτων Λιανικές" },
	{ value: "E3_881_003", label: "E3_881_003 Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Ενδοκοινοτικές" },
	{ value: "E3_881_004", label: "E3_881_004 Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Τρίτες Χώρες" },
	{ value: "E3_598_001", label: "E3_598_001 Πωλήσεις αγαθών που υπάγονται σε ΕΦΚ" },
	{ value: "E3_598_003", label: "E3_598_003 Πωλήσεις για λογαριασμό αγροτών μέσω αγροτικού συνεταιρισμού κλπ" }
];

const incomeClassificationCategories = [
	{ value: "category1_1", label: "category1_1 Έσοδα από Πώληση Εμπορευμάτων (+) / (-)" },
	{ value: "category1_2", label: "category1_2 Έσοδα από Πώληση Προϊόντων (+) / (-)" },
	{ value: "category1_3", label: "category1_3 Έσοδα από Παροχή Υπηρεσιών (+) / (-)" },
	{ value: "category1_4", label: "category1_4 Έσοδα από Πώληση Παγίων (+) / (-)" },
	{ value: "category1_5", label: "category1_5 Λοιπά Έσοδα/ Κέρδη (+) / (-)" },
	{ value: "category1_6", label: "category1_6 Αυτοπαραδόσεις / Ιδιοχρησιμοποιήσεις (+) / (-)" },
	{ value: "category1_7", label: "category1_7 Έσοδα για λ/σμο τρίτων (+) / (-)" },
	{ value: "category1_8", label: "category1_8 Έσοδα προηγούμενων χρήσεων (+)/ (-)" },
	{ value: "category1_9", label: "category1_9 Έσοδα επομένων χρήσεων (+) / (-)" },
	{ value: "category1_10", label: "category1_10 Λοιπές Εγγραφές Τακτοποίησης Εσόδων (+) / (-)" },
	{ value: "category1_95", label: "category1_95 Λοιπά Πληροφοριακά Στοιχεία Εσόδων (+) / (-)" },
	{ value: "category3", label: "category3 Διακίνηση" }
];



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


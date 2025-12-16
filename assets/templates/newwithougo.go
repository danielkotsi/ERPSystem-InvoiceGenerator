<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Invoice PDF</title>
  <style>
    @page { size: A4; margin: 20mm; }
    * { box-sizing: border-box; font-family: Arial, Helvetica, sans-serif; }
    body { margin: 0; font-size: 12px; color: #000; }

    .invoice { width: 100%; }

    /* HEADER */
    .invoice-head {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      margin-bottom: 15px;
    }

    .logo { width: 35%; }
    .logo-image { max-width: 180px; }
    .slogan { margin-top: 6px; font-style: italic; font-size: 11px; }

    .seller { width: 35%; font-size: 11px; }

    .qr-box { width: 20%; text-align: center; }
    .qr { width: 90px; height: 90px; }
    .qr-label { font-size: 9px; margin-top: 4px; }

    /* GENERIC TABLE */
    table { width: 100%; border-collapse: collapse; }
    th, td { border: 1px solid #000; padding: 5px; }
    th { background: #f2f2f2; text-align: center; }
    td { text-align: right; }
    td.left { text-align: left; }

    /* HEADER INFO TABLE */
    .invoice-header td { border: none; padding: 3px; }

    /* BUYER */
    .buyer { display: flex; justify-content: space-between; margin: 15px 0; }
    .buyer-info, .correlated-invoices { width: 48%; }

    /* ROWS */
    .invoice-rows { table-layout: fixed; margin-bottom: 20px; }
    .invoice-rows th.desc { width: 40%; }
    .invoice-rows th.qty { width: 10%; }
    .invoice-rows th.unit { width: 15%; }
    .invoice-rows th.value { width: 15%; }
    .invoice-rows th.vat { width: 10%; }
    .invoice-rows th.total { width: 10%; }

    .sub-row td { border-top: none; font-size: 11px; padding-left: 20px; }

    /* DETAILS */
    .invoice-details { display: flex; gap: 10px; }
    .balance-table, .vat-analysis { width: 40%; }
    .totals { width: 20%; }

    .totals td { border: none; }
    .totals tr:last-child td { border-top: 2px solid #000; font-weight: bold; }

    /* FOOTER */
    footer { margin-top: 25px; }
    .terms-of-sale { font-size: 10px; margin-bottom: 15px; }

    .footer-low { display: flex; justify-content: space-between; }
    .bank-accounts { width: 40%; }
    .publisher, .receiver { width: 25%; text-align: center; font-size: 11px; }
    .sign { margin-top: 40px; border-top: 1px solid #000; }
  </style>
</head>
<body>
<div class="invoice">

<header class="invoice-head">
  <div class="logo">
    <img class="logo-image" src="logo.png" alt="Logo">
    <div class="slogan">With respect to tradition. Since 1985</div>
  </div>

  <div class="seller">
    <strong>LIVERA</strong><br>
    Traditional Confectionery Products<br>
    VAT: 123456789<br>
    Address line<br>
  </div>

  <div class="qr-box">
    <img class="qr" src="data:image/png;base64,{{ .Invoice.QrBase64 }}">
    <div class="qr-label">Scan to verify invoice</div>
  </div>
</header>

<table class="invoice-header">
  <tr>
    <td class="left">Invoice No: </td>
    <td>Date: </td>
  </tr>
  <tr>
    <td class="left">UID: </td>
    <td>MARK: </td>
  </tr>
</table>

<section class="buyer">
  <div class="buyer-info">
    <strong>Buyer</strong><br>
    <br>
    <br>
    VAT: 
  </div>

  <div class="correlated-invoices">
    <strong>Related Invoices</strong><br>
    {}<br>
  </div>
</section>

<table class="invoice-rows">
  <thead>
    <tr>
      <th class="desc">Description</th>
      <th class="qty">Qty</th>
      <th class="unit">Unit</th>
      <th class="value">Value</th>
      <th class="vat">% VAT</th>
      <th class="total">Total</th>
    </tr>
  </thead>
  <tbody>
    
    <tr>
      <td class="left"></td>
      <td></td>
      <td></td>
      <td></td>
      <td>%</td>
      <td></td>
    </tr>
    
    <tr class="sub-row">
      <td class="left">– </td>
      <td></td><td></td><td></td><td></td><td></td>
    </tr>
    
    
  </tbody>
</table>

<section class="invoice-details">
  <table class="balance-table">
    <tr><td class="left">Previous Balance</td><td></td></tr>
    <tr><td class="left">New Balance</td><td></td></tr>
  </table>

  <table class="vat-analysis">
    <tr><th>Value</th><th>% VAT</th><th>VAT</th><th>Total</th></tr>
    
    <tr>
      <td></td>
      <td>%</td>
      <td></td>
      <td></td>
    </tr>
    
  </table>

  <table class="totals">
    <tr><td class="left">Value</td><td></td></tr>
    <tr><td class="left">Discount</td><td></td></tr>
    <tr><td class="left">Total VAT</td><td></td></tr>
    <tr><td class="left">Grand Total</td><td></td></tr>
  </table>
</section>

<footer>
  <div class="terms-of-sale">*Traditional confectionery products</div>

  <div class="footer-low">
    <table class="bank-accounts">
      <tr><td class="left">PIRAEUS</td></tr>
      <tr><td class="left">OPTIMA</td></tr>
      <tr><td class="left">EUROBANK</td></tr>
      <tr><td class="left">ALPHA BANK – GR8701401420142002002018579 – BIC: CRBAGRAA</td></tr>
    </table>

    <div class="publisher">
      THE PUBLISHER
      <div class="sign"></div>
    </div>

    <div class="receiver">
      THE RECEIVER
      <div class="sign"></div>
    </div>
  </div>
</footer>

</div>
</body>
</html>

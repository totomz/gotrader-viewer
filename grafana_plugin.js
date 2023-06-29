// ==UserScript==
// @name         Gotrader Mark Annotation
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       You
// @match        http://localhost:3000/d*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=undefined.localhost
// @grant        none
// ==/UserScript==

(function () {
	'use strict';

	const dashboard = 'ed0d1c26-e1f2-41fc-9ff3-e61de5bee663'; // TODO rigenera la dashboard con un uid statico
	const panelId = 1;	// Questo lo prendi facendo "edit panel"
	const symbol = $("#var-symbol").val();
	
	document.addEventListener("keydown", (event) => {
		if (event.key !== 'b' && event.key !== 's') {
			return;
		}

		const ts = $($('*[data-testid="SeriesTableRow"]')['0']).parent().children('*[aria-label="Timestamp"]').text();
		let action = "buy";
		if (event.key === 's') {
			action = "sell";
		}

		console.log(`${action}  at ${ts}`)

		fetch("/api/annotations", {
			method: "POST", 
			mode: "cors", 
			cache: "no-cache", 
			credentials: "same-origin", 
			headers: { "Content-Type": "application/json" },
			redirect: "follow", 
			referrerPolicy: "no-referrer", 
			body: JSON.stringify({
				dashboardUID: dashboard,
				panelId: panelId,
				time: new Date(ts).getTime(),
				tags:[action, `symbol:${symbol}`],
  				text:`Dajjee ${action}`
			}), 
		}).then(res => {
			console.log("mi sei mancato? no per niente. Mi manca lavorare con Gio? Non nel modo in cui credo");
			console.log(res);
		})


		// if (event.isComposing || event.keyCode === 229) {
		//   return;
		// }
		// do something
	});

})();
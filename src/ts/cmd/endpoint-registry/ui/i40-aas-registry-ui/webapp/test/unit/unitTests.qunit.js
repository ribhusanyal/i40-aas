/* global QUnit */
QUnit.config.autostart = false;

sap.ui.getCore().attachInit(function () {
	"use strict";

	sap.ui.require([
		"i40-aas-registry-ui/i40-aas-registry-ui/test/unit/AllTests"
	], function () {
		QUnit.start();
	});
});
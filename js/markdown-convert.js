
document.addEventListener('DOMContentLoaded', function() {
    var mdText = document.getElementById('md-text').innerText;
    var mdTarget = document.getElementById('md-target');

	var converter = new showdown.Converter();
	var generatedHTML = converter.makeHtml(mdText);

	mdTarget.innerHTML = generatedHTML;
}, false);

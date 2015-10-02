document.addEventListener('DOMContentLoaded', function() {
	var articleName = parse('a');
	getArticle(articleName, function(mdText) {
	    var mdTarget = document.getElementById('md-target');
	    var converter = new showdown.Converter();
	    var generatedHTML = converter.makeHtml(mdText);
	    mdTarget.innerHTML = generatedHTML;
	});

}, false);

function parse(val) {
    var result = "Not found",
        tmp = [];
    location.search
        //.replace ( "?", "" ) 
        // this is better, there might be a question mark inside
        .substr(1)
        .split("&")
        .forEach(function(item) {
            tmp = item.split("=");
            if (tmp[0] === val) result = decodeURIComponent(tmp[1]);
        });
    return result;
}

function getArticle(name, callback) {
	var url = '/articles/' + name + '.md';
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'text';
    xhr.onload = function() {
        var status = xhr.status;
        if (status == 200) {
            console.log(xhr.response);
            callback(xhr.response);
        }
    };
    xhr.send();
}

const contactUs = document.getElementById("contactUs")




var scrollingElement = (document.scrollingElement || document.body);
contactUs.onclick = function () {
    scrollingElement.scrollTop = scrollingElement.scrollHeight - scrollingElement.clientHeight
}

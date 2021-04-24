const contactUs = document.getElementById("contactUs")
const slider = document.getElementsByClassName("slider")
const banner = slider[0].getElementsByClassName("banner")
const nextBannerButton = document.getElementById("nextBanner")
var setBanner = 0;
var bannerInterval
const sendEmail = document.getElementById("sendEmail")
const customerName = document.getElementById("name")
const email = document.getElementById("email")
const query = document.getElementById("query")

function bannerChanger() {
    for (var i = 0; i < banner.length; ++i) {
        banner[i].style.display = "none"
    }
    banner[setBanner].style.display = "block";
    setBanner = (++setBanner) % banner.length
    bannerInterval = setTimeout(bannerChanger, 7000);
}

nextBannerButton.onclick = function () {
    for (var i = 0; i < banner.length; ++i) {
        banner[i].style.display = "none"
    }
    banner[setBanner].style.display = "block";
    setBanner = (++setBanner) % banner.length
    clearTimeout(bannerInterval)
    bannerInterval = setTimeout(bannerChanger, 7000);
}

bannerChanger()

var scrollingElement = (document.scrollingElement || document.body);
contactUs.onclick = function () {
    scrollingElement.scrollTop = scrollingElement.scrollHeight - scrollingElement.clientHeight
}

sendEmail.onclick = function () {
    var receiverName = customerName.value
    var receiver = email.value
    var message = query.value
    var URL = `https://queu-e.herokuapp.com/?resource=verify&function=sendOtp&email=${receiver}`
    var xhr = new XMLHttpRequest()
    xhr.open("GET", URL)
    xhr.send()
}

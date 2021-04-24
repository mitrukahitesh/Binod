const contactUs = document.getElementById("contactUs")
const slider = document.getElementsByClassName("slider")
const banner = slider[0].getElementsByClassName("banner")
const nextBannerButton = document.getElementById("nextBanner")
var setBanner = 0;
var bannerInterval
// const sendEmail = document.getElementById("sendEmail")
// const customerName = document.getElementById("name")
// const email = document.getElementById("email")
// const query = document.getElementById("query")
const grids = document.getElementsByClassName("grid")

function getBanners() {
    var xhr = new XMLHttpRequest()
    xhr.responseType = 'json'
    xhr.open("GET", "https://ecell-binod.herokuapp.com/?resource=banner")
    xhr.onload = () => {
        for (var i = 0; i < banner.length; ++i) {
            banner[i].src = xhr.response.banners[i].link
        }
    }
    xhr.send()
}

function getMenCategory() {
    var xhr = new XMLHttpRequest()
    xhr.responseType = 'json'
    xhr.open("GET", "https://ecell-binod.herokuapp.com/?resource=category&category=Men")
    xhr.onload = () => {
        for (var i = 0; i < 5; ++i) {
            var item = xhr.response.category.items[i]
            grids[i].getElementsByClassName("preview")[0].src = item.img
            grids[i].getElementsByClassName("item")[0].textContent = item.name
            grids[i].getElementsByClassName("offer")[0].textContent = item.offer
        }
    }
    xhr.send()
}

function getWomenCategory() {
    var xhr = new XMLHttpRequest()
    xhr.responseType = 'json'
    xhr.open("GET", "https://ecell-binod.herokuapp.com/?resource=category&category=Women")
    xhr.onload = () => {
        for (var i = 6; i < 11; ++i) {
            var item = xhr.response.category.items[i - 6]
            grids[i].getElementsByClassName("preview")[0].src = item.img
            grids[i].getElementsByClassName("item")[0].textContent = item.name
            grids[i].getElementsByClassName("offer")[0].textContent = item.offer
        }
    }
    xhr.send()
}

getBanners()
getMenCategory()
getWomenCategory()

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

// sendEmail.onclick = function () {
//     var receiverName = customerName.value
//     var receiver = email.value
//     var message = query.value
//     var URL = `https://queu-e.herokuapp.com/?resource=verify&function=sendOtp&email=${receiver}`
//     var xhr = new XMLHttpRequest()
//     xhr.open("GET", URL)
//     xhr.send()
// }

const contactUs = document.getElementById("contactUs")
const slider = document.getElementsByClassName("slider")
const banner = slider[0].getElementsByClassName("banner")
const nextBannerButton = document.getElementById("nextBanner")
var setBanner = 0;
var bannerInterval

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

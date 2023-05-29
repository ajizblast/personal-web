let hamburgerMenuOpen = false
function openHamburger() {
  let hamburgerMenu = document.getElementById('hamburger-menu')
  if(!hamburgerMenuOpen) {
    hamburgerMenu.style.display = "block";
    hamburgerMenuOpen = true;
  } else {
    hamburgerMenu.style.display = "none";
    hamburgerMenuOpen = false;
  }
}
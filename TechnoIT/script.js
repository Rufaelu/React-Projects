window.addEventListener("scroll", function () {
  const nav = document.getElementById("navbar");
  if (window.scrollY > 20) {
    nav.classList.add("bg-white");
    nav.classList.add("text-black");
    nav.classList.remove("bg-transparent");
    nav.classList.remove("text-white");
  } else {
    navbar.classList.add("bg-transparent");
    nav.classList.add("text-white");
    nav.classList.remove("text-black");

    navbar.classList.remove("bg-white");
  }
});

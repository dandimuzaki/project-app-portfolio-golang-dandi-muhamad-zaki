window.onscroll = function() {
    scrollFunction();
};

function scrollFunction() {
  const header = document.getElementById("header");

  if (document.body.scrollTop > 80 || document.documentElement.scrollTop > 80) {
      header.classList.add("scrolled");
  } else {
      header.classList.remove("scrolled");
  }
}

console.log("What")

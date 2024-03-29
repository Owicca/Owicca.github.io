let show_tab_class = "show";
let show_header_class = "active";

let tab_list = document.querySelectorAll(".tab_container .tab_content .tab");
let header = document.querySelector(".tab_container .tab_header");

function activate(target) {
  document.querySelector(".tab#"+target.id)
    .classList.add(show_tab_class);

  target.classList
    .add(show_header_class);
}

tab_list.forEach(elem => {
  let div = document.createElement("div");
  div.id = elem.id;
  div.innerHTML = elem.id.toUpperCase();

  header.appendChild(div);
});

// make first one active by default
activate(header.querySelector("div:first-of-type"));

let tab_header = document.querySelector(".tab_container .tab_header");
tab_header.addEventListener("click", e => {
  let clicked = e.target;

  clicked.parentElement
    .querySelectorAll("div")
    .forEach(elem => {
      elem.classList.remove(show_header_class)
    });

  tab_list.forEach(elem => {
    elem.classList.remove(show_tab_class);
  });

  activate(clicked);
});
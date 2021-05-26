var posts = document.getElementById("posts");
getPost();

function getPost() {
  posts.innerHTML = "";
  fetch("http://localhost:9090/posts")
    .then((response) => {
      return response.json();
    })
    .then((data) => {
      console.log(data);
      for (let i = 0; i < data.length; i++) {
        posts.innerHTML +=
          "\n<pre>" +
          data[i].body +
          "</pre>\n<small> email: " +
          data[i].email +
          "</small><hr>";
      }
    });
}

loginForm = document.getElementById("loginForm");
loginForm.addEventListener("submit", loginUser);
postForm = document.getElementById("form");
postForm.addEventListener("submit", makePost);

async function makePost(e) {
  e.preventDefault();
  var post = {
    body: document.getElementById("body").value,
    email: document.getElementById("email").value,
  };
  console.log(post);

  token = localStorage.getItem("token");
  console.log(token)

  await fetch("http://localhost:9090/posts", {
    mode: "no-cors",
    method: "POST",
    headers: { "Token": token },
    body: JSON.stringify(post),
  });

  document.getElementById("body").value = "";
  document.getElementById("email").value = "";
  getPost();
}

async function loginUser(e) {
  e.preventDefault();
  var user = {
    email: document.getElementById("emailLogin").value,
    password: document.getElementById("passLogin").value,
  };

  console.log(JSON.stringify(user));

  await fetch("http://localhost:9090/users/login", {
    mode: "cors",
    method: "POST",
    body: JSON.stringify(user),
  })
    .then((resp) => resp.text())
    .then((data) => localStorage.setItem("token", data));

  document.getElementById("emailLogin").value = "";
  document.getElementById("passLogin").value = "";
}

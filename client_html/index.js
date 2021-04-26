var posts = document.getElementById("posts");
getPost();

function getPost() {
  posts.innerHTML = "";
  fetch("http://localhost:9090/posts")
    .then(response => {
      return response.json();
    })
    .then(data => {
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

form = document.getElementById("form");
form.addEventListener("submit", makePost);

async function makePost(e) {
  e.preventDefault();
  var post = {
    body: document.getElementById("body").value,
    email: document.getElementById("email").value
  };
  console.log(post);

  await fetch("http://localhost:9090/posts", {
    mode: "no-cors",
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(post)
  });

  document.getElementById("body").value = "";
  document.getElementById("email").value = "";
  getPost();
}

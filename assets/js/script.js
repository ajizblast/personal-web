// data penting

let dataProject = [];
let i = 0;

function getData(event) {
  event.preventDefault();

  let nameProject = document.getElementById("nameProject").value;
  let startDate = document.getElementById("startDate").value;
  let endDate = document.getElementById("endDate").value;
  let description = document.getElementById("description").value;

  let nodeJs = document.getElementById("nodeJs").checked;
  let nextJs = document.getElementById("nextJs").checked;
  let reactJs = document.getElementById("reactJs").checked;
  let typeScript = document.getElementById("typeScript").checked;

  let fileImg = document.getElementById("fileImg").files[0];

  fileImg = URL.createObjectURL(fileImg);

  if (nameProject == "") {
    return alert("Tolong Lengkapi Form yg tersedia ya");
  } else if (startDate == "") {
    return alert("Tolong Lengkapi Form yg tersedia ya");
  } else if (endDate == "") {
    return alert("Tolong Lengkapi Form yg tersedia ya");
  } else if (description == "") {
    return alert("Tolong Lengkapi Form yg tersedia ya");
  } else if (fileImg == "") {
    return alert("Tolong Lengkapi Form yg tersedia ya");
  } else if (!nodeJs && !nextJs && !reactJs && !typeScript) {
    return alert("Mohon pilih tehnologi yang anda butuhkan ya");
  }
  let data = {
    nameProject,
    startDate,
    endDate,
    description,
    nodeJs,
    nextJs,
    reactJs,
    typeScript,
    fileImg,
  };

  dataProject.push(data);
  console.log(dataProject);

  showData();
}

function showData() {
  let containerResult = document.getElementById("containerResult");

  containerResult.innerHTML += `
        <div class="card-project">
            <img src="${dataProject[i].fileImg}" >
            <h2><a class="title-project" href="blog-project.html" target="_blank"
            >${dataProject[i].nameProject}</a></h2>
            <p class="durasi-project"> durasi : ${dataProject[i].startDate}<</p>
            <br>
            <p class="durasi-project"> durasi : ${dataProject[i].endDate}<</p>
            <p class="deskripsi-project">${dataProject[i].description}</p>
            <div class="icon-technology">
                <img id="iconNodeJs${i}" src="assets/images/node-js.png">
                <img id="iconReactJs${i}" src="assets/images/React-icon.svg.png">
                <img id="iconNextJs${i}" src="assets/images/next-js.png">
                <img id="iconTypeScript${i}" src="assets/images/typescript.png">
            </div>
            <div class="button-action">
                <button class="edit">edit</button>
                <button class="delete">delete</button>
            </div>
        </div>`;

  if (dataProject[i].nodeJs == false) {
    document.getElementById(`iconNodeJs${i}`).style.display = "none";
  }
  if (dataProject[i].nextJs == false) {
    document.getElementById(`iconNextJs${i}`).style.display = "none";
  }
  if (dataProject[i].reactJs == false) {
    document.getElementById(`iconReactJs${i}`).style.display = "none";
  }
  if (dataProject[i].typeScript == false) {
    document.getElementById(`iconTypeScript${i}`).style.display = "none";
  }

  i++;
}

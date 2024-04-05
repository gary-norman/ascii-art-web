const operation = document.getElementById("function")
const standardise = document.getElementById("standardise")
const colorize = document.getElementById("colorize")
const output = document.getElementById("output")
const align = document.getElementById("align")
const reverse = document.getElementById("reverse")
const color = document.getElementById("color")
const red = document.getElementById("red")
const green = document.getElementById("green")
const blue = document.getElementById("blue")
const yellow = document.getElementById("yellow")
const orange = document.getElementById("orange")
const standard = document.getElementById("standard")
const shadow = document.getElementById("shadow")
const thinkertoy = document.getElementById("thinkertoy")
const upload = document.getElementById("upload")
const alignment = document.getElementById("alignment")

 // drag and drop
const dropArea = document.querySelector("#drop_zone");
const dragText = document.querySelector(".dragText");
const dragButton = document.querySelector(".button");

let button = dropArea.querySelector(".button");
let input = dropArea.querySelector("input");
let file;
let filename
button.onclick = () => {
    input.click();
};

// when browse
input.addEventListener("change", function () {
    file = this.files[0];
    dropArea.classList.add("active");
    displayFile();
});

// when file is inside drag area
dropArea.addEventListener("dragover", (event) => {
    event.preventDefault();
    dropArea.classList.add("active");
    dragText.textContent = "Release to Upload";
    dragButton.style.display = "none";
    // console.log('File is inside the drag area');
});

// when file leave the drag area
dropArea.addEventListener("dragleave", () => {
    dropArea.classList.remove("active");
    // console.log('File left the drag area');
    dragText.textContent = "Drag your file here";
});

// when file is dropped
dropArea.addEventListener("drop", (event) => {
    event.preventDefault();
    dropArea.classList.add("dropped");
    // console.log('File is dropped in drag area');
    file = event.dataTransfer.files[0]; // grab single file even if user selects multiple files
    // console.log(file);
    displayFile();
});

function displayFile() {
    let fileType = file.type;
    // console.log(fileType);
    let validExtensions = ["text/plain"];
    if (validExtensions.includes(fileType)) {
    // console.log('This is an image file');
    let fileReader = new FileReader();
    fileReader.onload = () => {
        let fileURL = fileReader.result;
        // console.log(fileURL);
        dropArea.innerHTML = `<span class="dragText">uploaded ${file.name}</span>`;
    };
    fileReader.readAsDataURL(file);
} else {
    alert("This is not a Text File");
    dropArea.classList.remove("active");
}
}
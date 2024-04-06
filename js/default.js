// function selection and functions
const operation = document.getElementById("function")
const standardise = document.getElementById("standardise")
const colorize = document.getElementById("colorize")
const output = document.getElementById("output")
const align = document.getElementById("align")
const reverse = document.getElementById("reverse")
// selectors
const red = document.getElementById("red")
const green = document.getElementById("green")
const blue = document.getElementById("blue")
const yellow = document.getElementById("yellow")
const orange = document.getElementById("orange")
const txt_standard = document.getElementById("standard")
const txt_shadow = document.getElementById("shadow")
const txt_thinkertoy = document.getElementById("thinkertoy")
// ui sections
const color = document.getElementById("color")
const upload = document.getElementById("upload")
const alignment = document.getElementById("alignment")
const generator = document.getElementById("generator")
const asciiout = document.getElementById("asciiout")

// reset select and radios to default on page load
// TODO reset select and radios not functioning
document.addEventListener("DOMContentLoaded", reset);
function reset() {
    operation.selectedIndex = 0;
    // color.reset();
    operation.selectedIndex = 0;
    // standardise.reset();
    hideUI();
}
window.addEventListener("resize", hideUI);
// hide/unhide UI elements
operation.addEventListener("change", hideUI);
function hideUI() {
    // viewport size
    const viewport = window.innerWidth
    const value = operation.options[operation.selectedIndex].value;
    if (viewport <= 768) {
        if (value === "standard") {
            color.style.display = "none"
            upload.style.display = "none"
            alignment.style.display = "flex"
            generator.style.display = "flex"
            asciiout.style.display = "flex"
        }
        else if (value === "color") {
            color.style.display = "flex"
            upload.style.display = "none"
            alignment.style.display = "none"
            generator.style.display = "flex"
            asciiout.style.display = "flex"
        }
        else if (value === "output") {
            color.style.display = "none"
            upload.style.display = "none"
            alignment.style.display = "flex"
            generator.style.display = "flex"
            asciiout.style.display = "flex"
        }
        else {
            color.style.display = "none"
            upload.style.display = "flex"
            alignment.style.display = "none"
            generator.style.display = "none"
            asciiout.style.display = "none"
        }
    }
    else {
        color.style.display = "flex"
        upload.style.display = "flex"
        alignment.style.display = "flex"
        generator.style.display = "flex"
        asciiout.style.display = "flex"
    }
}

// drag and drop
// adapted from https://medium.com/@cwrworksite/drag-and-drop-file-upload-with-preview-using-javascript-cd85524e4a63
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
        dropArea.innerHTML = `<span class="dragText">uploaded ${file.name}</span>`;
    };
    fileReader.readAsDataURL(file);
} else {
    alert("This is not a Text File");
    dropArea.classList.remove("active");
    dragText.textContent = "Drag and drop your file, or";
    dragButton.style.display = "unset";
}
}
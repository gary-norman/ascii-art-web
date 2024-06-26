// options
const operation = document.getElementById("function");
const standardise = document.getElementById("standardise");
const colorize = document.getElementById("colorize");
const output = document.getElementById("output");
const align = document.getElementById("align");
const reverse = document.getElementById("reverse");
const darkmode = document.querySelector(".darkToggle");
const colorReset = document.querySelector(".whiteLabel");
// selectors
const red = document.getElementById("red");
const green = document.getElementById("green");
const blue = document.getElementById("blue");
const yellow = document.getElementById("yellow");
const orange = document.getElementById("orange");
const txt_standard = document.getElementById("standard");
const txt_shadow = document.getElementById("shadow");
const txt_thinkertoy = document.getElementById("thinkertoy");
const txt_right_radio = document.getElementById("radioright");
const txt_center_radio = document.getElementById("radiocenter");
const txt_justify_radio = document.getElementById("radiojustify");
const txt_left_input = document.querySelector("#left");
const txt_right_input = document.querySelector("#right");
const txt_center_input = document.querySelector("#center");
const txt_justify_input = document.querySelector("#justify");
const txt_right_image = document.querySelector(".iconRight");
const txt_center_image = document.querySelector(".iconCenter");
const txt_justify_image = document.querySelector(".iconJustify");
// ui sections
const color = document.getElementById("colorPicker");
const upload = document.getElementById("upload");
const alignment = document.getElementById("alignment");
const generator = document.getElementById("generator");
const asciiout = document.getElementById("asciiout");
const gopher = document.getElementById("gopher")
const uploadedFile = document.getElementById("uploadedFile")
// forms
const formMain = document.forms['mainForm'];
const textRadios = formMain.elements['text-align'];
const colorRadios = formMain.elements['colors'];
const generate = document.getElementById("generate")
const getText = document.getElementById("getText");
const getText2 = document.getElementById("getText2");
// buttons
const rev = document.getElementById("rev");
const genascii = document.getElementById("genascii")
const genasciiMob = document.getElementById("genasciiMob")
// text areas
const asciiOutput = document.getElementById("ascii-output");
// misc
let hasText = false
// functions
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
            genasciiMob.style.display = "block"
            asciiout.style.display = "flex"
        }
        else if (value === "color") {
            color.style.display = "flex"
            upload.style.display = "none"
            alignment.style.display = "none"
            generator.style.display = "flex"
            genasciiMob.style.display = "none"
            asciiout.style.display = "flex"
        }
        else if (value === "output") {
            color.style.display = "none"
            upload.style.display = "none"
            alignment.style.display = "flex"
            generator.style.display = "flex"
            genasciiMob.style.display = "block"
            asciiout.style.display = "flex"
        }
        else {
            color.style.display = "none"
            upload.style.display = "flex"
            alignment.style.display = "none"
            generator.style.display = "none"
            asciiout.style.display = "flex"
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
function reset() {
    operation.selectedIndex = 0;
    standardise.selectedIndex = 0;
    textRadios.value = 'left';
    resetColors();
    hideUI();
    hasText = false;
}
function resetColors() {
    for (let i = 0; i < colorRadios.length; i++ ) {
        colorRadios[i].checked = false;
        justifyDisable();
    }}
function justifyDisable() {
    let selected = false;
    for (let i = 0; i < colorRadios.length; i++ ) {
        if (colorRadios[i].checked === true) {
            selected = true
        }
    }
    if (selected === true) {
        txt_right_image.classList.add("checkmark2")
        txt_right_image.classList.remove("checkmark")
        txt_center_image.classList.add("checkmark2")
        txt_center_image.classList.remove("checkmark")
        txt_justify_image.classList.add("checkmark2")
        txt_justify_image.classList.remove("checkmark")
        txt_right_input.disabled = true
        txt_center_input.disabled = true
        txt_justify_input.disabled = true
        textRadios.value = 'left';
    }
    else {
        txt_right_image.classList.add("checkmark")
        txt_right_image.classList.remove("checkmark2")
        txt_center_image.classList.add("checkmark")
        txt_center_image.classList.remove("checkmark2")
        txt_justify_image.classList.add("checkmark")
        txt_justify_image.classList.remove("checkmark2")
        txt_right_input.disabled = false
        txt_center_input.disabled = false
        txt_justify_input.disabled = false
    }
}
function toggleReverse() {
    gopher.classList.toggle("gopherFlip");
    upload.classList.toggle("uploadFlip");
    rev.classList.toggle("revPressed")
}
function getTextFunc() {
    const viewport = window.innerWidth
    if (viewport <= 768) {
        getText2.classList.add("getText2Appear");
        getText2.classList.remove("getText2");
    }
    else {
        getText.classList.add("getTextAppear");
        getText.classList.remove("getText");
    }
}
function genasciiClick() {
    if (hasText === false) {
        genascii.disabled = false;
        genascii.classList.toggle("genasciiClickable");
        hasText = true;
    }
}
// ****** events ******
// make genascii clickable and highlighted
// generate.addEventListener("input", genasciiClick);
// show/hide upload
rev.addEventListener("click", toggleReverse);
gopher.addEventListener("click", toggleReverse);
// reset select and radios to default on page load
document.addEventListener("DOMContentLoaded", reset);
// disable align buttons if colour selected
color.addEventListener("change", justifyDisable);
//reset colours
colorReset.addEventListener("click", resetColors);
// hide UI elements if page resized
window.addEventListener("resize", hideUI);
// hide/unhide UI elements
operation.addEventListener("change", hideUI);
// set dark/light on output
darkmode.addEventListener("click", function () {
    asciiOutput.classList.toggle("asciiOutDark");
})

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
    let fileReader = new FileReader();
    fileReader.onload = () => {
        uploadedFile.innerHTML = `<div class="dragText">uploaded</div><div class="uploadedFile">${file.name}</div>`;
        dropArea.classList.add("dropped");
        genascii.disabled = true;
        genascii.classList.toggle("genasciiClickable");
        getTextFunc();
    };
    fileReader.readAsDataURL(file);
} else {
    alert("This is not a Text File");
    dropArea.classList.remove("active");
    dragText.textContent = "Drag and drop your file, or";
    dragButton.style.display = "unset";
}
}
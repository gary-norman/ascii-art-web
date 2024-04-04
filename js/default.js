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

 // drag and drop
function dropHandler(ev) {
    console.log("File(s) dropped");

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();

    if (ev.dataTransfer.items) {
        // Use DataTransferItemList interface to access the file(s)
        [...ev.dataTransfer.items].forEach((item, i) => {
            // If dropped items aren't files, reject them
            if (item.kind === "file") {
                const file = item.getAsFile();
                console.log(`… file[${i}].name = ${file.name}`);
            }
        });
    } else {
        // Use DataTransfer interface to access the file(s)
        [...ev.dataTransfer.files].forEach((file, i) => {
            console.log(`… file[${i}].name = ${file.name}`);
        });
    }
}
function dragOverHandler(ev) {
    console.log("File(s) in drop zone");

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
}

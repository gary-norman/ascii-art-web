function downloadFile() {
    let url = "http://localhost:8080/downloadFile"
    let options = {
        method: 'GET',
        headers: new Headers({
            'Content-Type': 'application/json',
        }),
        mode: 'cors',
        cache: 'default'
    };
    let strMimeType;
    let strFileName;

    //Perform a GET Request to server
    fetch(url, options)
        .then(function (response) {
            let contentType = response.headers.get("Content-Type"); //Get File name from content type
            strMimeType = "text/plain";
            strFileName = contentType.split(";")[1];
            return response.blob();

        }).then(function (myBlob) {
        let downloadLink = window.document.createElement('a');
        downloadLink.href = window.URL.createObjectURL(new Blob([myBlob], { type: strMimeType }));
        downloadLink.download = strFileName;
        document.body.appendChild(downloadLink);
        downloadLink.click();
        document.body.removeChild(downloadLink);
    }).catch((e) => {
        console.log("e", e);
    });
}
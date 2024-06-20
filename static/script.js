document.addEventListener("DOMContentLoaded", function() {
    const fileViewer = document.getElementById("file-viewer");

    document.querySelectorAll(".file-link").forEach(anchor => {
        anchor.addEventListener("click", function(event) {
            if (!this.target.includes("_blank")) {
                event.preventDefault();
                const filePath = this.getAttribute("href");

                if (filePath.endsWith(".mp4") || filePath.endsWith(".flv") || filePath.endsWith(".mkv")) {
                    fileViewer.innerHTML = `<video controls><source src="${filePath}" type="video/mp4">Your browser does not support the video tag.</video>`;
                } else {
                    window.location.href = filePath;
                }
            }
        });
    });
});

document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".file-link").forEach(anchor => {
        anchor.addEventListener("click", function(event) {
            if (!this.target.includes("_blank")) {
                event.preventDefault();
                const filePath = this.getAttribute("href");

                if (filePath.match(/\.(mp4|flv|mkv)$/)) {
                    window.open(filePath, '_blank');
                } else {
                    window.location.href = filePath;
                }
            }
        });
    });
});

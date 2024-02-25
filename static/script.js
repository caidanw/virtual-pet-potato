// Add JavaScript code here to handle button interactions and potato animations

const waterButton = document.getElementById("water-button");
waterButton.onclick = (event) => {
    document.body.classList.add("rain");
    setTimeout(() => {
        document.body.classList.remove("rain");
    }, 2000);
};

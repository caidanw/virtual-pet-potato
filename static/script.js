function triggerRain(color = null) {
    if (document.body.classList.contains("rain")) return;

    if (color !== null) {
        document.body.style.cssText = `--c: ${color};`;
    }

    document.body.classList.add("rain");
    setTimeout(() => {
        document.body.classList.remove("rain");
        document.body.style.cssText = "";
    }, 2000);
}

const waterButton = document.getElementById("water-button");
waterButton.addEventListener("click", () => triggerRain());

const fertilizeButton = document.getElementById("fertilize-button");
fertilizeButton.addEventListener("click", () => triggerRain("brown"));

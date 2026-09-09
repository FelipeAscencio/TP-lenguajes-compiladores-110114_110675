const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    const btn = document.getElementById("runBtn");
    btn.disabled = false;
    btn.innerText = "Ejecutar Código";
    btn.addEventListener("click", () => {
        const code = document.getElementById("code").value;
        const output = window.runLox(code); 
        document.getElementById("output").innerText = output;
    });
}).catch((err) => {
    console.error("Error cargando Wasm:", err);
});

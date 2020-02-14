if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("frontend.wasm"), go.importObject).then(
    async result => {
        mod = result.module;
        inst = result.instance;
        await go.run(inst);
    }
 );

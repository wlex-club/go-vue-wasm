let a ={
    module:{
        rules:[
            ...
            {
                test:/\.wasm$/,
                loaders: ['wasm-loader']
            }
        ]
    }
}
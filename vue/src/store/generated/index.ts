// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Shifty11TictactoeTictactoe from './shifty11.tictactoe.tictactoe'


export default { 
  Shifty11TictactoeTictactoe: load(Shifty11TictactoeTictactoe, 'shifty11.tictactoe.tictactoe'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
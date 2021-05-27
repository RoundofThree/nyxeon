// this doesn't work in VueJS 3 

import Vuetify from "vuetify" 
import "vuetify/dist/vuetify.min.css" 

const opts = {
    theme: {
        dark: false
    }, 
    options: {
        customProperties: true 
    },
    icons: {
        iconfont: "mdi"
    }
}

export default new Vuetify(opts)
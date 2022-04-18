import Vue from "vue";
import Vuex from "vuex";

import todoModule from "./modules/todo/todo"
import authentication from "./modules/auth/auth"


Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        todo: todoModule,
        auth: authentication,
    },
    state:{
        appTitle: process.env.VUE_APP_TITLE,
    }



});

import Vue from 'vue'
import VueRouter from 'vue-router'
import ToDo from '../views/ToDo.vue'
import BoardForCalendar from '../views/BoardForCalendar.vue'
import Registration from "@/views/RegistrationPage";
import LoginPage from "@/views/LoginPage";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'ToDo',
        component: ToDo
    },
    {
        path: '/calendar',
        name: 'BoardForCalendar',
        component: BoardForCalendar
    },
    {
        path: '/registration',
        name: 'Registration',
        component: Registration

    },
    {
        path: '/login',
        name: 'LoginPage',
        component: LoginPage

    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, from, next) => {
    document.title = `${process.env.VUE_APP_TITLE} - ${to.name}`
    next()
})

export default router

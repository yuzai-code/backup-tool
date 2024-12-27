import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

const routers = [
    {
        path: "/",
        name: "home",
        component: Home
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes: routers
});

export default router;

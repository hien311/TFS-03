import Vue from "vue";
import VueRouter from "vue-router";
import store from "../store/index"
import { CHECK_AUTHENTICATE } from "../store/actions.type";

Vue.use(VueRouter);

const routes = [
  {
    path: "/home",
    name: "home",
    component: () => import("../views/Home.vue"),
    children: [
      {
        path: "/list-articles/:type",
        name: "list-articles",
        component: () => import("../views/ListArticles.vue")
      }
    ],
    exact: true
  },
  {
    path: "/login",
    name: "login",
    component: () => import("../views/Login.vue"),
  },
  {
    path: "/register",
    name: "register",
    component: () => import("../views/Register.vue"),
  },
  {
    path: "/new-article",
    name: "new-article",
    component: () => import("../views/NewArticle.vue")
  },
  {
    path: "/article/:slug",
    name: "article",
    component: () => import("../views/Article.vue")
  },
  {
    path: "/setting",
    name: "setting",
    component: () => import("../views/Setting.vue")
  },
  {
    path: "/user",
    name: "user",
    component: () => import("../views/User.vue"),
    children: [
      { 
      path: "/list/:type", 
      name: 'user-articles-list', 
      component: () => import("../views/UserArticlesList.vue"
      )}
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to,from, next) => {
  if (to.path !== "/login" && to.path !== "/register") {
    store.dispatch(CHECK_AUTHENTICATE)
  }
  if (to.path=="/") {
    router.push('home')
  }
  next()
})

export default router;
